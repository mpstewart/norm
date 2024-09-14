package main

import (
	"github.com/mpstewart/norm"
)

type account struct {
	norm.Table `norm:"accounts"`
	Username   string `norm:"username"`
	Email      string `norm:"email"`
}

func main() {
	norm.Register[account]()

	norm.Search[account]{
		"username": "Greg",
		"email":    "greg@example.com",
	}.Query()
}
