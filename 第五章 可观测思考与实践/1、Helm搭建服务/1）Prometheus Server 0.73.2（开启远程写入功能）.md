# 1、Prometheus 简介


Prometheus 是一款开源的监控和警报系统，起源于 SoundCloud 并在2012年开始被广泛采用。它在2016年成为 Cloud Native Computing Foundation (CNCF) 的成员项目，与 Kubernetes 齐名，是云原生计算领域的重要组成部分。Prometheus 以其强大的灵活性、易用性和云原生友好特性，在现代IT基础设施监控中占据重要地位。

主要特点包括：
多维数据模型：Prometheus 使用基于度量名称和一组键/值对（标签）定义的时间序列数据模型，这种模型支持灵活的数据组织和查询。

PromQL查询语言：提供了一种强大而灵活的查询语言PromQL，允许用户执行复杂的数据查询、聚合和计算，以便进行数据分析和可视化。

数据收集机制：Prometheus 采用基于HTTP的拉取（Pull）模型来周期性地从目标处收集数据，同时也支持通过Pushgateway组件接收短期作业或批处理任务的数据推送。

服务发现与配置：支持自动服务发现机制，能够自动检测新的目标或服务，同时也允许静态配置监控目标。

本地存储：内置了一个高性能的时间序列数据库（TSDB），用于存储和检索监控数据，支持数据的持久化、备份与恢复。

告警管理：Prometheus Server负责生成告警，而告警的具体处理（如去重、分组、路由等）由独立的组件Alertmanager完成，提供高度可定制的告警策略和通知渠道。

生态系统丰富：Prometheus拥有广泛的出口商（Exporters），可以轻松集成监控各种服务和系统（如数据库、HTTP服务器、消息队列等），并与Grafana等可视化工具紧密集成，便于构建丰富的仪表板。

应用场景：
微服务监控：由于其轻量级、可伸缩的特性，Prometheus 特别适合微服务架构的监控，可以高效地跟踪服务间的依赖关系和性能指标。

容器与Kubernetes监控：Prometheus 成为Kubernetes生态系统的标准监控解决方案之一，能够无缝集成，监控容器的资源使用、健康状况和生命周期事件。

云基础设施监控：无论是公有云、私有云还是混合云环境，Prometheus都能提供全面的基础设施监控，包括服务器、网络设备和存储系统的性能监控。

应用程序性能监控（APM）：结合特定的Exporter和自定义指标，Prometheus可用于深入应用程序内部，监控事务延迟、错误率和资源消耗。

DevOps与运维：通过实时监控和即时告警，Prometheus帮助DevOps团队及时发现并解决问题，提高服务的稳定性和可用性。

综上所述，Prometheus凭借其强大功能和灵活的生态系统，成为了现代IT监控领域的首选工具之一，广泛适用于各种规模和复杂度的系统监控场景。


# 2、Github地址
[https://github.com/prometheus/prometheus](https://github.com/prometheus/prometheus)

# 3、Helm安装

> Helm是一个用于Kubernetes的应用程序包管理工具，它简化了Kubernetes应用的部署和管理过程。Helm的工作原理类似于Linux中的apt或者yum，但专门针对Kubernetes资源


## 3.1 添加Helm源

```bash
helm repo add grafana  https://grafana.github.io/helm-charts
```

## 3.2 下载chart包

```bash
helm pull prometheus-community/prometheus --version 25.21.0
tar -zxvf prometheus-25.21.0.tgz 
```
## 3.3 安装
### 3.3.1 开启远程可写入
建议开启这个功能，如果后续想使用grafana出的链路追踪工具tempo实现`Service Graph`，需要把数据推送到Prometheus，即写入接口。我尝试过使用`Prometheus Pushgateway`服务的方式，并没有成功
![在这里插入图片描述](https://img-blog.csdnimg.cn/direct/f3ccea0198c94a62a5ae7d944296da76.png)

![在这里插入图片描述](https://img-blog.csdnimg.cn/direct/969e71b02d444a54a7f71ec508c54240.png)
注⚠：使用Grafana Tempo实现的链路追踪图示，想要实现上图功能，欢迎订阅专栏：[APM领域](https://blog.csdn.net/zhanremo3062/category_12552674.html)


修改`vim prometheus/values.yaml`文件：

```bash
.
.
.
 251   remoteWrite:
 252   - url: http://prometheus-server.monitoring.svc.cluster.local/api/v1/write  # 暴露远程写入API
 .
 .
 .
 278   extraArgs:
 279     web.enable-remote-write-receiver: null   #开启远程写入功能
 .
 .
 .

```
### 3.3.2 安装

```bash
helm install prometheus prometheus/ -n monitoring
```


常见服务集成配置连接：
grafana datasource：http://prometheus-server.monitoring.svc.cluster.local
tempo remote_write：http://prometheus-server.monitoring.svc.cluster.local/api/v1/write
## 3.4 组件介绍

**Prometheus Server**: 负责从各个exporter收集时间序列数据，并存储这些数据以便后续查询和分析。它是整个监控系统的核心部分，负责数据的抓取、存储、查询和报警触发的基础数据处理。

**Alertmanager**: 负责处理由Prometheus Server生成的警报，包括去重、分组、路由警报到正确的接收者，并支持多种通知方式（如邮件、短信、聊天工具等）。它确保警报策略的执行，并且具有高可用性设计。

**Node Exporter**: 一个轻量级的守护进程，用于收集宿主机上的硬件和操作系统指标（如CPU负载、内存使用情况、磁盘I/O、网络流量等），并通过HTTP端点暴露这些数据给Prometheus Server抓取。这是监控Kubernetes节点的基本组件。

**Pushgateway**: 作为一个临时的中继站，允许短期作业或批处理任务推送它们的指标到Pushgateway，然后再由Prometheus Server拉取。这适用于不能被直接拉取数据的服务或作业
