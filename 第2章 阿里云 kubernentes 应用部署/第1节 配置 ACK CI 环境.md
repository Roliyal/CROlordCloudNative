# ACK CI配置

### CI起源

CI的概念最初由Martin Fowler在2000年的一篇博客文章中提出。他指出，CI是一种实践，旨在通过频繁地将代码集成到共享代码库中来降低开发成本和改善软件质量。这个想法很快得到了广泛的认可，并成为敏捷开发和DevOps中的核心实践之一。

在2000年代初，CI工具开始出现，帮助开发人员自动化构建、测试和部署代码。这些工具包括CruiseControl和Hudson等开源工具，以及商业工具如Jenkins和Travis CI等。（详见第一章第一节“配置企业CICD工具链”）

随着时间的推移，CI已经成为现代软件开发的重要实践之一。通过CI，团队可以更快地交付高质量的代码，并改进整个软件开发过程。随着持续交付、持续部署和DevOps的出现，CI变得更加重要和复杂，成为现代软件开发的不可或缺的一部分。

配置CI环境可以提高团队的效率和代码质量，与CD（持续交付）工具集成，以支持自动部署，加快交付速度，并提高生产环境的可靠性。

## ACK CI配置

### 前置条件

在开始配置ACK CI环境之前，您需要具备以下条件：

* 阿里云ACK集群：并具有适当的ACK集群权限，如果您没有创建过集群，请参考阿里云的文档创建一个ACK集群。（参考链接：<https://help.aliyun.com/document_detail/95108.html>）

* CI工具：需要选择一个CI工具来管理您的代码和自动化构建、测试和部署。本文我们以jenkins为例（构建方法详见第一章第一节“配置企业CICD工具链”）

* Docker镜像仓库：需要有一个Docker镜像仓库来存储您的应用程序镜像。阿里云提供了一个容器镜像服务ACR（Container Registry），本文中使用该服务作为镜像仓库服务。（参考链接：<https://help.aliyun.com/document_detail/300068.html>）

* Kubernetes插件：需要在您的CI工具中安装一个Kubernetes插件，以便可以使用CI工具来管理Kubernetes集群。例如，在Jenkins中，可以使用Kubernetes插件来管理Kubernetes集群。

* Git代码仓库：需要一个Git代码仓库来存储您的应用程序代码。可以使用GitHub、GitLab等Git代码托管服务，本文中使用“CROlordcodelibrary”作为代码仓库。（参考链接： <https://github.com/Roliyal/CROlordcodelibrary>）

这些是配置阿里云ACK的CI环境所需的主要前置条件。在准备好这些条件之后，就可以开始配置CI环境并使用它来自动化软件开发过程。

#### 部署Jenkins pipeline 流水线

```pipline
pipeline {
    // 定义使用的 Jenkins agent 类型
    agent { kubernetes { /* 配置省略 */ } }
    
    // 定义环境变量
    environment {
        GIT_BRANCH = 'main' // Git主分支的默认值
        MAJOR_VERSION = 'v1' // 主版本号
        MINOR_VERSION = '0'  // 次版本号
        PLATFORMS = 'linux/amd64' // 构建目标平台
        MAJOR = "${params.MAJOR_VERSION ?: env.MAJOR_VERSION ?: '1'}" // 主版本号，允许通过参数覆盖
        MINOR = "${params.MINOR_VERSION ?: env.MINOR_VERSION ?: '0'}" // 次版本号，允许通过参数覆盖
        PATCH = "${env.BUILD_NUMBER}" // 构建号，用作修订版本号
        VERSION_TAG = "${MAJOR}.${MINOR}.${PATCH}" // 组合版本标签
        IMAGE_REGISTRY = "${params.IMAGE_REGISTRY}" // 镜像仓库地址
        IMAGE_NAMESPACE = "${params.IMAGE_NAMESPACE}" // 镜像命名空间
        IMAGE_ID = "${params.IMAGE_NAMESPACE}" // 镜像ID
    }

    // 触发条件
    triggers { githubPush() }

    // 参数定义
    parameters {
        string(name: 'BRANCH', defaultValue: 'main', description: 'Git branch to build')
        choice(name: 'PLATFORMS', choices: ['linux/amd64', 'linux/amd64,linux/arm64,'], description: 'Target platforms for ACR registry')
        string(name: 'GIT_REPOSITORY', defaultValue: 'https://github.com/Roliyal/CROlordCodelibrary.git', description: 'Git repository URL')
        string(name: 'MAJOR_VERSION', defaultValue: '1', description: 'Major version number')
        string(name: 'MINOR_VERSION', defaultValue: '0', description: 'Minor version number')
        string(name: 'BUILD_DIRECTORY', defaultValue: 'Chapter2KubernetesApplicationBuild/Unit2CodeLibrary/FEBEseparation/go-guess-number', description: 'Build directory path')
        string(name: 'IMAGE_REGISTRY', defaultValue: 'crolord-registry-registry-vpc.cn-hongkong.cr.aliyuncs.com', description: 'The Alibaba ACR registry to use')
        string(name: 'IMAGE_NAMESPACE', defaultValue: 'febe', description: 'The Alibaba ACR image namespace')
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
                            --platforms linux/amd64,linux/arm64 \\
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


































