package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/BurntSushi/toml"
	. "github.com/logrusorgru/aurora"
)

type Manager struct {
	User     string
	Profiles map[string]Profile
	Dotfiles map[string]Dotfile
}

type Profile struct {
	Dotfiles []string
}

type Dotfile struct {
	Source string
	Target string
	Hooks  []string
}

func getCanonicalPath(path string, isTarget bool) string {
	usr, _ := user.Current()

	// This works fine with sudo and su -c
	// TODO:
	// Add doas support

	if usr.Name == "root" {
		sudoer := os.Getenv("SUDO_USER")
		if sudoer != "" {
			usr.HomeDir = filepath.Join("home", sudoer)
		}
	}

	dir := usr.HomeDir

	if isTarget {
		if path == "~" {
			path = dir
		} else if strings.HasPrefix(path, "~/") {
			path = filepath.Join(dir, path[2:])
		} else {
            path = " "
        }
	} else {
        wd, err := os.Getwd()

        if err != nil {
            panic(err)
        }

        path = filepath.Join(wd, path)
    }

	return path
}

func dotfileOperation(dotfile Dotfile, operation string) {
	source := getCanonicalPath(dotfile.Source, false)
	target := getCanonicalPath(dotfile.Target, true)

	if source == "" && target == "" && len(dotfile.Hooks) == 0 {
		fmt.Println(Sprintf(White("%s dotfile is not valid. Source, target and hooks are missed"), Red("ERROR:")))
		return
	} else if source == "" && target != "" {
		fmt.Println(Sprintf(White("%s dotfile is not valid. Target is missed"), Red("ERROR:")))
		return
	} else if target == "" && source != "" {
		fmt.Println(Sprintf(White("%s dotfile is not valid. Source is missed"), Red("ERROR:")))
		return
	}

	// Only hooks
	if source == "" && target == "" {
		runHooks(dotfile, operation)

		return
	}

	var err error

	if operation == "link" {
		err = os.Symlink(source, target)
	} else if operation == "unlink" {
		err = os.Remove(target)
	} else {
		err = fmt.Errorf("Unknown operation `%s`", operation)
		panic(err)
	}

	if err != nil {
		fmt.Println(Red("ERROR:"), err)
	} else {
		runHooks(dotfile, operation)
	}
}

func runHooks(dotfile Dotfile, operation string) {
	if operation == "unlink" {
		return
	}

	for _, hook := range dotfile.Hooks {
		out, err := exec.Command("sh", "-c", hook).Output()

		if err != nil {
			fmt.Println(err)
		}

		output := string(out[:])
		fmt.Print(output)
	}
}

func initDotfilesDir(name string) {
	err := os.MkdirAll(name, 0775)

	if err != nil {
		fmt.Println(err)
		return
	}

	file, err := os.OpenFile(filepath.Join(name, "manager.toml"), os.O_RDWR|os.O_CREATE, 0644)

	if err != nil {
		fmt.Println(err)
		os.RemoveAll(name)
		return
	}

	defer file.Close()

	fmt.Fprintf(file, "# This is example config. For more info see\n")
	fmt.Fprintf(file, "# https://github.com/beshenkaD/sween/tree/master/example\n\n")
	fmt.Fprintf(file, "[profiles]\n")
	fmt.Fprintf(file, "[profiles.main]\n")
	fmt.Fprintf(file, "\tdotfiles = [ \"vim\" ]\n")
	fmt.Fprintf(file, "[dotfiles]\n")
	fmt.Fprintf(file, "[dotfiles.vim]\n")
	fmt.Fprintf(file, "\tsource = \"vim\"\n")
	fmt.Fprintf(file, "\ttarget = \"~/.vimrc\"\n")
	fmt.Fprintf(file, "\thooks  = [ \"echo \"export EDITOR=vim\" > ~/.bashrc\" ]\n")

	fmt.Println(Green("Created"), "dotfile directory")
}

func main() {
	if runtime.GOOS == "windows" {
		fmt.Println("Go fuck yourself, fan of proprietary shit")
		os.Exit(1488)
	}

	version := "0.9.0 dev"

	// Parsing command line arguments
	var (
		isLink    bool
		isUnlink  bool
		isVersion bool
		Init      string
		profiles  string
		dotfiles  string
	)

	flag.StringVar(&profiles, "profiles", "", "use profile instead of dotfile names")
	flag.StringVar(&profiles, "p", "", "use profile instead of dotfile names (short)")

	flag.StringVar(&dotfiles, "dotfiles", "all", "dotfiles to install")
	flag.StringVar(&dotfiles, "d", "all", "dotfiles to install (short)")

	flag.BoolVar(&isLink, "link", false, "link dotfiles to destination point")
	flag.BoolVar(&isLink, "l", false, "link dotfiles to destination point (short)")

	flag.BoolVar(&isUnlink, "unlink", false, "unlink dotfiles")
	flag.BoolVar(&isUnlink, "u", false, "unlink dotfiles (short)")

	flag.BoolVar(&isVersion, "version", false, "show version")
	flag.BoolVar(&isVersion, "v", false, "show version (short)")

	flag.StringVar(&Init, "init", "", "init new dotfile directory")
	flag.StringVar(&Init, "i", "", "init new dotfile directory (short)")

	flag.Parse()

	if Init != "" {
		initDotfilesDir(Init)
		return
	}

	if isVersion {
		fmt.Println(Green("Sween"), version)
		return
	}

	// Parsing config file
	var manager Manager

	if _, err := toml.DecodeFile("manager.toml", &manager); err != nil {
		fmt.Println(Red("TOML ERROR:"), err)
		return
	}

	// Doing useful stuff here
	if !isLink && !isUnlink || isLink && isUnlink {
		fmt.Println(Red("ERROR:"), "You have not selected an operation")
		os.Exit(1)
	}

	var operation string
	if isLink {
		operation = "link"
	} else {
		operation = "unlink"
	}

	if profiles != "" {
		profiles := strings.Split(profiles, " ")

		for _, profile := range profiles {
			dots := manager.Profiles[profile]
			if len(dots.Dotfiles) == 0 {
				fmt.Println(Sprintf(White("%s Profile `%s` does not exist"), Red("ERROR:"), Red(profile)))
				os.Exit(1)
			}

			for _, dot := range dots.Dotfiles {
				fmt.Println(Sprintf(White("%sing %s..."), White(operation), Green(dot)))
				dotfileOperation(manager.Dotfiles[dot], operation)
			}
		}
		return
	}

	if dotfiles != "" {
		dots := strings.Split(dotfiles, " ")
		if dots[0] == "all" {
			for dotName, dot := range manager.Dotfiles {
				fmt.Println(Sprintf(White("%sing %s..."), White(operation), Green(dotName)))

				dotfileOperation(dot, operation)
			}
		} else {
			for _, dot := range dots {
				fmt.Println(Sprintf(White("%sing %s..."), White(operation), Green(dot)))

				dotfileOperation(manager.Dotfiles[dot], operation)
			}
		}

		return
	}
}
