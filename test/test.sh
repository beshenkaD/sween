#!/bin/sh

set_user() {
    sed -i "s/%%%/$USER/g" "manager.toml"
}

unset_user() {
    sed -i "s/$USER/%%%/g" "manager.toml"
}

restore() {
    sween -o unlink -d "all" >> /dev/null
}

# Colors
RED='\033[0;31m'
GRN='\033[0;32m'
RES='\033[0m'

check_output() {
    result=$($1)
    
    if echo "$result" | grep -q "$2"; then
        echo -e "TEST: ${FUNCNAME[1]} ${GRN}PASSED${RES}"
    else
        echo -e "TEST: ${FUNCNAME[1]} ${RED}FAILED${RES}"
    fi

    restore
}

without_target() {
    check_output "sween -o link -d without_target | grep 'ERROR:'" "Target is missed"
}

without_source() {
    check_output "sween -o link -d without_source | grep 'ERROR:'" "Source is missed"
}

full_path() {
    check_output "sween -o link -d full_path | grep 'ERROR:'" ""
}

tilda_path() {
    check_output "sween -o link -d tilda_path | grep 'ERROR:'" ""
}

only_hooks() {
    check_output "sween -o link -d only_hook" "this is HOOK!"
}

profile() {
    check_output "sween -o link -p main | grep 'ERROR:'" ""
}

multiple_dotfiles() {
    check_output "sween -o link -d 'tilda_path full_path' | wc -l" "2"
}

all_dotfiles() {
    check_output "sween -o link -d ALL | wc -l" "7"
}

bootstrap() {
    read -r -d '' text << EOM
# This is example config. For more info see
# https://github.com/beshenkaD/sween/tree/master/example
user = "beshenka"

[profiles]
[profiles.main]
    dotfiles = [ "vim" ]

[dotfiles]
[dotfiles.vim]
    source = "vim"
    target = "~/.vimrc"
    hooks  = [ "echo 'export EDITOR=vim' >> ~/.bashrc" ]

EOM

    sween -i testInitDir 
    check_output "cat testInitDir/manager.toml" $text

    rm -rf testInitDir
}

convert() {
    tmp="TMPFILEFORTEST"

    touch /tmp/$tmp >> /dev/null
    sween -c /tmp/$tmp >> /dev/null

    r=$(grep -Riwl 'manager.toml' -e 'dotfiles.TMPFILEFORTEST')
    check_output "echo $r" "manager.toml"

    rm -rf ./TMPFILEFORTEST

    sed -i '$ d' manager.toml
    sed -i '$ d' manager.toml
    sed -i '$ d' manager.toml
    sed -i '$ d' manager.toml
}

main() {
    set_user
    restore

    without_target
    without_source
    full_path
    tilda_path
    only_hooks
    profile
    multiple_dotfiles
    all_dotfiles
    bootstrap
    convert

    unset_user
}

main
