package main

import (
	"os"

	"github.com/kubeciio/koki/cmd"

	log "github.com/Sirupsen/logrus"
)

var GITCOMMIT = "HEAD"

func main() {
	// set version
	cmd.KokiCmd.Version = GITCOMMIT

	// run command
	if err := cmd.KokiCmd.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
