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

	config :=
`# This is example config. For more info see
# https://github.com/beshenkaD/sween/tree/master/example
user = "%s"

[profiles]
[profiles.main]
    dotfiles = [ "vim" ]

[dotfiles]
[dotfiles.vim]
    source = "vim"
    target = "~/.vimrc"
    hooks  = [ "echo 'export EDITOR=vim' >> ~/.bashrc" ]
`

	fmt.Fprintf(file, config, usr.Username)

	return nil
}
