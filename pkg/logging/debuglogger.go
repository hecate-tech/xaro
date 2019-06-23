package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// DebugLogger is the default logger used for the Quirk client.
// This logger will not print anything out.
type DebugLogger struct {
	logger *zap.Logger
}

var _ Logger = &DebugLogger{}

// NewDebugLogger returns a nil logging
// object for the Quirk client to use.
func NewDebugLogger() *DebugLogger {
	l, _ := zap.NewDevelopment()
	return &DebugLogger{
		logger: l,
	}
}

// Info does nothing.
func (l *DebugLogger) Info(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Info(msg, fields...)
}

// Debug logs nothing.
func (l *DebugLogger) Debug(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Debug(msg, fields...)
}

// Error logs nothing.
func (l *DebugLogger) Error(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Error(msg, fields...)
}

// Warn does nothing.
func (l *DebugLogger) Warn(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Warn(msg, fields...)
}

// Fatal logs nothing.
func (l *DebugLogger) Fatal(msg string, iFields ...interface{}) {
	fields := interfaceToZapField(iFields...)
	l.logger.Fatal(msg, fields...)
}

// interfaceToZapField takes the interfaces passed in and type asserts them
// into a zap.Field and returns a slice.
func interfaceToZapField(iFields ...interface{}) (fields []zapcore.Field) {
	for i := 0; i < len(iFields); i++ {
		fields = append(fields, iFields[i].(zapcore.Field))
	}
	return
}