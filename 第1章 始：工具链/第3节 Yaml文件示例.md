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