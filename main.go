package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fulgurant/datastore"
	"github.com/fulgurant/health"
	"github.com/fulgurant/server"
	"github.com/fulgurant/shitake/usermanager"
	"github.com/fulgurant/simplehash"

	"github.com/alecthomas/kong"
	"go.uber.org/zap"
)

type Config struct {
	Run struct{} `cmd:"" default:"1" help:"Run the web service"`

	Redis            string        `default:"redis" help:"IP Address of Redis Server" short:"r"`
	ListenAddress    string        `default:":8080" help:"Address on which the service will run" short:"p"`
	ShutdownDuration time.Duration `default:"1s" help:"Duration to take to allow existing connections to complete" short:"s"`
	WarningDuration  time.Duration `default:"1s" help:"Duration to take to allow routing algorithms to reroute" short:"w"`
}

func run(cfg *Config) error {
	logger, err := zap.NewProduction()
	if err != nil {
		return err
	}

	h := health.New(logger)

	svrOptions := server.
		DefaultOptions().
		WithListenAddress(cfg.ListenAddress).
		WithHealth(h).
		WithLogger(logger).
		WithShutdownDuration(cfg.ShutdownDuration).
		WithWarningDuration(cfg.WarningDuration)

	svr, err := server.New(svrOptions)
	if err != nil {
		return err
	}
	defer svr.Close()

	h.RegisterEndpoint(svr)

	ds := datastore.NewMock()
	sh := simplehash.NewMock("salt")

	umOptions := usermanager.
		DefaultOptions().
		WithLogger(logger).
		WithGetSetter(ds).
		WithHasher(sh)

	um, err := usermanager.New(umOptions)
	if err != nil {
		return err
	}

	um.RegisterEndpoints(svr)

	waitForShutdownSignal()

	return nil
}

func waitForShutdownSignal() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
}

func main() {
	var cfg Config

	ctx := kong.Parse(&cfg)

	switch ctx.Command() {
	case "run":
		err := run(&cfg)
		if err != nil {
			fmt.Printf("%v", err)
			os.Exit(1)
		}
	}
}
