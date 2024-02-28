### CICD工具链取舍之间

&emsp;&emsp;建立高效的企业 CICD 工具链对于现代软件开发至关重要。通过配置一个完善的 CICD 工具链，可以帮助您的企业实现持续交付、快速迭代、自动化测试和部署等目标，提高应用程序的质量和可靠性，提高团队的生产率和效率。然而，在配置企业 CICD 工具链之前，您需要了解各种 CI/CD 工具的优缺点，并考虑到团队的规模、技术栈、安全性和可靠性等因素，以确保选择适合您需求的工具和流程。在本文中，我们将介绍如何配置企业 CICD 工具链，包括选择适合您需求的工具、安装和配置工具、定义流程、确保安全性和可靠性等方面。我们希望这篇文章能够帮助您成功地配置和管理企业 CICD 工具链，实现高效的软件开发和交付。


| 工具                 | 语言   | 主要特点                               | 优点                                            | 缺点                                                   |
| -------------------- | ------ | -------------------------------------- | ----------------------------------------------- | ------------------------------------------------------ |
| Jenkins              | Java   | 开源、插件丰富、可定制化高             | 巨大的生态系统、强大的可扩展性和定制性          | UI界面较为陈旧、学习曲线较陡峭、需要手动管理安装和更新 |
| Travis CI            | Ruby   | 云托管、易于使用                       | 集成了许多CI/CD工具、易于配置和使用             | 定价较高、自定义能力较弱                               |
| CircleCI             | Python | 云托管、易于使用                       | 构建速度快、集成了许多CI/CD工具、易于配置和使用 | 定价较高、构建容器不能自定义                           |
| GitLab CI/CD         | Ruby   | 集成了CI/CD、代码仓库、Issue管理等功能 | 集成度高、易于使用、具有良好的可视化界面        | 部署相对较慢、文档和社区相对较弱                       |
| Bamboo               | Java   | 商业软件、企业级特性                   | 具有完整的DevOps生态系统、集成度高、易于使用    | 企业版定价较高、插件生态相对较弱                       |
| Azure DevOps         | .NET   | 集成了CI/CD、代码仓库、项目管理等功能  | 具有良好的可视化界面、易于使用、跨平台          | 定价较高、扩展和集成可能需要一定的学习成本             |
| CodeShip             | Ruby   | 云托管、易于使用                       | 部署速度快、易于配置和使用                      | 定价较高、功能相对较少                                 |
| AWS CodePipeline     | -      | AWS原生服务、易于使用                  | 集成了许多AWS服务、易于配置和使用               | 集成非AWS服务需要自定义插件                            |
| Aliyun CI/CD（云效） | -      | 阿里云原生服务、易于使用               | 集成了许多阿里云服务、与阿里云生态系统紧密结合  | 功能相对较少、不支持自定义镜像构建                     |
| Tencent Cloud CI/CD  | -      | 腾讯云原生服务、易于使用               | 集成了许多腾讯云服务、与腾讯云生态系统紧密结合  | 功能相对较少、不支持自定义镜像构建                     |
| Argo CD              | -      | 基于GitOps的Kubernetes应用程序管理工具 | 自动化部署和同步、可视化界面、应用程序版本控制  | 需要适应GitOps方式进行管理和部署                       |

&emsp;&emsp;希望这张表格能够帮助您更好地了解各种CI/CD工具的特点和优缺点。需要注意的是，每个企业的具体情况和需求不同，最终选择哪个工具应该根据实际情况。

### Argo CD 和 Jenkins 中选择 Jenkins 的理由

##### 用途和焦点:

&emsp;&emsp; Argo CD：Argo CD主要用于管理Kubernetes应用程序的生命周期。它基于GitOps的理念，通过将应用程序的配置和状态定义存储在Git版本控制系统中，并通过自动化来保持实际状态与期望状态的一致性。Argo CD专注于持续部署和管理云原生应用程序在Kubernetes集群上的运行。
&emsp;&emsp; Jenkins：Jenkins是一个用于持续集成和持续交付的自动化工具。它的主要用途是帮助开发团队在代码提交后自动构建、测试和部署应用程序，以实现快速、高质量的软件交付。
##### 环境：

&emsp;&emsp; Argo CD：Argo CD主要适用于Kubernetes环境，特别是用于部署和管理Kubernetes应用程序。
&emsp;&emsp; Jenkins：Jenkins是一个通用的自动化工具，可以用于构建和部署各种类型的应用程序，不仅限于Kubernetes环境。
##### 方式：

