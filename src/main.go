package main

import (
    "hostsdog"
)

func main() {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog")
    //dog.AddHosts("base")
    //dog.RmHosts("base")
    //dog.ForkHosts("base", "new")
}