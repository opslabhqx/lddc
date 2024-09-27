package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func usage() {
	fmt.Println("lddc: missing file arguments")
	fmt.Println("Try `lddc <path to the binary> <path to copy the dependencies>`")
	os.Exit(1)
}

func main() {
	if len(os.Args) < 3 {
		usage()
	}

	binaryPath := os.Args[1]
	destDir := os.Args[2]

	if _, e := os.Stat(binaryPath); os.IsNotExist(e) {
		fmt.Printf("Not a valid input %s\n", binaryPath)
		os.Exit(1)
	}

	if _, e := os.Stat(destDir); os.IsNotExist(e) {
		fmt.Printf("No such directory %s, creating...\n", destDir)
		if e := os.MkdirAll(destDir, os.ModePerm); e != nil {
			fmt.Printf("Failed to create directory %s: %v\n", destDir, e)
			os.Exit(1)
		}
	}

	fmt.Printf("Collecting the shared library dependencies for %s...\n", binaryPath)

	cmd := exec.Command("ldd", binaryPath)
	output, e := cmd.CombinedOutput()
	if e != nil {
		fmt.Printf("Failed to execute ldd: %v\n", e)
		os.Exit(1)
	}

	deps := []string{}
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) > 2 && (strings.HasPrefix(parts[0], "/") || strings.HasPrefix(parts[2], "/")) {
			if strings.HasPrefix(parts[0], "/") {
				deps = append(deps, parts[0])
			}
			if strings.HasPrefix(parts[2], "/") {
				deps = append(deps, parts[2])
			}
		}
	}

	fmt.Printf("Copying the dependencies to %s\n", destDir)
	for _, dep := range deps {
		fmt.Printf("Copying %s to %s\n", dep, destDir)
		destPath := filepath.Join(destDir, filepath.Base(dep))
		if e := copyFile(dep, destPath); e != nil {
			fmt.Printf("Failed to copy %s to %s: %v\n", dep, destPath, e)
		}
	}

	fmt.Println("Done!")
}

func copyFile(src, dst string) error {
	input, e := os.ReadFile(src)
	if e != nil {
		return e
	}

	e = os.WriteFile(dst, input, 0644)
	if e != nil {
		return e
	}

	return nil
}
