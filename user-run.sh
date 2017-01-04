#!/usr/bin/env bash

sudo -u ctf-$1 cgexec -g memory:ctflimit ~/${PROG_NAME:-service}
