package cli

import (
	"fmt"

	"refr/internal/config"

	"github.com/spf13/cobra"
)

var cfgFile string
var appConfig *config.Config

var rootCmd = &cobra.Command{
	Use:   "refr",
	Short: "Personal desktop reference for system commands",
	Long:  "A keyboard-launched CLI for browsing system command references with syntax-highlighted content.",
}

func init() {
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default: ~/.config/refr/config.toml)")
	rootCmd.PersistentPreRunE = func(cmd *cobra.Command, args []string) error {
		cfg, err := config.Load(cfgFile)
		if err != nil {
			return fmt.Errorf("loading config: %w", err)
		}
		appConfig = cfg
		return nil
	}
}

func Execute() error {
	return rootCmd.Execute()
}
