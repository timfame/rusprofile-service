package logger

import "go.uber.org/zap"

type Logger struct {
	config *Config
	logger *zap.Logger
}

// Default log level is DEBUG
func New(opts ...Option) *Logger {
	config := &Config{}
	for _, o := range opts {
		o(config)
	}

	result := &Logger{
		config: config,
	}

	zapOptions := []zap.Option{
		zap.AddCaller(),
		zap.AddCallerSkip(1),
	}

	if config.zapConfig != nil {
		result.logger, _ = config.zapConfig.Build(zapOptions...)
	} else {
		result.logger, _ = zap.NewProduction(zapOptions...)
	}

	return result
}

func (l *Logger) Sync() {
	_ = l.logger.Sync()
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, l.addServiceName(fields)...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.logger.Fatal(msg, l.addServiceName(fields)...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.logger.Error(msg, l.addServiceName(fields)...)
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	if l.config.level != DEBUG {
		return
	}
	l.logger.Debug(msg, l.addServiceName(fields)...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.logger.Warn(msg, l.addServiceName(fields)...)
}

func (l *Logger) addServiceName(fields []zap.Field) []zap.Field {
	if l.config.serviceName != "" {
		fields = append(fields, zap.String("service", l.config.serviceName))
	}
	return fields
}
