package main

import (
	"log"
	"os"
	"os/exec"
)

const (
	PNPM = "pnpm"
	BUN  = "bun"
	NPM  = "npm"
)

func main() {
	entries, err := os.ReadDir("./")
	packageManager := NPM
	args := os.Args

	if err != nil {
		log.Fatal("Something went wrong reading dir")
	}

	for _, name := range entries {
		if name.Name() == "pnpm-lock.yaml" {
			packageManager = PNPM
			break
		} else if name.Name() == "bun.lockb" || name.Name() == "bun.lock" {
			packageManager = BUN
			break
		}
	}

	cmd := exec.Command(packageManager, args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok {
			os.Exit(exitErr.ExitCode())
		}
		log.Fatal(err)
	}

}
