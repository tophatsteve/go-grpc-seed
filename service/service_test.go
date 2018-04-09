package sample

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateMessageReturnsExpectedMessage(t *testing.T) {
	m := generateMessage()
	assert.Equal(t, "Success", m, "received message is not correct")
}
