package log551

type Config struct {
	Debug       bool `json:"debug"`
	Information bool `json:"information"`
	Warning     bool `json:"warning"`
	Error       bool `json:"error"`
	Critical    bool `json:"critical"`
	Path        string `json:"path"`
}

type log551 struct {}

func New() *log551 {
	return &log551{}
}