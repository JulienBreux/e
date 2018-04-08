# e.Rror package

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
