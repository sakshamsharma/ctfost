#!/usr/bin/env bash

cgexec -g memory:ctflimit bash -c "sudo -u ctf-${1} /home/ctf-${1}/${PROG_NAME:-service}"
