package main

import (
	"flag"
)

func main() {
	flag.Parse()
	genAlfredResult("/Users/rjkroege/Library/DoNotDisturb/DB/ModeConfigurations.json", flag.Args())

}

const (
	prependflagstring = "alfredconnectionprepend"
	actionflagstring  = "alfredconnectionaction"
)

var actionflag = flag.Bool(actionflagstring, false, "run TaskPaper")
var prependflag = flag.Bool(prependflagstring, false, "prepend args to TaskPaper file")
