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

&emsp;&emsp;希望这张表格能够帮助您更好地了解各种CI/CD工具的特点和优缺点。需要注意的是，每个企业的具体情况和需求不同，最终选择哪个工具应该根据实际情况。

### Jenkins 的理由

&emsp;&emsp;Jenkins 是目前最流行和广泛使用的 CI/CD 工具之一，它有以下几个优点：

- 开源免费：Jenkins 是开源免费的工具，不需要支付任何授权费用。
- 插件丰富：Jenkins 拥有大量的插件，可以满足各种不同的需求，例如版本控制、构建工具、测试工具、集成工具等等。
- 可扩展性强：Jenkins 的架构非常灵活，可以轻松添加自定义插件和扩展，方便用户根据自己的需求进行定制。
- 易于安装和使用：Jenkins 安装和配置非常简单，可以在几分钟内完成部署，同时它的界面也非常友好，方便用户进行操作和管理。
- 社区活跃：Jenkins 拥有庞大的社区和用户群体，用户可以在社区中获得免费的技术支持、文档、教程等等资源。
- 支持多种语言和技术栈：Jenkins 支持多种编程语言和技术栈，如 Java、Python、Node.js、Docker 等等，方便用户进行多语言和多技术栈的项目构建和部署。

&emsp;&emsp;基于以上优点，Jenkins 成为了企业中广泛使用的 CI/CD 工具之一。

## 部署一个企业级 Jenkins 工具链

### 方案一基于阿里云 ACK（Kubernetes）构建 Jenkins

&emsp;&emsp;在 ACK 集群中用 helm 部署 Jenkins 并完成应用构建和部署

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
| 安全设置   | 部署后需要进行安全设置，如开启安全认证、限制插件安装等 |

#### 步骤一：部署 Jenkins


- 部署 Helm

·本文采用使用脚本安装

您可以获取这个脚本并在本地执行。它良好的文档会让您在执行之前知道脚本都做了什么。

```azure
$ curl -fsSL -o get_helm.sh https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3
$ chmod 700 get_helm.sh
$ ./get_helm.sh

```
如果想直接执行安装，运行以下命令
```azure
crul https://raw.githubusercontent.com/helm/helm/main/scripts/get-helm-3 | bash。
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
· 配置helm repo 地址以及更新本地索引

```azure
helm repo add jenkins https://charts.jenkins.io
helm repo update
```
使用以下命令获取是否正常返回安装结果
```azure
helm repo list
```
预期返回结果如下
````
[root@issac ~]# helm repo list
NAME    URL                      
jenkins https://charts.jenkins.io
[root@issac ~]#
````

创建 jenkins 配置文件
```yaml

# jenkins_values.yaml
service:
  type: ClusterIP
jenkins:
  Master:
    HostName: devops.roliyal.com
