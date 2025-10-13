package users

import (
	"users/internal/app/usecases/users"
	pb "users/pkg/api/users/v1"
)

type Deps struct {
	UsersUsecase users.Usecase
}

type Controller struct {
	pb.UnimplementedUserServiceServer
	Deps
}

func New(d Deps) *Controller {
	return &Controller{
		Deps: d,
	}
}
