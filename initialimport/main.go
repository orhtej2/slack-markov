package main

// Main entry point for the app. Handles command-line options, starts the web
// listener and any import, etc

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
	"github.com/orhtej2/slack-markov/utils"
)

var (
	prefixLen      int
	stateFile      string
	markovChain *utils.Chain
)

func init() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
}

func main() {
	// Parse command-line options
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: ./initialimport -importDir=directory\n")
		flag.PrintDefaults()
	}
	flag.IntVar(&prefixLen, "prefix", 2, "Prefix length in words")
	var importDir = flag.String("importDir", "", "The directory of a Slack export")
	flag.StringVar(&stateFile, "stateFile", "state", "File to use for maintaining our markov chain state")

	flag.Parse()

	if *importDir == "" {
		flag.Usage()
		os.Exit(2)
	}

	markovChain = utils.NewChain(prefixLen) // Initialize a new Chain.

	// Import into the chain
	err := StartImport(importDir)
	if err != nil {
		log.Fatal(err)
	} else { 
		log.Print("All done!")
	}
}
