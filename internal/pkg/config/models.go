package config

var paths = []string{
	"./config/config.json",
	"../config/config.json",
	"../../config/config.json",
}

var (
	Core   CoreConfig
	Db     DatabaseConfig
	Logger LoggerConfig
)

type CoreConfig struct {
	Port   string `json:"port"`
	Prefix string `json:"prefix"`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	Password string `json:"password"`
	Space    string `json:"space"`
}

type LoggerOut struct {
	Mode     string `json:"mode"`
	Filename string `json:"filename, omitempty"`
}

type LoggerConfig struct {
	Debug   LoggerOut `json:"debug"`
	Info    LoggerOut `json:"info"`
	Warning LoggerOut `json:"warning"`
	Error   LoggerOut `json:"error"`
	Fatal   LoggerOut `json:"fatal"`
}

type File struct {
	Core   CoreConfig     `json:"core"`
	Db     DatabaseConfig `json:"db"`
	Logger LoggerConfig   `json:"logger"`
}

func save(config File) {
	Core = config.Core
	Db = config.Db
	Logger = config.Logger
}
