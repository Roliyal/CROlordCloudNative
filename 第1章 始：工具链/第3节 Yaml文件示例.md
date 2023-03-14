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

请注意，此配置示例可能不适用于所有情况，并且可能需要进行适当的修改(存储卷、LB)才能满足您的特定需求。

