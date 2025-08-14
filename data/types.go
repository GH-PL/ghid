package data

type Hash struct {
	Regex string  `json:"regex"`
	Modes []Modes `json:"modes"`
}
type Modes struct {
	John    *string  `json:"john"`
	Hashcat *uint    `json:"hashcat"`
	Name    string   `json:"name"`
	Samples []string `json:"samples"`
}
