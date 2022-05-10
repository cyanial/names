package names

import (
	"testing"
)

var nmap map[string]struct{}

func setup() {
	nmap = make(map[string]struct{})
	for _, name := range names {
		nmap[name] = struct{}{}
	}
}

func TestGet(t *testing.T) {
	setup()

	name := Get()

	// fmt.Println(name)
	if _, ok := nmap[name]; !ok {
		t.Errorf("something wrong")
	}

}
