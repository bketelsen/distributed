# Distributed Computing with Go

## Schedule

* Distributed Communication

* Distributed Architecture

* Contexts

* Service and Peer Discovery

* Security, Authentication, Authorization

* Consensus

* Monitoring, Metrics, Logging

* Distributed Tracing

* Offline and Nearline Processing

* Frameworks and Toolkits

* Deployment


## Development Environment

You have two choices for working with this repository:

* You can use vagrant/virtualbox
* You can have a Go 1.7 installation on your computer 

If you are using Vagrant/Virtualbox, you need to install VirtualBox first, then Vagrant.

If you have Go 1.7 installed, you'll need to set your GOPATH to the root of this repository.  You can do that manually, or use `direnv`. I strongly recommend using `direnv`!  See below for instructions.



## Vagrant Usage

The vagrant setup in this repository will share the `src` directory to /home/vagrant/src and set your GOPATH in the virtual machine to `/home/vagrant`.

It also adds `/home/vagrant/bin` to your $PATH so that any executables that you build or install will be in your path when you're working inside the VM.

###  Starting

	vagrant up

### Entering the VM
	
	vagrant ssh

### Shared Directories & Editing

The `src` directory of this repo is available inside the VM as `/home/vagrant/src`.  Changes you make on your host computer will be available inside the vm immediately.  This means that you can edit using your favorite text editor (Sublime Text, Atom, Visual Studio Code, {neo}Vim, Emacs, etc) but use the vagrant ssh session to compile and run.

### Testing Vagrant Setup

	vagrant ssh
	go build hello

This should compile the `hello/main.go` file at `/home/vagrant/src/hello/main.go`

	go install hello
	hello

This should compile and install the hello app, then run it.  


## Local Setup

Local setup requires that the GOPATH be set to the root of this repository:

	export GOPATH=`pwd`

Test this by building the `hello` package:

	go build hello
	go install hello

Building packages and binaries will put compiled output in the `bin` and `pkg` directories.  

Add `bin` to your PATH:

	export PATH=`pwd`:$PATH

Test this by running `hello`.

There is an `environment.sh` file in the root of the course directory that will setup your GOPATH for 
the current terminal session if you type:
	
	source ./environment.sh

You will need to do this each time you create a new shell.  Alternatively, you can use `direnv`, a useful utility
that will read and source the contents of an `.envrc` file each time you enter a directory.  

## DIRENV Setup

These training materials are expected to be a standalone GOPATH.  You can make your life easy by installing `direnv`.  `direnv` must be located somewhere in your $PATH to work.  My suggestion is to add the `bin` directory of your $GOPATH to your path if you haven't already.

For bash, modify ~/.bashrc so that $GOPATH/bin is in your binary search path:

	export PATH=$GOPATH/bin:$PATH

Now install `direnv`
	cd to $HOME 
	go get github.com/direnv/direnv

For direnv to work properly it needs to be hooked into the shell. Each shell
has its own extension mechanism:

### BASH

Add the following line at the end of the "~/.bashrc" file:

```sh
eval "$(direnv hook bash)"
```

Make sure it appears even after rvm, git-prompt and other shell extensions
that manipulate the prompt.

### ZSH

Add the following line at the end of the "~/.zshrc" file:

```sh
eval "$(direnv hook zsh)"
```

### FISH

Add the following line at the end of the "~/.config/fish/config.fish" file:

```fish
eval (direnv hook fish)
```

### TCSH

Add the following line at the end of the "~/.cshrc" file:

```sh
eval `direnv hook tcsh`
```

### Restart Shell

After making any of these modifications, close and reopen your shell session.

The first time you enter a directory with an `.envrc` file you'll be prompted to allow direnv to make changes to your environment.




