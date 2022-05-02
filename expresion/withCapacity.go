package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"

	"github.com/colegno/arraylist"
)

type W_Capacity struct {
	Len interfaces.Expresion
}

func New_CapacityV(largo interfaces.Expresion) W_Capacity {
	return W_Capacity{largo}
}

func (p W_Capacity) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	name := env.(environment.Environment).Control.Id
	//ambito := env.(environment.Environment).DevAmbito()
	//gen.AddCodes("//CREANDO VECTOR CON CAPACIDAD", ambito)
	gen.NewComentario("CREANDO VECTOR CON CAPACIDAD FIJADA", name, true, false, "")
	retorno.Tipo2 = arraylist.New()
	dim_long := arraylist.New()
	dim_long.Add(interfaces.VOID)
	dimension := interfaces.Dimensions{Tipo: interfaces.VOID, Dimensions: dim_long}
	retorno.Tipo2.Add(dimension)
	largo := p.Len.Ejecutar(env, gen)
	pos := gen.NewTemp()
	iterador := gen.NewTemp()
	entrada := gen.NewTemp()
	salida := gen.NewTemp()
	//gen.AddCodes(pos+"=H;", ambito)
	gen.NewAsignacion(pos, "H", false, "", name, true, false, "")
	//gen.AddCodes("HEAP[(int)H]="+largo.Value+";", ambito)
	gen.NewHeap("H", largo.Value, false, "", name, true, false, "")
	//gen.AddCodes("H=H+1;", ambito)
	gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
	//gen.AddCodes(entrada+":", ambito)
	gen.NewLabels(entrada, false, "", name, true, false, "")
	//gen.AddCodes("if ("+iterador+"=="+largo.Value+") goto "+salida+";", ambito)
	gen.NewIf(iterador, "==", largo.Value, salida, false, "", name, true, false, "")
	//gen.AddCodes("HEAP[(int)H]=-123456;", ambito)
	gen.NewHeap("H", "-123456", false, "", name, true, false, "")
	//gen.AddCodes("H=H+1;", ambito)
	gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
	//gen.AddCodes(iterador+"="+iterador+"+1;", ambito)
	gen.NewOperacion(iterador, iterador, "+", "1", false, "", name, true, false, "")
	//gen.AddCodes("goto "+entrada+";", ambito)
	gen.NewSalto(entrada, false, "", name, true, false, "")
	//gen.AddCodes(salida+":", ambito)
	gen.NewLabels(salida, false, "", name, true, false, "")
	retorno.Type = interfaces.VECTOR
	retorno.Value = pos
	return retorno
}
