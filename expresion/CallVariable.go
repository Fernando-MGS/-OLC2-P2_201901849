package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
	//"fmt"
)

type CallVariable struct {
	Id    string
	Linea string
	Col   string
}

func NewCallVariable(id string, lin, col int) CallVariable {
	exp := CallVariable{Id: id, Linea: strconv.Itoa(lin), Col: strconv.Itoa(col)}
	return exp
}

func (p CallVariable) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value { // GUARDAR EL TEMPORAL PARA ELIMINARLO EN LA OPTIMIZACION
	//fmt.Println("aceso")
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	name := env.(environment.Environment).Control.Id
	result := env.(environment.Environment).GetVariable(p.Id, p.Linea, p.Col)
	//ambito := env.(environment.Environment).DevAmbito()
	if result.Tipo.Tipo != interfaces.NULL {
		//gen.AddCodes("//INICIO DE LLAMADA"+p.Id+"", ambito)
		gen.NewComentario("//INICIO DE LLAMADA"+p.Id, name, true, false, p.Linea)
		newTemp := gen.NewTemp()
		if interfaces.INTEGER == result.Tipo.Tipo || interfaces.FLOAT == result.Tipo.Tipo {
			/*value := "STACK[(int)" + result.Posicion + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)*/
			gen.NewCallStack(newTemp, result.Posicion, false, "", name, true, true, p.Linea)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.FLOAT == result.Tipo.Tipo {
			/*value := "STACK[(int)" + result.Posicion + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)*/
			gen.NewCallStack(newTemp, result.Posicion, false, "", name, true, true, p.Linea)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.USIZE == result.Tipo.Tipo {
			/*value := "STACK[(int)" + result.Posicion + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)*/
			gen.NewCallStack(newTemp, result.Posicion, false, "", name, true, true, p.Linea)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.CHAR == result.Tipo.Tipo {
			/*value := "STACK[(int)" + result.Posicion + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)*/
			gen.NewCallStack(newTemp, result.Posicion, false, "", name, true, true, p.Linea)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.BOOLEAN == result.Tipo.Tipo {
			/*value := "STACK[(int)" + result.Posicion + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)*/
			gen.NewCallStack(newTemp, result.Posicion, false, "", name, true, true, p.Linea)
			l1 := gen.NewLabel()
			l2 := gen.NewLabel()
			//gen.AddCodes("if("+newTemp+"=="+"1) goto "+l1+";", ambito)
			gen.NewIf(newTemp, "==", "1", l1, false, "", name, true, true, p.Linea)
			//gen.AddCodes("goto "+l2+";", ambito)
			gen.NewSalto(l2, false, "", name, true, true, p.Linea)
			/*conf := gen.GetConf()
			if conf == 0 {
				l2 = gen.NewLabel()
				value := "if(" + newTemp + "==" + "1) goto " + l1 + ";\n"
				value += "goto " + l2 + ";\n"
				gen.AddCodes(value, ambito)
				gen.AddTempBool(l1, l2)
				gen.SetConf()
			} else {
				l1 = gen.NewLabel()
				value := "if(" + newTemp + "==" + "0) goto "
				gen.AddTempBool(l1, value)
			}*/
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2, TrueLabel: l1, FalseLabel: l2}
		} else if interfaces.STR == result.Tipo.Tipo || result.Tipo.Tipo == interfaces.STRING {
			/*value := "STACK[(int)" + result.Posicion + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)*/
			gen.NewCallStack(newTemp, result.Posicion, false, "", name, true, true, p.Linea)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.ARRAY == result.Tipo.Tipo {
			retorno = interfaces.Value{Value: result.Posicion2, IsTemp: true, Type: interfaces.ARRAY, Tipo2: result.Tipo.Tipo2, TrueLabel: result.Posicion2}
			return retorno
		} else if interfaces.VECTOR == result.Tipo.Tipo {
			retorno = interfaces.Value{Value: result.Posicion2, IsTemp: true, Type: interfaces.VECTOR, Tipo2: result.Tipo.Tipo2, TrueLabel: result.Posicion2}
			return retorno
		}
	}
	return retorno
}
