![dogo logo](assets/Dogo.png)

[![Go Reference](https://pkg.go.dev/badge/github.com/XotoX1337/dogo.svg)](https://pkg.go.dev/github.com/XotoX1337/dogo)
[![Go Report Card](https://goreportcard.com/badge/github.com/XotoX1337/dogo)](https://goreportcard.com/report/github.com/XotoX1337/dogo)

A docker & docker-compose command line helper with shell autocompletion, written in
[go](https://go.dev/) and [cobra](https://github.com/spf13/cobra).

`dogo` wraps the most common docker and docker-compose tasks behind short, memorable
commands. You can address containers and services by name, by prefix, or by wildcard,
and act on many of them at once.

## Install

```sh
go install github.com/XotoX1337/dogo@latest
```

Alternatively, download a prebuilt binary from the latest
[release](https://github.com/XotoX1337/dogo/releases).

## Quick start

```sh
dogo list                  # show all containers & services
dogo start web db          # start one or many containers
dogo shell web             # open a shell in a running container
dogo logs web -n 50        # show the last 50 lines, then follow
```

## Commands

| Command      | Description                                                        |
| ------------ | ----------------------------------------------------------------- |
| `create`     | Create all or a specific service from a `docker-compose.yml` file  |
| `list`       | List containers & services (subcommands: `all`, `container`, `services`) |
| `start`      | Start one or many containers                                       |
| `stop`       | Stop one or many containers                                        |
| `restart`    | Restart one or many containers (stop, then start)                  |
| `remove`     | Remove one or many containers                                      |
| `rebuild`    | Rebuild one or many services (build, recreate and start)           |
| `logs`       | Show (and follow) the logs of one or many containers              |
| `exec`       | Execute a command in a running container                           |
| `shell`      | Open a shell (`/bin/bash`) in a running container                  |
| `completion` | Generate a shell completion script                                 |
| `version`    | Show the dogo version information                                  |
| `help`       | Help about any command                                            |

Run `dogo help <command>` for the full flag list of any command.

### Targeting containers & services

Most commands accept one or more names. Matching works as follows:

- **Exact match** wins if a container/service with that exact name exists.
- **Prefix match** otherwise — `dogo start infra_` starts every container whose name
  begins with `infra_`.
- **Wildcard** — a trailing `*` is also accepted as an explicit prefix wildcard
  (`dogo logs web*`). A bare `*` is rejected; you need at least one character.

### Examples

```sh
dogo shell myContainer
dogo start firstContainer secondContainer
dogo stop infra_                      # prefix: stops every infra_* container
dogo exec myContainer ls -la /app     # runs: docker exec -it myContainer ls -la /app
dogo logs myContainer                 # follow the logs (tail -f)
dogo logs myContainer -f=false        # print the logs once and exit
dogo logs myContainer -n 50           # last 50 lines, then follow
dogo logs myContainer -t              # include timestamps
dogo create                           # create services from ./docker-compose.yml
dogo create web -f ./infra            # create the "web" service from ./infra/docker-compose.yml
dogo rebuild web                      # rebuild, recreate and start the "web" service
```

### Command flags

**`logs`**

| Flag                 | Default | Description                                  |
| -------------------- | ------- | -------------------------------------------- |
| `-f, --follow`       | `true`  | Follow log output (`tail -f`)                |
| `-n, --tail`         | `all`   | Number of lines to show from the end         |
| `-t, --timestamps`   | `false` | Show timestamps                              |

**`create`**

| Flag         | Description                                                             |
| ------------ | ---------------------------------------------------------------------- |
| `-f, --file` | Path to the `docker-compose.yml` (file or directory; defaults to CWD)  |

**`rebuild`**

| Flag          | Description        |
| ------------- | ------------------ |
| `-p, --prune` | Prune build cache  |

## Completion

Generate a completion script and print it to stdout:

```sh
dogo completion [bash|zsh|fish|powershell]
```

You then save and source it according to your environment.

Alternatively, let `dogo` install it for you with `-f`/`--file`. It writes the script
to a default location and adds a line to your profile to source it.

### Bash / zsh / fish

```sh
dogo completion bash -f
```

Writes the script to `$HOME/.bash_completion.d/dogo-completion.sh` and sources it from
your `.bashrc`.

### Powershell

```powershell
dogo completion powershell -f
```

Writes the script to `$HOME\Documents\WindowsPowerShell\dogo-completion.ps1` and sources
it from your PowerShell profile.

### Custom destination

Override the default output directory with `-d`/`--dest`:

```sh
dogo completion bash -f -d <path>
```

## Credits

dogo mascot created by [@typomedia](https://github.com/typomedia).
