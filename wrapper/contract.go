package wrapper

import "os"

// schemaContract for encapsulation contract info
type schemaContract struct {
	rpc           string
	address       string
	method        string
	hash          string
	schemaBody    []byte
	schemaName    string
	privateKeyHex string
}

// contractBuilder for creating schemaContract
type contractBuilder struct {
	contract *schemaContract
}

//NewContractBuilder constructor for contract
func NewContractBuilder(rpc, address, method string) *contractBuilder {
	return &contractBuilder{contract: &schemaContract{
		rpc:     rpc,
		address: address,
		method:  method,
	}}
}

// WithWithRpc is for build contract with rpc parameter
func (c *contractBuilder) WithWithRpc(rpc string) *contractBuilder {
	c.contract.rpc = rpc
	return c
}

// WithMethod is for build contract with method parameter
func (c *contractBuilder) WithMethod(method string) *contractBuilder {
	c.contract.method = method
	return c
}

// WithAddress is for build contract with address parameter
func (c *contractBuilder) WithAddress(address string) *contractBuilder {
	c.contract.address = address
	return c
}

// WithAddress is for build contract with schema name parameter
func (c *contractBuilder) WithSchemaName(name string) *contractBuilder {
	c.contract.schemaName = name
	return c
}

// WithSchemaBytes is for build contract with schema body parameter
func (c *contractBuilder) WithSchemaBytes(body []byte) *contractBuilder {
	c.contract.schemaBody = body
	return c
}

// WithSchemaHash is for build contract with schema hash parameter
func (c *contractBuilder) WithSchemaHash(hash string) *contractBuilder {
	c.contract.hash = hash
	return c
}

// Build is for getting contract
func (c *contractBuilder) Build() *schemaContract {
	return c.contract
}

// WithPrivateKey is for build contract with private key parameter
func (c *contractBuilder) WithPrivateKey() *contractBuilder {
	key := os.Getenv("PRIVATE_KEY")
	c.contract.privateKeyHex = key
	return c
}
