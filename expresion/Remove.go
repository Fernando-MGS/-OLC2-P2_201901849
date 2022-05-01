package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Remove struct {
	Id    interfaces.Expresion
	Index interfaces.Expresion
	Line  string
	Col   string
}

func NewRemove(id, index interfaces.Expresion, line, col int) Remove {
	return Remove{id, index, strconv.Itoa(line), strconv.Itoa(col)}
}

func (p Remove) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	ambito := env.(environment.Environment).DevAmbito()
	access := p.Id.Ejecutar(env, gen)
	indexs := p.Index.Ejecutar(env, gen)
	if access.Type == interfaces.VECTOR {
		if indexs.Type == interfaces.INTEGER || indexs.Type == interfaces.USIZE {
			dimension := access.Tipo2.GetValue(0).(interfaces.Dimensions)
			index := gen.NewTemp()
			index_pos := gen.NewTemp()
			new_value := gen.NewTemp()
			gen.AddFuncExtra("REMOVEVECTOR")
			gen.AddCodes(index+"=P+1;", ambito)
			gen.AddCodes(index_pos+"=P+2;", ambito)
			gen.AddCodes("STACK[(int)"+index+"]="+indexs.Value+";", ambito)
			gen.AddCodes("STACK[(int)"+index_pos+"]="+access.Value+";", ambito)
			gen.AddCodes("proc_RemoveVector();", ambito)
			gen.AddFuncExtra("REMOVEVECTOR")
			gen.AddCodes(new_value+"=STACK[(int)P];", ambito)
			retorno.Value = new_value
			retorno.Type = dimension.Tipo
			retorno.IsTemp = true
		}

	} else {
		retorno.Type = interfaces.NULL
		env.(environment.Environment).NewError("EL ELEMENTO AL QUE SE QUIERE ACCEDER NO ES UN VECTOR", p.Line, p.Col)
	}
	return retorno
}
