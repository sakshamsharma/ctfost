#!/usr/bin/env bash

apt-get update
apt-get install cgroup-bin sudo

cp cgconfig.conf /etc/cgconfig.conf
