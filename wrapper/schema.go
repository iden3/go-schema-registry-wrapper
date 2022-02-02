package wrapper

import (
	"context"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	getSchemaBytesByHashMethod = "getBytesByHash"
	getSchemaBytesByNameMethod = "getBytesByName"
	getSchemaHashByNameMethod  = "getHashByName"
	saveMethod                 = "save"
)

// EncodeSchemaBytesByHash is used getting schema body by hash.
// hash string to retrieve schema body
func EncodeSchemaBytesByHash(hash string) ([]byte, error) {

	b := common.FromHex(hash)
	var arr [32]uint8
	copy(arr[:], b[:32])
	data, err := ABI.Pack(getSchemaBytesByHashMethod, arr)

	if err != nil {
		return nil, err
	}

	return data, nil
}

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
// name - schema name
func EncodeSchemaHashByName(name string) ([]byte, error) {
	return encode(getSchemaHashByNameMethod, name)
}

// DecodeSchemaHashByName is used getting decode hash by schema name.
// name - schema name
func DecodeSchemaHashByName(payload []byte) (*common.Hash, error) {
	outputs, err := decode(payload, getSchemaHashByNameMethod)
	if err != nil {
		return nil, err
	}
	output, ok := outputs[0].([32]uint8)
	if !ok {
		return nil, errors.New("expected result is not hash")
	}
	if isAllZeros(output) {
		return nil, errors.New("unexpected zero hash value")
	}

	b := output[:]
	h := common.BytesToHash(b)

	return &h, nil
}

func isAllZeros(arr [32]byte) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			return false
		}
	}
	return true
}

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

// EncodeSchemaBytesByName is used to get schema by name.
// name is schema name
func EncodeSchemaBytesByName(name string) ([]byte, error) {
	return encode(getSchemaBytesByNameMethod, name)
}

// DecodeSchemaBytesByName is decoding schema bytes
func DecodeSchemaBytesByName(payload []byte) ([]byte, error) {
	return decodeBytes(payload, getSchemaBytesByNameMethod)
}

func GetSaveTransaction(ctx context.Context, client *ethclient.Client, nonce uint64, contractAddress string, schemaName string, schemaBytes []byte) (*types.Transaction, error) {

	bytesData, err := ABI.Pack(saveMethod, schemaName, schemaBytes)
	if err != nil {
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}
	var gasLimit uint64 = 0

	tx := types.NewTransaction(nonce, common.HexToAddress(contractAddress), nil, gasLimit, gasPrice, bytesData)

	return tx, nil
}

//func callSave(ctx context.Context, client *ethclient.Client, crt *SchemaContract) (*types.Transaction, error) {
//
//	privateKey, err := crypto.HexToECDSA(crt.privateKeyHex)
//
//	if err != nil {
//		return nil, err
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		return nil, errors.New("error casting public key to ECDSA")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//
//	nonce, err := client.PendingNonceAt(ctx, fromAddress)
//	if err != nil {
//		return nil, err
//	}
//
//	gasPrice, err := client.SuggestGasPrice(ctx)
//	if err != nil {
//		return nil, err
//	}
//
//	address := common.HexToAddress(crt.address)
//
//	id, err := client.NetworkID(ctx)
//	if err != nil {
//		return nil, err
//	}
//	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, id)
//
//	if err != nil {
//		return nil, err
//	}
//
//	auth.Nonce = big.NewInt(int64(nonce))
//	auth.Value = big.NewInt(0) // in wei
//	auth.GasLimit = 0          // Gas limit to set for the transaction execution (0 = estimate)
//	auth.GasPrice = gasPrice
//
//	abiJ, err := abi.JSON(strings.NewReader(JSONABI))
//	if err != nil {
//		return nil, err
//	}
//
//	boundContract := bind.NewBoundContract(address, abiJ, client, client, client)
//	t, err := boundContract.Transact(auth, crt.method, crt.schemaName, crt.schemaBody)
//
//	if err != nil {
//		return nil, err
//	}
//
//	return t, nil
//
//}
