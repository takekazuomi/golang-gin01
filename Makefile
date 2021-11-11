help:	## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

run:	## Run
	go run main.go

test:	## Test
	go test takekazu.omi/golang-gin01 -v

cover:	## Coverage
	go test -cover takekazu.omi/golang-gin01 
