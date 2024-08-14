// ###########################
// ###########################
// ####  ================ ####
// #####  Nixy CLI Tool  #####
// ####  ================ ####
// ###########################
// ###########################

package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: nixy <command> <package>")
		fmt.Println("Commands:")
		fmt.Println("  install    Install a package using nix-env")
		fmt.Println("  uninstall  Uninstall a package using nix-env")
		fmt.Println("  virtual    Enter a Nix shell with a package installed using nix-shell -p")
		os.Exit(1)
	}

	command := os.Args[1]
	packageArg := os.Args[2]

	switch command {
	case "install":
		installPackage(packageArg)
	case "uninstall":
		uninstallPackage(packageArg)
	case "virtual":
		enterNixShell(packageArg, os.Args[3:])
	default:
		fmt.Println("Unknown command:", command)
		os.Exit(1)
	}
}

// func installNixENV() {
// 	fmt.Println("Install NixPKGS")
// }

func installPackage(packageArg string) {
	cmd := exec.Command("nix-env", "-iA", "nixos."+packageArg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error installing package %s: %v\n", packageArg, err)
		os.Exit(1)
	}
}

func uninstallPackage(packageArg string) {
	cmd := exec.Command("nix-env", "--uninstall", packageArg)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error uninstalling package %s: %v\n", packageArg, err)
		os.Exit(1)
	}
}

func enterNixShell(packageArg string, additionalArgs []string) {
	cmd := exec.Command("nix-shell", "-p", packageArg)

	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error entering Nix shell with package %s: %v\n", packageArg, err)
		os.Exit(1)
	}
}
