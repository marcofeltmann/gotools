package main

import (
	"fmt"
	"os"
	"path/filepath"
)

// Reflects sysexits.h as found at /Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/sysexits.h
const (
	posixExitOk            = 0  // Successful termination
	posixExitConfiguration = 78 // Some Configuraiton could not be read or set
)

func main() {
	configFilePath, outputDirectoryPath := parseArgumentsForPaths()

	fileStat, fileCheckError := os.Stat(configFilePath)
	if os.IsNotExist(fileCheckError) {
		fmt.Printf("No config file found at %s. Skipping.", configFilePath)
		// No file means no config. Sooo happy.
		os.Exit(posixExitOk)
	} else if fileStat.IsDir() {
		// File exists but is a directory. That is a misconfiguration!
		fmt.Printf("Unable to read file at %s: got directory instead of JSON file. Aborting!", configFilePath)
		os.Exit(posixExitConfiguration)
	}

	//SomeRoutine.SetConfigPath(configPath)
	fileStat, fileCheckError = os.Stat(outputDirectoryPath)

	if os.IsNotExist(fileCheckError) {
		// No file means no output. Bäm.
		fmt.Printf("Output directory at %s does not exist. Aborting!", outputDirectoryPath)
		os.Exit(posixExitConfiguration)
	} else if !fileStat.IsDir() {
		// File exists but is a directory. That is a misconfiguration!
		fmt.Printf("Unable to use output at %s: wanted directory. Aborting!", outputDirectoryPath)
		os.Exit(posixExitConfiguration)
	}

	// SomeRoutine.SetOutputDir(outputPath)

	// SomeRouinte.LinkContent()
}

func outputHelp() {
	println(".:Usage:.")
	println("Chatty version")

	/// go run sample.go configFile ~/.config/testbundler/sampleconfig.json targetDirectory dummyout/
	appCallPath := os.Args[0]

	appName := filepath.Base(appCallPath)

	fmt.Printf("\t%s configFile path/to/config.json targetDirectory path/to/output/directory\n", appName)
}

func parseArgumentsForPaths() (configPath string, outputDirectory string) {
	arguments := os.Args

	configPath = "configFile"
	outputDirectory = "outDir"

	if len(arguments) != 5 {
		outputHelp()
		os.Exit(posixExitOk)
	}

	expectIndexForConfigPath := -1
	expectIndexForOutputPath := -1

	for index, argument := range arguments {
		didFoundArgument := false
		switch argument {
		case "configFile":
			didFoundArgument = true
			expectIndexForConfigPath = index + 1
		case "targetDirectory":
			expectIndexForOutputPath = index + 1
			didFoundArgument = true
		}

		if !didFoundArgument {
			switch index {
			case expectIndexForConfigPath:
				configPath = arguments[expectIndexForConfigPath]
				if configPath == "configFile" || configPath == "targetDirectory" {
					outputHelp()
					os.Exit(posixExitConfiguration)
				}
				configPath, absPathError := filepath.Abs(configPath)
				if absPathError != nil {
					fmt.Printf("Unable to create path from configFile. input %#v produdes error %s", configPath, absPathError)
					os.Exit(posixExitConfiguration)
				}

			case expectIndexForOutputPath:
				outputDirectory = arguments[expectIndexForOutputPath]
				if outputDirectory == "configFile" || outputDirectory == "targetDirectory" {
					outputHelp()
					os.Exit(posixExitConfiguration)
				}
				outputDirectory, absPathError := filepath.Abs(outputDirectory)
				if absPathError != nil {
					fmt.Printf("Unable to create path from targetDirectory. Input %#v produces error %s", outputDirectory, absPathError)
					os.Exit(posixExitConfiguration)
				}
			}
		}
	}

	return
}
