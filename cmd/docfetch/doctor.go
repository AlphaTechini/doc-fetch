package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

type commandShim struct {
	path string
}

func runDoctor() {
	fmt.Println("DocFetch doctor")
	fmt.Println("Running version:", version)
	if executable, err := os.Executable(); err == nil {
		fmt.Println("Running binary:", executable)
	}

	shims := findCommandShims()
	if len(shims) == 0 {
		fmt.Println("No doc-fetch command shims were found in PATH.")
		return
	}

	fmt.Println("Command shims in PATH:")
	for index, shim := range shims {
		label := "duplicate"
		if index == 0 {
			label = "first match"
		}
		fmt.Printf("  [%s] %s\n", label, shim.path)
	}

	if len(shims) > 1 {
		fmt.Println("Multiple global installations can cause an older command to win when PATH order changes.")
		fmt.Println("Remove the stale installation with the package manager that owns it:")
		fmt.Println("  npm uninstall -g doc-fetch-cli")
		fmt.Println("  pnpm remove -g doc-fetch-cli")
	}
}

func findCommandShims() []commandShim {
	names := []string{"doc-fetch"}
	if runtime.GOOS == "windows" {
		names = []string{"doc-fetch.ps1", "doc-fetch.cmd", "doc-fetch.exe", "doc-fetch"}
	}

	seenDirectories := make(map[string]struct{})
	shims := make([]commandShim, 0)
	for _, directory := range filepath.SplitList(os.Getenv("PATH")) {
		if directory == "" {
			continue
		}

		directory = filepath.Clean(directory)
		if _, seen := seenDirectories[directory]; seen {
			continue
		}

		for _, name := range names {
			path := filepath.Join(directory, name)
			if info, err := os.Stat(path); err == nil && !info.IsDir() {
				shims = append(shims, commandShim{path: path})
				seenDirectories[directory] = struct{}{}
				break
			}
		}
	}
	return shims
}
