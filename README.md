Sween
=====
Convenient dotfile manager with toml configuration.

# Installation
```shell script
go get github.com/beshenkaD/sween
```
# Getting started
**1. Setup:**
``` shell script
git clone https://github.com/myOrg/dotfiles
cd myDotfiles && touch manager.toml
```
Or if you want to create a new dotfile directory
```shell script
sween --init dotfiles
```
**2. Configuration and usage:**
Take a look at the [example](example).

# TODO
- [x] More cli syntax and type detection Example: ```sween link dot/profile```

## License
All the code in this repository is released under the GPL License. Take a look
at the [LICENSE](LICENSE) for more info.
