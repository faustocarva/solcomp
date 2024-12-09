package main

import (
	"flag"
	"fmt"
	"github.com/faustocarva/solcomp/solc"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	// Define command-line parameters for Solidity file path and contract name
	solFilePath := flag.String("file", "", "Path to the Solidity (.sol) file")
	contractName := flag.String("contract", "", "Name of the contract to compile")
	flag.Parse()

	// Check if the required parameters are provided
	if *solFilePath == "" || *contractName == "" {
		fmt.Println("Usage: go run main.go -file=<path-to-solidity-file> -contract=<contract-name>")
		os.Exit(1)
	}

	// Read the Solidity source code from the file
	sourceCode, err := ioutil.ReadFile(*solFilePath)
	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	// Initialize the Solidity compiler
	compiler := solc.NewSolidityCompiler("/tmp/dogefuzz/")

	// Compile the Solidity source code
	contract, err := compiler.CompileSource(*contractName, string(sourceCode))
	if err != nil {
		log.Fatalf("Compilation failed: %v", err)
	}

	// Output the compiled contract
	fmt.Printf("Compiled Contract:\n%s\n", *contract)
}
