### vim 8.1 설치
* vim plugin "YouCompleteMe" 사용하기 위해 version 업그레이드 
* https://www.vim.org/download.php 확인
```
git clone https://github.com/vim/vim.git
cd vim/src
make
sudo make install
```


#### centos7
```
git clone https://github.com/vim/vim.git
cd vim
#make distclean  # if you build Vim before
make -j8
sudo make install
sudo cp src/vim /usr/bin
```
