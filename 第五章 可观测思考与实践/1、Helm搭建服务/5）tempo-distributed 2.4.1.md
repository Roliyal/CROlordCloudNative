

# 1、Tempo 简介

Grafana Tempo 是由 Grafana Labs 开源的分布式追踪系统，它作为可观测性栈的一部分，专注于提供高效、可扩展的追踪数据收集、存储和查询功能。Tempo 设计用于处理大规模的分布式应用程序中的追踪数据，帮助开发者理解服务间的调用关系、性能瓶颈和错误传播路径，从而提高系统的可靠性和性能。

Grafana Tempo 简介：

核心特性：
高性能存储： Tempo 使用对象存储（如 AWS S3、GCP GCS 或 Azure Blob Storage）来低成本、长时间地存储追踪数据，同时利用本地缓存提高查询速度。
高效压缩： 通过高效的块压缩和索引技术，大幅减少追踪数据的存储空间需求。
简易集成： 支持 OpenTelemetry、Jaeger、Zipkin 等多种追踪格式，便于与现有系统集成。
查询优化： 提供快速查询接口，即使是数 TB 的数据也能迅速检索。
与 Grafana 生态集成： 无缝集成 Grafana 数据可视化平台，便于创建追踪数据的可视化面板和告警。

常用场景：
微服务监控： 在复杂的微服务架构中，Tempo 能够帮助追踪服务间的调用链路，识别慢请求和故障点，这对于理解服务间依赖、优化服务性能至关重要。
性能诊断： 当系统响应变慢或出现故障时，开发和运维团队可以利用 Tempo 追踪事务，快速定位问题发生的源头，无论是数据库查询延迟、网络延迟还是代码逻辑问题。
故障排查： 在分布式系统中，故障可能涉及多个服务和层次，Tempo 提供了一种手段来追溯整个调用链，加速故障排查过程。
服务优化： 通过对追踪数据的持续分析，识别性能瓶颈，为服务优化和架构调整提供数据支持。
用户体验监控： 结合前端追踪数据，可以完整地重建用户会话的追踪记录，了解用户在使用应用过程中的实际体验，辅助提升产品体验。
安全审计： 分析追踪数据还能帮助识别潜在的安全威胁，如异常的访问模式或未经授权的访问尝试。
结合 Grafana 的可视化能力，Tempo 不仅能提供追踪数据的详细视图，还能通过丰富的图表和报警功能，将追踪洞察转化为可操作的智能，助力 DevOps 团队更好地管理现代云原生应用的复杂性。


# 2、Github地址
[https://github.com/grafana/tempo](https://github.com/grafana/tempo)

# 3、Helm安装

> Helm是一个用于Kubernetes的应用程序包管理工具，它简化了Kubernetes应用的部署和管理过程。Helm的工作原理类似于Linux中的apt或者yum，但专门针对Kubernetes资源


## 3.1 添加Helm源

```bash
helm repo add grafana  https://grafana.github.io/helm-charts
```

## 3.2 下载chart包

```bash
helm pull grafana/tempo-distributed --version 1.9.9
tar -zxvf tempo-distributed-1.9.9.tgz
```
## 3.3 安装
### 3.3.1 修改配置
配置文件`vim  tempo-distributedvalues.yaml`

```bash
 ...
  16   dnsService: 'coredns'  # 默认是kube-dns（Kubernetes dns服务器）
 ...
 229 metricsGenerator:
 230   # -- Specifies whether a metrics-generator should be deployed
 231   enabled: true           #启用metrics-generator，可以生产service graph
 ...
 324   ports:
 325     - name: grpc
 326       port: 9095
 327       service: true
 328     - name: http-memberlist
 329       port: 7946
 330       service: true
 331     - name: http-metrics
 332       port: 3100
 333       service: true
 ...
 361       remote_write:
 362       - url: http://prometheus-server.monitoring.svc.cluster.local/api/v1/write  #远程写入prometheus remote_write API地址
 ...
 975   otlp:   #这个眼熟不，记住吧，这个以后将非常强大，是一种协议规范，可以使用opentelemetry把数据写入进来，然后使用grafana展示
 976     http:
 977       # -- Enable Tempo to ingest Open Telemetry HTTP traces
 978       enabled: true
 979       # -- HTTP receiver advanced config
 980       receiverConfig: {}
 981     grpc:
 982       # -- Enable Tempo to ingest Open Telemetry GRPC traces
 983       enabled: true
 ...
1366 global_overrides:
1367   per_tenant_override_config: /runtime-config/overrides.yaml
1368   metrics_generator_processors:    #新增该配置（我没有找到单独的overrides.yml怎么配置，干脆直接在这里写了）
1369   - 'service-graphs'
1370   - 'span-metrics'
 ...
1585 minio:
1586   enabled: true
...
1603   persistence:
1604     size: 500Gi 
...
1613 gateway:
1614   # -- Specifies whether the gateway should be enabled
1615   enabled: true
```
### 3.3.2 安装

```bash
helm install tempo-distributed tempo-distributed/ -n trace
```


常见服务集成配置连接：
grafana datasource：`http://tempo-distributed-query-frontend-discovery.trace.svc.cluster.local:3100`

opentelemetry：`tempo-distributed-distributor-discovery.trace.svc.cluster.local:4317`
## 3.4 组件介绍
**Tempo Server**: 这是 Grafana Tempo 的核心组件，负责接收、处理和存储追踪数据。它支持从各种追踪源（如 Jaeger、OpenTelemetry 等）接收数据，并将数据高效地存储到后端存储系统中（通常是对象存储服务，如 AWS S3、Google Cloud Storage 或 Azure Blob Storage）。Tempo Server 还提供了查询 API，用于检索存储的追踪数据。

**Distributor**: 负责接收追踪数据并对其进行分配，实现负载均衡。它确保高并发写入时的服务稳定性。

**Ingester**: 处理从 Distributor 接收到的数据，对追踪数据进行压缩、归并和临时存储，之后将处理后的数据持久化到长期存储。

**Querier:** 负责处理追踪数据查询请求。当需要检索追踪数据时，它从对象存储中读取数据，并返回给用户或前端应用（如 Grafana）。

**Compactor**: 定期合并存储在对象存储中的小数据块，以优化存储效率和查询性能。

**Gateway**/Query Frontend: 这个组件作为查询请求的入口，负责路由查询请求到合适的 Querier，提供额外的安全控制或负载均衡。
