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
	//ambito := env.(environment.Environment).DevAmbito()
	name := env.(environment.Environment).Control.Id
	if access.Type == interfaces.ARRAY || access.Type == interfaces.VECTOR && access.Type != interfaces.NULL {
		dimension := access.Tipo2.GetValue(0).(interfaces.Dimensions)
		if value.Type == dimension.Tipo {
			p_valor := gen.NewTemp()
			p_tipo := gen.NewTemp()
			p_pos := gen.NewTemp()
			//gen.AddCodes(p_valor+"=P+1;", ambito)
			gen.NewOperacion(p_valor, "P", "+", "1", false, "", name, true, true, c.Line)
			//gen.AddCodes(p_tipo+"=P+2;", ambito)
			gen.NewOperacion(p_tipo, "P", "+", "2", false, "", name, true, true, c.Line)
			//gen.AddCodes(p_pos+"=P+3;", ambito)
			gen.NewOperacion(p_pos, "P", "+", "3", false, "", name, true, true, c.Line)
			//gen.AddCodes("STACK[(int)"+p_valor+"]="+value.Value+";", ambito)
			gen.NewStack(p_valor, value.Value, false, "", name, true, true, c.Line)
			if value.Type == interfaces.STRING || value.Type == interfaces.STR {
				//gen.AddCodes("STACK[(int)"+p_tipo+"]=0;", ambito)
				gen.NewStack(p_tipo, "0", false, "", name, true, true, c.Line)
			} else {
				//gen.AddCodes("STACK[(int)"+p_tipo+"]=1;", ambito)
				gen.NewStack(p_tipo, "1", false, "", name, true, true, c.Line)
			}
			//gen.AddCodes("STACK[(int)"+p_pos+"]="+access.Value+";", ambito)
			gen.NewStack(p_pos, access.Value, false, "", name, true, true, c.Line)
			gen.AddFuncExtra("CONTAINSVECTOR")
			//gen.AddCodes("proc_ContainsVector();", ambito)
			gen.NewLlamada("proc_ContainsVector", false, "", name, true, true, c.Line)
			true_l := gen.NewLabel()
			false_l := gen.NewLabel()
			resultado := gen.NewTemp()
			///gen.AddCodes(resultado+"=STACK[(int)P];", ambito)
			gen.NewCallStack(resultado, "P", false, "", name, true, true, c.Line)
			//gen.AddCodes("if("+resultado+"==1) goto "+true_l+";", ambito)
			gen.NewIf(resultado, "==", "1", true_l, false, "", name, true, true, c.Line)
			//gen.AddCodes("goto "+false_l+";", ambito)
			gen.NewSalto(false_l, false, "", name, true, false, c.Line)
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
