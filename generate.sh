#!/usr/bin/env sh
#
# Generates everything that needs to be generated.
#

mockery --all --keeptree
go generate ./...

