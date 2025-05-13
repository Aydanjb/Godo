# Godo

A lightweight command-line task manager built in Go.  

---

## Features

- Add, list, update, and delete tasks
- Mark tasks as `Todo`, `InProgress`, or `Done`
- Persistent storage via local JSON file (`tasks.json`)
- Intuitive CLI powered by `urfave/cli/v2`
- Fully unit-tested core with clean modular design

---

## Installation

Download the latest version from [Releases](https://github.com/Aydanjb/Godo/releases)

```bash
# Linux
curl -LO https://github.com/Aydanjb/Godo/releases/latest/download/godo-linux-amd64
chmod +x godo-linux-amd64
./godo-linux-amd64 list

# macOS
curl -LO https://github.com/Aydanjb/Godo/releases/latest/download/godo-darwin-amd64
chmod +x godo-darwin-amd64
./godo-darwin-amd64 add "Test from Mac"

# Windows
Download and run godo-windows-amd64.exe from PowerShell
```

## Usage

### Add a task
```bash
./godo add "Write unit tests"
```

### Update a task
```bash
./godo update taskID "Refactor task system"
```

### Delete a task
```bash
./godo delete taskID
```

### List all tasks

```bash
./godo list
```

### Mark task as todo, done, or in-progress
```bash
./godo mark-todo taskID
./godo mark-done taskID
./godo mark-in-progress taskID
```


