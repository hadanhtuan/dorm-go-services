package es

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/spf13/viper"
)

type ESEnv struct {
	Host     string `mapstructure:"ES_HOST"`
	Port     int    `mapstructure:"ES_PORT"`
	Username string `mapstructure:"ES_USER"`
	Password string `mapstructure:"ES_PWD"`
}

type ESClient struct {
	Client *elasticsearch.TypedClient
}

var (
	ES    *ESClient
	esEnv ESEnv
)

func ConnectElasticSearch() *ESClient {
	ES = new(ESClient)

	ParseENV(&esEnv)

	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	esUri := fmt.Sprintf(
		"http://%s:%d",
		esEnv.Host,
		esEnv.Port,
	)

	esCnf := elasticsearch.Config{
		Addresses: []string{esUri},
		Username:  esEnv.Username,
		Password:  esEnv.Password,
	}

	client, err := elasticsearch.NewTypedClient(esCnf)

	if err != nil {
		slog.Info(err.Error())
	}

	ES.Client = client

	return ES
}
func GetConnection() *ESClient {
	if ES != nil {
		return ES
	}
	return ConnectElasticSearch()
}

func ParseENV[T interface{}](object T) error {
	err := viper.Unmarshal(object)
	if err != nil {
		return err
	}
	return nil
}
