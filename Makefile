# Setup package manager
deps:
	@echo "Set your GOPATH then add this to .bashrc, .zshrc, etc"
	@echo "export PATH=\"\$$PATH:\$$GOPATH/bin\""
	-@go get github.com/tools/godep

# Build the samples
samples: main
	./fddp convert -i samples/sample.html samples/sample-indent.json
	./fddp convert samples/sample.html samples/sample.json