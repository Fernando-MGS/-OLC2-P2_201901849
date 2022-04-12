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
	result := env.(environment.Environment).GetVariable(p.Id, p.Linea, p.Col)
	ambito := true
	if env.(environment.Environment).Control.Id == "GLOBAL" || env.(environment.Environment).Control.Id == "main" {
		ambito = false
	}
	if result.Tipo.Tipo != interfaces.NULL {
		newTemp := gen.NewTemp()
		if interfaces.INTEGER == result.Tipo.Tipo || interfaces.FLOAT == result.Tipo.Tipo {
			value := "STACK[" + strconv.Itoa(result.Posicion) + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.FLOAT == result.Tipo.Tipo {
			value := "STACK[" + strconv.Itoa(result.Posicion) + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.USIZE == result.Tipo.Tipo {
			value := "STACK[" + strconv.Itoa(result.Posicion) + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.CHAR == result.Tipo.Tipo {
			value := "STACK[" + strconv.Itoa(result.Posicion) + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.BOOLEAN == result.Tipo.Tipo {
			value := "STACK[" + strconv.Itoa(result.Posicion) + "]"
			gen.AddExpression(newTemp, "", "", value, ambito)
			l1 := gen.NewLabel()
			l2 := ""
			conf := gen.GetConf()
			if conf == 0 {
				l2 = gen.NewLabel()
				value := "if(" + newTemp + "==" + "1) goto " + l1 + ";\n"
				value += "goto " + l2 + ";\n"
				gen.AddCodes(value, ambito)
				gen.AddTempBool(l1, l2)
				gen.SetConf()
			} else {
				l1 = gen.NewLabel()
				value := "if(" + newTemp + "==" + "0) goto " + l1 + ";\n"
				gen.AddTempBool(l1, value)
			}
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		}
	}
	return retorno
}
