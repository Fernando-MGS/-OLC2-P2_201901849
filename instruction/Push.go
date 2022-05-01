package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
)

type Push struct {
	Acceso interfaces.Expresion
	Valor  interfaces.Expresion
	Line   string
	Col    string
}

func NewPush(access, value interfaces.Expresion, line, col int) Push {
	return Push{access, value, strconv.Itoa(line), strconv.Itoa(col)}
}

func (p Push) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	ambito := env.(environment.Environment).DevAmbito()
	access := p.Acceso.Ejecutar(env, gen)
	value := p.Valor.Ejecutar(env, gen)
	fmt.Println(access)
	if access.Type != interfaces.NULL && value.Type != interfaces.NULL {
		//pos:=gen.NewTemp()
		dim := access.Tipo2.GetValue(0).(interfaces.Dimensions)

		if dim.Tipo == value.Type {
			new_pos := gen.NewTemp()
			stack_1 := gen.NewTemp()
			stack_2 := gen.NewTemp()
			gen.AddCodes(stack_1+"=P+1;", ambito)
			gen.AddCodes(stack_2+"=P+2;", ambito)
			gen.AddCodes("STACK[(int)"+stack_1+"]="+value.Value+";", ambito)
			gen.AddCodes("STACK[(int)"+stack_2+"]="+access.Value+";", ambito)
			gen.AddFuncExtra("PUSHVECTOR")
			gen.AddCodes("proc_vectorPush();", ambito)
			gen.AddCodes(new_pos+"=STACK[(int)P];", ambito)
			gen.AddCodes(access.Value+"="+new_pos+";", ambito)
			//gen.AddCodes("", ambito)
		} else {
			env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION VALIDA O EL ELEMENTO NO ES UN VECTOR", p.Line, p.Col)
		}
	} else {
		env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION VALIDA", p.Line, p.Col)
	}
	return retorno
}
