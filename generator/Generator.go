package generator

import (
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type Generator struct {
	temporal   int
	label      int
	Stack      int
	Heap       int
	code       *arrayList.List
	func_code  *arrayList.List
	tempList   *arrayList.List
	func_extra map[string]bool
	extra_code *arrayList.List
	temp_Bools *temporals
}

type Booleans struct {
	Expresiones *arrayList.List
	Last_Label  string
}

type temporals struct {
	Conf    int
	TrueL   string
	FalseL  string
	TrueL1  string
	FalseL1 string
}

func NewGenerator() *Generator {
	//fmt.Println("neuvo generator")
	generator := Generator{temporal: 0, label: 0, code: arrayList.New(), tempList: arrayList.New(), func_code: arrayList.New(), Stack: 0, Heap: 0, func_extra: make(map[string]bool), extra_code: arrayList.New(), temp_Bools: &temporals{Conf: 0, TrueL: "", FalseL: "", TrueL1: "", FalseL1: ""}}
	return &generator
}

func (g Generator) GetCode() *arrayList.List {

	return g.code
}

func (g Generator) GetTemps() *arrayList.List {

	return g.tempList
}

func (g Generator) GetFuncs() *arrayList.List {

	return g.func_code
}

func (g Generator) GetExtraFuncs() *arrayList.List {

	return g.extra_code
}

func (g Generator) AddTempBool(labelT, labelF string) {
	if g.temp_Bools.Conf == 0 {
		g.temp_Bools.TrueL = labelT
		g.temp_Bools.FalseL = labelF
	} else if g.temp_Bools.Conf == 1 {
		g.temp_Bools.TrueL1 = labelT
		g.temp_Bools.FalseL1 = labelF
	}
}

func (g Generator) GetConf() int {
	return g.temp_Bools.Conf
}

func (g Generator) SetConf() {
	if g.temp_Bools.Conf == 1 {
		g.temp_Bools.Conf = 0
	} else if g.temp_Bools.Conf == 0 {
		g.temp_Bools.Conf = 1
	}
}
func (g Generator) InvertirLabels() {
	aux := g.temp_Bools.FalseL
	g.temp_Bools.FalseL = g.temp_Bools.TrueL
	g.temp_Bools.TrueL = aux
}
func (g Generator) LabelsOr(l2 string) {
	g.temp_Bools.FalseL = l2
}
func (g Generator) RotarLabels() {
	g.temp_Bools.TrueL = g.temp_Bools.TrueL1
}

func (g Generator) GetTempsB() temporals {
	return *g.temp_Bools
}

//Generar un nuevo temporal
func (g *Generator) NewTemp() string {
	temp := "t" + fmt.Sprintf("%v", g.temporal)
	g.temporal = g.temporal + 1
	//Lo guardamos para declararlo
	g.tempList.Add(temp)
	return temp
}

//Generador de label
func (g *Generator) NewLabel() string {
	temp := g.label
	g.label = g.label + 1
	return "L" + fmt.Sprintf("%v", temp)
}

//Añade label al codigo
func (g *Generator) AddLabel(label string) {
	g.code.Add(label + ":")
}

func (g *Generator) AddExpression(target string, left string, right string, operator string, env bool) {
	if env {
		g.func_code.Add(target + " = " + left + " " + operator + " " + right + ";")
	} else {
		g.code.Add(target + " = " + left + " " + operator + " " + right + ";")
	}
}

//Añade un printf
func (g *Generator) AddPrintf(typePrint string, value string, ambito bool) {
	if ambito {
		g.func_code.Add("printf(\"%" + typePrint + "\"," + value + ");")
	} else {
		g.code.Add("printf(\"%" + typePrint + "\"," + value + ");")
	}
}

func (g *Generator) AddCode(code string) {
	g.code.Add(code)
}

func (g *Generator) AddCodes(code string, ambito bool) {
	if !ambito {
		g.code.Add(code)
	} else {
		g.func_code.Add(code)
	}
}

func (g *Generator) AddFuncExtra(id string) {
	if g.func_extra[id] {
		return
	} else {
		if id == "PRINTBOOL" {
			g.func_extra[id] = true
			g.extra_code.Add(print_bools())
		} else if id == "DIVZERO" {
			g.func_extra[id] = true
			g.extra_code.Add(print_divZero())
		} else if id == "NULL" {
			g.func_extra[id] = true
			g.extra_code.Add(print_Null())
		} else if id == "CONCATSTR" {
			g.func_extra[id] = true
			g.extra_code.Add(concat_STR(g))
		} else if id == "PRINTSTR" {
			g.func_extra[id] = true
			g.extra_code.Add(print_String(g))
		} else if id == "POW" {

		}
	}

}

func (g *Generator) AddFunc(code string) { //crea el codigo para las funciones
	g.func_code.Add(code)
}

func (g *Generator) AddIf(left string, right string, operator string, label string) {
	g.code.Add("if(" + left + " " + operator + " " + right + ") goto " + label + ";")
}

func (g *Generator) AddGoto(label string) {
	g.code.Add("goto " + label + ";")
}

func print_bools() string {
	code := "void proc_extra_printTrue(){\n"
	code += "printf(\"%c\",116);\n"
	code += "printf(\"%c\",114);\n"
	code += "printf(\"%c\",117);\n"
	code += "printf(\"%c\",101);\n"
	code += "return;\n"
	code += "}\n\n"
	code += "void proc_extra_printFalse(){\n"
	code += "printf(\"%c\",102);\n"
	code += "printf(\"%c\",97);\n"
	code += "printf(\"%c\",108);\n"
	code += "printf(\"%c\",115);\n"
	code += "printf(\"%c\",101);\n"
	code += "return;\n"
	code += "}\n\n"
	return code
}

func print_Null() string {
	code := "void proc_printNull(){\n"
	code += "printf(\"%c\",110);\n"
	code += "printf(\"%c\",117);\n"
	code += "printf(\"%c\",108);\n"
	code += "printf(\"%c\",108);\n"
	code += "return;\n}"
	return code
}

func concat_STR(gen *Generator) string {
	code := "void proc_concatSTR(){\n"
	tmp1 := gen.NewTemp()
	code += tmp1 + "=H;\n"
	tmp2 := gen.NewTemp()
	code += tmp2 + "=P+1;\n"
	tmp3 := gen.NewTemp()
	code += tmp3 + "=STACK[(int)" + tmp2 + "];\n"
	tmp4 := gen.NewTemp()
	code += tmp4 + "=P+2;\n"
	tmp5 := gen.NewTemp()
	code += tmp5 + "=STACK[(int)" + tmp4 + "];\n"
	l1 := gen.NewLabel()
	l2 := gen.NewLabel()
	l3 := gen.NewLabel()
	code += l1 + ":\n"
	tmp6 := gen.NewTemp()
	code += tmp6 + "=HEAP[(int)" + tmp3 + "];\n"
	code += "if(" + tmp6 + "!=-1) goto " + l3 + ";\n"
	code += "goto " + l2 + ";\n"
	code += l3 + ":\n"
	code += "HEAP[(int)H] =" + tmp6 + ";\n"
	code += "H=H+1;\n"
	code += tmp3 + "=" + tmp3 + "+1;\n"
	code += "goto " + l1 + ";\n"
	code += l2 + ":\n"
	code += tmp6 + "= HEAP[(int)" + tmp5 + "];\n"
	l4 := gen.NewLabel()
	l5 := gen.NewLabel()
	code += "if(" + tmp6 + "!=-1) goto " + l4 + ";\n"
	code += "goto " + l5 + ";\n"
	code += l4 + ":\n"
	code += "HEAP[(int)H] =" + tmp6 + ";\n"
	code += "H=H+1;\n"
	code += tmp5 + "=" + tmp5 + "+1;\n"
	code += "goto " + l2 + ";\n"
	code += l5 + ":\n"
	code += "HEAP[(int)H] = -1;"
	code += "H=H+1;\n"
	code += "STACK[(int)P] = " + tmp1 + ";\n"
	code += "return;\n}"
	return code
}

func print_String(gen *Generator) string {
	code := "void proc_printString(){\n"
	t5 := gen.NewTemp()
	code += t5 + "=P+1;\n"
	t6 := gen.NewTemp()
	l5 := gen.NewLabel()
	code += t6 + "=STACK[(int)" + t5 + "];\n"
	code += l5 + ":\n"
	t7 := gen.NewTemp()
	code += t7 + "=HEAP[(int)" + t6 + "];\n"
	l3 := gen.NewLabel()
	l4 := gen.NewLabel()
	code += "if(" + t7 + "!=-1) goto " + l3 + ";\n"
	code += "goto " + l4 + ";\n"
	code += l3 + ":\n"
	code += "printf(\"%c\",(int)" + t7 + ");\n"
	code += t6 + "=" + t6 + " + 1;\n"
	code += "goto " + l5 + ";\n"
	code += l4 + ":\n"
	code += "return;\n}"
	return code
}

func print_divZero() string {
	code := ""
	return code
}

func (g *Generator) Array_char(str string) []rune {
	runes := []rune(str)
	return runes
}
