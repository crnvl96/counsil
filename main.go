package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
	"sync"

	"github.com/spf13/cobra"
)

var MISE_TOOLS = []string{
	"node@24.9.0",
	"go@1.25.3",
	"rust@1.90.0",
	"uv@0.9.4",
	"stylua@2.3.0",
	"prettier@3.6.2",
	"lua-language-server@3.15.0",
	"python@3.14.0",
	"gofumpt@0.9.1",
}

func installMiseTool(t string) error {
	home, _ := os.UserHomeDir()
	cmd := exec.Command("mise", "use", "--cd", home, "--force", "--pin", t)
	return cmd.Run()
}

func syncMiseTools() {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 5) // limit to 5 concurrent installs
	for _, t := range MISE_TOOLS {
		wg.Add(1)
		go func(tool string) {
			defer wg.Done()
			sem <- struct{}{}        // acquire semaphore
			defer func() { <-sem }() // release
			if err := installMiseTool(tool); err != nil {
				fmt.Printf("Error installing %s: %v\n", tool, err)
			} else {
				fmt.Printf("Successfully installed %s\n", tool)
			}
		}(t)
	}
	wg.Wait()
}

func main() {
	rootCmd := &cobra.Command{
		Use:   "counsil",
		Short: "Counsil CLI tool",
		Long:  `A CLI tool for managing development tools.`,
		Run: func(cmd *cobra.Command, args []string) {
			syncMiseTools()
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
