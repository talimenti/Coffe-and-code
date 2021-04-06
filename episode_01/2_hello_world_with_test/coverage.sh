#!/bin/sh
go test -coverprofile=cover.out
sed -i -e "s#.*/\(.*\.go\)#\./\\1#" cover.out
go tool cover -html=cover.out
