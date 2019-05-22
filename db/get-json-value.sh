#!bin/sh

JSON=$1
VAL=$2

eval "cat $JSON | jq '$VAL'"