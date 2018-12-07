package main

import (
	"os"
	"io/ioutil"
)

func main() {
	// xxxxx为容器的ID
	pid := 1 // FIXME $pid应为容器xxxxx中第一个进程的ID
	os.MkdirAll("/sys/fs/cgroup/cpu/docker/xxxxx", 0755)
	ioutil.WriteFile("/sys/fs/cgroup/cpu/docker/xxxx/crgoup.procs", pid, os.O_CREATE)

	os.MkdirAll("/sys/fs/cgroup/cpuset/docker/xxxxx", 0755)
	ioutil.WriteFile("/sys/fs/cgroup/cpuset/docker/xxxxx/cgroup.procs", pid, os.O_CREATE)

	os.MkdirAll("/sys/fs/cgroup/blkio/docker/xxxxx", 0755)
	ioutil.WriteFile("/sys/fs/cgroup/blkio/docker/xxxxx/cgroup.procs", pid, os.O_CREATE)

	os.MkdirAll("/sys/fs/cgroup/memory/docker/xxxxx", 0755)
	ioutil.WriteFile("/sys/fs/cgroup/memory/docker/xxxxx/cgroup.procs", pid, os.O_CREATE)

	os.MkdirAll("/sys/fs/cgroup/devices/docker/xxxxx", 0755)
	ioutil.WriteFile("/sys/fs/cgroup/devices/docker/xxxxx/cgroup.procs", pid, os.O_CREATE)

	os.MkdirAll("/sys/fs/cgroup/freezer/docker/xxxxx", 0755)
	ioutil.WriteFile("/sys/fs/cgroup/freezer/docker/xxxxx/cgroup.procs", pid, os.O_CREATE)
}
