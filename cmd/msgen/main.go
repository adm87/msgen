package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path/filepath"
	"regexp"
	"syscall"

	"github.com/adm87/msgen/cmd/msgen/initialize"
	"github.com/adm87/msgen/cmd/msgen/update"
	"github.com/adm87/msgen/models"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	version = "0.0.0-unlreleased"
	args    = &models.MSGenArgs{
		MSGenYML:   ".",
		OutputPath: ".",
		AutoAccept: false,
	}
	cfg   = &models.MSGenCfg{}
	msgen = &cobra.Command{
		Use:           "msgen",
		Short:         "msgen is a code generator for microservices",
		Long:          "msgen is a code generator for microservices. Use it to manage initialization and updates of your microservice codebase.",
		Version:       version,
		SilenceErrors: true,
		SilenceUsage:  true,
		PersistentPreRunE: func(cmd *cobra.Command, _ []string) error {
			if err := validateArgs(args); err != nil {
				return fmt.Errorf("error validating arguments: %w", err)
			}
			if err := loadConfig(args.MSGenYML, cfg); err != nil {
				return fmt.Errorf("error loading config: %w", err)
			}
			if err := validateConfig(cfg); err != nil {
				return fmt.Errorf("error validating config: %w", err)
			}
			return nil
		},
	}
)

func validateArgs(args *models.MSGenArgs) error {
	path, err := filepath.Abs(args.MSGenYML)
	if err != nil {
		return fmt.Errorf("error resolving config path: %w", err)
	}
	args.MSGenYML = path

	info, err := os.Stat(args.MSGenYML)
	if os.IsNotExist(err) {
		return fmt.Errorf("config file not found at path: %s", args.MSGenYML)
	}
	if err != nil {
		return fmt.Errorf("error accessing config file: %w", err)
	}
	if !info.IsDir() {
		return fmt.Errorf("config path must be a directory containing msgen.yml: %s", args.MSGenYML)
	}

	path, err = filepath.Abs(args.OutputPath)
	if err != nil {
		return fmt.Errorf("error resolving output path: %w", err)
	}
	args.OutputPath = path

	return nil
}

func loadConfig(path string, cfg *models.MSGenCfg) error {
	viper.SetConfigName("msgen")
	viper.SetConfigType("yml")
	viper.AddConfigPath(path)

	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %w", err)
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("error unmarshalling config file: %w", err)
	}
	return nil
}

func validateConfig(cfg *models.MSGenCfg) error {
	if cfg.Service.Name == "" {
		return fmt.Errorf("service.name is required")
	}
	if cfg.Service.Module == "" {
		return fmt.Errorf("service.module is required")
	}
	if !regexp.MustCompile(`^[a-z0-9_.-]+(/[a-z0-9_.-]+)*$`).MatchString(cfg.Service.Module) {
		return fmt.Errorf("service.module must be a valid Go module path: %s", cfg.Service.Module)
	}
	return nil
}

func init() {
	msgen.AddCommand(initialize.Command(args, cfg))
	msgen.AddCommand(update.Command(args, cfg))

	msgen.PersistentFlags().StringVarP(&args.MSGenYML, "config", "c", args.MSGenYML, "path to msgen.yml configuration file")
	msgen.PersistentFlags().StringVarP(&args.OutputPath, "output", "o", args.OutputPath, "output path for generated code")
	msgen.PersistentFlags().BoolVarP(&args.AutoAccept, "yes", "y", args.AutoAccept, "automatically accept all prompts")
}

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer cancel()

	if err := msgen.ExecuteContext(ctx); err != nil {
		fmt.Printf("error executing command: %v\n", err)
		os.Exit(1)
	}
}
