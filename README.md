# LeetCode practice

My LeetCode solutions, unfinished attempts, and programming experiments.
This is a practice repository; code may be incomplete or incorrect.

## Go

The Go code lives in [`leet_go/`](leet_go/), using Go 1.23.4 or newer.
Numbered folders match LeetCode question IDs. Each uses `package main` so
problems can be run independently:

```sh
cd leet_go
go run ./2469
```

[`util/`](leet_go/util/) contains shared helpers. [`random/`](leet_go/random/)
contains standalone Go experiments; run those files individually.

## TODO

- [ ] Finish the attempts in `3`, `14`, `76`, `263`, and `3310`.
- [ ] Complete or revisit the empty `SquareTilingLayout` experiment.
- [ ] Separate the programs in `random/` so the folder can build as a package.
- [ ] Add tests and check solutions against examples and edge cases.
- [ ] Add the Java practice code later.

`go build ./...` currently fails because `263` has no active `main` function
and `random/` has multiple `main` functions. There are no Go test files yet.
