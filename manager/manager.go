package manager

import (
	"fmt"
	"strings"
    "os"
	. "sween/dotfiles"
	. "sween/profiles"

	"github.com/BurntSushi/toml"
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

func (m Manager) DotfilesOperation(dotfiles string, operation OperationType) {
    if dotfiles == "" {
        return
    }

	err := dotfilesOperation(dotfiles, operation, m)

	if err != nil {
		fmt.Println(Sprintf("%s %s", Red("ERROR:"), err))
	}
}

func (m Manager) ProfilesOperation(profiles string, operation OperationType) {
    if profiles == "" {
        return
    }

	err := profilesOperation(profiles, operation, m)

	if err != nil {
		fmt.Println(Sprintf("%s %s", Red("ERROR:"), err))
	}
}

// Helpers
func dotfilesOperation(dotfilesRaw string, operation OperationType, m Manager) error {
	dotfiles := strings.Split(dotfilesRaw, " ")

	if strings.ToLower(dotfiles[0]) == "all" {
		for dotfileName, dotfile := range m.Dotfiles {
			fmt.Println(Sprintf(White("%sing %s..."), White(operation), Green(dotfileName)))

			return dotfile.DotfileOperation(m.User, operation)
		}
	} else {
		for _, dotfile := range dotfiles {
			fmt.Println(Sprintf(White("%sing %s..."), White(operation), Green(dotfile)))

			return m.Dotfiles[dotfile].DotfileOperation(m.User, operation)
		}
	}

	return nil
}

func profilesOperation(profilesRaw string, operation OperationType, m Manager) error {
	profiles := strings.Split(profilesRaw, " ")

	for _, profile := range profiles {
		err := m.Profiles[profile].ProfileOperation(m.User, operation, m.Dotfiles)

		if err != nil {
			return err
		}
	}

	return nil
}
