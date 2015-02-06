package main

import (
    "os"
    "fmt"
    "hostsdog"
)

func usage() {
    fmt.Printf("hello\n")
}

func checkArgs(requireSize int) {
    if len(os.Args) != requireSize {
        usage()
        os.Exit(-1)
    }
}

func main() {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog", "/tmp/hosts")
    if len(os.Args) <= 1 {
        usage()
        return
    }

    switch act := os.Args[1]; act {
    case "add":
        checkArgs(3)
        config := os.Args[2]
        dog.AddHosts(config)
    case "rm":
        checkArgs(3)
        config := os.Args[2]
        dog.RmHosts(config)
    case "fork":
        checkArgs(4)
        oldConfig := os.Args[2]
        newConfig := os.Args[3]
        dog.ForkHosts(oldConfig, newConfig)
    case "switch":
        checkArgs(3)
        config := os.Args[2]
        dog.SwitchHosts(config)
    case "list":
        dog.ListHosts()
    default:
        usage()
    }
}