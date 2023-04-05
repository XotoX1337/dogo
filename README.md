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

## Completion

To generate a completion file simply run 

```shell
dogo completion bash -f
```


this will store the completion script in `$HOME/.bash_completion.d/dogo-completion.sh`.

After that you only need to source this file in your profile

## Credits
dogo mascot created by [@typomedia](https://github.com/typomedia)

