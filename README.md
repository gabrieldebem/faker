## Faker CLI

This is a CLI to generate faker data.

Download the binary [here](https://github.com/gabrieldebem/faker/releases/latest)

Also, you can build this using `go build` or `make selfbuild`.

You will need `xclip` installed to use the clipboard functionality.
You can run `sudo apt install xclip` on linux or `brew install xclip` on MacOS.

Then you can use it like this:

## Generate Fake CPF
```bash
faker cpf {--mask}
```

## Generate Fake CNPJ
```bash
faker cnpj {--mask}
```

In case you want this document with mask `###.###.###-##` or `##.###.###/####-##`
You can pass the flag `--mask`.

