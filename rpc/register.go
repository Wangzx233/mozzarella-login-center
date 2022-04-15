package rpc

import (
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/model"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"log"
	"time"
)

const (
	serviceName = "mozzarella-login-center"
	ip          = "8.142.81.74"
	port        = 8085
	version     = "0.0.1"
	grpcPort    = "8085"
)

var (
	KeyServiceHost string
)

func RegisterCenter() {
	// 连接服务器
	//create clientConfig
	clientConfig := constant.ClientConfig{
		NamespaceId:         "", //we can create multiple clients with different namespaceId to support multiple namespace.When namespace is public, fill in the blank string here.
		TimeoutMs:           5000,
		NotLoadCacheAtStart: true,
		LogLevel:            "error",
	}

	serverConfigs := []constant.ServerConfig{
		{
			IpAddr: "175.24.203.115",
			Port:   8848,
		},
	}
	namingClient, err := clients.CreateNamingClient(map[string]interface{}{
		"serverConfigs": serverConfigs,
		"clientConfig":  clientConfig,
	})
	if err != nil {
		log.Println(err)
		return
	}

	//订阅keycenter实时获取服务地址
	err = namingClient.Subscribe(&vo.SubscribeParam{
		ServiceName: "mozzarella-keycenter",
		Clusters:    nil,
		GroupName:   "",
		SubscribeCallback: func(services []model.SubscribeService, err error) {
			if err != nil {
				log.Println("SubscribeCallback err : ", err)
				return
			}
			//获取一个健康的服务地址
			for _, service := range services {
				if service.Healthy && service.Enable {
					KeyServiceHost = service.Ip + ":" + service.Metadata["gRPC_port"]
					NewKeyCenter()
					return
				}
			}
		},
	})

	_, err = namingClient.RegisterInstance(vo.RegisterInstanceParam{
		Ip:          ip,
		Port:        port,
		ServiceName: serviceName,
		Weight:      10,
		Metadata:    map[string]string{"version": version, "up-time": time.Now().String(), "gRPC_port": grpcPort},
		Enable:      true,
		Healthy:     true,
		Ephemeral:   true,
	})
	if err != nil {
		log.Println(err)
	}

}
