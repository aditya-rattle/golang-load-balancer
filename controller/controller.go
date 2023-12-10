package controller

import (
	"net/http"

	"loadbalancer.com/service"
)

type Controller struct {
	service service.ILoadBalancer
}

func NewController(service service.ILoadBalancer) *Controller {
	return &Controller{
		service: service,
	}
}

func (ctr *Controller) RouteRequest(wr http.ResponseWriter, req *http.Request) {
	ctr.service.Serve(wr, req)
}
