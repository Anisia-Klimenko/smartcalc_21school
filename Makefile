all:
	cd cmd; go run main.go

test:
	@cd internal; cd file; go test -coverprofile cover.out; rm cover.out
	@cd internal; cd file; go test -coverprofile cover.out; rm cover.out