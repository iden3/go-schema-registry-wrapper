package wrapper

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"

	"github.com/pkg/errors"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	getSchemaBytesByHashMethod         = "getBytesByHash"
	getSchemaBytesByNameMethod         = "getBytesByName"
	getSchemaHashByNameMethod          = "getHashByName"
	saveMethod                         = "save"
	errRPCClientCreationMessage        = "couldn't create rpc client"
	errCallArgumentEncodedErrorMessage = "wrong arguments were provided"
)

// SaveSchema is used to save schema to ethereum blockchain
// rpcURL - url to connect to the blockchain
// contractAddress is contract address
// sName schema name
// body json schema bytes
func SaveSchema(ctx context.Context, rpcURL, contractAddress, sName string, body []byte) (*types.Transaction, error) {

	c := NewContractBuilder(rpcURL, contractAddress, saveMethod).
		WithSchemaName(sName).
		WithSchemaBytes(body).
		WithPrivateKey().
		Build()

	cl, err := ethclient.DialContext(ctx, c.rpc)
	if err != nil {
		return nil, errors.Wrap(err, errRPCClientCreationMessage)
	}
	defer cl.Close()

	t, err := callSave(ctx, cl, c)

	if err != nil {
		return nil, err
	}
	return t, nil
}

// GetSchemaBytesByHash is used getting schema body by hash
// rpcURL - url to connect to the blockchain
// contractAddress is contract address
// hash hash string to retrieve schema body
func GetSchemaBytesByHash(ctx context.Context, rpcURL, contractAddress, hash string) ([]byte, error) {

	c := NewContractBuilder(rpcURL, contractAddress, getSchemaBytesByHashMethod).
		WithSchemaHash(hash).
		Build()

	outputs, err := contractCall(ctx, c)

	if err != nil {
		return nil, err
	}

	output, ok := outputs[0].([]byte)

	if !ok {
		return nil, errors.New("expected result is not []byte")
	}

	return output, nil
}

// GetSchemaHashByName is used getting schema hash by schema name
// rpcURL - url to connect to the blockchain
// contractAddress - contract address
// name - schema name
func GetSchemaHashByName(ctx context.Context, rpcURL, contractAddress, name string) (*common.Hash, error) {
	c := NewContractBuilder(rpcURL, contractAddress, getSchemaHashByNameMethod).
		WithSchemaName(name).
		Build()

	outputs, err := contractCall(ctx, c)

	if err != nil {
		return nil, err
	}

	output, ok := outputs[0].([32]uint8)
	if !ok {
		return nil, errors.New("expected result is not common.Hash")
	}

	b := output[:]
	h := common.BytesToHash(b)

	return &h, nil
}

// GetSchemaBytesByName is used to get schema by name
// rpcURL - url to connect to the blockchain
// contractAddress is contract address
// name is schema name
func GetSchemaBytesByName(ctx context.Context, rpcURL, contractAddress, name string) ([]byte, error) {
	c := NewContractBuilder(rpcURL, contractAddress, getSchemaBytesByNameMethod).
		WithSchemaName(name).
		Build()

	outputs, err := contractCall(ctx, c)

	if err != nil {
		return nil, err
	}

	output, ok := outputs[0].([]byte)

	if !ok {
		return nil, errors.New("expected result is not []byte")
	}

	return output, nil
}

func contractCall(ctx context.Context, crt *SchemaContract) ([]interface{}, error) {

	c, err := ethclient.DialContext(ctx, crt.rpc)
	if err != nil {
		return nil, errors.Wrap(err, errRPCClientCreationMessage)
	}
	defer c.Close()

	var data []byte
	switch crt.method {
	case getSchemaBytesByHashMethod:
		b := common.FromHex(crt.hash)
		var arr [32]uint8
		copy(arr[:], b[:32])
		data, err = StateABI.Pack(crt.method, arr)
	case getSchemaHashByNameMethod:
		data, err = StateABI.Pack(crt.method, crt.schemaName)
	case getSchemaBytesByNameMethod:
		data, err = StateABI.Pack(crt.method, crt.schemaName)
	default:
		return nil, errors.Errorf("Not supported method for ethereum contract: %s", crt.method)
	}

	if data == nil {
		return nil, errors.Wrapf(err, "%s contract info: %v", errCallArgumentEncodedErrorMessage, crt)
	}

	addr := common.HexToAddress(crt.address)

	res, err := c.CallContract(ctx, ethereum.CallMsg{
		To:   &addr,
		Data: data,
	}, nil)

	if err != nil {
		return nil, err
	}

	outputs, err := StateABI.Unpack(crt.method, res)

	if err != nil {
		return nil, err
	}

	return outputs, nil
}

func callSave(ctx context.Context, client *ethclient.Client, crt *SchemaContract) (*types.Transaction, error) {

	privateKey, err := crypto.HexToECDSA(crt.privateKeyHex)

	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, errors.New("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		return nil, err
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		return nil, err
	}

	address := common.HexToAddress(crt.address)

	id, err := client.NetworkID(ctx)
	if err != nil {
		return nil, err
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, id)

	if err != nil {
		return nil, err
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = 0          // Gas limit to set for the transaction execution (0 = estimate)
	auth.GasPrice = gasPrice

	abiJ, err := abi.JSON(strings.NewReader(JSONABI))
	if err != nil {
		return nil, err
	}

	boundContract := bind.NewBoundContract(address, abiJ, client, client, client)
	t, err := boundContract.Transact(auth, crt.method, crt.schemaName, crt.schemaBody)

	if err != nil {
		return nil, err
	}

	return t, nil

}
