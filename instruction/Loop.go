package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type Loop struct {
	Cuerpo *arrayList.List
}

func NewLoop(body *arrayList.List) Loop {
	//fmt.Println("ENTRO A LOOP1")
	instr := Loop{body}
	return instr
}

func (l Loop) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	//fmt.Println("ENTRO A LOOP")
	retorno.Type = interfaces.NULL
	entrada := gen.NewLabel()
	salida := gen.NewLabel()
	ambito := env.(environment.Environment).DevAmbito()
	gen.AddCodes("//INICIO DE LOOP", ambito)
	gen.AddCodes(entrada+":", ambito)
	tmpEnv := environment.NewEnvironment(env.(environment.Environment), env.(environment.Environment).Control.Id, entrada, salida, true)
	for _, s := range l.Cuerpo.ToArray() {
		res := s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
		/*fmt.Println("//////LOOP")
		fmt.Println("SE ENCONTRO UN RETORNO")*/

		ret := res.(interfaces.Value)
		fmt.Println(ret)
		if ret.Type != interfaces.NULL {
			/*fmt.Println("SE ENCONTRO UN RETORNO2")
			fmt.Println(res)*/
			retorno = ret
		}
		//fmt.Println("//////ENDLOOP")
	}
	gen.AddCodes("goto "+entrada+";", ambito)
	gen.AddCodes(salida+":", ambito)
	gen.AddCodes("//FIN DE LOOP", ambito)
	return retorno
}
