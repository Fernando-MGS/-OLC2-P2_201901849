package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"

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
	name := env.(environment.Environment).Control.Id
	retorno.Type = interfaces.NULL
	entrada := gen.NewLabel()
	salida := gen.NewLabel()
	//ambito := env.(environment.Environment).DevAmbito()
	///gen.AddCodes("//INICIO DE LOOP", ambito)
	gen.NewComentario("INICIO DE LOOP", name, true, false, "")
	//gen.AddCodes(entrada+":", ambito)
	gen.NewLabels(entrada, false, "", name, true, true, "")
	stack := env.(environment.Environment).Control.Stack
	tmpEnv := environment.NewEnvironment(env.(environment.Environment), env.(environment.Environment).Control.Id, entrada, salida, true, stack)
	for _, s := range l.Cuerpo.ToArray() {
		res := s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
		/*fmt.Println("//////LOOP")
		fmt.Println("SE ENCONTRO UN RETORNO")*/

		ret := res.(interfaces.Value)
		//fmt.Println(ret)
		if ret.Type != interfaces.NULL {
			/*fmt.Println("SE ENCONTRO UN RETORNO2")
			fmt.Println(res)*/
			retorno = ret
		}
		//fmt.Println("//////ENDLOOP")
	}
	//gen.AddCodes("goto "+entrada+";", ambito)
	gen.NewSalto(entrada, false, "", name, true, false, "")
	//gen.AddCodes(salida+":", ambito)
	gen.NewLabels(salida, false, "", name, true, true, "")
	//gen.AddCodes("//FIN DE LOOP", ambito)
	gen.NewComentario("FIN DE LOOP", name, true, false, "")
	return retorno
}
