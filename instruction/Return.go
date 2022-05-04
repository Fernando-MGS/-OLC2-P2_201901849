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
	ex := env.(environment.Environment).DevSalida()
	if valor.Type == interfaces.BOOLEAN {
		resultado := gen.NewTemp()
		salida := gen.NewLabel()
		gen.NewLabels(valor.TrueLabel, false, "", name, true, true, "line")
		gen.NewAsignacion(resultado, "1", false, "", name, true, false, "")
		gen.NewSalto(salida, false, "", name, true, true, "")
		gen.NewLabels(valor.FalseLabel, false, "", name, true, true, "line")
		gen.NewAsignacion(resultado, "0", false, "", name, true, false, "")
		gen.NewLabels(salida, false, "", name, true, true, "line")
		gen.NewStack("P", resultado, false, "", name, true, false, "")
		gen.NewSalto(ex, false, "", name, true, true, "")
	} else {
		gen.NewStack("P", valor.Value, false, "", name, true, false, "")
		gen.NewSalto(ex, false, "", name, true, true, "")
	}
	//gen.NewReturn("", false, "", name, true, false, "")
	retorno.Type = interfaces.NULL
	return retorno
}
