// Copyright 2019 acrazing <joking.young@gmail.com>. All rights reserved.
// Since 2019-12-27 13:27:58
package main

import (
	"github.com/acrazing/stmp-go/stmp"
	"os"
)

func main() {
	srv := stmp.NewServer()
	go srv.ServeTCP("127.0.0.1:5000")
	err := srv.Wait()
	if err != nil {
		println("listen error", err.Error())
		os.Exit(1)
	}
}
