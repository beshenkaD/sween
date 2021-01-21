package main

import (
	"fmt"
	"github.com/logrusorgru/aurora"
	"github.com/pborman/getopt/v2"
	"sween/bootstrap"
	m "sween/manager"
)

const VERSION = "0.0.5 dev"

var (
	isHelp       bool
	isVersion    bool
	initDirName  string
	rawOperation string
	profilesRaw  string
	dotfilesRaw  string
)

func main() {
	initArgs()
	getopt.Parse()

	if isHelp {
		getopt.Usage()
		return
	} else if isVersion {
		fmt.Println(aurora.Green("Sween"), VERSION)
		return
	} else if initDirName != "" {
		bootstrap.InitDotfilesDir(initDirName)
		return
	}

    operation := m.NewOperationType(rawOperation)
	manager := m.NewManager()
    // m.Convert("/home/beshenka/.config/cmus", manager.User)

	manager.Operation(dotfilesRaw, profilesRaw, operation)
}

func initArgs() {
	getopt.FlagLong(&isHelp, "help", '?', "Display help")
	getopt.FlagLong(&isVersion, "version", 'v', "Display version")
	getopt.FlagLong(&initDirName, "init", 'i', "Initialize new dotfile directory")
	getopt.FlagLong(&rawOperation, "operation", 'o', "Operation performed on dotfiles [link] [unlink]")
	getopt.FlagLong(&profilesRaw, "profiles", 'p', "Use profiles")
	getopt.FlagLong(&dotfilesRaw, "dotfiles", 'd', "Use dotfiles")
}
