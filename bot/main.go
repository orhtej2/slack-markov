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
	httpPort       int
	numWords       int
	prefixLen      int
	stateFile      string
	responseChance int
	botUsername    string

	markovChain *utils.Chain
)

func init() {
	rand.Seed(time.Now().UnixNano()) // Seed the random number generator.
}

func main() {
	// Parse command-line options
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "usage: ./slack-markov -port=8000\n")
		flag.PrintDefaults()
	}

	flag.IntVar(&httpPort, "port", 8000, "The HTTP port on which to listen")
	flag.IntVar(&numWords, "words", 100, "Maximum number of words in the output")
	flag.IntVar(&prefixLen, "prefix", 2, "Prefix length in words")
	flag.IntVar(&responseChance, "responseChance", 10, "Percent chance to generate a response on each request")
	flag.StringVar(&stateFile, "stateFile", "state", "File to use for maintaining our markov chain state")
	flag.StringVar(&botUsername, "botUsername", "markov-bot", "The name of the bot when it speaks")

	flag.Parse()

	if httpPort == 0 {
		flag.Usage()
		os.Exit(2)
	}

	markovChain = utils.NewChain(prefixLen) // Initialize a new Chain.

	// Import into the chain
	// Rebuild the markov chain from state
	err := markovChain.Load(stateFile)
	if err != nil {
		//log.Fatal(err)
		log.Fatal("Could not load from '%s'. This may be expected.", stateFile)
	}

	// Start the webserver
	StartServer(httpPort)
}
