# `nonstd`

[![PkgGoDev](https://img.shields.io/badge/-reference-blue?logo=go&logoColor=white&labelColor=505050)](https://pkg.go.dev/github.com/thediveo/nonstd)
[![License](https://img.shields.io/github/license/thediveo/nonstd)](https://img.shields.io/github/license/thediveo/nonstd)
![build and test](https://github.com/thediveo/nonstd/actions/workflows/buildandtest.yaml/badge.svg?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/thediveo/nonstd)](https://goreportcard.com/report/github.com/thediveo/nonstd)
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

A small collection of generic little helpers, deemed to be too "abstractive" (a
[portmaneau](https://en.wikipedia.org/wiki/Portmanteau) of _abstract_ and
_distractive_) to sully the stdlib. Of course, your mileage will vary widely.

`nonstd` is on purpose a much smaller and much less comprehensive collection
compared to all the existing kitchen sinks out there.

- `sets` when you're tired of `map[T]bool` and especially `map[T]struct{}` when
  just needing a set of `T`s.
- `xiter`:
  - `Filter` and `Filter2` to filter sequences.
  - `FirstOk` and `First2Ok`: get the first element of a sequence.
  - `Map` transforms sequences.
  - `Swap` (k,v) into (v,k) sequences.
- `xmaps`: `Collect`, `CollectT` where the same key maps to multiple values.
- `xslices`:
  - just `All` the values without the pesky indices please.
  - `Filter` returning a new slice with only the elements passing the test
    function.
  - `Map` transforms slices.
  - `Any` to convert slices of other types into `[]any` slices.
  - `ReflectValues` to convert an slice into a slice of `Reflect.Value`s.
  - `SortedCopy`, `SortedCopyFunc` and `StableSortedCopyFunc` sort copies of
    slices.
  - `Slice` turns a list of uniformly typed elements into a slice.
- `xstrings`:
  - `CutWhitespace` a string at the first white space.
  - `JoinSpaced` for joining string slice elements, ensuring there's only one
    space between them.
- `xslog`:
  - `Lazy` returns a `slog.Attr` that evaluates its value via a function only
    when actually needed.

# DevContainer

> [!CAUTION]
>
> Do **not** use VSCode's "~~Dev Containers: Clone Repository in Container
> Volume~~" command, as it is utterly broken by design, ignoring
> `.devcontainer/devcontainer.json`.

1. `git clone https://github.com/thediveo/nonstd`
2. in VSCode: Ctrl+Shift+P, "Dev Containers: Open Workspace in Container..."
3. select `nonstd.code-workspace` and off you go...

## Supported Go Versions

`nonstd` supports versions of Go that are noted by the [Go release
policy](https://golang.org/doc/devel/release.html#policy), that is, major
versions _N_ and _N_-1 (where _N_ is the current major version).

## Contributing

Please see [CONTRIBUTING.md](CONTRIBUTING.md).

## Copyright and License

`nonstd` is Copyright 2026 Harald Albrecht, and licensed under the Apache
License, Version 2.0.

