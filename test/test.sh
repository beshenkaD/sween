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

    unset_user
}

main
