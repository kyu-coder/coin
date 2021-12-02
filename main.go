package main

import (
	"github.com/kyu-coder/coin/cli"
	"github.com/kyu-coder/coin/db"
)

func main() {
	defer db.Close()
	cli.Start()
}
