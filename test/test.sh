#!/bin/sh

set_user() {
    sed -i "s/%%%/$USER/g" "manager.toml"
}

unset_user() {
    sed -i "s/$USER/%%%/g" "manager.toml"
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
}

without_target() {
    check_output "sween -o link -d without_target | grep 'ERROR:'" "Target is missed"
    sween -o unlink -d without_target >> /dev/null
}

without_source() {
    check_output "sween -o link -d without_source | grep 'ERROR:'" "Source is missed"
    sween -o unlink -d without_source >> /dev/null
}

full_path() {
    check_output "sween -o link -d full_path | grep 'ERROR:'" ""
    sween -o unlink -d full_path >> /dev/null
}

tilda_path() {
    check_output "sween -o link -d tilda_path | grep 'ERROR:'" ""
    sween -o unlink -d tilda_path >> /dev/null
}

only_hooks() {
    check_output "sween -o link -d only_hook" "this is HOOK!"
}

profile() {
    check_output "sween -o link -p main | grep 'ERROR:'" ""
    sween -o unlink -p main >> /dev/null
}

main() {
    set_user

    without_target
    without_source
    full_path
    tilda_path
    only_hooks
    profile

    unset_user
}

main
