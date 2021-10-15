package main

import (
	"flag"
	"fmt"
	"runtime"
)

func printUsage() {
	fmt.Printf(`Usage of gocode-generator:

gocode-generator -mod MODULE -dsn DSN -db DATABASE -OPTIONS

Examples:

gocode-generator -mod "awesomeProject" -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database"
gocode-generator -mod "awesomeProject" -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database" -o "/to/path" 
gocode-generator -mod "awesomeProject" -dsn "root:123@tcp(127.0.0.1:3306)/test" -db "Database" -au "bigboss" -o "/to/path" 

Supports options:
`)
	flag.PrintDefaults()
}

func printVersion() {
	fmt.Printf(`
github.com/billcoding/gocode-generator
%s
`, runtime.Version())
}
