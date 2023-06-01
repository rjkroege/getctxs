package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
)

const donotdisturbstate = "Library/DoNotDisturb/DB/ModeConfigurations.json"

func main() {
	flag.Parse()

	homedir, err := os.UserHomeDir()
	if err != nil {
		log.Fatalf("can't get a homedir: %v", err)
	}

	genAlfredResult(filepath.Join(homedir, donotdisturbstate), flag.Args())

}

const (
	prependflagstring = "alfredconnectionprepend"
	actionflagstring  = "alfredconnectionaction"
)

var actionflag = flag.Bool(actionflagstring, false, "run TaskPaper")
var prependflag = flag.Bool(prependflagstring, false, "prepend args to TaskPaper file")
