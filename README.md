# Congo

[![GoDoc](https://godoc.org/github.com/uraimo/congo?status.png)](https://godoc.org/github.com/uraimo/congo) [![GoReport A+](http://goreportcard.com/badge/uraimo/congo)](http://goreportcard.com/report/uraimo/congo)
[![Build Status](https://travis-ci.org/uraimo/congo.svg)](https://travis-ci.org/uraimo/congo)

Congo is a dead simple console Go package that mimic net/http server API to provide a simple way to add an interactive  console to your applications.

##Installation

`go get github.com/uraimo/congo`

##Usage

```go
package main

import (
	"fmt"
	"github.com/uraimo/congo"
)


func readit(args []string){
	fmt.Println("ls called with args:", args)
}

func err(cmd string){
	fmt.Println("Unknown command", cmd)
}

func main(){
  t :=congo.NewTerm("localhost>")
  t.AddCommand("ls",readit)
  t.AddUnknownHandler(err)
  t.ListenUntilExit()
} 
```

TODO
----
+ Configurable exit command/s
+ History support
+ Configurable autocomplete

