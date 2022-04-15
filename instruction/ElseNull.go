package instruction

import (
	"OLC2/generator"
	"OLC2/interfaces"
)

type ElseNull struct {
	null string
}

func NewElseNull(n string) ElseNull {
	expr := ElseNull{n}
	return expr
}

func (e ElseNull) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var result interfaces.Value
	result.Type = interfaces.NULL
	return result
}
