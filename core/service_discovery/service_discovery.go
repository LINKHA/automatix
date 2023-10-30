package service_discovery

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type Method struct {
	service string
	method  string
	addr    string
}

func (m *Method) GenMethod() {
	fmt.Sprintf("/automatix/services/%s/%s/%s", m.service, m.method, m.addr)
}

var (
	toRegisteredChSize = 100
)

type ServiceDiscovery struct {
	client   *clientv3.Client
	leaseID  clientv3.LeaseID
	leaseTTL int
	service  string
	addr     string
	ctx      context.Context

	//Registered services(已注册的服务)
	registeredMethod map[string]struct{}
	//Services to be registered(待注册的服务)
	toRegisteredMethod map[string]struct{}
	//Services in a cluster(集群中的服务)
	globalMethod map[string]struct{}

	lock sync.Mutex
}

func NewServiceDiscovery(endpoints []string, service, addr string, leaseTTL int) (*ServiceDiscovery, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

	return &ServiceDiscovery{
		client:             client,
		leaseID:            0,
		service:            service,
		addr:               addr,
		leaseTTL:           leaseTTL,
		ctx:                context.Background(),
		registeredMethod:   make(map[string]struct{}),
		toRegisteredMethod: make(map[string]struct{}),
		globalMethod:       make(map[string]struct{}),
	}, nil
}

func (sd *ServiceDiscovery) Serve() error {
	if sd.leaseID != 0 {
		return nil
	}

	leaseResp, err := sd.client.Grant(sd.ctx, int64(sd.leaseTTL))
	if err != nil {
		return err
	}

	sd.leaseID = leaseResp.ID
	// 启动一个定时器来定期续约租约
	go sd.keepAlive()
	// Check whether there is data in the queue to be registered
	// (判断待注册队列中是否有数据)
	go sd.putMethod()
	// 定期检查服务变化
	sd.watchService()

	return nil
}

// Wait for next heartbeat Register method(等待下次心跳注册方法)
func (sd *ServiceDiscovery) RegisterMethod(registerMethod map[string]struct{}) error {
	sd.lock.Lock()
	defer sd.lock.Unlock()
	for method, _ := range registerMethod {
		sd.toRegisteredMethod[method] = struct{}{}
	}
	return nil
}

// Immediate registration method(立即注册方法)
func (sd *ServiceDiscovery) RegisterMethodImmediate(registerMethod map[string]struct{}) error {
	sd.lock.Lock()
	defer sd.lock.Unlock()

	for method, _ := range registerMethod {
		//Check whether it has been registered(判断是否已经注册过)
		_, ok := sd.registeredMethod[method]
		if ok == false {
			continue
		}

		key := fmt.Sprintf("/automatix/services/%s/%s/%s", sd.service, sd.addr, method)
		_, err := sd.client.Put(sd.ctx, key, sd.addr, clientv3.WithLease(sd.leaseID))
		if err == nil {
			continue
		}

		sd.registeredMethod[method] = struct{}{}
	}

	return nil
}

func (sd *ServiceDiscovery) Close() (err error) {
	return sd.client.Close()
}

func (sd *ServiceDiscovery) keepAlive() {
	interval := time.Duration(sd.leaseTTL / 2)
	if interval <= 0 {
		interval = 1
	}
	heartbeat := time.NewTicker(interval * time.Second)
	defer heartbeat.Stop()

	for {
		select {
		case <-heartbeat.C:
			//续约租约，告诉 etcd 服务还在正常运行
			_, err := sd.client.KeepAlive(context.Background(), sd.leaseID)
			if err != nil {
				//失败三次就将数据放回待注册的服务
				log.Printf("Failed to renew lease: %v", err)
			}
		case <-sd.ctx.Done():
			return
		}
	}
}

// Register methods to Etcd(注册方法到Etcd)
func (sd *ServiceDiscovery) putMethod() error {
	interval := time.Duration(sd.leaseTTL / 2)
	if interval <= 0 {
		interval = 1
	}
	heartbeat := time.NewTicker(interval * time.Second)
	defer heartbeat.Stop()

loop:
	for {
		select {
		case <-heartbeat.C:
			//Simply determine whether there is data, and skip the case without data
			//(简单判断一下是否有数据，没有数据的情况直接跳过)
			if len(sd.toRegisteredMethod) == 0 {
				goto loop
			}

			sd.lock.Lock()
			//Prevents writes from affecting the sending of heartbeats
			//(防止写入影响到心跳的发送)
			for method, _ := range sd.toRegisteredMethod {
				//Check whether it has been registered(判断是否已经注册过)
				_, ok := sd.registeredMethod[method]
				if ok {
					continue
				}

				key := fmt.Sprintf("/automatix/services/%s/%s/%s", sd.service, sd.addr, method)
				_, err := sd.client.Put(sd.ctx, key, sd.addr, clientv3.WithLease(sd.leaseID))
				if err != nil {
					continue
				}

				sd.registeredMethod[method] = struct{}{}
				sd.globalMethod[method] = struct{}{}
				delete(sd.toRegisteredMethod, method)
			}

			sd.lock.Unlock()
		}
	}
}

func (sd *ServiceDiscovery) watchService() {
	serviceChangesCh := make(chan []string)
	serviceDeleteCh := make(chan []string)

	go func() {
		key := "/automatix/services"
		rch := sd.client.Watch(sd.ctx, key, clientv3.WithPrefix())
		for wresp := range rch {
			for _, ev := range wresp.Events {
				methods := []string{string(ev.Kv.Key)}
				switch ev.Type {
				case clientv3.EventTypePut:
					serviceChangesCh <- methods
				case clientv3.EventTypeDelete:
					serviceDeleteCh <- methods
				}
			}
		}
	}()

	go func() {
		for {
			select {
			case methods := <-serviceChangesCh:
				sd.lock.Lock()
				for _, method := range methods {
					sd.globalMethod[method] = struct{}{}
				}
				fmt.Printf("Service %s updated. New method: %v\n", sd.service, methods)
				sd.lock.Unlock()
			case methods := <-serviceDeleteCh:
				sd.lock.Lock()
				for _, method := range methods {
					delete(sd.globalMethod, method)
				}
				fmt.Printf("Service %s deleted. Deleted: %v\n", sd.service, methods)
				sd.lock.Unlock()
			case <-sd.ctx.Done():
				return
			}
		}
	}()
}