&emsp;&emsp; Argo CD：Argo CD采用GitOps的方式，通过Git仓库中的定义来持续监测和对比Kubernetes集群中的实际状态与期望状态，并自动进行部署、更新和回滚等操作。
&emsp;&emsp; Jenkins：Jenkins采用持续集成和持续交付的方式，通过自动化工作流来实现构建、测试和部署过程。
##### 界面：

&emsp;&emsp; Argo CD：Argo CD提供了直观的Web界面，方便用户查看应用程序的状态、历史版本和同步状态。
&emsp;&emsp; Jenkins：Jenkins也提供了Web界面，但更加灵活和定制化，用户可以根据需求配置和管理不同的构建和部署流程。
&emsp;&emsp; 综上所述，Argo CD和Jenkins是两种不同的工具，它们各自适用于不同的场景和用途。Argo CD主要用于持续部署和管理Kubernetes应用程序，而Jenkins是一个通用的自动化工具，可以用于构建和部署各种类型的应用程序。根据具体需求和环境，选择合适的工具能更好地满足项目的要求。

## 部署一个企业级 Jenkins 工具链

### 方案一基于阿里云 ACK（Kubernetes）Helm构建 Jenkins

&emsp;&emsp;在 ACK 集群中用 helm 部署 Jenkins 环境

#### 前提条件

已创建Kubernetes集群。具体操作，请参见创建Kubernetes托管版集群。
已通过kubectl连接到Kubernetes集群。具体操作，请参见通过kubectl工具连接集群。

#### 注意事项


| 事项       | 内容                                                   |
| ---------- | ------------------------------------------------------ |
| 版本兼容性 | Jenkins Helm 版本和 Kubernetes 版本、Helm 版本需要兼容 |
| 资源分配   | 合理分配资源，如内存、CPU、存储空间等                  |
| 配置参数   | 指定 Jenkins URL、Admin 用户名和密码等                 |
| 插件安装   | 可以通过配置 Helm chart 的 value 文件进行安装          |
| 数据持久化 | 需要配置存储卷以保证数据的持久化和可靠性               |
| 安全设置   | 部署后需要进行安全设置，如开启安全认证、插件安装等等   |

#### 步骤一：部署 Jenkins

- 部署 Helm

·本文采用使用脚本安装

您可以获取这个脚本并在本地执行。它良好的文档会让您在执行之前知道脚本都做了什么。

```shell
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh

```

如果想直接执行安装，运行以下命令

```shell
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

安装结果,并执行更新helm repo仓库Chart 中 Jenkins信息。

```
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/master/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh
[root@issac]# helm repo add jenkins https://charts.jenkins.io
"jenkins" has been added to your repositories
[root@issac ~]# helm repo update
Hang tight while we grab the latest from your chart repositories...
...Successfully got an update from the "jenkins" chart repository
Update Complete. ⎈Happy Helming!⎈
[root@issac~]#
```

至此，Helm CLI 安装完成，后续相关命令则可参考[helm]（https://helm.sh/zh/docs/intro/install/ ）

如需要补全命令则需要追加命令补全

```
helm completion bash > /etc/bash_completion.d/helm
```

#### 步骤二：部署  helm chart Jenkins

1. 配置helm repo 地址以及更新本地索引 （上述步骤完成后，此步骤可忽略）

```shell
helm repo add jenkins https://charts.jenkins.io
helm repo update
```

使用以下命令获取是否正常返回安装结果

```shell
helm repo list
```

预期返回结果如下

```shell
[root@issac ~]# helm repo list
NAME    URL            
jenkins https://charts.jenkins.io
[root@issac ~]#
```

2. 创建基础配置（需配置完成 kubectl 命令工具完成，可以参考官网[安装和设置 kubectl](https://kubernetes.io/docs/tasks/tools/?spm=5176.2020520152.0.0.49fd16ddyp09xv ）

2.1 创建 namespace 名称
```shell
kubectl create ns cicd
```
2.2 创建证书（可以根据需求配置）
```shell
kubectl create secret tls [YOUR_TLS_SECRET_NAME] \
  --cert=path/to/cert/file.crt \
  --key=path/to/key/file.key \
  -n cicd
```
2.3 创建密码凭据
```shell
kubectl create secret generic jenkins-admin-secret \
  --from-literal=jenkins-admin-user=admin \
  --from-literal=jenkins-admin-password= 用于登录 jenkins 密码 \
  -n cicd
