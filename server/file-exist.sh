#!/bin/bash

if [ -f "$1" ]; then
    echo "$1 exists."
else
    eval "$2"
fi

