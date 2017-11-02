package configuration

import (
	"fmt"
	"github.com/gogather/com"
	"gopkg.in/ini.v1"
	"strconv"
)

type Config struct {
	config *ini.File
	path   string
}

func NewConfig(path string) *Config {
	var cfg *ini.File
	cfgExist := com.FileExist(path)
	if cfgExist {
		cfg, _ = ini.Load(path)
	} else {
		cfg = ini.Empty()
	}

	dusCfg := &Config{
		config: cfg,
		path:   path,
	}

	if !cfgExist {
		dusCfg.initWithDefault()
	}

	return dusCfg
}

func (dc *Config) initWithDefault() (err error) {
	dc.Get("system", "port", "9080")
	dc.Get("mongodb", "host", "127.0.0.1")

	return nil
}

func (dc *Config) Get(section, key string, value string) string {
	sect, err := dc.config.GetSection(section)
	if err != nil {
		sect, _ = dc.config.NewSection(section)
	}

	val, err := sect.GetKey(key)
	if err != nil {
		sect.NewKey(key, value)
		dc.config.SaveTo(dc.path)
		return value
	}

	return val.String()
}

func (dc *Config) GetInt64(section, key string, value int64) int64 {
	intStr := dc.Get(section, key, fmt.Sprintf("%d", value))

	i, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return value
	} else {
		return i
	}
}
