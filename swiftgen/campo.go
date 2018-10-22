package swiftgen

import (
	"strings"
)

type TipoCampo uint8

const (
	CARACTER TipoCampo = iota + 1
	FLOTANTE
	ENTERO
)

type Campo struct {
	Tipo     TipoCampo
	Valor    string
	Posicion uint
}

func (c *Campo) New(tipo TipoCampo, valor string) *Campo {
	c.Tipo = tipo
	valor = strings.Replace(valor, "\"", "", -1)

	switch tipo {
	case FLOTANTE:
		c.Valor = strings.Replace(valor, ".", ",", -1)
	case ENTERO:
		c.Valor = strings.Replace(valor, ".", "", -1)
	default:
		c.Valor = valor
	}

	return c
}
