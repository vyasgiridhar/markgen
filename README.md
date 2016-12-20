# MarkGen
Markdown Generator in Go Lang

#

## How it works

to run `markgen`, run the command `markgen`
```
markgen README.md
listening :6060 ...
```

And this will open a browser window which shows the content of a file.
The files are the served in `https://localhost:6060/file name`

When you modify the file, `markgen` will send the modified data to the browser through a websocket connection. You don't even need to refresh the page.

To install, Run

```
$ go get github.com/vyasgiridhar/markgen/markgen
```


## Contribution

I welcome every kind of contribution.

If you have any problem using `markgen`, please file an issue in
[Issues](https://github.com/vyasgiridhar/markgen/issues).

If you'd like to contribute on source, please upload a pull request in
[Pull Requests](https://github.com/vyasgiridhar/markgen/pulls).