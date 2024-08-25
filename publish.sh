#!/usr/bin/env bash

if [ -z "$1" ]
  then
    echo "No argument supplied"
    exit 1
fi

git tag add $1
git push origin $1
