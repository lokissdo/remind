package middleware

import (
	"os"

	"github.com/rs/zerolog/log"

	"gopkg.in/yaml.v2"
)

type OpenAPISpec struct {
	Info struct {
		Title       string `yaml:"title"`
		Description string `yaml:"description"`
		Version 	string `yaml:"version"`
	} `yaml:"info"`
	Endpoint string `yaml:"endpoint"`
	Consumes []string `yaml:"consumes"`
	Produces []string `yaml:"produces"`
	Paths    map[string]struct {
		Security map[string]interface{} `yaml:"security"`
	} `yaml:"paths"`
}

func loadOpenAPISpec(filePath string) (*OpenAPISpec, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	var spec OpenAPISpec
	err = yaml.Unmarshal(data, &spec)
	if err != nil {
		return nil, err
	}

	return &spec, nil
}

type SecurityConfig struct {
	Path string
	AuthRequired bool
	AuthMethods []string 
}

func loadSecurityConfig(spec *OpenAPISpec) ([]SecurityConfig, error) {
	var configs []SecurityConfig
	for path, def := range spec.Paths {
		authRequired := false
		var authMethods []string
		for key := range def.Security {
			authRequired = true
			authMethods = append(authMethods, key)
		}
		configs = append(configs, SecurityConfig{
			Path: path,
			AuthRequired: authRequired,
			AuthMethods: authMethods,
		})
	}
	return configs, nil
}

func LoadAPISecurityConfig() (map[string]SecurityConfig, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Error().Err(err).Msg("failed to get working directory")
		return nil, err
	}
	dirPath := dir + "/api"

	allConfigs := make(map[string]SecurityConfig)

	folders, err := os.ReadDir(dirPath)
	if err != nil {
		log.Error().Err(err).Msg("failed to read dir")
		return nil, err
	}

	for _, folder := range folders {
		files, err := os.ReadDir(dirPath + "/" + folder.Name())
		if err != nil {
			log.Error().Err(err).Msg("failed to read dir")
			continue
		}

		for _, file := range files {
			path := dirPath + "/" + folder.Name() + "/" + file.Name()
			spec, err := loadOpenAPISpec(path)
			if err != nil {
				log.Error().Err(err).Msg("failed to load openapi spec")
				return nil, err
			}

			configs, err := loadSecurityConfig(spec)
			if err != nil {
				return nil, err
			}

			for _, config := range configs {
				allConfigs[config.Path] = config
			}
		}
	}

	return allConfigs, nil
}