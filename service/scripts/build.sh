#!/bin/bash

if [ ! -d "./bin" ]; then
    mkdir ./bin
fi

declare -A apps

apps["naivecatalog"]="go build -o ./bin/naivecatalog-cli cmd/naivecatalog/main.go"
apps["raglite"]="go build -o ./bin/raglite cmd/raglite/main.go"

build() {
    target=$1
    echo "Building $target."
    eval "${apps["$target"]}"
    success=$?
    if [[ $success -eq 0 ]]; then
        echo "Target $target successfully build."
        exit 0
    fi
    exit 1
}

# check if build receive a parameter to define what to build.
# if not, build all.
# if build_target run ok
if [[ $# -eq 1 ]]; then
    target=$1
    if [[ -v apps["$target"] ]]; then
        build $target
    fi
    echo -e "Build Fail!\nTarget '$target' not found.\nSee the list:"
    for target_to_build in "${!apps[@]}"; do
        echo "- $target_to_build"
    done
    exit 1
fi
exit 0


