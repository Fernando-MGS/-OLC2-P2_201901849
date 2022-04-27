package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"

	"github.com/colegno/arraylist"
)

type VectorA struct {
	Exp1 interfaces.Expresion
	Exp2 interfaces.Expresion
}

func NewVectorA(e1, e2 interfaces.Expresion) VectorA {
	expr := VectorA{e1, e2}
	return expr
}

func (v VectorA) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	var retorno interfaces.Value
	retorno.Type = interfaces.VECTOR
	ambito := env.(environment.Environment).DevAmbito()
	value := v.Exp1.Ejecutar(env, gen)
	long := v.Exp2.Ejecutar(env, gen)
	iterador := gen.NewTemp()
	//contador
	gen.AddCodes("//CREANDO VECTOR CON VALORES REPETIDOS", ambito)
	entrada := gen.NewLabel()
	salida := gen.NewLabel()
	//pos := gen.NewTemp()
	largo := gen.NewTemp()
	gen.AddCodes(largo+"=H;//POS DE INICIO VECTOR A", ambito)
	gen.AddCodes("HEAP[(int)H]="+long.Value+";", ambito)
	gen.AddCodes("H++;", ambito)
	gen.AddCodes(iterador+"=0;", ambito)
	//gen.AddCodes(pos+"=H;", ambito)
	gen.AddCodes(entrada+":", ambito)
	gen.AddCodes("if ("+iterador+"=="+long.Value+") goto "+salida+";", ambito)
	gen.AddCodes("HEAP[(int)H]="+value.Value+";", ambito)
	gen.AddCodes("H=H+1;", ambito)
	gen.AddCodes(iterador+"="+iterador+"+1;", ambito)
	gen.AddCodes("goto "+entrada+";", ambito)
	gen.AddCodes(salida+":", ambito)
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
