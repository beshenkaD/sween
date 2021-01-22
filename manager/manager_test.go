package manager

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
	"testing"
)

func changeUser() {
    // TODO: move it to utils
	f, err := ioutil.ReadFile("manager.toml")

	if err != nil {
		log.Fatal("Failed to open manager.toml")
	}

	newContents := strings.Replace(string(f), "%%%", os.Getenv("USER"), -1)

	if err = ioutil.WriteFile("manager.toml", []byte(newContents), 0); err != nil {
		log.Fatal("Failed to set user in manager.toml")
	}
}

func TestLink(t *testing.T) {
    m := NewManager()

    // TODO: hide output from this
    m.dotfilesOperation("all", Unlink)

    r1 := m.Dotfiles["without_target"].Operation(m.User, Link)
    r2 := m.Dotfiles["without_source"].Operation(m.User, Link)
    r3 := m.Dotfiles["full_path"].Operation(m.User, Link)
    r4 := m.Dotfiles["tilda_path"].Operation(m.User, Link)
    r5 := m.Dotfiles["only_hook"].Operation(m.User, Link)

    if r1 == nil {
        t.Errorf("incorrect result from dotfile without target")
    }
    if r2 == nil  {
        t.Error("incorrect result from dotfile without source")
    }
    if r3 != nil {
        t.Errorf("dotfile with full path failed: %e", r3)
    }
    if r4 != nil {
        t.Errorf("dotfile with tilda path failed: %e", r4)
    }
    if r5 != nil {
        t.Errorf("dotfile with only hook failed: %e", r5)
    }
}

func TestMain(m *testing.M) {
	os.Chdir("../test")
	changeUser()

	exitVal := m.Run()

	os.Exit(exitVal)
}
