package manager

import (
	"fmt"
	"os"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/beshenkaD/sween/utils"
	. "github.com/logrusorgru/aurora"
)

type Manager struct {
	User     string
	Profiles Profiles
	Dotfiles Dotfiles
}

func NewManager() Manager {
	var manager Manager

	if _, err := toml.DecodeFile("manager.toml", &manager); err != nil {
		fmt.Println(Red("TOML ERROR:"), err)
		os.Exit(1)
	}

	return manager
}

func (m *Manager) Operation(dotfiles string, profiles string, operation OperationType) {
	if dotfiles != "" {
		m.dotfilesOperation(dotfiles, operation)
	}

	if profiles != "" {
		m.profilesOperation(profiles, operation)
	}
}

func (m *Manager) dotfilesOperation(dotfilesRaw string, operation OperationType) {
	dotfiles := strings.Split(dotfilesRaw, " ")

	if strings.ToLower(dotfiles[0]) == "all" {
		for dotfileName, dotfile := range m.Dotfiles {
			fmt.Println(Sprintf(White("%sing %s..."), White(operation), Green(dotfileName)))

			err := dotfile.Operation(m.User, operation)
			utils.PrintError(err)
		}
	} else {
		for _, dotfile := range dotfiles {
			fmt.Println(Sprintf(White("%sing %s..."), White(operation), Green(dotfile)))

			err := m.Dotfiles[dotfile].Operation(m.User, operation)
			utils.PrintError(err)
		}
	}
}

func (m *Manager) profilesOperation(profilesRaw string, operation OperationType) {
	profiles := strings.Split(profilesRaw, " ")

	for _, profile := range profiles {
		err := m.Profiles[profile].ProfileOperation(m.User, operation, m.Dotfiles)
		utils.PrintError(err)
	}
}
