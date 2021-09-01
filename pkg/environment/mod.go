package environment

// type Mod struct {
// 	id          string
// 	name        string
// 	desc        string
// 	previewUrl  string
// 	minecraft   string
// 	filename    string
// 	screenshots []string
// 	repository  Repository
// }

type Mod struct {
	Id    string
	Mcver string
	Type  string
	File  string
}

type ModsList []Mod

func (ml *ModsList) GetById(id string) (int, Mod) {
	for idx, mod := range *ml {
		if mod.Id == id {
			return idx, mod
		}
	}
	return -1, Mod{}
}
