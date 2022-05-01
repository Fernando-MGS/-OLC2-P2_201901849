package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"

	"github.com/colegno/arraylist"
)

type Functions struct {
	Id        string
	Params    *arraylist.List
	Tipo      interfaces.TipoSimbolo
	Statments *arraylist.List
	Line      string
	Col       string
}

func NewFunctions(id string, tipo interfaces.TipoSimbolo, params, instr *arraylist.List, line, col int) Functions {
	return Functions{Id: id, Params: params, Tipo: tipo, Statments: instr, Line: strconv.Itoa(line), Col: strconv.Itoa(col)}
}

func (f Functions) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	newFunc := interfaces.Functions{Id: f.Id, Params: f.Params, Tipo: f.Tipo, Statments: f.Statments}
	tmp := environment.NewEnvironment(env.(environment.Environment), f.Id, "", "", false, 0)
	if env.(environment.Environment).SaveFunc(f.Line, f.Col, f.Id, newFunc) {
		for _, s := range f.Statments.ToArray() {
			s.(interfaces.Instruction).Ejecutar(tmp, gen)
		}
		retorno_conf := gen.GetReturn(f.Id)
		if retorno_conf {
			if f.Tipo.Tipo == interfaces.NULL {
				env.(environment.Environment).NewError("NO SE ESPERABA UN RETORNO", f.Line, f.Col)
			}
		} else {
			if f.Tipo.Tipo != interfaces.NULL {
				env.(environment.Environment).NewError("SE ESPERABA UN RETORNO", f.Line, f.Col)
			} else {
				code := generator.Obligatorio{"return;", false, ""}
				fragment := generator.Fragment{Valor: code, Tipo: 13, Valido: true, Mod: false}
				gen.AddFragment(f.Id, fragment)
			}
		}
	}
	return retorno
}
