Chameleon
===

![](https://dl.dropboxusercontent.com/u/7817937/_github/film/l-chameleons.jpg)

A dynamic image resizing server.

`http://example.com/images/53bcf9166b616e10c0020000/large.100x100.jpg` => 100x100 resized image of `http://s3.amazonaws.com/filmapp-development/images/53bcf9166b616e10c0020000/large.jpg`

usage
---

1. setup following environment variables `AWS_ACCESS_KEY`, `AWS_SECRET_ACCESS_KEY`
2. `gom install`
3. `gom build` to obtain a binary.

or just `go run main.go` to start listen `:3000`.

Related Projects
---

https://github.com/ReshNesh/pixlserv
