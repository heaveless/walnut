package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/zalando/go-keyring"
)

var setCmd = &cobra.Command {
	Use: "set [context] [value]",
	Args: cobra.ExactArgs(2),
	Short: "Create or update credentials",
	Run: func(cmd *cobra.Command, args []string) {
		var context = args[0]
		var value = args[1]
		
		err := keyring.Set(StorePrefix, context, value)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func init() {
	rootCmd.AddCommand(setCmd)
}
