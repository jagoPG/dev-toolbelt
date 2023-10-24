# 🛠️ Dev Toolbelt 
Scripts to set up the computer and multipe purpose tasks. 

Commands are focused to use with a launcher as Raycast, which allows to rapidly execute a command
and get the result in the same toolbar.


## Table Of Contents

1. [Requirements](#1-requirements)
2. [Folder Structure](#2-folder-structure)
3. [How to use](#3-how-to-use)
4. [Author](#4-author)


## 1. Requirements

* macOS 13.*
* [Go](https://go.dev)
* [Raycast](https://www.raycast.com)


## 2. Folder Structure

* **bash**: handy bash scripts.
* **dotfiles**: scripts to set up a new laptop.
* **programs**: Go application with multiple purpose scripts.
* **raycast**: environment to set up Script Commands in Raycast.


## 3. How to Use
Explanation about how to use each section.

### Bash Scripts
Each script contains in the comments an expanation about what is its purpose and how to use it.

### Dotfiles
Read each script to customise which applications you want to install, or how to set up a specific
application. To execute all of them at once run the `run.sh` script.

### Programs
The program receive as argument the name of the module you want to execute. Each module expects
a different set of arguments.

You can compile the application by running the `build-programs.sh` script - which also creates the
binary in the raycast folder - or by manually running:
```bash
go test ./... && go build
```

### Raycast
The `bash` folder contained in this section has to be added by: 
> Raycast > Settings > Extensions > Add Script Directory

The `bin` folder is expected to contain the `programs` binary created by compiling the application.


## 4. List of Commands
After adding the `raycast/bash` folder to the Script Directory in Raycast, you can execute the
following commands:

* Count letters in the passed argument
* Convert from hexadecimal notation to RGB
* Convert from RGB notation to hexadecimal
* Encode to Base64
* Decode from Base64
* Generate UUID
* Generate MD5 Hash

## 5. Author
👨‍💻 Jagoba Pérez-Gómez

📧 email: jagobapg@pm.me
