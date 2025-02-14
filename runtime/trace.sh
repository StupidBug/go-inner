#!/bin/bash

while getopts 'f:' option; do 
    case "$option" in
        f) 
            func=$OPTARG
            ;;
    esac
done

go test -v ./gc -run $func -trace=trace.out
go tool trace trace.out