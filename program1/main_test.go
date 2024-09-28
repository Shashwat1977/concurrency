package main

import (
	"io"
	"os"
	"strings"
	"sync"
	"testing"
)

func Test_printSomething(t *testing.T) {
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	wg := sync.WaitGroup{}
	wg.Add(1)
	go printSomething("testing", &wg)
	wg.Wait()
	_ = w.Close()
	result, _ := io.ReadAll(r)
	os.Stdout = oldStdout
	if !strings.Contains(string(result), "testing") {
		t.Errorf("Values don't match %s", string(result))
	}
}
