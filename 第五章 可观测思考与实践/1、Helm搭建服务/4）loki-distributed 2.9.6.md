
# 1、Loki简介

Grafana Loki 是一个开源的云原生日志聚合和分析系统，由 Grafana Labs 开发并维护。Loki 专注于为大规模的日志处理提供经济高效且易于管理的解决方案，尤其适用于微服务架构以及容器化和分布式环境。

以下是 Loki 的核心特性与设计原则：

1. **标签驱动存储**：
   - Loki 不对完整的日志内容进行索引，而是仅对每个日志流（log stream）定义的一组标签（tags）进行索引，类似于 Prometheus 对时间序列数据的处理方式。这种方式显著减少了存储开销，特别是对于包含大量相似或重复文本的日志数据。

2. **水平扩展**：
   - Loki 被设计成可以水平扩展以处理海量日志数据，通过多个工作节点共同分担存储和查询压力。

3. **组件结构**：
   - **Promtail**：是 Loki 生态系统的一部分，它是一个用于从服务器收集日志并将它们发送到 Loki 的日志代理。
   - **Distributor**：接收来自客户端（如 Promtail 或其他日志推送器）的日志数据，验证、分批，并将它们转发给 Ingester 进行进一步处理。
   - **Ingester**：负责接收经过 Distributor 分发来的日志块，将其持久化存储在对象存储中，并在日志被索引后删除临时数据。
   - **Querier**：响应用户发起的日志查询请求，从存储中检索相关日志数据。
   - **Grafana UI集成**：Loki 可以无缝集成到 Grafana 中，为用户提供可视化界面来搜索、过滤、分析和展示日志信息。

4. **经济高效**：
   - 由于其独特的索引策略和使用低成本的对象存储（例如 Amazon S3、MinIO 等），Loki 可以大幅减少存储成本。

5. **多租户支持**：
   - Loki 提供了多租户功能，允许不同团队或者项目在同一集群上独立管理和查看各自日志。

综上所述，Grafana Loki 旨在简化日志管理，并通过其创新的设计，为企业提供了在现代云环境中处理日志的一种高效且经济的方式。

# 2、Github
[https://github.com/grafana/loki](https://github.com/grafana/loki)


# 3、Helm安装

> Helm是一个用于Kubernetes的应用程序包管理工具，它简化了Kubernetes应用的部署和管理过程。Helm的工作原理类似于Linux中的apt或者yum，但专门针对Kubernetes资源

## 3.1 添加Helm源

```bash
helm repo add grafana  https://grafana.github.io/helm-charts
```

## 3.2 下载chart包

```bash
helm pull grafana/loki-distributed --version 0.78.5
tar zxvf loki-distributed-0.78.5.tgz 
```


vim loki-distributed/values.yaml
```bash
global:
  image:
    # -- Overrides the Docker registry globally for all images
    registry: null
  # -- Overrides the priorityClassName for all pods
  priorityClassName: null
  # -- configures cluster domain ("cluster.local" by default)
  clusterDomain: "cluster.local"
  # -- configures DNS service name
  dnsService: "coredns"       #注意下这个位置，要修改成自己集群的dns值，默认值是kube-dns
  # -- configures DNS service namespace
  dnsNamespace: "kube-system"
  ...
  ...
  ...
```
## 3.3 安装

```bash
helm upgrade --install  loki-distributed  loki-distributed/  -n logs
```

> 我把 sts/ingester服务和sts/querier服务持久化到了一个pvc，容器内挂载位置都是/var/loki，主要是解决 Failed to load chunks 问题https://github.com/grafana/helm-charts/issues/1111


## 3.4 修改配置
configmap/loki-distributed 

```bash
auth_enabled: false
chunk_store_config:
  max_look_back_period: 0s
common:
  compactor_address: http://loki-distributed-compactor:3100
compactor:
  shared_store: filesystem
distributor:
  ring:
    kvstore:
      store: memberlist
frontend:
  compress_responses: true
  log_queries_longer_than: 5s
  tail_proxy_url: http://loki-distributed-querier:3100
frontend_worker:
  frontend_address: loki-distributed-query-frontend-headless:9095
ingester:
  chunk_block_size: 262144
  chunk_encoding: snappy
  chunk_idle_period: 30m
  chunk_retain_period: 1m
  lifecycler:
    ring:
      kvstore:
        store: memberlist
      replication_factor: 1
  max_transfer_retries: 0
  wal:
    dir: /var/loki/wal
ingester_client:
  grpc_client_config:
    grpc_compression: gzip
limits_config:
  enforce_metric_name: false
  max_cache_freshness_per_query: 10m
  reject_old_samples: true
  reject_old_samples_max_age: 168h
  split_queries_by_interval: 15m
memberlist:
  join_members:
  - loki-distributed-memberlist
query_range:
  align_queries_with_step: true
  cache_results: true
  max_retries: 5
  results_cache:
    cache:
      embedded_cache:
        enabled: true
        ttl: 24h
ruler:
  alertmanager_url: https://alertmanager.xx
  external_url: https://alertmanager.xx
  ring:
    kvstore:
      store: memberlist
  rule_path: /tmp/loki/scratch
  storage:
    local:
      directory: /etc/loki/rules
    type: local
runtime_config:
  file: /var/loki-distributed-runtime/runtime.yaml
schema_config:
  configs:
  - from: "2020-09-07"
    index:
      period: 24h
      prefix: loki_index_
    object_store: filesystem
    schema: v11
    store: boltdb-shipper
server:
  http_listen_port: 3100
  grpc_server_max_recv_msg_size: 1048576000
  grpc_server_max_send_msg_size: 1048576000
storage_config:
  boltdb_shipper:
    active_index_directory: /var/loki/index
    cache_location: /var/loki/cache
    cache_ttl: 168h
    shared_store: filesystem
  filesystem:
    directory: /var/loki/chunks
table_manager:
  retention_deletes_enabled: false
  retention_period: 0s
```
最后重启loki，现在我们就有了一个分布式的Loki服务
