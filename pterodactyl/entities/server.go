package entities

type Server struct {
	Attributes ServerAttributes
}

type ServerAttributes struct {
	Identifier string `json:"identifier"`
	Name string `json:"name"`
	UUID string `json:"uuid"`
}