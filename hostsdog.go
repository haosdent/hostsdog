pacakge hostsdog

import (
    "os"
    "log"
)

func initDir() {
    err = os.Mkdir("/etc/hostsdog", 0755)
    if err != nil {
        log.Fatal(err)
    }
}

func generateHosts(config string) {
}

func AddHosts(config string) {
}

func RmHosts(config string) {
}

func ForkHosts(oldConfig string, newConfig string) {
}

func UseHosts(config string) {
}