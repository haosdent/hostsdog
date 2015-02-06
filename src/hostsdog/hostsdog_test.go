package hostsdog

import (
    _ "os"
    _ "log"
    _ "testing"
)

func TestAddHosts(t *testing.T) {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog", "/tmp/hosts")
    dog.AddHosts("new")
}

func TestEditHosts(t *testing.T) {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog", "/tmp/hosts")
    dog.EditHosts("new")
}

func TestRmHosts(t *testing.T) {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog", "/tmp/hosts")
    dog.RmHosts("new")
}

func TestForkHosts(t *testing.T) {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog", "/tmp/hosts")
    dog.ForkHosts("new", "newer")
}

func TestSwitchHosts(t *testing.T) {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog", "/tmp/hosts")
    dog.SwitchHosts("newer")
}

func TestListHosts(t *testing.T) {
    dog := hostsdog.NewHostsdog("/tmp/hostsdog", "/tmp/hosts")
    dog.ListHosts()
}