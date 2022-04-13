package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
	"time"
)

type Aritmetica struct {
	Op1      interfaces.Expresion
	Operador string
	Op2      interfaces.Expresion
	Unario   bool
	Line     string
	Col      string
}

func NewOperacion(Op1 interfaces.Expresion, Operador string, Op2 interfaces.Expresion, unario bool, line, col int) Aritmetica {

	exp := Aritmetica{Op1, Operador, Op2, unario, strconv.Itoa(line), strconv.Itoa(col)}
	return exp
}

func (p Aritmetica) Ejecutar(env interface{}, gen *generator.Generator) interfaces.Value {
	/*suma_resta_dominante := [5][5]interfaces.TipoExpresion{
		//INTEGER			//FLOAT			   //STRING			  //BOOLEAN		   //NULL
		//INTEGER
		{interfaces.INTEGER, interfaces.FLOAT, interfaces.STRING, interfaces.NULL, interfaces.NULL},
		//FLOAT
		{interfaces.FLOAT, interfaces.FLOAT, interfaces.STRING, interfaces.NULL, interfaces.NULL},
		//STRING
		{interfaces.STRING, interfaces.STRING, interfaces.STRING, interfaces.STRING, interfaces.NULL},
		//BOOLEAN
		{interfaces.NULL, interfaces.NULL, interfaces.STRING, interfaces.NULL, interfaces.NULL},
		//NULL
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
	}

	multi_division_dominante := [5][5]interfaces.TipoExpresion{
		{interfaces.INTEGER, interfaces.FLOAT, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.FLOAT, interfaces.FLOAT, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
		{interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL, interfaces.NULL},
	}*/
	var retornoIzq interfaces.Value
	var retornoDer interfaces.Value

	if p.Unario == true {
		retornoIzq = p.Op1.Ejecutar(env, gen)

	} else {
		retornoIzq = p.Op1.Ejecutar(env, gen)
		retornoDer = p.Op2.Ejecutar(env, gen)
	}
	if retornoIzq.Type == interfaces.NULL || retornoDer.Type == interfaces.NULL {
		return interfaces.Value{Type: interfaces.NULL}
	}
	ambito := true
	if env.(environment.Environment).Control.Id == "GLOBAL" || env.(environment.Environment).Control.Id == "main" {
		ambito = false
	}
	//var dominante interfaces.TipoExpresion

	newTemp := gen.NewTemp()

	switch p.Operador {
	case "+":
		{

			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {

					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {

					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else {
					fmt.Print("ERROR: No es posible sumar")
				}
			} else {
				if retornoIzq.Type == interfaces.STR && retornoDer.Type == interfaces.STRING {
					gen.AddCodes("//CONCATENACION DE STRINGS", ambito)
					gen.AddFuncExtra("CONCATSTR")
					tmp1 := gen.NewTemp()
					asignacion1 := tmp1 + "=P+1;"
					gen.AddCodes(asignacion1, ambito)
					tmp2 := gen.NewTemp()
					asignacion2 := tmp2 + "=P+2;"
					gen.AddCodes(asignacion2, ambito)
					asignacion1 = "STACK[(int)" + tmp1 + "]=" + retornoIzq.Value + ";"
					asignacion2 = "STACK[(int)" + tmp2 + "]=" + retornoDer.Value + ";"
					gen.AddCodes(asignacion1, ambito)
					gen.AddCodes(asignacion2, ambito)
					gen.AddCodes("proc_concatSTR();", ambito)
					tmp3 := gen.NewTemp()
					code := tmp3 + "=STACK[(int)P];\n"
					gen.AddCodes(code, ambito)
					gen.AddCodes("//FIN DE CONCATENACION DE STRINGS", ambito)
					return interfaces.Value{Value: tmp3, IsTemp: true, Type: interfaces.STR, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.STRING && retornoDer.Type == interfaces.STR {
					gen.AddCodes("//CONCATENACION DE STRINGS", ambito)
					gen.AddFuncExtra("CONCATSTR")
					tmp1 := gen.NewTemp()
					asignacion1 := tmp1 + "=P+1;"
					gen.AddCodes(asignacion1, ambito)
					tmp2 := gen.NewTemp()
					asignacion2 := tmp2 + "=P+2;"
					gen.AddCodes(asignacion2, ambito)
					asignacion1 = "STACK[(int)" + tmp1 + "]=" + retornoIzq.Value + ";"
					asignacion2 = "STACK[(int)" + tmp2 + "]=" + retornoDer.Value + ";"
					gen.AddCodes(asignacion1, ambito)
					gen.AddCodes(asignacion2, ambito)
					gen.AddCodes("proc_concatSTR();", ambito)
					tmp3 := gen.NewTemp()
					code := tmp3 + "=STACK[(int)P];\n"
					gen.AddCodes(code, ambito)
					gen.AddCodes("//FIN DE CONCATENACION DE STRINGS", ambito)
					return interfaces.Value{Value: tmp3, IsTemp: true, Type: interfaces.STR, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.INTEGER && retornoDer.Type == interfaces.USIZE {
					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.USIZE && retornoDer.Type == interfaces.INTEGER {
					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "NO ES POSIBLE SUMAR ESTE TIPO DATO", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}
		}

	case "-":
		{
			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {

					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {

					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				}
			} else {
				if retornoIzq.Type == interfaces.INTEGER && retornoDer.Type == interfaces.USIZE {
					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.USIZE && retornoDer.Type == interfaces.INTEGER {
					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "NO ES POSIBLE RESTAR ESTE TIPO DATO", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}

	case "*":
		{
			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {

					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {

					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				}
			} else {
				if retornoIzq.Type == interfaces.INTEGER && retornoDer.Type == interfaces.USIZE {
					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.USIZE && retornoDer.Type == interfaces.INTEGER {
					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "NO ES POSIBLE RESTAR ESTE TIPO DATO", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}

	case "/":
		{
			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {
					gen.AddCodes(Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0"), ambito)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {
					gen.AddCodes(Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0.0"), ambito)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				}
			} else {
				if retornoIzq.Type == interfaces.STRING && retornoDer.Type == interfaces.STR {
					gen.AddCodes("//CONCATENACION DE STRINGS", ambito)
					gen.AddFuncExtra("CONCATSTR")
					tmp1 := gen.NewTemp()
					asignacion1 := tmp1 + "=P+1;"
					gen.AddCodes(asignacion1, ambito)
					tmp2 := gen.NewTemp()
					asignacion2 := tmp2 + "=P+2;"
					gen.AddCodes(asignacion2, ambito)
					asignacion1 = "STACK[(int)" + tmp1 + "]=" + retornoIzq.Value + ";"
					asignacion2 = "STACK[(int)" + tmp2 + "]=" + retornoDer.Value + ";"
					gen.AddCodes(asignacion1, ambito)
					gen.AddCodes(asignacion2, ambito)
					gen.AddCodes("proc_concatSTR();", ambito)
					tmp3 := gen.NewTemp()
					code := tmp3 + "=STACK[(int)P];\n"
					gen.AddCodes(code, ambito)
					gen.AddCodes("//FIN DE CONCATENACION DE STRINGS", ambito)
					return interfaces.Value{Value: tmp3, IsTemp: true, Type: interfaces.STR, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.INTEGER && retornoDer.Type == interfaces.USIZE {
					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.USIZE && retornoDer.Type == interfaces.INTEGER {
					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "NO ES POSIBLE DIVIDIR ESTE TIPO DATO", Fecha: fecha}
			env.(environment.Environment).Errores(err)

		}
	case "<":
		{
			if retornoIzq.Type == retornoDer.Type {
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					l1 := ""
					l2 := ""
					if gen.GetConf() == 0 {
						l1 = gen.NewLabel()
						l2 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + "<" + retornoDer.Value + ") goto " + l1 + ";\n"
						value += "goto " + l2 + ";\n"
						gen.AddCodes(value, ambito)
						gen.AddTempBool(l1, l2)
						gen.SetConf()
					} else if gen.GetConf() == 1 {
						l1 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + "<" + retornoDer.Value + ") goto "
						gen.AddTempBool(l1, value)
					}
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				} else {
					t := time.Now()
					fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
						t.Year(), t.Month(), t.Day(),
						t.Hour(), t.Minute(), t.Second())
					err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
					env.(environment.Environment).Errores(err)
				}
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}

	case ">":
		{
			if retornoIzq.Type == retornoDer.Type {
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					l1 := ""
					l2 := ""
					if gen.GetConf() == 0 {
						l1 = gen.NewLabel()
						l2 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + ">" + retornoDer.Value + ") goto " + l1 + ";\n"
						value += "goto " + l2 + ";\n"
						gen.AddCodes(value, ambito)
						gen.AddTempBool(l1, l2)
						gen.SetConf()
					} else if gen.GetConf() == 1 {
						l1 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + ">" + retornoDer.Value + ") goto "
						gen.AddTempBool(l1, value)
					}
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				} else {
					t := time.Now()
					fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
						t.Year(), t.Month(), t.Day(),
						t.Hour(), t.Minute(), t.Second())
					err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
					env.(environment.Environment).Errores(err)
				}
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}
	case ">=":
		{
			if retornoIzq.Type == retornoDer.Type {
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					l1 := ""
					l2 := ""
					if gen.GetConf() == 0 {
						l1 = gen.NewLabel()
						l2 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + ">=" + retornoDer.Value + ") goto " + l1 + ";\n"
						value += "goto " + l2 + ";\n"
						gen.AddCodes(value, ambito)
						gen.AddTempBool(l1, l2)
						gen.SetConf()
					} else if gen.GetConf() == 1 {
						l1 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + ">=" + retornoDer.Value + ") goto "
						gen.AddTempBool(l1, value)
					}
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.STRING {

				}
			} else {
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}

	case "<=":
		{
			if retornoDer.Type == retornoIzq.Type {
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					l1 := ""
					l2 := ""
					if gen.GetConf() == 0 {
						l1 = gen.NewLabel()
						l2 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + "<=" + retornoDer.Value + ") goto " + l1 + ";\n"
						value += "goto " + l2 + ";\n"
						gen.AddCodes(value, ambito)
						gen.AddTempBool(l1, l2)
						gen.SetConf()
					} else if gen.GetConf() == 1 {
						l1 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + "<=" + retornoDer.Value + ") goto "
						gen.AddTempBool(l1, value)
					}
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				} else {
					t := time.Now()
					fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
						t.Year(), t.Month(), t.Day(),
						t.Hour(), t.Minute(), t.Second())
					err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
					env.(environment.Environment).Errores(err)
				}
			} else {
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}
	case "!=":
		{
			if retornoDer.Type == retornoIzq.Type {

				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					l1 := ""
					l2 := ""
					if gen.GetConf() == 0 {
						l1 = gen.NewLabel()
						l2 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + "!=" + retornoDer.Value + ") goto " + l1 + ";\n"
						value += "goto " + l2 + ";\n"
						gen.AddCodes(value, ambito)
						gen.AddTempBool(l1, l2)
						gen.SetConf()
					} else if gen.GetConf() == 1 {
						l1 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + "!=" + retornoDer.Value + ") goto "
						gen.AddTempBool(l1, value)

					}
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				} else if retornoDer.Type == interfaces.STRING || retornoIzq.Type == interfaces.STR {
					code := "//COMPARACION DE STRING\n"
					t1 := gen.NewTemp()
					t4 := gen.NewTemp()
					t2 := retornoIzq.Value
					t3 := retornoDer.Value
					resultado := gen.NewTemp()
					code += t1 + "=P+1;\n"
					code += "STACK[(int)" + t1 + "]=" + t2 + ";\n"
					code += t4 + "=P+2;\n"
					code += "STACK[(int)" + t4 + "]=" + t3 + ";\n"
					gen.AddFuncExtra("COMPARESTR")
					code += "proc_compareString();\n"
					code += resultado + "=STACK[(int)P];"
					gen.AddCodes(code, ambito)
					//
					l1 := ""
					l2 := ""
					if gen.GetConf() == 0 {
						l1 = gen.NewLabel()
						l2 = gen.NewLabel()
						value := "if (" + resultado + "==0) goto " + l1 + ";\n"
						value += "goto " + l2 + ";\n"
						gen.AddCodes(value, ambito)
						gen.AddTempBool(l1, l2)
						gen.SetConf()
					} else {
						l1 = gen.NewLabel()
						value := "if (" + resultado + "==0) goto "
						gen.AddTempBool(l1, value)
					}
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				}
			} else {
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
			env.(environment.Environment).Errores(err)

		}
	case "==":
		{
			if retornoDer.Type == retornoIzq.Type {
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					l1 := ""
					l2 := ""
					if gen.GetConf() == 0 {
						l1 = gen.NewLabel()
						l2 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + "==" + retornoDer.Value + ") goto " + l1 + ";\n"
						value += "goto " + l2 + ";\n"
						gen.AddCodes(value, ambito)
						gen.AddTempBool(l1, l2)
						gen.SetConf()
					} else if gen.GetConf() == 1 {
						l1 = gen.NewLabel()
						value := "if (" + retornoIzq.Value + "==" + retornoDer.Value + ") goto "
						gen.AddTempBool(l1, value)
					}
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				} else if retornoDer.Type == interfaces.STRING || retornoIzq.Type == interfaces.STR {
					code := "//COMPARACION DE STRING\n"
					t1 := gen.NewTemp()
					t4 := gen.NewTemp()
					t2 := retornoIzq.Value
					t3 := retornoDer.Value
					resultado := gen.NewTemp()
					code += t1 + "=P+1;\n"
					code += "STACK[(int)" + t1 + "]=" + t2 + ";\n"
					code += t4 + "=P+2;\n"
					code += "STACK[(int)" + t4 + "]=" + t3 + ";\n"
					gen.AddFuncExtra("COMPARESTR")
					code += "proc_compareString();\n"
					code += resultado + "=STACK[(int)P];"
					gen.AddCodes(code, ambito)
					//
					l1 := ""
					l2 := ""
					if gen.GetConf() == 0 {
						l1 = gen.NewLabel()
						l2 = gen.NewLabel()
						value := "if (" + resultado + "==1) goto " + l1 + ";\n"
						value += "goto " + l2 + ";\n"
						gen.AddCodes(value, ambito)
						gen.AddTempBool(l1, l2)
						gen.SetConf()
					} else {
						l1 = gen.NewLabel()
						value := "if (" + resultado + "==1) goto "
						gen.AddTempBool(l1, value)
					}
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				}
			} else {
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}

		}
	case "||":
		{
			if retornoIzq.Type == retornoDer.Type && retornoDer.Type == interfaces.BOOLEAN {
				anterior := gen.GetTempsB()
				/*fmt.Println("entro||-" + anterior.TrueL)
				fmt.Println(anterior.FalseL)
				fmt.Println(anterior.TrueL1)
				fmt.Println(anterior.FalseL1)
				fmt.Println("============")*/
				l2 := gen.NewLabel()
				value := anterior.FalseL + ":\n"
				value += anterior.FalseL1 + anterior.TrueL + ";\n"
				value += "goto " + l2 + ";"
				fmt.Println(value)
				gen.LabelsOr(l2)
				gen.AddCodes(value, ambito)
				//gen.RotarLabels()
				/*anterior = gen.GetTempsB()
				fmt.Println("salio-" + anterior.TrueL)
				fmt.Println(anterior.FalseL)
				fmt.Println(anterior.TrueL1)
				fmt.Println(anterior.FalseL1)
				fmt.Println("============")*/
				return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
			} else {
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}

		}
	case "&&":
		{
			fmt.Println("&&")
			fmt.Println(retornoDer)
			fmt.Println(retornoIzq)
			if retornoIzq.Type == retornoDer.Type && retornoDer.Type == interfaces.BOOLEAN {
				anterior := gen.GetTempsB()
				/*fmt.Println("entro-" + anterior.TrueL)
				fmt.Println(anterior.FalseL)
				fmt.Println(anterior.TrueL1)
				fmt.Println(anterior.FalseL1)
				fmt.Println("============")*/
				value := anterior.TrueL + ":\n"
				value += anterior.FalseL1 + anterior.TrueL1 + ";\n"
				value += "goto " + anterior.FalseL + ";"
				gen.AddCodes(value, ambito)
				gen.RotarLabels()
				/*anterior = gen.GetTempsB()
				fmt.Println("salio-" + anterior.TrueL)
				fmt.Println(anterior.FalseL)
				fmt.Println(anterior.TrueL1)
				fmt.Println(anterior.FalseL1)
				fmt.Println("============")*/
				return interfaces.Value{Value: newTemp, IsTemp: false, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				//gen.AddTempBool(anterior.TrueL1, anterior.FalseL1)
			} else {
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}

		}
	case "!":
		{
			if retornoIzq.Type == interfaces.BOOLEAN {
				gen.InvertirLabels()
				return interfaces.Value{Value: newTemp, IsTemp: false, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
			} else {
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "SE ESPERABA UNA EXPRESION BOOLEANA", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}
		}
	case "°":
		{
			if retornoIzq.Type == interfaces.INTEGER {
				//tmp:=gen.NewTemp()
				gen.AddExpression(newTemp, retornoDer.Value, "-1", "*", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
			} else if retornoIzq.Type == interfaces.FLOAT {
				gen.AddExpression(newTemp, retornoDer.Value, "-1", "*", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "SE ESPERABA UNA EXPRESION NUMERICA", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}
	case "^":
		{
			if retornoIzq.Type == retornoDer.Type {
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					tmp2 := gen.NewTemp()
					resultado := gen.NewTemp()
					//llamada a la funcion de potencias
					code := "//INICIO DE CALCULO DE POTENCIAS\n"
					code += newTemp + "=P+1;\n"
					code += tmp2 + "=P+2;\n"
					code += "STACK[(int)" + newTemp + "]=" + retornoIzq.Value + ";\n"
					code += "STACK[(int)" + tmp2 + "]=" + retornoDer.Value + ";\n"
					gen.AddFuncExtra("POW")
					code += "proc_potencia();\n"
					code += resultado + "=STACK[(int)P];"
					gen.AddCodes(code, ambito)
					return interfaces.Value{Value: resultado, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
					//code+=""
				}

			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "NO ES POSIBLE OPERAR ESTE TIPO DATO", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}
	case "%":
		{
			if retornoIzq.Type == retornoDer.Type {
				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.FLOAT {
					//resultado:=gen.NewTemp()
					code := "//CALCULO DE MODULO\n"
					code += newTemp + "=fmod(" + retornoIzq.Value + "," + retornoDer.Value + ");"
					gen.AddCodes(code, ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
			}
			t := time.Now()
			fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
				t.Year(), t.Month(), t.Day(),
				t.Hour(), t.Minute(), t.Second())
			err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "NO ES POSIBLE OBTENER EL MODULO DE ESTE TIPO DATO", Fecha: fecha}
			env.(environment.Environment).Errores(err)
		}
	}

	return interfaces.Value{Value: "0", Type: interfaces.NULL, IsTemp: false, TrueLabel: "", FalseLabel: ""}
}

func Comprobar_Div(g *generator.Generator, izq, der interfaces.Value, tmp, value string) string {
	l1 := g.NewLabel()
	l2 := g.NewLabel()
	code := "if(" + der.Value + "!=0) goto " + l1 + ";\n"
	g.AddFuncExtra("DIVZERO")
	code += "proc_divZero();\n"
	code += tmp + "=" + value + ";\n"
	code += "goto " + l2 + ";\n"
	code += l1 + ":\n"
	code += tmp + "=" + izq.Value + "/" + der.Value + ";\n"
	code += l2 + ":"
	return code
}
