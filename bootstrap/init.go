package bootstrap

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func InitDotfilesDir(name string) error {
	err := os.MkdirAll(name, 0755)

	if err != nil {
		return err
	}

	file, err := os.OpenFile(filepath.Join(name, "manager.toml"), os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		os.RemoveAll(name)
		return err
	}

	defer file.Close()

	usr, _ := user.Current()

	fmt.Fprintf(file, "# This is example config. For more info see\n")
	fmt.Fprintf(file, "# https://github.com/beshenkaD/sween/tree/master/example\n\n")
	fmt.Fprintf(file, "user = \"%s\"\n", usr.Username)
	fmt.Fprintf(file, "[profiles]\n")
	fmt.Fprintf(file, "[profiles.main]\n")
	fmt.Fprintf(file, "\tdotfiles = [ \"vim\" ]\n")
	fmt.Fprintf(file, "[dotfiles]\n")
	fmt.Fprintf(file, "[dotfiles.vim]\n")
	fmt.Fprintf(file, "\tsource = \"vim\"\n")
	fmt.Fprintf(file, "\ttarget = \"~/.vimrc\"\n")
	fmt.Fprintf(file, "\thooks  = [ \"echo 'export EDITOR=vim' >> ~/.bashrc\" ]\n")

	return nil
}
