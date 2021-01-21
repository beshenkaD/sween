package main

import (
	"fmt"
	"github.com/beshenkaD/sween/bootstrap"
	m "github.com/beshenkaD/sween/manager"
	"github.com/beshenkaD/sween/utils"
	"github.com/logrusorgru/aurora"
	"github.com/pborman/getopt/v2"
)

const VERSION = "0.1.6-dev"

var (
	isHelp       bool
	isVersion    bool
	initDirName  string
	OperationRaw string
	profilesRaw  string
	dotfilesRaw  string
	convertRaw   string
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
			utils.PrintError(err)
		}

		return
	}

	operation := m.NewOperationType(OperationRaw)
	manager := m.NewManager()

	if convertRaw != "" {
		m.Convert(convertRaw, manager.User)
		return
	}

	manager.Operation(dotfilesRaw, profilesRaw, operation)
}

func initArgs() {
	getopt.FlagLong(&isHelp, "help", '?', "Display help")
	getopt.FlagLong(&isVersion, "version", 'v', "Display version")
	getopt.FlagLong(&initDirName, "init", 'i', "Initialize new dotfile directory")
	getopt.FlagLong(&OperationRaw, "operation", 'o', "Operation performed on dotfiles [link] [unlink]")
	getopt.FlagLong(&profilesRaw, "profiles", 'p', "Use profiles")
	getopt.FlagLong(&dotfilesRaw, "dotfiles", 'd', "Use dotfiles")
	getopt.FlagLong(&convertRaw, "convert", 'c', "Convert a file or dir to sween compatable dotfile")
}
