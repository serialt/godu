package main

import (
	"log/slog"
	"os"

	"github.com/serialt/crab"
)

func service() {
	slog.Debug("debug msg")
	slog.Info("info msg")
	slog.Error("error msg")

}

func EnvGet(envName string, defaultValue string) (data string) {
	data = os.Getenv(envName)
	if len(data) == 0 {
		data = defaultValue
		return
	}
	return
}

func (c *Config) DecryptConfig() {
	if c.Encrypt {
		crab.AESDecryptCBCBase64(c.Token, AesKey)
		slog.Debug(c.Token)
	}
}
