package kss

import (
	"fmt"
	"testing"
)

func TestNget(t *testing.T) {
	msg := A{
		"user":    A{"name": "Wauo"},
		"address": "china",
	}
	v := Nget(msg, []string{"user", "name"}, "~_~")
	fmt.Println(v)

	v2 := Nget(msg, []string{"user", "job"}, "~_~")
	fmt.Println(v2)

	v3 := Nget(msg, []string{"address"}, "~_~")
	fmt.Println(v3)
}
