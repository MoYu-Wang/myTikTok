package setting

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Conf = new(AppConfig)

type MySQLConfig struct {
	Host         string `mapstructure:"host"`
	DB           string `mapstructure:"db"`
	User         string `mapstructure:"user"`
	Password     string `mapstructure:"password"`
	Port         int    `mapstructure:"port"`
	MaxOpenConns int    `mapstructure:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"maxIdleConns"`
}

// AppConfig 需要初始化的项目
type AppConfig struct {
	StartTime string `mapstructure:"startTime"`
	MachineID int64  `mapstructure:"machineID"`

	*MySQLConfig `mapstructure:"mysql"`
}

func Init(filePath string) (err error) {

	viper.SetConfigFile(filePath)
	err = viper.ReadInConfig()
	if err != nil {
		fmt.Printf("viper.ReadInConfig failed, err:%v\n", err)
		return err
	}

	// unmarshal the config info to the struct AppConfig(pointer Conf)
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Printf("viper.Unmarshal failed, err:%v\n", err)
		return err
	}
	// listen the Config
	viper.WatchConfig()
	// when it changes
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("config info be changed...")
		// update the AppConfig
		if err := viper.Unmarshal(Conf); err != nil {
			zap.L().Error("viper.Unmarshal failed, err:", zap.Error(err))
			return
		}
	})
	return
}
