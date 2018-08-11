package main

import (
	"net/http"
	"os"
	"os/signal"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/tony24681379/402/config"
)

func main() {
	config := config.Config()
	g := gin.New()
	g.Use(gin.Recovery())

	glog.Info("serve port", config.Port)
	server := &http.Server{
		Addr:    config.Port,
		Handler: g,
	}
	gracefulShutdown(server)
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			glog.Info("Server closed under request")
		} else {
			glog.Fatal("Server closed unexpect")
		}
	}

	glog.Info("Server exiting")
}

func gracefulShutdown(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		glog.Info("receive interrupt signal")
		if err := server.Close(); err != nil {
			glog.Fatal("Server Close:", err)
		}
	}()
}
