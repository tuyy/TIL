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

### ssh example
```
#!/bin/bash

set -e

if [ "x$1" == "x" ];
    then
        echo "Usage: ./install_go.sh <host>"
        exit
fi

GOVERSION=go1.12.6
GOPATH=/home1/irteam/tuyy/go

echo INSTALL $GOVERSION

# go 다운로드 및 설치(/usr/local/go)
ssh -l irteam $1 mkdir -p /home1/irteam/tmp
ssh -l irteam $1 "cd /home1/irteam/tmp;curl -O https://dl.google.com/go/$GOVERSION.linux-amd64.tar.gz"
ssh -l irteamsu $1 "cd /home1/irteam/tmp;sudo tar -C /usr/local -xzf go1.12.6.linux-amd64.tar.gz"

# gopath 지정
ssh -l irteam $1 "echo 'export GOPATH=$GOPATH' >> /home1/irteam/.bashrc"
ssh -l irteam $1 "echo 'export PATH=\$PATH:/usr/local/go/bin:$GOPATH/bin' >> /home1/irteam/.bashrc"
ssh -l irteam $1 source /home1/irteam/.bashrc

echo "GOPATH is $GOPATH"

ssh -l irteam $1 mkdir -p /home1/irteam/tuyy/go/bin
ssh -l irteam $1 mkdir -p /home1/irteam/tuyy/go/src

# go version 확인
ssh -l irteam $1 go version
```