```
2.4 创建持久卷
```shell
kubectl apply -f - <<EOF
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: jenkins-data
  namespace: cicd
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 50Gi
  storageClassName: alicloud-disk-ssd
EOF
```
示例
```shell
kubectl create secret tls jenkins-tls \
  --cert=/opt/tls/jenkins.roliyal.com.crt \
  --key=/opt/tls/jenkins.roliyal.com.key \
  -n cicd
---
[root@CROLord-To-ACK tls]# kubectl create ns cicd
namespace/cicd created
[root@CROLord-To-ACK tls]# kubectl create secret tls jenkins-tls   --cert=/opt/tls/jenkins.roliyal.com.crt   --key=/opt/tls/jenkins.roliyal.com.key   -n cicd
secret/jenkins-tls created
[root@CROLord-To-ACK tls]# 
```
###### 额外配置示例，用于初始化创建 secret ，此示例为 credentials 全局凭据信息，相关信息根据实际情况配置
```shell
kubectl create secret generic secret-credentials -n cicd \
  --from-file=k8s-prod-config=/opt/tls/config/crolord-ack-prod.yaml 本文演示使用生产环境 ACKkubeconfig 配置文件 \
  --from-file=k8s-uat-config=/opt/tls/config/crolord-ack-uat.yaml 本文演示使用测试环境 ACKkubeconfig 配置文件 \
  --from-file=github-token=/opt/tls/config/github.txt 本文演示使用github token 使用 \
  --from-literal=acr-username=TFRBSTV0OFJjcWkyeEtpNWtR 本文演示使用阿里云账号用于ACR登录账号（echo -n 'your_acr_username' | base64 生成，需要根据实际字段改变） \
  --from-literal=acr-password=TFRBSTV0OFJjcWkyeEtpNWtR /opt/tls/config本文演示使用阿里云账号用于ACR登录密码（echo -n 'your_acr_password' | base64 生成，需要根据实际字段改变）
```
###### 补充 ACR 容器镜像服务安全扫描全局凭据，以下数据通过 "echo -n 'your_access_key_id' | base64 生成，需要根据实际字段改变
```shell
kubectl patch secret secret-credentials -n cicd --patch='
 data:
   access_key_id: "TFRASDASDSADACCZ4NVhy"
   access_key_secret: "RHJuTASDASCZXCZCadasdbkRuMlR0"
   token: "M2VkMASCADASDSXSADASDASZXCVjNzVjasdA4"
 '
```
###示例  
[root@CROLord-To-ACK tls]# ll config/
total 20
-rw-r--r-- 1 root root 8003 Dec  6 16:28 crolord-ack-prod.yaml
-rw-r--r-- 1 root root 8005 Dec  6 16:29 crolord-ack-uat.yaml
-rw-r--r-- 1 root root   94 Dec  6 16:31 github.txt
[root@CROLord-To-ACK tls]# kubectl create secret generic secret-credentials \
> --from-file=k8s-prod-config=config/crolord-ack-prod.yaml \
> --from-file=k8s-uat-config=config/crolord-ack-uat.yaml \
> --from-literal=acr-username='本文演示使用阿里云账号用于ACR登录账号' \
> --from-literal=acr-password='本文演示使用阿里云账号用于ACR登录密码' \
> --from-file=github-token=config/github.txt \
> -n cicd
secret/secret-credentials created
[root@CROLord-To-ACK tls]# kubectl get secret secret-credentials -n cicd -o jsonpath="{.data.acr-username}" | base64 --decode

```

