package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"

	arrayList "github.com/colegno/arraylist"
)

type For struct {
	Cuerpo    *arrayList.List
	Variable  string
	Iteracion interfaces.For_Range
	line      string
	col       string
}

func NewFor(Cuerpo *arrayList.List, Variable string, Iteracion interfaces.For_Range, line, col int) For {
	instr := For{Cuerpo, Variable, Iteracion, strconv.Itoa(line), strconv.Itoa(col)}
	return instr
}

func (f For) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var result interfaces.Value
	result.Type = interfaces.NULL
	rangoInf := f.Iteracion.Range1.Ejecutar(env, gen)
	rangoSup := f.Iteracion.Range2.Ejecutar(env, gen)
	ambito := env.(environment.Environment).DevAmbito()
	stack := env.(environment.Environment).Control.Stack
	tmpEnv := environment.NewEnvironment(env, env.(environment.Environment).Control.Id, "", "", true, stack)
	if f.Iteracion.Tipo == 0 {
		if rangoInf.Type != interfaces.INTEGER && rangoInf.Type != interfaces.USIZE || rangoSup.Type != interfaces.INTEGER && rangoSup.Type != interfaces.USIZE {
			fmt.Println(rangoInf.Type)
			fmt.Println(rangoSup.Type)
			env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION NUMERICA", f.line, f.col)
			return result
		}
		tipos := interfaces.TipoSimbolo{Tipo: interfaces.INTEGER, Tipo2: arrayList.New()}
		entrada := gen.NewLabel()
		ciclo := gen.NewLabel()
		salida := gen.NewLabel()
		contador := gen.NewTemp()
		pre_contador := gen.NewTemp()
		//p := gen.Stack
		gen.AddCodes("//INICIO DE FOR-----", ambito)
		//env_stack := env.(environment.Environment).Control.Stack
		//env.(environment.Environment).Control.Stack++
		tam := gen.NewTemp()
		variable := interfaces.Symbol{Id: f.Variable, Tipo: tipos, Posicion: tam, Mutable: true, Line: f.line, Col: f.col}
		tmpEnv.SaveVariable(f.line, f.col, f.Variable, variable, tipos)
		gen.AddCodes(tam+"=P;", ambito)
		gen.AddCodes("STACK[(int)"+tam+"]="+rangoInf.Value+";//DECLARACION DE CONTADOR "+f.Variable, ambito)
		gen.AddCodes("P=P+1;", ambito)
		gen.AddCodes(entrada+":", ambito)
		llamada := gen.NewTemp()
		gen.AddCodes(llamada+"=STACK[(int)"+tam+"];", ambito)
		code := "if(" + llamada + "<" + rangoSup.Value + ") goto " + ciclo + ";"
		gen.AddCodes(code, ambito)
		gen.AddCodes("goto "+salida+";", ambito)
		gen.AddCodes(ciclo+":", ambito)
		for _, s := range f.Cuerpo.ToArray() {
			s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
		}
		contar := pre_contador + "=STACK[(int)" + tam + "];"
		incrementar := contador + "=" + pre_contador + "+1;//INCREMENTO DE CONTADOR"
		modificar := "STACK[(int)" + tam + "]=" + contador + ";"
		gen.AddCodes(contar, ambito)
		gen.AddCodes(incrementar, ambito)
		gen.AddCodes(modificar, ambito)
		gen.AddCodes("goto "+entrada+";", ambito)
		gen.AddCodes(salida+":", ambito)
		gen.AddCodes("//FIN DE FOR-----", ambito)
	} else {
		fmt.Println("hoal")
		if rangoInf.Type == interfaces.ARRAY || rangoInf.Type == interfaces.VECTOR {
			//i:=gen.NewTemp()
			index := gen.NewTemp()
			largo := gen.NewTemp()
			value := gen.NewTemp()
			aux := gen.NewTemp()
			indice := gen.NewTemp()
			entrada := gen.NewLabel()
			salida := gen.NewLabel()

			var tipo_set interfaces.TipoExpresion
			tipo2 := arrayList.New()
			dimensions := rangoInf.Tipo2.GetValue(0).(interfaces.Dimensions)
			if dimensions.Dimensions.Len() > 1 {
				tipo_set = rangoInf.Type
				tipo_dim := Disminucion_Tipos(dimensions.Dimensions)
				tipo2 = tipo_dim
			} else {
				tipo_set = dimensions.Tipo
			}
			tipos := interfaces.TipoSimbolo{Tipo: tipo_set, Tipo2: tipo2}
			variable := interfaces.Symbol{Id: f.Variable, Tipo: tipos, Posicion: value, Mutable: true, Line: f.line, Col: f.col}
			fmt.Println(variable, " hola soy variable")
			tmpEnv.SaveVariable(f.line, f.col, f.Variable, variable, tipos)
			//gen.AddCodes(value+"=HEAP", ambito)
			gen.AddCodes(index+"=0;", ambito)
			gen.AddCodes(indice+"="+rangoInf.Value+";", ambito)
			gen.AddCodes(largo+"=HEAP[(int)"+rangoInf.Value+"];//DETERMINANDO EL LARGO DEL CONJUNTO", ambito)
			gen.AddCodes(entrada+":", ambito)
			gen.AddCodes("if ("+index+">="+largo+") goto "+salida+";", ambito)
			gen.AddCodes(index+"="+index+"+1;//INCREMENTO DEL CONTADOR", ambito)
			gen.AddCodes(indice+"="+indice+"+1;//INCREMENTO DEL INDICE", ambito)
			gen.AddCodes(aux+"=HEAP[(int)"+indice+"];", ambito)
			gen.AddCodes("STACK[(int)"+value+"]="+aux+";", ambito)
			//gen.AddCodes("printf(\"%f\","+value+");", ambito)
			//gen.AddCodes(value+"=HEAP[(int)"+value+"];", ambito)
			for _, s := range f.Cuerpo.ToArray() {
				s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
			}
			gen.AddCodes("goto "+entrada+";", ambito)
			gen.AddCodes(salida+":", ambito)
			//gen.AddCodes("STACK[(int)"+i+"]="+index+";",ambito)
			//tmpEnv.SaveVariable()
		} else {
			env.(environment.Environment).NewError("SE ESPERABA UN ARRAY O UN VECTOR", f.line, f.col)
		}
	}
	return result
}

func Disminucion_Tipos(lista *arrayList.List) *arrayList.List {
	list := arrayList.New()
	conf := true
	for _, s := range list.ToArray() {
		if conf {
			conf = false
		}
		list.Add(s)
	}
	return list
}
