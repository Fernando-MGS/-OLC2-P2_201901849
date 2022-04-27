package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	arrayList "github.com/colegno/arraylist"
)

type Match struct {
	head  interfaces.Expresion
	line  string
	col   string
	cases *arrayList.List
}

func NewMatch(head interfaces.Expresion, cases *arrayList.List, line, col int) Match {
	instr := Match{head, strconv.Itoa(line), strconv.Itoa(col), cases}
	return instr
}

func (m Match) Ejecutar(env interface{}, gen *generator.Generator) interface{} {
	var retorno interfaces.Value
	retorno.Type = interfaces.NULL
	header := m.head.Ejecutar(env, gen)
	index := 0
	conf_tipo := false
	salida := gen.NewLabel()
	label_cond := ""
	ambito := env.(environment.Environment).DevAmbito()
	ret := gen.NewTemp()
	if header.Type != interfaces.STRUCT {
		conf_tipo = true
	} else {
		env.(environment.Environment).NewError("NO SE PUEDEN COMPARAR ELEMENTOS DE TIPO STRUCT", m.line, m.col)
		retorno.Type = interfaces.NULL
		return retorno
	}
	if conf_tipo && comprobar_tipo(m.cases) {
		gen.AddCodes("//INICIO DE MATCH", ambito)
		for index < m.cases.Len() {
			//entrada:=gen.NewLabel()
			//gen.AddCodes(entrada+":",ambito)
			actual_case := m.cases.GetValue(index).(interfaces.Cases)
			i := 0
			conditions := actual_case.Condicion.Len()
			if actual_case.Tipo == 1 {
				conditions = 1
			}
			conf := true
			num_cond := 0
			cond_true := gen.NewLabel()
			next_cond := gen.NewLabel()
			gen.AddCodes("//INICIO DEL CASO "+strconv.Itoa(index+1), ambito)
			gen.AddCodes("//REVISION DE CONDICIONES", ambito)
			if actual_case.Tipo == 0 {
				for i < conditions {

					if index > 0 {
						if i == 0 {
							gen.AddCodes(label_cond+":", ambito)
						}
					}
					//fmt.Println("el segundo for")
					//next_cond:=gen.NewLabel()
					var tmp interfaces.Value
					var aux interfaces.Expresion
					if actual_case.Tipo == 0 {
						conf = comprobar_conds(header, actual_case.Condicion, gen, env)
						aux = actual_case.Condicion.GetValue(i).(interfaces.Expresion)
						tmp = aux.Ejecutar(env, gen)
					}
					if conf {
						num_cond++
						code := "if(" + header.Value + "==" + tmp.Value + ") goto " + cond_true + ";"
						gen.AddCodes(code, ambito)
					} else {
						env.(environment.Environment).NewError("LOS TIPOS NO CONCUERDAN", m.line, m.col)
					}
					i++
				}
			} else {
				gen.AddCodes(label_cond+":", ambito)
			}

			if index < m.cases.Len()-1 {
				gen.AddCodes("goto "+next_cond+";", ambito)
				label_cond = next_cond
			}
			if num_cond > 0 || actual_case.Tipo == 1 {
				gen.AddCodes("//EJECUCION DE INSTRUCCIONES", ambito)
				gen.AddCodes(cond_true+":", ambito)
				in := env.(environment.Environment).Control.Entrada
				out := env.(environment.Environment).Control.Salida
				loop := env.(environment.Environment).Control.Ciclo
				stack := env.(environment.Environment).Control.Stack
				tmpEnv := environment.NewEnvironment(env.(environment.Environment), env.(environment.Environment).Control.Id, in, out, loop, stack)
				for _, s := range actual_case.Cuerpo.ToArray() {
					tipo := reflect.TypeOf(s)
					t := fmt.Sprintf("%v", tipo)
					if posicion := strings.Index(t, "Expresion"); posicion != -1 {
						res := s.(interfaces.Expresion).Ejecutar(tmpEnv, gen)
						code := ret + "=" + res.Value + ";"
						gen.AddCodes(code, ambito)
						//return ret
					} else {
						s.(interfaces.Instruction).Ejecutar(tmpEnv, gen)
					}

				}
				gen.AddCodes("//FIN DE EJECUCION DE INSTRUCCIONES", ambito)
				gen.AddCodes("goto "+salida+";", ambito)
			}
			gen.AddCodes("//FIN DEL CASO", ambito)
			index++

		}
	} else {
		env.(environment.Environment).NewError("EL CASO DEFAULT DEBE IR POR ULTIMO", m.line, m.col)
		retorno.Type = interfaces.NULL
	}
	gen.AddCodes(salida+":", ambito)
	return retorno
}

func comprobar_tipo(cases *arrayList.List) bool {
	a := true
	i := 0
	for _, s := range cases.ToArray() {
		if s.(interfaces.Cases).Tipo == 1 && i < cases.Len()-1 {
			return false
		}
		i++
	}
	return a
}

func comprobar_conds(head interfaces.Value, conditions *arrayList.List, gen *generator.Generator, env interface{}) bool {
	a := true
	for _, s := range conditions.ToArray() {
		aux := s.(interfaces.Expresion)
		tmp := aux.Ejecutar(env, gen)
		if tmp.Type != head.Type {
			return false
		}
	}
	return a
}
