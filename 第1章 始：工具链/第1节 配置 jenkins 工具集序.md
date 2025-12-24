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

#### 步骤一：部署 Helm

- 初始化集群工具配置
创建基础配置（需配置完成 kubectl 命令工具完成，可以参考官网 [安装和设置 kubectl](https://kubernetes.io/docs/tasks/tools/?spm=5176.2020520152.0.0.49fd16ddyp09xv) 可参考附加配置 Kubectl
```shell

# 下载kubectl二进制文件
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"

# 下载kubectl校验和文件
curl -LO "https://dl.k8s.io/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl.sha256"

# 验证下载的kubectl文件
echo "$(cat kubectl.sha256)  kubectl" | sha256sum --check

# 将kubectl安装到系统路径
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl

# 创建.kube目录并编辑config文件，配置集群凭证到计算机 $HOME/.kube/config 文件下。
mkdir -p $(dirname $HOME/.kube/config) && vim $HOME/.kube/config

#将文件的权限设置为只有拥有者可以读写，解决安全警告的问题。
chmod 600 /root/.kube/config

# 获取节点信息（可选，仅在kubectl配置正确时有效）
kubectl get node

```
- 部署 Helm

·本文采用使用脚本安装

您可以获取这个脚本并在本地执行。它良好的文档会让您在执行之前知道脚本都做了什么。

前置安装 Docker-ce 
```shell
echo -e "[docker-ce-stable]\nname=Docker CE Stable - \$basearch\nbaseurl=https://mirrors.aliyun.com/docker-ce/linux/centos/7/\$basearch/stable\nenabled=1\ngpgcheck=1\ngpgkey=https://mirrors.aliyun.com/docker-ce/linux/centos/gpg" | sudo tee /etc/yum.repos.d/docker-ce.repo && sudo yum clean all && sudo yum makecache && sudo yum install -y docker-ce && sudo systemctl start docker && sudo systemctl enable docker
```

```shell
curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
chmod 700 get_helm.sh
./get_helm.sh

```

（二选一）如果想直接执行安装，运行以下命令

```shell
curl https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash
```

安装结果示意,并执行更新helm repo仓库Chart 中 Jenkins信息。

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

至此，Helm CLI 安装完成，后续相关命令则可参考 [helm](https://helm.sh/zh/docs/intro/install/)

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

2.1 创建 namespace 名称
```shell
kubectl create ns cicd
```
2.2 jenkins域名访问创建 HTTPS 证书（可以根据需求配置）
```shell
kubectl create secret tls [YOUR_TLS_SECRET_NAME] \
  --cert=path/to/cert/file.crt \
  --key=path/to/key/file.key \
  -n cicd
```
2.3 jenkins Web页面创建密码凭据
```shell
kubectl create secret generic jenkins-admin-secret \
  --from-literal=jenkins-admin-user=admin \
  --from-literal=jenkins-admin-password= 用于登录 jenkins 密码 \
  -n cicd
```
2.4 创建 jenkins 持久卷
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
创建 namespace 、
```shell

[root@CROLord]# kubectl create ns cicd
namespace/cicd created
[root@CROLord]# kubectl create secret tls jenkinstls   --cert=/opt/tls/jenkins.roliyal.com.crt   --key=/opt/tls/jenkins.roliyal.com.key   -n cicd
secret/jenkinstls created
# kubectl create secret tls jenkinstls   --cert=/root/opt/jenkins.roliyal.com_bundle.crt   --key=/root/opt/jenkins.roliyal.com.key   -n cicd
secret/jenkinstls created
[root@CROLord opt]# kubectl create secret generic jenkins-admin-secret \
>   --from-literal=jenkins-admin-user=admin \
>   --from-literal=jenkins-admin-password=CRLord@123 \
>   -n cicd
secret/jenkins-admin-secret created
[root@CROLord opt]# 
[root@CROLord opt]# ll
total 12
-rw-r--r-- 1 root root 4097 Feb 28 17:46 jenkins.roliyal.com_bundle.crt
-rw-r--r-- 1 root root 1706 Feb 28 17:46 jenkins.roliyal.com.key
[root@CROLord opt]# pwd
/root/opt
[root@CROLord opt]# kubectl apply -f - <<EOF
> apiVersion: v1
> kind: PersistentVolumeClaim
> metadata:
>   name: jenkins-data
>   namespace: cicd
> spec:
>   accessModes:
>     - ReadWriteOnce
>   resources:
>     requests:
>       storage: 50Gi
>   storageClassName: alicloud-disk-ssd
> EOF
persistentvolumeclaim/jenkins-data created
[root@CROLord opt]# sudo docker login --username=eb** crolord ACR 镜像地址.cn-hongkong.cr.aliyuncs.com 
Password: 
WARNING! Your password will be stored unencrypted in /root/.docker/config.json.
Configure a credential helper to remove this warning. See
https://docs.docker.com/engine/reference/commandline/login/#credentials-store

Login Succeeded
[root@CROLord opt]#  ll /root/.docker/config.json 
-rw-r--r-- 1 root root 145 Mar 20 16:13 /root/.docker/config.json
[root@CROLord opt]# cat /root/.docker/config.json 
{
        "auths": {
                "crolord-rACR 镜像地址kong.cr.aliyuncs.com": {
                        "auth": "Z********认证信息***w=="
                }
        }
}
[root@CROLord opt]#kubectl  create secret generic kaniko-secret --from-file=/root/.docker/config.json -n cicd 
secret/kaniko-secret created
```
###### 提示
- 使用 podman 模拟 Docker CLI 的时候，它的配置文件位置可能与传统的 Docker 不同。要找到 config.json 文件，可以使用以下步骤：

- 检查 podman 配置文件路径
- podman 的配置文件可能存储在不同的位置，通常是用户的主目录下。以下是常见的路径：
```shell
[root@CROLord ~]# find / -name auth.json 2>/dev/null
/run/containers/0/auth.json
[root@CROLord ~]# cat /run/containers/0/auth.json
{
        "auths": {
                "crolord*******ng.cr.aliyuncs.com": {
                        "auth": "Z**********rKktwcDIz"
                }
        }
}[root@CROLord ~]#kubectl  create secret generic kaniko-secret --from-file=/run/containers/0/auth.json -n cicd 
secret/kaniko-secret created
[root@CROLord ~]# 
```


###### 配置示例，用于初始化创建 secret ，此示例为 Jenkins credentials 全局凭据信息，相关信息根据实际情况配置
```shell
kubectl create secret generic secret-credentials -n cicd \
  --from-file=k8s-prod-config=/root/opt/crolord_prod.dingding  \
  --from-file=k8s-uat-config=/root/opt/crolord-ack-uat.dingding  \
  --from-file=github-token=/root/opt/github.txt \
  --from-literal=acr-username=ZWJAasdasda2NTkyMQ== \
  --from-literal=acr-password=Q1JzxczxcxMjM= \
```
- `k8s-prod-config` 和 `k8s-uat-config` 参数用于添加生产环境和测试环境的 Kubernetes config配置文件。
- `github-token` 参数添加一个认证 GitHub 令牌文件。
- `acr-username` 和 `acr-password` 用于添加经过 Base64 编码的阿里云容器注册服务(ACR)的登录凭证。请确保根据实际情况替换这些值，此变量仅在配置使用 docker 构建生效。


Base64 编码是一种将二进制数据转换成纯文本格式的编码方法。要生成 Base64 编码的字符串，你可以使用命令行工具或在线服务。对于命令行，如在 Linux 或 macOS 上，可以使用 base64 命令。例如，要对 "your_text" 进行编码，可以使用以下命令：

```shell
echo -n 'your_text' | base64
```
这里，echo -n 确保不在输出中包含新行字符，base64 命令将输入的文本转换成 Base64 编码。对于 Windows，可以使用 PowerShell 的 ConvertTo-Base64 命令。


###### 说明补充 ACR 容器镜像服务安全扫描全局凭据，以下数据通过 "echo -n 'your_access_key_id' | base64 生成，需要根据实际字段改变
```shell
kubectl patch secret secret-credentials -n cicd --patch='
 data:
   access_key_id: "eW91cl9hY2Nlc3Nfa2V5X2lk"
   access_key_secret: "eW91cl9hY2Nlc3Nfa2V5X2lk"
   token: "eW91cl9hY2Nlc3Nfa2V5X2lk"
 '
```
### 预期验证 secret-credentials 生效
```shell
[root@CROLord opt]# ll
total 32
-rw-r--r-- 1 root root 8002 Mar 15 20:58 crolord_prod.dingding
-rw-r--r-- 1 root root 8002 Mar 15 20:57 crolord_uat.dingding
-rw-r--r-- 1 root root   22 Mar 15 21:00 github.txt
-rw-r--r-- 1 root root 4097 Feb 28 17:46 jenkins.roliyal.com_bundle.crt
-rw-r--r-- 1 root root 1706 Feb 28 17:46 jenkins.roliyal.com.key
[root@CROLord opt]# kubectl create secret generic secret-credentials -n cicd \
>   --from-file=k8s-prod-config=/root/opt/crolord_prod.dingding  \
>   --from-file=k8s-uat-config=/root/opt/crolord_uat.dingding \
>   --from-file=github-token=/root/opt/github.txt \
>   --from-literal=acr-username=ZW***********yMQ== \
>   --from-literal=acr-password=Q1******xMjM=
secret/secret-credentials created
[root@CROLord opt]# kubectl get secret secret-credentials -n cicd -o jsonpath="{.data.acr-username}" | base64 --decode
```

2.5 配置 Helm Jenkins values 配置清单

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
      tag: "2.459"
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
      - configuration-as-code:1836.vccda_4a_122a_a_e
      - job-dsl:1.87
      - docker-build-publish:1.4.0
      - sshd:3.322.v159e91f6a_550
      - ws-cleanup:0.45
      - gson-api:2.11.0-41.v019fcf6125dc
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
               image: docker.io/crolord/kanikomanifest-tool:v1.1.0 #镜像版本为AMD架构，其中封装 kaniko、 trivy 、Manifest-tools 工具
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
helm -n cicd install    jenkins jenkins/jenkins -f /opt/tls/jenkins-values.dingding
```

4. 安装完成后，可以使用以下命令查看 Jenkins 的状态，以及配置 jenkins 初始化

4.1 查看 Jenkins 密码，本文默认密码则是步骤2.3 创建 secrets 凭据所设置密码
```
 kubectl exec --namespace cicd -it svc/jenkins -c jenkins -- /bin/cat /run/secrets/additional/chart-admin-password && echo
```

正常访问上述配置域名地址查看是否符合预期

#####  Helm SonarQube 部署
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
helm show values sonarqube/sonarqube > values.dingding
```
这会将当前图表的默认配置输出到values.yaml文件中。然后，你可以使用任何文本编辑器打开这个文件，并根据需要修改配置。比如，你可能想要修改以下一些配置：
```yaml
# Default values for sonarqube.
deploymentType: "StatefulSet"
replicaCount: 1
revisionHistoryLimit: 10
deploymentStrategy: {}
OpenShift:
  enabled: false
  createSCC: true
edition: "community"
image:
  repository: sonarqube
  tag: 10.4.1-{{ .Values.edition }}
  pullPolicy: IfNotPresent
securityContext:
  fsGroup: 0
containerSecurityContext:
  allowPrivilegeEscalation: false
  runAsNonRoot: true
  runAsUser: 1000
  runAsGroup: 0
  seccompProfile:
    type: RuntimeDefault
  capabilities:
    drop: ["ALL"]
elasticsearch:
  configureNode: false
  bootstrapChecks: true
service:
  type: ClusterIP
  externalPort: 9000
  internalPort: 9000
  labels:
  annotations: {}
networkPolicy:
  enabled: false
  prometheusNamespace: "monitoring"
sonarWebContext: ""
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
  tls:
    termination: edge
  annotations: {}
affinity: {}
tolerations: []
nodeSelector: {}
hostAliases: []
readinessProbe:
  initialDelaySeconds: 60
  periodSeconds: 30
  failureThreshold: 6
  timeoutSeconds: 1
livenessProbe:
  initialDelaySeconds: 60
  periodSeconds: 30
  failureThreshold: 6
  timeoutSeconds: 1
startupProbe:
  initialDelaySeconds: 30
  periodSeconds: 10
  failureThreshold: 24
  timeoutSeconds: 1
initContainers:
  securityContext:
    allowPrivilegeEscalation: false
    runAsNonRoot: true
    runAsUser: 1000
    runAsGroup: 0
    seccompProfile:
      type: RuntimeDefault
    capabilities:
      drop: ["ALL"]
  resources: {}
extraInitContainers: {}
extraContainers: []
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
  securityContext:
    privileged: true
    runAsUser: 0
initFs:
  enabled: true
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
  version: "0.17.2"
  noCheckCertificate: false
  webBeanPort: 8000
  ceBeanPort: 8001
  config:
    rules:
      - pattern: ".*"
prometheusMonitoring:
  podMonitor:
    enabled: false
    namespace: "default"
    interval: 30s
plugins:
  install: []
  noCheckCertificate: false
jvmOpts: ""
jvmCeOpts: ""
monitoringPasscode: "define_it"
annotations: {}
resources:
  limits:
    cpu: 800m
    memory: 4Gi
  requests:
    cpu: 400m
    memory: 2Gi
persistence:
  enabled: true
  annotations: {}
  storageClass: "alicloud-disk-topology-alltype"
  accessMode: ReadWriteOnce
  size: 50Gi
  uid: 1000
  guid: 0
  volumes: []
  mounts: []
emptyDir: {}
jdbcOverwrite:
  enable: false
  jdbcUrl: "jdbc:postgresql://myPostgress/myDatabase?socketTimeout=1500"
  jdbcUsername: "sonarUser"
  jdbcPassword: "sonarPass"
postgresql:
  enabled: true
  postgresqlUsername: "sonarUser"
  postgresqlPassword: "sonarPass"
  postgresqlDatabase: "sonarDB"
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
    enabled: true
    fsGroup: 1001
  containerSecurityContext:
    enabled: true
    runAsUser: 1001
    allowPrivilegeEscalation: false
    runAsNonRoot: true
    seccompProfile:
      type: RuntimeDefault
    capabilities:
      drop: ["ALL"]
  volumePermissions:
    enabled: false
    securityContext:
      runAsUser: 0
  shmVolume:
    chmod:
      enabled: false
  serviceAccount:
    enabled: false
podLabels: {}
sonarqubeFolder: /opt/sonarqube
tests:
  image: ""
  enabled: true
  resources: {}
serviceAccount:
  create: false
  annotations: {}
extraConfig:
  secrets: []
  configmaps: []
terminationGracePeriodSeconds: 60
```
- 持久化存储选择阿里云 storageClass 存储类，
- SonarQube版本 选择10.4.1
- 插件安装
- 资源限制（CPU、内存）
- 服务类型（比如，使用 clusterIP 以 mse ingress 访问）
4. 使用自定义Values文件安装SonarQube
- 修改values.yaml文件后，使用以下命令，通过自定义的values.yaml文件安装SonarQube：
```shell
helm upgrade --install -n cicd --version '~8' sonarqube sonarqube/sonarqube -f values.dingding
```
- 确保将<-n cicd >替换为实际的命名空间。
5. 访问SonarQube
   - 安装完成后，你可能需要执行一些额外的步骤来访问SonarQube界面，使用 MSE 查看SonarQube的IP地址，需要MSE控制台登录查看，并做域名映射。
```shell
[root@CROLord ~]#helm upgrade --install -n cicd --version '~8' sonarqube sonarqube/sonarqube -f values.dingding
NAMESPACE: cicd
STATUS: deployed
REVISION: 1
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


