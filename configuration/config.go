package configuration

type Configuration struct {
	Server   Server `json:"server"`
	Database DB     `json:"database"`
}

type Server struct {
	Port string `json:"port"`
}
type DB struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Dbname   string `json:"dbname"`
}
