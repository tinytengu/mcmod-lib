package environment

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
