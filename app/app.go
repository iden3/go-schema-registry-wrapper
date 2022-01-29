package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/iden3/go-schema-registry-wrapper/wrapper"
	"github.com/pkg/errors"
)

func main() {
	testSave()
	testGetSchemaBytesByName()
}

func testSave() {
	addr := "0xC93143C6dd477268133CDFD4ba117aC8293362F2"
	url := "https://ropsten.infura.io/v3/6796c7488d784b0bb3ffafd2e2696fef"
	ctx := context.Background()
	b, err := json.Marshal(wrapper.JsonABI)
	if err != nil {
		errors.New("fail to marshal json")
	}

	err = wrapper.SaveSchema(ctx, url, addr, "test2", b)

	if err != nil {
		errors.Wrap(err, "fail to marshal json")
		fmt.Println(err)
	}

}

func testGetSchemaBytesByName() {
	addr := "0xC93143C6dd477268133CDFD4ba117aC8293362F2"
	url := "https://ropsten.infura.io/v3/6796c7488d784b0bb3ffafd2e2696fef"
	ctx := context.Background()

	h, err := wrapper.GetSchemaBytesByName(ctx, url, addr, "test2")

	if err != nil {
		errors.Wrap(err, "fail to marshal json")
		fmt.Println(err)
	}

	var a interface{}

	err = json.Unmarshal(h, &a)

	fmt.Println(err.Error())

	if err != nil {
		errors.Wrap(err, "fail to marshal json")
	}

	fmt.Println(h)
}
