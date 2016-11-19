BINARY=inspect

build:
	go build -o ${BINARY} 

test:
	go test -v ./...

install:
	go install

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean build
