package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"

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
	ambito := env.(environment.Environment).DevAmbito()
	len := p.size.Ejecutar(env, gen)
	valor := p.expresion.Ejecutar(env, gen)
	if len.Type == interfaces.NULL || valor.Type == interfaces.NULL {
		retorno.Type = interfaces.NULL
		return retorno
	}
	entrada := gen.NewLabel()
	salida := gen.NewLabel()
	contador := gen.NewTemp()
	inicio := gen.NewTemp()
	mesures := arraylist.New()
	mesures.Add(len.Value)
	gen.AddCodes(inicio+"=H;", ambito)
	dim := interfaces.Dimensions{Tipo: valor.Type, Dimensions: mesures}
	gen.AddCodes(contador+"=0;", ambito)
	gen.AddCodes(entrada+":", ambito)
	gen.AddCodes("if ("+contador+"=="+len.Value+") goto "+salida+";", ambito)
	gen.AddCodes("HEAP[(int)H]="+valor.Value+";", ambito)
	gen.AddCodes("H=H+1;", ambito)
	gen.Heap++
	gen.AddCodes(contador+"="+contador+"+1;", ambito)
	gen.AddCodes("goto "+entrada+";", ambito)
	gen.AddCodes(salida+":", ambito)
	gen.AddCodes("HEAP[(int)H]=-123456;", ambito)
	gen.AddCodes("H=H+1;", ambito)
	gen.Heap++
	retorno.Type = interfaces.ARRAY
	retorno.Tipo2 = arraylist.New()
	retorno.Tipo2.Add(dim)
	retorno.TrueLabel = len.Value
	fmt.Println("asadsdas")
	retorno.Value = inicio
	retorno.IsTemp = true
	return retorno
}
