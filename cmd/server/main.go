package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/priscilf/shortener_url/internal/config"
	"github.com/priscilf/shortener_url/server"
)

type MainHandler struct{}

func (m *MainHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func main() {
	ctx := context.Background()
	confifile, err := os.Getwd()
	if err != nil {
		fmt.Println("Err get directory")
		return
	}

	conf, err := config.MustLoadConfig(confifile + "/config/server.yaml")

	if err != nil {
		fmt.Println("Err get config: %s", err.Error())
		return
	}

	handler := MainHandler{}

	address := fmt.Sprintf("%s:%s", conf.HTTPServerConfig.Address, conf.HTTPServerConfig.Port)
	println(address)
	server := server.NewServer(address, conf.HTTPServerConfig, &handler)
	go func() {
		if err := server.Start(); err != nil {
			fmt.Println(err.Error())
			return
		}
	}()

	log.Printf("Server started on %s\n", address)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	defer server.Stop(ctx)

	fmt.Println(conf)
}
