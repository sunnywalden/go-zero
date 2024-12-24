// Code generated by goctl. DO NOT EDIT!
// Source: hi.proto

package server

import (
	"context"

	eventlogic "github.com/sunnywalden/go-zero/tools/goctl/example/rpc/hi/internal/logic/event"
	"github.com/sunnywalden/go-zero/tools/goctl/example/rpc/hi/internal/svc"
	"github.com/sunnywalden/go-zero/tools/goctl/example/rpc/hi/pb/hi"
)

type EventServer struct {
	svcCtx *svc.ServiceContext
	hi.UnimplementedEventServer
}

func NewEventServer(svcCtx *svc.ServiceContext) *EventServer {
	return &EventServer{
		svcCtx: svcCtx,
	}
}

func (s *EventServer) AskQuestion(ctx context.Context, in *hi.EventReq) (*hi.EventResp, error) {
	l := eventlogic.NewAskQuestionLogic(ctx, s.svcCtx)
	return l.AskQuestion(in)
}
