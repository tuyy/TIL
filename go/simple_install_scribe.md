## go install script
```
#!/bin/bash

set -e

GOVERSION=go1.12.6

echo INSTALL $GOVERSION

# go 다운로드 및 설치(/usr/local/go)
curl -O https://dl.google.com/go/$GOVERSION.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz

# gopath 지정
echo 'export GOPATH=$HOME/tuyy/go' >> $HOME/.bashrc
echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> $HOME/.bashrc
source /home1/irteam/.bashrc

echo GOPATH is $GOPATH

mkdir -p $HOME/tuyy/go/bin;mkdir -p $HOME/tuyy/go/src

# go version 확인
go version
```
