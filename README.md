# `dogo` 
docker command line helper with autocomplete written in [Go](https://go.dev/) and [cobra](https://github.com/spf13/cobra)

## Usage 
```shell
Usage:
  dogo [command]

Available Commands:
  completion  Generate completion script
  exec        execute a command in a running container
  help        Help about any command
  list        list all containers & services
  remove      remove one or many containers
  ssh         connect to a running container
  start       start one or many containers
  stop        stop one or many containers
```
## Examples
    dogo ssh yourContainer
    dogo start firstContainer secondContainer ...

## Completion

To generate a completion file simply run 

```shell
dogo completion bash -f
```


this will store the completion script in `$HOME/.bash_completion.d/dogo-completion.sh`.

After that you only need to source this file in your profile

