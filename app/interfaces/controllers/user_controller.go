package controllers

import (
	"go_http_api_template/domain"
	"go_http_api_template/interfaces/database"
	"go_http_api_template/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create() {
	u := domain.User{}
	controller.Interactor.Add(u)

	return
}

func (controller *UserController) GetUser() domain.Users {
	res := controller.Interactor.GetInfo()
	return res
}

func (controller *UserController) Delete(id string) {
	controller.Interactor.Delete(id)
}
