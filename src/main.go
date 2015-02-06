package main

import (
    "os"
    "hostsdog"
)

func main() {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog", "/tmp/hosts")
    if len(os.Args) <= 1 {
        usage()
        return
    }

    //dog.AddHosts("base")
    //dog.RmHosts("base")
    //dog.ForkHosts("base", "new")
    //dog.SwitchHosts("new")
    //dog.ListHosts()
}