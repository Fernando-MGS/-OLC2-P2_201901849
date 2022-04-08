package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
	//"fmt"
)

type CallVariable struct {
	Id string
}

func NewCallVariable(id string) CallVariable {
	exp := CallVariable{Id: id}
	return exp
}

func (p CallVariable) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value { // GUARDAR EL TEMPORAL PARA ELIMINARLO EN LA OPTIMIZACION
	//fmt.Println("aceso")
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	result := env.(environment.Environment).GetVariable(p.Id)
	if result.Tipo.Tipo != interfaces.NULL {
		newTemp := gen.NewTemp()
		if interfaces.INTEGER == result.Tipo.Tipo {
			value := "STACK[" + strconv.Itoa(result.Posicion) + "]"
			gen.AddExpression(newTemp, "", "", value)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.FLOAT == result.Tipo.Tipo {
			value := "STACK[" + strconv.Itoa(result.Posicion) + "]"
			gen.AddExpression(newTemp, "", "", value)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		} else if interfaces.USIZE == result.Tipo.Tipo {
			value := "STACK[" + strconv.Itoa(result.Posicion) + "]"
			gen.AddExpression(newTemp, "", "", value)
			retorno = interfaces.Value{Value: newTemp, IsTemp: true, Type: result.Tipo.Tipo, Tipo2: result.Tipo.Tipo2}
		}
	}
	return retorno
}
