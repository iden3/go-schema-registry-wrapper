package wrapper

import (
	geth "github.com/ethereum/go-ethereum/mobile"
	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
)

const (
	getSchemaBytesByHashMethod = "getBytesByHash"
	getSchemaBytesByNameMethod = "getBytesByName"
	getSchemaHashByNameMethod  = "getHashByName"
	saveMethod                 = "save"
)

var errorSchemaNameNotFound = errors.New("schema name not found")
var errorDecodeSchemaHash = errors.New("can't decode schema hash")

// EncodeSchemaBytesByHash is used getting schema body by hash.
// hash is a hex string to retrieve schema body
func EncodeSchemaBytesByHash(hash string) ([]byte, error) {

	h, err := geth.NewHashFromHex(hash)

	if err != nil {
		return nil, err
	}

	data, err := ABI.Pack(getSchemaBytesByHashMethod, h.GetBytes())

	if err != nil {
		return nil, err
	}

	return data, nil
}

// DecodeSchemaBytesByHash is for decoding payload bytes for getSchemaBytesByHashMethod contract.
func DecodeSchemaBytesByHash(payload []byte) ([]byte, error) {
	return decodeBytes(payload, getSchemaBytesByHashMethod)
}

func decodeBytes(payload []byte, method string) ([]byte, error) {
	outputs, err := decode(payload, method)
	if err != nil {
		return nil, err
	}

	output, ok := outputs[0].([]byte)

	if !ok {
		return nil, errors.New("expected result is not []byte")
	}

	return output, nil
}

// EncodeSchemaHashByName is used getting schema hash by schema name.
func EncodeSchemaHashByName(name string) ([]byte, error) {
	return encode(name, getSchemaHashByNameMethod)
}

// DecodeSchemaHashByName is used getting decode hash by schema name.
func DecodeSchemaHashByName(payload []byte) (*common.Hash, error) {
	outputs, err := decode(payload, getSchemaHashByNameMethod)
	if err != nil {
		return nil, err
	}
	output, ok := outputs[0].([32]uint8)
	if !ok {
		return nil, errorDecodeSchemaHash
	}
	if isAllZeros(output) {
		return nil, errorSchemaNameNotFound
	}

	b := output[:]
	h := common.BytesToHash(b)

	return &h, nil
}

// isAllZeros check if array contains only zeros
func isAllZeros(arr [32]byte) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

// encode is helper for ABI.Pack function
func encode(payload, method string) ([]byte, error) {
	data, err := ABI.Pack(method, payload)

	if err != nil {
		return nil, err
	}

	return data, nil
}

func decode(payload []byte, method string) ([]interface{}, error) {
	outputs, err := ABI.Unpack(method, payload)

	if err != nil {
		return nil, err
	}

	return outputs, nil
}

// EncodeSchemaBytesByName is used to get encoded bytes from name for using in getSchemaBytesByNameMethod contract.
func EncodeSchemaBytesByName(name string) ([]byte, error) {
	return encode(name, getSchemaBytesByNameMethod)
}

// DecodeSchemaBytesByName is used to get decoded bytes from name for using in getSchemaBytesByNameMethod contract.
func DecodeSchemaBytesByName(payload []byte) ([]byte, error) {
	return decodeBytes(payload, getSchemaBytesByNameMethod)
}

// EncodeSaveTransaction is used to get encoded bytes from name and schema body for using in saveMethod contract.
func EncodeSaveTransaction(schemaName string, schemaBytes []byte) ([]byte, error) {

	bytesData, err := ABI.Pack(saveMethod, schemaName, schemaBytes)
	if err != nil {
		return nil, err
	}

	return bytesData, nil
}
