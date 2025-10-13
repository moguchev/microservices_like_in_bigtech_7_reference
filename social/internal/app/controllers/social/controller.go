package social

import (
	"social/internal/app/usecases/social"
	pb "social/pkg/api/social/v1"
)

type Deps struct {
	SocialUsecase social.Usecase
}

type Controller struct {
	pb.UnimplementedSocialServiceServer
	Deps
}

func New(d Deps) *Controller {
	return &Controller{
		Deps: d,
	}
}
