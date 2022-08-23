package recovery

import (
	"context"
	"errors"
	"fmt"

	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/log"
	"go.uber.org/zap"
	"lib/core/v1/logger"
	"runtime/debug"
)

func castError(raw any) error {
	var err error
	switch v := raw.(type) {
	case error:
		err = v
	case string:
		err = errors.New(fmt.Sprintf("panic recovered: %s", v))
	default:
		err = errors.New("undefined error")
	}
	return err
}

func RecoverAndLog(ctx context.Context) {
	if r := recover(); r != nil {

		err := castError(r)
		stack := string(debug.Stack())

		if ctx != nil {
			span := opentracing.SpanFromContext(ctx)
			if span != nil {
				span.SetTag("error", true)
				span.SetTag("panic", true)
				span.LogFields(log.Error(err))
			}
		} else {
			ctx = context.Background()
		}
		logger.ErrorKV(
			ctx,
			fmt.Sprintf("%v", err),
			zap.String("x_recovery_stack", stack),
			zap.Bool("x_recovery_panic", true),
		)
	}
}
