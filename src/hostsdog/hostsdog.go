package hostsdog

import (
    "fmt"
    "os"
    "io"
    "bufio"
    "io/ioutil"
    "os/exec"
    "log"
)

type Hostsdog struct {
    dir string
    hosts string
}

const (
    baseConfig = "base"
)

func NewHostsdog(dir string, hosts string) *Hostsdog {
    var instance = &Hostsdog{
        dir,
        hosts,
    }
    instance.initDir()
    return instance
}

func (self *Hostsdog) initDir() {
    err := os.Mkdir(self.dir, 0755)
    if err == nil {
        self.copy(self.hosts, self.getConfigPath(baseConfig))
    } else if !os.IsExist(err) {
        log.Fatal(err)
    }
}

func checkErr(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func (self *Hostsdog) getConfigPath(config string) string {
    return fmt.Sprintf("%s/%s", self.dir, config)
}

func (self *Hostsdog) copy(src string, dst string) {
    data, err := ioutil.ReadFile(src)
    checkErr(err)
    err = ioutil.WriteFile(dst, data, 0644)
    checkErr(err)
}

func (self *Hostsdog) getHostsItems(config string) []string {
    items := make([]string, 0, 100)
    file, err := os.Open(self.getConfigPath(config))
    defer file.Close()
    if err == nil {
        reader := bufio.NewReader(file)
        for {
            line, _, err := reader.ReadLine()
            if err == io.EOF {
                break
            }
            items = append(items, string(line))
        }
    } else if !os.IsNotExist(err) {
        checkErr(err)
    }

    return items
}

func (self *Hostsdog) writeItems(writer io.Writer, items []string, config string) {
    if (len(items) > 0) {
        fmt.Fprintf(writer, "====== Start of %s ======\n", config)
    }
    for _, item := range items {
        fmt.Fprintf(writer, "%s\n", item)
    }
    if (len(items) > 0) {
        fmt.Fprintf(writer, "====== End of %s ======\n\n", config)
    }
}

func (self *Hostsdog) generateHosts(config string) {
    file, err := os.Create(self.hosts)
    defer file.Close()
    checkErr(err)
    writer := bufio.NewWriter(file)

    items := self.getHostsItems(baseConfig)
    self.writeItems(writer, items, baseConfig)

    items = self.getHostsItems(config)
    self.writeItems(writer, items, config)

    writer.Flush()
}

func (self *Hostsdog) EditHosts(config string) {
    cmd := exec.Command("vim", self.getConfigPath(config))
    cmd.Stdin = os.Stdin;
    cmd.Stdout = os.Stdout;
    cmd.Stderr = os.Stderr;
    err := cmd.Run()
    checkErr(err)
}

func (self *Hostsdog) AddHosts(config string) {
    self.EditHosts(config)
}

func (self *Hostsdog) RmHosts(config string) {
    err := os.Remove(self.getConfigPath(config))
    checkErr(err)
}

func (self *Hostsdog) ForkHosts(oldConfig string, newConfig string) {
    self.copy(self.getConfigPath(oldConfig), self.getConfigPath(newConfig))
    self.EditHosts(newConfig)
}

func (self *Hostsdog) SwitchHosts(config string) {
    self.generateHosts(config)
}

func (self *Hostsdog) ListHosts() {
    files, _ := ioutil.ReadDir(self.dir)
    if (len(files) > 0) {
        fmt.Printf("Config hosts:\n")
    }
    for _, file := range files {
        fmt.Printf("\t\t%s\n", file.Name())
    }
}