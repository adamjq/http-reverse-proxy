package config

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestConfig_valid(t *testing.T) {
	assert := require.New(t)
	t.Setenv("PORT", "9000")
	t.Setenv("SERVICE_A_URL", "https://my-service-a.com")
	t.Setenv("SERVICE_B_URL", "https://my-service-b.com")

	cfg := NewConfig()

	err := cfg.validate()
	assert.NoError(err)

	assert.Equal("9000", cfg.Port)
	assert.Equal("https://my-service-a.com", cfg.ServiceAUrl)
	assert.Equal("https://my-service-b.com", cfg.ServiceBUrl)
}

func TestConfig_defaults(t *testing.T) {
	assert := require.New(t)

	cfg := NewConfig()

	err := cfg.validate()
	assert.NoError(err)

	assert.Equal("5000", cfg.Port)
	assert.Equal("http://localhost:5001", cfg.ServiceAUrl)
	assert.Equal("http://localhost:5002", cfg.ServiceBUrl)
}
