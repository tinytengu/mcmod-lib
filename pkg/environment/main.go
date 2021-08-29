package environment

import (
	"mcmodlib/pkg/utils"
	"os"
	"path"

	"gopkg.in/yaml.v2"
)

type Environment struct {
	configFile string
	path       string
	Storage
}

type Storage struct {
	Properties   PropertiesList `yaml:"properties"`
	Repositories RepositoryList
	Mods         []Mod `yaml:"mods"`
}

func NewEnvironment(dirPath string) Environment {
	env := Environment{}

	env.configFile = "mcmod.yaml"
	env.path = dirPath

	env.Properties = PropertiesList{}
	env.Repositories = RepositoryList{}
	env.Mods = []Mod{}

	return env
}

func (env *Environment) GetConfigPath() string {
	return path.Join(env.path, env.configFile)
}

func (env *Environment) Write() error {
	data, err := yaml.Marshal(&env.Storage)
	if err != nil {
		return err
	}

	err = os.MkdirAll(env.path, os.ModePerm)
	if err != nil {
		return err
	}

	err = utils.WriteFile(env.GetConfigPath(), data)
	if err != nil {
		return err
	}
	return nil
}

func (env *Environment) Read() error {
	data, err := utils.ReadFile(env.GetConfigPath())
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(data, &env.Storage)
	if err != nil {
		return err
	}

	return nil
}
