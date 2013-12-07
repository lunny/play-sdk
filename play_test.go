package play

import (
	"fmt"
	"testing"
)

var codes string = `
	package main

import "fmt"

func main() {
	fmt.Println("fdsafsafdsafsadf")
}
	`

func TestCompileCode(t *testing.T) {
	res, err := CompileCode(codes)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(res)
}

func TestFormat(t *testing.T) {
	res, err := Format(codes)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(res)
}

func TestCompile(t *testing.T) {
	res, err := Compile(codes)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(res)
}

func TestShare(t *testing.T) {
	formatCodes, err := Format(codes)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(formatCodes)

	res, err := Share(formatCodes)
	if err != nil {
		t.Error(err)
		return
	}

	fmt.Println(res)
}
