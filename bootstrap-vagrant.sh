#!/bin/bash

echo Installing Basics
sudo DEBIAN_FRONTEND=noninteractive apt-get update
sudo DEBIAN_FRONTEND=noninteractive apt-get install mercurial bzr curl git -y

echo Setting Paths
export PATH=$PATH:/home/vagrant/bin:/usr/local/go/bin
export GOPATH=/home/vagrant

echo Installing Go
cd /tmp
curl -s -O https://storage.googleapis.com/golang/go1.7.linux-amd64.tar.gz
sudo tar -C /usr/local -xzf go1.8.3.linux-amd64.tar.gz
rm /tmp/go1.8.3.linux-amd64.tar.gz


echo Setting GOPATH and PATH
cat <<EOT >> /home/vagrant/.bashrc
export PATH=$PATH:/home/vagrant/bin:/usr/local/go/bin
export GOPATH=/home/vagrant
EOT

echo Installing vim-go
curl -fLo /home/vagrant/.vim/autoload/plug.vim --create-dirs https://raw.githubusercontent.com/junegunn/vim-plug/master/plug.vim
git clone https://github.com/fatih/vim-go.git /home/vagrant/.vim/plugged/vim-go

cat <<EOT >> /home/vagrant/.vimrc
call plug#begin()
Plug 'fatih/vim-go'
call plug#end()

set autowrite
let mapleader = ","

map <C-n> :cnext<CR>
map <C-m> :cprevious<CR>

nnoremap <leader>a :cclose<CR>
autocmd FileType go nmap <leader>r  <Plug>(go-run)
autocmd FileType go nmap <Leader>c <Plug>(go-coverage-toggle)
autocmd BufNewFile,BufRead *.go setlocal noexpandtab tabstop=4 shiftwidth=4 

let g:go_list_type = "quickfix"
let g:go_fmt_command = "goimports"
let g:go_highlight_types = 1
let g:go_highlight_fields = 1
let g:go_highlight_functions = 1
let g:go_highlight_methods = 1

" run :GoBuild or :GoTestCompile based on the go file
function! s:build_go_files()
  let l:file = expand('%')
  if l:file =~# '^\f\+_test\.go$'
    call go#cmd#Test(0, 1)
  elseif l:file =~# '^\f\+\.go$'
    call go#cmd#Build(0)
  endif
endfunction

autocmd FileType go nmap <leader>b :<C-u>call <SID>build_go_files()<CR>
EOT

chown -R vagrant:vagrant /home/vagrant/.vim
chown -R vagrant:vagrant /home/vagrant/.vimrc
