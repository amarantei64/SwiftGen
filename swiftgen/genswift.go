package swiftgen

import (
	"fmt"
)

type swiftgenEngine struct {
	Variables *map[string]Campo
	Tipo      string
	Cantidad  uint32
	Plantilla string
}

func (o *swiftgenEngine) New(tipo string, cantidad uint32, variables map[string]Campo) (*swiftgenEngine, error) {
	if variables == nil || len(variables) == 0 {
		return nil, fmt.Errorf("ValueError: Debe indicar las variables que se reemplazarán")
	}

	switch tipo {
	case "REMC", "REMD", "INT", "LBTR":
		o.Tipo = tipo
		o.Variables = &variables
		o.Cantidad = cantidad
		//Falta definir la configuracion
	default:
		return nil, fmt.Errorf("ValueError: Los tipos de mensajes admitidos son REMC, REMD, INT, LBTR")
	}

	return nil, nil
}
