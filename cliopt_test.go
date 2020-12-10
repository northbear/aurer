package main

import (
	"testing"
)

// func TestCliOpt(t *testing.T) {
// 	var options CmdOptValues = GetCmdOpt()
// 	if len(*options.SearchPattern) == 0 {
// 		t.Log("Empty search pattern")
// 	} else {
// 		t.Log("Search pattern: ", options.SearchPattern)
// 	}
// }

func TestCliOptionSearching(t *testing.T) {
	args := []string{"aurer", "-Ss", "testpattern"}
	cmdOpts, err := GetActionConfig(args)
	if err != nil {
		t.Error("unexpected parsing error")
	}
	
	if cmdOpts.Action != SEARCH {
		t.Errorf("wrong value for Action: %d", cmdOpts.Action)
	}
	if cmdOpts.Param != "testpattern" {
		t.Errorf("wrong value for Param: %s", cmdOpts.Param)
	}
}

func TestCliOptionSearchInfo(t *testing.T) {
	args := []string{"aurer", "-Si", "testpattern"}
	cmdOpts, err := GetActionConfig(args)
	if err != nil {
		t.Error("unexpected parsing error")
	}
	
	if cmdOpts.Action != SEARCH_INFO {
		t.Errorf("wrong value for Action: %d", cmdOpts.Action)
	}
	if cmdOpts.Param != "testpattern" {
		t.Errorf("wrong value for Param: %s", cmdOpts.Param)
	}
}
