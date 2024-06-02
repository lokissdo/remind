package configs

import (
	"io/ioutil"

	"github.com/creasty/defaults"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)


type YAML struct {
	Port int `yaml:"port" validate:"required"`
	JWT  struct {
		Secret               string `yaml:"secret" validate:"required"`
		AccessTokenDuration  int    `yaml:"access_token_duration" default:"10"`
		RefreshTokenDuration int    `yaml:"refresh_token_duration" default:"60"`
		Iuser                string `yaml:"iuser" default:"todo"`
	} `yaml:"jwt" validate:"required"`
	PostgresQL struct {
		Connection string `yaml:"connection" validate:"required"`
		Driver     string `yaml:"driver" default:"postgres"`
	} `yaml:"postgres" validate:"required"`
	
	AppEnvironment string `yaml:"app_environment" default:"development"`
	Log            struct {
		Level string `yaml:"level" default:"debug"`
		File  string `yaml:"file" default:""`
	} `yaml:"log"`
	// SchoolID string `yaml:"school_id" validate:"required"`
}

func (s *YAML) UnmarshalYAML(unmarshal func(interface{}) error) error {
	defaults.Set(s)

	type plain YAML
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}

	return nil
}

var Yaml YAML

func init() {
	LoadYAML()
}

func LoadYAML() {

	yamlFile, err := ioutil.ReadFile("conf.yaml")
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Init env error")
		panic(err)
	}
	err = yaml.Unmarshal(yamlFile, &Yaml)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Init env error")
		panic(err)
	}
	err = validator.New().Struct(Yaml)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("Init env error")
		panic(err)
	}
}
