package manager

import (
	"fmt"
	. "github.com/otiai10/copy"
	"os"
	"path/filepath"
	"strings"
)

// copy target to wd
// add target to manager.toml
// remove target
// link dotfile

func Convert(path string, user string) {
	// 1st
	wd, err := os.Getwd()
	usr, err := os.UserHomeDir()

	if err != nil {
		panic(err)
	}

	name := filepath.Base(path) // == source

	err = Copy(path, wd+"/"+name)

	if err != nil {
		panic(err)
	}

	// 2nd
	s := `
[dotfiles.%s]
    source = "%s"
    target = "%s"
`
	if ext := filepath.Ext(path); ext != "" {
		name = name[0 : len(name)-len(ext)]
	}
	if strings.HasSuffix(path, "rc") {
		name = strings.ReplaceAll(name, "rc", "")
	}

	ext := filepath.Ext(path)

	target := path

	if strings.HasPrefix(path, usr) {
		target = strings.ReplaceAll(target, usr, "~")
	}

	str := fmt.Sprintf(s, name, name, target)

	fmt.Println(ext)

	f, err := os.OpenFile("manager.toml", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		panic(err)
	}

	defer f.Close()

	if _, err = f.WriteString(str); err != nil {
		panic(err)
	}

	//3rd Remove target
	os.RemoveAll(path)

	//4th link
	newDotfile := Dotfile{
		Source: name,
		Target: target,
		Hooks:  nil,
	}

    err = newDotfile.DotfileOperation(user, Link)

    if err != nil {
        panic(err)
    }

	os.Exit(0)
}
