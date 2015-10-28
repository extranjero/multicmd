package cmd

import (
	"fmt"
	"testing"
)

var (
	cmd *Cmd
)

func TestSingle(t *testing.T) {
	cmd = Command("ls")

	b, err := cmd.Start()
	if err != nil {
		t.Error("start error:", err)
		return
	}

	fmt.Printf("%s \n ", b)

}

func TestMuliple(t *testing.T) {
	cmd = Command("ls")

	_, err := cmd.Pipe("wc", "-c")
	if err != nil {
		t.Error("pipe error: ", err)
		return
	}

	b, err := cmd.Start()
	if err != nil {
		t.Error("start error:", err)
		return
	}

	fmt.Printf("%s \n ", b)

}
