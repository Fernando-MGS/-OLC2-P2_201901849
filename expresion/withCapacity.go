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
	ambito := env.(environment.Environment).DevAmbito()
	gen.AddCodes("//CREANDO VECTOR CON CAPACIDAD", ambito)
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
	gen.AddCodes(pos+"=H;", ambito)
	gen.AddCodes("HEAP[(int)H]="+largo.Value+";", ambito)
	gen.AddCodes("H=H+1;", ambito)
	gen.AddCodes(entrada+":", ambito)
	gen.AddCodes("if ("+iterador+"=="+largo.Value+") goto "+salida+";", ambito)
	gen.AddCodes("HEAP[(int)H]=-123456;", ambito)
	gen.AddCodes("H=H+1;", ambito)
	gen.AddCodes(iterador+"="+iterador+"+1;", ambito)
	gen.AddCodes("goto "+entrada+";", ambito)
	gen.AddCodes(salida+":", ambito)
	retorno.Type = interfaces.VECTOR
	retorno.Value = pos
	return retorno
}
