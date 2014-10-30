package comtool

import "testing"

func TestShRun(t *testing.T) {
	if err := ShRun("echo", "hello"); err != nil {
		t.Error(err)
	}
}
