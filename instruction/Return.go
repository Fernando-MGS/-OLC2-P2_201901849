package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Return struct {
	Valor interfaces.Expresion
	Line  string
	Col   string
}

func NewReturn(valor interfaces.Expresion, line, col int) Return {
	return Return{valor, strconv.Itoa(line), strconv.Itoa(col)}
}

func (p Return) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	name := env.(environment.Environment).Control.Id
	valor := p.Valor.Ejecutar(env, gen)
	gen.NewStack("P", valor.Value, false, "", name, true, false, "")
	gen.NewReturn("", false, "", name, true, false, "")
	retorno.Type = interfaces.NULL
	return retorno
}
