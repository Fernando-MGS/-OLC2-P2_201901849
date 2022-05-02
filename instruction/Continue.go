package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type Continue struct {
	extra string
	line  string
	col   string
}

func NewContinue(extra string, line, col int) Continue {
	brk := Continue{extra, strconv.Itoa(line), strconv.Itoa(col)}
	return brk
}

func (c Continue) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var ret interfaces.Value
	name := env.(environment.Environment).Control.Id
	//ambito := env.(environment.Environment).DevAmbito()
	if env.(environment.Environment).Control.Ciclo {
		//code := "//CONTINUE\n" + "goto " + env.(environment.Environment).Control.Entrada + ";\n//FIN DE CONTINUE\n"
		gen.NewComentario("INICIO DEL CONTINUE", name, true, false, c.line)
		gen.NewSalto(env.(environment.Environment).Control.Entrada, false, "", name, true, false, c.line)
		gen.NewComentario("FIN DEL CONTINUE", name, true, false, c.line)
		//gen.AddCodes(code, ambito)
	} else {
		env.(environment.Environment).NewError("NO SE PUEDE USAR UN CONTINUE FUERA DE UN CICLO", c.line, c.col)
	}
	//code:="goto "+env.(environment.Environment)
	ret.Type = interfaces.NULL
	return ret
}
