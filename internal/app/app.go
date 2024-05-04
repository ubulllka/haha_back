package app

import (
	"haha/internal/config"
	"haha/internal/db"
	"haha/internal/handlers"
	"haha/internal/logger"
	"haha/internal/server"
	"haha/internal/service"
)

func Run() error {
	logg := logger.GetLogger()

	conf, err := config.InitConfig()
	if err != nil {
		logg.Error(err)
		return err
	}

	db, err := db.InitializeDB(logg, conf.DB.Host, conf.DB.User, conf.DB.Password, conf.DB.Name, conf.DB.Port)
	if err != nil {
		logg.Error(err)
		return err
	}
	defer db.Close()

	repo := service.NewRepository(db)
	serv := handlers.NewService(repo)
	hand := handlers.NewHandler(serv)

	logg.Info("Init repositories, services, handlers")

	srv := new(server.Server)
	if err := srv.Run(conf.Server.URL, hand.InitRouter()); err != nil {
		logg.Fatalf("server did not start work: %s", err.Error())
		return err
	}

	logg.Info("Server listening url " + conf.Server.URL)

	return nil
}
