package models

import "fmt"

type Repository struct {
	Tag string
	Url string
}

type RepositoryList []Repository

func (rl *RepositoryList) GetByTag(tag string) (Repository, error) {
	for _, repo := range *rl {
		if repo.Tag == tag {
			return repo, nil
		}
	}
	return Repository{}, fmt.Errorf("tag '%v' not found", tag)
}
