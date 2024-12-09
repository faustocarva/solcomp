# Solcomp - Solidity Compiler

solcomp is a Go-based tool that compiles Solidity contracts from a given file and saves the resulting contract details (ABI and bytecode) into files.

## Requirements:
- Go 1.18+ (for building and running the project)

## Installation:
1. Clone the repository:
```
   git clone https://github.com/faustocarva/solcomp.git
   cd solcomp
```

3. Build the project:
```
   make build
```
   This will produce the solcomp binary in the ./bin directory.

## Usage:
To compile a Solidity contract, use the following command:
```
   ./bin/solcomp -file=<path-to-solidity-file> -contract=<contract-name>
```

## Parameters:
- -file: Path to the Solidity .sol file to be compiled.
- -contract: (optional) The name of the contract within the Solidity file to compile. 

## Example:
```
   ./bin/solcomp -file=./contracts/MyContract.sol -contract=MyContract
```

This command will compile the MyContract contract in the MyContract.sol file and output the contract's ABI and bytecode to the current directory.
NOTE: If no contract name is provided, the tool will output the ABI and bytecode (BIN) files for all contracts in the current directory.

## File Output:
- The compiled ABI and bytecode will be saved in files named \<contract-name\>_abi.json, \<contract-name\>_deployment_bytecode.bin, and \<contract-name\>_runtime_bytecode.bin.

## Development:

Running the project locally:
To run the project without building the binary, you can use:
```
   go run main.go -file=<path-to-solidity-file> -contract=<contract-name>
```
