package main

import (
    "os"
    "fmt"
    "runtime"
    "hostsdog"
)

var (
    defaultConfigDir = "/etc/hostsdog"
    defaultHosts = "/etc/hosts"
)

func usage() {
    fmt.Printf("Hostsdog Usage:\n")
    fmt.Printf("\tadd     [hosts]\n")
    fmt.Printf("\tedit    [hosts]\n")
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
    if runtime.GOOS == "windows" {
        defaultConfigDir = "C:/Windows/System32/drivers/etc/hostsdog"
        defaultHosts = "C:/Windows/System32/drivers/etc/hosts"
    }

    dog := hostsdog.NewHostsdog(defaultConfigDir, defaultHosts)
    if len(os.Args) <= 1 {
        usage()
        return
    }

    dog.CheckPermission()

    switch act := os.Args[1]; act {
    case "add":
        checkArgs(3)
        config := os.Args[2]
        dog.AddHosts(config)
    case "edit":
        checkArgs(3)
        config := os.Args[2]
        dog.EditHosts(config)
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