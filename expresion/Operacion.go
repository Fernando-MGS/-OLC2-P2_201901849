package expresion

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
	"strconv"
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
	/*if p.Unario {
		retornoIzq = p.Op1.Ejecutar(env, gen)

	} else {
		retornoIzq = p.Op1.Ejecutar(env, gen)
		retornoDer = p.Op2.Ejecutar(env, gen)
	}
	if retornoIzq.Type == interfaces.NULL || retornoDer.Type == interfaces.NULL {
		return interfaces.Value{Type: interfaces.NULL}
	}*/

	ambito := true
	if env.(environment.Environment).Control.Id == "GLOBAL" || env.(environment.Environment).Control.Id == "main" {
		ambito = false
	}
	//var dominante interfaces.TipoExpresion

	newTemp := gen.NewTemp()

	switch p.Operador {
	case "+":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {

					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {

					gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else {
					env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
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
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}
		}

	case "-":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
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
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}

	case "*":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
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
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}

	case "/":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {
					gen.AddCodes(Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0", 0), ambito)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {
					gen.AddCodes(Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0.0", 0), ambito)
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
			env.(environment.Environment).NewError("NO ES POSIBLE DIVIDIR ESTE TIPO DE DATO", p.Line, p.Col)

		}
	case "<":
		{
			gen.AddCodes("//--------<", ambito)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
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
				} else if retornoIzq.Type == interfaces.STRING {
					code := "//COMPARACION DE STRING\n"
					dir1 := gen.NewTemp() //dirección del primer parametro
					dir2 := gen.NewTemp() //dirección del segundo parametro
					code += dir1 + "=P+1;\n"
					code += dir2 + "=P+2;\n"
					code += "STACK[(int)" + dir1 + "]=" + retornoIzq.Value + ";\n"
					code += "STACK[(int)" + dir2 + "]=" + retornoDer.Value + ";\n"
					gen.AddFuncExtra("COMPARELONG")
					code += "compareLong_String();\n"
					resultado := gen.NewTemp()
					//code += "STACK[(int)P]=" + retornoDer.Value + ";\n"
					code += resultado + "=STACK[(int)P];\n"
					gen.AddCodes(code, ambito)
					addCodeBool(gen, "==", "0", resultado, ambito)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				}
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}

	case ">":
		{
			gen.AddCodes("//-------->", ambito)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
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
				} else if retornoIzq.Type == interfaces.STRING {
					code := "//COMPARACION DE STRING\n"
					dir1 := gen.NewTemp() //dirección del primer parametro
					dir2 := gen.NewTemp() //dirección del segundo parametro
					code += dir1 + "=P+1;\n"
					code += dir2 + "=P+2;\n"
					code += "STACK[(int)" + dir1 + "]=" + retornoIzq.Value + ";\n"
					code += "STACK[(int)" + dir2 + "]=" + retornoDer.Value + ";\n"
					gen.AddFuncExtra("COMPARELONG")
					code += "compareLong_String();\n"
					resultado := gen.NewTemp()
					//code += "STACK[(int)P]=" + retornoDer.Value + ";\n"
					code += resultado + "=STACK[(int)P];\n"
					gen.AddCodes(code, ambito)
					addCodeBool(gen, "==", "2", resultado, ambito)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				}
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}
	case ">=":
		{
			gen.AddCodes("//-------->=", ambito)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
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
					code := "//COMPARACION DE STRING\n"
					dir1 := gen.NewTemp() //dirección del primer parametro
					dir2 := gen.NewTemp() //dirección del segundo parametro
					code += dir1 + "=P+1;\n"
					code += dir2 + "=P+2;\n"
					code += "STACK[(int)" + dir1 + "]=" + retornoIzq.Value + ";\n"
					code += "STACK[(int)" + dir2 + "]=" + retornoDer.Value + ";\n"
					gen.AddFuncExtra("COMPARELONG")
					code += "compareLong_String();\n"
					resultado := gen.NewTemp()
					//code += "STACK[(int)P]=" + retornoDer.Value + ";\n"
					code += resultado + "=STACK[(int)P];\n"
					gen.AddCodes(code, ambito)
					addCodeBool(gen, ">=", "1", resultado, ambito)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				}
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}

	case "<=":
		{
			gen.AddCodes("//--------<=", ambito)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
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
				} else if interfaces.STR == retornoDer.Type || interfaces.STRING == retornoIzq.Type {
					code := "//COMPARACION DE STRING\n"
					dir1 := gen.NewTemp() //dirección del primer parametro
					dir2 := gen.NewTemp() //dirección del segundo parametro
					code += dir1 + "=P+1;\n"
					code += dir2 + "=P+2;\n"
					code += "STACK[(int)" + dir1 + "]=" + retornoIzq.Value + ";\n"
					code += "STACK[(int)" + dir2 + "]=" + retornoDer.Value + ";\n"
					gen.AddFuncExtra("COMPARELONG")
					code += "compareLong_String();\n"
					resultado := gen.NewTemp()
					//code += "STACK[(int)P]=" + retornoDer.Value + ";\n"
					code += resultado + "=STACK[(int)P];\n"
					gen.AddCodes(code, ambito)
					addCodeBool(gen, "<=", "1", resultado, ambito)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				} else {
					env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
				}
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}
	case "!=":
		{
			gen.AddCodes("//--------!=", ambito)
			resultado := ""
			resultado2 := ""
			gen.Conf2()
			//retornoIzq,retornoDer=DevTipos(p,env,gen)
			retornoIzq = p.Op1.Ejecutar(env, gen)
			//gen.AddCodes("//PRACTICA", ambito)
			if retornoIzq.Type == interfaces.BOOLEAN {
				gen.SetConf()
				code, resultado1 := True_NotTrue("!=", retornoIzq, gen)
				resultado = resultado1
				gen.AddCodes(code+"\n//FIN DE IZQ", ambito)
			}
			retornoDer = p.Op2.Ejecutar(env, gen)
			if retornoDer.Type == interfaces.BOOLEAN {
				//gen.SetConf()
				code, resultado1 := True_NotTrue("!=", retornoIzq, gen)
				resultado2 = resultado1
				gen.AddCodes(code+"\n//FIN DE DERTYPE", ambito)
				/*fmt.Println("ruta 2")
				fmt.Println(gen.GetTempsB())*/
			}
			//retornoIzq,retornoDer=DevTipos(p,env,gen)
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
						if gen.GetConf2() {
							gen.AddCodes("//practica", ambito)
						}
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
				} else if interfaces.BOOLEAN == retornoDer.Type {
					l1 := gen.NewLabel()
					l2 := gen.NewLabel()
					code := "if(" + resultado + "!=" + resultado2 + ") goto " + l1 + ";\n"
					code += "goto " + l2 + ";"
					gen.AddCodes(code, ambito)
					gen.SetConf()
					gen.AddTempBool(l1, l2)
					gen.SetConf()
					//gen.RotarLabels()
					//fmt.Println(gen.GetTempsB())
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
				}
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)

		}
	case "==":
		{
			gen.AddCodes("//--------==", ambito)
			resultado := ""
			resultado2 := ""
			gen.Conf2()
			//retornoIzq,retornoDer=DevTipos(p,env,gen)
			retornoIzq = p.Op1.Ejecutar(env, gen)
			//gen.AddCodes("//PRACTICA", ambito)
			if retornoIzq.Type == interfaces.BOOLEAN {
				gen.SetConf()
				code, resultado1 := True_NotTrue("==", retornoIzq, gen)
				resultado = resultado1
				gen.AddCodes(code+"\n//FIN DE IZQ", ambito)
			}
			retornoDer = p.Op2.Ejecutar(env, gen)
			if retornoDer.Type == interfaces.BOOLEAN {
				//gen.SetConf()
				code, resultado1 := True_NotTrue("==", retornoIzq, gen)
				resultado2 = resultado1
				gen.AddCodes(code+"\n//FIN DE DERTYPE", ambito)
				fmt.Println("ruta 2")
				fmt.Println(gen.GetTempsB())
			}
			/*fmt.Println(resultado)
			fmt.Println(resultado2)*/
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
				} else if interfaces.BOOLEAN == retornoIzq.Type {

					//fmt.Println(gen.GetConf())
					l1 := gen.NewLabel()
					l2 := gen.NewLabel()
					code := "if(" + resultado + "==" + resultado2 + ") goto " + l1 + ";\n"
					code += "goto " + l2 + ";"
					gen.AddCodes(code, ambito)
					gen.SetConf()
					gen.AddTempBool(l1, l2)
					gen.SetConf()
					//gen.RotarLabels()
					//fmt.Println(gen.GetTempsB())
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
					//gen.AddCodes(code, ambito)*/
				}
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}

		}
	case "||":
		{
			gen.AddCodes("//--------||", ambito)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == retornoDer.Type && retornoDer.Type == interfaces.BOOLEAN {
				anterior := gen.GetTempsB()
				l2 := gen.NewLabel()
				value := "//INICIO DE ||\n" + anterior.FalseL + ":\n"
				value += anterior.FalseL1 + anterior.TrueL + ";\n"
				value += "goto " + l2 + ";"
				fmt.Println(value)
				gen.LabelsOr(l2)
				gen.AddCodes(value, ambito)
				return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: "", FalseLabel: ""}
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}

		}
	case "&&":
		{
			gen.AddCodes("//--------&&", ambito)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			/*fmt.Println("&&")
			fmt.Println(retornoDer)
			fmt.Println(retornoIzq)*/
			if retornoIzq.Type == retornoDer.Type && retornoDer.Type == interfaces.BOOLEAN {
				anterior := gen.GetTempsB()
				newTrue := gen.NewLabel()
				value := "//INICIO DE &&\n" + anterior.TrueL + ":\n"
				value += anterior.FalseL1 + newTrue + ";\n"
				value += "goto " + anterior.FalseL + ";"
				gen.AddCodes(value, ambito)
				//fmt.Println(gen.GetTempsB())
				//gen.RotarLabels()

				gen.SetTrueFalse(newTrue, anterior.FalseL)
				fmt.Println(gen.GetTempsB())
				return interfaces.Value{Value: newTemp, IsTemp: false, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				//gen.AddTempBool(anterior.TrueL1, anterior.FalseL1)
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}

		}
	case "!":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == interfaces.BOOLEAN {
				gen.InvertirLabels()
				return interfaces.Value{Value: newTemp, IsTemp: false, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
			} else {
				env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION BOOLEANA", p.Line, p.Col)
			}
		}
	case "°":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == interfaces.INTEGER {
				//tmp:=gen.NewTemp()
				gen.AddExpression(newTemp, retornoIzq.Value, "-1", "*", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
			} else if retornoIzq.Type == interfaces.FLOAT {
				gen.AddExpression(newTemp, retornoIzq.Value, "-1", "*", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
			}
			env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION NUMERICA", p.Line, p.Col)
		}
	case "^":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
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
			env.(environment.Environment).NewError("NO ES POSIBLE OPERAR ESTE TIPO DATO", p.Line, p.Col)
		}
	case "%":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == retornoDer.Type {
				if retornoDer.Type == interfaces.INTEGER {
					//resultado:=gen.NewTemp()
					gen.AddCodes(Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0", 1), ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoDer.Type == interfaces.FLOAT {
					gen.AddCodes(Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0.0", 1), ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
			}
			env.(environment.Environment).NewError("NO ES POSIBLE OBTENER EL MODULO DE ESTE TIPO DATO", p.Line, p.Col)
		}
	}

	return interfaces.Value{Value: "0", Type: interfaces.NULL, IsTemp: false, TrueLabel: "", FalseLabel: ""}
}

func Comprobar_Div(g *generator.Generator, izq, der interfaces.Value, tmp, value string, tipo int) string {
	l1 := g.NewLabel()
	l2 := g.NewLabel()
	code := "if(" + der.Value + "!=0) goto " + l1 + ";\n"
	g.AddFuncExtra("DIVZERO")
	code += "proc_divZero();\n"
	code += tmp + "=" + value + ";\n"
	code += "goto " + l2 + ";\n"
	code += l1 + ":\n"
	if tipo == 0 {
		code += tmp + "=" + izq.Value + "/" + der.Value + ";\n"
	} else {
		code += tmp + "=fmod(" + izq.Value + "," + der.Value + ");\n"
	}
	code += l2 + ":"
	return code
}

func True_NotTrue(signo string, izq interfaces.Value, g *generator.Generator) (string, string) {
	anterior := g.GetTempsB()
	fmt.Println(anterior)
	tmpIzq := g.NewTemp()
	//tmpDer:=g.NewTemp()
	l1 := anterior.TrueL
	l2 := anterior.FalseL
	l3 := g.NewLabel() //salida de la primera asignacion
	/*l4:=anterior.TrueL1
	l5:=g.NewLabel()//FALSE 2*/
	value := "//INICIO DE " + signo + "\n" + l1 + ":\n"
	value += tmpIzq + "=1;\ngoto " + l3 + ";\n"
	value += l2 + ":\n"
	value += tmpIzq + "=0;\n"
	value += l3 + ":\n"

	return value, tmpIzq
}

func DevTipos(p Aritmetica, env interface{}, gen *generator.Generator) (interfaces.Value, interfaces.Value) {
	var retornoIzq interfaces.Value
	var retornoDer interfaces.Value
	if p.Unario {
		retornoIzq = p.Op1.Ejecutar(env, gen)

	} else {
		retornoIzq = p.Op1.Ejecutar(env, gen)
		retornoDer = p.Op2.Ejecutar(env, gen)
	}
	return retornoIzq, retornoDer
}

func addCodeBool(gen *generator.Generator, signo, value, temp string, ambito bool) {
	l1 := ""
	l2 := ""

	if gen.GetConf() == 0 {
		l1 = gen.NewLabel()
		l2 = gen.NewLabel()
		value := "if (" + temp + signo + value + ") goto " + l1 + ";\n"
		value += "goto " + l2 + ";\n"
		gen.AddCodes(value, ambito)
		gen.AddTempBool(l1, l2)
		gen.SetConf()

	} else if gen.GetConf() == 1 {
		l1 = gen.NewLabel()
		value := "if (" + temp + signo + value + ") goto "
		gen.AddTempBool(l1, value)
	}
}
