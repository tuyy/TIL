## vim 설정
### vim plugin 설치
```
git clone https://github.com/VundleVim/Vundle.vim.git ~/.vim/bundle/Vundle.vim
```

### vimrc 설정
```
set nocompatible
filetype off

set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

" vim 파일 트리
Plugin 'scrooloose/nerdtree'
Plugin 'Xuyuanp/nerdtree-git-plugin'

Plugin 'majutsushi/tagbar'

Plugin 'paradigm/textobjectify'

Plugin 'vim-airline/vim-airline'
"Plugin 'airblade/vim-gitgutter'

" 파일검색
Plugin 'junegunn/fzf'
Plugin 'junegunn/fzf.vim'

" grep
Plugin 'yegappan/grep'
" 현재 커서 단어 검색 기본단축키 <leader>gw
Plugin 'aghareza/vim-gitgrep'

call vundle#end()
filetype plugin indent on

map <leader>` :cclose <CR>
map <leader>2 :FZF <CR>
map <leader>3 :tabe <CR>
map <leader>4 :make 
map <leader>q :q! <CR>
map <leader>aq :wqa! <CR>
map <F4> :NERDTreeToggle<CR>
nmap <F5> :TagbarToggle<CR>

nnoremap <silent> <S-F2> :tp<cr>            " Go to previous tag
nnoremap <silent> <S-F3> :tn<cr>            " Go to next tag
nnoremap <silent> <F2>   :cp<cr>            " Go to previous error line (compile result)
nnoremap <silent> <F3>   :cn<cr>            " Go to next error line (compile result)

noremap <F9> :Rgrep<CR>

nmap <F6> :set mouse=<CR>

if &term == 'xterm-256color'
    set <S-F1>=
    set <S-F2>=
    set <S-F3>=
    set <S-F4>=

endif

set t_ti= t_te=

set backspace=indent,eol,start

set autoindent

set nobackup
set history=50
set ruler
set showcmd
set laststatus=2

set titlestring=%F\ %m\ %r\ %*
set stl=%<%f\ %h%m%r%=0x%B\ %-10.((col:%c,line:%l)\ %L\ Lines%)

set fileformat=unix
set foldcolumn=0

set hlsearch
set ignorecase
set incsearch

set tabstop=4
set softtabstop=4
set shiftwidth=4
set shiftround
set expandtab

syntax on

" Keep buffer of 3 lines at each end of screen: top and bottom.
" set scrolloff=3

set encoding=utf-8
set termencoding=utf-8
set fileencoding=utf-8
set fileencodings=utf-8
set number

" Making vim load c file plugin under the directory of filetye plugin
filetype plugin indent on

" When editing a file, always jump to the last known cursor position.
" Don't do it when the position is invalid or when inside an event handler
" (happens when dropping a file on gvim).
autocmd BufReadPost *
\ if line("'\"") > 0 && line("'\"") <= line("$") |
\   exe "normal g`\"" |
\ endif

autocmd FileType Makefile                  setlocal noexpandtab
autocmd BufNewFile,BufRead *.i             setf c
autocmd BufNewFile,BufRead *.sc            setf c
autocmd BufNewFile,BufRead *.th,*.tc,*.ts  setf ruby
autocmd BufNewFile,BufRead *.re            setf lex
autocmd BufNewFile,BufRead *.ly            setf yacc
autocmd BufNewFile,BufRead *.idl           set ft=c

" Note that grep's options are set by env variable GREP_OPTIONS
let Grep_Path = '/bin/grep'
let Grep_OpenQuickfixWindow = 1
let Grep_Default_Options = '-rn'

" Suppress beeping and flashing screen
if $OS_NAME == "osx"
    set vb " macvim beeps unless this option is set
    set noeb
else
    set vb t_vb=
    set novb " gvim flashes screen without this option
    set noeb
endif

" Decorate doxygen comment.
let g:load_doxygen_syntax=1
let g:doxygen_enhanced_color=1

set path+=.,/usr/include/,/usr/local/include,./include
set path+=~/shawn/index-seal.trunk/box-model-generator/include,~/shawn/index-seal.trunk/include

set imactivatekey=S-space

" Because MacVim has its own color scheme,
" it does not display my own color scheme.
" This line keep MacVim from applying its own color scheme.
colorscheme default

if &diff
    hi DiffChange ctermbg=153 " Light Light blue
    hi DiffText   ctermbg=33 " Light blue
endif

hi Search cterm=NONE ctermfg=white ctermbg=blue

```
