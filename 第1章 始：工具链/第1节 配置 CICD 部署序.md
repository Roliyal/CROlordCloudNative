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

至此，Helm CLI 安装完成，后续相关命令则可参考[helm]:https://helm.sh/zh/docs/intro/install/

如需要补全命令则需要追加命令补全

```
helm completion bash > /etc/bash_completion.d/helm
```

#### 步骤二：部署  helm chart Jenkins

1. 配置helm repo 地址以及更新本地索引

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

2. 创建 jenkins values 配置文件

```yaml
# jenkins_values.yaml
service:
  type: LoadBalancer
jenkins:
  Master:
    HostName: devops.roliyal.com
adminUser: crolord
adminKey: crolord
persistence:
  enabled: true
  size: 100Gi
image:
  tag: latest-jdk11
volumeMounts:
  - name: jenkins-data
    mountPath: /opt/jenkins
volumes:
  - name: jenkins-data
    persistentVolumeClaim:
      claimName: jenkins-pvc
sslCert:
  enabled: true
  certFileName: cert.crt
  keyFileName: key.key

```

提示：根据需要修改 values.yaml 中的配置项

```shell
jenkinsAdminPassword: Jenkins 管理员密码。
persistence.enabled: 是否启用持久化存储。
persistence.size: 持久化存储卷的大小。
persistence.storageClass: 存储卷的 StorageClass。
persistence.mountPath: 持久化存储卷挂载的路径。
service.type: Jenkins Service 的类型，可以设置为 ClusterIP、NodePort 或 LoadBalancer。
ingress.enabled: 是否启用 Ingress。
ingress.hosts: Ingress 的域名列表。
```

3. 部署 helm Jenkins

```
helm install jenkins jenkins/jenkins -f values.yaml
```

4. 安装完成后，可以使用以下命令查看 Jenkins 的状态，以及配置 jenkins 初始化

```
kubectl get pods -l "component=jenkins-master"
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
至此，整个Jenkins配置完成，后续插件配置使用以及 CICD 链路在后续章节展示。
