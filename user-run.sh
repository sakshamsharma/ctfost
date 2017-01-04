#!/usr/bin/env bash

sudo -u $1 cgexec -g memory:ctflimit ~/${PROG_NAME:-service}
