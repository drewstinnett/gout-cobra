# Gout Cobra

[![Go Reference](https://pkg.go.dev/badge/github.com/drewstinnett/gout-cobra.svg)](https://pkg.go.dev/github.com/drewstinnett/gout-cobra)
[![Go Report Card](https://goreportcard.com/badge/github.com/drewstinnett/gout-cobra)](https://goreportcard.com/report/github.com/drewstinnett/gout-cobra)

Integrates [gout](https://github.com/drewstinnett/gout) and [cobra](https://github.com/spf13/cobra) with some low code helpers.

For a full example, take a look at [./examples/basic/main.go](./examples/basic/main.go)

## Usage

Bind the flags needed for configuring Gout to your cobra.Command like so:

```go
goutbra.Bind(cmd)
```

By default, this creates `--format` and `--format-template` flags. You can change this by passing in options, like:

```go
goutbra.Bind(cmd, goutbra.WithField("my-special-format"))
```

Which would produce the flags `--my-special-format` and `--my-special-format-template`

When you are ready to use Gout, apply the cobra.Command options like:

```go
goutbra.Cmd(cmd)
```

This will set the builtIn Gout instance to the options provived by your cobra.Command
