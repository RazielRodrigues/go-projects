package main

import (
	"context"
	"flag"
	"kv-store/internal/container"
	"kv-store/internal/infrastructure/persistence"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	port      = flag.String("port", "6379", "Port to listen")
	enableAOF = flag.Bool("aof", false, "enable aof")
)

func main() {
	flag.Parse()

	opt := persistence.AOFProviderOption{
		EnableAOF: *enableAOF,
		FilePath:  "data.aof",
	}

	ctn, err := container.InitializeContainer(opt)
	if err != nil {
		log.Fatalf("error initializing the container: %v", err)
	}
	defer ctn.Close()

	if *enableAOF && ctn.Persistence != nil {
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		if err := ctn.Persistence.Replay(ctx, ctn.Store); err != nil {
			log.Printf("warning: failed to replay aof: %v", err)
		} else {
			log.Printf("replayed AOF")
		}
	}

	ctn.Store.StartCleanup(1000)

	listener, err := net.Listen("tcp", ":"+*port)
	if err != nil {
		log.Fatalf("error listening on port %v: %v", *port, err)
	}
	defer listener.Close()

	log.Printf("Server listening on port %v (AOF: %v)", *port, *enableAOF)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Printf("caught interrupt, shutting down...")
		listener.Close()
	}()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("error accepting connection: %v", err)
			break
		}
		go ctn.TCPHandler.HandleConnection(conn)
	}
}
