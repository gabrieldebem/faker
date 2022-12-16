selfbuild:
	@go build -o faker -v
	@mkdir -p ~/.local/bin
	@mv -f faker ~/.local/bin/faker