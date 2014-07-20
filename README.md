Chameleon
===

A dynamic image resizing server.

`http://example.com/images/53bcf9166b616e10c0020000/large.100x100.jpg` => 100x100 resized image of `http://s3.amazonaws.com/filmapp-development/images/53bcf9166b616e10c0020000/large.jpg`

usage
---

1. setup envitonment variable `AWS_ACCESS_KEY_ID`, `AWS_SECRET_ACCESS_KEY`
2. `gom install`
3. `gom build` to obtain a binary.
