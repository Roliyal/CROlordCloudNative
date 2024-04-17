# ACK CI配置

# CI的起源

持续集成（CI）是一种软件开发实践，旨在频繁地（通常是每天多次）将代码变更集成到共享存储库中。这一概念最早由Grady Booch在1991年提出，并随着敏捷软件开发的兴起而广泛应用。CI的主要目的是通过自动化构建和测试来快速发现集成错误，从而减少集成问题，提高软件质量。
# CI的前置要求
要成功实现持续集成（CI），需要满足几个关键的前置条件：

* **版本控制系统（VCS）**：CI的基石是版本控制。Git是目前最流行的VCS，它允许多名开发者同时工作在同一个项目上，而不会互相干扰。
* **测试驱动开发（TDD）**：TDD是一种软件开发过程，要求开发者在编写实际代码之前先编写测试用例。
* **集成环境的要求**：CI流程需要专门的服务器或服务来执行构建和测试任务。工具如Jenkins、Travis CI和CircleCI等提供了灵活的配置选项来满足不同项目的需求。

## Helm Jenkins 使用与配置
### 查看Helm Values字段含义
- 自定义Jenkins环境配置

 - 使用以下命令查看具体的Helm values字段含义，以及参考在线文档：

```shell
helm show values <Helm仓库>/<Helm镜像名>
```

 - 在线文档参考：[Jenkins Helm 参数参考集](https://github.com/jenkinsci/helm-charts/blob/main/charts/jenkins/VALUES.md)

###### 在values.yaml中配置Jenkins环境，以下是配置示例：

```yaml
# 持久化存储配置
persistence:
  enabled: true
  storageClass: "alicloud-disk-topology-alltype" #依次尝试创建指定的存储类型，并且使用WaitForFirstConsumer模式，可以兼容多可用区集群。 [参考StorageClass](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/disk-volume-overview-3#p-y0r-qmp-hxh)
  accessMode: ReadWriteOnce
  size: "100Gi"
  #subPath: "jenkins_home"
  # 以下部分为高级配置，如不需要可保持当前状态
  #existingClaim: jenkins-data
  annotations: {}
  labels: {}
  dataSource:
  volumes:
  mounts:

# JCasC 配置项
controller:
  usePodSecurityContext: true
  containerSecurityContext:
    runAsUser: 1000
    runAsGroup: 1000
    readOnlyRootFilesystem: true
    allowPrivilegeEscalation: true
  componentName: "jenkins-controller"
  image:
    registry: "docker.io"
    repository: "jenkins/jenkins"
    tag: "2.453"
    #tag: "2.444"
    #tagLabel: jdk17
    pullPolicy: "Always"
  # JCasC - 用户登录密码用户名配置
  #adminSecret: true
  #createSecret: true
  admin:
    createSecret: true
    existingSecret: "jenkins-admin-secret"
    userKey: "jenkins-admin-user"
    passwordKey: "jenkins-admin-password"
  #配置时区
  javaOpts: "-Duser.timezone=Asia/Shanghai -Dorg.jenkinsci.plugins.durabletask.BourneShellScript.USE_BINARY_WRAPPER=true"
  #配置插件
  installPlugins:
    - persistent-parameter:1.3
    - workflow-multibranch:783.va_6eb_ef636fb_d
    - kubernetes:4203.v1dd44f5b_1cf9
    - workflow-aggregator:latest
    - git:5.2.1
    - configuration-as-code:1775.v810dc950b_514
    - job-dsl:1.87
    - docker-build-publish:1.4.0
    - sshd:3.322.v159e91f6a_550
    - ws-cleanup:0.45
    - gson-api:2.10.1-15.v0d99f670e0a_7
    - github:1.38.0
    - build-name-setter:2.4.1
    - versionnumber:1.11
    - dingding-notifications:2.7.3
    - docker-workflow:572.v950f58993843
    - kubernetes-credentials-provider:1.262.v2670ef7ea_0c5
    - github-branch-source:1787.v8b_8cd49a_f8f1
    - build-name-setter:2.4.2
    - build-environment:1.7
    - sonar:2.17.2
    - pipeline-utility-steps:2.16.2
    - http_request:1.18
  existingSecret: "secret-credentials"
  additionalExistingSecrets:
    - name: secret-credentials
      keyName: ACR_USERNAME
    - name: secret-credentials
      keyName: ACR_PASSWORD
    - name: secret-credentials
      keyName: k8s-uat-config
    - name: secret-credentials
      keyName: k8s-prod-config
    - name: secret-credentials
      keyName: github-token
    - name: secret-credentials
      keyName: access_key_id
    - name: secret-credentials
      keyName: access_key_secret
    - name: secret-credentials
      keyName: token
  JCasC:
    defaultConfig: true
    configScripts:
      #用于定义预设jenkins 全局凭据，此选项为可选，如不需要则删除这段配置
      jenkins-casc-configs-credentials: |
        credentials:
          system:
            domainCredentials:
            - credentials:
              - string:
                  description: "AccessKey ID "
                  id: "access_key_id"
                  scope: GLOBAL
                  secret: ${secret-credentials-access_key_id}
              - string:
                  description: "AccessKey Secret"
                  id: "access_key_secret"
                  scope: GLOBAL
                  secret: ${secret-credentials-access_key_secret}
              - string:
                  description: "token"
                  id: "token"
                  scope: GLOBAL
                  secret: ${secret-credentials-token}
              - string:
                  description: "Kubernetes Token PROD"
                  id: "k8s_token-prod"
                  scope: GLOBAL
                  secret: ${secret-credentials-k8s-prod-config}
              - string:
                  description: "Kubernetes Token UAT"
                  id: "k8s_token_uat"
                  scope: GLOBAL
                  secret: ${secret-credentials-k8s-uat-config}
              - string:
                  description: "GitHub Token or Gitlab Token"
                  id: "github_token"
                  scope: GLOBAL
                  secret: ${secret-credentials-github-token}
              - usernamePassword:
                  description: "ACR Registry Credentials"
                  id: "acr_registry_credentials"
                  password: ${secret-credentials-ACR_PASSWORD}
                  scope: GLOBAL
                  username: ${secret-credentials-ACR_USERNAME}
      my-jobs: |
        jobs:
          - script: >
              job('demo-job') {
                  steps {
                      shell('echo Hello World')
                  }
              }
      update-center: |
        jenkins:
          updateCenter:
            sites:
              - id: "default"
                #如 dingding通知则需要配置官方源
                url: "https://updates.jenkins.io/update-center.json"
                #url: "https://mirrors.aliyun.com/jenkins/updates/update-center.json"
      my-jenkins-views: |
        jenkins:
          views:
            - list:
                name: "FEBEseparation-UAT"
                description: "FEBE separation UAT view"
               #jobNames:
                # - "demo-job"
                columns:
                  - "status"
                  - "weather"
                  - "jobName"
                  - "lastSuccess"
                  - "lastFailure"
                  - "lastDuration"
                  - "buildButton"
            - list:
                name: "FEBEseparation-Prod"
                description: "FEBE separation Production view"
                # jobNames:
            - list:
                name: "Microservice-UAT"
                description: "Microservice UAT view"
            - list:
                name: "Microservice-Prod"
                description: "Microservice Production view"

  #ingress 控制器设置
  serviceType: ClusterIP
  servicePort: 8080
  ingress:
    enabled: true
    apiVersion: "networking.k8s.io/v1"
    hostName: "jenkins.roliyal.com"
    annotations:
      nginx.ingress.kubernetes.io/rewrite-target: /
      nginx.ingress.kubernetes.io/ssl-redirect: "true"
    ingressClassName: mse # 因使用 mse 云原生网关，ingressclass名需根据实际填写
    tls:
      - secretName: "jenkins-tls"
        hosts:
          - "jenkins.roliyal.com"

# RBAC配置-默认开启
rbac:
  create: true
  readSecrets: true

# Prometheus监控配置
prometheus:
  enabled: true

# Jenkins代理配置
agent:
  enabled: true
  defaultsProviderTemplate: ""
  jenkinsUrl:
  jenkinsTunnel:
  kubernetesConnectTimeout: 5
  kubernetesReadTimeout: 15
  maxRequestsPerHostStr: "32"
  namespace: cicd
  image:
    repository: "jenkins/inbound-agent"
    tag: "latest"
  workingDir: "/home/jenkins/agent"
  nodeUsageMode: "NORMAL"
  customJenkinsLabels: []
  imagePullSecretName:
  componentName: "jenkins-agent"
  websocket: false
  privileged: true
  runAsUser: 0
  runAsGroup: 0
  resources:
    requests:
      cpu: "1024m"
      memory: "1024Mi"
    limits:
      cpu: "4096"
      memory: "4096Mi"
  alwaysPullImage: true
  podRetention: "Never"
  showRawYaml: false
  workspaceVolume: {}
  envVars: []
  #nodeSelector: {}
  command:
  args: "${computer.jnlpmac} ${computer.name}"
  sideContainerName: "jnlp"
  TTYEnabled: true
  containerCap: 10
  podName: "jenkins-slave"
  idleMinutes: 0
  yamlTemplate: ""
  yamlMergeStrategy: "override"
  connectTimeout: 100
  annotations: {}
  podTemplates:
    kanikoarm: |
      - name: kanikoarm
        label: kanikoarm
        serviceAccount: jenkins
        nodeSelector: "beta.kubernetes.io/arch=arm64"
        containers:
          - name: kanikoarm
            image: gcr.io/kaniko-project/executor:v1.22.0-debug 
            command: "sh -c"
            args: "cat"
            ttyEnabled: true
            privileged: false
            #resourceRequestCpu: "1000m"
            #resourceRequestMemory: "512Mi"
            #resourceLimitCpu: "2"
            #resourceLimitMemory: "2048Mi"
          - name: jnlp
            image: jenkins/inbound-agent:latest
            command: "jenkins-agent"
            args: ""
            ttyEnabled: true
            privileged: false
        volumes:
          - secretVolume:
              secretName: kaniko-secret
              mountPath: "/kaniko/.docker"
    kanikoamd: |
      - name: kanikoamd
        label: kanikoamd
        serviceAccount: jenkins
        nodeSelector: "beta.kubernetes.io/arch=amd64" #lables取自阿里云kubernetes 容器服务节点标签 
        containers:
          - name: kanikoamd
            image: docker.io/crolord/kanikomanifest-tool:v1.1.5 #镜像版本为AMD架构，其中封装 kaniko、 trivy 、Manifest-tools、 sonarqunbe 工具
            command: "sh -c"
            args: "cat"
            ttyEnabled: true
            privileged: false
            #resourceRequestCpu: "1000m"
            #resourceRequestMemory: "512Mi" 
            #resourceLimitCpu: "2"
            #resourceLimitMemory: "2048Mi"
          - name: jnlp
            image: jenkins/inbound-agent:latest
            command: "jenkins-agent"
            args: ""
            ttyEnabled: true
            privileged: false
        volumes:
          - secretVolume:
              secretName: kaniko-secret
              mountPath: "/kaniko/.docker"         
    podman: |
      - name: podman
        label: podman
        serviceAccount: default
        nodeSelector: "beta.kubernetes.io/arch=amd64" #lables取自阿里云kubernetes 容器服务节点标签 
        containers:
          - name: podman
            image: crolord/podman:latest #该镜像目前支持 amd64 架构
            command: "/bin/sh -c"
            args: "cat"
            ttyEnabled: true
            privileged: true
        volumes:
          - hostPathVolume:
              hostPath: "/sys"
              mountPath: "/sys"
    docker: |
      - name: docker
        label: docker
        serviceAccount: default
        containers:
          - name: docker
            image: docker:dind
            command: "/bin/sh -c"
            args: "cat"
            ttyEnabled: true
            privileged: true
        volumes:
          - hostPathVolume:
              hostPath: "/var/run/docker.sock"
              mountPath: "/var/run/docker.sock"
  volumes:
    - type: HostPath
      hostPath: /var/lib/containers
      mountPath: /var/lib/containers
    - type: HostPath
      hostPath: /sys
      mountPath: /sys
```

###### 持久化存储配置 (persistence)

- enabled: 开关项，当设置为true时，启用持久化存储，保证Jenkins数据（如作业配置、历史记录等）在Pod重新部署后依然保留。
- storageClass: 指定Kubernetes中定义的存储类别名称，此处为alicloud-disk-topology-alltype。不同的存储类别可能支持不同的性能、价格和可用区部署策略。WaitForFirstConsumer模式意味着卷的创建将延迟到Pod被调度到具体节点上，详细 StorageClass 配置，参考[StorageClass](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/disk-volume-overview-3#p-y0r-qmp-hxh)
- accessMode: 定义如何访问存储卷。ReadWriteOnce仅允许单个节点以读写模式挂载卷。
- size: 请求的存储卷大小，这里是100Gi。
- annotations, labels: 用于添加额外的注解和标签到存储卷，这些可以用于存储策略或成本中心等的管理。
- dataSource, volumes, mounts: 这些高级配置项用于指定数据来源、额外卷和挂载点，通常用于特殊存储需求或数据迁移。

###### JCasC配置项 (controller)

- usePodSecurityContext, containerSecurityContext: 这些配置定义了Pod和容器的安全约束，例如以指定用户身份运行（避免使用root），禁用文件系统的写入权限等，以增强安全性。
- componentName: 用于Kubernetes标签，有助于资源的组织和管理。
- image: 定义使用的Jenkins镜像，包括镜像仓库地址、镜像标签等。pullPolicy: "Always"意味着Kubernetes将在每次启动容器时尝试拉取最新的镜像。
- admin: 管理Jenkins的管理员用户和密码，可选地从Kubernetes Secret获取。
- javaOpts: Java虚拟机选项，用于自定义Jenkins运行时的Java环境，如设置时区和脚本执行方式。
- installPlugins: 列出需要在Jenkins中安装的插件及其版本。
- JCasC: 使用Jenkins Configuration as Code插件预配置Jenkins，包括凭证、作业、视图等，使配置可以作为代码管理。

###### RBAC配置 (rbac)

- create: 是否自动创建RBAC资源，如角色和角色绑定，以授予Jenkins必要的Kubernetes API权限。
- readSecrets: 是否允许Jenkins读取Kubernetes Secrets，通常用于管理凭证。
- Prometheus监控配置 (prometheus)
- enabled: 是否集成Prometheus监控，用于收集和查看Jenkins的运行指标。

###### Jenkins代理配置 (agent)

- enabled, image, workingDir等: 配置Jenkins代理的基本信息，如启用状态、代理使用的Docker镜像、工作目录等。代理用于执行构建任务，以分担主节点负载。
- resources: 定义代理Pod的资源请求和限制，确保Kubernetes能够根据负载合理调度资源。
- podTemplates: 定义不同场景下的代理模板，如基于不同架构（AMD64/ARM64）或不同工具（Kaniko、Docker、Podman）的构建环境。这些模板使得在不同环境下执行构建任务成为可能。

###### 其他配置

- ingress: 配置外部访问Jenkins服务的入口，包括主机名、TLS证书等。
  - 证书方式可以采用ALB ingress 方式选择三种证书方式， Secret证书 、自动发现证书、AlbConfig指定证书，参考链接 [使用ALB Ingress配置HTTPS监听证书](https://help.aliyun.com/zh/ack/ack-managed-and-ack-dedicated/user-guide/use-an-alb-ingress-to-configure-certificates-for-an-https-listener#context-1h8-lvh-2o5)
- serviceType, servicePort: 指定Jenkins服务的类型（如ClusterIP）和端口，决定如何在Kubernetes内部访问Jenkins。

+ 通过这些配置，可以详细控制Jenkins在Kubernetes环境中的部署方式，包括安全、存储、插件、构建代理等各个方面。配置时，应根据具体需求调整参数，以达到最佳的运行效果和资源利用率。

###### 自定义工具镜像补充

+ docker.io/crolord/kanikomanifest-tool:v1.1.5 镜像 dockerfile 示例

```dockerfile
# 使用 Kaniko 的最新版镜像作为构建阶段名为 plugin
FROM gcr.io/kaniko-project/executor:latest AS plugin

# 使用 manifest-tool 的 Docker 镜像
FROM docker.io/mplatform/manifest-tool:latest AS manifest-tool

# 添加一个构建阶段用来从 bitnami/trivy 镜像中复制 trivy
FROM docker.io/bitnami/trivy:latest AS trivy

# 使用 SonarScanner 的 Docker 镜像
FROM sonarsource/sonar-scanner-cli:latest AS sonar-scanner

# 基础镜像 ubuntu 并指定标签，确保构建的一致性
FROM ubuntu:20.04
RUN apt-get update && apt-get install -y software-properties-common && \
    add-apt-repository ppa:openjdk-r/ppa && \
    apt-get update && apt-get install -y openjdk-17-jre-headless && \
    rm -rf /var/lib/apt/lists/*

# Set JAVA_HOME environment variable
ENV JAVA_HOME /usr/lib/jvm/java-17-openjdk-amd64
ENV PATH $PATH:$JAVA_HOME/bin

# 设置 Docker 配置环境变量

ENV DOCKER_CONFIG /kaniko/.docker

# 从 plugin 阶段复制 Kaniko 执行程序
COPY --from=plugin /kaniko/executor /usr/local/bin/kaniko

# 从 manifest-tool 镜像复制 manifest-tool 二进制到当前镜像
COPY --from=manifest-tool /manifest-tool /usr/local/bin/manifest-tool

# 从 trivy 阶段复制 trivy 二进制到当前镜像
COPY --from=trivy /opt/bitnami/trivy/bin/trivy /usr/local/bin/trivy

# 从 sonar-scanner 阶段复制 sonar-scanner 二进制到当前镜像
COPY --from=sonar-scanner /opt/sonar-scanner /opt/sonar-scanner
ENV PATH="/opt/sonar-scanner/bin:${PATH}"


```
## Helm SonarQube 使用与配置
1. 添加Helm仓库
  - 首先，需要将存放SonarQube Helm图表的仓库添加到Helm中。可以使用下面的命令添加官方的SonarQube Helm仓库：

```bash
helm repo add sonarqube https://SonarSource.github.io/helm-chart-sonarqube
helm repo update
```

2. 安装SonarQube
 - 接下来，使用Helm安装SonarQube。你可以直接安装，也可以先下载values.yaml文件进行修改，然后再安装。先简单安装看看：

```shell
helm install sonarqube sonarqube/sonarqube --version <chart_version> --namespace <your_namespace> --create-namespace

```
+ 将<chart_version>替换为你想安装的版本号，将<your_namespace>替换为你想在其上安装SonarQube的Kubernetes命名空间。

3. 修改Values
   Helm图表的安装和配置可以通过修改values.yaml文件来定制。你可以从图表仓库中下载默认的values.yaml文件，进行修改：

```shell
helm show values sonarqube/sonarqube > values.yaml
```
这会将当前图表的默认配置输出到values.yaml文件中。然后，你可以使用任何文本编辑器打开这个文件，并根据需要修改配置。比如，你可能想要修改以下一些配置：
```yaml
# Default values for sonarqube.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

# If the deployment Type is set to Deployment sonarqube is deployed as a replica set.
deploymentType: "StatefulSet"

# There should not be more than 1 sonarqube instance connected to the same database. Please set this value to 1 or 0 (in case you need to scale down programmatically).
replicaCount: 1

# How many revisions to retain (Deployment ReplicaSets or StatefulSets)
revisionHistoryLimit: 10

# This will use the default deployment strategy unless it is overriden
deploymentStrategy: {}
# Uncomment this to scheduler pods on priority
# priorityClassName: "high-priority"

## Use an alternate scheduler, e.g. "stork".
## ref: https://kubernetes.io/docs/tasks/administer-cluster/configure-multiple-schedulers/
##
# schedulerName:

## Is this deployment for OpenShift? If so, we help with SCCs
OpenShift:
  enabled: false
  createSCC: true

edition: "community"

image:
  repository: sonarqube
  tag: 10.4.1-{{ .Values.edition }}
  pullPolicy: IfNotPresent
  # If using a private repository, the imagePullSecrets to use
  # pullSecrets:
  #   - name: my-repo-secret

# Set security context for sonarqube pod
securityContext:
  fsGroup: 0

# Set security context for sonarqube container
containerSecurityContext:
  # Sonarqube dockerfile creates sonarqube user as UID and GID 0
  # Those default are used to match pod security standard restricted as least privileged approach
  allowPrivilegeEscalation: false
  runAsNonRoot: true
  runAsUser: 1000
  runAsGroup: 0
  seccompProfile:
    type: RuntimeDefault
  capabilities:
    drop: ["ALL"]

# Settings to configure elasticsearch host requirements
elasticsearch:
  # DEPRECATED: Use initSysctl.enabled instead
  configureNode: false
  bootstrapChecks: true

service:
  type: ClusterIP
  externalPort: 9000
  internalPort: 9000
  labels:
  annotations: {}
  # May be used in example for internal load balancing in GCP:
  # cloud.google.com/load-balancer-type: Internal
  # loadBalancerSourceRanges:
  #   - 0.0.0.0/0
  # loadBalancerIP: 1.2.3.4

# Optionally create Network Policies
networkPolicy:
  enabled: false

  # If you plan on using the jmx exporter, you need to define where the traffic is coming from
  prometheusNamespace: "monitoring"

  # If you are using a external database and enable network Policies to be created
  # you will need to explicitly allow egress traffic to your database
  # expects https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/#networkpolicyspec-v1-networking-k8s-io
  # additionalNetworkPolicys:

# will be used as default for ingress path and probes path, will be injected in .Values.env as SONAR_WEB_CONTEXT
# if .Values.env.SONAR_WEB_CONTEXT is set, this value will be ignored
sonarWebContext: ""

# also install the nginx ingress helm chart
nginx:
  enabled: false

ingress:
  enabled: true
  apiVersion: "networking.k8s.io/v1"
  hosts:
    - name: sonarqube.roliyal.com
      path: /
      servicePort: 9000
      pathType: ImplementationSpecific
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /
  ingressClassName: mse

  tls: []
  # - secretName: chart-example-tls
  #   hosts:
  #     - chart-example.local

route:
  enabled: false
  host: ""
  # Add tls section to secure traffic. TODO: extend this section with other secure route settings
  # Comment this out if you want plain http route created.
  tls:
    termination: edge

  annotations: {}
  # See Openshift/OKD route annotation
  # https://docs.openshift.com/container-platform/4.10/networking/routes/route-configuration.html#nw-route-specific-annotations_route-configuration
  # haproxy.router.openshift.io/timeout: 1m

  # Additional labels for Route manifest file
  # labels:
  #  external: 'true'

# Affinity for pod assignment
# Ref: https://kubernetes.io/docs/concepts/configuration/assign-pod-node/#affinity-and-anti-affinity
affinity: {}

# Tolerations for pod assignment
# Ref: https://kubernetes.io/docs/concepts/configuration/taint-and-toleration/
# taint a node with the following command to mark it as not schedulable for new pods
# kubectl taint nodes <node> sonarqube=true:NoSchedule
# The following statement will tolerate this taint and as such reverse a node for sonarqube
tolerations: []
#  - key: "sonarqube"
#    operator: "Equal"
#    value: "true"
#    effect: "NoSchedule"

# Node labels for pod assignment
# Ref: https://kubernetes.io/docs/user-guide/node-selection/
# add a label to a node with the following command
# kubectl label node <node> sonarqube=true
nodeSelector: {}
#  sonarqube: "true"

# hostAliases allows the modification of the hosts file inside a container
hostAliases: []
# - ip: "192.168.1.10"
#   hostnames:
#   - "example.com"
#   - "www.example.com"

readinessProbe:
  initialDelaySeconds: 60
  periodSeconds: 30
  failureThreshold: 6
  # Note that timeoutSeconds was not respected before Kubernetes 1.20 for exec probes
  timeoutSeconds: 1
  # If an ingress *path* other than the root (/) is defined, it should be reflected here
  # A trailing "/" must be included
  # deprecated please use sonarWebContext at the value top level
  # sonarWebContext: /

livenessProbe:
  initialDelaySeconds: 60
  periodSeconds: 30
  failureThreshold: 6
  # Note that timeoutSeconds was not respected before Kubernetes 1.20 for exec probes
  timeoutSeconds: 1
  # If an ingress *path* other than the root (/) is defined, it should be reflected here
  # A trailing "/" must be included
  # deprecated please use sonarWebContext at the value top level
  # sonarWebContext: /

startupProbe:
  initialDelaySeconds: 30
  periodSeconds: 10
  failureThreshold: 24
  # Note that timeoutSeconds was not respected before Kubernetes 1.20 for exec probes
  timeoutSeconds: 1
  # If an ingress *path* other than the root (/) is defined, it should be reflected here
  # A trailing "/" must be included
  # deprecated please use sonarWebContext at the value top level
  # sonarWebContext: /

initContainers:
  # image: busybox:1.36
  # We allow the init containers to have a separate security context declaration because
  # the initContainer may not require the same as SonarQube.
  # Those default are used to match pod security standard restricted as least privileged approach
  securityContext:
    allowPrivilegeEscalation: false
    runAsNonRoot: true
    runAsUser: 1000
    runAsGroup: 0
    seccompProfile:
      type: RuntimeDefault
    capabilities:
      drop: ["ALL"]
  # We allow the init containers to have a separate resources declaration because
  # the initContainer does not take as much resources.
  resources: {}

# Extra init containers to e.g. download required artifacts
extraInitContainers: {}

## Array of extra containers to run alongside the sonarqube container
##
## Example:
## - name: myapp-container
##   image: busybox
##   command: ['sh', '-c', 'echo Hello && sleep 3600']
##
extraContainers: []

## Provide a secret containing one or more certificate files in the keys that will be added to cacerts
## The cacerts file will be set via SONARQUBE_WEB_JVM_OPTS and SONAR_CE_JAVAOPTS
##
caCerts:
  enabled: false
  image: adoptopenjdk/openjdk11:alpine
  secret: your-secret

initSysctl:
  enabled: true
  vmMaxMapCount: 524288
  fsFileMax: 131072
  nofile: 131072
  nproc: 8192
  # image: busybox:1.36
  securityContext:
    # Compatible with podSecurity standard privileged
    privileged: true
    # if run without root permissions, error "sysctl: permission denied on key xxx, ignoring"
    runAsUser: 0
  # resources: {}

# This should not be required anymore, used to chown/chmod folder created by faulty CSI driver that are not applying properly POSIX fsgroup.
initFs:
  enabled: true
  # Image: busybox:1.36
  # Compatible with podSecurity standard baseline.
  securityContext:
    privileged: false
    runAsNonRoot: false
    runAsUser: 0
    runAsGroup: 0
    seccompProfile:
      type: RuntimeDefault
    capabilities:
      drop: ["ALL"]
      add: ["CHOWN"]

prometheusExporter:
  enabled: false
  # jmx_prometheus_javaagent version to download from Maven Central
  version: "0.17.2"
  # Alternative full download URL for the jmx_prometheus_javaagent.jar (overrides prometheusExporter.version)
  # downloadURL: ""
  # if you need to ignore TLS certificates for whatever reason enable the following flag
  noCheckCertificate: false

  # Ports for the jmx prometheus agent to export metrics at
  webBeanPort: 8000
  ceBeanPort: 8001

  config:
    rules:
      - pattern: ".*"
  # Overrides config for the CE process Prometheus exporter (by default, the same rules are used for both the Web and CE processes).
  # ceConfig:
  #   rules:
  #     - pattern: ".*"
  # image: curlimages/curl:8.2.1
  # For use behind a corporate proxy when downloading prometheus
  # httpProxy: ""
  # httpsProxy: ""
  # noProxy: ""
  # Reuse default initcontainers.securityContext that match restricted pod security standard
  # securityContext: {}

prometheusMonitoring:
  # Generate a Prometheus Pod Monitor (https://github.com/coreos/prometheus-operator)
  #
  podMonitor:
    # Create PodMonitor Resource for Prometheus scraping
    enabled: false
    # Specify a custom namespace where the PodMonitor will be created
    namespace: "default"
    # Specify the interval how often metrics should be scraped
    interval: 30s
    # Specify the timeout after a scrape is ended
    # scrapeTimeout: ""
    # Name of the label on target services that prometheus uses as job name
    # jobLabel: ""

# List of plugins to install.
# For example:
# plugins:
#  install:
#    - "https://github.com/AmadeusITGroup/sonar-stash/releases/download/1.3.0/sonar-stash-plugin-1.3.0.jar"
#    - "https://github.com/SonarSource/sonar-ldap/releases/download/2.2-RC3/sonar-ldap-plugin-2.2.0.601.jar"
#
plugins:
  install: []

  # For use behind a corporate proxy when downloading plugins
  # httpProxy: ""
  # httpsProxy: ""
  # noProxy: ""

  # image: curlimages/curl:8.2.1
  # resources: {}

  # .netrc secret file with a key "netrc" to use basic auth while downloading plugins
  # netrcCreds: ""

  # Set to true to not validate the server's certificate to download plugin
  noCheckCertificate: false
  # Reuse default initcontainers.securityContext that match restricted pod security standard
  # securityContext: {}

## (DEPRECATED) The following value sets SONAR_WEB_JAVAOPTS (e.g., jvmOpts: "-Djava.net.preferIPv4Stack=true"). However, this is deprecated, please set SONAR_WEB_JAVAOPTS or sonar.web.javaOpts directly instead.
jvmOpts: ""

## (DEPRECATED) The following value sets SONAR_CE_JAVAOPTS. However, this is deprecated, please set SONAR_CE_JAVAOPTS or sonar.ce.javaOpts directly instead.
jvmCeOpts: ""

## a monitoring passcode needs to be defined in order to get reasonable probe results
# not setting the monitoring passcode will result in a deployment that will never be ready
monitoringPasscode: "define_it"
# Alternatively, you can define the passcode loading it from an existing secret specifying the right key
# monitoringPasscodeSecretName: "pass-secret-name"
# monitoringPasscodeSecretKey: "pass-key"

## Environment variables to attach to the pods
##
# env:
#   # If you use a different ingress path from /, you have to add it here as the value of SONAR_WEB_CONTEXT
#   - name: SONAR_WEB_CONTEXT
#     value: /sonarqube
#   - name: VARIABLE
#     value: my-value

# Set annotations for pods
annotations: {}

## We usually don't make specific ressource recommandations, as they are heavily dependend on
## The usage of SonarQube and the surrounding infrastructure.
## Adjust these values to your needs, but make sure that the memory limit is never under 4 GB
resources:
  limits:
    cpu: 800m
    memory: 4Gi
  requests:
    cpu: 400m
    memory: 2Gi

persistence:
  enabled: true
  ## Set annotations on pvc
  annotations: {}

  ## Specify an existing volume claim instead of creating a new one.
  ## When using this option all following options like storageClass, accessMode and size are ignored.
  # existingClaim:

  ## If defined, storageClassName: <storageClass>
  ## If set to "-", storageClassName: "", which disables dynamic provisioning
  ## If undefined (the default) or set to null, no storageClassName spec is
  ##   set, choosing the default provisioner.  (gp2 on AWS, standard on
  ##   GKE, AWS & OpenStack)
  ##
  storageClass: "alicloud-disk-topology-alltype"
  accessMode: ReadWriteOnce
  size: 50Gi
  uid: 1000
  guid: 0

  ## Specify extra volumes. Refer to ".spec.volumes" specification : https://kubernetes.io/fr/docs/concepts/storage/volumes/
  volumes: []
  ## Specify extra mounts. Refer to ".spec.containers.volumeMounts" specification : https://kubernetes.io/fr/docs/concepts/storage/volumes/
  mounts: []

# In case you want to specify different resources for emptyDir than {}
emptyDir: {}
  # Example of resouces that might be used:
  # medium: Memory
# sizeLimit: 16Mi

# A custom sonar.properties file can be provided via dictionary.
# For example:
# sonarProperties:
#   sonar.forceAuthentication: true
#   sonar.security.realm: LDAP
#   ldap.url: ldaps://organization.com

# Additional sonar properties to load from a secret with a key "secret.properties" (must be a string)
# sonarSecretProperties:

# Kubernetes secret that contains the encryption key for the sonarqube instance.
# The secret must contain the key 'sonar-secret.txt'.
# The 'sonar.secretKeyPath' property will be set automatically.
# sonarSecretKey: "settings-encryption-secret"

## Override JDBC values
## for external Databases
jdbcOverwrite:
  # If enable the JDBC Overwrite, make sure to set `postgresql.enabled=false`
  enable: false
  # The JDBC url of the external DB
  jdbcUrl: "jdbc:postgresql://myPostgress/myDatabase?socketTimeout=1500"
  # The DB user that should be used for the JDBC connection
  jdbcUsername: "sonarUser"
  # Use this if you don't mind the DB password getting stored in plain text within the values file
  jdbcPassword: "sonarPass"
  ## Alternatively, use a pre-existing k8s secret containing the DB password
  # jdbcSecretName: "sonarqube-jdbc"
  ## and the secretValueKey of the password found within that secret
  # jdbcSecretPasswordKey: "jdbc-password"

## (DEPRECATED) Configuration values for postgresql dependency
## ref: https://github.com/bitnami/charts/blob/master/bitnami/postgresql/README.md
postgresql:
  # Enable to deploy the bitnami PostgreSQL chart
  enabled: true
  ## postgresql Chart global settings
  # global:
  #   imageRegistry: ''
  #   imagePullSecrets: ''
  ## bitnami/postgres image tag
  # image:
  #   tag: 11.7.0-debian-10-r9
  # existingSecret Name of existing secret to use for PostgreSQL passwords
  # The secret has to contain the keys postgresql-password which is the password for postgresqlUsername when it is
  # different of postgres, postgresql-postgres-password which will override postgresqlPassword,
  # postgresql-replication-password which will override replication.password and postgresql-ldap-password which will be
  # used to authenticate on LDAP. The value is evaluated as a template.
  # existingSecret: ""
  #
  # The bitnami chart enforces the key to be "postgresql-password". This value is only here for historic purposes
  # existingSecretPasswordKey: "postgresql-password"
  postgresqlUsername: "sonarUser"
  postgresqlPassword: "sonarPass"
  postgresqlDatabase: "sonarDB"
  # Specify the TCP port that PostgreSQL should use
  service:
    port: 5432
  resources:
    limits:
      cpu: 2
      memory: 2Gi
    requests:
      cpu: 100m
      memory: 200Mi
  persistence:
    enabled: true
    accessMode: ReadWriteOnce
    size: 40Gi
    storageClass: "alicloud-disk-topology-alltype"
  securityContext:
    # For standard Kubernetes deployment, set enabled=true
    # If using OpenShift, enabled=false for restricted SCC and enabled=true for anyuid/nonroot SCC
    enabled: true
    # fsGroup specification below are not applied if enabled=false. enabled=false is the required setting for OpenShift "restricted SCC" to work successfully.
    # postgresql dockerfile sets user as 1001
    fsGroup: 1001
  containerSecurityContext:
    # For standard Kubernetes deployment, set enabled=true
    # If using OpenShift, enabled=false for restricted SCC and enabled=true for anyuid/nonroot SCC
    enabled: true
    # runAsUser specification below are not applied if enabled=false. enabled=false is the required setting for OpenShift "restricted SCC" to work successfully.
    # postgresql dockerfile sets user as 1001, the rest aim at making it compatible with restricted pod security standard.
    runAsUser: 1001
    allowPrivilegeEscalation: false
    runAsNonRoot: true
    seccompProfile:
      type: RuntimeDefault
    capabilities:
      drop: ["ALL"]
  volumePermissions:
    # For standard Kubernetes deployment, set enabled=false
    # For OpenShift, set enabled=true and ensure to set volumepermissions.securitycontext.runAsUser below.
    enabled: false
    # if using restricted SCC set runAsUser: "auto" and if running under anyuid/nonroot SCC - runAsUser needs to match runAsUser above
    securityContext:
      runAsUser: 0
  shmVolume:
    chmod:
      enabled: false
  serviceAccount:
    ## If enabled = true, and name is not set, postgreSQL will create a serviceAccount
    enabled: false
    # name:

# Additional labels to add to the pods:
# podLabels:
#   key: value
podLabels: {}
# For compatibility with 8.0 replace by "/opt/sq"
# For compatibility with 8.2, leave the default. They changed it back to /opt/sonarqube
sonarqubeFolder: /opt/sonarqube

tests:
  image: ""
  enabled: true
  resources: {}

# For OpenShift set create=true to ensure service account is created.
serviceAccount:
  create: false
  # name:
  # automountToken: false # default
  ## Annotations for the Service Account
  annotations: {}

# extraConfig is used to load Environment Variables from Secrets and ConfigMaps
# which may have been written by other tools, such as external orchestrators.
#
# These Secrets/ConfigMaps are expected to contain Key/Value pairs, such as:
#
# apiVersion: v1
# kind: ConfigMap
# metadata:
#   name: external-sonarqube-opts
# data:
#   SONARQUBE_JDBC_USERNAME: foo
#   SONARQUBE_JDBC_URL: jdbc:postgresql://db.example.com:5432/sonar
#
# These vars can then be injected into the environment by uncommenting the following:
#
# extraConfig:
#   configmaps:
#     - external-sonarqube-opts

extraConfig:
  secrets: []
  configmaps: []

# account:
# The values can be set to define the current and the (new) custom admin passwords at the startup (the username will remain "admin")
#   adminPassword: admin
#   currentAdminPassword: admin
# The above values can be also provided by a secret that contains "password" and "currentPassword" as keys. You can generate such a secret in your cluster
# using "kubectl create secret generic admin-password-secret-name --from-literal=password=admin --from-literal=currentPassword=admin"
#   adminPasswordSecretName: ""
# # Reuse default initcontainers.securityContext that match restricted pod security standard
# #   securityContext: {}
#   resources:
#     limits:
#       cpu: 100m
#       memory: 128Mi
#     requests:
#       cpu: 100m
#       memory: 128Mi
# curlContainerImage: curlimages/curl:8.2.1
# adminJobAnnotations: {}
# deprecated please use sonarWebContext at the value top level
#   sonarWebContext: /
terminationGracePeriodSeconds: 60
```
- 持久化存储选择阿里云 storageClass 存储类，
- SonarQube版本 选择8.9.9
- 插件安装 
- 资源限制（CPU、内存）
- 服务类型（比如，使用 clusterIP 以 mse ingress 访问）
4. 使用自定义Values文件安装SonarQube
 - 修改values.yaml文件后，使用以下命令，通过自定义的values.yaml文件安装SonarQube：
```shell
helm upgrade --install -n cicd --version '~8' sonarqube sonarqube/sonarqube -f values.yaml```
 - 确保将<-n cicd >替换为实际的命名空间。

5. 访问SonarQube
   - 安装完成后，你可能需要执行一些额外的步骤来访问SonarQube界面，使用 MSE 查看SonarQube的IP地址，需要MSE控制台登录查看，并做域名映射。
   
```shell
[root@CROLord ~]#helm upgrade --install -n cicd --version '~8' sonarqube sonarqube/sonarqube -f values.yaml```
NAMESPACE: cicd
STATUS: deployed
REVISION: 2
NOTES:
1. Get the application URL by running these commands:
  http://sonarqube.roliyal.com
[root@CROLord ~]# 
```
- 然后，使用返回的 http://sonarqube.roliyal.com 在浏览器中访问SonarQube。

6. 清理
    - 如果需要，你可以通过以下命令删除SonarQube实例：
```shell
helm uninstall  -n cicd  sonarqube
```

```pipline
pipeline {
    // 定义使用的 Jenkins agent 类型
    agent { kubernetes { /* 配置省略 */ } }
    
    // 定义环境变量
    environment {
        GIT_BRANCH = 'main' // Git主分支的默认值
        MAJOR_VERSION = 'v1' // 主版本号
        MINOR_VERSION = '0'  // 次版本号
        PLATFORMS = 'linux/amd64,linux/arm64' // 构建目标平台
        MAJOR = "${params.MAJOR_VERSION ?: env.MAJOR_VERSION ?: '1'}" // 主版本号，允许通过参数覆盖
        MINOR = "${params.MINOR_VERSION ?: env.MINOR_VERSION ?: '0'}" // 次版本号，允许通过参数覆盖
        PATCH = "${env.BUILD_NUMBER}" // 构建号，用作修订版本号
        VERSION_TAG = "${MAJOR}.${MINOR}.${PATCH}" // 组合版本标签
        IMAGE_REGISTRY = "${params.IMAGE_REGISTRY}" // 镜像仓库地址
        IMAGE_NAMESPACE = "${params.IMAGE_NAMESPACE}" // 镜像命名空间
        IMAGE_ID = "${params.IMAGE_NAMESPACE}" // 镜像ID
        SONARQUBE_DOMAIN = "${params.SONARQUBE_DOMAINE}" // Sonarqube 域名配置
    }

    // 触发条件
    triggers { githubPush() }

    // 参数定义
    parameters {
        persistentString(name: 'BRANCH', defaultValue: 'main', description: 'Initial default branch: main')
        persistentChoice(name: 'PLATFORMS', choices: ['linux/amd64', 'linux/amd64,linux/arm64'], description: 'Target platforms, initial value: linux/amd64,linux/arm64')
        persistentString(name: 'GIT_REPOSITORY', defaultValue: 'https://github.com/Roliyal/CROlordCodelibrary.git', description: 'Git repository URL, default: https://github.com/Roliyal/CROlordCodelibrary.git')
        persistentString(name: 'MAJOR_VERSION', defaultValue: '1', description: 'Major version number, default: 1')
        persistentString(name: 'MINOR_VERSION', defaultValue: '0', description: 'Minor version number, default: 0')
        persistentString(name: 'BUILD_DIRECTORY', defaultValue: 'Chapter2KubernetesApplicationBuild/Unit2CodeLibrary/FEBEseparation/go-guess-number', description: 'Build directory path, default path: Chapter2KubernetesApplicationBuild/Unit2CodeLibrary/FEBEseparation/go-guess-number')
        persistentString(name: 'IMAGE_REGISTRY', defaultValue: 'crolord-registry-registry-vpc.cn-hongkong.cr.aliyuncs.com', description: 'Image registry address, default: crolord-registry-registry-vpc.cn-hongkong.cr.aliyuncs.com')
        persistentString(name: 'IMAGE_NAMESPACE', defaultValue: 'febe', description: 'Image namespace, default: febe')
        persistentString(name: 'SONARQUBE_DOMAINE', defaultValue: 'sonarqube.roliyal.com', description: 'SonarQube domain, default: sonarqube.roliyal.com')

    }
    
        // 构建流程定义
        stages {
            // 设置版本信息
            stage('Version') {
                steps {
                    script {
                        env.PATCH_VERSION = env.BUILD_NUMBER
                        env.VERSION_NUMBER = "${env.MAJOR}.${env.MINOR}.${env.PATCH_VERSION}"
                        echo "Current Version: ${env.VERSION_NUMBER}"
                    }
                }
            }
            
        // 检出代码
        stage('Checkout') {
            steps {
                cleanWs() // 清理工作空间
                script {
                    env.GIT_BRANCH = params.BRANCH
                }
                // 检出Git仓库
                checkout scm: [
                    $class: 'GitSCM',
                    branches: [[name: "*/${env.GIT_BRANCH}"]],
                    userRemoteConfigs: [[url: params.GIT_REPOSITORY]],
                    extensions: [[$class: 'CloneOption', depth: 1, noTags: false, reference: '', shallow: true]]
                ]
                echo '代码检出完成'
            }
        }
        
        // 检查目录和Dockerfile
        stage('Check Directory') {
            steps {
                echo "Current working directory: ${pwd()}"
                sh 'ls -la'
                stash includes: '**', name: 'source-code' // 存储工作空间，包括Dockerfile和应用代码
            }
        }
        stage('SonarQube analysis') {
            agent { kubernetes { inheritFrom 'kanikoamd' } }
            steps {
                // 从之前的阶段恢复存储的源代码
                unstash 'source-code'
        
                // 指定在特定容器中执行
                container('kanikoamd') {
                    // 设置SonarQube环境
                    withSonarQubeEnv('sonar') {
                        script {
                            // 使用withCredentials从Jenkins凭据中获取SonarQube token
                            withCredentials([string(credentialsId: 'sonar', variable: 'SONAR_TOKEN')]) {
                                // 执行sonar-scanner命令
                                sh """
                                sonar-scanner \
                                  -Dsonar.projectKey=${JOB_NAME} \
                                  -Dsonar.projectName='${env.IMAGE_NAMESPACE}' \
                                  -Dsonar.projectVersion=${env.VERSION_TAG} \
                                  -Dsonar.sources=. \
                                  -Dsonar.exclusions='**/*_test.go,**/vendor/**' \
                                  -Dsonar.language=go \
                                  -Dsonar.host.url=http://${env.SONARQUBE_DOMAIN} \
                                  -Dsonar.login=${SONAR_TOKEN} \
                                  -Dsonar.projectBaseDir=${env.BUILD_DIRECTORY}
                                """
                            }
                            
                            // 使用script块处理HTTP请求和JSON解析
                            withCredentials([string(credentialsId: 'sonar', variable: 'SONAR_TOKEN')]) {
                                def authHeader = "Basic " + ("${SONAR_TOKEN}:".bytes.encodeBase64().toString())
                                def response = httpRequest(
                                    url: "http://${env.SONARQUBE_DOMAIN}/api/qualitygates/project_status?projectKey=${JOB_NAME}",
                                    customHeaders: [[name: 'Authorization', value: authHeader]],
                                    consoleLogResponseBody: true,
                                    acceptType: 'APPLICATION_JSON',
                                    contentType: 'APPLICATION_JSON'
                                )
                                def json = readJSON text: response.content
                                if (json.projectStatus.status != 'OK') {
                                    error "SonarQube quality gate failed: ${json.projectStatus.status}"
                                } else {
                                    echo "Quality gate passed successfully."
                                }
                            }
                        }
                    }
                }
            }
        }


        // 并行构建阶段
        stage('Parallel Build') {
            parallel {
                // 为 amd64 构建镜像
                stage('Build for amd64') {
                    agent { kubernetes { inheritFrom 'kanikoamd' } }
                    steps {
                        unstash 'source-code' // 恢复之前存储的代码
                        container('kanikoamd') {
                            sh """
                                kaniko \
                                  --context ${env.WORKSPACE}/${params.BUILD_DIRECTORY} \
                                  --dockerfile ${params.BUILD_DIRECTORY}/Dockerfile \
                                  --destination ${env.IMAGE_REGISTRY}/${env.IMAGE_NAMESPACE}/${env.JOB_NAME}:${VERSION_TAG}-amd64 \
                                  --cache=true \
                                  --cache-repo=${env.IMAGE_REGISTRY}/${env.IMAGE_NAMESPACE}/cache \
                                  --skip-tls-verify \
                                  --skip-unused-stages=true \
                                  --custom-platform=linux/amd64 \
                                  --build-arg BUILDKIT_INLINE_CACHE=1 \
                                  --snapshotMode=redo \
                                  --log-format=text \
                                  --verbosity=info
                            """
                        }
                    }
                }
                // 为 arm64 构建镜像
                stage('Build for arm64') {
                    agent { kubernetes { inheritFrom 'kanikoarm' } }
                    steps {
                        unstash 'source-code'
                        container('kanikoarm') {
                            sh """
                            /kaniko/executor \
                              --context ${env.WORKSPACE}/${params.BUILD_DIRECTORY} \
                              --dockerfile ${params.BUILD_DIRECTORY}/Dockerfile \
                              --destination ${env.IMAGE_REGISTRY}/${env.IMAGE_NAMESPACE}/${env.JOB_NAME}:${VERSION_TAG}-arm64 \
                              --cache=true \
                              --cache-repo=${env.IMAGE_REGISTRY}/${env.IMAGE_NAMESPACE}/cache \
                              --skip-tls-verify \
                              --skip-unused-stages=true \
                              --custom-platform=linux/arm64 \
                              --build-arg BUILDKIT_INLINE_CACHE=1 \
                              --snapshotMode=redo \
                              --log-format=text \
                              --verbosity=info
                            """
                        }
                    }
                }
            }
        }
        
        // 推送多架构镜像 Manifest-tools
        stage('Push Multi-Arch Manifest') {
            agent { kubernetes { inheritFrom 'kanikoamd' } }
            steps {
                container('kanikoamd') {
                    script {
                        sh "manifest-tool --version "
                        // 创建并推送多架构镜像的manifest
                        sh """
                            manifest-tool --insecure push from-args \\
                            --platforms '${env.PLATFORMS}' \\
                            --template '${env.IMAGE_REGISTRY}/${env.IMAGE_NAMESPACE}/${env.JOB_NAME}:${env.VERSION_TAG}-ARCHVARIANT' \\
                            --target '${env.IMAGE_REGISTRY}/${env.IMAGE_NAMESPACE}/${env.JOB_NAME}:${env.VERSION_TAG}'
                        """
                        sh "trivy image --exit-code 1 --severity HIGH,CRITICAL --ignore-unfixed --no-progress --insecure --timeout 5m '${env.IMAGE_REGISTRY}/${env.IMAGE_NAMESPACE}/${env.JOB_NAME}:${env.VERSION_TAG}'"
                    }
                }
            }
        }

    }
}

```
