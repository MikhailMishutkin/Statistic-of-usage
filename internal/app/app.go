package app

import (
	"testtask/internal/controller"
	"testtask/internal/storage"
	"testtask/internal/usecase"
)

func Run() {

	repository := storage.NewStorage(storage.Repository{})
	useCase := usecase.New(repository)
	c := controller.New(useCase)

	useCase.Separate(c.UserAnsw())
	useCase.ToCompare()
	useCase.ResToSrcn()
}
