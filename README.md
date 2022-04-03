# todo
todo is an Golang CLI To-Do list manager

    todo is a CLI To-Do list manager

    Usage:
    todo [command]

    Available Commands:
    add         Adds a task to your task list.
    del         Permanently remove/delete the task from your task list.
    done        Marks the task as done/completed
    help        Help about any command
    list        Lists all the tasks in your task list.
    rm          Remove the task from your to-do list.

    Flags:
    -h, --help   help for todo

    Use "todo [command] --help" for more information about a command.




####todo list -adr

You have the following tasks in your to-do list:
Task No. Task
      1. 11112222333354444
      2. 11112222333354444555555

You have removed the following tasks from your to-do list:
Task No. Task
      1. 1111

You have Done the following tasks in your to-do list:
Task No. Task
      1. 11112222

####todo del 2
Task "2. 11112222333354444555555" is permanently deleted from your To-Do list

You have the following tasks in your to-do list:
Task No. Task
      1. 11112222333354444
                                                                                                                                        
####todo list -adr

You have the following tasks in your to-do list:
Task No. Task
      1. 11112222333354444

You have removed the following tasks from your to-do list:
Task No. Task
      1. 1111

You have Done the following tasks in your to-do list:
Task No. Task
      1. 11112222
