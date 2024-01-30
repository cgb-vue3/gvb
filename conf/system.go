package conf

import "fmt"

type System struct {
	//Host string `yaml:"host" json:"host,omitempty"`
	Port int
	Env  string
}

func (Sys System) Addr() string {
	return fmt.Sprintf(":%d", Sys.Port)
}
