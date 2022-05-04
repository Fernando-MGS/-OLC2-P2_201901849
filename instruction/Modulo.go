package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Mod struct {
	id   string
	Line string
	Col  string
}

func NewMod(id string, line, col int) Mod {
	return Mod{id, strconv.Itoa(line), strconv.Itoa(col)}
}

func (m Mod) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	a := interfaces.Bases{Nombre: m.id, NoTables: "", Linea: m.Line}
	env.(environment.Environment).LogBase(a)
	return retorno
}
