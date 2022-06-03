package main

import (
	"testing"
	//	"fmt"
)

var (
	argsWithoutParams []string = []string{"aurer"}
)

func compare_string_list(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i, p := range a {
		if p != b[i] {
			return false
		}
	}
	return true
}

func TestCompareStringListTrue(t *testing.T) {
	a := []string{"a", "b", "c"}
	if !compare_string_list(a, a) {
		t.Error("do not detect matching")
	}
}

func TestCompareStringListFalseByContent(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"c", "b", "a"}
	if compare_string_list(a, b) {
		t.Error("do not detect mismatching by content")
	}
}

func TestCompareStringListFalseByLength(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"a", "b"}
	if compare_string_list(a, b) {
		t.Error("do not detect mismatching by length")
	}
}

func TestAssingingArgs(t *testing.T) {
	var c *Config = NewConfig()
	err := c.Init(argsWithoutParams)
	if err != nil {
		t.Error("config initialization fails", err)
	}
	if !compare_string_list(c.Args, argsWithoutParams) {
		t.Error("Arg list is not assigned to an attribute Args", c, argsWithoutParams)
	}
}
