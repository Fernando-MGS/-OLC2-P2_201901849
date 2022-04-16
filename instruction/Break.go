package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"

	//"fmt"
	"strconv"
)

type Break struct {
	Retorno interfaces.Expresion
	Tipo    bool
	line    string
	col     string
}

func NewBreak(retorno interfaces.Expresion, tipo bool, line, col int) Break {
	brk := Break{retorno, tipo, strconv.Itoa(line), strconv.Itoa(col)}
	return brk
}

func (b Break) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var ret interfaces.Value
	ret.Type = interfaces.NULL
	ambito := env.(environment.Environment).DevAmbito()
	if env.(environment.Environment).Control.Ciclo {
		if b.Tipo {
			gen.AddCodes("//INICIO DEL BREAK", ambito)
			res := b.Retorno.Ejecutar(env, gen)
			//fmt.Println("EL RETORNO ES-BREAK ")
			ret = res
			/*fmt.Println(res)
			fmt.Println("FIN-BREAK ")*/
			code := "goto " + env.(environment.Environment).Control.Salida + ";"
			gen.AddCodes(code, ambito)
			gen.AddCodes("//FIN DEL BREAK", ambito)
			return res
		} else {
			code := "//BREAK\n" + "goto " + env.(environment.Environment).Control.Salida + ";\n//FIN DE BREAK\n"
			gen.AddCodes(code, ambito)
		}
	} else {
		env.(environment.Environment).NewError("NO SE PUEDE USAR UN CONTINUE FUERA DE UN CICLO", b.line, b.col)
	}

	//fmt.Println("El retorno 2 es ")
	return ret
}
