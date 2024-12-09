package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/faustocarva/solcomp/solc"
)

// Helper function to save a contract's data to files
func saveContractToFiles(contract *solc.Contract) error {
	// Save ABI
	err := os.WriteFile(contract.Name+"_abi.json", []byte(contract.AbiDefinition), 0644)
	if err != nil {
		return fmt.Errorf("error saving ABI for contract %s: %w", contract.Name, err)
	}

	// Save Deployment Bytecode
	err = os.WriteFile(contract.Name+"_deployment_bytecode.bin", []byte(contract.DeploymentBytecode), 0644)
	if err != nil {
		return fmt.Errorf("error saving Deployment Bytecode for contract %s: %w", contract.Name, err)
	}

	// Save Runtime Bytecode
	err = os.WriteFile(contract.Name+"_runtime_bytecode.bin", []byte(contract.RuntimeBytecode), 0644)
	if err != nil {
		return fmt.Errorf("error saving Runtime Bytecode for contract %s: %w", contract.Name, err)
	}

	return nil
}

func main() {
	// Define command-line parameters for Solidity file path and contract name
	solFilePath := flag.String("file", "", "Path to the Solidity (.sol) file")
	contractName := flag.String("contract", "", "Name of the contract to compile")
	flag.Parse()

	// Check if the required parameters are provided
	if *solFilePath == "" {
		fmt.Println("Usage: solcomp -file=<path-to-solidity-file> -contract=<contract-name>")
		os.Exit(1)
	}

	// Read the Solidity source code from the file
	sourceCode, err := os.ReadFile(*solFilePath)
	if err != nil {
		log.Fatalf("Failed to read the file: %v", err)
	}

	// Initialize the Solidity compiler
	compiler := solc.NewSolidityCompiler("/tmp/dogefuzz/")

	// Compile the Solidity source code
	result, err := compiler.CompileSource(*contractName, string(sourceCode))
	if err != nil {
		log.Fatalf("Compilation failed: %v", err)
	}

	if result.SingleContract != nil {
		err := saveContractToFiles(result.SingleContract)
		if err != nil {
			fmt.Println("Error:", err)
		}
	} else if result.AllContracts != nil {
		for _, contract := range result.AllContracts {
			err := saveContractToFiles(contract)
			if err != nil {
				fmt.Println("Error:", err)
			}
		}
	} else {
		fmt.Println("No contracts were returned.")
	}
}
