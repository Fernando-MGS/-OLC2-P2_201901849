package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"

	"github.com/colegno/arraylist"
)

type Struct struct {
	Id        string
	Atributos *arraylist.List
	Line      string
	Col       string
}

func NewStruct(id string, list *arraylist.List, line, col int) Struct {
	instr := Struct{id, list, strconv.Itoa(line), strconv.Itoa(col)}
	return instr
}

func (p Struct) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	new_struct := interfaces.Structs{Id: p.Id, Atributo: p.Atributos}
	env.(environment.Environment).SaveStruct(p.Line, p.Col, p.Id, new_struct)
	return retorno
}