adminUser: crolord
adminKey: crolord123
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
  certData: -----BEGIN CERTIFICATE-----
  MIIGbjCCBNagAwIBAgIRAOVk3l1yU+I3g7vs7GY4g9gwDQYJKoZIhvcNAQEMBQAw
  WTELMAkGA1UEBhMCQ04xJTAjBgNVBAoTHFRydXN0QXNpYSBUZWNobm9sb2dpZXMs
  IEluYy4xIzAhBgNVBAMTGlRydXN0QXNpYSBSU0EgRFYgVExTIENBIEcyMB4XDTIy
  MDgwMTAwMDAwMFoXDTIzMDgwMTIzNTk1OVowHTEbMBkGA1UEAxMSZGV2b3BzLnJv
  bGl5YWwuY29tMIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEA1KNPn7Qq
  rScXQYr44YIsSGwfPZNLE3y+PY++fDKeBfI9kkoUQARZQHI/Rj4ainlpo5uSOvV3
  RoP062La1/Vk5+JJS7R6xRBmlVOEG8Cmk1j0/YlQIgtP1qvpR3+lINn4VQJKR9D0
  PZrDWkac2DrgMrVs6YPxx9JeQFjnbQoJRn9BM7m5w1lLvFT+VgCIMgsOsVrQfgx1
  S4VTji1PdfkhRIKZdK7j0Uk9yZ43xcp+GEizc2Hc3UxbRGxsSzKQFOFUW0etrB6t
  wLn2SW566Uo6WG9LFcrIKLk7Hm1VnbJhfb8kqjIljckYAD1rE2aCLBRXOpY7A0kR
  IwfV7uTRgc+ttQIDAQABo4IC6zCCAucwHwYDVR0jBBgwFoAUXzp8ERB+DGdxYdyL
  o7UAA2f1VxwwHQYDVR0OBBYEFO2xaUXX6hZ6FDwYBEYEW2lKH5R/MA4GA1UdDwEB
  /wQEAwIFoDAMBgNVHRMBAf8EAjAAMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEF
  BQcDAjBJBgNVHSAEQjBAMDQGCysGAQQBsjEBAgIxMCUwIwYIKwYBBQUHAgEWF2h0
  dHBzOi8vc2VjdGlnby5jb20vQ1BTMAgGBmeBDAECATB9BggrBgEFBQcBAQRxMG8w
  QgYIKwYBBQUHMAKGNmh0dHA6Ly9jcnQudHJ1c3QtcHJvdmlkZXIuY24vVHJ1c3RB
  c2lhUlNBRFZUTFNDQUcyLmNydDApBggrBgEFBQcwAYYdaHR0cDovL29jc3AudHJ1
  c3QtcHJvdmlkZXIuY24wHQYDVR0RBBYwFIISZGV2b3BzLnJvbGl5YWwuY29tMIIB
  fQYKKwYBBAHWeQIEAgSCAW0EggFpAWcAdgCt9776fP8QyIudPZwePhhqtGcpXc+x
  DCTKhYY069yCigAAAYJXDNIYAAAEAwBHMEUCIDW9slRabxaCamyDu5UAw2n6NpZ1
  l+kfi7pnHWH3ZsuyAiEAsX6OY56M/QLrvoBHVJj6pxgW0XfnW/xqIOG2b8JiTmUA
  dQB6MoxU2LcttiDqOOBSHumEFnAyE4VNO9IrwTpXo1LrUgAAAYJXDNHhAAAEAwBG
  MEQCICfcbxjka6xrHrtbFAusG4Isjj/nFgFB43/RzZI2JcxBAiBrH+oZnSWmvvMQ
  35vjDb0QqkagYFn38BihWp82+bD9cwB2AOg+0No+9QY1MudXKLyJa8kD08vREWvs
  62nhd31tBr1uAAABglcM0bUAAAQDAEcwRQIhAIJKdvzbiztRQZ4+9hNZUWrnMMtc
  enGxKOwrUUZbVXj0AiB761o2uOJ8IT11B145sDIyFRk9RTaokv89s7scDxTPbjAN
  BgkqhkiG9w0BAQwFAAOCAYEAExq12sOwD0dcHm8FSzUbsazbXDbzDaMCrDI9PWCZ
  x5H7cfeUjmpMqUHEdXBvmiDJB5kJMsjkskfApS2wr/E9BGF7CwU8RkgDgKzw58xZ
  gdgJ0+diMGkuvoo87u8+5yY0j6NasUtmTzSEshp8y7lWbIU4jcuHpOZ73XTZDoGs
  BuUlclEE22Ye7YFxoEJyg+o2WK7oDRK+GtI1QjhpmyAg9d775xMzgcx7P59xBeL0
  94EEcIRON95IQmHYzjEVIgw3Ugux/XZ+/2rEljJL5sXmPozLRKqN+Tsgo77IjxZk
  c+KOM69Tqq7sb9OJOOWZ8fKxhcM+M6e7d1W7CQK6W8FtcyCdEf3HHtGQHVNUvfbp
  m9hHn0KgiK/s+gkA6uPGdk7dBZe5caIJoaT6MrrRnHq5IIjOzPUqe+t0HvPbJTGy
  KFHsSiEj/LxGSdkFOfpPLBDJahLqTHLb1h+53Wu4XYA3aJrjyhO5snzohT+ZbXLT
  J87QziYcTVVRyAb320C+ZebJ
  -----END CERTIFICATE-----
  -----BEGIN CERTIFICATE-----
  MIIFBzCCA++gAwIBAgIRALIM7VUuMaC/NDp1KHQ76aswDQYJKoZIhvcNAQELBQAw
  ezELMAkGA1UEBhMCR0IxGzAZBgNVBAgMEkdyZWF0ZXIgTWFuY2hlc3RlcjEQMA4G
  A1UEBwwHU2FsZm9yZDEaMBgGA1UECgwRQ29tb2RvIENBIExpbWl0ZWQxITAfBgNV
  BAMMGEFBQSBDZXJ0aWZpY2F0ZSBTZXJ2aWNlczAeFw0yMjAxMTAwMDAwMDBaFw0y
  ODEyMzEyMzU5NTlaMFkxCzAJBgNVBAYTAkNOMSUwIwYDVQQKExxUcnVzdEFzaWEg
  VGVjaG5vbG9naWVzLCBJbmMuMSMwIQYDVQQDExpUcnVzdEFzaWEgUlNBIERWIFRM
  UyBDQSBHMjCCAaIwDQYJKoZIhvcNAQEBBQADggGPADCCAYoCggGBAKjGDe0GSaBs
  Yl/VhMaTM6GhfR1TAt4mrhN8zfAMwEfLZth+N2ie5ULbW8YvSGzhqkDhGgSBlafm
  qq05oeESrIJQyz24j7icGeGyIZ/jIChOOvjt4M8EVi3O0Se7E6RAgVYcX+QWVp5c
  Sy+l7XrrtL/pDDL9Bngnq/DVfjCzm5ZYUb1PpyvYTP7trsV+yYOCNmmwQvB4yVjf
  IIpHC1OcsPBntMUGeH1Eja4D+qJYhGOxX9kpa+2wTCW06L8T6OhkpJWYn5JYiht5
  8exjAR7b8Zi3DeG9oZO5o6Qvhl3f8uGU8lK1j9jCUN/18mI/5vZJ76i+hsgdlfZB
  Rh5lmAQjD80M9TY+oD4MYUqB5XrigPfFAUwXFGehhlwCVw7y6+5kpbq/NpvM5Ba8
  SeQYUUuMA8RXpTtGlrrTPqJryfa55hTuX/ThhX4gcCVkbyujo0CYr+Uuc14IOyNY
  1fD0/qORbllbgV41wiy/2ZUWZQUodqHWkjT1CwIMbQOY5jmrSYGBwwIDAQABo4IB
  JjCCASIwHwYDVR0jBBgwFoAUoBEKIz6W8Qfs4q8p74Klf9AwpLQwHQYDVR0OBBYE
  FF86fBEQfgxncWHci6O1AANn9VccMA4GA1UdDwEB/wQEAwIBhjASBgNVHRMBAf8E
  CDAGAQH/AgEAMB0GA1UdJQQWMBQGCCsGAQUFBwMBBggrBgEFBQcDAjAiBgNVHSAE
  GzAZMA0GCysGAQQBsjEBAgIxMAgGBmeBDAECATBDBgNVHR8EPDA6MDigNqA0hjJo
  dHRwOi8vY3JsLmNvbW9kb2NhLmNvbS9BQUFDZXJ0aWZpY2F0ZVNlcnZpY2VzLmNy
  bDA0BggrBgEFBQcBAQQoMCYwJAYIKwYBBQUHMAGGGGh0dHA6Ly9vY3NwLmNvbW9k
  b2NhLmNvbTANBgkqhkiG9w0BAQsFAAOCAQEAHMUom5cxIje2IiFU7mOCsBr2F6CY
  eU5cyfQ/Aep9kAXYUDuWsaT85721JxeXFYkf4D/cgNd9+hxT8ZeDOJrn+ysqR7NO
  2K9AdqTdIY2uZPKmvgHOkvH2gQD6jc05eSPOwdY/10IPvmpgUKaGOa/tyygL8Og4
  3tYyoHipMMnS4OiYKakDJny0XVuchIP7ZMKiP07Q3FIuSS4omzR77kmc75/6Q9dP
  v4wa90UCOn1j6r7WhMmX3eT3Gsdj3WMe9bYD0AFuqa6MDyjIeXq08mVGraXiw73s
  Zale8OMckn/BU3O/3aFNLHLfET2H2hT6Wb3nwxjpLIfXmSVcVd8A58XH0g==
  -----END CERTIFICATE-----
  keyData: -----BEGIN RSA PRIVATE KEY-----
  MIIEowIBAAKCAQEA1KNPn7QqrScXQYr44YIsSGwfPZNLE3y+PY++fDKeBfI9kkoU
  QARZQHI/Rj4ainlpo5uSOvV3RoP062La1/Vk5+JJS7R6xRBmlVOEG8Cmk1j0/YlQ
  IgtP1qvpR3+lINn4VQJKR9D0PZrDWkac2DrgMrVs6YPxx9JeQFjnbQoJRn9BM7m5
  w1lLvFT+VgCIMgsOsVrQfgx1S4VTji1PdfkhRIKZdK7j0Uk9yZ43xcp+GEizc2Hc
  3UxbRGxsSzKQFOFUW0etrB6twLn2SW566Uo6WG9LFcrIKLk7Hm1VnbJhfb8kqjIl
  jckYAD1rE2aCLBRXOpY7A0kRIwfV7uTRgc+ttQIDAQABAoIBAA2F9WJqyuwIOGpq
  tDljVf2lIrd/zp6GqHKx2aN8dKBcL55GJ9OKc0KuAWguOvHjltxY4IvvYI6ThdgS
  iWiCqtA0jATMjaJK0LtefGBneDCWz73wJbCEl6dHd6acb5wPQMPnSWIX3/CrDxGG
  vCLkW63d6/dN9OqIboYZIV7F7KEXCRiDBB4Fd5JmLHrxZ6ZA90BBkOmPIJ/1krS0
  zB3fVyJOdyHhYhMUdg4SEjsA0bALpAS44CRZkO5CRQ//WtVkl6QpAJyw4hFjXbU+
  CyZugRayhdsfELbvnncwzZspkOJVzZY6coRCsFSGbaM6Czu/McS5+VWx1SlzXLcQ
  38xQKGECgYEA7Jo+GqGAWrX3dIFuH7nJhxinFVF3QICT6Er1Nhow9q0d+nnGNtEd
  5N3YpOACvF8IOyJuHpsMAFdztNXckVwYaoxqFkmkOjlROQKtJwpOQs1lW4T7Dzwp
  6h2LB8eQ+gUsyDMvIJldL+MAmx49R8Q0XP4g6xO2+g5gWQQpv+mNMQ8CgYEA5hIb
  B7qQXh2ZD+dfo3RJi3PMuAi1/k69K7ydwrb2idCJ4E9kQZwFcVVoMkdfyoPbfbLs
  kLuDWPiRU5KL3vulOdqeZ19wNXpMyYXpGpkb71yg/vFOv+0H3CHZqADuDc2oejC+
  rJa+LHZAHnlbcRUtRigVH8142u7JgTCm/hzuLPsCgYEAh8p6bDRWgzk+XUpPVrv9
  MqDue+i2hXmF6eLjWvqrMVfoBbJQFXPtMUY1qWK7jzsHcVDwXHZl6+hFCvtWzMJL
  bRNLa6E2NQhiWlLz550dj29shZsLsBG6iJgODBf4V9YSfpAJsy7x8aLZ3Sz8xKyR
  1PExGVnGQTtxBoXCJFe5ZfcCgYBVdQ4zPboYK1hKTv/4P959fQLirOGk12xuzX2v
  8LP8lshP2E1+DUz8PuQYIOjU2UtzEj3KuMveBV49s6ZeqgxCRBEohouwYYAaLrJa
  HdsBet+WMt20bn/H5Y7qV4YU/HoDAQ4iH0/+ReIlL6CmjV4mvAa0rGais6WHZiHx
  K5/QdwKBgHdLTZSNcKn+a+hO3NG2eWjpN+KGqsmb6760sY/2sArPx7FerLT8x9tv
  DXezSHkTbnHblWGa3ng2isC+TVf0BenRx5y2YgL4ED6QgxGpkgF+S5Iam1BpHRR4
  H74HIA/602OW0O3srV4SdpIkI/fOM8hNhh5Cy/w4GOVjE0o14GRy
  -----END RSA PRIVATE KEY-----

```

然后可以根据需要修改 values.yaml 中的配置项，例如存储卷的配置、管理员密码等等。

以下是一些常用的配置项：

jenkinsAdminPassword: Jenkins 管理员密码。
persistence.enabled: 是否启用持久化存储。
persistence.size: 持久化存储卷的大小。
persistence.storageClass: 存储卷的 StorageClass。
persistence.mountPath: 持久化存储卷挂载的路径。
service.type: Jenkins Service 的类型，可以设置为 ClusterIP、NodePort 或 LoadBalancer。
ingress.enabled: 是否启用 Ingress。
ingress.hosts: Ingress 的域名列表。

```
helm install jenkins jenkins/jenkins -f values.yaml
```

等待安装完成后，可以使用以下命令查看 Jenkins 的状态：

```
kubectl get pods -l "component=jenkins-master"
```

Jenkins 会自动创建一个初始管理员用户，可以使用管理员用户名 admin 和密码来登录 Jenkins。同时，可以使用 kubectl port-forward 命令将 Jenkins Service 暴露出来，以便访问 Jenkins Web UI：

```
kubectl port-forward svc/jenkins 8080:8080
```

##### 方案二基于 ECS 服务器构建 Jenkins
