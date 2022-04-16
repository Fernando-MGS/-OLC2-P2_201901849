// Code generated from Chems.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// BaseChemsListener is a complete listener for a parse tree produced by Chems.
type BaseChemsListener struct{}

var _ ChemsListener = &BaseChemsListener{}

// VisitTerminal is called when a terminal node is visited.
func (s *BaseChemsListener) VisitTerminal(node antlr.TerminalNode) {}

// VisitErrorNode is called when an error node is visited.
func (s *BaseChemsListener) VisitErrorNode(node antlr.ErrorNode) {}

// EnterEveryRule is called when any rule is entered.
func (s *BaseChemsListener) EnterEveryRule(ctx antlr.ParserRuleContext) {}

// ExitEveryRule is called when any rule is exited.
func (s *BaseChemsListener) ExitEveryRule(ctx antlr.ParserRuleContext) {}

// EnterStart is called when production start is entered.
func (s *BaseChemsListener) EnterStart(ctx *StartContext) {}

// ExitStart is called when production start is exited.
func (s *BaseChemsListener) ExitStart(ctx *StartContext) {}

// EnterInstrucciones is called when production instrucciones is entered.
func (s *BaseChemsListener) EnterInstrucciones(ctx *InstruccionesContext) {}

// ExitInstrucciones is called when production instrucciones is exited.
func (s *BaseChemsListener) ExitInstrucciones(ctx *InstruccionesContext) {}

// EnterInstruccion is called when production instruccion is entered.
func (s *BaseChemsListener) EnterInstruccion(ctx *InstruccionContext) {}

// ExitInstruccion is called when production instruccion is exited.
func (s *BaseChemsListener) ExitInstruccion(ctx *InstruccionContext) {}

// EnterAsignacion_var is called when production asignacion_var is entered.
func (s *BaseChemsListener) EnterAsignacion_var(ctx *Asignacion_varContext) {}

// ExitAsignacion_var is called when production asignacion_var is exited.
func (s *BaseChemsListener) ExitAsignacion_var(ctx *Asignacion_varContext) {}

// EnterDeclaracion_var is called when production declaracion_var is entered.
func (s *BaseChemsListener) EnterDeclaracion_var(ctx *Declaracion_varContext) {}

// ExitDeclaracion_var is called when production declaracion_var is exited.
func (s *BaseChemsListener) ExitDeclaracion_var(ctx *Declaracion_varContext) {}

// EnterMutable is called when production mutable is entered.
func (s *BaseChemsListener) EnterMutable(ctx *MutableContext) {}

// ExitMutable is called when production mutable is exited.
func (s *BaseChemsListener) ExitMutable(ctx *MutableContext) {}

// EnterTypes is called when production types is entered.
func (s *BaseChemsListener) EnterTypes(ctx *TypesContext) {}

// ExitTypes is called when production types is exited.
func (s *BaseChemsListener) ExitTypes(ctx *TypesContext) {}

// EnterTipo_d is called when production tipo_d is entered.
func (s *BaseChemsListener) EnterTipo_d(ctx *Tipo_dContext) {}

// ExitTipo_d is called when production tipo_d is exited.
func (s *BaseChemsListener) ExitTipo_d(ctx *Tipo_dContext) {}

// EnterAsignar_Array is called when production asignar_Array is entered.
func (s *BaseChemsListener) EnterAsignar_Array(ctx *Asignar_ArrayContext) {}

// ExitAsignar_Array is called when production asignar_Array is exited.
func (s *BaseChemsListener) ExitAsignar_Array(ctx *Asignar_ArrayContext) {}

// EnterDimensiones is called when production dimensiones is entered.
func (s *BaseChemsListener) EnterDimensiones(ctx *DimensionesContext) {}

// ExitDimensiones is called when production dimensiones is exited.
func (s *BaseChemsListener) ExitDimensiones(ctx *DimensionesContext) {}

