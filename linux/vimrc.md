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



**vimrc22**

```
set nocompatible

set shell=/bin/bash

set rtp+=~/.vim/bundle/Vundle.vim

call vundle#begin()

Plugin 'gmarik/Vundle.vim'
Plugin 'ervandew/supertab'
Plugin 'nanotech/jellybeans.vim'
Plugin 'majutsushi/tagbar'
Plugin 'scrooloose/nerdtree'
Plugin 'nathanaelkane/vim-indent-guides'
Plugin 'airblade/vim-gitgutter' " vim with git status(added, modified, and removed lines)
Plugin 'tpope/vim-fugitive' " vim with git command(e.g., Gdiff)
Plugin 'vim-airline/vim-airline' " vim status bar
Plugin 'vim-airline/vim-airline-themes'
Plugin 'blueyed/vim-diminactive'
Plugin 'danro/rename.vim'

Plugin 'jiangmiao/auto-pairs'
Plugin 'preservim/nerdcommenter' " 주석 생성 \ci
Plugin 'junegunn/fzf'
Plugin 'junegunn/fzf.vim'
Plugin 'tpope/vim-surround' " ysiw\" cs\"\' ds\" yss\"
Plugin 't9md/vim-choosewin'

"Plugin 'jnwhiteh/vim-golang'

"Plugin 'fatih/vim-go'
Plugin 'scrooloose/syntastic'

call vundle#end()

set t_Co=256

" for jellybeans
colorscheme jellybeans

" =================== 단축키 ===================

" 현재 파일 기준 NERDTree 열기
nmap ,t :NERDTreeFind<CR>

map <leader>` :cclose <CR>
map <leader>2 :FZF <CR>
map <leader>3 :tabe <CR>
map <leader>4 :make -j40<CR>
map <leader>5 :make dist; make check -j40<CR>
map <leader>6 :!pytest % -x -s -vvv
map <leader>q :q! <CR>
map <leader>aq :wqa <CR>

nmap  -  <Plug>(choosewin)

nnoremap <silent> <leader>lp  :lprevious<CR>     " Go to previous syntastic error
nnoremap <silent> <leader>ln  :lnext<CR>         " Go to next syntastic error
nnoremap <silent> <C-d> :SyntasticReset<CR>

nnoremap <silent> <F2>   :cp<CR>            " Go to previous error line (compile result)
nnoremap <silent> <F3>   :cn<CR>            " Go to next error line (compile result)
map <F4> :NERDTreeToggle<CR>
nmap <F5> :TagbarToggle<CR>
"nmap <F6> :SyntasticCheck gcc<CR>

noremap <F9> :Rgrep<CR>

nnoremap <leader>bp :bp<CR>
nnoremap <leader>bn :bn<CR>

"" TODO============================== linter
"set statusline+=%#warningmsg#
"set statusline+=%{SyntasticStatuslineFlag()}
"set statusline+=%*
"let g:syntastic_always_populate_loc_list = 1
"let g:syntastic_auto_loc_list = 1
"let g:syntastic_check_on_open = 1
"let g:syntastic_check_on_wq = 0
"let g:syntastic_auto_jump = 1
"
""let g:syntastic_cpp_checkers = ['cpplint']
""let g:syntastic_cpp_cpplint_exec = 'cpplint'
"let g:syntastic_cpp_compiler = 'g++'
"let g:syntastic_cpp_compiler_options = "-std=c++17 -Wall -Werror -fconcepts -Wno-error=unused-local-typedefs"
"let g:syntastic_c_compiler_options = "-std=c17 -Wall -Werror -fconcepts -Wno-error=unused-local-typedefs"
"let g:syntastic_cpp_check_header = 0
"let g:syntastic_cpp_no_include_search = 1
"let g:syntastic_cpp_include_dirs = ["include", "externals/owfs/include", "externals/carbon/include", "externals/boost/include", "externals/nmqueue/include", "model/include", "externals/synapfilter/include", "externals/json/include", "externals/googletest-master/include", "externals/libzip/include", "externals/mysql-connector-c++-8.0.12-linux-glibc2.12-x86-64bit/include/jdbc", "externals/icu/include", "externals/chardet/include"]
let g:syntastic_mode_map = {"mode": "passive"}

" ==============================

" Note that grep's options are set by env variable GREP_OPTIONS
let Grep_Path = '/bin/grep'
let Grep_OpenQuickfixWindow = 1
let Grep_Default_Options = '-rn'

