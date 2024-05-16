#!/bin/bash
# 设置命名空间
NAMESPACE=default
# 获取所有部署的名称
DEPLOYMENTS=$(kubectl get deployments -n $NAMESPACE -o=jsonpath='{.items[*].metadata.name}')
# 遍历所有部署，并更新标签
for DEPLOYMENT in $DEPLOYMENTS; do
  # 构建 patch 的 JSON 字符串，确保使用了正确的转义符号
  PATCH_STRING="{\"metadata\": {\"labels\": {\"armsPilotAutoEnable\": \"on\", \"armsSecAutoEnable\": \"on\", \"armsPilotCreateAppName\": \"$DEPLOYMENT\"}}}"
  # 使用 kubectl patch 命令更新标签
  kubectl patch deployment "$DEPLOYMENT" -n $NAMESPACE --type merge -p "$PATCH_STRING"
done
echo "Updated all deployments with the specified labels in namespace $NAMESPACE…"
