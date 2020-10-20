# Sween

## Why?
because I can.
And also no one has ever written a dotfile manager in pure C99 before me.

## Installation
```shell script
make release
make install #You can specify prefix via PREFIX=
```
## Getting started
**1. Setup :**
``` shell script
git clone https://github.com/my_org/dotfiles
cd my_dotfiles && touch manager.toml
```

OR if you want to create a new dotfile directory
```shell script
sween --init dotfiles
```
**2. Configuration :**
Take a look at the [example](example).

## License
All the code in this repository is released under the GPL License. Take a look
at the [LICENSE](LICENSE) for more info.
