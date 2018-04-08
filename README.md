# e.Rror package

[![CircleCI](https://circleci.com/gh/JulienBreux/e.svg?style=svg)](https://circleci.com/gh/JulienBreux/e)

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
