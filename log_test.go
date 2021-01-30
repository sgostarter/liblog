package liblog

import (
	"context"
	"testing"

	"github.com/jiuzhou-zhao/go-fundamental/interfaces"
	"github.com/stretchr/testify/assert"
)

func test1(log interfaces.Logger) {
	log.Record(context.Background(), 1, interfaces.LogLevelDebug, "test1")
}

func TestDepth(t *testing.T) {
	log, err := NewZapLogger()
	assert.Nil(t, err)
	log.Record(context.Background(), 0, interfaces.LogLevelDebug, "test")
	test1(log)
}
