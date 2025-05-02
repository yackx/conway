# conway

A simple [Conway's Game of Life](http://en.wikipedia.org/wiki/Conway's_Game_of_Life) wrtten in [Go](http://golang.org/).

* No cycle detection
* Infinite grid
* No GUI

## Build and run

```bash
$ go build
$ ./conway
```

    [(0, 0) (1, 0) (2, 0) (1, 1) (2, 1)]
    [(0, 0) (2, 0) (2, 1) (1, -1) (0, 1)]
    [(0, 0) (2, 0) (1, -1)]
    [(1, -1) (1, 0)]
    []

## Unit tests

    $ go test
    PASS
    ok      github.com/yackx/conway 0.171s

## License

[GNU GPL v.3](LICENSE)
