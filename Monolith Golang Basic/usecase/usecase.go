package usecase

import (
	"golang-basic/config"
	_interface "golang-basic/repository/interface"
	"golang-basic/repository/mysql"
)

type usecase struct {
	main     *config.MainConfig
	orderRepo _interface.OrderRepository
}

func InitUsecase(conf *config.Config) (usecase, error) {
	db, err := config.ConnectDB(conf.DB)
	if err != nil {
		return usecase{}, err
	}

	return usecase{
		main:     conf.Main,
		orderRepo: mysql.NewOrderRepo(db),
	}, nil
}
