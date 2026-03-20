package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter project name: ")
	projectName, _ := reader.ReadString('\n')
	projectName = trim(projectName)

	fmt.Print("Enter target directory: ")
	targetDir, _ := reader.ReadString('\n')
	targetDir = trim(targetDir)

	projectPath := filepath.Join(targetDir, projectName)

	// check if exists
	if _, err := os.Stat(projectPath); !os.IsNotExist(err) {
		fmt.Println("Project already exists. Abort.")
		return
	}

	// create base folder
	err := os.MkdirAll(projectPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	createStructure(projectPath)

	fmt.Println("Project created at:", projectPath)
}

func trim(input string) string {
	return string([]byte(input)[:len(input)-1])
}

func createStructure(base string) {
	dirs := []string{
		"cmd/app",
		"internal/domain/entity",
		"internal/domain/repository",
		"internal/application/service",
		"internal/ports/inbound",
		"internal/ports/outbound",
		"internal/adapters/inbound/http",
		"internal/adapters/outbound/persistence",
		"internal/infrastructure/database",
		"configs",
		"migrations",
	}

	for _, dir := range dirs {
		fullPath := filepath.Join(base, dir)
		err := os.MkdirAll(fullPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
