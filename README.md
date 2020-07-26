# todocli <!-- omit in toc -->

ToDoCli is a simple cli saving your time by managing your tasks from your terminal.

![goreleaser](https://github.com/wstnbrg/todocli/workflows/goreleaser/badge.svg?event=release)

- [Install](#install)
  - [Using homebrew](#using-homebrew)
  - [Other](#other)
- [Usage](#usage)
  - [List active tasks](#list-active-tasks)
  - [Add tasks](#add-tasks)
    - [to current week](#to-current-week)
    - [to a specific week](#to-a-specific-week)
  - [Mark a task as done](#mark-a-task-as-done)
  - [Config paths](#config-paths)
    - [Read config paths](#read-config-paths)
    - [Set config paths](#set-config-paths)
- [Issues](#issues)
- [Planed features](#planed-features)
- [License](#license)

# Install

## Using homebrew

Add the formula to your system

```
brew tap wstnbrg/todocli https://github.com/wstnbrg/todocli
```

And then install it via:

```
brew install todocli
```

## Other

Download [latest releases](https://github.com/wstnbrg/todocli/releases/) archive file that fit's your system best, extract the archive and place todocli into your path. 


# Usage

## List active tasks

```
todocli
```

will show your current active tasks. Done will not be displayed. Output looks like:

```
General:
    [ ] create this awesome cli tool 
    [ ] add some tasks to it 

somemagicproject:
    [ ] add some feature
```

## Add tasks

Adding a task without the optional +{projectname} argument will add it to a "General" project

### to current week

```
todocli add some task [+someproject]
```

### to a specific week

```
todocli add some task [+someproject] @31/2020
```

## Mark a task as done

```
todocli done 2 +someproject
```

This will mark the second listed task of the +someproject as done. Omitting +someproject will use the "General" project instead.

## Config paths

### Read config paths

```
todocli conf get configPath
```

Reading out config path 'taskdir'

```
todocli conf get taskdir
```

### Set config paths

```
todocli conf set configPath configValue
```

Please note that the argument configValue is the only argument that allows to have spaces in it.
Setting the configPath taskdir

```
todocli conf set taskdir /some/path/to/save/to
```

# Issues
- [ ] cant handle backlog from 2 weeks ago
- [ ] cant handle backlog after year change

# Planed features
- [ ] show tasks of a specific week
- [ ] configure to save todos per week or day
- [ ] multi language support

If you like anything else feel free to contact me or hack it in your self ;)

# License

This project is licensed under the MIT License - see the [LICENSE](https://github.com/wstnbrg/todocli/blob/master/LICENSE) file for details