package hostsdog

import (
    "fmt"
    "os"
    "os/exec"
    "log"
)

type Hostsdog struct {
    dir string
}

func NewHostsdog(dir string) *Hostsdog {
    var instance = &Hostsdog{
        dir,
    }
    instance.initDir()
    return instance
}

func (self *Hostsdog) initDir() {
    err := os.Mkdir(self.dir, 0755)
    if err != nil && !os.IsExist(err) {
        log.Fatal(err)
    }
}

func (self *Hostsdog) generateHosts(config string) {
    //TODO
}

func (self *Hostsdog) getConfigPath(config string) string {
    return fmt.Sprintf("%s/%s", self.dir, config)
}

func (self *Hostsdog) EditHosts(config string) {
    cmd := exec.Command("vim", self.getConfigPath(config))
    cmd.Stdin = os.Stdin;
    cmd.Stdout = os.Stdout;
    cmd.Stderr = os.Stderr;
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
}

func (self *Hostsdog) AddHosts(config string) {
    self.EditHosts(config)
}

func (self *Hostsdog) RmHosts(config string) {
    err := os.Remove(self.getConfigPath(config))
    if err != nil {
        log.Fatal(err)
    }
}

func (self *Hostsdog) ForkHosts(oldConfig string, newConfig string) {
    cmd := exec.Command("cp", "-rf", self.getConfigPath(oldConfig), self.getConfigPath(newConfig))
    err := cmd.Run()
    if err != nil {
        log.Fatal(err)
    }
    self.EditHosts(newConfig)
}

func (self *Hostsdog) UseHosts(config string) {
}

func (self *Hostsdog) ListHosts() {
}