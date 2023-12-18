#!/usr/bin/sh
set -xe

NUMBER_OF_NODES=2

oarsub -l nodes=$NUMBER_OF_NODES -I
