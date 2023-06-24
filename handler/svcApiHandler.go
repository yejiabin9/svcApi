package handler

import (
	"context"
	"encoding/json"
	log "github.com/asim/go-micro/v3/logger"
	svc "github.com/yejiabin9/svc/proto/svc"
	svcApi "github.com/yejiabin9/svcApi/proto/svcApi"
)

type SvcApi struct {
	SvcService svc.SvcService
}

// svcApi.FindSvcById 通过API向外暴露为/svcApi/findSvcById，接收http请求
// 即：/svcApi/FindSvcById 请求会调用go.micro.api.svcApi 服务的svcApi.FindSvcById 方法
func (e *SvcApi) FindSvcById(ctx context.Context, req *svcApi.Request, rsp *svcApi.Response) error {
	log.Info("Received svcApi.FindSvcById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/svcApi/FindSvcById'}")
	rsp.Body = string(b)
	return nil
}

// svcApi.AddSvc 通过API向外暴露为/svcApi/AddSvc，接收http请求
// 即：/svcApi/AddSvc 请求会调用go.micro.api.svcApi 服务的svcApi.AddSvc 方法
func (e *SvcApi) AddSvc(ctx context.Context, req *svcApi.Request, rsp *svcApi.Response) error {
	log.Info("Received svcApi.AddSvc request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/svcApi/AddSvc'}")
	rsp.Body = string(b)
	return nil
}

// svcApi.DeleteSvcById 通过API向外暴露为/svcApi/DeleteSvcById，接收http请求
// 即：/svcApi/DeleteSvcById 请求会调用go.micro.api.svcApi 服务的 svcApi.DeleteSvcById 方法
func (e *SvcApi) DeleteSvcById(ctx context.Context, req *svcApi.Request, rsp *svcApi.Response) error {
	log.Info("Received svcApi.DeleteSvcById request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/svcApi/DeleteSvcById'}")
	rsp.Body = string(b)
	return nil
}

// svcApi.UpdateSvc 通过API向外暴露为/svcApi/UpdateSvc，接收http请求
// 即：/svcApi/UpdateSvc 请求会调用go.micro.api.svcApi 服务的svcApi.UpdateSvc 方法
func (e *SvcApi) UpdateSvc(ctx context.Context, req *svcApi.Request, rsp *svcApi.Response) error {
	log.Info("Received svcApi.UpdateSvc request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问/svcApi/UpdateSvc'}")
	rsp.Body = string(b)
	return nil
}

// 默认的方法svcApi.Call 通过API向外暴露为/svcApi/call，接收http请求
// 即：/svcApi/call或/svcApi/ 请求会调用go.micro.api.svcApi 服务的svcApi.FindSvcById 方法
func (e *SvcApi) Call(ctx context.Context, req *svcApi.Request, rsp *svcApi.Response) error {
	log.Info("Received svcApi.Call request")
	rsp.StatusCode = 200
	b, _ := json.Marshal("{success:'成功访问：Call'}")
	rsp.Body = string(b)
	return nil
}
