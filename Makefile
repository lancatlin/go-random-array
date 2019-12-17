GO := go
NAME := random-array

all: windows linux
windows: 
	GOOS=windows $(GO) build -o $(NAME)_windows.exe

linux:
	GOOS=linux $(GO) build -o $(NAME)_linux
