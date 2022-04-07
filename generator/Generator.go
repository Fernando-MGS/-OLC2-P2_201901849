package generator

import (
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type Generator struct {
	temporal int
	label    int
	Stack    int
	Heap     int
	code     *arrayList.List
	tempList *arrayList.List
}

func NewGenerator() *Generator {
	//fmt.Println("neuvo generator")
	generator := Generator{temporal: 0, label: 0, code: arrayList.New(), tempList: arrayList.New()}
	return &generator
}

func (g Generator) GetCode() *arrayList.List {

	return g.code
}

func (g Generator) GetTemps() *arrayList.List {

	return g.tempList
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

func (g *Generator) AddExpression(target string, left string, right string, operator string) {
	g.code.Add(target + " = " + left + " " + operator + " " + right + ";")
}

//Añade un printf
func (g *Generator) AddPrintf(typePrint string, value string) {
	g.code.Add("printf(\"%" + typePrint + "\"," + value + ");")
}

func (g *Generator) AddCode(code string) {
	g.code.Add(code)
}

func (g *Generator) AddFunc() {

}
