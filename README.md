# e.Rror package

[![CircleCI](https://circleci.com/gh/JulienBreux/e.svg?style=svg)](https://circleci.com/gh/JulienBreux/e) 
[![codecov](https://codecov.io/gh/JulienBreux/e/branch/master/graph/badge.svg)](https://codecov.io/gh/JulienBreux/e)
[![codebeat badge](https://codebeat.co/badges/153232b0-eca1-4384-959f-6ed2bea64772)](https://codebeat.co/projects/github-com-julienbreux-e-master)
[![Go Report Card](https://goreportcard.com/badge/github.com/JulienBreux/e)](https://goreportcard.com/report/github.com/JulienBreux/e)
[![GoDoc](https://godoc.org/github.com/JulienBreux/e?status.svg)](http://godoc.org/github.com/JulienBreux/e)
[![GitHub tag](https://img.shields.io/github/tag/JulienBreux/e.svg)](Tag)

## BEFORE

```go
_, err := my.Awesome().Method()
if err != nil {
    // ...
}
```

## AFTER

```go
if err := e.Rror(my.Awesome().Method()); err != nil {
    // ...
}
```

## Maintainer

- :octocat: [Julien Breux](https://github.com/JulienBreux) - [@JulienBreux](https://twitter.com/JulienBreux)
