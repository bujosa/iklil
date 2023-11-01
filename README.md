# iklil
This is private company assigment related with golang position

## Instructions
###  package zap

A Logger provides fast, leveled, structured logging. All methods are safe
for concurrent use. [The documentation for the Logger type](https://pkg.go.dev/go.uber.org/zap#Logger)

The Logger is designed for contexts in which every microsecond and every
allocation matters, so its API intentionally favors performance and type
safety over brevity. For most applications, the SugaredLogger strikes a
better balance between performance and ergonomics.

The methods for Logger are:
```go
//  Sugar() *SugaredLogger
//  Named(s string) *Logger
//  WithOptions(opts ...Option) *Logger
//  With(fields ...Field) *Logger
//  Level() zapcore.Level
//  Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
//  Log(lvl zapcore.Level, msg string, fields ...Field)
//  Debug(msg string, fields ...Field)
//  Info(msg string, fields ...Field)
//  Warn(msg string, fields ...Field)
//  Error(msg string, fields ...Field)
//  DPanic(msg string, fields ...Field)
//  Panic(msg string, fields ...Field)
//  Fatal(msg string, fields ...Field)
//  Sync() error
//  Core() zapcore.Core
//  Name() string
//  clone() *Logger
//  check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry
type Logger {
  // omitted
}
```


```go
// A SugaredLogger wraps the base Logger functionality in a slower, but less
// verbose, API. Any Logger can be converted to a SugaredLogger with its Sugar
// method.
//
// Unlike the Logger, the SugaredLogger doesn't insist on structured logging.
// For each log level, it exposes four methods:
//
//   - methods named after the log level for log.Print-style logging
//   - methods ending in "w" for loosely-typed structured logging
//   - methods ending in "f" for log.Printf-style logging
//   - methods ending in "ln" for log.Println-style logging
//
// For example, the methods for InfoLevel are:
//
//  Info(...any)           Print-style logging
//  Infow(...any)          Structured logging (read as "info with")
//  Infof(string, ...any)  Printf-style logging
//  Infoln(...any)         Println-style logging

Cannot use 'zapLogger' (type *zap.Logger) as the type retryablehttp.LeveledLogger Type does not implement 'retryablehttp.LeveledLogger' need the method: Error(msg string, keysAndValues ...interface{}) have the method: Error(msg string, fields ...Field)

Cannot use 'zapLogger.Sugar()' (type *SugaredLogger) as the type retryablehttp.LeveledLogger Type does not implement 'retryablehttp.LeveledLogger' need the method: Error(msg string, keysAndValues ...interface{}) have the method: Error(args ...interface{})

type SugaredLogger struct {
  base *Logger
}
```


### package retryablehttp

[LeveledLogger](github.com/hashicorp/go-retryablehttp) is an interface that can be implemented by any logger or a
logger wrapper to provide leveled logging. The methods accept a message
string and a variadic number of key-value pairs. For log.Printf style
formatting where message string contains a format specifier, use Logger
interface.
```go
type LeveledLogger interface {
  Error(msg string, keysAndValues ...interface{})
  Info(msg string, keysAndValues ...interface{})
  Debug(msg string, keysAndValues ...interface{})
  Warn(msg string, keysAndValues ...interface{})
}
```

Logger interface allows to use other loggers than standard log.Logger.
```go
type Logger interface {
    Printf(string, ...interface{})
}
```