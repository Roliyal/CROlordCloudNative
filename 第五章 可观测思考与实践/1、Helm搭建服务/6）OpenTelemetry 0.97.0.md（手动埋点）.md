>手动埋点会让我们明白其中的逻辑关系，我个人并不排程自动埋点，关于可观测的路还很长，并不是所有组件都能做到自动适配，以防牵一发而动全身，建议大家不要舍远求近~新手使用手动埋点是最佳选择

# 1、OpenTelemetry 简介


OpenTelemetry 是一个开源的、全面的可观测性框架，旨在为云原生及传统应用提供标准化的方式来生成、收集、处理和导出遥测数据，包括跟踪（Traces）、度量（Metrics）和日志（Logs）。这个项目由 Cloud Native Computing Foundation (CNCF) 孵化，目标是简化和统一观测性数据的收集与分析，从而提高开发人员在分布式系统中调试和监控应用的能力。

核心特点包括：

与供应商无关：OpenTelemetry 提供了一套标准的 API 和 SDK，使得开发者能够采用统一的方式集成观测性功能，而不用绑定到特定的供应商或工具。
多语言支持：支持多种编程语言，确保跨语言的应用程序可以有一致的可观测性体验。
全面的遥测数据类型：涵盖跟踪、度量和日志，满足不同场景下的观测需求。
可插拔的后端：数据导出机制允许将数据发送到任意符合规范的后端服务，如 Jaeger、Prometheus、Elasticsearch 等，便于用户根据自身需求选择合适的存储和分析工具。
常见使用场景：

分布式系统追踪：在微服务架构中，请求可能会经过多个服务。OpenTelemetry 能够跨服务边界追踪请求的完整路径，帮助识别性能瓶颈和错误源头。

性能监控：通过收集和分析度量数据，如响应时间、错误率和吞吐量，来评估系统的整体健康状况和性能表现。

故障排查：当系统出现问题时，利用详细的跟踪信息和日志记录，快速定位问题发生的具体环节和原因。

容量规划：长期收集和分析度量数据，可以帮助团队更好地理解资源使用情况，为未来的容量规划提供数据支持。

安全审计：日志数据可用于审计目的，帮助识别潜在的安全威胁或异常访问模式。

合规性报告：确保系统操作符合特定行业或地区的合规要求，通过日志和度量数据来证明和报告合规状态。

总之，OpenTelemetry 通过提供统一的可观测性解决方案，帮助组织在复杂的现代软件环境中提升运维效率，加速问题解决，并优化应用性能


# 2、Github地址
[https://github.com/open-telemetry](https://github.com/open-telemetry)

# 3、Helm安装

> Helm是一个用于Kubernetes的应用程序包管理工具，它简化了Kubernetes应用的部署和管理过程。Helm的工作原理类似于Linux中的apt或者yum，但专门针对Kubernetes资源


## 3.1 添加Helm源

```bash
helm repo add grafana  https://grafana.github.io/helm-charts
```

## 3.2 下载chart包

```bash
helm pull open-telemetry/opentelemetry-collector --version 0.86.0
tar -zxvf opentelemetry-collector-0.86.0.tgz
```
## 3.3 安装

```bash
helm upgrade --install otel opentelemetry-collector/ -n otel
```
http端口号：4318
grpc端口号：4317

业务app访问地址：http://otel-opentelemetry-collector.otel.svc.cluster.local:4317/4318

