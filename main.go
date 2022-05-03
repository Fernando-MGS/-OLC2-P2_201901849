package main

import (
	"OLC2/environment"
	"OLC2/generator"
	"OLC2/interfaces"
	"OLC2/parser"
	"fmt"
	"log"
	"os"
	"reflect"

	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/container"
	"fyne.io/fyne/layout"
	"fyne.io/fyne/theme"
	"fyne.io/fyne/widget"
	"github.com/antlr/antlr4/runtime/Go/antlr"
)

var consola string
var salida string
var RepErr []interfaces.Errores
var RepSim []interfaces.Simbolos
var RepBase []interfaces.Bases
var RepOptimizacion []interfaces.Optimizacion
var Data widget.Label
var list *widget.Table
var list1 *widget.Table
var list2 *widget.Table
var list3 *widget.Table

type TreeShapeListener struct {
	*parser.BaseChemsListener
}

func NewTreeShapeListener() *TreeShapeListener {
	return new(TreeShapeListener)
}

func (this *TreeShapeListener) ExitStart(ctx *parser.StartContext) {
	result := ctx.GetLista()

	var globalEnv environment.Environment
	globalEnv = environment.NewEnvironment(nil, "GLOBAL", "", "", false, 1)

	var gen *generator.Generator
	gen = generator.NewGenerator()
	//var salida string
	salida = ""

	for _, s := range result.ToArray() {
		tipo := reflect.TypeOf(s)
		t := fmt.Sprintf("%v", tipo)
		if t == "instruction.Functions" {
			s.(interfaces.Instruction).Ejecutar(globalEnv, gen)
		} else if t == "instruction.Struct" {
			s.(interfaces.Instruction).Ejecutar(globalEnv, gen)
		}
		//s.(interfaces.Instruction).Ejecutar(globalEnv, gen)
	}

	//Escribir salida

	f, err := os.Create("salida.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	salida += "#include <stdio.h>\n"
	salida += "#include <math.h>\n"
	salida += "float HEAP[82000];\n"
	salida += "float STACK[82000];\n"
	salida += "float P;\n"
	salida += "float H;\n"
	if gen.GetTemps().Len() > 0 {
		salida += "float "

		salida += fmt.Sprintf("%v", gen.GetTemps().GetValue(0))
		gen.GetTemps().RemoveAtIndex(0)

		for _, s := range gen.GetTemps().ToArray() {
			salida += ", "
			salida += fmt.Sprintf("%v", s)
		}

		salida += ";\n\n"
	}
	for _, s := range gen.GetExtraFuncs().ToArray() {
		salida += fmt.Sprintf("%v", s)
		salida += "\n"
	}
	salida += "\n"
	/*for _, s := range gen.GetFuncs().ToArray() {
		salida += fmt.Sprintf("%v", s)
		salida += "\n"
	}*/
	salida += gen.GetFunciones()

	/*salida += "\nvoid main(){\n"

	for _, s := range gen.GetCode().ToArray() {
		salida += fmt.Sprintf("%v", s)
		salida += "\n"
	}

	salida += "return;\n}\n"*/

	errores := globalEnv.DevErrores()
	simbolos := globalEnv.DevSimbolos()
	bases := globalEnv.DevBases()
	var reset []interfaces.Errores
	var resetS []interfaces.Simbolos
	var resetB []interfaces.Bases
	var resetO []interfaces.Optimizacion
	RepErr = reset
	RepSim = resetS
	RepBase = resetB
	RepOptimizacion = resetO
	RepErr = append(RepErr, interfaces.Errores{Line: "LINEA", Col: "COLUMNA", Mess: "MENSAJE", Fecha: "FECHA"})
	RepSim = append(RepSim, interfaces.Simbolos{ID: "ID", Tipo: "TIPO", Ambito: "AMBITO", Fila: "FILA"})
	RepBase = append(RepBase, interfaces.Bases{Nombre: "NOMBRE", NoTables: "TABLAS", Linea: "LINEA"})
	RepOptimizacion = append(RepOptimizacion, interfaces.Optimizacion{Tipo: "TIPO", Regla: "REGLA", Expr_Original: "EXPRESION ORIGINAL", Expr_Optima: "EXPRESION OPTIMA", Fila: "FILA"})
	for _, s := range bases {
		RepBase = append(RepBase, s.(interfaces.Bases))
	}
	for _, s := range simbolos {
		RepSim = append(RepSim, s.(interfaces.Simbolos))
	}
	for _, s := range errores {
		RepErr = append(RepErr, s.(interfaces.Errores))
	}
	_, err2 := f.WriteString(salida)

	if err2 != nil {
		log.Fatal(err2)
	}

}

func ejecutar_antlr() {
	is := antlr.NewInputStream(consola)
	// Create the Lexer
	lexer := parser.NewChemsLexer(is)
	stream := antlr.NewCommonTokenStream(lexer, antlr.TokenDefaultChannel)

	// Create the Parser
	p := parser.NewChems(stream)

	p.BuildParseTrees = true
	tree := p.Start()

	antlr.ParseTreeWalkerDefault.Walk(NewTreeShapeListener(), tree)
}

func show_Data() {

	if Data.Text == "" {
		Data.Text = "Fernando Mauricio Gómez Santos -201901849-OLC2-N"
	} else {
		Data.Text = ""
	}
}

func set_Reportes() {
	err1 := interfaces.Errores{Line: "LINEA", Col: "COLUMNA", Mess: "DETALLES", Fecha: "FECHA"}
	sim1 := interfaces.Simbolos{ID: "ID", Tipo: "TIPO", Ambito: "AMBITO", Fila: "FILA", Columna: "COLUMNA"}
	bases := interfaces.Bases{Nombre: "MODULOS", NoTables: "NO. TABLAS", Linea: "LINEA"}
	optimos := interfaces.Optimizacion{Tipo: "TIPO", Regla: "REGLA", Expr_Original: "EXPRESION ORIGINAL", Expr_Optima: "EXPRESION OPTIMA", Fila: "FILA"}
	RepErr = []interfaces.Errores{err1}
	RepSim = []interfaces.Simbolos{sim1}
	RepBase = append(RepBase, bases)
	RepOptimizacion = append(RepOptimizacion, optimos)
}

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("DB-RUST")
	myWindow.Resize(fyne.NewSize(1300, 650))
	Data.Text = ""
	editor := widget.NewMultiLineEntry()
	editor.SetPlaceHolder("//OLC2-DB RUST")
	titulos := widget.NewLabel("EDITOR                                                                                                                                                   CONSOLA")
	result := widget.NewMultiLineEntry()
	set_Reportes()
	result.Disable()
	result.SetPlaceHolder("CONSOLA")
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.MediaPlayIcon(), func() {
			consola = editor.Text
			ejecutar_antlr()
			//result.SetText(salida)
			list.Refresh()
			list1.Refresh()
			list2.Refresh()
			list3.Refresh()
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.ConfirmIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), show_Data),
	)
	list = widget.NewTable(
		func() (int, int) {
			return len(RepErr), 4
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("----------------------------------------------")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == 0 {
				o.(*widget.Label).SetText(RepErr[i.Row].Line)
				o.Refresh()
			} else if i.Col == 1 {
				o.(*widget.Label).SetText(RepErr[i.Row].Col)
				o.Refresh()
			} else if i.Col == 2 {
				o.(*widget.Label).SetText(RepErr[i.Row].Fecha)
				o.Refresh()
			} else if i.Col == 3 {
				o.(*widget.Label).SetText(RepErr[i.Row].Mess)
				o.Refresh()
			}

		})
	//list.Refresh()
	list1 = widget.NewTable(
		func() (int, int) {
			return len(RepSim), 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("-------------------------------------------------------")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == 0 {
				o.(*widget.Label).SetText(RepSim[i.Row].Fila)
			} else if i.Col == 1 {
				o.(*widget.Label).SetText(RepSim[i.Row].Columna)
			} else if i.Col == 2 {
				o.(*widget.Label).SetText(RepSim[i.Row].Ambito)
			} else if i.Col == 3 {
				o.(*widget.Label).SetText(RepSim[i.Row].Tipo)
			} else if i.Col == 4 {
				o.(*widget.Label).SetText(RepSim[i.Row].ID)
			}
		})
	list2 = widget.NewTable(
		func() (int, int) {
			return len(RepBase), 3
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("-----------------------------------------------------")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == 0 {
				o.(*widget.Label).SetText(RepBase[i.Row].Nombre)
			} else if i.Col == 1 {
				o.(*widget.Label).SetText(RepBase[i.Row].NoTables)
			} else if i.Col == 2 {
				o.(*widget.Label).SetText(RepBase[i.Row].Linea)
			} else if i.Col == 3 {
				o.(*widget.Label).SetText("-----")
			} else if i.Col == 4 {
				o.(*widget.Label).SetText("-----")
			} else if i.Col == 5 {
				o.(*widget.Label).SetText("-----")
			}

		})
	list3 = widget.NewTable(
		func() (int, int) {
			return len(RepBase), 5
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("------------------------------------------------------")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			if i.Col == 0 {
				o.(*widget.Label).SetText(RepOptimizacion[i.Row].Fila)
			} else if i.Col == 1 {
				o.(*widget.Label).SetText(RepOptimizacion[i.Row].Regla)
			} else if i.Col == 2 {
				o.(*widget.Label).SetText(RepOptimizacion[i.Row].Tipo)
			} else if i.Col == 3 {
				o.(*widget.Label).SetText(RepOptimizacion[i.Row].Expr_Original)
			} else if i.Col == 4 {
				o.(*widget.Label).SetText(RepOptimizacion[i.Row].Expr_Optima)
			} else if i.Col == 5 {
				o.(*widget.Label).SetText("-----")
			}

		})
	console := container.NewVScroll(editor)
	console15 := container.NewHScroll(console)
	console1 := container.NewGridWrap(fyne.NewSize(600, 300), console15)
	console3 := container.NewVScroll(result)
	console35 := container.NewHScroll(console3)
	console4 := container.NewGridWrap(fyne.NewSize(600, 300), console35)

	report1 := container.NewHScroll(list1)
	report2 := container.NewGridWrap(fyne.NewSize(1200, 270), report1)
	report3 := container.NewHScroll(list)
	report4 := container.NewGridWrap(fyne.NewSize(1200, 270), report3)
	report5 := container.NewHScroll(list2)
	report6 := container.NewGridWrap(fyne.NewSize(1200, 270), report5)
	report7 := container.NewHScroll(list3)
	report8 := container.NewGridWrap(fyne.NewSize(1200, 270), report7)
	//editor.
	/*console.Resize(fyne.NewSize(500, 600))
	editor.Resize(fyne.NewSize(500, 250))*/
	//content := container.NewGridWrap(fyne.NewSize(300, 400), editor)
	tabs := container.NewAppTabs(
		container.NewTabItem("SIMBOLOS", report2),
		container.NewTabItem("ERRORES", report4),
		container.NewTabItem("BASES", report6),
		container.NewTabItem("OPTIMIZACION", report8),
	)
	myWindow.SetContent(
		fyne.NewContainerWithLayout(
			layout.NewVBoxLayout(),
			fyne.NewContainerWithLayout(layout.NewHBoxLayout(), toolbar), &Data, titulos,
			fyne.NewContainerWithLayout(
				layout.NewVBoxLayout(),
				fyne.NewContainerWithLayout(layout.NewHBoxLayout(), console1, console4),
			), tabs,
		),
	)
	myWindow.CenterOnScreen()

	//myWindow.SetContent(content)
	myWindow.ShowAndRun()
}
