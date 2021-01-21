package manager

import (
	"fmt"
    "errors"
    . "github.com/logrusorgru/aurora"
)

type Profile struct {
    Dotfiles []string
}
type Profiles map[string]Profile

func (p Profile) ProfileOperation(user string, operation OperationType, dotfiles Dotfiles) error {
	if len(p.Dotfiles) == 0 {
		return errors.New("profile does not exist")
	}

	for _, dotfile := range p.Dotfiles {
        fmt.Println(Sprintf(White("%sing %s..."), White(operation), Green(dotfile)))

        dotfiles[dotfile].DotfileOperation(user, operation)
	}

	return nil
}
