package manager

import (
	"fmt"
	"github.com/beshenkaD/sween/utils"
	. "github.com/otiai10/copy"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func createDotfile(path string, user string) Dotfile {
	var (
		source string
		target string
		hd     string // home dir
	)

	if h, err := os.UserHomeDir(); err != nil {
		utils.PrintError(err)
		os.Exit(1)
	} else {
		hd = h
	}

	source = filepath.Base(path)
	target = path

	if ext := filepath.Ext(path); ext != "" {
		source = source[0 : len(source)-len(ext)]
	}
	if strings.HasSuffix(path, "rc") {
		source = strings.ReplaceAll(source, "rc", "")
	}
	if strings.HasPrefix(path, hd) {
		target = strings.ReplaceAll(target, hd, "~")
	}

	newDotfile := Dotfile{
		Source: source,
		Target: target,
		Hooks:  nil,
	}

	return newDotfile
}

func copyToDest(source string, target string) {
	wd, err := os.Getwd()

	if err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}

	if err := Copy(source, wd+"/"+target); err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}
}

func writeToManager(template string, d Dotfile) {
	f, err := os.OpenFile("manager.toml", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}

	if _, err = f.WriteString(template); err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}

	f.Close()
}

func removeFromManager(strToRemove string) error {
	f, err := ioutil.ReadFile("manager.toml")

	if err != nil {
		return err
	}

	newContents := strings.Replace(string(f), strToRemove, "", -1)

	err = ioutil.WriteFile("manager.toml", []byte(newContents), 0)

	if err != nil {
		return err
	}

	return nil
}

func Convert(path string, user string) {
	dot := createDotfile(path, user)

	template := `
[dotfiles.%s]
    source = "%s"
    target = "%s"
`
	str := fmt.Sprintf(template, dot.Source, dot.Source, dot.Target)

	copyToDest(path, dot.Source)
	writeToManager(str, dot)

	if err := os.RemoveAll(path); err != nil {
		utils.PrintError(err)
		os.Exit(1)
	}

	if err := dot.Operation(user, Link); err != nil {
		if err := os.Rename(dot.Source, path); err != nil {
			panic(err)
		}

		if err := removeFromManager(str); err != nil {
			panic(err)
		}
		utils.PrintError(err)
		os.Exit(1)
	}
}
