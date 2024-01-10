# LSH
LSH (Lazy SSH Human) is a CLI tool to manage and simplify SSH connections, 
designed for lazy humans like me.

## Installation
- Download latest release
- Extract the `lsh` binary to `~/local/bin` or any other directory in your `$PATH`

### Installation for lazy humans
```
curl -sL https://github.com/mirkobrombin/LSH/releases/latest/download/lsh -o ~/.local/bin/lsh && chmod +x ~/.local/bin/lsh && echo "Installation done, start using LSH by typing lsh and remember there's nothing wrong with being lazy 🥱!"
```

## Usage
```
lsh [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  connect     Connect to a saved bookmark
  export      Export bookmarks
  help        Help about any command
  import      Import bookmarks
  list        List all saved bookmarks
  remove      Remove a saved bookmark
```

### How to create bookmarks
```
lsh connect [user@host]
```

LSH will identify the connection string and prompt you for a bookmark name, 
which will be used to save the connection details.

You can use any SSH compatible connection string, for example:

```
lsh connect -- user@host:port
```

the full connection string will be stored in the bookmark.

## TODO
This is a list of features I'd like to implement in the future:
- [ ] Export SSH keys along with bookmarks
- [ ] Store bookmarks in a secure way
- [ ] Improve SSH command parsing