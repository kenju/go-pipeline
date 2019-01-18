# go-pipeline

[![CircleCI](https://circleci.com/gh/kenju/go-pipeline.svg?style=svg)](https://circleci.com/gh/kenju/go-pipeline)

- [Documentation](https://godoc.org/github.com/kenju/go-pipeline)
- [Changelog](https://github.com/kenju/go-pipeline/blob/master/CHANGELOG.md)

# Install

```sh
$ go get github.com/kenju/go-pipeline
```

# Usage

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

for v := range pipeline.TakeInt(ctx, pipeline.GeneratorInt(ctx, 1, 2, 3, 4, 5), 3) {
    fmt.Printf("%v ", v)
    // 1 2 3
}
```

```go
ctx, cancel := context.WithCancel(context.Background())
defer cancel()

mapFn := func(v string) string { return "**" + v + "**" }
for v := range pipeline.MapString(ctx, mapFn, pipeline.GeneratorString(ctx, "hello", "world")) {
    fmt.Println(v)
    // **hello**
    // **world**
}
```

# Development

## Update CHANGELOG.md

```sh
$ npm install -g auto-changelog
$ git tag vx.x.x
$ make changelog
$ git push
```
