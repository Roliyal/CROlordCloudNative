# ACK ingress 选型配置

### ingress应用场景：
&emsp;&emsp;在K8S中，Kubernetes 集群内的网络环境与外部是隔离的，也就是说 Kubernetes 集群外部的客户端无法直接访问到集群内部的服务，Kubernetes通过 Kube-proxy服务实现了Service的对外发布及负载均衡。 但是在实际的互联网应用场景中，不仅要实现单纯的转发，还有更为细致的策略需求，基于该需求，Kubernetes引入资源对象ingress，我们会从多个角度对比Nginx ingress，ALB ingress，MSE ingress ，同时介绍如何配置这三个ingress对象，我们希望这篇文章能够帮助您选择到更贴合您业务场景的ingress对象，并且能快速方便的使用它。


| 比较项    | Nginx ingress                                                                                                                 | ALB ingress                                                                                                              | MSE ingress                                                                                                                                    |
|--------|-------------------------------------------------------------------------------------------------------------------------------|--------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------------------------------------------------------------------------------------------|
| 特点     | 使用Nginx作为Ingress Controller，提供了丰富的功能，例如请求限速、基于主机名、路径或请求方法的路由、SSL/TLS终止等。它是Kubernetes社区最受欢迎的Ingress Controller之一，已经广泛应用于生产环境中。 | 是阿里云提供的一种Ingress Controller，它使用ALB作为Ingress Controller，可以利用ALB服务的高可用性和强大的负载均衡功能。ALB Ingress支持HTTP和HTTPS流量的路由，以及基于主机名、路径或请求方法的路由。 | 是阿里云云提供的一种Ingress Controller，它使用微服务引擎MSE网关作为Ingress Controller，可以提供HTTP、HTTPS和GRPC流量的路由。MSE Ingress支持基于主机名、路径、请求方法、请求头等多种条件的路由，同时也支持灰度发布和限流等功能。 | 
| 典型应用场景 | 网关高度定制化；云原生应用金丝雀发布、蓝绿发布。                                                                                                      | 网关全托管、免运维；互联网应用七层高性能自动弹性；云原生应用金丝雀发布、蓝绿发布；超大QPS、超大并发连接。                                                                   | 网关全托管、免运维；南北向、东西向流量统一管理，微服务网关，全链路灰度；多个容器集群、PaaS平台、ECS服务共用一个网关实例；混合云、多数据中心、多业务域的内部互通；认证鉴权，灵活设置，安全防护要求高；超大流量、高并发业务。                              |
| 架构     | 基于Nginx+Lua插件扩展。                                                                                                              | 基于阿里洛神云网络平台；基于CyberStar自研平台，支持自动弹性伸缩。                                                                                    | 基于开源Higress项目，控制面使用Istiod，数据面使用Envoy                                                                                                           |用户独享实例。|
| 支持协议   | 支持HTTP和HTTPS协议；支持WebSocket、WSS和gRPC协议。                                                                                        | 支持HT支持HTTP和HTTPS协议；支持WebSocket、WSS和gRPC协议。                                                                               | 支持HTTP和HTTPS协议；支持WebSocket、WSS和gRPC协议；支持HTTP/HTTPS转Dubbo协议。                                                                                    |
| 认证鉴权   | 支持Basic Auth认证方式；支持oAuth协议。                                                                                                   | 	支持TLS身份认证。                                                                                                              | 支持Basic Auth、oAuth、JWT、OIDC认证；集成阿里云IDaaS；支持自定义认证。                                                                                              |
| 运维能力   | 自行维护组件；通过配置HPA进行扩缩容；需要主动配置规格进行调优。                                                                                             | 全托管、免运维；自动弹性，免配置支持超大容量；处理能力随业务峰值自动伸缩。                                                                                    | 全托管，免运维。                                                                                                                                       |
| 安全     | 支持HTTPS协议；支持黑白名单功能。                                                                                                           | HTTPS（集成SSL）支持全链路HTTPS、SNI多证书、RSA、ECC双证、TLS 1.3协议和TLS算法套件选择；支持WAF防护（对接阿里云Web防火墙）；支持DDos防护（对接阿里云DDoS防护服务）；支持黑白名单功能。       | 支持全链路HTTPS、SNI多证书（集成SSL），可配置TLS版本；支持路由级WAF防护（对接阿里云Web防火墙）；支持路由级黑白名单功能。                                                                         |
| 服务治理   | 服务发现支持K8s；服务灰度支持金丝雀发布；服务高可用支持限流。                                                                                              | 服务发现支持K8s；服务灰度支持金丝雀发布；服务高可用支持限流。                                                                                         | 服务发现支持K8s、Nacos、ZooKeeper、EDAS、SAE、DNS、固定IP；支持2个以上版本的金丝雀发布、标签路由，与MSE服务治理结合可实现全链路灰度发布；内置集成MSE服务治理中的Sentinel，支持限流、熔断、降级；服务测试支持服务Mock。            |
| 扩展性    | 使用Lua脚本。| 使用AScript自研脚本。| 使用Wasm插件，实现多语言编写；使用Lua脚本。|                                                                                                                        


#### Nginx ingress：
创建Nginx Ingress Controller服务  

方式一：  
创建ACK集群时，组件选择设置是否安装Ingress组件。默认Nginx Ingress。  
若选择Nginx Ingress，则自动安装nginx Ingress组件。  

