package main

import (
	"context"
	"log"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/duyledat197/go-gen-tools/cmd"
	_ "github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
)

var srv server

func load(ctx context.Context) error {

	if err := srv.loadConfig(ctx); err != nil {
		return err
	}

	if err := srv.loadLogger(); err != nil {
		return err
	}

	if err := srv.loadDatabaseClients(ctx); err != nil {
		return err
	}

	if err := srv.loadThirdPartyClients(ctx); err != nil {
		return err
	}

	if err := srv.loadRepositories(); err != nil {
		return err
	}

	if err := srv.loadServices(); err != nil {
		return err
	}

	if err := srv.loadDeliveries(); err != nil {
		return err
	}

	if err := srv.loadPublishers(ctx); err != nil {
		return err
	}

	if err := srv.loadSubscribers(ctx); err != nil {
		return err
	}

	if err := srv.loadClients(ctx); err != nil {
		return err
	}

	if err := srv.loadServers(ctx); err != nil {
		return err
	}

	return nil
}
func start(ctx context.Context) error {
	errChan := make(chan error)

	for _, database := range srv.databases {
		if err := database.Connect(ctx); err != nil {
			return err
		}
	}

	for _, p := range srv.processors {
		go func(p processor) {
			if err := p.Start(ctx); err != nil {
				errChan <- err
			}
		}(p)
	}
	go func() {
		err := <-errChan
		srv.logger.Sugar().Errorf("start error: %w\n", err)
	}()
	return nil
}

func stop(ctx context.Context) error {
	for _, processor := range srv.processors {
		if err := processor.Stop(ctx); err != nil {
			return err
		}
	}

	for _, database := range srv.databases {
		if err := database.Stop(ctx); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	cmd.Execute()
	ctx := context.Background()
	// TODO: with graceful shutdown
	timeWait := 15 * time.Second
	signChan := make(chan os.Signal, 1)

	if err := load(ctx); err != nil {
		log.Fatal(err)
	}

	if err := start(ctx); err != nil {
		log.Fatal(err)
	}
	signal.Notify(signChan, os.Interrupt, syscall.SIGTERM)
	<-signChan
	log.Println("Shutting down")
	ctx, cancel := context.WithTimeout(context.Background(), timeWait)
	defer func() {
		log.Println("Close another connection")
		cancel()
	}()
	if err := stop(ctx); err == context.DeadlineExceeded {
		log.Print("Halted active connections")
	}
	close(signChan)
	log.Printf("Server down Completed")
}
