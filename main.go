package main

import (
	"fmt"
	"os"
	"os/exec"
	"sync"

	"github.com/spf13/cobra"
)

var MISE_TOOLS = []string{
	"node@latest",
	"go@latest",
	"rust@latest",
	"uv@latest",
	"stylua@latest",
	"lua-language-server@latest",
	"python@latest",
	"gofumpt@latest",
	"opencode@latest",
	"bob@latest",
	"ruff@latest",
	"prettier@latest",
}

var UV_TOOLS = []string{
	"pyright@latest",
}

var NPM_TOOLS = []string{
	"vscode-langservers-extracted@latest",
	"typescript-language-server@latest",
}

var GO_TOOLS = []string{
	"golang.org/x/tools/gopls@latest",
	// personal tools
	"github.com/crnvl96/dirt@latest",
}

func installMise() error {
	cmd := exec.Command("sh", "-c", "curl https://mise.run/bash | sh")
	return cmd.Run()
}

func installMiseTool(t string) error {
	home, _ := os.UserHomeDir()
	cmd := exec.Command("mise", "use", "--cd", home, "--force", "--pin", t)
	return cmd.Run()
}

func installUvTool(t string) error {
	cmd := exec.Command("uv", "tool", "install", "--upgrade", t)
	return cmd.Run()
}

func installNpmTool(t string) error {
	cmd := exec.Command("npm", "i", "-g", t)
	return cmd.Run()
}

func installGoTool(t string) error {
	cmd := exec.Command("go", "install", t)
	return cmd.Run()
}

type InstallFunc func(string) error

func syncTools(tools []string, install InstallFunc) {
	var wg sync.WaitGroup
	sem := make(chan struct{}, 10)
	for _, t := range tools {
		wg.Add(1)
		go func(tool string) {
			defer wg.Done()
			sem <- struct{}{}
			fmt.Printf("Starting install for %s\n", tool)
			if err := install(tool); err != nil {
				fmt.Printf("Error installing %s: %v\n", tool, err)
			} else {
				fmt.Printf("Successfully installed %s\n", tool)
			}
			<-sem
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
			if err := installMise(); err != nil {
				fmt.Printf("Error installing mise: %v\n", err)
			}

			syncTools(MISE_TOOLS, installMiseTool)
			syncTools(UV_TOOLS, installUvTool)
			syncTools(NPM_TOOLS, installNpmTool)
			syncTools(GO_TOOLS, installGoTool)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
