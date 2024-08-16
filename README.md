# Terminal Todo

A terminal cli based todo app written in go and cobra.

## Commands

- List - Print all todos

options -

--status = completed/pending

```sh
terminal-todo list
```

```sh
terminal-todo list --status=completed
```

- Add - Add a new todo

```sh
terminal-todo add <todo here>
```

- Update - Update a todo by id

options - 

id - todo id
--todo = todo to update 
--status = completed/pending

```sh
terminal-todo update <todoId> --todo=<todo to update> --status=<completed/pending>
```

- Delete - Delete a todo by id

options - 

id - todo id
--todo = todo to update 
--status = completed/pending

```sh
terminal-todo delete <todoId>
```

## Running Locally

### Clone repo

```sh
git clone https://github.com/pratikmane1299/terminal-todo.git
```

### Install dependencies

```sh
go mod tidy
```

### Build and run the cli

```sh
go build
```

```sh
./terminal-todo
```
