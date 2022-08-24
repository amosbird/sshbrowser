#!/usr/bin/env bash

env GOOS=windows go build -ldflags -H=windowsgui
