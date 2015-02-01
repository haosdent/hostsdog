package hostsdog

import (
    "os"
    "os/exec"
    "log"
)

func initDir() {
    err := os.Mkdir("/etc/hostsdog", 0755)
    if err != nil && !os.IsExist(err) {
        log.Fatal(err)
    }
}

func generateHosts(config string) {
}

func AddHosts(config string) {
    cmd := exec.Command("vi")
    cmd.Stdin = os.Stdin;
    cmd.Stdout = os.Stdout;
    cmd.Stderr = os.Stderr;
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

func RmHosts(config string) {
}

func ForkHosts(oldConfig string, newConfig string) {
}

func UseHosts(config string) {
}