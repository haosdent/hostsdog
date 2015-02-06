package main

import (
    "os"
    "fmt"
    "hostsdog"
)

const (
    defaultConfigDir "/etc/hostsdog"
    defaultHosts "/etc/hosts"
)

func usage() {
    fmt.Printf("Hostsdog Usage:\n")
    fmt.Printf("\tadd     [hosts]\n")
    fmt.Printf("\trm      [hosts]\n")
    fmt.Printf("\tfork    [old hosts] [new hosts]\n")
    fmt.Printf("\tswitch  [hosts]\n")
    fmt.Printf("\tlist\n")
}

func checkArgs(requireSize int) {
    if len(os.Args) != requireSize {
        usage()
        os.Exit(-1)
    }
}

func main() {
    dog := hostsdog.NewHostsdog(defaultConfigDir, defaultHosts)
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