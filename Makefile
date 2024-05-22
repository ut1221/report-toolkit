# Makefile for building Go project targeting multiple OS

# The name of the output binary
OUTPUT_BINARY=app

# The main entry point of the program
MAIN=cmd/main.go

.PHONY: all clean windows linux macos

all: windows linux macos

# Build for Windows x86-64 (64-bit)
windows:
	GOOS=windows GOARCH=amd64 go build -o $(OUTPUT_BINARY).exe $(MAIN)
	@echo "编译成功，生成的 Windows 二进制文件为: $(OUTPUT_BINARY).exe"

# Build for Linux x86-64 (64-bit)
linux:
	GOOS=linux GOARCH=amd64 go build -o $(OUTPUT_BINARY)-linux $(MAIN)
	@echo "编译成功，生成的 Linux 二进制文件为: $(OUTPUT_BINARY)-linux"

# Build for macOS ARM64 (M1/M2)
macos:
	GOOS=darwin GOARCH=arm64 go build -o $(OUTPUT_BINARY)-macos $(MAIN)
	@echo "编译成功，生成的 macOS 二进制文件为: $(OUTPUT_BINARY)-macos"

clean:
	@rm -f $(OUTPUT_BINARY).exe $(OUTPUT_BINARY)-linux $(OUTPUT_BINARY)-macos
	@echo "清理完成"
