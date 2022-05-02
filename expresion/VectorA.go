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
	//ambito := env.(environment.Environment).DevAmbito()
	value := v.Exp1.Ejecutar(env, gen)
	long := v.Exp2.Ejecutar(env, gen)
	iterador := gen.NewTemp()
	name := env.(environment.Environment).Control.Id
	//contador
	//gen.AddCodes("//CREANDO VECTOR CON VALORES REPETIDOS", ambito)
	gen.NewComentario("CREANDO VECTOR CON VALORES REPETIDOS", name, true, false, "")
	entrada := gen.NewLabel()
	salida := gen.NewLabel()
	//pos := gen.NewTemp()
	largo := gen.NewTemp()
	//gen.AddCodes(largo+"=H;//POS DE INICIO VECTOR A", ambito)
	gen.NewAsignacion(largo, "H", true, "POS DE INICIO DE VECTOR A", name, true, false, "")
	//gen.AddCodes("HEAP[(int)H]="+long.Value+";", ambito)
	gen.NewHeap("H", long.Value, false, "", name, true, false, "")
	//gen.AddCodes("H++;", ambito)
	gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
	//gen.AddCodes(iterador+"=0;", ambito)
	gen.NewAsignacion(iterador, "0", false, "", name, true, false, "")
	//gen.AddCodes(pos+"=H;", ambito)
	//gen.AddCodes(entrada+":", ambito)
	gen.NewLabels(entrada, false, "", name, true, false, "")
	//gen.AddCodes("if ("+iterador+"=="+long.Value+") goto "+salida+";", ambito)
	gen.NewIf(iterador, "==", long.Value, salida, false, "", name, true, false, "")
	//gen.AddCodes("HEAP[(int)H]="+value.Value+";", ambito)
	gen.NewHeap("H", value.Value, false, "", name, true, false, "")
	//gen.AddCodes("H=H+1;", ambito)
	gen.NewOperacion("H", "H", "+", "1", false, "", name, true, false, "")
	//gen.AddCodes(iterador+"="+iterador+"+1;", ambito)
	gen.NewOperacion(iterador, iterador, "+", "1", false, "", name, true, true, "")
	//gen.AddCodes("goto "+entrada+";", ambito)
	gen.NewSalto(entrada, false, "", name, true, false, "")
	//gen.AddCodes(salida+":", ambito)
	gen.NewLabels(salida, false, "", name, true, true, "")
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
