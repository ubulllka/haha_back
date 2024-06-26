package app

import (
	"haha/internal/config"
	"haha/internal/db"
	"haha/internal/handlers"
	"haha/internal/logger"
	"haha/internal/server"
	"haha/internal/service"
)

func Run(logg logger.Logger) error {

	conf := config.InitConfig(logg)

	dataBase, err := db.InitializeDB(conf.DB.Host, conf.DB.Port, conf.DB.User, conf.DB.Password, conf.DB.Name, logg)
	if err != nil {
		logg.Error(err)
		return err
	}
	defer dataBase.Close()

	repo := service.NewRepository(dataBase, &logg)
	serv := handlers.NewService(repo, &logg)
	hand := handlers.NewHandler(serv, &logg)

	logg.Info("Init repositories, services, handlers")

	srv := new(server.Server)
	if err := srv.Run(conf.Server.URL, hand.InitRouter(conf.Client.URL), &logg); err != nil {
		logg.Fatalf("server did not start work: %s", err.Error())
		return err
	}

	logg.Info("Server listening url " + conf.Server.URL)

	return nil
}
