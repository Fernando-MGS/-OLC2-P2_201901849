package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
)

type StructAccess struct {
	Id       interfaces.Expresion
	Atributo string
	Line     string
	Col      string
	//Valor    interfaces.Expresion
}

func NewStructAccess(id interfaces.Expresion, att string, line, col int) StructAccess {
	return StructAccess{Id: id, Atributo: att, Line: strconv.Itoa(line), Col: strconv.Itoa(col)}
}

func (p StructAccess) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	name := env.(environment.Environment).Control.Id
	access := p.Id.Ejecutar(env, gen)
	if access.Type == interfaces.STRUCT {
		conf := false
		var atributo interfaces.Atributos
		i := 0
		//fmt.Println("Access 2 es ", access.Tipo2.ToArray())
		for _, s := range access.Tipo2.ToArray() {
			atributo = s.(interfaces.Atributos)
			if atributo.Id == p.Atributo {
				conf = true
				break
			}
			i++
		}
		if conf {
			fmt.Println("EL ATRITBUTO A ACCEDER ES ", atributo)
			fmt.Println("EL VALUE ES ", access.Value)
			pos_valor := gen.NewTemp()
			res := gen.NewTemp()
			gen.NewOperacion(pos_valor, access.Value, "+", strconv.Itoa(i), false, "", name, true, true, "")
			gen.NewCallHeap(res, pos_valor, true, "POSICION EN EL HEAP DEL ATRIBUTO "+atributo.Id, name, true, false, "")
			retorno.Value = res
			retorno.Type = atributo.Tipo.Tipo
			retorno.Tipo2 = atributo.Tipo.Tipo2
			retorno.IsTemp = true
			retorno.FalseLabel = "2"
		} else {
			env.(environment.Environment).NewError("EL ATRIBUTO "+p.Atributo+" NO EXISTE EN EL STRUCT", p.Line, p.Col)
		}
	} else {
		env.(environment.Environment).NewError("EL ELEMENTO AL QUE SE DESEA ACCEDER NO ES UN STRUCT", p.Line, p.Col)
	}
	return retorno
}
