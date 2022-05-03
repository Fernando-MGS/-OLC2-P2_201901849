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
	//ambito := env.(environment.Environment).DevAmbito()
	stack := env.(environment.Environment).Control.Stack
	tmpEnv := environment.NewEnvironment(env, env.(environment.Environment).Control.Id, "", "", true, stack)
	name := env.(environment.Environment).Control.Id
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
		//gen.AddCodes("//INICIO DE FOR-----", ambito)
		gen.NewComentario("INICIO DE FOR", name, true, false, f.line)
		//env_stack := env.(environment.Environment).Control.Stack
		//env.(environment.Environment).Control.Stack++
		tam := gen.NewTemp()
		variable := interfaces.Symbol{Id: f.Variable, Tipo: tipos, Posicion: tam, Mutable: true, Line: f.line, Col: f.col}
		tmpEnv.SaveVariable(f.line, f.col, f.Variable, variable, tipos)
		//gen.AddCodes(tam+"=P;", ambito)
		gen.NewAsignacion(tam, "P", false, "", name, true, true, f.line)
		fmt.Println("eL NAME DE TAM ES ", tam)
		//gen.AddCodes("STACK[(int)"+tam+"]="+rangoInf.Value+";//DECLARACION DE CONTADOR "+f.Variable, ambito)
		gen.NewStack(tam, rangoInf.Value, true, "DECLARACION DE CONTADOR "+f.Variable, name, true, false, f.line)
		//gen.AddCodes("P=P+1;", ambito)
		gen.NewOperacion("P", "P", "+", "1", false, "", name, true, true, f.line)
		//gen.AddCodes(entrada+":", ambito)
		gen.NewLabels(entrada, false, "", name, true, true, "")
		llamada := gen.NewTemp()
		//gen.AddCodes(llamada+"=STACK[(int)"+tam+"];", ambito)
		gen.NewCallStack(llamada, tam, false, "", name, true, true, f.line)
		//code := "if(" + llamada + "<" + rangoSup.Value + ") goto " + ciclo + ";"
		gen.NewIf(llamada, "<", rangoSup.Value, ciclo, false, "", name, true, true, f.line)
		//gen.AddCodes(code, ambito)
		//gen.AddCodes("goto "+salida+";", ambito)
		gen.NewSalto(salida, false, "", name, true, true, f.line)
		//gen.AddCodes(ciclo+":", ambito)
		gen.NewLabels(ciclo, false, "", name, true, true, "")
		for _, s := range f.Cuerpo.ToArray() {
			s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
		}
		//contar := pre_contador + "=STACK[(int)" + tam + "];"
		gen.NewCallStack(pre_contador, tam, false, "", name, true, true, f.line)
		//incrementar := contador + "=" + pre_contador + "+1;//INCREMENTO DE CONTADOR"
		gen.NewOperacion(contador, pre_contador, "+", "1", false, "", name, true, true, f.line)
		//modificar := "STACK[(int)" + tam + "]=" + contador + ";"
		gen.NewStack(tam, contador, false, ""+f.Variable, name, true, true, f.line)
		/*gen.AddCodes(contar, ambito)
		gen.AddCodes(incrementar, ambito)
		gen.AddCodes(modificar, ambito)*/
		//gen.AddCodes("goto "+entrada+";", ambito)
		gen.NewSalto(entrada, false, "", name, true, false, f.line)
		//gen.AddCodes(salida+":", ambito)
		gen.NewLabels(salida, false, "", name, true, true, "")
		//gen.AddCodes("//FIN DE FOR-----", ambito)
		gen.NewComentario("FIN DE FOR", name, true, false, f.line)
	} else {
		//fmt.Println("hoal")
		if rangoInf.Type == interfaces.ARRAY || rangoInf.Type == interfaces.VECTOR {
			//i:=gen.NewTemp()
			fmt.Println("ESTOY EN UN FOR")
			index := gen.NewTemp()
			largo := gen.NewTemp()
			value := gen.NewTemp()
			aux := gen.NewTemp()
			indice := gen.NewTemp()
			entrada := gen.NewLabel()
			salida := gen.NewLabel()

			var tipo_set interfaces.TipoExpresion
			tipo2 := arrayList.New()
			fmt.Println("TIPO 2")
			fmt.Println(rangoInf.Tipo2.ToArray()...)
			dimensions := rangoInf.Tipo2.GetValue(0).(interfaces.Dimensions)
			conf_tipo := false
			if dimensions.Dimensions.Len() > 1 {
				tipo_set = rangoInf.Type
				tipo_dim := Disminucion_Tipos(dimensions.Dimensions)
				new_dim := interfaces.Dimensions{Tipo: dimensions.Tipo, Dimensions: tipo_dim}
				list := arrayList.New()
				list.Add(new_dim)
				tipo2 = list
			} else {
				conf_tipo = true
				tipo_set = dimensions.Tipo
			}
			tipos := interfaces.TipoSimbolo{Tipo: tipo_set, Tipo2: tipo2}
			variable := interfaces.Symbol{Id: f.Variable, Tipo: tipos, Posicion: value, Posicion2: value, Mutable: true, Line: f.line, Col: f.col}
			//fmt.Println(variable, " hola soy variable")
			tmpEnv.SaveVariable(f.line, f.col, f.Variable, variable, tipos)
			//gen.AddCodes(value+"=HEAP", ambito)
			//gen.AddCodes(index+"=0;", ambito)
			gen.NewComentario("INICIO DE FOR", name, true, false, f.line)
			gen.NewAsignacion(index, "0", false, "", name, true, true, f.line)
			//gen.AddCodes(indice+"="+rangoInf.Value+";", ambito)
			gen.NewAsignacion(indice, rangoInf.Value, false, "", name, true, true, f.line)
			//gen.AddCodes(largo+"=HEAP[(int)"+rangoInf.Value+"];//DETERMINANDO EL LARGO DEL CONJUNTO", ambito)
			gen.NewCallHeap(largo, rangoInf.Value, false, "", name, true, false, f.line)
			if conf_tipo {
				gen.NewAsignacion(value, "P", true, "ASIGNACION DE LA POSICION DE "+f.Variable, name, true, false, "")
				gen.NewOperacion("P", "P", "+", "1", false, "", name, true, false, "")
			}

			//en.AddCodes(entrada+":", ambito)
			gen.NewLabels(entrada, false, "", name, true, true, "")
			//gen.AddCodes("if ("+index+">="+largo+") goto "+salida+";", ambito)
			gen.NewIf(index, ">=", largo, salida, false, "", name, true, true, f.line)
			//gen.AddCodes(index+"="+index+"+1;//INCREMENTO DEL CONTADOR", ambito)
			gen.NewOperacion(index, index, "+", "1", true, "INCREMENTO DEL CONTADOR", name, true, true, f.line)
			//gen.AddCodes(indice+"="+indice+"+1;//INCREMENTO DEL INDICE", ambito)
			gen.NewOperacion(indice, indice, "+", "1", true, "INCREMENTO DEL INDICE", name, true, true, f.line)
			//gen.AddCodes(aux+"=HEAP[(int)"+indice+"];", ambito)
			gen.NewCallHeap(aux, indice, false, "", name, true, true, f.line)
			//gen.AddCodes("STACK[(int)"+value+"]="+aux+";", ambito)
			if conf_tipo {
				gen.NewStack(value, aux, true, "DECLARACION DE CONTADOR "+f.Variable, name, true, false, f.line)
			} else {
				gen.NewAsignacion(value, aux, true, "ACTUALIZACION DE CONTADOR "+f.Variable, name, true, false, f.line)
			}
			//gen.AddCodes("printf(\"%f\","+value+");", ambito)
			//gen.AddCodes(value+"=HEAP[(int)"+value+"];", ambito)
			for _, s := range f.Cuerpo.ToArray() {
				s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
			}
			//gen.AddCodes("goto "+entrada+";", ambito)
			gen.NewSalto(entrada, false, "", name, true, false, f.line)
			//gen.AddCodes(salida+":", ambito)
			gen.NewLabels(salida, false, "", name, true, true, "")
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
