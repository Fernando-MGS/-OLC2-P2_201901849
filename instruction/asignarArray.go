package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
)

type ModArray struct {
	Access interfaces.Expresion
	Value  interfaces.Expresion
	Line   string
	Col    string
}

func NewModArray(index, val interfaces.Expresion, line, col int) ModArray {
	instr := ModArray{index, val, strconv.Itoa(line), strconv.Itoa(col)}
	return instr
}

func (m ModArray) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	///ambito := env.(environment.Environment).DevAmbito()
	bounds := m.Access.Ejecutar(env, gen)
	value := m.Value.Ejecutar(env, gen)
	name := env.(environment.Environment).Control.Id
	if bounds.Type == interfaces.NULL || value.Type == interfaces.NULL {
		return retorno
	}
	//gen.AddCodes("//INICIANDO LA ASIGNACION DE UN ARRAY O VECTOR", ambito)

	if bounds.Type == value.Type {
		if bounds.Type == interfaces.ARRAY || bounds.Type == interfaces.VECTOR {
			fmt.Println("AVERS I ENTRA")
			gen.NewComentario("INICIO DE ASIGNACION DE ARRAY O VECTOR", name, true, false, m.Line)
			//env.(environment.Environment).NewError("EL ELEMENTO QUE SE QUIERE MODIFICAR NO ES UN VECTOR", m.Line, m.Col)
			dim_bounds := bounds.Tipo2.GetValue(0).(interfaces.Dimensions)
			dim_value := value.Tipo2.GetValue(0).(interfaces.Dimensions)
			if dim_value.Dimensions.Len() != dim_bounds.Dimensions.Len() {
				env.(environment.Environment).NewError("LAS DIMENSIONES NO CONCUERDAN", m.Line, m.Col)
				return retorno
			}

			//gen.AddCodes(bounds.Value+"="+value.Value+";", ambito)
			//return retorno
		}
		//if bounds.FalseLabel == "1" {
		gen.NewHeap(bounds.TrueLabel, value.Value, false, "", name, true, false, m.Line)
		/*} else if bounds.FalseLabel == "2" {
			gen.NewAsignacion(bounds.Value, value.Value, true, "ASIGNACION DE STRUCT", name, true, true, "")
		}*/

		//gen.AddCodes("HEAP[(int)"+bounds.TrueLabel+"]="+value.Value+";", ambito)

		/*entrada := gen.NewLabel()
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
		}*/
	} else {
		env.(environment.Environment).NewError("LOS TIPOS NO CONCUERDAN EN LA ASIGNACION", m.Line, m.Col)
	}
	//gen.AddCodes("//FINALIZANDO LA ASIGNACION", ambito)
	gen.NewComentario("FINALIZANDO LA ASIGNACION", name, true, false, m.Line)
	return retorno
}
