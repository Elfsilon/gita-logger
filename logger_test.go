package gita

import (
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDefaultLogger(t *testing.T) {
	logger := NewLogger(nil)
	assert.Equal(t, logger.level, InfoLevel)
	assert.Equal(t, logger.ctx.eventsCount, 0)
	assert.False(t, logger.formatter.DisplayID)
	assert.False(t, logger.formatter.DisplayTimestamp)
	assert.True(t, logger.formatter.DisplayStackTrace)
	assert.True(t, logger.formatter.DisplayLocation)
	assert.Equal(t, logger.formatter.TimestampFormat, "15:01:02")
	assert.Nil(t, logger.formatter.IDStyle)
	assert.Nil(t, logger.formatter.TimestampStyle)
	assert.Nil(t, logger.formatter.MessageStyle)
	assert.NotNil(t, logger.formatter.LevelStyle)
	assert.NotNil(t, logger.formatter.StackTraceStyle)
}

func TestLoggerLog(t *testing.T) {
	logger := NewLogger(nil)
	err := logger.log("Test", InfoLevel, 0)
	require.Nil(t, err)
	require.Equal(t, logger.ctx.eventsCount, 1)
}

func BenchmarkLog(b *testing.B) {
	b.SetParallelism(100)
	logger := NewLogger(&Config{
		Out: io.Discard,
		Err: io.Discard,
	})

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			logger.log("Test", InfoLevel, 0)
		}
	})
}
