package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"

	"github.com/iden3/go-schema-registry-wrapper/wrapper"
	"github.com/pkg/errors"
)

func main() {
	// testSave()
	// testGetSchemaBytesByName()
	// testGetSchemaHashByName()
	testGetSchemaBytesByHash()
}

func testSave() {
	addr := os.Getenv("TEST_ADDR")
	url := os.Getenv("RPC_URL")
	ctx := context.Background()
	b, err := json.Marshal(wrapper.JsonABI)
	if err != nil {
		errors.New("fail to marshal json")
	}

	t, err := wrapper.SaveSchema(ctx, url, addr, "test", b)

	fmt.Println(t.Hash())

	if err != nil {
		errors.Wrap(err, "fail to marshal json")
		fmt.Println(err)
	}

}

func testGetSchemaBytesByName() {
	addr := os.Getenv("TEST_ADDR")
	url := os.Getenv("RPC_URL")
	ctx := context.Background()

	h, err := wrapper.GetSchemaBytesByName(ctx, url, addr, "test2")

	if err != nil {
		errors.Wrap(err, "fail to marshal json")
		fmt.Println(err)
	}

	var a interface{}

	err = json.Unmarshal(h, &a)

	if err != nil {
		fmt.Println(err.Error())
		errors.Wrap(err, "fail to marshal json")
	}

	fmt.Println(a)
}

func testGetSchemaHashByName() {
	addr := os.Getenv("TEST_ADDR")
	url := os.Getenv("RPC_URL")
	ctx := context.Background()

	h, err := wrapper.GetSchemaHashByName(ctx, url, addr, "test2")

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(h.Hex())
}

func testGetSchemaBytesByHash() {
	addr := os.Getenv("TEST_ADDR")
	url := os.Getenv("RPC_URL")
	ctx := context.Background()

	h, err := wrapper.GetSchemaHashByName(ctx, url, addr, "test2")

	if err != nil {
		fmt.Println(err)
	}

	hash := h.Hex()

	b, err := wrapper.GetSchemaBytesByHash(ctx, url, addr, hash)

	if err != nil {
		errors.Wrap(err, "fail to marshal json")
		fmt.Println(err)
	}

	var a interface{}

	err = json.Unmarshal(b, &a)

	if err != nil {
		fmt.Println(err.Error())
		errors.Wrap(err, "fail to marshal json")
	}

	fmt.Println(a)
}
