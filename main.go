package main

import (
	"fmt"
	"os"
	"os/exec"
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
	"opencode@0.15.8",
	"bob@4.1.4",
	"ruff@0.14.1",
}

var YAY_TOOLS = []string{
	"calibre-bin",
}

var UV_TOOLS = []string{
	"pyright",
}

var NPM_TOOLS = []string{
	"vscode-langservers-extracted@latest",
	"typescript-language-server@latest",
}

var GO_TOOLS = []string{
	"golang.org/x/tools/gopls@latest",
}

func installMiseTool(t string) error {
	home, _ := os.UserHomeDir()
	cmd := exec.Command("mise", "use", "--cd", home, "--force", "--pin", t)
	return cmd.Run()
}

func installYayTool(t string) error {
	cmd := exec.Command("yay", "--sudoloop", "-Sy", "--needed", "--noconfirm", t)
	return cmd.Run()
}

func installUvTool(t string) error {
	cmd := exec.Command("uv", "tool", "install", "--upgrade", t)
	return cmd.Run()
}

func installNpmTool(t string) error {
	cmd := exec.Command("npm", "i", "-g", t+"@latest")
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
			syncTools(MISE_TOOLS, installMiseTool)
			syncTools(YAY_TOOLS, installYayTool)
			syncTools(UV_TOOLS, installUvTool)
			syncTools(NPM_TOOLS, installNpmTool)
			syncTools(GO_TOOLS, installGoTool)
		},
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}
}
