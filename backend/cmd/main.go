package main

import (
	"backend/config"
	"backend/internal/app"
	"backend/internal/repo/postgres"
	"backend/pkg/logger"
	"backend/pkg/token/paseto"
	"context"
	"fmt"
	"log"
	"os"
)

func setLogger(loggerFile *os.File, cfg *config.Logger) (*logger.Logger, error) {
	fmt.Println("trying to create logger")
	loggerFile, err := os.OpenFile(
		cfg.File,
		os.O_APPEND|os.O_CREATE|os.O_WRONLY,
		0664,
	)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func(loggerFile *os.File) {
		err := loggerFile.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(loggerFile)

	l := logger.New(cfg.Level, loggerFile)
	return l, nil
}

func main() {
	ctx := context.Background()

	fmt.Println("trying to read config")
	c, err := config.New()
	if err != nil {
		log.Fatal(err)
	}

	var loggerFile *os.File
	l, err := setLogger(loggerFile, &c.Logger)
	if err != nil {
		return
	}

	dbConnector, err := postgres.NewDbConn(ctx, &c.DB.Postgres)
	if err != nil {
		log.Fatal("failed to postgres connect", err)
		return
	}

	pasetoStruct, err := paseto.NewPaseto(paseto.KEY)
	if err != nil {
		log.Fatal("failed to generate paseto struct", err)
		return
	}

	fmt.Println("trying to make new app")
	_ = app.NewApp(l, dbConnector, pasetoStruct, c)

}
