

# 1、Grafana简介
Grafana是一款流行的开源数据可视化和分析平台，它特别擅长处理时序数据，即随时间变化的数据，因此在监控系统性能、应用程序指标、物联网（IoT）设备数据等领域有着广泛的应用

# 2、常用场景
Kubernetes监控集成：
Grafana通常与Prometheus协同工作，后者是专为云原生和微服务架构设计的开源监控系统。在K8s集群中，Prometheus负责从节点、Pods、服务以及其它组件收集各类资源指标，如CPU使用率、内存消耗、网络I/O、磁盘空间等。这些数据随后被Grafana用于创建实时、动态的可视化仪表板，帮助运维团队直观地监控集群健康状况、资源分配和应用性能。

资源优化与故障排查：
通过Grafana，Kubernetes管理员可以快速识别资源瓶颈、异常行为和潜在故障点。定制化的仪表板能够展示集群内各个层面的详细信息，从整体负载到单个容器的性能，为资源规划和优化提供数据支持。

告警与通知：
集成Prometheus的告警规则后，Grafana能在检测到预定义条件（如资源使用超标或服务不可达）时触发告警。这些告警可以通过电子邮件、Slack、PagerDuty等多种渠道发送，确保团队能及时响应并采取措施。

日志分析集成：
与Loki等日志聚合工具结合，Grafana还能实现对Kubernetes集群日志的可视化分析。Loki专门设计用于处理大规模日志数据，将日志与指标数据在同一个界面下关联展示，极大提升了故障诊断的效率。

可扩展性和易用性：
Grafana的界面友好且高度可配置，支持通过其丰富的API和插件生态进行扩展。针对Kubernetes，Grafana有专门的应用和仪表板模板，能够快速部署并开始监控，降低了入门门槛。

总之，在Kubernetes监控体系中，Grafana不仅是数据可视化的前端展示工具，更是连接数据源、告警管理和决策制定的关键桥梁，对于保障K8s集群稳定运行和优化资源利用至关重要。

# 3、Github
[https://github.com/grafana/grafana](https://github.com/grafana/grafana)

# 4、Helm安装

> Helm是一个用于Kubernetes的应用程序包管理工具，它简化了Kubernetes应用的部署和管理过程。Helm的工作原理类似于Linux中的apt或者yum，但专门针对Kubernetes资源

## 4.1 添加Helm源

```bash
helm repo add grafana  https://grafana.github.io/helm-charts
```
## 4.2 下载chart包

```bash
helm pull loki/grafana --version 6.56.2
tar -zxvf grafana-6.56.2.tgz
```
## 4.3 修改配置
vim grafana/values.yaml
 

```bash
----------
#开启数据持久化
persistence:
  type: pvc
  enabled: true
  # storageClassName: default
  accessModes:
    - ReadWriteMany
  size: 100Gi
...
...
...
#设置密码
# Administrator credentials when not using an existing secret (see below)
adminUser: admin
adminPassword: W7xelEv6i5lkFSEn58WtdUuEgJkv46Bxxxxxxxxx
```
## 4.4 安装

```bash
helm install grafana95 grafana/ -n monitoring
```
集群内部访问地址：http://grafana95.monitoring.svc.cluster.local

```bash
[root@ops grafana]# kubectl get pods,pvc,configmap -n monitoring  |grep grafana |awk '{print $1,$2}'
pod/grafana95-864c75b994-q79r2 1/1
persistentvolumeclaim/grafana95 Bound
configmap/grafana95 1
```
