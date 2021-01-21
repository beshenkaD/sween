package main

import (
	"fmt"
	"github.com/beshenkaD/sween/bootstrap"
	m "github.com/beshenkaD/sween/manager"
	"github.com/logrusorgru/aurora"
	"github.com/pborman/getopt/v2"
)

const VERSION = "0.1.5-dev"

var (
	isHelp       bool
	isVersion    bool
	initDirName  string
	OperationRaw string
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
		err := bootstrap.InitDotfilesDir(initDirName)

		if err != nil {
		}

		return
	}

	operation := m.NewOperationType(OperationRaw)
	manager := m.NewManager()
	// m.Convert("/home/beshenka/.config/cmus", manager.User)

	manager.Operation(dotfilesRaw, profilesRaw, operation)
}

func initArgs() {
	getopt.FlagLong(&isHelp, "help", '?', "Display help")
	getopt.FlagLong(&isVersion, "version", 'v', "Display version")
	getopt.FlagLong(&initDirName, "init", 'i', "Initialize new dotfile directory")
	getopt.FlagLong(&OperationRaw, "operation", 'o', "Operation performed on dotfiles [link] [unlink]")
	getopt.FlagLong(&profilesRaw, "profiles", 'p', "Use profiles")
	getopt.FlagLong(&dotfilesRaw, "dotfiles", 'd', "Use dotfiles")
}
