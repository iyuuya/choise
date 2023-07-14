package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"

	"github.com/iyuuya/go/exitcode"
)

func usage() {
	fmt.Println("usage: choise item1 item2 [itemN ...]")
}

func main() {
	os.Exit(int(realMain()))
}

func realMain() exitcode.ExitCode {
	rand.Seed(time.Now().UnixNano())
	argc := len(os.Args)
	if argc <= 1 {
		usage()
		return exitcode.ExitError
	}

	pick := rand.Intn(len(os.Args) - 1)
	fmt.Println(os.Args[pick+1])

	return exitcode.ExitOK
}
