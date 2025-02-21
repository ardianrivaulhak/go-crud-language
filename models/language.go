package models


type Relation struct {
	InfluencedBy []string `json:"influenced-by"`
	Influences   []string `json:"influences"`
}

type Language struct {
	Language       string   `json:"language"`
	Appeared       int      `json:"appeared"`
	Created        []string `json:"created"`
	Functional     bool     `json:"functional"`
	ObjectOriented bool     `json:"object-oriented"`
	Relation       Relation `json:"relation"`
}


var Languages = []Language{
	{
		Language:      "C",
		Appeared:      1972,
		Created:       []string{"Dennis Ritchie"},
		Functional:    true,
		ObjectOriented: false,
		Relation: Relation{
			InfluencedBy: []string{"B", "ALGOL 68", "Assembly", "FORTRAN"},
			Influences:   []string{"C++", "Objective-C", "C#", "Java", "JavaScript", "PHP", "Go"},
		},
	},
}


func GetLanguages() []Language {
	return Languages
}

func AddLanguage(language Language) {
	Languages = append(Languages, language)
}

func UpdateLanguage(index int, language Language) {
	Languages[index] = language
}

func DeleteLanguage(index int) {
	Languages = append(Languages[:index], Languages[index+1:]...)
}