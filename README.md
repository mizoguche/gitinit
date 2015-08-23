gitinit
---

git init and copy hooks for langage

## Usage
### 1. create hooks templates
create langage hooks template directory in `~/.gitinit`. and create `hooks` in it.

```sh
$ mkdir -p ~/.gitnit/go/hooks
$ touch ~/.gitnit/go/hooks/pre-commit
```

### 2. run
```sh
$ gitinit go
```

so make current directory git repository, copy `~/.gitinit/go/hooks/*` to `.git/hooks/` and chmod all copied files 755..

## License
MIT