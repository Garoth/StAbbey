#!/bin/bash
# Rebuilds & runs the server

export GOPATH="$(pwd)"
export MYNAME="$0"

JS_DIR="resources/js/"
JS_BUILD_TARGETS=("hand.js" "map.js")

function stabbey_compile_css() {
    # Disabled for now, since we're using the client-side less js script
    return

    which lessc &> /dev/null
    if [[ "$?" -ne 0 ]]; then
        echo "Less compiler isn't installed! Skipping compilation of CSS files."
        echo "(You need to have a lessc binary in your PATH)"
    fi

    find . -name "*.less" | while read line; do
        local target="$(echo ${line} | sed 's/\.less/.css/')"
        lessc "${line}" "${target}"
    done
}

function stabbey_install() {
    go install $@ stabbey
}

function stabbey_run() {
    bin/stabbey $@
}

function stabbey_build_js_deps() {
    # Generating dependencies file
    echo "Writing Dependency File"
    python tools/closure/depswriter.py \
        --root_with_prefix="resources/js .." \
        --output_file="resources/js/3rd-party/deps.js"
}

function stabbey_build_js() {
    # local compilation_level="WHITESPACE_ONLY"
    local compilation_level="SIMPLE_OPTIMIZATIONS"
    # local compilation_level="ADVANCED_OPTIMIZATIONS"

    stabbey_build_js_deps

    for target in ${JS_BUILD_TARGETS[*]}; do
        # Running Google Closure Compiler
        echo "Building file ${target}"
        java -jar tools/closure/compiler.jar \
            --compilation_level "${compilation_level}" \
            --accept_const_keyword \
            --language_in "ECMASCRIPT5" \
            --summary_detail_level 3 \
            --externs $JS_DIR/externs.js \
            --warning_level=VERBOSE \
            --js_output_file "${JS_DIR}/compiled/${target}" \
            --js "${JS_DIR}/${target}" ${JS_DIR}/lib/*
    done
}

function stabbey_race() {
    stabbey_install -race
    echo "Build with race detector, running:"
    stabbey_run
}

function stabbey_deps() {
    go get code.google.com/p/go.net/websocket
    echo "Installed go.net websocket lib"
    go get github.com/oleiade/lane
    echo "Installed oleiade's lane (data structures) lib"
    echo "All deps installed"
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
    echo "   run:     builds and runs the stabbey server in JS compiled mode"
    echo "   runloop: like run, but restarts on successful exit. This speeds up"
    echo "            the development loop, since Control-C exits with 0 (and"
    echo "            triggers a rebuild), while Control-\\ exits with 1 (and"
    echo "            terminates the rebuild-run loop)."
    echo "            Also runs the JS in uncompiled debugging mode"
    echo "   js:      runs the javascript compiler only"
    echo "   jsdeps:  prepares the js dependencies file for closure"
    echo "   race:    builds and runs stabbey server with race detector"
    echo "   deps:    installs necessary dependencies"
    echo "   clean:   deletes currently built binaries and cache"
}

if [[ "$1" == "run" ]]; then
    stabbey_compile_css
    stabbey_install
    stabbey_build_js
    stabbey_run -compiledjs
elif [[ "$1" == "runloop" ]]; then
    while true; do
        stabbey_compile_css
        stabbey_install
        stabbey_build_js_deps
        stabbey_run
        if [[ "$?" -ne "0" ]]; then
            break
        fi
        sleep 1
    done
elif [[ "$1" == "js" ]]; then
    stabbey_build_js
elif [[ "$1" == "jsdeps" ]]; then
    stabbey_build_js_deps
elif [[ "$1" == "race" ]]; then
    stabbey_race
elif [[ "$1" == "deps" ]]; then
    stabbey_deps
elif [[ "$1" == "clean" ]]; then
    stabbey_clean
else
    stabbey_usage
fi
