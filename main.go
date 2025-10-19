package main

import (
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	rootCmd := &cobra.Command{
		Use:   "counsil",
		Short: "Counsil CLI tool",
		Long:  `A CLI tool for managing development tools.`,
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Counsil CLI executed")
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
