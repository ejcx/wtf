#!/bin/bash
./crashbot -project=$1 -bucket=whatthefuzz &
make -f Makefile.fuzz $1
