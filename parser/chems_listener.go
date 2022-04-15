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

	// EnterBreaks is called when entering the breaks production.
	EnterBreaks(c *BreaksContext)

	// EnterContinues is called when entering the continues production.
	EnterContinues(c *ContinuesContext)

	// EnterExpression is called when entering the expression production.
	EnterExpression(c *ExpressionContext)

	// EnterExpr_arit is called when entering the expr_arit production.
	EnterExpr_arit(c *Expr_aritContext)

	// EnterPrimitivo is called when entering the primitivo production.
	EnterPrimitivo(c *PrimitivoContext)

	// ExitStart is called when exiting the start production.
	ExitStart(c *StartContext)

	// ExitInstrucciones is called when exiting the instrucciones production.
	ExitInstrucciones(c *InstruccionesContext)

	// ExitInstruccion is called when exiting the instruccion production.
	ExitInstruccion(c *InstruccionContext)

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

	// ExitBreaks is called when exiting the breaks production.
	ExitBreaks(c *BreaksContext)

	// ExitContinues is called when exiting the continues production.
	ExitContinues(c *ContinuesContext)

	// ExitExpression is called when exiting the expression production.
	ExitExpression(c *ExpressionContext)

	// ExitExpr_arit is called when exiting the expr_arit production.
	ExitExpr_arit(c *Expr_aritContext)

	// ExitPrimitivo is called when exiting the primitivo production.
	ExitPrimitivo(c *PrimitivoContext)
}
