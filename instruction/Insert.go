package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Insert struct {
	Id    interfaces.Expresion
	Valor interfaces.Expresion
	Index interfaces.Expresion
	Line  string
	Col   string
}

func NewInsert(id, index interfaces.Expresion, valor interfaces.Expresion, line, col int) Insert {
	instr := Insert{Id: id, Valor: valor, Index: index, Line: strconv.Itoa(line), Col: strconv.Itoa(col)}
	return instr
}

func (p Insert) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	access := p.Id.Ejecutar(env, gen)
	value := p.Valor.Ejecutar(env, gen)
	indexs := p.Index.Ejecutar(env, gen)
	//ambito := env.(environment.Environment).DevAmbito()
	name := env.(environment.Environment).Control.Id
	if access.Type == interfaces.VECTOR {
		conf := false
		//fmt.Println(value)
		dimension_acc := access.Tipo2.GetValue(0).(interfaces.Dimensions)
		//fmt.Println(indexs)
		if indexs.Type == interfaces.INTEGER || indexs.Type == interfaces.USIZE {
			dimension := access.Tipo2.GetValue(0).(interfaces.Dimensions)
			if dimension.Dimensions.Len() > 1 && value.Type == interfaces.VECTOR {
				if dimension.Tipo == dimension_acc.Tipo && dimension.Dimensions.Len() == dimension_acc.Dimensions.Len() {
					conf = true
				} else {
					env.(environment.Environment).NewError("NO SE PUEDE INSERTAR EL VALOR PORQUE LOS TIPOS NO CONCUERDAN", p.Line, p.Col)
				}
			} else if value.Type == dimension.Tipo {
				conf = true
			} else {
				env.(environment.Environment).NewError("NO SE PUEDE INSERTAR EL VALOR PORQUE LOS TIPOS NO CONCUERDAN", p.Line, p.Col)
			}
			if conf {
				valor := gen.NewTemp()
				index := gen.NewTemp()
				pos := gen.NewTemp()
				gen.AddFuncExtra("INSERTVECTOR")
				//gen.AddCodes(valor+"=P+1;", ambito)
				gen.NewOperacion(valor, "P", "+", "1", false, "", name, true, true, p.Line)
				//gen.AddCodes(index+"=P+2;", ambito)
				gen.NewOperacion(index, "P", "+", "2", false, "", name, true, true, p.Line)
				//gen.AddCodes(pos+"=P+3;", ambito)
				gen.NewOperacion(pos, "P", "+", "3", false, "", name, true, true, p.Line)
				//gen.AddCodes("STACK[(int)"+valor+"]="+value.Value+";", ambito)
				gen.NewStack(valor, value.Value, false, "DECLARACION DE CONTADOR ", name, true, true, p.Line)
				//gen.AddCodes("STACK[(int)"+index+"]="+indexs.Value+";", ambito)
				gen.NewStack(index, indexs.Value, false, "", name, true, false, p.Line)
				//gen.AddCodes("STACK[(int)"+pos+"]="+access.Value+";", ambito)
				gen.NewStack(pos, access.Value, false, "DECLARACION DE CONTADOR ", name, true, true, p.Line)
				//gen.AddCodes("proc_InsertVector();", ambito)
				gen.NewLlamada("proc_InsertVector", false, "", name, true, false, p.Line)
			}
		}
	} else {
		env.(environment.Environment).NewError("SE ESPERABA UN VECTOR", p.Line, p.Col)
	}
	return retorno
}
