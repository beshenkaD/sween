package manager

import (
	"github.com/otiai10/copy"
	"os"
)

/* Скопировать папку из source в target (тут они меняются местами)
   Создать запись в manager.toml заменив /home/$USER на ~/
   удалить папку source и выполнить link
*/

func Convert(source string, target string) {
	w, _ := os.Getwd()
    h, _ := os.UserHomeDir()
}
