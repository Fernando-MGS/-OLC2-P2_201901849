package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"

	"github.com/colegno/arraylist"
)

type SetData struct {
	expresion interfaces.Expresion
	size      interfaces.Expresion
}

func NewArray2(expr, tam interfaces.Expresion) SetData {
	exp := SetData{expresion: expr, size: tam}
	return exp
}

func (p SetData) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.ARRAY
	//ambito := env.(environment.Environment).DevAmbito()
	name := env.(environment.Environment).Control.Id
	value := p.expresion.Ejecutar(env, gen)
	long := p.size.Ejecutar(env, gen)
	iterador := gen.NewTemp()
	//contador
	//gen.AddCodes("//CREANDO ARRAY CON VALORES REPETIDOS", ambito)
	gen.NewComentario("CREANDO ARRAY CON VALORES REPETIDOS", name, true, false, "")
	entrada := gen.NewLabel()
	salida := gen.NewLabel()
	//pos := gen.NewTemp()
	largo := gen.NewTemp()
	//gen.AddCodes(largo+"=H;//POS DE INICIO ARRAY A", ambito)
	gen.NewAsignacion(largo, "H", false, "", name, true, false, "")
	//gen.AddCodes("HEAP[(int)H]="+long.Value+";", ambito)
	gen.NewHeap("H", value.Value, false, "", name, true, false, "")
	//gen.AddCodes("H++;", ambito)
	gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
	//gen.AddCodes(iterador+"=0;", ambito)
	gen.NewAsignacion(iterador, "0", false, "", name, true, true, "")
	//gen.AddCodes(pos+"=H;", ambito)
	//gen.AddCodes(entrada+":", ambito)
	gen.NewLabels(entrada, false, "", name, true, false, "")
	//gen.AddCodes("if ("+iterador+"=="+long.Value+") goto "+salida+";", ambito)
	gen.NewIf(iterador, "==", long.Value, salida, false, "", name, true, true, "")
	//gen.AddCodes("HEAP[(int)H]="+value.Value+";", ambito)
	gen.NewHeap("H", value.Value, false, "", name, true, false, "")
	//gen.AddCodes("H=H+1;", ambito)
	gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
	//gen.AddCodes(iterador+"="+iterador+"+1;", ambito)
	gen.NewOperacion(iterador, iterador, "+", "1", false, "", name, true, false, "")
	//gen.AddCodes("goto "+entrada+";", ambito)
	gen.NewSalto(entrada, false, "", name, true, false, "")
	//gen.AddCodes(salida+":", ambito)
	gen.NewLabels(salida, false, "", name, true, false, "")
	//retorno.Type=value.Type
	retorno.Tipo2 = arraylist.New()
	longs := arraylist.New()
	longs.Add(long.Value)
	dimensions := interfaces.Dimensions{Tipo: value.Type, Dimensions: longs}
	retorno.Tipo2.Add(dimensions)
	retorno.TrueLabel = largo
	retorno.Value = largo
	//retorno.IsTemp=true
	return retorno
}
