package main

import (
	"fmt"
	"github.com/atotto/clipboard"
	"github.com/gabrieldebem/faker/factories"
	"os"
	"os/exec"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		fmt.Println("Hi, Hellcome to Faker CLI")
		return
	}

	var value string

	switch args[0] {
		case "install":
			fmt.Println("Installing xclip...")

			bt, err := exec.Command("/bin/sh", "./commands/xclip.sh").Output()

			if err != nil {
				fmt.Println("An error occurred while installing xclip. Check if you already have it installed.")
				return
			}
			value = string(bt)

			fmt.Println(value)

			fmt.Println("Done!")

			return
		case "cpf":
			value = factories.Cpf()
		break
		case "cnpj":
			value = factories.Cnpj()
		break
		default:
			value = "Invalid option"
		break
	}

	clipboard.WriteAll(value)

	fmt.Println(value)
}

