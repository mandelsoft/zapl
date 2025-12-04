package zapl

import (
	"fmt"

	"go.uber.org/zap"
)

func x() {
	var l *zap.Logger

	l, _ = zap.NewDevelopment()
	fmt.Printf("name: %s\n", l.Name())
}
