# Zerolog

[Zerolog](https://github.com/rs/zerolog) logger implementation

## Usage

```go
func ExampleWithOut() {
  l := zero.NewLogger(zero.WithOut(os.Stdout), zero.WithLevel(logger.DebugLevel))

  l.Logf(logger.InfoLevel, "testing: %s", "logf")

  // Output:
  // {"level":"info","message":"testing: logf"}
}
```
