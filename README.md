# Using ZAP logging with Mandelsoft Logging

This package provides a wrapper for the [zap logging system](https://github.com/uber-go/zap)
which uses the [mandelsoft logging](https://github.com/mandelsoft/logging) as backend
for a zap logger.

> **Note** If you want to use the zap logging system as backend for the mandelsoft logging
you can use the module https://github.com/go-logr/zapr to create a logr logger with
`zapr.NewLogger(zapLog)`, which can then be passed as base logger to the mandelsoft
logging system.

The wrapper maps named zap loggers
to the realm `zap/<logger name>`.

With `zapl.NewStatic` a new static mapping is created. The log level for the used mandelsoft loggers are fixed at the time the logger is first used to log a message.
To create a wrapper with unbound loggers, use `NewDynamic`.

Log levels are mapped according to their names. `Fatal`and `Panic` are mapped to `Error` level.

The wrapper by default sets the `Debug`level for the zap logger to pass the control about enabled levels to the mandelsoft logging configuration.

If you need to change this, set another level for the realm `zap` before creating the wrapper.

## Example

For an exaple how to use this module, see [example](example/main.go).


