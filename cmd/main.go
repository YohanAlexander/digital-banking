package main

import (
	"github.com/yohanalexander/desafio-banking-go/cmd/banking"
)

func init() {
	banking.InitENV()
}

func main() {
	db := banking.InitDB()
	banking.InitHTTP()
	defer db.Close()
}