// EnterTipo_vector is called when production tipo_vector is entered.
func (s *BaseChemsListener) EnterTipo_vector(ctx *Tipo_vectorContext) {}

// ExitTipo_vector is called when production tipo_vector is exited.
func (s *BaseChemsListener) ExitTipo_vector(ctx *Tipo_vectorContext) {}

// EnterVectores is called when production vectores is entered.
func (s *BaseChemsListener) EnterVectores(ctx *VectoresContext) {}

// ExitVectores is called when production vectores is exited.
func (s *BaseChemsListener) ExitVectores(ctx *VectoresContext) {}

// EnterLoops is called when production loops is entered.
func (s *BaseChemsListener) EnterLoops(ctx *LoopsContext) {}

// ExitLoops is called when production loops is exited.
func (s *BaseChemsListener) ExitLoops(ctx *LoopsContext) {}

// EnterIfs is called when production ifs is entered.
func (s *BaseChemsListener) EnterIfs(ctx *IfsContext) {}

// ExitIfs is called when production ifs is exited.
func (s *BaseChemsListener) ExitIfs(ctx *IfsContext) {}

// EnterLlaves_if is called when production llaves_if is entered.
func (s *BaseChemsListener) EnterLlaves_if(ctx *Llaves_ifContext) {}

// ExitLlaves_if is called when production llaves_if is exited.
func (s *BaseChemsListener) ExitLlaves_if(ctx *Llaves_ifContext) {}

// EnterElses is called when production elses is entered.
func (s *BaseChemsListener) EnterElses(ctx *ElsesContext) {}

// ExitElses is called when production elses is exited.
func (s *BaseChemsListener) ExitElses(ctx *ElsesContext) {}

// EnterBreaks is called when production breaks is entered.
func (s *BaseChemsListener) EnterBreaks(ctx *BreaksContext) {}

// ExitBreaks is called when production breaks is exited.
func (s *BaseChemsListener) ExitBreaks(ctx *BreaksContext) {}

// EnterContinues is called when production continues is entered.
func (s *BaseChemsListener) EnterContinues(ctx *ContinuesContext) {}

// ExitContinues is called when production continues is exited.
func (s *BaseChemsListener) ExitContinues(ctx *ContinuesContext) {}

// EnterImpr is called when production impr is entered.
func (s *BaseChemsListener) EnterImpr(ctx *ImprContext) {}

// ExitImpr is called when production impr is exited.
func (s *BaseChemsListener) ExitImpr(ctx *ImprContext) {}

// EnterFormato is called when production formato is entered.
func (s *BaseChemsListener) EnterFormato(ctx *FormatoContext) {}

// ExitFormato is called when production formato is exited.
func (s *BaseChemsListener) ExitFormato(ctx *FormatoContext) {}

// EnterExpression is called when production expression is entered.
func (s *BaseChemsListener) EnterExpression(ctx *ExpressionContext) {}

// ExitExpression is called when production expression is exited.
func (s *BaseChemsListener) ExitExpression(ctx *ExpressionContext) {}

// EnterExpr_arit is called when production expr_arit is entered.
func (s *BaseChemsListener) EnterExpr_arit(ctx *Expr_aritContext) {}

// ExitExpr_arit is called when production expr_arit is exited.
func (s *BaseChemsListener) ExitExpr_arit(ctx *Expr_aritContext) {}

// EnterPrimitivo is called when production primitivo is entered.
func (s *BaseChemsListener) EnterPrimitivo(ctx *PrimitivoContext) {}

// ExitPrimitivo is called when production primitivo is exited.
func (s *BaseChemsListener) ExitPrimitivo(ctx *PrimitivoContext) {}

// EnterListValues is called when production listValues is entered.
func (s *BaseChemsListener) EnterListValues(ctx *ListValuesContext) {}

// ExitListValues is called when production listValues is exited.
func (s *BaseChemsListener) ExitListValues(ctx *ListValuesContext) {}
