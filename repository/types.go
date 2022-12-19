package repository

type Picture struct {
	Name   string `json:"name"`
	Binary string `json:"bin"`
}
type SourceCode struct {
	Name    string `json:"name"`
	Content string `json:"content"`
}
type Commit struct {
	Id       string       `json:"id"`
	Message  string       `json:"message"`
	Codes    []SourceCode `json:"codes"`
	Pictures []Picture    `json:"pictures"`
}
type Section struct {
	Id      string   `json:"id"`
	Name    string   `json:"name"`
	Commits []Commit `json:"commits"`
}
type Data struct {
	Sections []Section `json:"sections"`
}
type RemoteRepository struct {
	Name string `json:"name"`
	Uri  string `json:"uri"`
}
type Config struct {
	Remote []RemoteRepository `json:"remote"`
}
