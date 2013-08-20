# Congo

Congo is a dead simple console Go package that mimic net/http server API to provide a simple way to add a console to your applications.

Installation
------------

	go get github.com/uraimo/congo

Usage
-----

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

