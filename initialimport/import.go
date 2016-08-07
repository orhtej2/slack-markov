package main

// Implements importing history from a Slack export from:
// https://my.slack.com/services/export

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
)

// Start an import from a slack export directory
// Does some basic error checking then imports the data in the background
func StartImport(dir *string) (err error) {
	log.Printf("Starting import from %s", *dir)

	// Does this directory exist? Get its contents
	contents, err := ioutil.ReadDir(*dir)
	if err != nil {
		log.Fatal(err)
		return err
	}
	log.Printf("Here %s", *dir)

	// Looks good, import each directory/channel in a goroutine
	log.Printf("Here %s", *dir)
	for _, file := range contents {
		log.Printf("Found %s", file.Name())
		if file.IsDir() {
			ImportDir(*dir + "/" + file.Name())
		}
	}

	// Write the state file
	err = markovChain.Save(stateFile)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Import complete. %d suffixes", len(markovChain.Chain))

	return nil
}

// Handles the import of a channel/directory
func ImportDir(dir string) {
	contents, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range contents {
		if !file.IsDir() {
			fo, err := os.Open(dir + "/" + file.Name())
			if err != nil {
				log.Fatal(err)
			} else {
				log.Printf("Importing %s", dir + "/" + file.Name())
			}
			defer fo.Close()

			scanner := bufio.NewScanner(fo)
			for scanner.Scan() {
				markovChain.Write(scanner.Text())
			}

			if err := scanner.Err(); err != nil {
				log.Fatal(err)
			}
		}
	}
}
