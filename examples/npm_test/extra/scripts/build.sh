#!/usr/bin/env sh

# transpiles everything
npx swc src -d dist

# copies json files that were not handled by transpilation
cp src/code/services/sayHello/sayHello.json dist/code/services/sayHello/

# copies system.json into transpiled directory
cp system.json dist/
