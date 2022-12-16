package repository

type File struct {
	Name    string "json:name"
	Type    string "json:type"
	Content string "json:content"
}
type Commit struct {
	Id      string "json:id"
	Message string "json:message"
	Files   []File "json:files"
}
type Tag struct {
	Id      string   "json:id"
	Name    string   "json:name"
	Commits []Commit "json:commits"
}
type Data struct {
	Tags []Tag "json:tags"
}
type RemoteRepository struct {
	Name string "json:name"
	Uri  string "json:uri"
}
type Config struct {
	Remote []RemoteRepository "json:remote"
}
