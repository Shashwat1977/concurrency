package main

import (
	"testing"
	"time"
)

func Test_dine(t *testing.T) {
	eatTime = 0 * time.Second
	thinkTime = 0 * time.Second

	for i := 0; i < 10; i++ {
		order = []string{}
		dine()
		if len(order) != 5 {
			t.Errorf("The length is not 5")
		}
	}
}
