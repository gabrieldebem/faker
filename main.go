package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/gabrieldebem/faker/factories"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Hi, Hellcome to Faker CLI")
		return
	}

	var value string

	switch args[0] {
	case "cpf":
		hasMask := false

		if len(args) > 1 && args[1] == "--mask" {
			hasMask = true
		}

		value = factories.Cpf(hasMask)
		break
	case "cnpj":
		hasMask := false

		if len(args) > 1 && args[1] == "--mask" {
			hasMask = true
		}

		value = factories.Cnpj(hasMask)
		break
	default:
		value = "Invalid option"
		break
	}

	clipboard.WriteAll(value)

	fmt.Println(value)
}
