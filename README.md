# Welldone

Welldone is a CLI todolist.

## 🚀 Quickstart

Prerequisite: You need to have [Go 1.15.5](https://golang.org/dl/) or higher on your computer.

### Install script

```bash
curl "monscript" | bash
```

### Manual installation

```bash
go get -u github.com/doliG/welldone
go install github.com/doliG/welldone
welldone version
# Should print "welldone vX"
```

**Troubleshooting**

zsh: command not found: welldone

Just make sure that `GOPATH` is in your env. If not:
```
echo 'export PATH=$PATH:$(go env GOPATH)/bin' >> ~/.zshrc
source ~/.zshrc
# Or ".bashrc" or any ".rc" that you use to configure your shell.
```


<!-- ```
brew install godo
``` -->

## 📰 Usage

```bash
# Add a todo
welldone add "Do something"

# List todos
welldone list [-a]

# Mark as done / undone
welldone toggle [id id2 id3...]

# Edit
welldone edit [id] [-m "New name"]
```

## ℹ️ What's next ?

I do have a lot of stuff to do, and a lot of ideas. If you're interessed in develpment of this package, feel free to open an issue or a PR.

\> Read more in [MILESTONES.md](./MILESTONES.md) 
