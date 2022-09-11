package logger

import (
	"fmt"

	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	logger, err := zap.NewProduction()
	if err != nil {
		return nil, fmt.Errorf("zap logger failed: %w", err)
	}

	return logger, nil
}
