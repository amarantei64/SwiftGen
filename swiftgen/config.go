package swiftgen

import (
	"path/filepath"
	"path"
)

type Config struct {
	ArchConfig      string
	PlRemesa        string
	PlLBTR          string
	PlInternacional string
	PlRemesaCaja    string
	DatosClientes   string
	RutaDestino     string
}

func (c *Config) Load() {
	c.ArchConfig = filepath.Base()
}