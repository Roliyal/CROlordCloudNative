# 1、Promtail 简介
Promtail 是一款由 Grafana Labs 开发的日志收集代理软件，设计用于与 Loki 日志聚合系统紧密协作。Loki 是一个高度可扩展、高可用性且支持多租户的日志管理系统，其设计理念受到了 Prometheus 监控系统的启发，特别注重效率和易用性。不同于传统日志管理系统，Loki 不对日志内容进行全文索引，而是依赖标签（labels）来索引和查询日志，从而实现了资源的有效利用。

Promtail 的核心职责如下：

日志采集：它负责从各种数据源（如应用程序日志文件、标准输出、容器日志等）收集日志数据。
结构化日志：Promtail 可以自动或通过配置添加标签到日志条目中，这些标签用于在 Loki 中高效地索引和查询日志。
日志传输：收集到的日志数据会被 Promtail 格式化并发送到 Loki 服务器，期间可能采用批处理和压缩等优化措施以减少网络传输开销和提升效率。
灵活性与配置：Promtail 通常通过一个名为 config.yaml 的配置文件进行配置，该文件允许用户定义日志源、目标 Loki 服务器的地址、标签规则以及其他高级设置。配置支持环境变量引用，可以通过命令行参数动态调整。
综上所述，Promtail 作为日志收集的前端组件，在现代化日志管理架构中扮演着关键角色，特别是与 Grafana 和 Loki 配合使用时，能够提供强大的日志可视化、分析和故障排查能力


# 2、Github地址
[https://github.com/grafana/puppet-promtail](https://github.com/grafana/puppet-promtail)

# 3、Helm安装

> Helm是一个用于Kubernetes的应用程序包管理工具，它简化了Kubernetes应用的部署和管理过程。Helm的工作原理类似于Linux中的apt或者yum，但专门针对Kubernetes资源


## 3.1 添加Helm源

```bash
helm repo add grafana  https://grafana.github.io/helm-charts
```

## 3.2 下载chart包

```bash
helm pull grafana/promtail --version 6.15.3
tar -zxvf promtail-6.15.3.tgz 
```
## 3.3 安装

```bash
helm install promtail promtail/ -n logs
```
## 3.4 修改配置

```bash
server:
  log_level: info
  log_format: logfmt
  http_listen_port: 3101
  

clients:
  - url: http://loki-distributed-gateway/loki/api/v1/push

positions:
  filename: /run/promtail/positions.yaml

scrape_configs:
  # See also https://github.com/grafana/loki/blob/master/production/ksonnet/promtail/scrape_config.libsonnet for reference
  - job_name: kubernetes-pods
    pipeline_stages:
      - match:
          selector: '{namespace=~"dev|test|uat"}'  #只有当namespace的值是test或uat时执行该匹配
          stages:
          - multiline:
              firstline: '^*\[\w+-\w+-\w+:\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}:\d{5}\]' #多行合并
          - regex:
              expression: .*(?P<level>INFO|WARN|ERROR).*   #level的值来自查找日志中的INFO、WARN、ERROR
          - timestamp:
              format: RFC3339Nano
              source: timestamp
          - labels:
              level: 
              timestamp:
      - match:
          selector: '{namespace=~"dev|test|uat"}'  #只有当namespace的值是test或uat时执行该匹配
          stages:
          - multiline:
              firstline: '^*\[\w+-\w+-\w+:\d{1,3}.\d{1,3}.\d{1,3}.\d{1,3}:\d{5}\]'
          - regex:
              expression: .*(?P<apserror>ERROR.*TID:.{300}).*   #匹配错误字段样式，最多显示300个字符：ERROR 1[TID:70019cdf8aa1416f8e02326adfc32a43.399207.17121108001200001][uId:][sId:][tId:][KafkaC...
          - timestamp:
              format: RFC3339Nano
              source: timestamp
          - labels:
              level: 
              apserror: 
              timestamp:
      - cri: {}
    kubernetes_sd_configs:
      - role: pod
    relabel_configs:
      - source_labels:
          - __meta_kubernetes_pod_controller_name
        regex: ([0-9a-z-.]+?)(-[0-9a-f]{8,10})?
        action: replace
        target_label: __tmp_controller_name
      - source_labels:
          - __meta_kubernetes_pod_label_app_kubernetes_io_component
          - __meta_kubernetes_pod_label_component
        regex: ^;*([^;]+)(;.*)?$
        action: replace
        target_label: component
      - action: replace
        source_labels:
        - __meta_kubernetes_pod_node_name
        target_label: node_name
      - action: replace
        source_labels:
        - __meta_kubernetes_namespace
        target_label: namespace
      - action: replace
        source_labels:
        - __meta_kubernetes_pod_name
        target_label: pod
      - action: replace
        source_labels:
        - __meta_kubernetes_pod_container_name
        target_label: container
      - action: replace
        regex: (dev|test|uat)/(.*)
        replacement: /var/log/pods/$1*$2*/$2/*.log          #只匹配指定命名空间的日志
        separator: /
        source_labels:
        - __meta_kubernetes_namespace
        - __meta_kubernetes_pod_container_name
        target_label: __path__
      - action: replace
        regex: true/(dev|test|uat)/(.*)
        replacement: /var/log/pods/$1*$2*/$2/*.log           #只匹配指定命名空间的日志
        separator: /
        source_labels:
        - __meta_kubernetes_pod_annotationpresent_kubernetes_io_config_hash
        - __meta_kubernetes_pod_annotation_kubernetes_io_config_hash
        - __meta_kubernetes_namespace
        - __meta_kubernetes_pod_container_name
        target_label: __path__
  
  

limits_config:
  

tracing:
  enabled: false

```
总结：
1、Loki服务端URL，我这里的是http://loki-distributed-gateway/loki/api/v1/push，需要修改成自己的，loki分布式安装地址：[https://dongweizhen.blog.csdn.net/article/details/138799616](https://dongweizhen.blog.csdn.net/article/details/138799616)
2、多行匹配详情可以参考：[https://dongweizhen.blog.csdn.net/article/details/135284386](https://dongweizhen.blog.csdn.net/article/details/135284386)
3、日志匹配相关可以参考：[https://blog.ossq.cn/2674.html](https://blog.ossq.cn/2674.html)
⚠4、这里只加了2个标签，加多了没法保证日志实时读取（我调过性能参数，很遗憾并没有解决）
