#!/bin/sh
set -xe

USER=gfaccinhuth
SSHKEY=~/.ssh/g5k
SITE=lyon


# copy files and connect
rsync -avz -e "ssh -i $SSHKEY" . $USER@access.grid5000.fr:$SITE/project
ssh -i $SSHKEY $USER@access.grid5000.fr
