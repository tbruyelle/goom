package goom

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	c, err := New()

	if err != nil {
		t.Fatal(err)
	}
	if len(c.Lists) == 0 {
		t.Error("No cmds found")
	}
}

func TestGet(t *testing.T) {
	c, _ := New()

	cmd, err := c.Get("screenshot")

	if err != nil {
		t.Fatalf("Get() returns an error %s", err)
	}
	fmt.Println(cmd)
}
