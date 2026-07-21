package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type staleInstallation struct {
	manager string
	shim    commandShim
}

func runFix(skipConfirmation bool) {
	shims := findCommandShims()
	if len(shims) < 2 {
		fmt.Println("No duplicate doc-fetch command shims were found.")
		return
	}

	installations := staleInstallations(shims[1:])
	if len(installations) == 0 {
		fmt.Println("Duplicate shims were found, but their package managers could not be identified.")
		return
	}

	fmt.Println("Keeping the first PATH match:", shims[0].path)
	fmt.Println("Removing stale global installations:")
	for _, installation := range installations {
		fmt.Printf("  %s: %s\n", installation.manager, installation.shim.path)
	}

	if !skipConfirmation && !confirmFix() {
		fmt.Println("No installations were changed.")
		return
	}

	for _, installation := range installations {
		if err := uninstallGlobalPackage(installation.manager); err != nil {
			fmt.Printf("Failed to remove the %s installation at %s: %v\n", installation.manager, installation.shim.path, err)
			continue
		}
		fmt.Printf("Removed the stale %s installation at %s\n", installation.manager, installation.shim.path)
	}
}

func staleInstallations(shims []commandShim) []staleInstallation {
	seenManagers := make(map[string]struct{})
	installations := make([]staleInstallation, 0, len(shims))
	for _, shim := range shims {
		manager, ok := packageManagerForShim(shim.path)
		if !ok {
			continue
		}
		if _, seen := seenManagers[manager]; seen {
			continue
		}
		seenManagers[manager] = struct{}{}
		installations = append(installations, staleInstallation{manager: manager, shim: shim})
	}
	return installations
}

func packageManagerForShim(path string) (string, bool) {
	normalizedPath := strings.ToLower(filepath.ToSlash(path))
	if strings.Contains(normalizedPath, "/pnpm/") {
		return "pnpm", true
	}
	if strings.Contains(normalizedPath, "/npm/") {
		return "npm", true
	}
	return "", false
}

func confirmFix() bool {
	fmt.Print("Continue? [y/N]: ")
	input, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return false
	}
	input = strings.ToLower(strings.TrimSpace(input))
	return input == "y" || input == "yes"
}

func uninstallGlobalPackage(manager string) error {
	args := []string{"uninstall", "--global", "doc-fetch-cli"}
	if manager == "pnpm" {
		args = []string{"remove", "--global", "doc-fetch-cli"}
	}

	command := exec.Command(manager, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr
	return command.Run()
}
