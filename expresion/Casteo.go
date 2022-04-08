package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"time"
)

type Cast struct {
	Arg         interfaces.Expresion
	Tipo_casteo interfaces.TipoExpresion
}

func NewCast(args interfaces.Expresion, tipo interfaces.TipoExpresion) Cast {
	//fmt.Println("castera1")
	expr := Cast{args, tipo}
	return expr
}

func (c Cast) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	argumento := c.Arg.Ejecutar(env, gen)
	fmt.Println(argumento)
	fmt.Println(c.Tipo_casteo)
	if argumento.Type == interfaces.INTEGER && c.Tipo_casteo == interfaces.FLOAT { // de entero a float
		//value:=argumento.
		//argumento.Value="(int)"+argumento.Value
		argumento.Type = interfaces.FLOAT
		return argumento
	} else if argumento.Type == interfaces.FLOAT && c.Tipo_casteo == interfaces.INTEGER { //de float a entero
		//value:=argumento.
		argumento.Value = "(int)" + argumento.Value
		argumento.Type = interfaces.INTEGER
		fmt.Println(argumento)
		return argumento
	} else if argumento.Type == interfaces.INTEGER && c.Tipo_casteo == interfaces.USIZE { //de entero a usize
		//value:=argumento.
		//argumento.Value="(int)"+argumento.Value
		argumento.Type = interfaces.USIZE
		return argumento
	} else if argumento.Type == interfaces.USIZE && c.Tipo_casteo == interfaces.INTEGER { //de usize a entero
		//argumento.Value="(int)"+argumento.Value
		argumento.Type = interfaces.INTEGER
		return argumento
	} else if argumento.Type == interfaces.FLOAT && c.Tipo_casteo == interfaces.USIZE { //de float a usize
		argumento.Value = "(int)" + argumento.Value
		argumento.Type = interfaces.USIZE
		return argumento
	} else if argumento.Type == interfaces.USIZE && c.Tipo_casteo == interfaces.FLOAT { //usize a float
		//argumento.Value="(int)"+argumento.Value
		argumento.Type = interfaces.FLOAT
		return argumento
	}
	t := time.Now()
	fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
		t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
	err := interfaces.Errores{Line: "0", Col: "0", Mess: "EL CASTEO NO ES VALIDO ENTRE ESTOS TIPOS", Fecha: fecha}
	env.(environment.Environment).Errores(err)
	fmt.Println(err)
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	return retorno
}
