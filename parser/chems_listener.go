// Code generated from Chems.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Chems

import "github.com/antlr/antlr4/runtime/Go/antlr"

// ChemsListener is a complete listener for a parse tree produced by Chems.
type ChemsListener interface {
	antlr.ParseTreeListener

	// EnterStart is called when entering the start production.
	EnterStart(c *StartContext)

	// EnterInstrucciones is called when entering the instrucciones production.
	EnterInstrucciones(c *InstruccionesContext)

	// EnterInstruccion is called when entering the instruccion production.
	EnterInstruccion(c *InstruccionContext)

	// EnterInstruccion_wc is called when entering the instruccion_wc production.
	EnterInstruccion_wc(c *Instruccion_wcContext)

	// EnterAsignacion_var is called when entering the asignacion_var production.
	EnterAsignacion_var(c *Asignacion_varContext)

	// EnterDeclaracion_var is called when entering the declaracion_var production.
	EnterDeclaracion_var(c *Declaracion_varContext)

	// EnterMutable is called when entering the mutable production.
	EnterMutable(c *MutableContext)

	// EnterTypes is called when entering the types production.
	EnterTypes(c *TypesContext)

	// EnterTipo_d is called when entering the tipo_d production.
	EnterTipo_d(c *Tipo_dContext)

	// EnterAsignar_Array is called when entering the asignar_Array production.
	EnterAsignar_Array(c *Asignar_ArrayContext)

	// EnterDimensiones is called when entering the dimensiones production.
	EnterDimensiones(c *DimensionesContext)

	// EnterTipo_vector is called when entering the tipo_vector production.
	EnterTipo_vector(c *Tipo_vectorContext)

	// EnterVectores is called when entering the vectores production.
	EnterVectores(c *VectoresContext)

	// EnterLoops is called when entering the loops production.
	EnterLoops(c *LoopsContext)

	// EnterIfs is called when entering the ifs production.
	EnterIfs(c *IfsContext)

	// EnterLlaves_if is called when entering the llaves_if production.
	EnterLlaves_if(c *Llaves_ifContext)

	// EnterElses is called when entering the elses production.
	EnterElses(c *ElsesContext)

	// EnterBreaks is called when entering the breaks production.
	EnterBreaks(c *BreaksContext)

	// EnterContinues is called when entering the continues production.
	EnterContinues(c *ContinuesContext)

	// EnterImpr is called when entering the impr production.
	EnterImpr(c *ImprContext)

	// EnterFormato is called when entering the formato production.
	EnterFormato(c *FormatoContext)

	// EnterMatches is called when entering the matches production.
	EnterMatches(c *MatchesContext)

	// EnterCasos is called when entering the casos production.
	EnterCasos(c *CasosContext)

	// EnterCases is called when entering the cases production.
	EnterCases(c *CasesContext)

	// EnterConditions is called when entering the conditions production.
	EnterConditions(c *ConditionsContext)

	// EnterDefaults is called when entering the defaults production.
	EnterDefaults(c *DefaultsContext)

	// EnterSet_match is called when entering the set_match production.
	EnterSet_match(c *Set_matchContext)

	// EnterRfor is called when entering the rfor production.
	EnterRfor(c *RforContext)

	// EnterIter_for is called when entering the iter_for production.
	EnterIter_for(c *Iter_forContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpr_arit is called when entering the expr_arit production.
	EnterExpr_arit(c *Expr_aritContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// EnterListValues is called when entering the listValues production.
	EnterListValues(c *ListValuesContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

	// ExitInstruccion_wc is called when exiting the instruccion_wc production.
	ExitInstruccion_wc(c *Instruccion_wcContext)

	// ExitAsignacion_var is called when exiting the asignacion_var production.
	ExitAsignacion_var(c *Asignacion_varContext)

	// ExitDeclaracion_var is called when exiting the declaracion_var production.
	ExitDeclaracion_var(c *Declaracion_varContext)

	// ExitMutable is called when exiting the mutable production.
	ExitMutable(c *MutableContext)

	// ExitTypes is called when exiting the types production.
	ExitTypes(c *TypesContext)

	// ExitTipo_d is called when exiting the tipo_d production.
	ExitTipo_d(c *Tipo_dContext)

	// ExitAsignar_Array is called when exiting the asignar_Array production.
	ExitAsignar_Array(c *Asignar_ArrayContext)

	// ExitDimensiones is called when exiting the dimensiones production.
	ExitDimensiones(c *DimensionesContext)

	// ExitTipo_vector is called when exiting the tipo_vector production.
	ExitTipo_vector(c *Tipo_vectorContext)

	// ExitVectores is called when exiting the vectores production.
	ExitVectores(c *VectoresContext)

	// ExitLoops is called when exiting the loops production.
	ExitLoops(c *LoopsContext)

	// ExitIfs is called when exiting the ifs production.
	ExitIfs(c *IfsContext)

	// ExitLlaves_if is called when exiting the llaves_if production.
	ExitLlaves_if(c *Llaves_ifContext)

	// ExitElses is called when exiting the elses production.
	ExitElses(c *ElsesContext)

	// ExitBreaks is called when exiting the breaks production.
	ExitBreaks(c *BreaksContext)

	// ExitContinues is called when exiting the continues production.
	ExitContinues(c *ContinuesContext)

	// ExitImpr is called when exiting the impr production.
	ExitImpr(c *ImprContext)

	// ExitFormato is called when exiting the formato production.
	ExitFormato(c *FormatoContext)

	// ExitMatches is called when exiting the matches production.
	ExitMatches(c *MatchesContext)

	// ExitCasos is called when exiting the casos production.
	ExitCasos(c *CasosContext)

	// ExitCases is called when exiting the cases production.
	ExitCases(c *CasesContext)

	// ExitConditions is called when exiting the conditions production.
	ExitConditions(c *ConditionsContext)

	// ExitDefaults is called when exiting the defaults production.
	ExitDefaults(c *DefaultsContext)

	// ExitSet_match is called when exiting the set_match production.
	ExitSet_match(c *Set_matchContext)

	// ExitRfor is called when exiting the rfor production.
	ExitRfor(c *RforContext)

	// ExitIter_for is called when exiting the iter_for production.
	ExitIter_for(c *Iter_forContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpr_arit is called when exiting the expr_arit production.
	ExitExpr_arit(c *Expr_aritContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)

	// ExitListValues is called when exiting the listValues production.
	ExitListValues(c *ListValuesContext)
}
