// Package configuration 配置模块
package configuration

import (
	"fmt"
	"strconv"
	"sync"

	"github.com/gogather/com"
	"gopkg.in/ini.v1"
)

type Config struct {
	config   *ini.File
	path     string
	writeLck *sync.Mutex
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
		config:   cfg,
		path:     path,
		writeLck: &sync.Mutex{},
	}

	if !cfgExist {
		dusCfg.initWithDefault()
	}

	return dusCfg
}

func (dc *Config) initWithDefault() (err error) {
	dc.Get("system", "port", "9080")

	return nil
}

func (dc *Config) write(path string, content string) error {
	defer dc.writeLck.Unlock()
	dc.writeLck.Lock()
	return com.WriteFile(path, content)
}

// GetSectionAsMap 将配置区加载为map
func (dc *Config) GetSectionAsMap(section string) map[string]string {
	sect, err := dc.config.GetSection(section)
	if err != nil {
		sect, _ = dc.config.NewSection(section)
		dc.writeLck.Lock()
		dc.config.SaveTo(dc.path)
		dc.writeLck.Unlock()
	}

	result := map[string]string{}
	keys := sect.Keys()
	for _, key := range keys {
		result[key.Name()] = key.Value()
	}
	return result
}

// GetSection 获取配置区
func (dc *Config) GetSection(section string) *ini.Section {
	sect, err := dc.config.GetSection(section)
	if err != nil {
		sect, _ = dc.config.NewSection(section)
	}
	return sect
}

// Get 获取配置项，section,key为配置区与键，value为默认值，当配置项不存在时返回value默认值，且初始化该值到配置文件
func (dc *Config) Get(section, key string, value string) string {
	sect, err := dc.config.GetSection(section)
	if err != nil {
		sect, _ = dc.config.NewSection(section)
	}

	val, err := sect.GetKey(key)
	if err != nil {
		sect.NewKey(key, value)
		dc.writeLck.Lock()
		dc.config.SaveTo(dc.path)
		dc.writeLck.Unlock()
		return value
	}

	return val.String()
}

// Set 运行时主动设置配置项值，并存入配置文件
func (dc *Config) Set(section, key string, value string) {
	sect, err := dc.config.GetSection(section)
	if err != nil {
		sect, _ = dc.config.NewSection(section)
	}

	sect.NewKey(key, value)
	dc.writeLck.Lock()
	dc.config.SaveTo(dc.path)
	dc.writeLck.Unlock()
}

// SectionExist 配置区是否存在
func (dc *Config) SectionExist(section string) bool {
	_, err := dc.config.GetSection(section)
	if err != nil {
		return false
	} else {
		return true
	}
}

// GetInt64 获取配置项值为int64类型
func (dc *Config) GetInt64(section, key string, value int64) int64 {
	intStr := dc.Get(section, key, fmt.Sprintf("%d", value))

	i, err := strconv.ParseInt(intStr, 10, 64)
	if err != nil {
		return value
	} else {
		return i
	}
}
