package main

import (
	"flag"

	"github.com/golang/glog"

	"pixur.org/pixur/be/server"
	"pixur.org/pixur/be/server/config"
)

func main() {
	flag.Parse()
	defer glog.Flush()

	s := new(server.Server)

	glog.Fatal(s.StartAndWait(config.Conf))
}
