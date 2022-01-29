package wrapper

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

const JsonABI = `[
	{
		"inputs": [
			{
				"internalType": "bytes32",
				"name": "hash",
				"type": "bytes32"
			}
		],
		"name": "getBytesByHash",
		"outputs": [
			{
				"internalType": "bytes",
				"name": "",
				"type": "bytes"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "name",
				"type": "string"
			}
		],
		"name": "getBytesByName",
		"outputs": [
			{
				"internalType": "bytes",
				"name": "",
				"type": "bytes"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "name",
				"type": "string"
			}
		],
		"name": "getHashByName",
		"outputs": [
			{
				"internalType": "bytes32",
				"name": "",
				"type": "bytes32"
			}
		],
		"stateMutability": "view",
		"type": "function"
	},
	{
		"inputs": [
			{
				"internalType": "string",
				"name": "schemaName",
				"type": "string"
			},
			{
				"internalType": "bytes",
				"name": "schemaBody",
				"type": "bytes"
			}
		],
		"name": "save",
		"outputs": [],
		"stateMutability": "nonpayable",
		"type": "function"
	}
]`

var StateABI abi.ABI

// nolint // common approach to register default supported circuit
func init() {
	var err error
	StateABI, err = abi.JSON(strings.NewReader(JsonABI))
	if err != nil {
		panic(fmt.Errorf("can't parse state contract abi %v", err))
	}
}
