package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var getCmd = &cobra.Command {
	Use: "get",
	Args: cobra.ExactArgs(1),
	Short: "Get credentials",
	Run: func(cmd *cobra.Command, args []string) {
		var context = args[0]

		value, err := keyring.Get(StorePrefix, context)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		fmt.Println(value)
	},
}

func init() {
	rootCmd.AddCommand(getCmd)
}
