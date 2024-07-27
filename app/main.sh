#!/bin/bash

usage() {
    cat << EOU
lddc: missing file arguments
Try \`lddc <path to the binary> <path to copy the dependencies>\`
EOU
    exit 0
}

if [ "$#" -lt 2 ]; then
    usage
fi

if [ ! -e "$1" ]; then
    echo "Not a valid input $1"
    exit 1
fi

if [ ! -d "$2" ]; then
    echo "No such directory $2, creating..."
    mkdir -p "$2"
fi

echo "Collecting the shared library dependencies for $1..."
deps=$(ldd "$1" | awk 'BEGIN{ORS=" "}$1~/^\//{print $1}$3~/^\//{print $3}' | sed 's/,$/\n/')
echo "Copying the dependencies to $2"
for dep in $deps; do
    echo "Copying $dep to $2"
    cp "$dep" "$2"
done

echo "Done!"
