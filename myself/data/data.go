package data

type List struct {
	Name     string `json:"name"`
	LastName string `json:"lastname"`
	Year     int    `json:"year"`
}

type JsonWork struct {
	Lists []List `json:"lists"`
}

func NewJsonWork() *JsonWork {
	variable := &JsonWork{
		Lists: []List{
			{
				Name:     "1",
				LastName: "2",
				Year:     1981},
			{
				Name:     "3",
				LastName: "4",
				Year:     1982},
		},
	}
	return variable
}
