/**
  @author: $(USER)
  @data:$(DATE)
  @note:
**/
package log

import (
	"context"
	"errors"
	"go.uber.org/zap"
	"testing"
)

func TestZap(t *testing.T) {
	zap1 := NewLog("test", false)
	zap1.Start(context.Background())
	logger := zap1.Logger
	logger.Debug("测试", zap.Error(errors.New("test errors")))
	logger.Info("测试", zap.Error(errors.New("test errors")))
	logger.Warn("测试", zap.Error(errors.New("test errors")))
	logger.Error("测试", zap.Error(errors.New("test errors")))
}
