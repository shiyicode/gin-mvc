package config

import (
	"io/ioutil"
	"os"
	"sync"

	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

type Config struct {
	Run   RunConfig   `toml:"run"`
	Log   LogConfig   `toml:"log"`
	Mysql MysqlConfig `toml:"mysql"`
	Jwt   JwtConfig   `toml:"jwt"`
}

type RunConfig struct {
	WaitTimeout int    `toml:"waitTimeout"`
	HTTPPort    int    `toml:"httpPort"`
	Mode        string `toml:"mode"`
	MaxAllowed  int    `toml:"maxAllowed"`
}

type LogConfig struct {
	Enable     bool   `toml:"enable"`
	Path       string `toml:"path"`
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
	MaxEffectiveTime int64  `toml:"maxEffectiveTime"`
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
		log.Fatalln("config file not specified: use -c $filename")
	}
	return _config
}

// 加载配置文件
func Load(configFile string) {
	if err := loadConfig(); err != nil {
		log.Fatal(errors.Wrapf(err, "load file from %s failed", _configFile))
	}

	_configFile = configFile

	log.Infof("load file from %s success: %s, config: %#v", _configFile, _config)
}

func ReLoad() {
	if err := loadConfig(); err != nil {
		log.Error(errors.Wrapf(err, "load file from %s failed", _configFile))
	}
	log.Infof("reload file from %s success: %s, config: %#v", _configFile, _config)
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
