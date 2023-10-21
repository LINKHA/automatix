package service_discovery

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/linkha/automatix/common"
	"go.etcd.io/etcd/clientv3"
)

type Method struct {
	service string
	method  string
	addr    string
}

func (m *Method) GenMethod() {
	fmt.Sprintf("/automatix/services/%s/%s/%s", m.service, m.method, m.addr)
}

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
	}, nil
}

func (sd *ServiceDiscovery) Serve() error {
	if sd.leaseID != 0 {
		return common.Asddff
		// return &common.Error{ErrCode: 1001, ErrMsg: "Dss"}
	}

	leaseResp, err := sd.client.Grant(sd.ctx, int64(sd.leaseTTL))
	if err != nil {
		return err
	}

	sd.leaseID = leaseResp.ID

	// 启动一个定时器来定期续约租约
	go sd.keepAlive()

	// 定期检查服务变化
	go sd.watchService()

	return nil
}

// Wait for next heartbeat Register method(等待下次心跳注册方法)
func (sd *ServiceDiscovery) RegisterMethod(registerMethod map[string]struct{}) error {
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

func (sd *ServiceDiscovery) DiscoverService() ([]string, error) {
	//watch
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := sd.client.Get(ctx, fmt.Sprintf("/services/%s", sd.service), clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	endpoints := make([]string, 0)
	for _, kv := range resp.Kvs {
		endpoints = append(endpoints, string(kv.Value))
	}

	return endpoints, nil
}

func (sd *ServiceDiscovery) Close() (err error) {
	return sd.client.Close()
}

func (sd *ServiceDiscovery) keepAlive() {
	heartbeat := time.NewTicker(time.Duration(sd.leaseTTL/2) * time.Second)
	defer heartbeat.Stop()

	for {
		select {
		case <-heartbeat.C:
			// Check whether there is data in the queue to be registered
			// (判断待注册队列中是否有数据)
			sd.putMethod()

			// 续约租约，告诉 etcd 服务还在正常运行
			_, err := sd.client.KeepAlive(context.Background(), sd.leaseID)
			if err != nil {
				//失败三次就将数据放回待注册的服务
				log.Printf("Failed to renew lease: %v", err)
			}
		}
	}
}

// Register methods to Etcd(注册方法到Etcd)
func (sd *ServiceDiscovery) putMethod() error {
	//Simply determine whether there is data, and skip the case without data
	//(简单判断一下是否有数据，没有数据的情况直接跳过)
	if len(sd.toRegisteredMethod) == 0 {
		return nil
	}

	sd.lock.Lock()
	defer sd.lock.Unlock()

	//Prevents writes from affecting the sending of heartbeats
	//(防止写入影响到心跳的发送)
	go func() {
		for method, _ := range sd.toRegisteredMethod {
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
			delete(sd.toRegisteredMethod, method)
		}
	}()

	return nil
}

func (sd *ServiceDiscovery) watchService() {
	serviceChangesCh := make(chan []string)

	go func() {
		key := fmt.Sprintf("/services/%s", sd.service)
		rch := sd.client.Watch(sd.ctx, key, clientv3.WithPrefix())

		for wresp := range rch {
			for _, ev := range wresp.Events {
				switch ev.Type {
				case clientv3.EventTypePut:
					// 服务新增或更新
					endpoints := []string{string(ev.Kv.Value)}
					serviceChangesCh <- endpoints
				case clientv3.EventTypeDelete:
					// 服务删除
					serviceChangesCh <- nil
				}
			}
		}
	}()

	go func() {
		for {
			endpoints := <-serviceChangesCh
			if endpoints != nil {
				fmt.Printf("Service %s updated. New endpoints: %v\n", sd.service, endpoints)
			} else {
				fmt.Printf("Service %s deleted.\n", sd.service)
			}
		}
	}()
}
