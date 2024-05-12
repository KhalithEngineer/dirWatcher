package models

import (
	"sync"
	"time"
)

type Config struct {
	Dir         string `json:"dir"`
	Interval    int    `json:"interval"`
	MagicString string `json:"magic_string"`
	CongigMut   sync.RWMutex
}

type TaskRun struct {
	ID           int
	StartTime    time.Time
	EndTime      time.Time
	Runtime      string
	Status       string
	FilesAdded   []string
	FilesDeleted []string
	Occurrences  int
	Directory    string
	MagicString  string
}

// SetConfig sets the configuration
func (c *Config) SetConfig(dir string, interval int, magicString string) {
	c.CongigMut.Lock()
	defer c.CongigMut.Unlock()
	c.Dir = dir
	c.Interval = interval
	c.MagicString = magicString
}

// GetConfig returns a copy of the configuration
func (c *Config) GetConfig() (string, int, string) {
	c.CongigMut.RLock()
	defer c.CongigMut.RUnlock()
	return c.Dir, c.Interval, c.MagicString
}
