package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"strconv"
)

type ModArray struct {
	id     string
	Access interfaces.Expresion
	Value  interfaces.Expresion
	Line   string
	Col    string
}

func NewModArray(id string, index, val interfaces.Expresion, line, col int) ModArray {
	instr := ModArray{id, index, val, strconv.Itoa(line), strconv.Itoa(col)}
	return instr
}

func (m ModArray) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	ambito := env.(environment.Environment).DevAmbito()
	bounds := m.Access.Ejecutar(env, gen)
	value := m.Value.Ejecutar(env, gen)
	if bounds.Type == interfaces.NULL || value.Type == interfaces.NULL {
		return retorno
	}
	gen.AddCodes("//INICIANDO LA ASIGNACION DE "+m.id, ambito)
	if bounds.Type == value.Type {
		if bounds.Type == interfaces.ARRAY {
			entrada := gen.NewLabel()
			salida := gen.NewLabel()
			//gen.AddCodes("if("+bounds.TrueLabel+") ",ambito)
			iterador := gen.NewTemp()
			contador := gen.NewTemp()
			contador2 := gen.NewTemp()
			aux := gen.NewTemp()
			gen.AddCodes(iterador+"=0;//iterador", ambito)
			gen.AddCodes(contador+"="+bounds.Value+";//contador1", ambito)
			gen.AddCodes(contador2+"="+value.Value+";//contador2", ambito)
			gen.AddCodes(entrada+":", ambito)
			gen.AddCodes("if("+iterador+"=="+bounds.TrueLabel+") goto "+salida+";", ambito)
			gen.AddCodes(aux+"=HEAP[(int)"+contador2+"];", ambito)
			gen.AddCodes("HEAP[(int)"+contador+"]="+aux+";", ambito)
			gen.AddCodes(contador+"="+contador+"+1;", ambito)
			gen.AddCodes(contador2+"="+contador2+"+1;", ambito)
			gen.AddCodes(iterador+"="+iterador+"+1;", ambito)
			gen.AddCodes("goto "+entrada+";", ambito)
			gen.AddCodes(salida+":", ambito)
		} else {
			gen.AddCodes("HEAP[(int)"+bounds.TrueLabel+"]="+value.Value+";", ambito)
		}
	} else {
		env.(environment.Environment).NewError("LOS TIPOS NO CONCUERDAN EN LA ASIGNACION", m.Line, m.Col)
	}
	gen.AddCodes("//FINALIZANDO LA ASIGNACION DE "+m.id, ambito)
	return retorno
}
