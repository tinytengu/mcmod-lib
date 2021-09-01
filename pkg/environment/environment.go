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

func NewEnvironment(dirPath string) Environment {
	env := Environment{
		configFile: "mcmod.yaml",
		path:       dirPath,
	}

	env.Storage = Storage{
		Properties:   PropertiesList{},
		Repositories: RepositoryList{},
		Mods:         []Mod{},
	}

	return env
}

func (env *Environment) GetConfigPath() string {
	return path.Join(env.path, env.configFile)
}

func (env *Environment) GetPath() string {
	return env.path
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
