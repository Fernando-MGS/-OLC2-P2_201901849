package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"

	"github.com/colegno/arraylist"
)

type Struct struct {
	Id    string
	Lista *arraylist.List
	Line  string
	Col   string
}

func NewStruct(id string, l *arraylist.List, line, col int) Struct {
	return Struct{Id: id, Lista: l, Line: strconv.Itoa(line), Col: strconv.Itoa(col)}
}

func (p Struct) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	conf, result := env.(environment.Environment).GetStruct(p.Id, p.Line, p.Col)
	name := env.(environment.Environment).Control.Id
	if conf {
		gen.NewComentario("EJECUTANDO STRUCT "+p.Id, name, true, false, "")
		i := 0
		conf1 := true
		expresiones := arraylist.New()
		if result.Atributo.Len() == p.Lista.Len() {
			gen.NewComentario("EJECUTANDO ATRIBUTOS", name, true, false, "")
			for _, s := range result.Atributo.ToArray() {
				att1 := s.(interfaces.Atributos)
				att2 := p.Lista.GetValue(i).(interfaces.Atributo)
				gen.NewComentario("EJECUCION DE  ATRIBUTO "+att1.Id, name, true, false, "")
				exp := att2.Expr.Ejecutar(env, gen)
				if exp.Type == att1.Tipo.Tipo && att1.Id == att2.Id {
					expresiones.Add(exp)
				} else {
					conf1 = false
					break
				}
				i++
			}
			if conf1 {
				atributos := arraylist.New()
				tipo := interfaces.STRUCT
				i = 0
				pos_origin := gen.NewTemp()
				gen.NewAsignacion(pos_origin, "H", false, "", name, true, false, "")
				for _, s := range result.Atributo.ToArray() {
					valor := expresiones.GetValue(i).(interfaces.Value)
					gen.NewHeap("H", valor.Value, true, "ATRIBUTO "+s.(interfaces.Atributos).Id, name, true, false, "")
					gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
					//att := interfaces.AttSymbol{Id: s.(interfaces.Atributos).Id, Valor: valor}
					atributos.Add(s.(interfaces.Atributos))
					i++
				}
				retorno.Value = pos_origin
				retorno.Type = tipo
				retorno.Tipo2 = atributos
				retorno.IsTemp = false

			} else {
				env.(environment.Environment).NewError("LOS ATRIBUTOS NO CONCUERDAN EN LA DECLARACION", p.Line, p.Col)
			}
		} else {
			env.(environment.Environment).NewError("LA CANTIDAD DE ATRIBUTOS NO CONCUERDA ", p.Line, p.Col)
		}
	}
	return retorno
}
