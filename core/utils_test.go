package core

import (
	"testing"
)

func TestGetWildCardHost(t *testing.T) {
	hosts, err := GetWildcardHost("xxx.zcorky.com")
	if err != nil {
		t.Fatal(err)
	}

	if len(hosts) != 1 {
		t.Errorf("expect 1, but got %d", len(hosts))
	}

	if hosts[0] != "*.zcorky.com" {
		t.Errorf("expect *.zcorky.com, but got %s", hosts[0])
	}

	hosts2, err := GetWildcardHost("xxx.inlets.zcorky.com")
	if err != nil {
		t.Fatal(err)
	}

	// fmt.Println(hosts2)

	if len(hosts2) != 2 {
		t.Errorf("expect 1, but got %d", len(hosts2))
	}

	if hosts2[0] != "*.inlets.zcorky.com" {
		t.Errorf("expect *.inlets.zcorky.com, but got %s", hosts2[0])
	}

	if hosts2[1] != "*.zcorky.com" {
		t.Errorf("expect *.zcorky.com, but got %s", hosts2[1])
	}
}
