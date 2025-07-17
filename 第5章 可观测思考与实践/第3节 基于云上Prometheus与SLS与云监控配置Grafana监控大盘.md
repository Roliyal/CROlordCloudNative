
# 基于云上 Prometheus 与 SLS 与 云监控配置 Grafana 监控大盘详细文档

## 目录

1. [背景与目标](#背景与目标)
2. [名词解释](#名词解释)
3. [接入中心与接入方式](#接入中心与接入方式)
4. [数据接入流程](#数据接入流程)
5. [Prometheus 实例与存储策略](#prometheus-实例与存储策略)
6. [Grafana 配置与监控大盘](#grafana-配置与监控大盘)
7. [常见问题与故障排查](#常见问题与故障排查)
8. [参考链接](#参考链接)

## 背景与目标

### 产品详情
通过 **Prometheus** 和 **SLS** 集成，结合 **Grafana** 可视化，用户能够高效地实现全栈应用和基础设施监控。该平台支持对 **容器**、**ECS**、**数据库** 等资源的实时数据采集、存储、展示及分析。通过 **ARMS**（应用实时监控服务），用户还可以配置灵活的告警规则来确保应用的健康运行。

- **目标**：
    1. 通过 Prometheus 实例采集和存储不同服务的监控数据（容器、ECS、数据库等）。
    2. 配置 **Grafana** 大盘，展示实时监控数据并进行分析。
    3. 通过 **PromQL** 编写查询，获取高质量的监控数据并进行可视化。
## 架构图（Mermaid 格式）

```mermaid
flowchart LR
  subgraph Prometheus
    A[Prometheus Server] --> B[Grafana]
  end
  subgraph ARMS
    C[ARMS Metrics Collection]
  end
  subgraph Data Sources
    D[Cloud Services (ECS, RDS, etc.)]
    E[Custom Metrics]
  end
  A --> D
  A --> E
  B --> C
```


## 名词解释

| 名称                      | 说明                                               |
|-------------------------|--------------------------------------------------|
| **Prometheus**           | 开源的监控系统，专门设计用于时序数据的收集和存储。                        |
| **Grafana**              | 数据可视化工具，通过 Prometheus 数据源生成监控图表和大盘。                      |
| **SLS**                  | 简单日志服务（Simple Log Service），阿里云提供的日志管理与查询服务。              |
| **ARMS**                 | 阿里云的应用实时监控服务，提供应用链路追踪、日志监控及告警等功能。                   |
| **接入中心**             | 提供快速接入各种阿里云资源（如容器、ECS、数据库）的统一平台。                     |
| **Metricstore**          | Prometheus 数据存储库，保存监控指标数据并支持高效查询。                          |

## 接入中心与接入方式

### 接入中心入口

1. **登录 ARMS 控制台**：  
   进入 **阿里云 ARMS 控制台**，在左侧导航栏中选择 **接入中心**。

2. **选择接入的监控目标**：  
   在接入中心中，您可以看到各种支持的服务和组件，如 **容器服务**、**ECS 主机监控**、**数据库监控** 等。

   ![接入中心页面](https://your-screenshot-link)

#### 支持的接入方式

1. **容器集群监控**：支持 **ACK** 集群及容器化应用的监控，自动集成 Prometheus 进行监控。
2. **ECS 主机监控**：支持通过 **Prometheus** 监控 ECS 实例的性能，采用 **node-exporter** 和 **process-exporter**。
3. **云服务监控**：支持从云服务（如 **RDS**、**OSS**、**Kafka**）接入 Prometheus 数据。
4. **自定义监控**：可以通过配置文件接入任意的自定义监控数据。

## 数据接入流程

### 1. 数据接入步骤

**步骤 1：创建 Prometheus 实例**

1. 登录 **ARMS 控制台**。
2. 在 **接入中心** 中，选择您需要接入的监控目标（如容器集群、ECS、数据库等）。
3. 按照向导操作，创建 **Prometheus 实例**，并将相关云服务进行接入。

**步骤 2：配置 Prometheus 数据源**

1. 登录 **Grafana** 控制台。
2. 在左侧菜单中选择 **Data Sources**（数据源）。
3. 点击 **Add Data Source**（添加数据源），选择 **Prometheus**。
4. 配置 **Prometheus URL**（如 `http://prometheus-server:9090`）并保存数据源。

#### 配置 Prometheus 数据源的截图

![Grafana 数据源配置](https://your-screenshot-link)

### 2. Prometheus 查询（PromQL）

在 **Grafana** 中，您可以使用 **Prometheus Query Language (PromQL)** 编写查询来获取监控数据。

#### 示例：Prometheus 查询

- **获取 CPU 使用率**：
  ```prometheus
  avg(rate(node_cpu_seconds_total{mode="idle"}[5m])) by (instance)
  ```

- **获取内存使用情况**：
  ```prometheus
  avg(node_memory_MemTotal_bytes - node_memory_MemFree_bytes) by (instance)
  ```

- **获取网络传输量**：
  ```prometheus
  avg(rate(node_network_transmit_bytes_total[5m])) by (instance)
  ```

在 **Grafana** 中，通过编写 **PromQL** 查询语句，用户能够高效查询 Prometheus 中的数据并将其展示为图表或其他可视化元素。

## Prometheus 实例与存储策略

### 1. Prometheus 存储策略

在 **Prometheus** 服务中，每个接入的云产品（容器集群、ECS 主机、云服务等）都会创建一个对应的 **Prometheus 实例** 用于存储相关的监控数据。

- **容器集群**：为每个容器集群创建独立的 Prometheus 实例。
- **ECS 主机**：为每个 ECS 主机及其 VPC 网络创建独立的 Prometheus 实例。
- **云服务监控**：同一区域的云服务监控数据将存储在统一的 Prometheus 实例中。

## Grafana 配置与监控大盘

### 1. 默认 Grafana 大盘

阿里云 Prometheus 服务集成了多种 **Grafana 大盘**，包括但不限于：

- **ECS 总览大盘**：监控 ECS 实例的 CPU、内存、网络资源。
- **Node 进程级监控大盘**：监控节点上的进程资源消耗。
- **GPU 监控大盘**：用于监控具备 GPU 的 ECS 实例。

#### 创建自定义 Grafana 大盘

1. 登录 **Grafana** 控制台。
2. 在 **Dashboards** 页面，点击 **+ Create Dashboard**。
3. 添加 **Prometheus** 数据源，并创建相应的查询。

#### 示例：创建 ECS CPU 使用情况图表

1. 创建 **新面板**，选择 **Graph** 类型。
2. 输入以下 **PromQL** 查询：
   ```prometheus
   avg(rate(node_cpu_seconds_total{mode="idle"}[5m])) by (instance)
   ```
3. 在 **Legend** 配置中自定义面板名称，如：`CPU Usage for {instance}`。

## 常见问题与故障排查

### 1. 数据不显示
**原因**：数据接入延迟或 Prometheus 数据源配置不正确。  
**排查步骤**：
1. 等待 1-2 分钟，确保数据加载。
2. 检查 **Prometheus** 数据源配置，确保数据源 URL 和端口配置正确。

### 2. Prometheus 查询不返回数据
**原因**：查询语法错误，或者查询的时间范围内没有数据。  
**排查步骤**：
1. 确保 **PromQL** 查询语法正确。
2. 在查询界面调整 **时间范围**，确保查询的时间段内有数据。

## 参考链接

- [Prometheus 官方文档](https://prometheus.io/docs/)
- [Grafana 官方文档](https://grafana.com/docs/)
- [ARMS 产品手册](https://help.aliyun.com/document_detail/139265.html)
- [云上 Prometheus 配置与使用](https://help.aliyun.com/document_detail/139270.html)

## 架构图（Mermaid 格式）

```mermaid
flowchart LR
  subgraph Prometheus
    A[Prometheus Server] --> B[Grafana]
  end
  subgraph ARMS
    C[ARMS Metrics Collection]
  end
  subgraph Data Sources
    D[Cloud Services (ECS, RDS, etc.)]
    E[Custom Metrics]
  end
  A --> D
  A --> E
  B --> C
```
