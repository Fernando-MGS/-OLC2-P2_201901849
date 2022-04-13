package instruction

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"fmt"
)

type Imprimir struct {
	Expresion interfaces.Expresion
}

func NewImprimir(val interfaces.Expresion) Imprimir {
	exp := Imprimir{val}
	return exp
}

func (p Imprimir) Ejecutar(env interface{}, gen *generator.Generator) interface{} {

	result := p.Expresion.Ejecutar(env, gen)
	fmt.Println(result)
	ambito := true
	if env.(environment.Environment).Control.Id == "GLOBAL" || env.(environment.Environment).Control.Id == "main" {
		ambito = false
	}
	if result.Type == interfaces.INTEGER || result.Type == interfaces.USIZE {
		gen.AddPrintf("d", "(int)"+fmt.Sprintf("%v", result.Value), ambito)
	} else if result.Type == interfaces.FLOAT {
		gen.AddPrintf("f", fmt.Sprintf("%v", result.Value), ambito)
	} else if result.Type == interfaces.CHAR {
		gen.AddPrintf("c", "(int)"+fmt.Sprintf("%v", result.Value), ambito)
	} else if result.Type == interfaces.BOOLEAN {
		value := ""
		l1 := gen.GetTempsB().TrueL  //ln	true
		l2 := gen.GetTempsB().FalseL //ln+1	false
		l3 := gen.NewLabel()         //ln+2	exit
		/*if result.IsTemp {
		value = "if (" + result.Value + "==1) goto " + l1 + ";\n"
		value += "goto " + l2 + ";\n"*/
		value += l1 + ":\n"
		value += "proc_extra_printTrue();\n"
		value += "goto " + l3 + ";\n"
		value += l2 + ":\n"
		value += "proc_extra_printFalse();\n"
		value += l3 + ":\n"
		gen.AddCodes(value, ambito)
		gen.SetConf()
		gen.AddFuncExtra("PRINTBOOL")
	} else if result.Type == interfaces.STR || result.Type == interfaces.STRING {
		l1 := gen.NewLabel()
		l2 := gen.NewLabel()
		code := "if (" + result.Value + "==-1) goto " + l1 + ";\n"
		code += "P=P+1;\n"
		gen.Stack++
		t1 := gen.NewTemp()
		code += t1 + "=P+1;\n"
		code += "STACK[(int)" + t1 + "] =" + result.Value + ";\n"
		code += "proc_printString();\n"
		code += "P=P-1;\n"
		//gen.Stack--
		code += "goto " + l2 + ";\n"
		code += l1 + ":\n"
		code += l2 + ":\n"
		gen.AddCodes(code, ambito)
		gen.AddFuncExtra("PRINTSTR")
	}
	//gen.AddPrintf("c", fmt.Sprintf("%v", result.Value))
	//salto de línea
	gen.AddPrintf("c", "10", ambito)
	var retorno interfaces.Symbol
	retorno.Tipo.Tipo = interfaces.NULL
	return retorno
}
