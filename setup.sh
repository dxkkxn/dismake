#!/usr/bin/sh
# set -xe
rm .bashrc

echo "Installing go..."
wget -O go1.21.4.linux-amd64.tar.gz https://go.dev/dl/go1.21.4.linux-amd64.tar.gz;
chmod -R u+w ~/go && rm -rf ~/go
tar -C ~ -xzf go1.21.4.linux-amd64.tar.gz 2>&1 > /dev/null
rm go1.21.4.linux-amd64.tar.gz;

export PATH=$PATH:~/go/bin;

echo "Installing goyacc..."
go install -C ~/dismake/client modernc.org/goyacc 2>&1 > /dev/null

echo "Installing blender..."
wget -O blender-4.0.2-linux-x64.tar.xz https://ftp.halifax.rwth-aachen.de/blender/release/Blender4.0/blender-4.0.2-linux-x64.tar.xz
tar -xvf blender-4.0.2-linux-x64.tar.xz  2>&1 > /dev/null
PATH=$PATH:~/blender-4.0.2-linux-x64
rm blender-4.0.2-linux-x64.tar.xz

echo "Installing magick..."
wget -O magick https://imagemagick.org/archive/binaries/magick
chmod u+x magick
mv magick ~/blender-4.0.2-linux-x64
echo $PATH > .path
echo "PATH=$PATH" > .bashrc
