package config

import (
	"io/ioutil"
	"os"
	"sync"
	"time"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Run     RunConfig     `toml:"run"`
	Log     LogConfig     `toml:"log"`
	Mysql   MysqlConfig   `toml:"mysql"`
	Jwt     JwtConfig     `toml:"jwt"`
	Endless EndlessConfig `toml:"endless"`
}

type RunConfig struct {
	HTTPAddr   string `toml:"httpAddr"`
	Mode       string `toml:"mode"`
	MaxAllowed int    `toml:"maxAllowed"`
}

type LogConfig struct {
	Enable     bool   `toml:"enable"`
	Path       string `toml:"path"`
	Level      string `toml:"level"`
	RotateTime int    `toml:"rotateTime"`
	MaxAge     int    `toml:"maxAge"`
}

type MysqlConfig struct {
	MaxIdle int    `toml:"maxIdle"`
	MaxOpen int    `toml:"maxOpen"`
	Debug   bool   `toml:"debug"`
	WebAddr string `toml:"webAddr"`
}

type JwtConfig struct {
	EncodeMethod     string `toml:"encodeMethod"`
	MaxEffectiveTime string `toml:"maxEffectiveTime"`
}

type EndlessConfig struct {
	ReadTimeOut    time.Duration `toml:"readTimeOut"`
	WriteTimeOut   time.Duration `toml:"writeTimeOut"`
	MaxHeaderBytes int           `toml:"maxHeaderBytes"`
	HammerTime     time.Duration `toml:"hammerTime"`
}

var (
	_configFile string
	_config     Config
	_lock       = new(sync.RWMutex)
)

func Get() Config {
	_lock.RLock()
	defer _lock.RUnlock()

	if _configFile == "" {
		log.Fatal("config file not load")
	}
	return _config
}

// 加载配置文件
func Load(configFile string) {
	_configFile = configFile

	if err := loadConfig(); err != nil {
		log.Fatal(errors.Wrapf(err, "load file from %s failed", _configFile))
	}

	log.Infof("load file from %s success; config: %#v", _configFile, _config)
}

func ReLoad() {
	if err := loadConfig(); err != nil {
		log.Error(errors.Wrapf(err, "load file from %s failed", _configFile))
	}
	log.Infof("reload file from %s success; config: %#v", _configFile, _config)
}

func loadConfig() error {
	_lock.Lock()
	defer _lock.Unlock()
	// 配置文件是否存在
	if _, err := os.Stat(_configFile); os.IsNotExist(err) {
		return err
	}

	bs, err := ioutil.ReadFile(_configFile)
	if err != nil {
		return err
	}

	if _, err := toml.Decode(string(bs), &_config); err != nil {
		return err
	}
	return nil
}
