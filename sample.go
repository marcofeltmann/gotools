package main

import (
	"log"
	"os"
	"os/user"
	"path/filepath"
)


// Reflects sysexits.h as found at /Library/Developer/CommandLineTools/SDKs/MacOSX10.15.sdk/usr/include/sysexits.h
const (
	posixExitOk = 0				// Successful termination
	posixExitConfiguration = 78	// Some Configuraiton could not be read or set
)

func main() {
	currentUser, userRetrievalError := user.Current()
	if userRetrievalError != nil {
		log.Panicf("Unable to retrieve user: error %s. Panicking!", userRetrievalError)
	}

	userHomeDirectory := currentUser.HomeDir
	configPath := filepath.Join(userHomeDirectory, ".config/testbundler/sampleconfig.json")

	fileStat, fileCheckError := os.Stat(configPath)
	if os.IsNotExist(fileCheckError) {
		log.Printf("No config file found at %s. Skipping.", configPath)
		// No file means no config. Sooo happy.
		os.Exit(posixExitOk)
	} else if fileStat.IsDir() {
		// File exists but is a directory. That is a misconfiguration!
		log.Printf("Unable to read file at %s: got directory instead of JSON file. Aborting!", configPath)
		os.Exit(posixExitConfiguration)
	}

	//SomeRoutine.SetConfigPath(configPath)

	outputPath := "./dummyout"
	outputPath, outputDirectoryAbsPathError := filepath.Abs(outputPath)
	if outputDirectoryAbsPathError != nil {
		log.Panicf("Unable to get absolute path fo output directroy %s: error %s. Panicking!", outputPath, outputDirectoryAbsPathError)
	}

	fileStat, fileCheckError = os.Stat(outputPath)

	if os.IsNotExist(fileCheckError) {
		// No file means no output. Bäm.
		log.Printf("Output directory at %s does not exist. Aborting!", outputPath)
		os.Exit(posixExitConfiguration)
	} else if !fileStat.IsDir() {
		// File exists but is a directory. That is a misconfiguration!
		log.Printf("Unable to use output at %s: wanted directory. Aborting!", outputPath)
		os.Exit(posixExitConfiguration)
	}

	// SomeRoutine.SetOutputDir(outputPath)

	log.Printf("\nsomeRoutine.SetConfigFilePath(\n\t%#v\n)\n\nsomeRoutine.SetOutputDir(\n\t%#v\n)", configPath, outputPath)
}