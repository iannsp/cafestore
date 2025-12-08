#!/bin/bash

if [ ! -d "./bin" ]; then
    mkdir ./bin
fi

if [[ "$1" == "naiveCatalog" ]]; then
    echo "Building ./bin/naivecatalog-cli..."
    go build -o ./bin/naivecatalog-cli cmd/naivecatalog/main.go 
    echo "Done!"
    exit 0
fi

go build -o ./bin/service-naivecatalog cmd/naivecatalog/main.go 
go build -o ./bin/service-webui cmd/webui/main.go 
go build -o ./bin/service-whatsapp cmd/whatsapp/main.go 
