#!/usr/bin/env sh

npx swc src -d .
cp src/code/services/helloWorld/helloWorld.json code/services/helloWorld/

