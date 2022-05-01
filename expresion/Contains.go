package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Contains struct {
	Acceso interfaces.Expresion
	Value  interfaces.Expresion
	Line   string
	Col    string
}

func NewContains(id, value interfaces.Expresion, line, col int) Contains {
	return Contains{id, value, strconv.Itoa(line), strconv.Itoa(col)}
}

func (c Contains) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	access := c.Acceso.Ejecutar(env, gen)
	value := c.Value.Ejecutar(env, gen)
	ambito := env.(environment.Environment).DevAmbito()
	if access.Type == interfaces.ARRAY || access.Type == interfaces.VECTOR && access.Type != interfaces.NULL {
		dimension := access.Tipo2.GetValue(0).(interfaces.Dimensions)
		if value.Type == dimension.Tipo {
			p_valor := gen.NewTemp()
			p_tipo := gen.NewTemp()
			p_pos := gen.NewTemp()
			gen.AddCodes(p_valor+"=P+1;", ambito)
			gen.AddCodes(p_tipo+"=P+2;", ambito)
			gen.AddCodes(p_pos+"=P+3;", ambito)
			gen.AddCodes("STACK[(int)"+p_valor+"]="+value.Value+";", ambito)
			if value.Type == interfaces.STRING || value.Type == interfaces.STR {
				gen.AddCodes("STACK[(int)"+p_tipo+"]=0;", ambito)
			} else {
				gen.AddCodes("STACK[(int)"+p_tipo+"]=1;", ambito)
			}
			gen.AddCodes("STACK[(int)"+p_pos+"]="+access.Value+";", ambito)
			gen.AddFuncExtra("CONTAINSVECTOR")
			gen.AddCodes("proc_ContainsVector();", ambito)
			true_l := gen.NewLabel()
			false_l := gen.NewLabel()
			resultado := gen.NewTemp()
			gen.AddCodes(resultado+"=STACK[(int)P];", ambito)
			gen.AddCodes("if("+resultado+"==1) goto "+true_l+";", ambito)
			gen.AddCodes("goto "+false_l+";", ambito)
			retorno.Type = interfaces.BOOLEAN
			retorno.IsTemp = false
			retorno.TrueLabel = true_l
			retorno.FalseLabel = false_l
		} else {
			env.(environment.Environment).NewError("EL TIPO DATO QUE SE DESEA ENCONTRA NO CONCUERDA CON EL TIPO DEL VECTOR", c.Line, c.Col)
		}
	} else {
		env.(environment.Environment).NewError("EL ELEMENTO AL QUE SE QUIERE ACCEDER NO ES UN VECTOR O UN ARRAY", c.Line, c.Col)
	}
	return retorno
}