##### 将 path/to/cert/file.crt 和 path/to/key/file.key 替换为您的证书文件和密钥文件的实际路径，并将 [YOUR_TLS_SECRET_NAME] 替换为您想要给 Secret 的名称。更新 Helm 命令中的 [YOUR_TLS_SECRET_NAME] 为您刚刚创建的 Secret 的名称。
2.5 配置 values 配置清单
```yaml
# jenkins-values.yaml

# 持久化存储配置
persistence:
  enabled: true
  storageClass: "alicloud-disk-ssd"
  accessMode: ReadWriteOnce
  size: "50Gi"
  # subPath: "jenkins_home"
  # 以下部分为高级配置，如不需要可保持当前状态
  existingClaim:
  annotations: {}
  labels: {}
  dataSource:
  volumes:
  mounts:

# JCasC 配置项
controller:
  # JCasC - 用户登录密码用户名配置
  adminSecret: true
  admin:
    existingSecret: "jenkins-admin-secret"
    userKey: "jenkins-admin-user"
    passwordKey: "jenkins-admin-password"
  #配置时区
  javaOpts: "-Duser.timezone=Asia/Shanghai"
  #配置插件
  installPlugins:
    - workflow-multibranch:770.v1a_d0708dd1f6     
    - kubernetes:4151.v6fa_f0fb_0b_4c9
    - workflow-aggregator:latest
    - git:5.2.1
    - configuration-as-code:1756.v2b_7eea_874392
    - job-dsl:1.87
    - docker-build-publish:1.3.1
    - sshd:3.312.v1c601b_c83b_0e
    - ws-cleanup:0.45
    - github:1.37.3.1
    - build-name-setter:2.4.0
    - versionnumber:1.11
    - dingding-notifications:2.6.2
    - docker-workflow:572.v950f58993843
  existingSecret: "secret-credentials"
  additionalExistingSecrets:
    - name: secret-credentials
      keyName: acr-username
    - name: secret-credentials
      keyName: acr-password
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
                  password: ${secret-credentials-acr-password}
                  scope: GLOBAL
                  username: ${secret-credentials-acr-username}
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
    tls:
      - secretName: "jenkins-tls"
        hosts:
          - "jenkins.roliyal.com"

# RBAC配置-默认开启
rbac:
  create: true

# Prometheus监控配置
prometheus:
  enabled: true

# 备份配置
backup:
  enabled: true
  schedule: "0 2 * * *"

# Jenkins代理配置
agent:
   enabled: true
   defaultsProviderTemplate: ""
   jenkinsUrl:
   jenkinsTunnel:
   kubernetesConnectTimeout: 5
   kubernetesReadTimeout: 15
   maxRequestsPerHostStr: "32"
   namespace:
   image: "jenkins/inbound-agent"
   tag: "latest"
   workingDir: "/home/jenkins/agent"
   nodeUsageMode: "NORMAL"
   customJenkinsLabels: []
   imagePullSecretName:
   componentName: "jenkins-agent"
   websocket: false
   privileged: false
   runAsUser:
   runAsGroup:
   resources:
      requests:
         cpu: "512m"
         memory: "512Mi"
      limits:
         cpu: "512m"
         memory: "512Mi"
   alwaysPullImage: false
   podRetention: "Never"
   showRawYaml: true
   workspaceVolume: {}
   envVars: []
   nodeSelector: {}
   command:
   args: "${computer.jnlpmac} ${computer.name}"
   sideContainerName: "jnlp"
   TTYEnabled: false
   containerCap: 10
   podName: "jenkins-slave"
   idleMinutes: 0
   yamlTemplate: ""
   yamlMergeStrategy: "override"
   connectTimeout: 100
   annotations: {}
   podTemplates:
      podman: |
        - name: podman
          label: podman
          serviceAccount: jenkins
          containers:
            - name: podman
              image: crolord/podman:latest
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
          serviceAccount: jenkins
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

提示：根据需要修改 values.yaml 中的配置项

```shell
jenkinsAdminPassword: 设置用于 Jenkins 管理员账户的密码。在安装 Jenkins 时，您将使用这个密码登录 Jenkins 控制台。

persistence.enabled: 该设置决定是否为 Jenkins 启用持久化存储。启用后，Jenkins 的数据（如作业配置、构建历史等）将存储在持久化卷中，这样即使 Pod 重启，数据也不会丢失。

persistence.size: 指定持久化卷的大小。需要根据您预期的数据量来设置，确保有足够的空间存储 Jenkins 的数据和构建工件。

persistence.storageClass: 该字段指定要使用的存储类。StorageClass 由您的 Kubernetes 环境提供，定义了使用哪种存储后端以及如何创建持久化卷。

persistence.mountPath: 持久化卷在 Jenkins Pod 内部的挂载路径。这是 Jenkins 数据存储的目录，应确保 Jenkins 配置指向此路径。

service.type: 定义如何暴露 Jenkins 服务。ClusterIP 仅在内部集群网络中可用，NodePort 在节点的特定端口上对外提供服务，而 LoadBalancer 会请求云提供者创建一个负载均衡器，为服务提供单一的接入点。

ingress.enabled: 确定是否为 Jenkins 创建 Kubernetes Ingress 资源。Ingress 允许您通过定义的域名访问 Jenkins，通常与 Ingress 控制器一起使用，如 nginx 或 traefik。

