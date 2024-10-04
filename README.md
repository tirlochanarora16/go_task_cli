# Task Tracker

This is a simple task tracker application written in Go. It allows you to add, update, list, delete, and change the status of tasks. The tasks are stored in a JSON file.

## Features

- Add a new task
- List all tasks or filter by status (done, in-progress, not-done)
- Update a task's name
- Mark a task as in-progress or done
- Delete a task

## Commands

- `add "task name"`: Adds a new task with the given name.
- `list`: Lists all tasks.
- `list done`: Lists all tasks that are done.
- `list in-progress`: Lists all tasks that are in-progress.
- `list not-done`: Lists all tasks that are not done.
- `update <id> "new task name"`: Updates the name of the task with the given ID.
- `mark-in-progress <id>`: Marks the task with the given ID as in-progress.
- `mark-done <id>`: Marks the task with the given ID as done.
- `delete <id>`: Deletes the task with the given ID.
- `exit`: Exits the application.

## Usage

1. Run the application using `go run main.go`.
2. Enter commands to manage your tasks.

## Example

```sh
add "Buy groceries"
update 0 "Buy groceries and cook dinner"
mark-in-progress 0
list in-progress
mark-done 0
list done
delete 0
list
exit
```

## Project URL

```sh
https://roadmap.sh/projects/task-tracker
```

## License

This project is licensed under the MIT License.
