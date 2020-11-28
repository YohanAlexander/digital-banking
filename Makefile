BINARY=banking

build:
	go build -o ${BINARY} github.com/yohanalexander/desafio-banking-go/cmd/banking

clean:
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean