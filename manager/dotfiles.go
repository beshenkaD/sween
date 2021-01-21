package manager

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

type Dotfile struct {
	Source string
	Target string
	Hooks  []string
}

type Dotfiles map[string]Dotfile

func (d Dotfile) Operation(user string, operation OperationType) error {
	source := resolvePath(d.Source, user, false)
	target := resolvePath(d.Target, user, true)

	// fmt.Println(source)
	// fmt.Println(target)
	// fmt.Println(d.Hooks)

	if source == "" && target == "" && len(d.Hooks) == 0 {
		return errors.New("dotfile is not valid. Source, target and hooks are missed")
	} else if source != "" && target == "" {
		return errors.New("dotfile is not valid. Target is missed")
	} else if target != "" && source == "" {
		return errors.New("dotfile is not valid. Source is missed")
	} else if source == "" && target == "" && operation != Unlink {
		RunHooks(d.Hooks)
		return nil
	} else if source == "" && target == "" && operation == Unlink {
		return nil
	}

	if operation != Unlink {
		defer RunHooks(d.Hooks)
	}

	switch operation {
	case Link:
		return os.Symlink(source, target)
	case Unlink:
		return os.Remove(target)
	default:
		return fmt.Errorf("unknown operation `%s`\n", operation)
	}
}

func RunHooks(hooks []string) {
	for _, hook := range hooks {
		out, err := exec.Command("sh", "-c", hook).Output()

		if err != nil {
			fmt.Println(err)
		}

		output := string(out[:])
		fmt.Print(output)
	}
}

func resolvePath(path string, user string, isTarget bool) string {
	if path == "" {
		return ""
	}

	if isTarget {
		if path == "~" {
			return filepath.Join("/home", user)
		} else if strings.HasPrefix(path, "~/") {
			return filepath.Join("/home", user, path[2:])
		}
	} else {
		wd, err := os.Getwd()

		if err != nil {
			panic(err)
		}

		return filepath.Join(wd, path)
	}

	if strings.HasPrefix(path, "/") {
		return path
	}

	return ""
}
