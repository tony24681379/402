package config

import (
	"flag"

	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

type Configs struct {
	Port     string
	MongoURL string
}

func Config() *Configs {
	flag.Set("alsologtostderr", "true")
	flag.Set("v", "2")
	flag.CommandLine.Parse([]string{})
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.String("port", "4000", "serve port")
	pflag.String("mongo_url", "127.0.0.1:27017", "serve port")
	pflag.Parse()
	viper.BindPFlags(pflag.CommandLine)

	return &Configs{
		Port:     ":" + viper.GetString("port"),
		MongoURL: viper.GetString("mongo_url"),
	}
}
