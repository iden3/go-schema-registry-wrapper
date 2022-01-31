package wrapper

import (
	"context"
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var mockRPCURL = os.Getenv("RPC_URL")
var mockContractAddress = os.Getenv("TEST_ADDR") // before transition
var schemaName = "test"

// TODO test for schema saving

// func TestSaveSchema(t *testing.T) {
// 	b, _ := json.Marshal(JsonABI)
// 	tr, err := SaveSchema(context.Background(), mockRPCURL, mockContractAddress, schemaName, b)
// 	assert.Nil(t, err)
// 	assert.NotNil(t, tr)
// }

func TestSaveSchemaExists(t *testing.T) {

	b, _ := json.Marshal(JSONABI)

	tr, err := SaveSchema(context.Background(), mockRPCURL, mockContractAddress, schemaName, b)

	assert.NotNil(t, err)
	assert.Nil(t, tr)
	assert.Contains(t, err.Error(), "Schema already exists")
}

func TestVerifyInvalidRPC(t *testing.T) {

	invalidURL := "test://invalidurl1234.com"
	b, _ := json.Marshal(JSONABI)

	_, err := SaveSchema(context.Background(), invalidURL, mockContractAddress, schemaName, b)

	assert.NotNil(t, err)
	assert.Contains(t, err.Error(), errRPCClientCreationMessage)
}

func TestGetSchemaBytesByHash(t *testing.T) {
	ctx := context.Background()

	h, err := GetSchemaHashByName(ctx, mockRPCURL, mockContractAddress, schemaName)

	assert.Nil(t, err)
	assert.NotNil(t, h)

	hash := h.Hex()

	b, err := GetSchemaBytesByHash(ctx, mockRPCURL, mockContractAddress, hash)

	assert.Nil(t, err)
	assert.NotNil(t, b)

	var j string

	err = json.Unmarshal(b, &j)

	assert.Nil(t, err)

	assert.Equal(t, j, JSONABI)

}

func TestGetSchemaHashByName(t *testing.T) {
	h, err := GetSchemaHashByName(context.Background(), mockRPCURL, mockContractAddress, schemaName)

	assert.Nil(t, err)
	assert.NotNil(t, h)
}

func TestGetSchemaBytesByName(t *testing.T) {

	b, err := GetSchemaBytesByName(context.Background(), mockRPCURL, mockContractAddress, schemaName)

	assert.Nil(t, err)
	assert.NotNil(t, b)

	var j string

	err = json.Unmarshal(b, &j)
	assert.Nil(t, err)

	assert.Equal(t, j, JSONABI)
}
