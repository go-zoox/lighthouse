package hosts

import (
	"fmt"
	"testing"

	"github.com/go-zoox/fs"
)

func TestHosts(t *testing.T) {
	hosts := &Hosts{
		FilePath: fs.JoinPath(fs.CurrentDir(), "tests/hosts"),
		Mapping:  make(map[string]string),
	}

	if err := hosts.Load(); err != nil {
		t.Fatal(err)
	}

	for key, ip := range hosts.Mapping {
		fmt.Printf("%s => %s\n", key, ip)
	}

	if hosts.Length() != 8 {
		t.Fatal("invalid hosts length, expected 8, got", hosts.Length())
	}

	if ip, err := hosts.LookUp("localhost"); err != nil || ip != "127.0.0.1" {
		t.Fatal("expected localhost => 127.0.0.1, but got", ip, err)
	}

	if ip, err := hosts.LookUp("localhost", 4); err != nil || ip != "127.0.0.1" {
		t.Fatal("expected localhost:4 => 127.0.0.1, but got", ip, err)
	}

	if ip, err := hosts.LookUp("localhost", 6); err == nil {
		t.Fatal("expected localhost:6 not found, but got", ip, err)
	}

	if ip, err := hosts.LookUp("over"); err != nil || ip != "127.0.1.1" {
		t.Fatal("expected over => 127.0.1.1, but got", ip, err)
	}

	if ip, err := hosts.LookUp("over", 4); err != nil || ip != "127.0.1.1" {
		t.Fatal("expected over:4 => 127.0.1.1, but got", ip, err)
	}

	if ip, err := hosts.LookUp("over", 6); err == nil {
		t.Fatal("expected over:6 not found, but got", ip, err)
	}

	if ip, err := hosts.LookUp("ip6-localhost"); err == nil {
		t.Fatal("expected ip6-localhost not found, but got", ip, err)
	}

	if ip, err := hosts.LookUp("ip6-localhost", 4); err == nil {
		t.Fatal("expected ip6-localhost:4 not found, but got", ip, err)
	}

	if ip, err := hosts.LookUp("ip6-localhost", 6); err != nil || ip != "::1" {
		t.Fatal("expected ip6-localhost:6 => ::1, but got", ip, err)
	}
}
