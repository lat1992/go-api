/*
 * File: X:\go-api\main.go
 * Created At: 2019-11-03 18:19:28
 * Created By: Mauhoi WU
 * 
 * Modified At: 2019-11-19 17:19:13
 * Modified By: Mauhoi WU
 */

package main

import (
	"log"
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
	"net/http"
	"crypto/tls"
	"golang.org/x/crypto/acme/autocert"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"./configuration"
)

func serve(TLS string, server *http.Server) {
	if TLS == "no" {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	} else {
		if err := server.ListenAndServeTLS("", ""); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}	
	}
}

func safeQuit(server *http.Server) {
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 3 * time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	select {
		case <-ctx.Done():
			log.Println("Timeout of 3 seconds.")
	}
	log.Println("Server exiting")
}

func main() {
	var server *http.Server
	config := configuration.GetServer()
	router := gin.Default()
	if config["tls"] == "no" {
		router.Use(cors.Default())
		GetRouter(router)
		server = &http.Server{
			Addr: ":"+ config["port"],
			Handler: router,
		}
	} else {
		corsConfig := cors.DefaultConfig()
		corsConfig.AllowOrigins = []string{config["web-url"]}
		router.Use(cors.New(corsConfig))
		GetRouter(router)
		certManager := autocert.Manager{
			Prompt: autocert.AcceptTOS,
			HostPolicy: autocert.HostWhitelist(config["domain"]),
			Cache: autocert.DirCache(config["certificat-directory"]),
		}
		server = &http.Server{
			Addr: ":"+ config["port"],
			Handler: router,
			TLSConfig: &tls.Config{
				GetCertificate: certManager.GetCertificate,
			},
		}
	}
	go serve(config["tls"], server)
	safeQuit(server)
}
