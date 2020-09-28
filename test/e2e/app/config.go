package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	ChainID          string `toml:"chain_id"`
	Listen           string
	GRPC             bool `toml:"grpc"`
	Dir              string
	PersistInterval  uint64                      `toml:"persist_interval"`
	SnapshotInterval uint64                      `toml:"snapshot_interval"`
	RetainBlocks     uint64                      `toml:"retain_blocks"`
	ValidatorUpdates map[string]map[string]uint8 `toml:"validator_update"`
	PrivValServer    string                      `toml:"privval_server"`
	PrivValKey       string                      `toml:"privval_key"`
	PrivValState     string                      `toml:"privval_state"`
}

func LoadConfig(file string) (*Config, error) {
	cfg := &Config{
		Listen:          "unix:///var/run/app.sock",
		GRPC:            false,
		PersistInterval: 1,
	}
	r, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("failed to open app config %q: %w", file, err)
	}
	_, err = toml.DecodeReader(r, &cfg)
	if err != nil {
		return nil, fmt.Errorf("invalid config: %w", err)
	}
	return cfg, cfg.Validate()
}

func (cfg Config) Validate() error {
	// We don't do exhaustive config validation here, instead relying on Testnet.Validate()
	// to handle it.
	switch {
	case cfg.ChainID == "":
		return errors.New("chain_id parameter is required")
	case cfg.Listen == "":
		return errors.New("listen parameter is required")
	default:
		return nil
	}
}