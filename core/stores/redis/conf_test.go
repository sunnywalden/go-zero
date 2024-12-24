package redis

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sunnywalden/go-zero/core/stringx"
)

func TestRedisConf(t *testing.T) {
	tests := []struct {
		name string
		RedisConf
		ok bool
	}{
		{
			name: "missing host",
			RedisConf: RedisConf{
				Host:     "",
				Type:     NodeType,
				Username: "",
				Password: "",
			},
			ok: false,
		},
		{
			name: "missing type",
			RedisConf: RedisConf{
				Host:     "localhost:6379",
				Type:     "",
				Username: "",
				Password: "",
			},
			ok: false,
		},
		{
			name: "ok",
			RedisConf: RedisConf{
				Host:     "localhost:6379",
				Type:     NodeType,
				Username: "",
				Password: "",
			},
			ok: true,
		},
		{
			name: "ok",
			RedisConf: RedisConf{
				Host:     "localhost:6379",
				Type:     ClusterType,
				Username: "user",
				Password: "pwd",
				Tls:      true,
			},
			ok: true,
		},
	}

	for _, test := range tests {
		t.Run(stringx.RandId(), func(t *testing.T) {
			if test.ok {
				assert.Nil(t, test.RedisConf.Validate())
				assert.NotNil(t, test.RedisConf.NewRedis())
			} else {
				assert.NotNil(t, test.RedisConf.Validate())
			}
		})
	}
}

func TestRedisKeyConf(t *testing.T) {
	tests := []struct {
		name string
		RedisKeyConf
		ok bool
	}{
		{
			name: "missing host",
			RedisKeyConf: RedisKeyConf{
				RedisConf: RedisConf{
					Host:     "",
					Type:     NodeType,
					Username: "",
					Password: "",
				},
				Key: "foo",
			},
			ok: false,
		},
		{
			name: "missing key",
			RedisKeyConf: RedisKeyConf{
				RedisConf: RedisConf{
					Host:     "localhost:6379",
					Type:     NodeType,
					Username: "",
					Password: "",
				},
				Key: "",
			},
			ok: false,
		},
		{
			name: "ok",
			RedisKeyConf: RedisKeyConf{
				RedisConf: RedisConf{
					Host:     "localhost:6379",
					Type:     NodeType,
					Username: "",
					Password: "",
				},
				Key: "foo",
			},
			ok: true,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if test.ok {
				assert.Nil(t, test.RedisKeyConf.Validate())
			} else {
				assert.NotNil(t, test.RedisKeyConf.Validate())
			}
		})
	}
}
