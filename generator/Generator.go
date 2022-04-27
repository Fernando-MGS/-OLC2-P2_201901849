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
	salida     *arrayList.List
	retornos   *arrayList.List
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
	Cequal  bool
	TrueL   string
	FalseL  string
	TrueL1  string
	FalseL1 string
}

func NewGenerator() *Generator {
	//fmt.Println("neuvo generator")
	generator := Generator{salida: arrayList.New(), temporal: 0, label: 0, code: arrayList.New(), tempList: arrayList.New(), func_code: arrayList.New(), Stack: 0, Heap: 0, func_extra: make(map[string]bool), extra_code: arrayList.New(), temp_Bools: &temporals{Conf: 0, TrueL: "", FalseL: "", TrueL1: "", FalseL1: "", Cequal: false}, retornos: arrayList.New()}
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

func (g *Generator) SetSalida(exit string) {
	//gen := Generator{salida: exit}
	g.salida.Add(exit)
}
func (g Generator) GetSalida() string {
	index := g.salida.Len() - 1
	r1 := fmt.Sprintf("%v", g.salida.GetValue(index))
	/*fmt.Println(r1 + "salida")
	fmt.Println()*/
	return r1
}
func (g Generator) DelSalida() {
	index := g.salida.Len() - 1
	g.salida.RemoveAtIndex(index)
}

func (g *Generator) SetRetorno(retorno string) {
	//gen := Generator{salida: exit}
	g.retornos.Add(retorno)
}
func (g Generator) GetRetorno() string {
	index := g.retornos.Len() - 1
	r1 := fmt.Sprintf("%v", g.retornos.GetValue(index))
	/*fmt.Println(r1 + "salida")
	fmt.Println()*/
	return r1
}
func (g Generator) DelRetorno() {
	index := g.retornos.Len() - 1
	if index >= 0 {
		g.retornos.RemoveAtIndex(index)
	}
}

func (g Generator) GetConf() int {
	return g.temp_Bools.Conf
}

func (g Generator) SetTrueFalse(v, f string) {
	g.temp_Bools.TrueL = v
	g.temp_Bools.FalseL = f
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
func (g Generator) Conf2() {
	g.temp_Bools.Cequal = !g.temp_Bools.Cequal
}
func (g Generator) GetConf2() bool {
	return g.temp_Bools.Cequal
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
			g.func_extra[id] = true
			g.extra_code.Add(pow_num(g))
		} else if id == "COMPARESTR" {
			g.func_extra[id] = true
			g.extra_code.Add(compare_string(g))
		} else if id == "BOUNDS" {
			g.func_extra[id] = true
			g.extra_code.Add(bounds_error())
		} else if id == "COMPARELONG" {
			g.func_extra[id] = true
			g.extra_code.Add(compareLong_Str(g))
		} else if id == "MOD" {
			g.func_extra[id] = true
			g.extra_code.Add(calcularMod(g))
		} else if id == "PRINTINT" {
			g.func_extra[id] = true
			g.extra_code.Add(printArrayNumbers(g))
		} else if id == "PRINTFLOAT" {
			g.func_extra[id] = true
			g.extra_code.Add(printArrayFloat(g))
		} else if id == "PRINTSTRING" {
			g.func_extra[id] = true
			g.extra_code.Add(printArrayString(g))
		} else if id == "PRINTCHAR" {
			g.func_extra[id] = true
			g.extra_code.Add(printArrayChar(g))
		} else if id == "SETINT" {
			g.func_extra[id] = true
			g.extra_code.Add(printSetInt(g))
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

func print_divZero() string {
	code := "void proc_divZero(){\n"
	code += "printf(\"%c\",77);\n"  //m
	code += "printf(\"%c\",97);\n"  //a
	code += "printf(\"%c\",116);\n" //t
	code += "printf(\"%c\",104);\n" //h
	code += "printf(\"%c\",69);\n"  //e
	code += "printf(\"%c\",114);\n" //r
	code += "printf(\"%c\",114);\n" //r
	code += "printf(\"%c\",111);\n" //o
	code += "printf(\"%c\",114);\n" //r
	code += "printf(\"%c\",10);\n"  //r
	code += "return;\n}"
	return code
}

func pow_num(g *Generator) string {
	t0 := g.NewTemp()
	t1 := g.NewTemp()
	t2 := g.NewTemp()
	t3 := g.NewTemp()
	t4 := g.NewTemp()
	l0 := g.NewLabel()
	l1 := g.NewLabel()
	l2 := g.NewLabel()
	code := "void proc_potencia(){\n"
	code += t0 + "=P+1;\n"
	code += t1 + " = STACK[(int)" + t0 + "];\n"
	code += t2 + "= P + 2;\n"
	code += t3 + " = STACK[(int)" + t2 + "];\n"
	code += t4 + "=" + t1 + ";\n"
	code += l2 + ":\n"
	code += "if(" + t3 + " > 1) goto " + l0 + ";\n"
	code += "goto " + l1 + ";\n"
	code += l0 + ":\n"
	code += t4 + " = " + t4 + "*" + t1 + ";\n"
	code += t3 + " = " + t3 + "-1;\n"
	code += "goto " + l2 + ";\n"
	code += l1 + ":\n"
	code += "STACK[(int)P] = " + t4 + ";\n"
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
	gen.Heap++
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
	gen.Heap++
	code += tmp5 + "=" + tmp5 + "+1;\n"
	code += "goto " + l2 + ";\n"
	code += l5 + ":\n"
	code += "HEAP[(int)H] = -1;"
	code += "H=H+1;\n"
	gen.Heap++
	code += "STACK[(int)P] = " + tmp1 + ";\n"
	code += "return;\n}"
	return code
}

func calcularMod(gen *Generator) string {
	code := "void proc_calcularMod(){\n"
	par1 := gen.NewTemp()
	par2 := gen.NewTemp()
	dir1 := gen.NewTemp()
	dir2 := gen.NewTemp()
	pre_res := gen.NewTemp()
	res1 := gen.NewTemp()
	res := gen.NewTemp()
	code += dir1 + "=P+1;\n"
	code += dir2 + "=P+2;\n"
	code += par1 + "=STACK[(int)" + dir1 + "];\n"
	code += par2 + "=STACK[(int)" + dir2 + "];\n"
	code += pre_res + "=" + par1 + "/" + par2 + ";\n"
	code += res1 + "=" + par2 + "*" + "(int)" + pre_res + ";\n"
	code += res + "=" + par1 + "-" + res1 + ";\n"
	code += "STACK[(int)P]=" + res + ";\n"
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

func compare_string(g *Generator) string {
	code := "void proc_compareString(){\n"
	t39 := g.NewTemp()
	t40 := g.NewTemp()
	t41 := g.NewTemp()
	t42 := g.NewTemp()
	t43 := g.NewTemp()
	t44 := g.NewTemp()
	t45 := g.NewTemp()
	l35 := g.NewLabel()
	l36 := g.NewLabel()
	l34 := g.NewLabel()
	l33 := g.NewLabel()
	l37 := g.NewLabel()
	l38 := g.NewLabel()
	l39 := g.NewLabel()
	code += t39 + "=P+1;\n"
	code += t40 + "=STACK[(int)" + t39 + "];\n"
	code += t41 + "=P+2;\n"
	code += t42 + "=STACK[(int)" + t41 + "];\n"
	code += t43 + "=1;\n"
	code += t44 + "=HEAP[(int)" + t40 + "];\n"
	code += t45 + "=HEAP[(int)" + t42 + "];\n"
	code += l35 + ":\n"
	code += "if(" + t44 + "!=-1) goto " + l33 + ";\n"
	code += "goto " + l34 + ";\n"
	code += l33 + ":\n"
	code += "if(" + t44 + "!=" + t45 + ") goto " + l36 + ";\n"
	code += "goto " + l37 + ";\n"
	code += l36 + ":\n"
	code += t43 + "= 0  ;\n"
	code += "goto " + l34 + ";\n"
	code += l37 + ":\n"
	code += t40 + " = " + t40 + " + 1;\n"
	code += t42 + " = " + t42 + " + 1;\n"
	code += t44 + " = HEAP[(int)+" + t40 + "];\n"
	code += t45 + " = HEAP[(int)+" + t42 + "];\n"
	code += "goto " + l35 + ";\n"
	code += l34 + ":\n"
	code += "if(" + t45 + "!=-1) goto " + l38 + ";\n"
	code += "goto " + l39 + ";\n"
	code += l38 + ":\n"
	code += t43 + "=0;\n"
	code += l39 + ":\n"
	code += "STACK[(int)P] =" + t43 + ";\n"
	code += "return ;\n}"
	return code
}

func compareLong_Str(g *Generator) string {
	//conteo del primer string
	code := "void compareLong_String(){\n"
	param1 := g.NewTemp()
	guia1 := g.NewTemp()
	code += param1 + "=P+1;\n"
	code += guia1 + "=STACK[(int)" + param1 + "];\n"
	l1 := g.NewLabel()
	l2 := g.NewLabel()
	l3 := g.NewLabel()
	size1 := g.NewTemp()
	code += size1 + "=0;\n"
	code += l1 + ":\n"
	chr := g.NewTemp()
	code += chr + "=HEAP[(int)" + guia1 + "];\n"
	code += "if (" + chr + "!=-1) goto " + l2 + ";\n"
	code += "goto " + l3 + ";\n"
	code += l2 + ":\n"
	code += guia1 + "=" + guia1 + "+1;\n"
	code += size1 + "=" + size1 + "+1;\n"
	code += "goto " + l1 + ";\n"
	code += l3 + ":\n"
	//conteo del segundo string
	param2 := g.NewTemp()
	guia2 := g.NewTemp()
	code += param2 + "=P+2;\n"
	code += guia2 + "=STACK[(int)" + param2 + "];\n"
	l4 := g.NewLabel()
	l5 := g.NewLabel()
	l6 := g.NewLabel()
	size2 := g.NewTemp()
	code += size2 + "=0;\n"
	code += l4 + ":\n"
	chr1 := g.NewTemp()
	code += chr1 + "=HEAP[(int)" + guia2 + "];\n"
	code += "if (" + chr1 + "!=-1) goto " + l5 + ";\n"
	code += "goto " + l6 + ";\n"
	code += l5 + ":\n"
	code += guia2 + "=" + guia2 + "+1;\n"
	code += size2 + "=" + size2 + "+1;\n"
	code += "goto " + l4 + ";\n"
	code += l6 + ":\n"
	l7 := g.NewLabel()
	l8 := g.NewLabel()
	l9 := g.NewLabel()
	//l10 := g.NewLabel()
	l11 := g.NewLabel()
	code += "if(" + size1 + "==" + size2 + ") goto " + l7 + ";\n"
	code += "if(" + size1 + ">" + size2 + ") goto " + l8 + ";\n"
	code += "if(" + size1 + "<" + size2 + ") goto " + l9 + ";\n"
	//code += "if(" + size1 + "<=" + size2 + ") goto " + l10 + ";\n"
	code += l7 + ":\n"
	code += "STACK[(int)P] = 1;\n"
	code += "goto " + l11 + ";\n"
	code += l8 + ":\n"
	code += "STACK[(int)P] = 2;\n"
	code += "goto " + l11 + ";\n"
	code += l9 + ":\n"
	code += "STACK[(int)P] = 0;\n"
	/*code += "goto " + l11 + ";\n"
	/*code += l10 + ":\n"
	code += "STACK[(int)P] = 3;\n"*/
	code += l11 + ":\n"
	code += "return ;\n}"
	return code
}

func printArrayNumbers(gen *Generator) string {
	code := "void proc_printSetNumbers(){\n"
	code += "printf(\"%c\",91);\n" //[
	param := gen.NewTemp()
	param2 := gen.NewTemp()
	i := gen.NewTemp()
	iterador := gen.NewTemp()
	contador := gen.NewTemp()
	entrada := gen.NewLabel()
	num := gen.NewTemp()
	salida := gen.NewLabel()
	code += param + "=P+1;\n"
	code += param2 + "=P+2;\n"
	code += iterador + "=0;\n"
	code += contador + "=STACK[(int)" + param + "];\n"
	code += i + "=STACK[(int)" + param2 + "];\n"
	code += entrada + ":\n"
	code += "if(" + iterador + "==" + i + ") goto " + salida + ";\n"
	//gen.AddCodes("",am)
	code += num + "=HEAP[(int)" + contador + "];\n"
	code += "printf(\"%d\",(int)" + num + ");\n" //[
	code += contador + "=" + contador + "+1;\n"
	code += "printf(\"%c\",44);\n" //]
	code += iterador + "=" + iterador + "+1;\n"
	code += "goto " + entrada + ";\n"
	code += salida + ":\n"
	code += "printf(\"%c\",93);\n" //]
	code += "return ;}"
	return code
}

func printArrayFloat(gen *Generator) string {
	code := "void proc_printSetFloat(){\n"
	code += "printf(\"%c\",91);\n" //[
	param := gen.NewTemp()
	param2 := gen.NewTemp()
	i := gen.NewTemp()
	iterador := gen.NewTemp()
	contador := gen.NewTemp()
	entrada := gen.NewLabel()
	num := gen.NewTemp()
	salida := gen.NewLabel()
	code += param + "=P+1;\n"
	code += param2 + "=P+2;\n"
	code += iterador + "=0;\n"
	code += contador + "=STACK[(int)" + param + "];\n"
	code += i + "=STACK[(int)" + param2 + "];\n"
	code += entrada + ":\n"
	code += "if(" + iterador + "==" + i + ") goto " + salida + ";\n"
	//gen.AddCodes("",am)
	code += num + "=HEAP[(int)" + contador + "];\n"
	code += "printf(\"%f\"," + num + ");\n" //[
	code += contador + "=" + contador + "+1;\n"
	code += "printf(\"%c\",44);\n" //]
	code += iterador + "=" + iterador + "+1;\n"
	code += "goto " + entrada + ";\n"
	code += salida + ":\n"
	code += "printf(\"%c\",93);\n" //]
	code += "return ;}"
	return code
}

func printArrayString(gen *Generator) string {
	code := "void proc_printSetString(){\n"
	code += "printf(\"%c\",91);\n" //[
	param := gen.NewTemp()
	param2 := gen.NewTemp()
	i := gen.NewTemp()
	iterador := gen.NewTemp()
	contador := gen.NewTemp()
	entrada := gen.NewLabel()
	num := gen.NewTemp()
	param3 := gen.NewTemp()
	salida := gen.NewLabel()
	code += param + "=P+1;\n"
	code += param2 + "=P+2;\n"
	code += iterador + "=0;\n"
	code += contador + "=STACK[(int)" + param + "];\n"
	code += i + "=STACK[(int)" + param2 + "];\n"
	code += entrada + ":\n"
	code += "if(" + iterador + "==" + i + ") goto " + salida + ";\n"
	//gen.AddCodes("",am)
	code += num + "=HEAP[(int)" + contador + "];\n"
	//code += "printf(\"%f\"," + num + ");\n" //[
	code += param3 + "=P+1;\n"
	code += "printf(\"%c\",34);\n" //]
	code += "STACK[(int)" + param3 + "]=" + num + ";\n"
	gen.AddFuncExtra("PRINTSTR")
	code += "proc_printString();\n"
	code += "printf(\"%c\",34);\n" //]
	code += contador + "=" + contador + "+1;\n"
	code += "printf(\"%c\",44);\n" //]
	code += iterador + "=" + iterador + "+1;\n"
	code += "goto " + entrada + ";\n"
	code += salida + ":\n"
	code += "printf(\"%c\",93);\n" //]
	code += "return ;}"
	return code
}

func printArrayChar(gen *Generator) string {
	code := "void proc_printSetChar(){\n"
	code += "printf(\"%c\",91);\n" //[
	param := gen.NewTemp()
	param2 := gen.NewTemp()
	i := gen.NewTemp()
	iterador := gen.NewTemp()
	contador := gen.NewTemp()
	entrada := gen.NewLabel()
	num := gen.NewTemp()
	salida := gen.NewLabel()
	code += param + "=P+1;\n"
	code += param2 + "=P+2;\n"
	code += iterador + "=0;\n"
	code += contador + "=STACK[(int)" + param + "];\n"
	code += i + "=STACK[(int)" + param2 + "];\n"
	code += entrada + ":\n"
	code += "if(" + iterador + "==" + i + ") goto " + salida + ";\n"
	//gen.AddCodes("",am)
	code += num + "=HEAP[(int)" + contador + "];\n"
	code += "printf(\"%c\"," + num + ");\n" //[
	code += contador + "=" + contador + "+1;\n"
	code += "printf(\"%c\",44);\n" //]
	code += iterador + "=" + iterador + "+1;\n"
	code += "goto " + entrada + ";\n"
	code += salida + ":\n"
	code += "printf(\"%c\",93);\n" //]
	code += "return ;}"
	return code
}

func printSetInt(gen *Generator) string {
	code := "void proc_printInt(){\n"
	index_p1 := gen.NewTemp()
	index_p2 := gen.NewTemp()
	param := gen.NewTemp()  //tipo: si es 0 entonces empieza a imprimir datos, si es mayor entonces recursiva de procInt
	param2 := gen.NewTemp() //la posicion en el heap
	tipo := gen.NewTemp()   //temporal que se usara para comparar el tipo
	longitud := gen.NewTemp()
	iterador := gen.NewTemp()
	valor := gen.NewTemp()
	pos := gen.NewTemp()
	entrada := gen.NewLabel()
	salida := gen.NewLabel()
	tipo1 := gen.NewLabel()
	tipo2 := gen.NewLabel()
	size_return := gen.NewTemp()
	//retorno:=gen.NewTemp()
	index_ret := gen.NewTemp()
	//pos_ret:=gen.NewTemp()
	code += index_p1 + "=P+1;//tipo\n"
	code += index_p2 + "=P+2;//posicion en heap\n"
	code += param + "=STACK[(int)" + index_p1 + "];//TIPO\n"
	code += param2 + "=STACK[(int)" + index_p2 + "];//POSICION EN HEAP\n"
	code += longitud + "=HEAP[(int)" + param2 + "];\n"
	code += pos + "=" + param2 + "+1;//POSICION DEL PRIMER ELEMENTO DEL VECTOR\n"
	code += tipo + "=" + param + "-1;//DISMINUCION DEL TIPO\n"
	code += iterador + "=0;\n"
	code += "printf(\"%c\",91);\n" //]
	code += entrada + ":\n"

	code += "if(" + iterador + ">=" + longitud + ") goto " + salida + ";\n"
	code += valor + "=HEAP[(int)" + pos + "];\n"

	//comprobacion de tipo
	code += "if (" + tipo + ">0) goto " + tipo1 + ";\n"

	//ES UN NUMERO
	code += "printf(\"%d\",(int)" + valor + ");\n" //]
	code += "goto " + tipo2 + ";\n"

	//ES UN VECTOR
	code += tipo1 + ":\n"
	//PARAMETROS PARA LLAMAR DE NUEVO AL PROC PRINTINT
	code += "STACK[(int)" + index_p1 + "]=" + tipo + ";\n"
	code += "STACK[(int)" + index_p2 + "]=" + valor + ";\n"
	//DEJANDO UN RETORNO DONDE SE GUARDA LA POSICION ACTUAL
	code += size_return + "=STACK[81999];\n"
	code += size_return + "=1+" + size_return + ";\n"
	code += index_ret + "=81999-" + size_return + ";\n"
	code += "STACK[(int)" + index_ret + "]=" + pos + ";\n"
	code += size_return + "=1+" + size_return + ";\n"
	code += index_ret + "=81999-" + size_return + ";\n"
	code += "STACK[(int)" + index_ret + "]=" + iterador + ";\n"
	code += size_return + "=1+" + size_return + ";\n"
	code += index_ret + "=81999-" + size_return + ";\n"
	code += "STACK[(int)" + index_ret + "]=" + tipo + ";\n"
	code += size_return + "=1+" + size_return + ";\n"
	code += index_ret + "=81999-" + size_return + ";\n"
	code += "STACK[(int)" + index_ret + "]=" + longitud + ";\n"
	code += "STACK[81999]=" + size_return + ";\n"
	//LLAMADO DE LA FUNCION
	code += "proc_printInt();\n"
	//RECUPERACION DE LA LONGITUD
	code += size_return + "=STACK[81999];\n"
	code += index_ret + "=81999-" + size_return + ";\n"    //CALCULO DEL INDICE DONDE ESTA EL RETORNO
	code += size_return + "=" + size_return + "-1;\n"      //DISMINUCION DEL NUMERO DE RETORNOS ESPERADOS
	code += longitud + "=STACK[(int)" + index_ret + "];\n" //RECUPERACION FINAL DE LA POSICION
	//code += size_return + "=" + size_return + "-1;\n"
	//RECUPERACION DE LA TIPO
	code += index_ret + "=81999-" + size_return + ";\n"
	code += size_return + "=" + size_return + "-1;\n"
	code += tipo + "=STACK[(int)" + index_ret + "];\n"
	//RECUPERACION DEL ITERADOD
	code += index_ret + "=81999-" + size_return + ";\n"
	code += size_return + "=" + size_return + "-1;\n"
	code += iterador + "=STACK[(int)" + index_ret + "];\n"
	//RECUPERACION DE LA POSICION
	code += index_ret + "=81999-" + size_return + ";\n"
	code += size_return + "=" + size_return + "-1;\n"
	code += pos + "=STACK[(int)" + index_ret + "];\n"
	//MODIFICACION DEL TAMAÑO
	code += "STACK[81999]=" + size_return + ";\n"
	//SON DATOS
	code += tipo2 + ":\n"
	code += "printf(\"%c\",44);\n" //","
	code += pos + "=" + pos + "+1;\n"
	code += iterador + "=" + iterador + "+1;\n"
	code += "goto " + entrada + ";\n"
	code += salida + ":\n"
	code += "printf(\"%c\",93);\n" //]
	code += "return;\n}\n "
	return code
}

func bounds_error() string {
	code := "void proc_boundsError(){\n"
	code += "printf(\"%c\",66);\n"  //b
	code += "printf(\"%c\",111);\n" //o
	code += "printf(\"%c\",117);\n" //u
	code += "printf(\"%c\",110);\n" //n
	code += "printf(\"%c\",100);\n" //d
	code += "printf(\"%c\",115);\n" //s
	code += "printf(\"%c\",69);\n"  //e
	code += "printf(\"%c\",114);\n" //r
	code += "printf(\"%c\",114);\n" //r
	code += "printf(\"%c\",111);\n" //o
	code += "printf(\"%c\",114);\n" //r
	code += "printf(\"%c\",110);\n" //r
	code += "return;\n}"
	return code
}

func (g *Generator) Array_char(str string) []rune {
	runes := []rune(str)
	return runes
}
