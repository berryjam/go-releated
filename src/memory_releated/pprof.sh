#/bin/sh

go tool pprof --alloc_space readBytes bytes.mprof
