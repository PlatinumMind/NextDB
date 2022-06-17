package database;

type NextDatabase struct {
	name string `yaml:"name,omitempty"`
	username string `yaml:"username,omitempty"`
	password string	`yaml:"password,omitempty"`
	port int32 `yaml:"port,omitempty"`
}

type NextData struct {

}