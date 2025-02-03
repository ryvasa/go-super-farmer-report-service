package main

import (
	"os"
	"os/signal"
	"syscall"

	_ "github.com/lib/pq"
	"github.com/ryvasa/go-super-farmer-report-service/pkg/logrus"
	wire_excel "github.com/ryvasa/go-super-farmer-report-service/pkg/wire"
)

func main() {
	logrus.Log.Info("Starting Report service...")

	// Inisialisasi aplikasi menggunakan Wire
	app, err := wire_excel.InitializeReportApp()
	if err != nil {
		logrus.Log.Fatalf("Failed to initialize app: %v", err)
	}

	// Jalankan gRPC Server dalam goroutine
	go app.StartGRPCServer()

	// Membuat channel untuk menangkap sinyal OS
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Menunggu sinyal untuk menghentikan program
	<-stop

	logrus.Log.Info("Shutting down gracefully...")
	// Cleanup code here if needed
}
