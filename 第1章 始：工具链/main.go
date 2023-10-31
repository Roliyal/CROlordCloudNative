package main

import (
	"fmt"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

func main() {
	// 初始化客户端
	client, err := ecs.NewClientWithAccessKey("cn-hangzhou", "LTAI5tC1K6AQXjBByGMbGpYi", "HjYzTyG3XfQ71uPHXqzxuQIDDhw8k0")
	if err != nil {
		fmt.Println("Init error: ", err)
		return
	}

	// 创建API请求
	request := ecs.CreateDescribeSecurityGroupsRequest()
	request.Scheme = "https"

	// 发起请求
	response, err := client.DescribeSecurityGroups(request)
	if err != nil {
		fmt.Println("Error while requesting API: ", err)
		return
	}

	// 处理响应
	for _, group := range response.SecurityGroups.SecurityGroup {
		fmt.Println("Security Group ID: ", group.SecurityGroupId)
		// 在这里，你可以进一步检查每个安全组的规则
	}
}
