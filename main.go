package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/nsg3355/ph-cafe-manager/config"
	route "github.com/nsg3355/ph-cafe-manager/services"

	"github.com/spf13/cast"
)

func main() {
	// Step 1: Config Load
	config.GetInstance().LoadConfig()

	// Step 2: DB 초기화
	config.InitDB(config.GetInstance().PayhereDB)

	// Database 초기화

	// Step 3: API 서버 초기화
	r := route.InitRouter()

	srv := &http.Server{
		Addr:    ":" + cast.ToString(config.GetInstance().ServicePort),
		Handler: r,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Printf("Starting server... port : %v", config.GetInstance().ServicePort)

	// Step 4: Graceful Shutdown
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	// go-gin 서비스 종료 (타임아웃 10초)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}
}