ingress.hosts: 一个或多个您希望 Jenkins 响应的域名列表。您需要确保这些域名在 DNS 中正确解析到您的 Kubernetes Ingress 控制器的 IP 地址。
```

3. 部署 helm Jenkins

```
helm -n cicd install jenkins jenkins/jenkins -f jenkins-values.yaml
```

示例
```shell
helm -n cicd install    jenkins jenkins/jenkins -f /opt/tls/jenkins-values.yaml
```

4. 安装完成后，可以使用以下命令查看 Jenkins 的状态，以及配置 jenkins 初始化

4.1 查看 Jenkins 密码，本文默认密码则是步骤2.3 创建 secrets 凭据所设置密码
```
 kubectl exec --namespace cicd -it svc/jenkins -c jenkins -- /bin/cat /run/secrets/additional/chart-admin-password && echo
```

正常访问上述配置域名地址查看是否符合预期
##### 方案二基于 ECS 服务器构建 docker-compose-Jenkins

- 部署 docker && docker-compose

1. 安装Docker
   首先，您需要安装Docker。在 Centos系统中，可以通过运行以下命令来安装Docker：

```
sudo yum update
sudo yum install docker
```

在其他操作系统中，请参考Docker官方文档以了解如何安装Docker。

2. 配置Docker
   启动Docker服务并将其配置为在系统启动时自动启动。在Centos中，可以通过运行以下命令来完成此操作：

```

sudo mkdir -p /etc/docker
sudo tee /etc/docker/daemon.json <<-'EOF'
{
  "registry-mirrors": ["https://此处应为您的镜像加速地址.com"]
}
EOF
sudo systemctl daemon-reload
sudo systemctl restart docker

sudo systemctl start docker
sudo systemctl enable docker
```

3. 安装Docker Compose

```shell
sudo yum install docker-compose
```

4. 验证Docker和Docker Compose安装
   您可以通过运行以下命令来验证Docker和Docker Compose是否正确安装：

```shell
docker --version
docker-compose --version
```

如果两个命令都返回版本信息，则说明Docker和Docker Compose已成功安装并准备好使用。

- 部署docker-compose 环境下 jenkins

1. 创建 Jenkins 目录以及对应 Jenkins 权限,并配置 https 证书

```shell
mkdir -p /opt/jenkins/jenkins_data
mkdir -p /opt/jenkins/certs
sudo chown -R 1000:1000 /opt/jenkins/jenkins_data
sudo chown -R 1000:1000 /opt/jenkins/certs

openssl req -x509 -newkey rsa:4096 -keyout /opt/jenkins/certs/roliyal.key -out /opt/jenkins/certs/roliyal.crt -days 365 -nodes

```

2. 创建 docker-compose 配置 yaml 文件信息

```yaml
version: '3'

services:
  jenkins:
    hostname: devops.roliyal.com //您的域名
    image: jenkins/jenkins:latest
    container_name: jenkins
    restart: always
    ports:
      - "80:8080"
      - "443:8443"
    volumes:
      - /opt/jenkins/jenkins_data:/var/jenkins_home
      - /opt/jenkins/certs:/certs
    user: "1000:1000"
 
    environment:
      - JENKINS_OPTS=--prefix=/jenkins -Dorg.apache.commons.jelly.tags.fmt.timeZone=Asia/Shanghai --httpsPort=8443 --httpsCertificate=/certs/roliyal.crt --httpsPrivateKey=/certs/roliyal.key
      - JENKINS_UC=https://mirrors.aliyun.com/jenkins/updates/update-center.json
    networks:
      - jenkins-net
    depends_on:
      - jenkins-slave

  jenkins-slave:
    image: jenkins/jnlp-slave:alpine
    container_name: jenkins-slave
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
    networks:
      - jenkins-net

volumes:
  jenkins-data:

networks:
  jenkins-net:
    driver: bridge
```

重要信息

1. 配置 environment 时，需要注意certs目录，Jenkins_data目录位置是否与当前环境一致，否则会提示目录无法写入。
2. 证书需要正确获取 key，crt 信息。
3. 启动Jenkins容器：
   在终端中，导航到项目根目录并运行以下命令以启动Jenkins容器：

```shell
docker-compose up -d
```

此命令将创建并启动Jenkins容器。 -d参数告诉Docker在后台运行容器。

4、获取Jenkins 初始密码

```shell
 docker exec jenkins cat /var/jenkins_home/secrets/initialAdminPassword
```

此处 Jenkins 名为容器名，根据实际情况灵活变动。
至此，整个Jenkins配置完成，后续插件配置使用以及 CI/CD 链路在后续章节展示。
