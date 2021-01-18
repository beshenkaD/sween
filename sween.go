package main

import (
	"fmt"
	. "github.com/logrusorgru/aurora"
	"github.com/pborman/getopt/v2"
	. "sween/dotfiles"
	. "sween/init"
	. "sween/manager"
)

const VERSION = "0.0.4 dev"

var (
	isHelp       bool
	isVersion    bool
	initDirName  string
	rawOperation string
	profiles     string
	dotfiles     string
)

func main() {
	initArgs()
	getopt.Parse()

	if isHelp {
		getopt.Usage()
		return
	} else if isVersion {
		fmt.Println(Green("Sween"), VERSION)
		return
	} else if initDirName != "" {
		InitDotfilesDir(initDirName)
		return
	}

	manager := NewManager()
	operation := NewOperationType(rawOperation)

	manager.DotfilesOperation(dotfiles, operation)
	manager.ProfilesOperation(profiles, operation)
}

func initArgs() {
	getopt.FlagLong(&isHelp, "help", '?', "Display help")
	getopt.FlagLong(&isVersion, "version", 'v', "Display version")
	getopt.FlagLong(&initDirName, "init", 'i', "Initialize new dotfile directory")
	getopt.FlagLong(&rawOperation, "operation", 'o', "Operation performed on dotfiles [link] [unlink]")
	getopt.FlagLong(&profiles, "profiles", 'p', "Use profiles")
	getopt.FlagLong(&dotfiles, "dotfiles", 'd', "Use dotfiles")
}
