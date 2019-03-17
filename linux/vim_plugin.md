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
set nocompatible
filetype off

set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

" vim 파일 트리
Plugin 'scrooloose/nerdtree'
Plugin 'Xuyuanp/nerdtree-git-plugin'

" airline
Plugin 'vim-airline/vim-airline'

" vim-cpp-enhanced-highlight
Plugin 'octol/vim-cpp-enhanced-highlight'

" tagbar
Plugin 'majutsushi/tagbar'

call vundle#end()
filetype plugin indent on
 ```

* https://vimawesome.com/ 에서 플러그인을 골라 vundle#begin() .. end() 사이에 놓는다.
* plugin을 설치한다
```
// vim 에서
:PluginInstall

// bash에서
vim +PluginInstall +qall
```
