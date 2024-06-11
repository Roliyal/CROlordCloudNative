# 常见问题解答（FAQ）

## 简介

本FAQ文档旨在回答用户在使用过程中常见的问题。通过本文件，您可以快速找到解决方案，提高使用体验。本文件主要面向新用户和使用中遇到问题的用户。

## 目录

1. [提示 `Also: org.jenkinsci.plugins.workflow.actions.ErrorAction$ErrorId: ce3be74c-0e12-49bc-bd31-e555e69d4c28org.jenkinsci.plugins.scriptsecurity.sandbox.RejectedAccessException: Scripts not permitted to use method groovy.lang.GString getBytes` 如何解决？](#1-常见问题)
2. [提示 `Also: org.jenkinsci.plugins.workflow.actions.ErrorAction$ErrorId: 894f0031-9c2d-4bef-ae99-41c56e466d2dorg.jenkinsci.plugins.scriptsecurity.sandbox.RejectedAccessException: Scripts not permitted to use staticMethod org.codehaus.groovy.runtime.EncodingGroovyMethods encodeBase64 byte[]` 如何解决？](#2-常见问题)
3. [提示`error checking push permissions -- make sure you entered the correct tag name, and that you are authenticated correctly, and try again: checking push permission for "***********************": POST https://******/blobs/uploads/: UNAUTHORIZED: authentication required; [map[Action:pull Class: Name:m**o/m***t Type:repository] map[Action:push Class: Name:m**o/m**t Type:repository]]
   06:54:33.624697 common.go:40: exit status 1` 如何解决？](#3-常见问题)

## 问题与解答

### 1. 常见问题

**问题**：提示 `Also: org.jenkinsci.plugins.workflow.actions.ErrorAction$ErrorId: ce3be74c-0e12-49bc-bd31-e555e69d4c28org.jenkinsci.plugins.scriptsecurity.sandbox.RejectedAccessException: Scripts not permitted to use method groovy.lang.GString getBytes` 如何解决？

**解答**：点击Dashboard > Manage Jenkins > ScriptApproval, 点击`Approve`批准`method groovy.lang.GString getBytes`, 重新运行流水线即可。

### 2. 常见问题

**问题**：提示 `Also: org.jenkinsci.plugins.workflow.actions.ErrorAction$ErrorId: 894f0031-9c2d-4bef-ae99-41c56e466d2dorg.jenkinsci.plugins.scriptsecurity.sandbox.RejectedAccessException: Scripts not permitted to use staticMethod org.codehaus.groovy.runtime.EncodingGroovyMethods encodeBase64 byte[]` 如何解决？

**解答**：点击Dashboard > Manage Jenkins > ScriptApproval, 点击`Approve`批准`staticMethod org.codehaus.groovy.runtime.EncodingGroovyMethods encodeBase64 byte[]`, 重新运行流水线即可。

### 3. 常见问题
**问题**: 提示`error checking push permissions -- make sure you entered the correct tag name, and that you are authenticated correctly, and try again: checking push permission for "***********************": POST https://******/blobs/uploads/: UNAUTHORIZED: authentication required; [map[Action:pull Class: Name:m**o/m***t Type:repository] map[Action:push Class: Name:m**o/m**t Type:repository]]
06:54:33.624697 common.go:40: exit status 1` 如何解决？

**解答**: 在阿里云镜像控制台，[容器镜像服务](https://cr.console.aliyun.com/cn-hongkong/instances), 选择需要上传的镜像仓库实例，创建命名空间（需要与本文IMAGE_NAMESPACE名称一致），并开启自动创建仓库。


## 联系信息

如需更多帮助，请通过以下方式联系我们：
- **邮箱**：issac@roliyal.com
- **官方网站**：[www.crolord.com](http://www.crolrod.com)
