package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
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
	fmt.Println(env.(environment.Environment).Control.Id)
	ex := gen.NewLabel()
	tmp := environment.NewEnvironment(env.(environment.Environment), f.Id, "", "", false, 1, ex)
	name := f.Id
	if env.(environment.Environment).SaveFunc(f.Line, f.Col, f.Id, newFunc) {
		gen.NewObligatorio("void "+f.Id+" (){", false, "", f.Id, true, false, f.Line)
		//list_pos:=arraylist.New()
		//gen.NewOperacion("P", "P", "+", "1", false, "", name, true, false, "")
		for _, s := range f.Params.ToArray() {
			tmp_pos := gen.NewTemp()
			tmp_val := gen.NewTemp()
			param := s.(interfaces.Parametros)
			stack := strconv.Itoa(tmp.Control.Stack)
			gen.NewOperacion(tmp_pos, "P", "+", stack, false, "", name, true, false, "")
			gen.NewCallStack(tmp_val, tmp_pos, false, "", name, true, false, "")
			//tmp.Control.Stack++
			//gen.NewOperacion("P", "P", "+", "1", false, "", name, true, false, "")
			tmp.Control.Stack++
			//list_pos.Add(tmp_val)
			if param.Tipo.Tipo == interfaces.ARRAY || param.Tipo.Tipo == interfaces.VECTOR || param.Tipo.Tipo == interfaces.STRUCT {
				tipo2 := arraylist.New()
				if param.Tipo.Tipo == interfaces.VECTOR {
					dim := createDimension(param.Tipo)
					tipo2.Add(dim)
					param.Tipo.Tipo2 = tipo2
				}

				sym := interfaces.Symbol{Id: param.Id, Tipo: param.Tipo, Posicion: tmp_val, Posicion2: tmp_val, Mutable: true}
				tmp.SaveVariable(f.Line, f.Col, param.Id, sym, param.Tipo)
			} else {
				gen.NewStack(tmp_pos, tmp_val, true, "DECLARACION DEL PARAMETRO "+param.Id, name, true, false, "")
				sym := interfaces.Symbol{Id: param.Id, Tipo: param.Tipo, Posicion: tmp_pos, Mutable: true}
				tmp.SaveVariable(f.Line, f.Col, param.Id, sym, param.Tipo)
			}

		}
		for _, s := range f.Statments.ToArray() {
			s.(interfaces.Instruction).Ejecutar(tmp, gen)
		}
		retorno_conf := gen.GetReturn(f.Id)
		if retorno_conf {
			if f.Tipo.Tipo == interfaces.NULL {
				env.(environment.Environment).NewError("NO SE ESPERABA UN RETORNO", f.Line, f.Col)
			}
		} else {
			/*if f.Tipo.Tipo != interfaces.NULL {
				env.(environment.Environment).NewError("SE ESPERABA UN RETORNO", f.Line, f.Col)
				gen.NewLabels(ex, false, "", name, true, true, "")
				gen.NewObligatorio("return;", false, "", f.Id, true, false, f.Line)
			} else {
				gen.NewObligatorio("return;", false, "", f.Id, true, false, f.Line)
			}*/
			gen.NewLabels(ex, false, "", name, true, true, "")
			gen.NewObligatorio("return;", false, "", f.Id, true, false, f.Line)
		}
		gen.NewObligatorio("}", false, "", f.Id, true, false, f.Line)
	}
	return retorno
}

func createDimension(t interfaces.TipoSimbolo) interfaces.Dimensions {
	tipo := t.Tipo2.GetValue(0).(interfaces.TipoExpresion)
	dim := interfaces.Dimensions{Tipo: tipo, Dimensions: t.Tipo2}
	/*fmt.Println("********")
	fmt.Println(dim.Dimensions.Len())*/
	return dim
}
