package cmd

import (
	"os"
	"github.com/spf13/cobra"
)

var Version string = "1.0.0"

const StorePrefix string = "WALNUT"

var rootCmd = &cobra.Command {
	Use: "walnut",
	Short: "Credentials Manager",
	Version: Version,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {

}
