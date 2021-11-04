package abraxas

import (
	"testing"

	log "github.com/sirupsen/logrus"
	logt "github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func Test_Invoke(t *testing.T) {
	logger := log.New()
	hook := logt.NewLocal(logger)

	Invoke(logger)

	assert.Equal(t, answer, hook.LastEntry().Message)
}

func Test_Mediate(t *testing.T) {
	logger := log.New()
	hook := logt.NewLocal(logger)

	Mediate(logger)

	assert.Equal(t, answer, hook.LastEntry().Message)
}
