# gohashdeep
A cousin of hashdeep written in Go.

## Example run

### Suppose we have the following folder structure
```
~/go/src/github.com/ren-zxcyq/gohashdeep$ tree .
.
├── filelist.go
├── filelist.go.bak
└── testinguh
    └── asdf

1 directory, 3 files
```

### Running results in the following
```
~/go/bin/gohashdeep
Path              UID    GID    Size   Hash
----              ---    ---    ----   ----
.                 1000   1000   4096   
.gitignore        1000   1000   6      bb1aa666fd6ad90d04ecff9c442ea7fe66d702991388efa2512df1bb723e117f
filelist.go       1000   1000   2426   8048fdc4f911b95b1bdad2742e0770d5e3c59aaf170fc65eb81c55e45675b14a
filelist.go.bak   1000   1000   2501   b1bdb5752e3ed8df20af1038a03236de00ede3be1af27413bc7e0188c5d85ebd
testinguh         1000   1000   4096   
testinguh/asdf    1000   1000   0      e3b0c44298fc1c149afbf4c8996fb92427ae41e4649b934ca495991b7852b855
```