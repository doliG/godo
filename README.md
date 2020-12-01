# Godo

CLI interface for managing todolist.

## Quickstart

At the moment, we only have a version for "developper", meaning you have some knowledge of what a package manager is, and is used to CLI.

We will work on a proper installation with a package manager in the future, maybe with version 2 or 3.

**Requirements:**

- You must have go version 1.15 or higher installed. [> Install Go](https://golang.org/dl/)

**Then, run:**

```
# Install

# Add in path
```

## Usage

```bash
# Add a todo
godo add "Do something"

# List todos
godo list [-a]

# Mark as done / undone
godo toggle [id id2 id3...]

# Edit
godo edit [id] [-m "New name"]
```

## What's next ?

I do have a lot of stuff to do, and a lot of ideas. If you're interessed in develpment of this package, feel free to open an issue or a PR.

\> Read more in [MILESTONES.md](./MILESTONES.md) 