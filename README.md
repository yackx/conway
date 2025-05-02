# conway

A simple [Conway's Game of Life](http://en.wikipedia.org/wiki/Conway's_Game_of_Life) written in Go and in Haskell.

* No cycle detection
* Infinite grid
* No GUI

## golang

### Build and run

```bash
$ go build
$ ./conway
```

### Unit tests

    $ go test
    PASS
    ok      github.com/yackx/conway 0.171s

## Haskell

### Build and run

```bash
$ ghc conway.hs
$ ./conway
```

## Sample output

    [(0, 0) (1, 0) (2, 0) (1, 1) (2, 1)]
    [(0, 0) (2, 0) (2, 1) (1, -1) (0, 1)]
    [(0, 0) (2, 0) (1, -1)]
    [(1, -1) (1, 0)]
    []

## License

[GNU GPL v.3](LICENSE)
