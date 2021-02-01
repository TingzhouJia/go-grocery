package config
import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
)
type Config struct {
	Name string
}

func (c *Config) init() error {
	if c.Name !=""{
		viper.SetConfigFile(c.Name)
	}else{
		viper.AddConfigPath("db")
		viper.SetConfigName("config")
	}
	viper.SetConfigType("yaml")
	if err:=viper.ReadInConfig();err!=nil{
		return err
	}
	return nil
}

func (c *Config) watchConfig()  {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("log change")
	})
}

func Init(name string)  error{
	c:=Config{Name: name}
	if err:=c.init();err!=nil{
		return err
	}
	c.watchConfig()
	return nil
}

