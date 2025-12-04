package main

import (
	"github.com/mandelsoft/logging"
	"github.com/mandelsoft/logging/logrusl"
	"github.com/mandelsoft/zapl"
	"go.uber.org/zap"
)

func main() {

	ctx := logrusl.Human().New()
	l := ctx.Logger(logging.NewRealm("test"))

	l.Info("a test message")
	l.Info("an {{type}} message", logging.KeyValue("type", "inline"))

	ctx.AddRule(logging.NewConditionRule(logging.DebugLevel, logging.NewRealmPrefix("zap")))
	ctx.Logger(logging.NewRealm("zap")).Debug("a debug message")
	log := zapl.NewStatic(ctx, "test")

	log.Info("a zap test message")
	log.Info("a second zap message with fields", zap.String("ctx", "first"))
	log.Info("combining field formatting: {{value}}", zap.String("value", "inline field"))

	log.Debug("a zap debug message")

}
