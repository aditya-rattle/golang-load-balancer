package service

import (
	"log"
	"net/http"

	"loadbalancer.com/models"
)

type ILoadBalancer interface {
	GetNextAvailableServer() *models.Server
	Serve(wr http.ResponseWriter, req *http.Request)
}

type LoadBalancer struct {
	servers         []*models.Server
	roundRobinCount int
}

func NewLoadBalancer(servers []*models.Server) ILoadBalancer {
	return &LoadBalancer{
		roundRobinCount: 0,
		servers:         servers,
	}
}

func (lb *LoadBalancer) GetNextAvailableServer() *models.Server {
	targerServer := lb.servers[lb.roundRobinCount%len(lb.servers)]
	for !targerServer.IsAlive() {
		lb.roundRobinCount++
		targerServer = lb.servers[lb.roundRobinCount%len(lb.servers)]
	}
	lb.roundRobinCount++
	return targerServer
}

func (lb *LoadBalancer) Serve(wr http.ResponseWriter, req *http.Request) {
	// wr.Write([]byte("Hello World"))
	target := lb.GetNextAvailableServer()
	log.Printf("Serving request by %s", target.GetAddress())
	target.Proxy.ServeHTTP(wr, req)
}
