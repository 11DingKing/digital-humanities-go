package config

import "fmt"

func (c Config) Validate() error {
	if c.Addr == "" {
		return fmt.Errorf("address required")
	}
	if c.DatabasePath == "" {
		return fmt.Errorf("database path required")
	}
	if c.WorkerIntervalSeconds < 1 {
		return fmt.Errorf("worker interval invalid")
	}
	return nil
}
