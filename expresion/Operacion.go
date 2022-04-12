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
	suma_resta_dominante := [5][5]interfaces.TipoExpresion{
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
	}
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
	var dominante interfaces.TipoExpresion

	newTemp := gen.NewTemp()

	switch p.Operador {
	case "+":
		{

			dominante = suma_resta_dominante[retornoIzq.Type][retornoDer.Type]

			if dominante == interfaces.INTEGER {

				gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: dominante, TrueLabel: "", FalseLabel: ""}

			} else if dominante == interfaces.FLOAT {

				gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "+", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: dominante, TrueLabel: "", FalseLabel: ""}

			} else {
				fmt.Print("ERROR: No es posible sumar")
			}

		}

	case "-":
		{
			dominante = suma_resta_dominante[retornoIzq.Type][retornoDer.Type]

			if dominante == interfaces.INTEGER {

				gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: dominante, TrueLabel: "", FalseLabel: ""}

			} else if dominante == interfaces.FLOAT {
				gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "-", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: dominante, TrueLabel: "", FalseLabel: ""}

			} else {
				fmt.Print("ERROR: No es posible restar")
			}
		}

	case "*":
		{
			dominante = multi_division_dominante[retornoIzq.Type][retornoDer.Type]

			if dominante == interfaces.INTEGER {
				gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: dominante, TrueLabel: "", FalseLabel: ""}

			} else if dominante == interfaces.FLOAT {
				gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "*", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: dominante, TrueLabel: "", FalseLabel: ""}

			} else {
				fmt.Print("ERROR: No es posible Multiplicar")
			}

		}

	case "/":
		{
			dominante = multi_division_dominante[retornoIzq.Type][retornoDer.Type]

			if dominante == interfaces.INTEGER {
				gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "/", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: dominante, TrueLabel: "", FalseLabel: ""}

			} else if dominante == interfaces.FLOAT {
				gen.AddExpression(newTemp, retornoIzq.Value, retornoDer.Value, "/", ambito)
				return interfaces.Value{Value: newTemp, IsTemp: true, Type: dominante, TrueLabel: "", FalseLabel: ""}

			} else {
				t := time.Now()
				fecha := fmt.Sprintf("%d-%02d-%02dT%02d:%02d:%02d",
					t.Year(), t.Month(), t.Day(),
					t.Hour(), t.Minute(), t.Second())
				err := interfaces.Errores{Line: p.Line, Col: p.Col, Mess: "LOS TIPOS NO SE PUEDEN DIVIDIR ENTRE SI", Fecha: fecha}
				env.(environment.Environment).Errores(err)
			}

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
						value := "if (" + retornoIzq.Value + "<" + retornoDer.Value + ") goto " + l1 + ";\n"
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
						value := "if (" + retornoIzq.Value + ">" + retornoDer.Value + ") goto " + l1 + ";\n"
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
						value := "if (" + retornoIzq.Value + ">=" + retornoDer.Value + ") goto " + l1 + ";\n"
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
						value := "if (" + retornoIzq.Value + "<=" + retornoDer.Value + ") goto " + l1 + ";\n"
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
						value := "if (" + retornoIzq.Value + "!=" + retornoDer.Value + ") goto " + l1 + ";\n"
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
	case "==":
		{
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
					value := "if (" + retornoIzq.Value + "==" + retornoDer.Value + ") goto " + l1 + ";\n"
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
	case "||":
		{
			if retornoIzq.Type == retornoDer.Type && retornoDer.Type == interfaces.BOOLEAN {
				anterior := gen.GetTempsB()
				value := anterior.FalseL + ":\n"
				value += anterior.FalseL1
				value += "goto " + anterior.FalseL + ";\n"
				gen.AddCodes(value, ambito)
				gen.RotarLabels()
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
				fmt.Println("entro-" + anterior.TrueL)
				fmt.Println(anterior.FalseL)
				fmt.Println(anterior.TrueL1)
				fmt.Println(anterior.FalseL1)
				fmt.Println("============")
				value := anterior.TrueL + ":\n"
				value += anterior.FalseL1
				value += "goto " + anterior.FalseL + ";\n"
				gen.AddCodes(value, ambito)
				gen.RotarLabels()
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
	}

	return interfaces.Value{Value: "0", Type: interfaces.NULL, IsTemp: false, TrueLabel: "", FalseLabel: ""}
}
