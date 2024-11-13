package entity

type Request struct {
	Num1      float64 `json:"num1"`
	Num2      float64 `json:"num2"`
	Operation string  `json:"operation"`
}

func (r Request) GetInput() string {
	return ""
}
