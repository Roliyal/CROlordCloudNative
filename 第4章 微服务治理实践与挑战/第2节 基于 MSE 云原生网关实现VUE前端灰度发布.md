# 基于 MSE 云原生网关实现前端灰度发布

### 目录

1. [概述](#1-概述)
2. [前端灰度发布的基本概念](#2-前端灰度发布的基本概念)
3. [MSE 云原生网关的配置](#3-MSE-云原生网关的配置)
4. [前端应用部署及灰度策略](#4-前端应用部署及灰度策略)
5. [Cookie 传递及灰度流量路由](#5-Cookie-传递及灰度流量路由)
6. [YAML 配置细节](#6-YAML-配置细节)
7. [网关配置与插件](#7-网关配置与插件)
8. [注意事项](#8-注意事项)
9. [总结](#9-总结)

---

### 1. 概述

本文介绍如何通过 **MSE 云原生网关** 实现前端灰度发布。前端灰度发布是一种流量控制策略，它能够在新版本的前端应用发布时，通过将部分用户的流量引导到新版本，确保版本更新的平稳过渡。通过结合 MSE 云原生网关和前端的灰度策略，我们能够实现更加灵活的流量管理和版本控制。

### 2. 前端灰度发布的基本概念

前端灰度发布的主要目标是将应用的新版本仅暴露给部分用户，确保新版本稳定后再向全体用户发布。其基本概念包括：

- **灰度流量控制**：将一部分用户的流量路由到新版本的应用。
- **平滑过渡**：通过分阶段发布，避免全量发布时发生系统崩溃。
- **用户标识**：通过用户唯一标识（如 `userId`）来判断哪些用户将接入灰度版本。

### 3. MSE 云原生网关的配置

微服务引擎MSE（Microservices Engine）是一个面向业界主流开源微服务生态的一站式微服务平台，提供注册配置中心（原生支持Nacos、ZooKeeper、Eureka）、云原生网关（原生支持Ingress、Envoy）、微服务治理（原生支持Spring Cloud、Dubbo、Sentinel、遵循OpenSergo服务治理规范）、分布式任务调度（兼容开源XXL-JOB、ElasticJob、Spring Schedule）的能力。


#### 步骤 1：创建和配置 MSE 云原生网关

在 MSE 控制台创建并配置云原生网关，设置流量路由规则和灰度发布策略。以下是创建网关的步骤：

1. 登录 MSE 控制台，创建一个新的云原生网关。
2. 配置网关的路由规则，使用 `frontend-gray`[frontend-gray下载](https://github.com/alibaba/higress/tree/main/plugins/wasm-go/extensions/frontend-gray) 插件，按照用户的唯一标识（如 `userId`）进行流量路由。

#### 步骤 2：安装 `frontend-gray` 插件

1. 在 MSE 网关控制台中选择 "插件市场"，然后选择 `frontend-gray` 插件。
2. 点击 `创建插件`，上传编译好的插件文件 `frontend-gray.wasm`，并配置插件参数。

```yaml
grayKey: X-User-ID  # 从 Cookie 中提取用户ID
rules:
   - name: gray-user
     grayKeyValue:
        - '000002'  # 白名单：用户ID等于 000020

baseDeployment:
   version: base  # 默认版本
grayDeployments:
   - name: gray-user
     version: gray
     enabled: true    
```

#### 步骤 3：配置网关路由规则

网关路由规则配置如下：

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
   annotations:
      nginx.ingress.kubernetes.io/canary: 'true'
      nginx.ingress.kubernetes.io/canary-by-header: x-higress-tag
      nginx.ingress.kubernetes.io/canary-by-header-value: gray   
  labels:
    ingress-controller: mse
  namespace: crolord
  name: frontend-base-ingress
spec:
   ingressClassName: mse
   rules:
      - host: micro.roliyal.com
        http:
           paths:
              - backend:
                   service:
                      name: micro-vue-front-service
                      port:
                         number: 80
                path: /
                pathType: ImplementationSpecific
   rules:
      - host: micro.roliyal.com
        http:
           paths:
              - backend:
                   service:
                      name: micro-vue-front-canary-service
                      port:
                         number: 80
                path: /
                pathType: Prefix
```

在上述配置中，前端流量将根据请求中的 `X-User-ID` 路由到相应的灰度或主版本服务。

### 4. 前端应用部署及灰度策略

#### 步骤 1：前端应用的灰度发布

前端应用部署涉及将用户请求通过 MSE 云原生网关转发到不同的后端版本。具体流程如下：

1. 用户访问前端应用时，首先通过云原生网关进行流量分配。
2. 如果用户属于灰度用户（根据 `userId` 和灰度标识），流量将路由到灰度版本的后端应用。

前端代码中通过在登录时设置灰度标识，具体实现如下：

```javascript
// 在登录时设置灰度标识
document.cookie = `X-User-ID=${userId}; path=/;`;
document.cookie = `x-pre-higress-tag=gray; path=/;`;
```

#### 步骤 2：前端应用的灰度控制

在登录或注册接口中，前端会设置用户的 `userId` 和 `authToken`，并将这些信息存储在 `localStorage` 和 Cookie 中。后端接收到这些信息后，会根据用户的标识判断是否将其流量路由到灰度版本。以下是前端的请求设置：

```javascript
// 在请求时设置用户标识和灰度标签
axiosInstance.interceptors.request.use(config => {
    const userId = store.state.userId || localStorage.getItem('userId');
    const authToken = store.state.authToken || localStorage.getItem('authToken');

    if (userId) {
        deleteCookie('X-User-ID');
        document.cookie = `X-User-ID=${userId}; path=/;`;
    }

    deleteCookie('x-pre-higress-tag');
    document.cookie = `x-pre-higress-tag=gray; path=/;`;

    if (authToken) {
        config.headers['Authorization'] = authToken;
    }

    return config;
}, error => Promise.reject(error));
```

### 5. Cookie 传递及灰度流量路由

#### 步骤 1：通过 Cookie 实现前端灰度流量路由

前端应用会通过 **Cookie** 和 **请求头** 传递用户的灰度标识。具体操作如下：

1. 用户登录时，`X-User-ID` 和 `x-pre-higress-tag` 会被存储到 Cookie 中。
2. 每次用户请求时，前端应用会检查 Cookie 中的标识，并将其传递到后端服务。
3. 云原生网关根据这些标识决定是否将用户流量路由到灰度版本。

#### 步骤 2：灰度流量的路由实现

通过 MSE 网关插件，前端灰度流量可以根据用户的标识进行动态路由。MSE 网关会根据配置的规则（如 `userid`）判断是否将流量路由到灰度版本。

### 6. YAML 配置细节

#### 服务与应用部署配置

```yaml
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
   name: app-vue-front-pvc
   namespace: crolord
spec:
   accessModes:
      - ReadWriteOnce
   storageClassName: crolord-cnfs-uat-nas-tlv4v # 使用初始化配置CNFS文件系统
   resources:
      requests:
         storage: 20Gi
---
apiVersion: apps/v1
kind: Deployment
metadata:
   name: micro-vue-front
   namespace: crolord
   labels:
      app: micro-vue-front
spec:
   replicas: 2
   strategy:
      type: RollingUpdate
      rollingUpdate:
         maxSurge: 1
         maxUnavailable: 0
   selector:
      matchLabels:
         app: micro-vue-front
   template:
      metadata:
         labels:
            app: micro-vue-front
      spec:
         terminationGracePeriodSeconds: 30
         securityContext:
            fsGroup: 101
         containers:
            - name: micro-vue-front
              image: crolord-uat-registry-vpc.cn-hongkong.cr.aliyuncs.com/micro/micro-front-uat:1.0.17
              imagePullPolicy: IfNotPresent
              ports:
                 - containerPort: 80
              resources:
                 requests:
                    cpu: "250m"
                    memory: "128Mi"
                 limits:
                    cpu: "1"
                    memory: "512Mi"
              lifecycle:
                 postStart:
                    exec:
                       command:
                          - /bin/sh
                          - -c
                          - |
                             echo "$(date) postStart hook" >> /var/log/nginx/hook.log
                 preStop:
                    exec:
                       command:
                          - /bin/sh
                          - -c
                          - |
                             echo "$(date) preStop graceful quit" >> /var/log/nginx/hook.log
                             nginx -s quit
                             sleep 10
              livenessProbe:
                 httpGet:
                    path: /healthz
                    port: 80
                 initialDelaySeconds: 30
                 periodSeconds: 20
                 failureThreshold: 3
              readinessProbe:
                 httpGet:
                    path: /healthz
                    port: 80
                 initialDelaySeconds: 10
                 periodSeconds: 10
                 failureThreshold: 3
              env:
                 - name: TZ
                   value: Asia/Shanghai
              volumeMounts:
                 - name: nginx-config
                   mountPath: /etc/nginx/conf.d
                 - name: app-vue-front-storage
                   mountPath: /usr/share/nginx/html/app-vue-front-storage
                 - name: nginx-logs
                   mountPath: /var/log/nginx
         volumes:
            - name: nginx-config
              configMap:
                 name: micro-vue-front-nginx-config
            - name: app-vue-front-storage
              persistentVolumeClaim:
                 claimName: app-vue-front-pvc
            - name: nginx-logs
              emptyDir: {}

```

#### Service 配置

```yaml
apiVersion: v1
kind: Service
metadata:
   name: micro-vue-front-service
   namespace: crolord
spec:
   selector:
      app: micro-vue-front
   ports:
      - protocol: TCP
        port: 80
        targetPort: 80
   type: ClusterIP
```

### 7. 网关配置与插件

在 MSE 云原生网关中创建插件时，配置如下规则以实现灰度发布：

```yaml
grayKey: X-User-ID  # 从 Cookie 中提取用户ID
rules:
   - name: gray-user
     grayKeyValue:
        - '000002'  # 白名单：用户ID等于 000020

baseDeployment:
   version: base  # 默认版本
grayDeployments:
   - name: gray-user
     version: gray
     enabled: true    
```

该配置通过判断用户的 `X-User-ID`，将部分用户的流量路由到灰度版本。

### 8. 注意事项

在实现前端灰度发布时，以下是一些需要注意的事项：

1. **用户标识（userId）管理**：
    - 确保前端应用正确地传递用户标识。灰度流量的控制依赖于用户标识的正确传递和存储。务必通过 Cookie 或请求头传递 `X-User-ID` 和 `x-pre-higress-tag`。
    - 登录后将 `userId` 和 `authToken` 存储在 `localStorage` 和 Cookie 中，以便后续请求时使用。

2. **灰度流量的准确控制**：
    - 配置网关时，要确保灰度规则配置正确。灰度发布的关键在于通过 `userId` 等标识判断流量应该路由到哪个版本的应用。
    - 在部署新版本时，确保旧版本和灰度版本的兼容性，以避免出现功能故障或数据丢失。

3. **前端和后端配合**：
    - 前端需要根据网关路由规则设置正确的灰度标签，并且根据请求头和 Cookie 的信息判断是否启用灰度版本。
    - 后端应用需要根据传递的 `X-User-ID` 和 `x-pre-higress-tag` 决定是否处理灰度请求，并且在必要时提供灰度版本的接口。

4. **调试和日志**：
    - 在开发过程中，建议开启详细的日志记录功能，以便追踪请求的流向和问题。
    - 使用浏览器的开发者工具（如 Chrome DevTools）检查请求头和 Cookie 信息，确保灰度标识正确传递。

### 9. 总结

通过结合 MSE 云原生网关的强大流量控制能力与前端灰度发布策略，我们能够实现高效且可控的前端灰度发布。在此过程中，前端通过 Cookie 和请求头传递灰度标识，云原生网关根据这些标识判断是否将流量路由到灰度版本，确保新版本发布的平滑过渡。

---