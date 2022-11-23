package driver

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"runtime"
	"strings"
)

type Config struct {
	Http struct {
		Mode  string `json:"mode"`
		Host  string `json:"host"`
		Port  string `json:"port"`
		Pprof struct {
			Port string `json:"port"`
		} `json:"pprof"`
	}

	Logger struct {
		Path string `json:"path"`
	}

	Jwt struct {
		Secret string `json:"secret"`
		Expire string `json:"expire"`
	}
	Database struct {
		Driver   string `json:"driver"`
		Host     string `json:"host"`
		Database string `json:"database"`
		User     string `json:"user"`
		Password string `json:"password"`
		Port     string `json:"port"`
	}

	Redis struct {
		Host     string `json:"host"`
		Port     string `json:"port"`
		Password string `json:"password"`
		Db       string `json:"db"`
	}
}

var (
	Conf *Config
)

// ReadConfig 解析配置文件并监听配置文件变化
func ReadConfig(path string) error {
	if path == "" {
		return errors.New("cannot find config file ")
	}
	//解析path路径
	ps := strings.Split(path, "/")
	//文件路径前缀
	p := strings.Join(ps[:len(ps)-1], "/")

	//读取指定配置文件
	Conf = new(Config)
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	if runtime.GOOS == "windows" {
		v.AddConfigPath("configs")
	} else {
		v.AddConfigPath(p)
	}
	if err := v.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}
	if err := v.Unmarshal(Conf); err != nil {
		return errors.WithStack(err)
	}

	//监听配置文件变化
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		if err := ReadConfig(path); err != nil {
			fmt.Println("ConfigServer config file change has an error")
		}
	})

	return nil
}
