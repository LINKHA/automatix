package service_discovery

import (
	"log"
	"testing"
	"time"
)

var (
	serviceDiscovery *ServiceDiscovery
)

func ClientTest() {
	etcdEndpoints := []string{"http://localhost:2379"} // 替换为您的 etcd 服务器地址
	serviceName := "example-service"
	serviceAddress := "127.0.0.1:8080" // 替换为您的服务地址
	leaseTTL := 2                      // 租约的过期时间，以秒为单位

	var err error
	serviceDiscovery, err = NewServiceDiscovery(etcdEndpoints, serviceName, serviceAddress, leaseTTL)
	if err != nil {
		log.Fatalf("Failed to create service discovery: %v", err)
	}

	if err := serviceDiscovery.Serve(); err != nil {
		log.Fatalf("Failed serve service: %v", err)
	}
}

func TestServiceDiscovery(t *testing.T) {
	ClientTest()
	registerMethod := make(map[string]struct{})
	registerMethod["a1"] = struct{}{}
	registerMethod["a2"] = struct{}{}
	registerMethod["a3"] = struct{}{}

	if err := serviceDiscovery.RegisterMethod(registerMethod); err != nil {
		log.Fatalf("Failed to register service: %v", err)
	}

	time.Sleep(50 * time.Second)
	// select {}
}
