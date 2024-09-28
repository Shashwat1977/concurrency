package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_main(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	main()
	result, _ := io.ReadAll(r)
	output := string(result)
	_ = w.Close()
	os.Stdout = oldOut
	if !strings.Contains(output, "31200") {
		t.Error("Values dont match")
	}
}
