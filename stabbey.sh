#!/bin/bash
# Rebuilds & runs the server

# TODO Figure out a good way to install dependencies with go get. Deps:
#      - code.google.com/p/go.net/websocket

export GOPATH="$(pwd)"
export MYNAME="$0"

function stabbey_install() {
    pushd src/stabbey &> /dev/null
    go install
    popd &> /dev/null
}

function stabbey_run() {
    bin/stabbey
}

function stabbey_clean() {
    rm -Rf bin pkg
    pushd src/stabbey &> /dev/null
    go clean
    popd &> /dev/null
}

function stabbey_usage() {
    echo "Usage: $MYNAME <command>"
    echo ""
    echo "Where <command> is one of:"
    echo " - run:     builds and runs the stabbey server"
    echo " - runloop: like run, but restarts on successful exit. This speeds up"
    echo "            the development loop, since Control-C exits with 0 (and"
    echo "            triggers a rebuild), while Control-\\ exits with 1 (and"
    echo "            terminates the rebuild-run loop)"
    echo " - clean:   deletes currently built binaries and cache"
}

if [[ "$1" == "run" ]]; then
    stabbey_install
    stabbey_run
elif [[ "$1" == "runloop" ]]; then
    while true; do
        stabbey_install
        stabbey_run
        if [[ "$?" -ne "0" ]]; then
            break
        fi
        sleep 1
    done
elif [[ "$1" == "clean" ]]; then
    stabbey_clean
else
    stabbey_usage
fi
