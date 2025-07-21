
# 鳄霸 · CROlord

## 序 · Prologue

&emsp;&emsp;有一群闲来无事的年轻人，想要深入探索云原生技术的边界，并打算从零构建一个示范性的最佳实践项目。他们既是资深外包技术专家，又对新兴技术充满热情。

&emsp;&emsp;一次偶然的灵感，他们决定用一个具有领域特征的名字来命名这个项目 —— **CROLord**，它象征着 "Crocodile Overlord（鳄鱼霸主）" 的缩写，寓意这个项目将在云原生治理中占据一席之地。

&emsp;&emsp;在项目建设过程中，团队使用容器化技术隔离模块，采用微服务架构解耦系统，并利用自动化工具提升部署与运维效率，打造了一个端到端的最佳实践样板。尽管过程中挑战重重，但通过持续学习与迭代，他们打造出一个高可靠、高可扩展的系统，收获了宝贵的云原生经验。

&emsp;&emsp;他们希望通过分享此项目，帮助更多开发者理解云原生概念，推动这一技术生态的发展。

---

## 目录 · Table of Contents

### 第1章 始：工具链
- 1.1 配置 Jenkins 工具集序
- 1.2 镜像优化与制作
- 1.3 YAML 文件编写规则
- 1.4 利用 ACR 镜像仓库构建镜像安全
- 1.5 小结

### 第2章 阿里云 Kubernetes 应用部署
- 2.0 ACK 一些初始化
- 2.1 配置CICD环境优化配置
- 2.2 Monolithic application GO VUE项目前后分离部署 ACK 环境
- 2.3 Monolithic application GO VUE项目前后分离部署 ACK 环境
- 2.4 Micro services JAVA微服务部署 ACK 环境

### 第3章 下一代 Serverless(SAE) 应用构建
- 3.1 配置 SAE Jenkins 环境
- 3.2 Monolithic application GO VUE项目前后分离部署 部署 SAE 环境
- 3.4 Monolithic application GO VUE项目前后分离部署 部署 SAE 环境

### 第4章 微服务治理实践与挑战
- 4.1 基于阿里云 MSE 实现 Go 服务全链路灰度发布 — 从代码到生产的全流程
- 4.2 基于 MSE 云原生网关实现VUE前端灰度发布
- 4.3 配置限流熔断环境
- 4.4 安全防护与优雅上下线配置指南

### 第5章 可观测思考与实践
- 5.1 使用 Helm 安装 Prometheus Server（远程写入）
- 5.2 基于阿里云 ARMS 的端到端可观测全链路追踪
- 5.3 基于 SLS 收集 ACK 业务日志实践（LoongCollector AliyunPipelineConfig）
- 5.4 基于阿里云 Prometheus、SLS 与云监控构建多数据源 Grafana 监控大盘


### 第6章 应用高可用与性能压测
- 6.1  基于阿里云 PTS 压测场景（登录 + 游戏操作 + 查询积分）

### 第7章 解决方案啊
- 7.1 附录 解决方案


---

## 项目介绍 · Introduction

本项目旨在构建一个贯穿云原生全生命周期的自动化实践平台，涵盖从容器构建、CI/CD、微服务治理、Serverless 架构部署、到可观测与压测的完整流程，支持：
- 企业级 Kubernetes（ACK）部署方案
- 阿里云 Serverless（SAE）全生命周期支持
- 多种语言（Go/Vue/）混合部署场景
- 微服务治理：流量灰度、限流熔断、安全策略
- 统一可观测方案（Prometheus、Grafana、ARMS、SLS）
- 应用高可用与性能测试方案

---

## 使用说明 · How to Use

1. 克隆仓库：
   ```bash
   git clone https://github.com/Roliyal/CROlordCloudNative.git
   git clone https://github.com/Roliyal/CROlordCodelibrary
   cd crolord
   ```

2. 初始化环境：
  - 按照各章节的 Jenkinsfile / YAML / code / 等模板执行
  - 提供代码细节片段位于 `CROlordCodelibrary` 目录

3. 参考章节文档操作：
  - 每章提供逐步实践文档（位于章节目录下）

---

## 贡献者 · Contributors

感谢以下贡献者在项目中的协作与支持：

- **@issac** - 项目骨架设计与微服务部署+Jenkins + Helm + YAML 构建实践+Vue 端灰度与监控配置+Go 微服务治理与 MSE 实践
- **@Rich-Liuxf** - 项目项目骨架设计+CICD优化+可观测配置+性能压测实践+微服务治理实践+Serverless部署实践

欢迎更多感兴趣的朋友参与！

---

## 开源许可 · License

本项目采用 [Attribution-ShareAlike 4.0 International](https://creativecommons.org/licenses/by-sa/4.0/) 协议：

您可以自由地：
- **分享** — 在任何媒介以任何形式复制、发行本作品
- **演绎** — 修改、转换本作品，用于任何用途，包括商业用途

但须遵守以下条件：
- **署名** — 必须给予适当的署名
- **相同方式共享** — 若修改或构建自本作品，您必须采用相同的许可协议进行共享

---

## 浅聊摸鱼之道 · Work-Life Balance

<div style="display: flex; flex-direction: row;">
  <img src="resource/images/Dingcode.png" alt="鳄霸日常" width="150" height="200" style="margin-right: 10px;">
</div>

---
欢迎 star 🌟、fork 🍴、issue 💬 与 PR 🤝！
