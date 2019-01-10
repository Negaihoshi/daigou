package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/negaihoshi/daigou/models"
	"github.com/negaihoshi/daigou/pkg/gredis"
	"github.com/negaihoshi/daigou/pkg/logging"
	"github.com/negaihoshi/daigou/pkg/setting"
	"github.com/negaihoshi/daigou/routers"
)

// @title Golang Gin API
// @version 1.0
// @description An example of gin
// @termsOfService https://github.com/negaihoshi/daigou

// @license.name MIT
// @license.url https://github.com/negaihoshi/daigou/blob/master/LICENSE
func main() {
	setting.Config()
	models.Setup()
	logging.Setup()
	gredis.Setup()

	routersInit := routers.InitRouter()
	readTimeout := setting.AppConfig.Server.ReadTimeout
	writeTimeout := setting.AppConfig.Server.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.AppConfig.Server.HttpPort)
	maxHeaderBytes := 1 << 20

	server := &http.Server{
		Addr:           endPoint,
		Handler:        routersInit,
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			log.Printf("Listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	log.Println("Serv	wer exiting")

	// endless.DefaultReadTimeOut = readTimeout
	// endless.DefaultWriteTimeOut = writeTimeout
	// endless.DefaultMaxHeaderBytes = maxHeaderBytes
	// server := endless.NewServer(endPoint, routersInit)
	// server.BeforeBegin = func(add string) {
	// 	log.Printf("Actual pid is %d", syscall.Getpid())
	// }

	// err := server.ListenAndServe()
	// if err != nil {
	// 	log.Printf("Server err: %v", err)
	// }
}
