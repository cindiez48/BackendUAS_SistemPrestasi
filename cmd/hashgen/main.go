package main

import (
	"flag"
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	pw := flag.String("password", "", "password to hash")
	flag.Parse()

	if *pw == "" {
		if flag.NArg() > 0 {
			*pw = flag.Arg(0)
		} else {
			fmt.Fprintln(os.Stderr, "Provide a password with -password or as first argument")
			os.Exit(2)
		}
	}

	hash, _ := bcrypt.GenerateFromPassword([]byte(*pw), bcrypt.DefaultCost)
	fmt.Println("HASH:", string(hash))
}
