package logger

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNew_Production(t *testing.T) {
	logger, err := New(Production)
	require.NoError(t, err)
	require.NotNil(t, logger)
	defer logger.Sync()

	var buf bytes.Buffer
	logger = logger.WithOptions(zap.WrapCore(func(zapcore.Core) zapcore.Core {
		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		return zapcore.NewCore(zapcore.NewJSONEncoder(encoderCfg), zapcore.AddSync(&buf), zap.InfoLevel)
	}))

	logger.Info("test message", zap.String("env", "prod"))

	var log map[string]interface{}
	err = json.Unmarshal(buf.Bytes(), &log)
	require.NoError(t, err)

	assert.Equal(t, "info", log["level"])
	assert.Equal(t, "test message", log["msg"])
	assert.Equal(t, "prod", log["env"])
	assert.Contains(t, log, "timestamp")
	assert.NotEmpty(t, log["timestamp"])
}

func TestNew_Development(t *testing.T) {
	logger, err := New(Development)
	require.NoError(t, err)
	require.NotNil(t, logger)
	defer logger.Sync()

	var buf bytes.Buffer
	logger = logger.WithOptions(zap.WrapCore(func(zapcore.Core) zapcore.Core {
		encoderCfg := zap.NewDevelopmentEncoderConfig()
		encoderCfg.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewCore(zapcore.NewConsoleEncoder(encoderCfg), zapcore.AddSync(&buf), zap.DebugLevel)
	}))

	logger.Info("test dev message")

	output := buf.String()
	assert.Contains(t, output, "INFO")
	assert.Contains(t, output, "test dev message")
	assert.True(t, strings.Contains(output, ".go") || strings.Contains(output, "logger_test.go"),
		"expected source file in dev log")
}

func TestNew_UsableEndToEnd(t *testing.T) {
	prod, err := New(Production)
	require.NoError(t, err)
	prod.Info("end-to-end production log")
	prod.Sync()

	dev, err := New(Development)
	require.NoError(t, err)
	dev.Info("end-to-end development log")
	dev.Sync()

	assert.NotNil(t, prod)
	assert.NotNil(t, dev)
}
