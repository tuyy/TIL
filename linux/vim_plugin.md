### Vundle이란?
* vim plugin 관리자

### 설치
```
git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim
```

### plugin install
* ~/.vimrc 파일 제일 상단에 아래와 같이 놓는다.
```
set nocompatible              " be iMproved, required
filetype off                  " required

set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

Plugin 'scrooloose/nerdtree'
Plugin 'Xuyuanp/nerdtree-git-plugin'

call vundle#end()            " required
filetype plugin indent on    " required

```

* 예시
```
  1 set nocompatible
  2 filetype off
  3
  4 set rtp+=~/.vim/bundle/Vundle.vim
  5 call vundle#begin()
  6
  7 " vim 파일 트리
  8 Plugin 'scrooloose/nerdtree'
  9 Plugin 'Xuyuanp/nerdtree-git-plugin'
 10
 11 " 문법체크
 12 Plugin 'scrooloose/syntastic'
 13
 14 " 코드 자동완성
 15 Plugin 'valloric/youcompleteme'
 16
 17 call vundle#end()
 18 filetype plugin indent on
 19
 20 " Syntastic
 21 let g:syntastic_cpp_compiler = "g++"
 22 let g:syntastic_cpp_compiler_options = "-std=c++17 -Wall -Wextra -Wpedantic"
 ```

* https://vimawesome.com/ 에서 플러그인을 골라 vundle#begin() .. end() 사이에 놓는다.
* plugin을 설치한다
```
// vim 에서
:PluginInstall

// bash에서
vim +PluginInstall +qall
```
