package main

import (
	_ "fmt"
)

const (
	SEARCH = 1 << iota
	INFO
	DOWNLOAD
	INSTALL
)

type Config struct {
	Args    []string
	options int
	DestDir string
}

var DefaultConfig Config = Config{}

func (c *Config) ReadConfig(path string) error { return nil }
func (c *Config) ParseCLI(args []string) error { return nil }

var usageMessage string = "Usage: aurer <pattern>|<package-name>... [<options>]"
var helpMessage string = "Here should be help message"

type Action interface {
	Name() string
	Execute() error
}

type ActionSearch struct {
	pattern string
	by      string
}

// func New(pattern, by string) *ActionSearch { return &ActionSearch{ pattern: pattern, by: by } }
func (a ActionSearch) Name() string    { return "Search" }
func (a ActionSearch) Execute() error  { return nil }
func (a ActionSearch) Pattern() string { return a.pattern }
func (a ActionSearch) ByField() string { return a.by }


type ActionGetInfo struct {
	pkg_list []string
}

// func New(pkg string) ActionGetInfo { return &ActionGetInfo{ pkg_name: pkg } }
func (a ActionGetInfo) Name() string      { return "GetInfo" }
func (a ActionGetInfo) Execute() error  { return nil }
func (a ActionGetInfo) Pattern() []string { return a.pkg_list }

type ActionUsage struct {
	message string
}

func (a ActionUsage) Name() string { return "Usage" }
func (a ActionUsage) Execute() error  { return nil }
func (a ActionUsage) Pattern() []string {
	return []string{a.message}
}

type ActionDownload struct {}

func (a ActionDownload) Name() string { return "Download" }
func (a ActionDownload) Execute() error  { return nil }

type ActionInstall struct {}

func (a ActionInstall) Name() string { return "Install" }
func (a ActionInstall) Execute() error  { return nil }


func NewConfig() *Config { return &Config{} }

func (c *Config) Init(args []string) error {
	c.Args = args
	return nil
}

func (c Config) GetAction() (Action, error) {
	var action Action
	if len(c.Args) == 1 {
		action = ActionUsage{message: usageMessage}
		return action, nil
	}
	action = ActionSearch{pattern: c.Args[1]}

	return action, nil
}
