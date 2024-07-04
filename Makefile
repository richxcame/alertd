dev:
	@echo "Starting development server..."
	@go run main.go

buildwin:
	@echo "Building for windows..."
	@GOOS=windows GOARCH=amd64 go build -o alertd.exe
	@echo "Build complete"

buildlinux:
	@echo "Building for linux..."
	@GOOS=windows GOARCH=amd64 go build -o alertd
	@echo "Build complete"

buildmac:
	@echo "Building for mac with apple silicon chip..."
	@GOOS=darwin GOARCH=arm64 go build -o alertd
	@echo "Build complete"	