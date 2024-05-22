###### 1. 常见问题提示 `Also:   org.jenkinsci.plugins.workflow.actions.ErrorAction$ErrorId: ce3be74c-0e12-49bc-bd31-e555e69d4c28
org.jenkinsci.plugins.scriptsecurity.sandbox.RejectedAccessException: Scripts not permitted to use method groovy.lang.GString getBytes` 如何解决？
- 解决方案：点击Dashboard > Manage Jenkins > ScriptApproval, 点击`Approve`批准`method groovy.lang.GString getBytes`, 重新运行流水线即可。
###### 2. 常见问题提示 `Also:   org.jenkinsci.plugins.workflow.actions.ErrorAction$ErrorId: 894f0031-9c2d-4bef-ae99-41c56e466d2d
org.jenkinsci.plugins.scriptsecurity.sandbox.RejectedAccessException: Scripts not permitted to use staticMethod org.codehaus.groovy.runtime.EncodingGroovyMethods encodeBase64 byte[]` 如何解决？
- 解决方案：点击Dashboard > Manage Jenkins > ScriptApproval, 点击`Approve`批准`staticMethod org.codehaus.groovy.runtime.EncodingGroovyMethods encodeBase64 byte[]`, 重新运行流水线即可。
