package wrapper

import "os"

// SchemaContract for encapsulation contract info
type SchemaContract struct {
	rpc           string
	address       string
	method        string
	hash          string
	schemaBody    []byte
	schemaName    string
	privateKeyHex string
}

// ContractBuilder for creating schemaContract
type ContractBuilder struct {
	contract *SchemaContract
}

// NewContractBuilder constructor for contract
func NewContractBuilder(rpc, address, method string) *ContractBuilder {
	return &ContractBuilder{contract: &SchemaContract{
		rpc:     rpc,
		address: address,
		method:  method,
	}}
}

// WithWithRPC is for build contract with rpc parameter
func (c *ContractBuilder) WithWithRPC(rpc string) *ContractBuilder {
	c.contract.rpc = rpc
	return c
}

// WithMethod is for build contract with method parameter
func (c *ContractBuilder) WithMethod(method string) *ContractBuilder {
	c.contract.method = method
	return c
}

// WithAddress is for build contract with address parameter
func (c *ContractBuilder) WithAddress(address string) *ContractBuilder {
	c.contract.address = address
	return c
}

// WithSchemaName is for build contract with schema name parameter
func (c *ContractBuilder) WithSchemaName(name string) *ContractBuilder {
	c.contract.schemaName = name
	return c
}

// WithSchemaBytes is for build contract with schema body parameter
func (c *ContractBuilder) WithSchemaBytes(body []byte) *ContractBuilder {
	c.contract.schemaBody = body
	return c
}

// WithSchemaHash is for build contract with schema hash parameter
func (c *ContractBuilder) WithSchemaHash(hash string) *ContractBuilder {
	c.contract.hash = hash
	return c
}

// Build is for getting contract
func (c *ContractBuilder) Build() *SchemaContract {
	return c.contract
}

// WithPrivateKey is for build contract with private key parameter
func (c *ContractBuilder) WithPrivateKey() *ContractBuilder {
	key := os.Getenv("PRIVATE_KEY")
	c.contract.privateKeyHex = key
	return c
}
