package fastforex

type History struct {
	Base    string      `json:"base"`
	Updated string      `json:"updated"`
	Result  interface{} `json:"result"`
}