" Making source code line foo(a, b, c,
"                         foo(a,
"                             b,
"                             c,
" meaning: , not followed by 1 or more traling white spaces
" :help \@!
nnoremap <S-F11> mr:s/,\(\s*$\)\@!/,\r/g<CR>=`r

" Removing trailing white spaces
nnoremap <S-F12> :%s/\s\+$//g<CR>

nnoremap <silent> // :noh<CR>

" ==============================================

set autoindent
set cindent
set smartindent " 스마트한 들여쓰기"

set showcmd "visual mode 로 선택된 라인수 표기됨
set tabstop=4
set softtabstop=4
set shiftwidth=4
set shiftround
set expandtab
set encoding=utf-8
set termencoding=utf-8
set fileencoding=utf-8
set fileencodings=utf-8
set number
set history=1000
set mouse=
set hlsearch
" for indent guide
let g:indentguides_spacechar = '┆'
let g:indentguides_tabchar = '|'
let g:indent_guides_enable_on_vim_startup = 1
let g:indent_guides_start_level=2
let g:indent_guides_guide_size=1

" for vim-airline
let g:airline#extensions#tabline#enabled = 1 " turn on buffer list
let g:airline_theme='hybrid'
set laststatus=2 " turn on bottom bar
let mapleader = ","

" for blueyed/vim-diminactive
let g:diminactive_enable_focus = 1

syntax enable
filetype indent on
highlight Comment term=bold cterm=bold ctermfg=4

" TODO 여유가되면 문법 학습 후 제대로 적용하자;
" TODO 여유가되면 문법 학습 후 제대로 적용하자;
" TODO 여유가되면 문법 학습 후 제대로 적용하자;
function! ReplaceTuyy(src)
  let src_text= a:src
  let dest_text = input("Enter dest word: ", src_text)
  let rep_type = input("Enter range: ", "block")

  if (rep_type == 'block')
      let space_num = str2nr(input("Enter block space size: ", "0"))
      let start_block   = "{"
      let end_block     = "}"
      let new_end_block = "};"
      if space_num == 4
        let start_block   = "    {"
        let end_block     = "    }"
        let new_end_block = "    };"
      elseif space_num == 8
        let start_block     = "        {"
        let end_block       = "        }"
        let new_end_block   = "        };"
      elseif space_num == 12
        let start_block     = "            {"
        let end_block       = "            }"
        let new_end_block   = "            };"
      elseif space_num == 16
        let start_block     = "                {"
        let end_block       = "                }"
        let new_end_block   = "                };"
      elseif space_num == 20
        let start_block     = "                    {"
        let end_block       = "                    }"
        let new_end_block   = "                    };"
      endif

      let current_line_no = line(".")
      let index = current_line_no
      let start = 1
      while index > 0
        if getline(index) == start_block
            let start = index
            break
        endif
        let index = index - 1
      endwhile
      
      
      let index = current_line_no
      let end = line(".")
      while index <= line('$')
        if getline(index) == end_block
          let end = index
          break
        endif

        if getline(index) == new_end_block
          let end = index
          break
        endif
        let index = index + 1
      endwhile
  else
      let start = 1
      let end = line("$")
  endif

  while start <= end
    call setline(start, substitute(getline(start), src_text, dest_text, "g"))
    let start = start + 1
  endwhile

  redraw

  if (rep_type != 'block')
    let rep_type = 'file'
  endif
  echom "[" . rep_type . "] " . src_text . " -> " . dest_text . " (Done!)"

endfunction
" block 단위로 단어바꾸기
map <S-F6> :call ReplaceTuyy(input("Enter src word: ", expand("<cword>")))<CR>

function! GetFuncLineSize(src)
  let space_num= str2nr(a:src)

  let start_block   = "{"
  let end_block     = "}"
  let new_end_block = "};"
  if space_num == 4
    let start_block   = "    {"
    let end_block     = "    }"
    let new_end_block = "    };"
  elseif space_num == 8
    let start_block     = "        {"
    let end_block       = "        }"
    let new_end_block   = "        };"
  elseif space_num == 12
    let start_block     = "            {"
    let end_block       = "            }"
    let new_end_block   = "            };"
  elseif space_num == 16
    let start_block     = "                {"
    let end_block       = "                }"
    let new_end_block   = "                };"
  elseif space_num == 20
    let start_block     = "                    {"
    let end_block       = "                    }"
    let new_end_block   = "                    };"
  endif

  let current_line_no = line(".")
  let index = current_line_no
  let start = 1
    while index > 0
    if getline(index) == start_block
        let start = index
        break
    endif
    let index = index - 1
  endwhile

  let index = current_line_no
  let end = line(".")

  while index <= line('$')
    if getline(index) == end_block
      let end = index
      break
    endif

    if getline(index) == new_end_block
      let end = index
      break
    endif
    let index = index + 1
  endwhile

  redraw
  echom "[FuncLineSize] start(" . start . ") end(" . end . ") result(". (end - start - 1) . ")"
endfunction

" block 단위로 Func line 수 확인
map <S-F7> :call GetFuncLineSize(input("Enter block space size: ", "0"))<CR>
```

