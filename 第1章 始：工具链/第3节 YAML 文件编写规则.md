### 什么是 yaml

YAML (YAML Ain't Markup Language)是一种轻量级的文本格式，用于表示数据。YAML 文件通常用于配置文件和数据序列化，它使用缩进和换行符来表示层次结构数据。

YAML 文件可以包含以下内容：

- 简单数据类型，如字符串、整数、浮点数和布尔值。
- 列表，表示一组值。
- 映射，表示一组键值对。
- 注释，用于解释和说明 YAML 文件中的内容。
  YAML 文件的格式非常简洁和易读，它不像其他标记语言（如 XML 或 JSON）那样需要使用标签或符号。因此，它在配置文件和数据序列化方面非常受欢迎，被广泛应用于各种编程语言和应用程序中。


| 名称       | 描述                                     | 示例                                                                                                          |
| ---------- | ---------------------------------------- | ------------------------------------------------------------------------------------------------------------- |
| HTML       | 超文本标记语言，用于创建网页             | `<html><head><title>My Page</title></head><body><h1>Hello, world!</h1></body></html>`                         |
| CSS        | 层叠样式表，用于美化网页                 | `body { background-color: #f2f2f2; }`                                                                         |
| JavaScript | 脚本语言，用于实现交互和动态效果         | `var x = 5; console.log(x);`                                                                                  |
| Python     | 面向对象的编程语言，用于开发各种应用程序 | `print("Hello, world!")`                                                                                      |
| Java       | 面向对象的编程语言，用于开发大型应用程序 | `public class HelloWorld { public static void main(String[] args) { System.out.println("Hello, world!"); } }` |
| SQL        | 结构化查询语言，用于管理关系型数据库     | `SELECT * FROM customers WHERE city = 'New York';`                                                            |
| YAML       | 轻量级的文本格式，用于表示数据           | `name: John Smith\nage: 30\ncity: New York`                                                                   |
| JSON       | 轻量级的文本格式，用于表示数据           | `{"name": "John Smith", "age": 30, "city": "New York"}`                                                       |
| XML        | 可扩展标记语言，用于表示结构化数据       | `<book><title>My Book</title><author>John Smith</author><price>10.99</price></book>`                          |

以上是一个简单的常见语言示例，它列出了几种编程语言和标记语言的名称、描述和示例。

综上所述，Kubernetes（k8s）使用YAML作为配置文件的格式，因为YAML语言简洁易读，具有良好的可读性和可维护性。它是一种人类可读的数据序列化格式，相比于其他类似的格式，如JSON或XML，YAML更加易于阅读和编写。
在Kubernetes中，YAML文件包含了对于应用程序和其它资源的定义。通过这些YAML文件，用户可以定义应用程序所需要的Pod、Service、Deployment等资源，从而使Kubernetes能够自动管理这些资源的生命周期。此外，YAML文件还能够定义应用程序的环境变量、卷挂载、端口映射等信息。
总的来说，Kubernetes选择使用YAML作为配置文件的格式，主要是因为YAML语言简洁易读，同时也能够满足Kubernetes需要的定义应用程序和资源的需求。

### 经典架构 LNMP yaml示例

LNMP 架构之所以被称为经典架构，是因为它是一种被广泛应用于构建 Web 应用程序的经典架构之一。LNMP 架构由 Linux、Nginx、MySQL 和 PHP 四个开源组件组成，每个组件都具有一定的特点和优点：

- Linux 操作系统是开源的、免费的，并且具有高度稳定性和可靠性。它是 LNMP 架构的基石，提供了稳定的运行环境。
- Nginx 是一个高性能的 Web 服务器和反向代理服务器，具有低内存占用、高并发能力、可扩展性好等特点，适合于处理大量静态资源请求和高并发连接请求。
- MySQL 是一种开源的关系型数据库管理系统（RDBMS），具有高度稳定性、可靠性和安全性，并且能够处理大量的事务请求。
- PHP 是一种流行的开源脚本语言，具有简单易用、快速开发、可移植性好等特点，适合于开发动态 Web 应用程序。
  这四个组件的结合使得 LNMP 架构具有高度稳定性、可靠性、安全性和高性能的特点，适合于构建大规模的 Web 应用程序和服务。由于 LNMP 架构已经存在多年，并被广泛应用于各种 Web 应用程序中，因此它被称为经典架构。

1. 部署 LNMP 资源清单

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx
spec:
  type: LoadBalancer
  ports:
    - name: http
      port: 80
      targetPort: 80
  selector:
    app: nginx
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx
          ports:
            - containerPort: 80
          volumeMounts:
            - name: nginx-conf
              mountPath: /etc/nginx/conf.d
        - name: php-fpm
          image: php-fpm
          ports:
            - containerPort: 9000
          volumeMounts:
            - name: php-fpm-conf
              mountPath: /usr/local/etc/php-fpm.d
          env:
            - name: MYSQL_HOST
              value: mysql
            - name: MYSQL_DATABASE
              value: example_db
            - name: MYSQL_USER
              value: db_user
            - name: MYSQL_PASSWORD
              value: db_password
        - name: mysql
          image: mysql:5.7
          ports:
            - containerPort: 3306
          env:
            - name: MYSQL_ROOT_PASSWORD
              value: root_password
            - name: MYSQL_DATABASE
              value: example_db
            - name: MYSQL_USER
              value: db_user
            - name: MYSQL_PASSWORD
              value: db_password
          volumeMounts:
            - name: mysql-data
              mountPath: /var/lib/mysql
      volumes:
        - name: nginx-conf
          configMap:
            name: nginx-config
        - name: php-fpm-conf
          configMap:
            name: php-fpm-config
        - name: mysql-data
          persistentVolumeClaim:
            claimName: mysql-pvc
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mysql-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: nginx-config
data:
  default.conf: |
    server {
      listen 80;
      server_name localhost;
      location / {
        root /var/www/html;
        index index.php;
        try_files $uri $uri/ /index.php?$query_string;
      }
      location ~ \.php$ {
        fastcgi_pass php-fpm:9000;
        fastcgi_index index.php;
        fastcgi_param SCRIPT_FILENAME $document_root$fastcgi_script_name;
        include fastcgi_params;
      }
    }
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: php-fpm-config
data:
  www.conf: |
    [www]
    user = www-data
    group = www-data
    listen = 0.0.0.0:9000
    listen.owner = www-data
    listen.group = www-data
    chdir = /var/www
    env[MYSQ_HOST] = $MYSQL_HOST
    env[MYSQL_DATABASE] = $MYSQL_DATABASE
    env[MYSQL_USER] = $MYSQL_USER
    env[MYSQL_PASSWORD] = $MYSQL_PASSWORD
    php_admin_value[error_log] = /var/log/fpm-php.www.log

```

该配置示例包括了以下资源：

- 一个 Nginx Service，类型为 LoadBalancer，负责将流量引导到 Nginx Deployment
- 一个 Nginx Deployment，该 Deployment 包含一个 Nginx 容器和一个 PHP-FPM 容器。Nginx 容器将处理 HTTP 请求，并将 PHP 请求传递给 PHP-FPM 容器。PHP-FPM 容器将运行 PHP 代码并与 MySQL 数据库进行通信
- 一个 MySQL Deployment，该 Deployment 包含一个 MySQL 容器，用于存储应用程序数据
- 一个 PersistentVolumeClaim（PVC），用于将 MySQL 数据存储在持久卷中

此外，该配置示例还包括两个 ConfigMap，一个用于存储 Nginx 配置，另一个用于存储 PHP-FPM 配置。

请注意，此配置示例可能不适用于所有情况，并且可能需要进行适当的修改(存储卷、LB)才能满足您的特定需求，当然您也可以将以上资源清单做对应的拆分。

### YAML 资源清单优化方向

在 Kubernetes 中如何优化 YAML 配置文件：


| 配置项         | 描述                                                                                                         | 示例                                                                                                                                                             |
| -------------- | ------------------------------------------------------------------------------------------------------------ | ---------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| 环境变量       | 使用环境变量可以使 YAML 文件更加简洁和易于维护。                                                             | `env: - name: DB_HOST value: mysql`                                                                                                                              |
| ConfigMap      | 使用 ConfigMap 可以将应用程序的配置信息与 YAML 文件分离开来，从而使得配置更加灵活和可维护。                  | `envFrom: - configMapRef: name: app-config`                                                                                                                      |
| Secrets        | 使用 Secrets 可以将敏感信息与 YAML 文件分离开来，从而提高安全性。                                            | `envFrom: - secretRef: name: app-secrets`                                                                                                                        |
| 资源限制       | 在 YAML 文件中设置资源限制可以确保应用程序不会消耗过多的资源，并且可以提高应用程序的稳定性和可靠性。         | `resources: limits: memory: "1Gi" cpu: "1"`                                                                                                                      |
| liveness 探针  | 在 YAML 文件中设置 liveness 探针可以确保容器在正常运行和准备好处理流量之前，先进行一些必要的检查和准备工作。 | `livenessProbe: httpGet: path: /healthz port: 8080`                                                                                                              |
| readiness 探针 | 在 YAML 文件中设置 readiness 探针可以确保容器已准备好处理流量。                                              | `readinessProbe: httpGet: path: /readyz port: 8080`                                                                                                              |
| Deployment     | 使用 Deployment 可以确保应用程序始终处于所需的副本数，并且可以进行滚动更新。                                 | `apiVersion: apps/v1 kind: Deployment metadata: name: app-deployment spec: replicas: 3 template: spec: containers: - name: app image: myapp:latest`              |
| StatefulSet    | 使用 StatefulSet 可以确保有状态应用程序的稳定性和顺序性。                                                    | `apiVersion: apps/v1 kind: StatefulSet metadata: name: mysql spec: serviceName: mysql replicas: 3 template: spec: containers: - name: mysql image: mysql:latest` |

注意，以上示例仅供参考，实际情况可能会因应用程序的不同而有所变化，接下来我们实际应用上面LNMP示例。

### 1. 更加适用的LNMP资源清单

```yaml
apiVersion: v1
kind: Service
metadata:
  name: nginx-service
spec:
  selector:
    app: nginx
  ports:
    - protocol: TCP
      port: 80
      targetPort: 80
---
apiVersion: v1
kind: ConfigMap
metadata:
  name: app-config
data:
  DB_HOST: mysql
  DB_PORT: "3306"
  DB_NAME: mydatabase
---
apiVersion: v1
kind: Secret
metadata:
  name: app-secrets
data:
  DB_PASSWORD: dGVzdA==
---
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx-deployment
spec:
  replicas: 3
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
        - name: nginx
          image: nginx:alpine
          ports:
            - containerPort: 80
          envFrom:
            - configMapRef:
                name: app-config
            - secretRef:
                name: app-secrets
          livenessProbe:
            httpGet:
              path: /
              port: 80
          readinessProbe:
            httpGet:
              path: /
              port: 80
          resources:
            limits:
              memory: "256Mi"
              cpu: "0.1"
---
apiVersion: v1
kind: Service
metadata:
  name: mysql-service
spec:
  selector:
    app: mysql
  ports:
    - protocol: TCP
      port: 3306
      targetPort: 3306
---
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: mysql-pvc
spec:
  accessModes:
    - ReadWriteOnce
  resources:
    requests:
      storage: 1Gi
---
apiVersion: apps/v1
kind: StatefulSet
metadata:
  name: mysql-statefulset
spec:
  replicas: 1
  serviceName: mysql-service
  selector:
    matchLabels:
      app: mysql
  template:
    metadata:
      labels:
        app: mysql
    spec:
      containers:
        - name: mysql
          image: mysql:5.7
          ports:
            - containerPort: 3306
          envFrom:
            - configMapRef:
                name: app-config
            - secretRef:
                name: app-secrets
          livenessProbe:
            exec:
              command: ["mysqladmin", "ping", "-h", "127.0.0.1"]
            initialDelaySeconds: 30
            periodSeconds: 10
            timeoutSeconds: 5
          readinessProbe:
            exec:
              command: ["mysql", "-h", "127.0.0.1", "-e", "SELECT 1"]
            initialDelaySeconds: 5
            periodSeconds: 10
            timeoutSeconds: 1
          resources:
            limits:
              memory: "512Mi"
              cpu: "0.5"
          volumeMounts:
            - name: mysql-pvc
              mountPath: /var/lib/mysql
 
```

这个 YAML 文件相对于之前的版本，做出了以下优化：

- Nginx 容器使用了更轻量的 Alpine 版本，减小了容器镜像的大小，从而提高了部署效率
- Nginx 容器的资源配额也进行了调整
- 将 MySQL 的配置信息和密码放到了 ConfigMap 和 Secret 中，分别进行管理
- 为 MySQL 数据库配置了一个 Persistent Volume Claim，从而实现了数据的持久化存储，避免了数据的丢失
- MySQL 容器的 livenessProbe 和 readinessProbe 也进行了调整，从而更加准确地检测容器的健康状态
  这些优化措施可以提高应用程序的可靠性和性能，使应用程序更加稳定和高效地运行在 Kubernetes 平台上。

### 2. 应用的可用性

在上述的健康检查以及存活就绪检查中，我们也可以添加 preStop 和 preStart 确保应用的可用性，具体可以看下面的示例。

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: my-container
      image: nginx
      lifecycle:
        preStop:
          exec:
            command: ["/bin/sh","-c","nginx -s quit && sleep 5s"]
        preStart:
          exec:
            command: ["/bin/sh","-c","echo Starting nginx..."]

```

在这个示例 YAML 文件中，我们为一个名为 my-pod 的 Pod 配置了一个名为 my-container 的容器。这个容器使用了 nginx 镜像，并且在 lifecycle 字段下配置了 preStop 和 preStart。

- preStop：在 Pod 被停止之前执行的命令。在这个示例中，我们使用了 exec 类型的钩子，并且在 command 字段中指定了要执行的命令。具体来说，我们使用 /bin/sh 命令打开一个 Shell，然后执行了 nginx -s quit 命令停止 Nginx 服务，并且在停止之前等待 5 秒钟，以确保所有请求都被处理完毕。
- preStart：在容器启动之前执行的命令。在这个示例中，我们同样使用了 exec 类型的钩子，并且在 command 字段中指定了要执行的命令。具体来说，我们使用 /bin/sh 命令打开一个 Shell，并且在 Shell 中执行了一个 echo 命令，输出了一条提示信息。

这些钩子可以用于在容器的生命周期中执行一些特定的操作，例如在容器启动之前或停止之后执行一些脚本或命令。这些操作可以帮助我们更好地管理容器，提高容器的可靠性和稳定性。
在 YAML 文件中配置这些钩子时，需要注意以下几点：
钩子的定义需要在容器的 lifecycle 字段下；

- 钩子的类型为 exec 时，需要在 command 字段中指定要执行的命令；
- 钩子的类型为 httpGet 或 tcpSocket 时，需要在 httpGet 或 tcpSocket 字段中分别指定要访问的 URL 或主机和端口号。
  您可以根据应用程序的需要，自行定义并配置这些钩子。例如，可以使用 preStop 钩子在容器停止之前保存一些数据，或者使用 preStart 钩子在容器启动之前初始化一些资源。

### 2.1 preStop 停止保存一些数据，您可以指定一些您需要的信息

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: my-container
      image: nginx
      volumeMounts:
        - name: data
          mountPath: /data
      lifecycle:
        preStop:
          exec:
            command: ["/bin/sh", "-c", "echo 'Saving data...'; cp -r /data /data-backup"]
  volumes:
    - name: data
      emptyDir: {}

```

在这个示例中，我们定义了一个名为 my-pod 的 Pod，其中包含一个名为 my-container 的容器，它使用 Nginx 镜像，并将一个名为 data 的空目录挂载到 /data 目录下。在容器的 lifecycle 字段下，我们配置了 preStop 钩子，使用 exec 执行了一个命令，将 /data 目录下的数据复制到 /data-backup 目录中，从而保存了数据。当 Kubernetes 调度器删除这个 Pod 时，该容器的 preStop 钩子就会被执行，从而实现了数据的保存。

### 2.2 preStart 开始启动应用前

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: my-pod
spec:
  containers:
    - name: my-container
      image: nginx
      lifecycle:
        preStart:
          exec:
            command: ["/bin/sh", "-c", "echo 'Initializing resources...'; mkdir -p /data"]
      volumeMounts:
        - name: data
          mountPath: /data
  volumes:
    - name: data
      emptyDir: {}

```

在这个示例中，我们定义了一个名为 my-pod 的 Pod，其中包含一个名为 my-container 的容器，它使用 Nginx 镜像，并将一个名为 data 的空目录挂载到 /data 目录下。在容器的 lifecycle 字段下，我们配置了 preStart 钩子，使用 exec 执行了一个命令，创建了一个名为 /data 的目录，从而初始化了资源。当 Kubernetes 调度器创建这个 Pod 并启动容器时，该容器的 preStart 钩子就会被执行，从而实现了资源的初始化。

### 2.3 配置 coredump 以便更好捕获主 机异常信息

- 1、配置 core 文件转存

```shell
echo "/tmp/dumps/core.%e.%p" > /proc/sys/kernel/core_pattern
```
注意提示，如果是 Kubernetes 环境中，需要给每个节点配置 core_pattern，建议您使用运维编排工具

/proc/sys/kernel/core_pattern 文件定义了当程序崩溃时核心转储文件的命名格式和位置。

配置释义说明:

/tmp/dumps/ 指定了转储文件应该存放的目录。
core 是转储文件的前缀。
%e 是一个模板，代表崩溃的程序的执行文件名。
%p 是一个模板，代表崩溃的程序的进程ID。
所以，如果一个名为 myapp 的程序（进程ID为 1234）崩溃，其核心转储文件的名字将会是 /tmp/dumps/core.myapp.1234

- 永久生效：

修改 /proc/sys/kernel/core_pattern 只会在当前的系统运行期间生效。如果您想要让这个改变在系统重启后仍然生效，您需要做以下的配置：
```shell
kernel.core_pattern = /tmp/dumps/core.%e.%p

```
然后，运行以下命令来应用更改：
```shell
sudo sysctl -p
```

- 使用 systemd：

如果您的系统使用的是 systemd, 你也可以通过 systemd 的 sysctl.d 目录来配置。创建一个文件，例如 /etc/sysctl.d/50-core-pattern.conf，并添加上面同样的行。

无论选择哪种方法，都确保 /tmp/dumps/ 目录存在，并具有适当的权限，以便可以写入核心转储文件。

- 2、配置 yaml 文件
```yaml
apiVersion: v1
kind: Pod
metadata:
  name: nginx-coredump-test
spec:
  containers:
  - name: nginx-container
    image: nginx
    securityContext:
      privileged: true
    volumeMounts:
    - name: coredump-storage
      mountPath: /tmp
  volumes:
  - name: coredump-storage
    hostPath:
      path: /tmp
      type: Directory

```
- 使用shell命令模拟错误：
进入nginx容器：
```shell
kubectl exec -it nginx-coredump-test -- /bin/sh

```
然后，使用以下命令模拟一个段错误：
```shell
sh -c 'kill -s SEGV $$'

```
在对应主机中，检查是否生成对应 core 文件
```shell
[root@issac]# ls -lrt /tmp/dumps/
total 140
-rw------- 1 root root 454656 Nov  1 15:36 core.sh.52
```
- 利用 gbd 查看错误信息
```shell
# gdb /bin/sh /tmp/dumps/core.sh.52
GNU gdb (GDB) Red Hat Enterprise Linux 9.2-7.1.0.4.al8
Copyright (C) 2020 Free Software Foundation, Inc.
License GPLv3+: GNU GPL version 3 or later <http://gnu.org/licenses/gpl.html>
This is free software: you are free to change and redistribute it.
There is NO WARRANTY, to the extent permitted by law.
Type "show copying" and "show warranty" for details.
This GDB was configured as "x86_64-redhat-linux-gnu".
Type "show configuration" for configuration details.
For bug reporting instructions, please see:
<http://www.gnu.org/software/gdb/bugs/>.
Find the GDB manual and other documentation resources online at:
    <http://www.gnu.org/software/gdb/documentation/>.

For help, type "help".
Type "apropos word" to search for commands related to "word"...
Reading symbols from /bin/sh...
Reading symbols from .gnu_debugdata for /usr/bin/bash...
(No debugging symbols found in .gnu_debugdata for /usr/bin/bash)

warning: core file may not match specified executable file.
[New LWP 52]
Core was generated by `sh -c kill -s SEGV $$'.
Program terminated with signal SIGSEGV, Segmentation fault.
#0  0x00007f475de171e7 in ?? ()
(gdb) bt
#0  0x00007f475de171e7 in ?? ()
#1  0x0000558dfaa3feda in ?? ()
#2  0x0000000000000000 in ?? ()
(gdb) 
```
- 核心转储是由 sh -c kill -s SEGV $$ 命令生成的,然而，回溯（bt 命令的输出）只提供了有限的信息，因为 /bin/sh 可执行文件没有包含调试符号，而且这种故意的段错误产生的堆栈可能不是非常有用。
- 可以使用更多的 gdb 命令（如 list、info locals、info registers 等）来获取更多关于崩溃上下文的信息。如果有其他应用程序或具体的故障场景，可能会得到更详细的回溯。