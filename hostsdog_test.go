package hostsdog

import (
    _ "os"
    _ "log"
    "testing"
)

/*func TestMkdir(t *testing.T) {
    err := os.Mkdir("/etc/hostsdog", 0755)
    if err != nil && !os.IsExist(err) {
        log.Fatal(err)
    }
}*/

func TestAddHosts(t *testing.T) {
    AddHosts("")
}