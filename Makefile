all:
	cd cmd; go run main.go

test:
	@cd internal/file; go test -coverprofile cover.out; rm cover.out
	@cd internal/math; go test -coverprofile cover.out; rm cover.out