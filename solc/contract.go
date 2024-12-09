package solc

type Contract struct {
	Name               string
	AbiDefinition      string
	DeploymentBytecode string
	RuntimeBytecode    string
	Address            string
}

func NewContract(name, abiDefinition, deploymentBytecode string, runtimeBytecode string) *Contract {
	return &Contract{
		Name:               name,
		AbiDefinition:      abiDefinition,
		DeploymentBytecode: deploymentBytecode,
		RuntimeBytecode:    runtimeBytecode,
	}
}

type CompileResult struct {
	SingleContract *Contract   // Holds a single contract if one is requested
	AllContracts   []*Contract // Holds all contracts if no specific contract is requested
}
