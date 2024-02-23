package cmd

import (
	"os"

	"github.com/spf13/cobra"
)


var rootCmd = &cobra.Command{
	Use:   "token-price-tracker",
	Short: "Get real time token price of any token on uniswap v3.",
	Long: `Paste address of token live on ethereum mainnet and see the magic.`,
}


func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}


