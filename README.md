# Gout Cobra

Integrates [gout](https://github.com/drewstinnett/gout) and [cobra](https://github.com/spf13/cobra) with some low code helpers.

For a full example, take a look at [./examples/basic/main.go](./eamples/basic/main.go)

## Usage

Bind the flags needed for configuring Gout to your cobra.Command like so:

```go
goutbra.Bind(cmd)
```

By default, this creates `--format` and `--format-template` flags. You can change this by passing in options, like:

```go
goutbra.Bind(cmd, WithField("my-special-format"))
```

Which would produce the flags `--my-special-format` and `--my-special-format-template`

When you are ready to use Gout, apply the cobra.Command options like:

```go
goutbra.Cmd(cmd)
```

This will set the builtIn Gout instance to the options provived by your cobra.Command