方式二：  
若创建集群时未选择，则需要手动创建Nginx Ingress Controller  
通过heml安装Nginx Ingress Controller：
下载并解压Chart压缩包
~~~shell
$ wget https://aliacs-app-catalog.oss-cn-hangzhou.aliyuncs.com/charts-incubator/ack-ingress-nginx-v1-4.0.16.tgz
$ tar -zxvf ack-ingress-nginx-v1-4.0.16.tgz
~~~
修改values.yaml文件  

|参数| 描述                                               |
|--------|--------------------------------------------------|
|controller.image.repository| ingress-nginx镜像地址。                               |
|controller.image.tag| ingress-nginx镜像版本                                |
|controller.ingressClassResource.name| 设置Ingress Controller所对应的IngressClass的名称          |
|controller.ingressClassResource.controllerValue| controller.ingressClassResource.controllerValue  |
|controller.replicaCount| 设置该Ingress Controller Pod的副本数                    |
|controller.service.external.enabled| 	是否开启公网SLB访问，不需要开启则设置为false                      |
|controller.service.internal.enabled| 	是否开启私网SLB访问，需要开启则设置为true                        |
|controller.kind| 设置IngressController部署形态，可选值：Deployment和DaemonSet |
|service.beta.kubernetes.io/alibaba-cloud-loadbalancer-id| 如果需要使用已有的SLB，添加该参数，配置已有SLB的实例ID                  |
|service.beta.kubernetes.io/alibaba-cloud-loadbalancer-force-override-listeners| true表示强制覆盖SLB监听                                  |
values.yaml参考：
>CROlord-CloudNative\第2章 阿里云 kubernentes 应用部署\values.yaml  

修改完values.yaml后打包helm chart包：
~~~shell
$ helm package ack-ingress-nginx-v1
~~~
执行结果参考：
~~~shell
$ ls
ack-ingress-nginx-v1   ack-ingress-nginx-v1-4.0.16.tgz
~~~
使用helm创建nginx-ingress-controller:
~~~shell
$ helm install ack-ingress-nginx-v1 ack-ingress-nginx-v1-4.0.16.tgz

NAME: ack-ingress-nginx-v1
LAST DEPLOYED: Thu Apr  6 16:06:46 2023
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE: None
NOTES:
The ingress-nginx controller has been installed.
It may take a few minutes for the LoadBalancer IP to be available.
You can watch the status by running 'kubectl --namespace default get services -o wide -w ack-ingress-nginx-v1-controller'

An example Ingress that makes use of the controller:
  apiVersion: networking.k8s.io/v1
  kind: Ingress
  metadata:
    name: example
    namespace: foo
    annotations:
      kubernetes.io/ingress.class: 
  spec:
    rules:
      - host: www.example.com
        http:
          paths:
            - backend:
                service:
                  name: exampleService
                  port:
                    number: 80
              path: /
    # This section is only required if TLS is to be enabled for the Ingress
    tls:
      - hosts:
        - www.example.com
        secretName: example-tls

If TLS is enabled for the Ingress, a Secret containing the certificate and key must also be provided:

  apiVersion: v1
  kind: Secret
  metadata:
    name: example-tls
    namespace: foo
  data:
    tls.crt: <base64 encoded cert>
    tls.key: <base64 encoded key>
  type: kubernetes.io/tls
~~~
查看ack-ingress-nginx-v1-controller：
~~~shell
$ kubectl describe  svc ack-ingress-nginx-v1-controller
Name:                     ack-ingress-nginx-v1-controller
Namespace:                default
Labels:                   app.kubernetes.io/component=controller
                          app.kubernetes.io/instance=ack-ingress-nginx-v1
                          app.kubernetes.io/managed-by=Helm
                          app.kubernetes.io/name=ack-ingress-nginx-v1
                          app.kubernetes.io/version=v1.2.1-aliyun.1
                          helm.sh/chart=ack-ingress-nginx-v1-4.0.16
                          service.beta.kubernetes.io/hash=aa40771efacafee16369261d8f726a4a2ddb3973cdc693efc9343a91
                          service.k8s.alibaba/loadbalancer-id=lb-j6cd99urbi9fvj4gab69m
Annotations:              meta.helm.sh/release-name: ack-ingress-nginx-v1
                          meta.helm.sh/release-namespace: default
                          service.beta.kubernetes.io/alibaba-cloud-loadbalancer-force-override-listeners: true
                          service.beta.kubernetes.io/alibaba-cloud-loadbalancer-id: lb-j6cd99urbi9fvj4gab69m
Selector:                 app.kubernetes.io/component=controller,app.kubernetes.io/instance=ack-ingress-nginx-v1,app.kubernetes.io/name=ack-ingress-nginx-v1
Type:                     LoadBalancer
IP Family Policy:         SingleStack
IP Families:              IPv4
IP:                       192.168.197.246
IPs:                      192.168.197.246
LoadBalancer Ingress:     10.0.0.110
Port:                     http  80/TCP
TargetPort:               80/TCP
NodePort:                 http  32412/TCP
Endpoints:                10.0.1.106:80,10.0.1.107:80
Port:                     https  443/TCP
TargetPort:               443/TCP
NodePort:                 https  30532/TCP
Endpoints:                10.0.1.106:443,10.0.1.107:443
Session Affinity:         None
External Traffic Policy:  Local
HealthCheck NodePort:     30387
Events:
  Type    Reason                   Age                From                Message
  ----    ------                   ----               ----                -------
  Normal  UnAvailableLoadBalancer  31s (x4 over 34s)  service-controller  There are no available nodes for LoadBalancer
  Normal  EnsuredLoadBalancer      16s (x4 over 31s)  service-controller  Ensured load balancer [lb-j6cd99urbi9fvj4gab69m]
~~~
