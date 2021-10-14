.PHONY: test
test:
	go test -v ./... -count=1

test-with-coverage:
	go test -v ./... -count=1 -covermode=atomic -coverprofile=profile.out