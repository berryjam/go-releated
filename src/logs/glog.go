package main

import (
	"flag"
	"github.com/golang/glog"
)

func main() {
	flag.Parse()

	glog.Info("This is a Info log")
	glog.Warning("This is a Warning log")
	glog.Error("This is a Error log")

	glog.V(1).Infoln("level 1")
	glog.V(2).Info("level 2")

	glog.Flush()
}
