package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

// DefaultCommands is a type that defines commands that should be run depending on prefered javascript bundler
type DefaultCommands struct {
	Build     string
	Start     string
	Serve     string
	Install   string
	Directory string
}

func setupBuildCommands(isYarn bool, pwd string) DefaultCommands {
	var buildCommands DefaultCommands
	if isYarn {
		buildCommands = DefaultCommands{
			Build:     "build",
			Start:     "start",
			Serve:     "serve",
			Install:   "",
			Directory: pwd,
		}
	} else {
		buildCommands = DefaultCommands{
			Build:     "run build",
			Start:     "start",
			Serve:     "run serve",
			Install:   "install",
			Directory: "",
		}
	}
	return buildCommands
}

// TODO FIX THIS BUNDLER CONNECTION NOT WORKING PROPERLY
func startReactApp() {
	os.Chdir("client")
	isYarn := true

	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal("Error geting current working directory")
	}

	// absolute path to yarn executible
	p, err := exec.LookPath("yarn")
	if err != nil {
		isYarn = false
		// absolute path to npm executible
		p, err = exec.LookPath("npm")
		if err != nil {
			panic("You need to have either yarn or npm installed. If both are installed yarn will be preferred over")
		}
	}
	fmt.Printf("\n[+]Found Javascript package bundler at: %v\n", p)
	buildCommands := setupBuildCommands(isYarn, pwd)
	fmt.Println("Installing necessary Javascript packages in client folder...")
	fmt.Printf("\n%v\n", buildCommands.Directory)
	cwdFlag := fmt.Sprintf("--cwd=\"%v\"", buildCommands.Directory)
	cmd := exec.Command(p, buildCommands.Install, cwdFlag)
	err = cmd.Run()
	if err != nil {
		fmt.Printf("Error installing and checking Javascript packages:\n%v", err)
	}
	fmt.Println("\n[+]Building necessary Javascript packages in client folder...")
	cmd = exec.Command(p, buildCommands.Build, cwdFlag)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error building Javascript packages:\n%v\nExiting...", err))
	}
	fmt.Println("\n[+]Serving Front End...")
	cmd = exec.Command(p, buildCommands.Serve, cwdFlag)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error serving Javascript packages:\n%v\nExiting...", err))
	}
}
