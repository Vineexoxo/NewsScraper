package configurations

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	"github.com/spf13/viper"
)

type Llm_config struct {
	ApiKey  string
	BaseURL string
}
var configPath string 
type Config struct {
	ServiceName string
	LLMConfig   *Llm_config
	Grpc *grpc.GrpcConfig
}
func init(){
	flag.StringVar(&configPath, "config", "", "generating description microservices")
}

func initConfig()(*Config, *Llm_config, *grpc.GrpcConfig, error){
	env:=os.Getenv("APP_ENV")
	if env == ""{
		env="development"
	}
	if configPath==""{
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv!="" {
			configPath=configPathFromEnv
		}else {
			d,err:=dirname()
			if err!=nil{
				return nil ,err
			}
			configPath=d
		}
	}
	cfg := &Config{}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}
	if err := viper.Unmarshal(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}

func GetMicroserviceName(serviceName string) string {
	return fmt.Sprintf("%s", strings.ToUpper(serviceName))
}

func filename() (string, error) {
	_, filename, _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("unable to get the current filename")
	}
	return filename, nil
}

func dirname() (string, error) {
	filename, err := filename()
	if err != nil {
		return "", err
	}
	return filepath.Dir(filename), nil
}
