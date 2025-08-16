package config

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/pkg/errors"
	gormpgsql "github.com/shishir54234/NewsScraper/backend/pkg/database"
	"github.com/shishir54234/NewsScraper/backend/pkg/grpc"
	echoserver "github.com/shishir54234/NewsScraper/backend/pkg/http/echo/server"
	"github.com/shishir54234/NewsScraper/backend/pkg/logger"
	"github.com/shishir54234/NewsScraper/backend/pkg/otel"
	"github.com/shishir54234/NewsScraper/backend/pkg/rabbitmq"
	"github.com/spf13/viper"
)


type ConfigLLMClient struct {
	ConnAddr string
}

var configPath string

type Config struct {
	ServiceName  string                        `mapstructure:"serviceName"`
	Logger       *logger.LoggerConfig          `mapstructure:"logger"`
	Rabbitmq     *rabbitmq.RabbitMQConfig      `mapstructure:"rabbitmq"`
	Echo         *echoserver.EchoConfig        `mapstructure:"echo"`
	Grpc         *grpc.GrpcConfig              `mapstructure:"grpc"`
	ConfigLLMClient *ConfigLLMClient          `mapstructure:"configLLMClient"`
	GormPostgres *gormpgsql.PostgresConfig `mapstructure:"gormPostgres"`
	Jaeger       *otel.JaegerConfig            `mapstructure:"jaeger"`
}


func init() {
	flag.StringVar(&configPath, "config", "", "products write microservice config path")
}

func InitConfig() (*Config, *logger.LoggerConfig, *otel.JaegerConfig, *gormpgsql.PostgresConfig,
	*grpc.GrpcConfig, *echoserver.EchoConfig, *rabbitmq.RabbitMQConfig, *ConfigLLMClient, error) {
	fmt.Println("JLOKER", configPath)
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	if configPath == "" {
		configPathFromEnv := os.Getenv("CONFIG_PATH")
		if configPathFromEnv != "" {
			configPath = configPathFromEnv
		} else {
			//https://stackoverflow.com/questions/31873396/is-it-possible-to-get-the-current-root-of-package-structure-as-a-string-in-golan
			//https://stackoverflow.com/questions/18537257/how-to-get-the-directory-of-the-currently-running-file
			d, err := dirname()
			if err != nil {
				fmt.Println("JLOKER", configPath)	
				return nil, nil, nil, nil, nil, nil, nil,nil, err
			}

			configPath = d
		}
	}

	cfg := &Config{}

	viper.SetConfigName(fmt.Sprintf("config.%s", env))
	viper.AddConfigPath(configPath)
	viper.SetConfigType("json")

	// fmt.Println("JLOKER", cfg.configLLMClient.ConnAddr)
	if err := viper.ReadInConfig(); err != nil {
	// fmt.Println("JLOKER", cfg.configLLMClient.ConnAddr)
		fmt.Println("error in reading config", err)
		return nil, nil, nil, nil, nil, nil, nil,nil, errors.Wrap(err, "viper.ReadInConfig")
	}

	if err := viper.Unmarshal(cfg); err != nil {
		// fmt.Println("JLOKER", cfg.configLLMClient.ConnAddr)
		fmt.Println("error in unmarshalling", err)
		return nil, nil, nil, nil, nil, nil, nil, nil,errors.Wrap(err, "viper.Unmarshal")
	}
	if(cfg.ConfigLLMClient==nil){
		fmt.Println("config llm client is nil")
	}
	if(cfg.GormPostgres==nil){
		fmt.Println("gorm postgres is nil")
	}
	if(cfg.Logger==nil){
		fmt.Println("logger is nil")
	}
	if(cfg.Jaeger==nil){
		fmt.Println("jaeger is nil")
	}
	if(cfg.Grpc==nil){
		fmt.Println("grpc is nil")
	}
	if(cfg.Echo==nil){
		fmt.Println("echo is nil")
	}
	if(cfg.Rabbitmq==nil){
		fmt.Println("rabbitmq is nil")
	}
	return cfg, cfg.Logger, cfg.Jaeger, cfg.GormPostgres, cfg.Grpc, cfg.Echo, cfg.Rabbitmq, cfg.ConfigLLMClient, nil
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
