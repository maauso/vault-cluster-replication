// Package logs provides a structured logs for the whole project
package logs

import (
	"time"

	"github.com/go-logr/logr"

	uzap "go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"sigs.k8s.io/controller-runtime/pkg/log/zap"
)

// Logger is the common var logs.

var Logger logr.Logger //nolint:gochecknoglobals

func init() { //nolint:gochecknoinits
	configLog := uzap.NewProductionEncoderConfig()
	configLog.EncodeTime = func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(ts.UTC().Format(time.RFC3339Nano))
	}
	Logger = zap.New(
		zap.Encoder(
			zapcore.NewJSONEncoder(configLog),
		),
	)
}
