#!/bin/bash
go run crashbot/crashbot.go
make -f Makefile.fuzz jq

