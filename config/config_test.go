package config

import "testing"

func TestConfig(t *testing.T) {
	var testfile string = "../test.yaml"

	t.Run("Parse()", func(t *testing.T) {
		cfg, err := Parse(testfile)
		if err != nil {
			t.Error("parse file failed")
		}
		if cfg.Username != "test" {
			t.Error("parse bad user")
		}
		if cfg.Password != "test" {
			t.Error("parse bad password")
		}
		if cfg.Url != "test" {
			t.Error("parse bad url")
		}
	})
}