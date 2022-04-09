## todo

```
todo is a CLI To-Do list manager

Usage:
  todo [command]

Available Commands:
  add         Adds a task to your task list.
  clear       Clears the done/removed list
  del         Permanently remove/delete the task from your task list.
  done        Marks the task as done/completed
  help        Help about any command
  list        Lists all the tasks in your task list.
  rm          Remove the task from your to-do list.

Flags:
  -h, --help   help for todo

Use "todo [command] --help" for more information about a command.

```

## todo add

```
Adds a task to your task list.

Usage:
  todo add [flags]

Examples:

        todo add your task

        todo add "your task"
```

## todo list

shows the active tasks in the list

### There are 3 lists in the to-do app

- active tasks list
      - can be listed by **todo list** or **todo list -a** commands
-  done tasks list

      - can be listed by **todo list -d** command
      - done tasks list contains the list of tasks that are marked as done
- removed tasks list
      - can be listed by **todo list -r**
      - removed tasks list contains the list of tasks that are marked as removed

_3 lists can be displayed by **todo list -dra** command_


#### todo del

```
Permanently remove/delete the task from your task list.

Usage:
  todo del [flags]

Examples:

        todo del 1 2
                - here 1 and 2 are the Task number of the active task
        todo del 3
                - here 3 is the Task number of the active task

```

## todo done

```
Marks the task as done/completed

Usage:
  todo done [flags]

Examples:

        todo done 1 2
                - here 1 and 2 are the Task number of the active task
        todo done 3
                - here 3 is the Task number of the active task
```

## todo clear

```
Clears the done/removed list

Usage:
  todo clear [flags]
  todo clear [command]

Examples:

        todo clear done
                - clears the done lists
        todo clear rm
                - clears the removed task list


Available Commands:
  done        Clears the done tasks list
  rm          Clears the removed tasks list
```

## todo rm

```
Remove the task from your to-do list.

Usage:
  todo rm [flags]

Examples:

        todo rm 1 2
                - here 1 and 2 are the Task number of the active task
        todo rm 3
                - here 3 is the Task number of the active task
```

