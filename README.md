# todocli <!-- omit in toc -->

ToDoCli is a simple cli saving your time by managing your tasks from your terminal.

- [Install](#install)
- [Usage](#usage)
  - [List active tasks](#list-active-tasks)
  - [Add tasks](#add-tasks)
    - [to current week](#to-current-week)
    - [to a specific week](#to-a-specific-week)
  - [Mark a task as done](#mark-a-task-as-done)
- [Issues](#issues)
- [Todo for first release](#todo-for-first-release)
- [Planed features](#planed-features)
- [License](#license)

# Install

{ToBeFilled}

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

# Issues

# Todo for first release
- [x] adding tasks
- [x] displaying current weeks tasks
- [x] check done tasks
- [ ] handle backlog

# Planed features
- [ ] multi language support

If you like anything else feel free to contact me or hack it in your self ;)

# License

This project is licensed under the MIT License - see the LICENSE file for details