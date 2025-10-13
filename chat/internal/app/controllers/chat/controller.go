package chat

import (
	"chat/internal/app/usecases/chat"
	pb "chat/pkg/api/chat/v1"
)

type Deps struct {
	ChatUsecase chat.Usecase
}

type Controller struct {
	pb.UnimplementedChatServiceServer
	Deps
}

func New(d Deps) *Controller {
	return &Controller{
		Deps: d,
	}
}
