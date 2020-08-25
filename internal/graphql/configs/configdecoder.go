package config

// Code generated by github.com/gokultp/go-envparser. DO NOT EDIT.
import (
	"errors"
	"os"

	commonconfigs "github.com/firstcontributions/firstcontributions/internal/configs"
	"github.com/gokultp/go-envparser/pkg/envdecoder"
)

func (t *Config) DecodeEnv() error {
	_recLog := commonconfigs.LogConfig{}
	if err := envdecoder.Decode(&_recLog); err != nil {
		return errors.New("env Decoder interface is not implemented for type commonconfigs.LogConfig")
	}
	t.Log = &_recLog
	if _recPortStr := os.Getenv("PROFILE_PORT"); _recPortStr != "" {
		_recPort := _recPortStr
		t.Port = &_recPort
	}
	return nil
}
