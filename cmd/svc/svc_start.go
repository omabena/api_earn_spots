package svc

import (
	"context"
	"os"
	"os/signal"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/omabena/api_earn_spots/internal/config"
)

// Config bundles all local configs from environment vars into a structure
type Config struct {
	BaseURL  string `default:"http://localhost:7299" split_words:"true"`
	LogLevel string `default:"info" split_words:"true"`
}

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "start challenge",
	Run:   runStart,
}

func runStart(_ *cobra.Command, _ []string) {
	ctx := cancelationSignal()

	// Acquire config
	var cfg Config
	envPath := config.GetEnvOrDefault("ENV_FILE_PATH", ".env")
	config.MustParseConfig(envPath, &cfg)

	level, err := log.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.WithError(err).Fatal("failed to parse logger level")
	}

	log.SetLevel(level)
	// cat := categorize.NewCategorize(cfg.BaseURL)
	// cat.Execute(ctx)

	<-ctx.Done()
}

// cancelationSignal constructs a cancellation context that gets triggered on os.Interrupt
// Other goroutines can wait for ctx.Done() signal to begin graceful shutdown
func cancelationSignal() context.Context {
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		<-c
		cancel()
	}()

	return ctx
}
