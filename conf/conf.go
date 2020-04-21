package conf

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"time"
)

var (

	RunMode string
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	PageSize int
	JwtSecret 	string
)

type Config struct {
	Name string
}

var conf Config
func Init() error {
	conf = Config{
		Name: "app.yaml",
	}
	err := conf.initConfig()
	if err != nil {
		log.Println("读取配置文件错误")
		return err
	}
	conf.initAppConfig()
	// 热加载配置文件
	conf.watchConfig()
	return nil
}

func (c *Config)initConfig() error {
	// 设置viper
	if c.Name != "" {
		viper.SetConfigFile(c.Name)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("app.yaml")
	}
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()
	viper.SetEnvPrefix("")
	if err := viper.ReadInConfig(); err != nil {
		log.Println("ReadInConfig err:", err)
		return err
	}
	log.Println("读取配置文件成功")
	return nil
}

func (c *Config)initAppConfig() error  {
	RunMode = viper.GetString("RUN_MODE")
	HTTPPort = viper.GetInt("HTTP_PORT")
	ReadTimeout = viper.GetDuration("READ_TIMEOUT")
	WriteTimeout = viper.GetDuration("WRITE_TIMEOUT")
	PageSize = viper.GetInt("PAGE_SIZE")
	JwtSecret 	= viper.GetString("JWT_SECRET")

	return nil
}

// 热加载配置
func (c *Config) watchConfig()  {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Println("OnConfigChange:", in.Name)
	})
}