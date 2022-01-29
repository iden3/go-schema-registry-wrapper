package wrapper

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

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
	errRPCClientCreationMessage        = "couldn't create rpc client: %s"
	errCallArgumentEncodedErrorMessage = "wrong arguments were provided"
)

// SaveSchema is used to save schema to ethereum blockchain
// rpcURL - url to connect to the blockchain
// contractAddress is contract address
// sName schema name
// body json schema bytes
func SaveSchema(ctx context.Context, rpcURL, contractAddress, sName string, body []byte) error {

	c := NewContractBuilder(rpcURL, contractAddress, saveMethod).
		WithSchemaName(sName).
		WithSchemaBytes(body).
		Build()

	cl, err := ethclient.DialContext(ctx, c.rpc)
	if err != nil {
		return errors.Errorf(errRPCClientCreationMessage, err.Error())
	}
	defer cl.Close()

	err = callSave(ctx, cl, c)

	if err != nil {
		return err
	}
	return nil
}

// GetSchemaBytesByHash is used geting schema body by hash
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

// GetSchemaBytesByHash is used geting schema hash by schema name
// rpcURL - url to connect to the blockchain
// contractAddress - contract address
// name - schema name
func GetSchemaHashByName(ctx context.Context, rpcURL, contractAddress, name string) (*common.Hash, error) {
	c := NewContractBuilder(rpcURL, contractAddress, getSchemaHashByNameMethod).
		WithSchemaHash(name).
		Build()

	outputs, err := contractCall(ctx, c)

	if err != nil {
		return nil, err
	}

	output, ok := outputs[0].([]byte)

	h := common.BytesToHash(output)

	if !ok {
		return nil, errors.New("expected result is not common.Hash")
	}

	return &h, nil
}

// VerifyState is used to verify identity state
// rpcURL - url to connect to the blockchain
// contractAddress is contract address
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

func contractCall(ctx context.Context, crt *schemaContract) ([]interface{}, error) {

	c, err := ethclient.DialContext(ctx, crt.rpc)
	if err != nil {
		return nil, errors.Errorf(errRPCClientCreationMessage, err.Error())
	}
	defer c.Close()

	var data []byte
	switch crt.method {
	case saveMethod:
		data, err = StateABI.Pack(crt.method, crt.schemaName, crt.schemaBody)
	case getSchemaBytesByHashMethod:
		data, err = StateABI.Pack(crt.method, crt.hash)
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

func callSave(ctx context.Context, client *ethclient.Client, crt *schemaContract) error {
	privateKey, err := crypto.HexToECDSA("1833d74a66dd5b6a9243e740de14f8f47c18bef101adb9b06103a0f882bbff4f")
	if err != nil {
		return err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(fromAddress.Hex())
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Fatal(err)
	}

	gasLimit := client.EstimateGas(ctx, )

	// gasCap, err := client.SuggestGasTipCap(ctx)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	inc := new(big.Int).Set(gasPrice)
	inc.Div(inc, new(big.Int).SetUint64(10))
	gasPrice.Add(gasPrice, inc)
	fmt.Fprintln("Transaction metadata %s", "gasPrice", gasPrice)
	auth := bind.NewKeyedTransactor(privateKey)

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = gasLimit
	auth.GasPrice = gasPrice

	address := common.HexToAddress(crt.address)
	instance, err := NewWrapper(address, client)
	if err != nil {
		log.Fatal(err)
	}
	t, err := instance.Save(auth, crt.schemaName, crt.schemaBody)
	fmt.Println(t)

	if err != nil {
		log.Fatal(err)
	}
	return nil

}
