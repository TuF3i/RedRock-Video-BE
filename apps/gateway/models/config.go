package models

type Config struct {
	Hertz Hertz
}

type Hertz struct {
	IPAddr string
	Port   string
}
