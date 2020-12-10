package main

import (
	"errors"
	"flag"
	"strings"
)

// function for processing CLI options
// Examples of CLI calls:
//     -s <string> search packages matched to template given in <string>
//     -d <name>   download AUR files of package with name <name>
//     -u          update AUR files for installed AUR packages

type ActionType int

const (
	SEARCH ActionType = iota + 1
	SEARCH_INFO
	DOWNLOAD
	UPDATE
	UPDATE_INFO
)

var (
	ErrorNoParams error = errors.New("Empty parameter list")	
)

type CmdOptValues struct {
	Action        ActionType
	Param string
	SearchPattern string
	PackageName   string
}



func Usage() string {
	return `Usage: aurer -[S|D|U][i] <param>

Print Usage...`
}

func GetActionConfig(a []string) (CmdOptValues, error) {
	cmdOpt := CmdOptValues{}
	argc := len(a)
	if argc < 2 {
		return cmdOpt, ErrorNoParams
	}
	
	switch {
	case strings.HasPrefix(a[1], "-S"):
		switch {
		case strings.Contains(a[1], "i"):
			cmdOpt.Action = SEARCH_INFO
			if argc >= 2 {
				cmdOpt.Param = a[2]
			} else {
				return cmdOpt, errors.New("name of package is not defined")
			}
		case strings.Contains(a[1], "s"):
			cmdOpt.Action = SEARCH
			if argc >= 2 {
				cmdOpt.Param = a[2]
			} else {
				return cmdOpt, errors.New("search pattern is not defined")
			}
		default:
			cmdOpt.Action = DOWNLOAD
			if argc >= 2 {
				cmdOpt.Param = a[2]
			} else {
				return cmdOpt, errors.New("name of package is not defined")
			}
		}
	case strings.HasPrefix(a[1], "-U"):
		switch {
		case strings.Contains(a[1], "i"):
			cmdOpt.Action = UPDATE_INFO
			cmdOpt.Param = a[2]
		default:
			cmdOpt.Action = UPDATE
		}
		cmdOpt.Action = UPDATE
	}
	return cmdOpt, nil
}

func GetCmdOpt() CmdOptValues {
	cmdOpt := CmdOptValues{}

	flag.StringVar(&cmdOpt.SearchPattern, "s", "", "search AUR packages matched to given pattern")
	flag.StringVar(&cmdOpt.PackageName, "d", "", "download AUR package with given name")
	flag.Parse()
	return cmdOpt
}
