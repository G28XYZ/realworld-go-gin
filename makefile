start:
	GO_ENV=prod go run main.go

dev:
	GO_ENV=dev air

dev_win:
	set GO_ENV=dev && air