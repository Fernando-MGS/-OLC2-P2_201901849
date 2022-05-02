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
	var retornoIzq interfaces.Value
	var retornoDer interfaces.Value
	//ambito := true
	name := env.(environment.Environment).Control.Id
	/*if env.(environment.Environment).Control.Id == "GLOBAL" || env.(environment.Environment).Control.Id == "main" {
		ambito = false
	}*/
	//var dominante interfaces.TipoExpresion

	newTemp := gen.NewTemp()

	switch p.Operador {
	case "+":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			/*fmt.Println(retornoIzq.Type, "  vs")
			fmt.Println(retornoDer.Type)*/
			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {

					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					gen.NewOperacion(newTemp, retornoIzq.Value, "+", retornoDer.Value, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {

					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					gen.NewOperacion(newTemp, retornoIzq.Value, "+", retornoDer.Value, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else {
					env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN SUMAR ENTRE SI", p.Line, p.Col)
				}
			} else {
				if retornoIzq.Type == interfaces.STR && retornoDer.Type == interfaces.STRING {
					//gen.AddCodes("//CONCATENACION DE STRINGS", ambito)
					gen.NewComentario("CONCATENACION DE STRINGS", name, true, false, p.Line)
					gen.AddFuncExtra("CONCATSTR")
					tmp1 := gen.NewTemp()
					//asignacion1 := tmp1 + "=P+1;"
					//gen.AddCodes(asignacion1, ambito)
					gen.NewOperacion(tmp1, "P", "+", "1", false, "", name, true, true, p.Line)
					tmp2 := gen.NewTemp()
					//asignacion2 := tmp2 + "=P+2;"
					//gen.AddCodes(asignacion2, ambito)
					gen.NewOperacion(tmp2, "P", "+", "2", false, "", name, true, true, p.Line)
					//asignacion1 = "STACK[(int)" + tmp1 + "]=" + retornoIzq.Value + ";"
					gen.NewStack(tmp1, retornoIzq.Value, false, "", name, true, true, p.Line)
					//asignacion2 = "STACK[(int)" + tmp2 + "]=" + retornoDer.Value + ";"
					gen.NewStack(tmp2, retornoDer.Value, false, "", name, true, true, p.Line)
					/*gen.AddCodes(asignacion1, ambito)
					gen.AddCodes(asignacion2, ambito)*/
					//gen.AddCodes("proc_concatSTR();", ambito)
					gen.NewLlamada("proc_concatSTR", false, "", name, true, false, p.Line)
					tmp3 := gen.NewTemp()
					/*code := tmp3 + "=STACK[(int)P];\n"
					gen.AddCodes(code, ambito)*/
					gen.NewCallStack(tmp3, "P", false, "", name, true, true, p.Line)
					//gen.AddCodes("//FIN DE CONCATENACION DE STRINGS", ambito)
					gen.NewComentario("FIN DE CONCATENACION DE STRINGS", name, true, false, p.Line)
					return interfaces.Value{Value: tmp3, IsTemp: true, Type: interfaces.STR, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.STRING && retornoDer.Type == interfaces.STR {
					//gen.AddCodes("//CONCATENACION DE STRINGS", ambito)
					gen.NewComentario("CONCATENACION DE STRINGS", name, true, false, p.Line)
					gen.AddFuncExtra("CONCATSTR")
					tmp1 := gen.NewTemp()
					/*asignacion1 := tmp1 + "=P+1;"
					gen.AddCodes(asignacion1, ambito)*/
					gen.NewOperacion(tmp1, "P", "+", "1", false, "", name, true, true, p.Line)
					tmp2 := gen.NewTemp()
					/*asignacion2 := tmp2 + "=P+2;"
					gen.AddCodes(asignacion2, ambito)*/
					gen.NewOperacion(tmp2, "P", "+", "2", false, "", name, true, true, p.Line)
					/*asignacion1 = "STACK[(int)" + tmp1 + "]=" + retornoIzq.Value + ";"
					asignacion2 = "STACK[(int)" + tmp2 + "]=" + retornoDer.Value + ";"
					gen.AddCodes(asignacion1, ambito)
					gen.AddCodes(asignacion2, ambito)*/
					gen.NewStack(tmp1, retornoIzq.Value, false, "", name, true, true, p.Line)
					gen.NewStack(tmp2, retornoDer.Value, false, "", name, true, true, p.Line)
					//gen.AddCodes("proc_concatSTR();", ambito)
					gen.NewLlamada("proc_concatSTR", false, "", name, true, false, p.Line)
					tmp3 := gen.NewTemp()
					//code := tmp3 + "=STACK[(int)P];\n"
					//gen.AddCodes(code, ambito)
					gen.NewCallStack(tmp3, "P", false, "", name, true, true, p.Line)
					//gen.AddCodes("//FIN DE CONCATENACION DE STRINGS", ambito)
					gen.NewComentario("FIN DE CONCATENACION DE STRINGS", name, true, false, p.Line)
					return interfaces.Value{Value: tmp3, IsTemp: true, Type: interfaces.STR, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.INTEGER && retornoDer.Type == interfaces.USIZE {
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					gen.NewOperacion(newTemp, retornoIzq.Value, "+", retornoDer.Value, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.USIZE && retornoDer.Type == interfaces.INTEGER {
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
					gen.NewOperacion(newTemp, retornoIzq.Value, "+", retornoDer.Value, false, "", name, true, true, p.Line)
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
					gen.NewOperacion(newTemp, retornoIzq.Value, "-", retornoDer.Value, false, "", name, true, true, p.Line)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {
					gen.NewOperacion(newTemp, retornoIzq.Value, "-", retornoDer.Value, false, "", name, true, true, p.Line)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				}
			} else {
				if retornoIzq.Type == interfaces.INTEGER && retornoDer.Type == interfaces.USIZE {
					gen.NewOperacion(newTemp, retornoIzq.Value, "-", retornoDer.Value, false, "", name, true, true, p.Line)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.USIZE && retornoDer.Type == interfaces.INTEGER {
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
					gen.NewOperacion(newTemp, retornoIzq.Value, "-", retornoDer.Value, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN RESTAR ENTRE SI", p.Line, p.Col)
		}

	case "*":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			fmt.Println(retornoIzq)
			fmt.Println(retornoDer)
			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {
					gen.NewOperacion(newTemp, retornoIzq.Value, "*", retornoDer.Value, false, "", name, true, true, p.Line)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {
					gen.NewOperacion(newTemp, retornoIzq.Value, "*", retornoDer.Value, false, "", name, true, true, p.Line)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				}
			} else {
				if retornoIzq.Type == interfaces.INTEGER && retornoDer.Type == interfaces.USIZE {
					gen.NewOperacion(newTemp, retornoIzq.Value, "*", retornoDer.Value, false, "", name, true, true, p.Line)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoIzq.Type == interfaces.USIZE && retornoDer.Type == interfaces.INTEGER {
					gen.NewOperacion(newTemp, retornoIzq.Value, "*", retornoDer.Value, false, "", name, true, true, p.Line)
					//gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN MULTIPLICAR ENTRE SI", p.Line, p.Col)
		}

	case "/":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == retornoDer.Type {

				if retornoDer.Type == interfaces.INTEGER || retornoDer.Type == interfaces.USIZE {
					Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0", 0, name)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				} else if retornoDer.Type == interfaces.FLOAT {
					Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0.0", 0, name)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoDer.Type, TrueLabel: "", FalseLabel: ""}

				}
			} /*else {
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
			}*/
			env.(environment.Environment).NewError("NO ES POSIBLE DIVIDIR ESTE TIPO DE DATO", p.Line, p.Col)

		}
	case "<":
		{
			//gen.AddCodes("//--------<", ambito)
			gen.NewComentario("--------<", name, true, false, p.Line)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == retornoDer.Type {
				l1 := gen.NewLabel()
				l2 := gen.NewLabel()
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					//gen.AddCodes("if ("+retornoIzq.Value+"<"+retornoDer.Value+") goto "+l1+";", ambito)
					gen.NewIf(retornoIzq.Value, "<", retornoDer.Value, l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				} else if retornoIzq.Type == interfaces.STRING {
					//code := "//COMPARACION DE STRING\n"
					gen.NewComentario("COMPARACION DE STRING", name, true, true, p.Line)
					dir1 := gen.NewTemp() //dirección del primer parametro
					dir2 := gen.NewTemp() //dirección del segundo parametro
					//code += dir1 + "=P+1;\n"
					gen.NewOperacion(dir1, "P", "+", "1", true, "DIRECCION DEL PRIMER PARAMETRO", name, true, true, p.Line)
					//code += dir2 + "=P+2;\n"
					gen.NewOperacion(dir1, "P", "+", "2", true, "DIRECCION DEL SEGUNDO PARAMETRO", name, true, true, p.Line)
					//code += "STACK[(int)" + dir1 + "]=" + retornoIzq.Value + ";\n"
					gen.NewStack(dir1, retornoIzq.Value, false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + dir2 + "]=" + retornoDer.Value + ";\n"
					gen.NewStack(dir2, retornoDer.Value, false, "", name, true, true, p.Line)
					gen.AddFuncExtra("COMPARELONG")
					//code += "compareLong_String();\n"
					gen.NewLlamada("compareLong_String", false, "", name, true, false, p.Line)
					resultado := gen.NewTemp()
					//code += "STACK[(int)P]=" + retornoDer.Value + ";\n"
					//code += resultado + "=STACK[(int)P];\n"
					gen.NewCallStack(resultado, "P", false, "", name, true, true, p.Line)
					//gen.AddCodes(code, ambito)
					//gen.AddCodes("if ("+resultado+"==0) goto "+l1+";", ambito)
					gen.NewIf(resultado, "==", "0", l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				}
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}

	case ">":
		{
			//gen.AddCodes("//-------->", ambito)
			gen.NewComentario("-------->", name, true, false, p.Line)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == retornoDer.Type {
				l1 := gen.NewLabel()
				l2 := gen.NewLabel()
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					//gen.AddCodes("if ("+retornoIzq.Value+">"+retornoDer.Value+") goto "+l1+";", ambito)
					gen.NewIf(retornoIzq.Value, ">", retornoDer.Value, l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				} else if retornoIzq.Type == interfaces.STRING {
					//code := "//COMPARACION DE STRING\n"
					gen.NewComentario("COMPARACION DE STRING", name, true, false, p.Line)
					dir1 := gen.NewTemp() //dirección del primer parametro
					dir2 := gen.NewTemp() //dirección del segundo parametro
					//code += dir1 + "=P+1;\n"
					gen.NewOperacion(dir1, "P", "+", "1", false, "", name, true, true, p.Line)
					//code += dir2 + "=P+2;\n"
					gen.NewOperacion(dir2, "P", "+", "2", false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + dir1 + "]=" + retornoIzq.Value + ";\n"
					gen.NewStack(dir1, retornoIzq.Value, false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + dir2 + "]=" + retornoDer.Value + ";\n"
					gen.NewStack(dir2, retornoDer.Value, false, "", name, true, true, p.Line)
					gen.AddFuncExtra("COMPARELONG")
					//code += "compareLong_String();\n"
					gen.NewLlamada("compareLong_String", false, "", name, true, false, p.Line)
					resultado := gen.NewTemp()
					//code += resultado + "=STACK[(int)P];\n"
					gen.NewCallStack(resultado, "P", false, "", name, true, true, p.Line)
					//gen.AddCodes(code, ambito)
					//gen.AddCodes("if ("+resultado+"==2) goto "+l1+";", ambito)
					gen.NewIf(resultado, "==", "2", l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				}
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}
	case ">=":
		{
			//			gen.AddCodes("//-------->=", ambito)
			gen.NewComentario("-------->=", name, true, false, p.Line)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == retornoDer.Type {
				l1 := gen.NewLabel()
				l2 := gen.NewLabel()
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					//gen.AddCodes("if ("+retornoIzq.Value+">="+retornoDer.Value+") goto "+l1+";", ambito)
					gen.NewIf(retornoIzq.Value, ">=", retornoDer.Value, l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				} else if retornoIzq.Type == interfaces.STRING {
					//code := "//COMPARACION DE STRING\n"
					gen.NewComentario("COMPARACION DE STRING", name, true, false, p.Line)
					dir1 := gen.NewTemp() //dirección del primer parametro
					dir2 := gen.NewTemp() //dirección del segundo parametro
					//code += dir1 + "=P+1;\n"
					gen.NewOperacion(dir1, "P", "+", "1", false, "", name, true, true, p.Line)
					//code += dir2 + "=P+2;\n"
					gen.NewOperacion(dir2, "P", "+", "2", false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + dir1 + "]=" + retornoIzq.Value + ";\n"
					gen.NewStack(dir1, retornoIzq.Value, false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + dir2 + "]=" + retornoDer.Value + ";\n"
					gen.NewStack(dir2, retornoDer.Value, false, "", name, true, true, p.Line)
					gen.AddFuncExtra("COMPARELONG")
					//code += "compareLong_String();\n"
					gen.NewLlamada("compareLong_String", false, "", name, true, false, p.Line)
					resultado := gen.NewTemp()
					//code += resultado + "=STACK[(int)P];\n"
					gen.NewCallStack(resultado, "P", false, "", name, true, true, p.Line)
					//gen.AddCodes(code, ambito)
					//gen.AddCodes("if ("+resultado+">=1) goto "+l1+";", ambito)
					gen.NewIf(resultado, ">=", "1", l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					//addCodeBool(gen, ">=", "1", resultado, ambito)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				}
			}
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}
	case "<=":
		{
			//gen.AddCodes("//--------<=", ambito)
			gen.NewComentario("--------<=", name, true, false, p.Line)
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoDer.Type == retornoIzq.Type {
				l1 := gen.NewLabel()
				l2 := gen.NewLabel()
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					//gen.AddCodes("if ("+retornoIzq.Value+"<="+retornoDer.Value+") goto "+l1+";", ambito)
					gen.NewIf(retornoIzq.Value, "<=", retornoDer.Value, l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				} else if interfaces.STR == retornoDer.Type || interfaces.STRING == retornoIzq.Type {
					//code := "//COMPARACION DE STRING\n"
					gen.NewComentario("COMPARACION DE STRING", name, true, false, p.Line)
					dir1 := gen.NewTemp() //dirección del primer parametro
					dir2 := gen.NewTemp() //dirección del segundo parametro
					//code += dir1 + "=P+1;\n"
					gen.NewOperacion(dir1, "P", "+", "1", false, "", name, true, true, p.Line)
					//code += dir2 + "=P+2;\n"
					gen.NewOperacion(dir2, "P", "+", "2", false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + dir1 + "]=" + retornoIzq.Value + ";\n"
					gen.NewStack(dir1, retornoIzq.Value, false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + dir2 + "]=" + retornoDer.Value + ";\n"
					gen.NewStack(dir2, retornoDer.Value, false, "", name, true, true, p.Line)
					gen.AddFuncExtra("COMPARELONG")
					//code += "compareLong_String();\n"
					gen.NewLlamada("compareLong_String", false, "", name, true, false, p.Line)
					resultado := gen.NewTemp()
					//code += "STACK[(int)P]=" + retornoDer.Value + ";\n"
					//code += resultado + "=STACK[(int)P];\n"
					gen.NewCallStack(resultado, "P", false, "", name, true, true, p.Line)
					//gen.AddCodes(code, ambito)
					//gen.AddCodes("if ("+resultado+"<=1) goto "+l1+";", ambito)
					gen.NewIf(resultado, "<=", "1", l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					//addCodeBool(gen, "<=", "1", resultado, ambito)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
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
			//gen.AddCodes("//--------!=", ambito)
			gen.NewComentario("-------------!=--------", name, true, false, p.Line)
			resultado := ""
			resultado2 := ""
			retornoIzq = p.Op1.Ejecutar(env, gen)
			if retornoIzq.Type == interfaces.BOOLEAN {
				salida := gen.NewLabel()
				resultado = gen.NewTemp()
				//gen.AddCodes(retornoIzq.TrueLabel+":", ambito)
				gen.NewLabels(retornoIzq.TrueLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes(resultado+"=1;", ambito)
				gen.NewAsignacion(resultado, "1", false, "", name, true, true, p.Line)
				//gen.AddCodes("goto "+salida+";", ambito)
				gen.NewSalto(salida, false, "", name, true, true, p.Line)
				//gen.AddCodes(retornoIzq.FalseLabel+":", ambito)
				gen.NewLabels(retornoIzq.FalseLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes(resultado+"=0;", ambito)
				gen.NewAsignacion(resultado, "0", false, "", name, true, true, p.Line)
				//gen.AddCodes(salida+":", ambito)
				gen.NewLabels(salida, false, "", name, true, true, p.Line)
			}
			retornoDer = p.Op2.Ejecutar(env, gen)
			if retornoDer.Type == interfaces.BOOLEAN {
				salida := gen.NewLabel()
				resultado2 = gen.NewTemp()
				//gen.AddCodes(retornoDer.TrueLabel+":", ambito)
				gen.NewLabels(retornoDer.TrueLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes(resultado2+"=1;", ambito)
				gen.NewAsignacion(resultado2, "1", false, "", name, true, true, p.Line)
				//gen.AddCodes("goto "+salida+";", ambito)
				gen.NewSalto(salida, false, "", name, true, true, p.Line)
				//gen.AddCodes(retornoDer.FalseLabel+":", ambito)
				gen.NewLabels(retornoDer.FalseLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes(resultado2+"=0;", ambito)
				gen.NewAsignacion(resultado2, "0", false, "", name, true, true, p.Line)
				//gen.AddCodes(salida+":", ambito)
				gen.NewLabels(salida, false, "", name, true, true, p.Line)
			}
			/*fmt.Println(resultado)
			fmt.Println(resultado2)*/
			if retornoDer.Type == retornoIzq.Type {
				l1 := gen.NewLabel()
				l2 := gen.NewLabel()
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					//gen.AddCodes("if ("+retornoIzq.Value+"!="+retornoDer.Value+") goto "+l1+";", ambito)
					gen.NewIf(retornoIzq.Value, "!=", retornoDer.Value, l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				} else if retornoDer.Type == interfaces.STRING || retornoIzq.Type == interfaces.STR {
					//code := "//COMPARACION DE STRING\n"
					gen.NewComentario("COMPARACION DE STRING", name, true, false, p.Line)
					t1 := gen.NewTemp()
					t4 := gen.NewTemp()
					t2 := retornoIzq.Value
					t3 := retornoDer.Value
					resultado := gen.NewTemp()
					//code += t1 + "=P+1;\n"
					gen.NewOperacion(t1, "P", "+", "1", false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + t1 + "]=" + t2 + ";\n"
					gen.NewStack(t1, t2, false, "", name, true, true, p.Line)
					//code += t4 + "=P+2;\n"
					gen.NewOperacion(t4, "P", "+", "2", false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + t4 + "]=" + t3 + ";\n"
					gen.NewStack(t4, t3, false, "", name, true, true, p.Line)
					gen.AddFuncExtra("COMPARESTR")
					//code += "proc_compareString();\n"
					gen.NewLlamada("proc_compareString", false, "", name, true, false, p.Line)
					//code += resultado + "=STACK[(int)P];"
					gen.NewCallStack(resultado, "P", false, "", name, true, true, p.Line)
					//gen.AddCodes(code, ambito)
					//gen.AddCodes("if ("+resultado+"==0) goto "+l1+";", ambito)
					gen.NewIf(resultado, "==", "0", l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				} else if interfaces.BOOLEAN == retornoIzq.Type {
					//code := "if(" + resultado + "!=" + resultado2 + ") goto " + l1 + ";\n"
					gen.NewIf(resultado, "!=", resultado2, l1, false, "", name, true, true, p.Line)
					//code += "goto " + l2 + ";"
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				}
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}
		}
	case "==":
		{
			gen.NewComentario("-------------==--------", name, true, false, p.Line)
			resultado := ""
			resultado2 := ""
			retornoIzq = p.Op1.Ejecutar(env, gen)
			if retornoIzq.Type == interfaces.BOOLEAN {
				salida := gen.NewLabel()
				resultado = gen.NewTemp()
				//gen.AddCodes(retornoIzq.TrueLabel+":", ambito)
				gen.NewLabels(retornoIzq.TrueLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes(resultado+"=1;", ambito)
				gen.NewAsignacion(resultado, "1", false, "", name, true, true, p.Line)
				//gen.AddCodes("goto "+salida+";", ambito)
				gen.NewSalto(salida, false, "", name, true, true, p.Line)
				//gen.AddCodes(retornoIzq.FalseLabel+":", ambito)
				gen.NewLabels(retornoIzq.FalseLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes(resultado+"=0;", ambito)
				gen.NewAsignacion(resultado, "0", false, "", name, true, true, p.Line)
				//gen.AddCodes(salida+":", ambito)
				gen.NewLabels(salida, false, "", name, true, true, p.Line)
			}
			retornoDer = p.Op2.Ejecutar(env, gen)
			if retornoDer.Type == interfaces.BOOLEAN {
				salida := gen.NewLabel()
				resultado2 = gen.NewTemp()
				//gen.AddCodes(retornoDer.TrueLabel+":", ambito)
				gen.NewLabels(retornoDer.TrueLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes(resultado2+"=1;", ambito)
				gen.NewAsignacion(resultado2, "1", false, "", name, true, true, p.Line)
				//gen.AddCodes("goto "+salida+";", ambito)
				gen.NewSalto(salida, false, "", name, true, true, p.Line)
				//gen.AddCodes(retornoDer.FalseLabel+":", ambito)
				gen.NewLabels(retornoDer.FalseLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes(resultado2+"=0;", ambito)
				gen.NewAsignacion(resultado2, "0", false, "", name, true, true, p.Line)
				//gen.AddCodes(salida+":", ambito)
				gen.NewLabels(salida, false, "", name, true, true, p.Line)
			}
			/*fmt.Println(resultado)
			fmt.Println(resultado2)*/
			if retornoDer.Type == retornoIzq.Type {
				l1 := gen.NewLabel()
				l2 := gen.NewLabel()
				if retornoIzq.Type == interfaces.INTEGER || retornoIzq.Type == interfaces.FLOAT {
					//gen.AddCodes("if ("+retornoIzq.Value+"=="+retornoDer.Value+") goto "+l1+";", ambito)
					gen.NewIf(retornoIzq.Value, "==", retornoDer.Value, l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				} else if retornoDer.Type == interfaces.STRING || retornoIzq.Type == interfaces.STR {
					//code := "//COMPARACION DE STRING\n"
					gen.NewComentario("COMPARACION DE STRING", name, true, false, p.Line)
					t1 := gen.NewTemp()
					t4 := gen.NewTemp()
					t2 := retornoIzq.Value
					t3 := retornoDer.Value
					resultado := gen.NewTemp()
					//code += t1 + "=P+1;\n"
					gen.NewOperacion(t1, "P", "+", "1", false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + t1 + "]=" + t2 + ";\n"
					gen.NewStack(t1, t2, false, "", name, true, true, p.Line)
					//code += t4 + "=P+2;\n"
					gen.NewOperacion(t4, "P", "+", "2", false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + t4 + "]=" + t3 + ";\n"
					gen.NewStack(t4, t3, false, "", name, true, true, p.Line)
					gen.AddFuncExtra("COMPARESTR")
					//code += "proc_compareString();\n"
					gen.NewLlamada("proc_compareString", false, "", name, true, false, p.Line)
					//code += resultado + "=STACK[(int)P];"
					gen.NewCallStack(resultado, "P", false, "", name, true, true, p.Line)
					//gen.AddCodes(code, ambito)
					//gen.AddCodes("if ("+resultado+"==1) goto "+l1+";", ambito)
					gen.NewIf(resultado, "==", "1", l1, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+l2+";", ambito)
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				} else if interfaces.BOOLEAN == retornoIzq.Type {
					//code := "if(" + resultado + "==" + resultado2 + ") goto " + l1 + ";\n"
					gen.NewIf(resultado, "==", resultado2, l1, false, "", name, true, true, p.Line)
					//code += "goto " + l2 + ";"
					gen.NewSalto(l2, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: l1, FalseLabel: l2}
				}
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}
		}
	case "||":
		{
			//gen.AddCodes("//--------||", ambito)
			gen.NewComentario("-------------||--------", name, true, false, p.Line)
			retornoIzq := p.Op1.Ejecutar(env, gen)
			if retornoIzq.Type == interfaces.BOOLEAN {
				//gen.AddCodes(retornoIzq.FalseLabel+":", ambito)
				gen.NewLabels(retornoIzq.FalseLabel, false, "", name, true, true, p.Line)
				retornoDer := p.Op2.Ejecutar(env, gen)
				//gen.AddCodes(retornoDer.TrueLabel+":", ambito)
				gen.NewLabels(retornoDer.TrueLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes("goto "+retornoIzq.TrueLabel+";", ambito)
				gen.NewSalto(retornoIzq.TrueLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes("goto "+retornoIzq.TrueLabel+";", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: false, Type: interfaces.BOOLEAN, TrueLabel: retornoIzq.TrueLabel, FalseLabel: retornoDer.FalseLabel}
			} else {
				env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
			}

		}
	case "&&":
		{
			//gen.AddCodes("//--------&&", ambito)
			gen.NewComentario("--------&&", name, true, false, p.Line)
			//gen.AddCodes("//--------RetornoIZQ", ambito)
			gen.NewComentario("--------RetornoIZQ", name, true, false, p.Line)
			retornoIzq := p.Op1.Ejecutar(env, gen)
			if retornoIzq.Type == interfaces.BOOLEAN {
				//gen.AddCodes(retornoIzq.TrueLabel+":", ambito)
				gen.NewLabels(retornoIzq.TrueLabel, false, "", name, true, true, p.Line)
				//gen.AddCodes("//--------RetornoDER", ambito)
				gen.NewComentario("--------RetornoDER", name, true, false, p.Line)
				retornoDer := p.Op2.Ejecutar(env, gen)
				if retornoDer.Type == interfaces.BOOLEAN {
					/*gen.AddCodes(retornoDer.TrueLabel+":", ambito)
					gen.AddCodes("goto "+retornoIzq.TrueLabel+";", ambito)*/
					//gen.AddCodes(retornoDer.FalseLabel+":", ambito)
					gen.NewLabels(retornoDer.FalseLabel, false, "", name, true, true, p.Line)
					//gen.AddCodes("goto "+retornoIzq.FalseLabel+";", ambito)
					gen.NewSalto(retornoIzq.FalseLabel, false, "", name, true, true, p.Line)
					return interfaces.Value{Value: newTemp, IsTemp: false, Type: retornoIzq.Type, TrueLabel: retornoDer.TrueLabel, FalseLabel: retornoIzq.FalseLabel}
				}
			}
			//gen.AddCodes(retornoIzq.FalseLabel+":", ambito)
			gen.NewLabels(retornoIzq.FalseLabel, false, "", name, true, true, p.Line)
			env.(environment.Environment).NewError("LOS TIPOS NO SE PUEDEN OPERAR ENTRE SI", p.Line, p.Col)
		}
	case "!":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == interfaces.BOOLEAN {
				//gen.InvertirLabels()
				aux := retornoIzq.TrueLabel
				retornoIzq.TrueLabel = retornoIzq.FalseLabel
				retornoIzq.FalseLabel = aux
				return retornoIzq
			} else {
				env.(environment.Environment).NewError("SE ESPERABA UNA EXPRESION BOOLEANA", p.Line, p.Col)
			}
		}
	case "°":
		{
			retornoIzq, retornoDer = DevTipos(p, env, gen)
			if retornoIzq.Type == interfaces.INTEGER {
				//tmp:=gen.NewTemp()
				//gen.AddExpression(newTemp, retornoIzq.Value, "-1", "*", ambito)
				gen.NewOperacion(newTemp, "0", "-", retornoIzq.Value, false, "", name, true, true, p.Line)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
			} else if retornoIzq.Type == interfaces.FLOAT {
				//gen.AddExpression(newTemp, retornoIzq.Value, "-1", "*", ambito)
				gen.NewOperacion(newTemp, "0", "-", retornoIzq.Value, false, "", name, true, true, p.Line)
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
					//code := "//INICIO DE CALCULO DE POTENCIAS\n"
					gen.NewComentario("INICIO DE CALCULO DE POTENCIAS", name, true, false, p.Line)
					//code += newTemp + "=P+1;\n"
					gen.NewOperacion(newTemp, "P", "+", "1", false, "", name, true, true, p.Line)
					//code += tmp2 + "=P+2;\n"
					gen.NewOperacion(tmp2, "P", "+", "2", false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + newTemp + "]=" + retornoIzq.Value + ";\n"
					gen.NewStack(newTemp, retornoIzq.Value, false, "", name, true, true, p.Line)
					//code += "STACK[(int)" + tmp2 + "]=" + retornoDer.Value + ";\n"
					gen.NewStack(tmp2, retornoDer.Value, false, "", name, true, true, p.Line)
					gen.AddFuncExtra("POW")
					//code += "proc_potencia();\n"
					gen.NewLlamada("proc_potencia", false, "", name, true, false, p.Line)
					//code += resultado + "=STACK[(int)P];"
					gen.NewCallStack(resultado, "P", false, "", name, true, false, p.Line)
					//gen.AddCodes(code, ambito)
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
					Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0", 1, name)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				} else if retornoDer.Type == interfaces.FLOAT {
					Comprobar_Div(gen, retornoIzq, retornoDer, newTemp, "0.0", 1, name)
					return interfaces.Value{Value: newTemp, IsTemp: true, Type: retornoIzq.Type, TrueLabel: "", FalseLabel: ""}
				}
			}
			env.(environment.Environment).NewError("NO ES POSIBLE OBTENER EL MODULO DE ESTE TIPO DATO", p.Line, p.Col)
		}
	}

	return interfaces.Value{Value: "0", Type: interfaces.NULL, IsTemp: false, TrueLabel: "", FalseLabel: ""}
}

func Comprobar_Div(g *generator.Generator, izq, der interfaces.Value, tmp, value string, tipo int, name string) {
	l1 := g.NewLabel()
	l2 := g.NewLabel()
	//code := "if(" + der.Value + "!=0) goto " + l1 + ";\n"
	g.NewIf(der.Value, "!=", "0", l1, false, "", name, true, true, "")
	g.AddFuncExtra("DIVZERO")
	//code += "proc_divZero();\n"
	g.NewLlamada("proc_divZero", false, "", name, true, true, "")
	//code += tmp + "=" + value + ";\n"
	g.NewAsignacion(tmp, value, false, "", name, true, true, "")
	//code += "goto " + l2 + ";\n"
	g.NewSalto(l2, false, "", name, true, true, "")
	//code += l1 + ":\n"
	g.NewLabels(l1, false, "", name, true, true, "")
	if tipo == 0 {
		//code += tmp + "=" + izq.Value + "/" + der.Value + ";\n"
		g.NewOperacion(tmp, izq.Value, "/", der.Value, false, "", name, true, true, "")
	} else {
		//code += tmp + "=fmod(" + izq.Value + "," + der.Value + ");\n"
		g.NewFmod(tmp, izq.Value, der.Value, false, "", name)
	}
	//code += l2 + ":"
	g.NewLabels(l2, false, "", name, true, true, "")
}

/*func True_NotTrue(signo string, izq interfaces.Value, g *generator.Generator) (string, string) {
	anterior := g.GetTempsB()
	fmt.Println(anterior)
	tmpIzq := g.NewTemp()
	//tmpDer:=g.NewTemp()
	l1 := anterior.TrueL
	l2 := anterior.FalseL
	l3 := g.NewLabel() //salida de la primera asignacion
	//l4:=anterior.TrueL1
	//l5:=g.NewLabel()//FALSE 2
	value := "//INICIO DE " + signo + "\n" + l1 + ":\n"
	value += tmpIzq + "=1;\ngoto " + l3 + ";\n"
	value += l2 + ":\n"
	value += tmpIzq + "=0;\n"
	value += l3 + ":\n"

	return value, tmpIzq
}*/

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

/*func addCodeBool(gen *generator.Generator, signo, value, temp string, ambito bool) {
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
}*/
