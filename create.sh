#!/bin/bash
if [[ -z $1 ]];
then 
    echo "No parameter passed."
    exit 1
fi
touch2() { mkdir -p "$(dirname "$1")" && touch "$1" ; }
pkg=$1
touch2 $pkg/$pkg.go;

echo -en "package $pkg\n\nfunc Run() {\n\n}\n" > ./$pkg/$pkg.go; pkg=""