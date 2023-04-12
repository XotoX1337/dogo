![dogo logo](assets/Dogo.png)
[![Go Reference](https://pkg.go.dev/badge/github.com/XotoX1337/dogo.svg)](https://pkg.go.dev/github.com/XotoX1337/dogo)
[![Go Report Card](https://goreportcard.com/badge/github.com/XotoX1337/dogo)](https://goreportcard.com/report/github.com/XotoX1337/dogo)

docker (compose) command line helper with autocomplete written in [go](https://go.dev/) and [cobra](https://github.com/spf13/cobra)

## Install

```shell
go install github.com/XotoX1337/dogo@latest
```

## Usage 
```shell
Usage:
  dogo [command]

Available Commands:
  completion  Generate completion script
  create      Create all or a specific service from a docker-compose.yml file
  exec        Execute a command in a running container
  help        Help about any command
  list        List all containers & services
  rebuild     Rebuild one or many services
  remove      Remove one or many containers
  restart     Restart one or many containers
  shell       Use shell of a running container
  start       Start one or many containers
  stop        Stop one or many containers

Flags:
  -h, --help     help for dogo
  -t, --toggle   Help message for toggle
```
## Examples
    dogo shell yourContainer
    dogo start firstContainer secondContainer ...

# Completion
To generate a completion script run 
```shell
dogo completion [bash|zsh|fish|powershell]
```
This generates a completion script and prints it to stdout. After that, the script must be saved and loaded according to the environment (Windows, Linux).

Alternativley you can let `dogo` do all that for you. At the moment the following terminals are supported:
## Bash

```shell
dogo completion bash -f
```
this will generate the completion script and place it by default under `$HOME/.bash_completion.d/dogo-completion.sh` and also add a line to source the file for your profile.

You can change the default path for the completion script with 
```shell
dogo completion bash -f -d <path>
```

## Powershell

```shell
dogo completion bash -f
```
this will generate the completion script and place it by default under `$HOME\Documents\PowerShell\Microsoft.PowerShell_profile.ps1`.

If you already have a Powershell profile the completions cript will be added to the end of that file.

You can change the default path for the completion script with 
```shell
dogo completion bash -f -d <path>
```


## Credits
dogo mascot created by [@typomedia](https://github.com/typomedia)

