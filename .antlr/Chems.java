// Generated from c:\Users\ferna\OneDrive\Documentos\Go\src\github.com\Fernando-MGS\-OLC2-P2_201901849\Chems.g4 by ANTLR 4.8

    import "OLC2/interfaces"
    import "OLC2/expresion"
    import "OLC2/instruction"
    import arrayList "github.com/colegno/arraylist"

import org.antlr.v4.runtime.atn.*;
import org.antlr.v4.runtime.dfa.DFA;
import org.antlr.v4.runtime.*;
import org.antlr.v4.runtime.misc.*;
import org.antlr.v4.runtime.tree.*;
import java.util.List;
import java.util.Iterator;
import java.util.ArrayList;

@SuppressWarnings({"all", "warnings", "unchecked", "unused", "cast"})
public class Chems extends Parser {
	static { RuntimeMetaData.checkVersion("4.8", RuntimeMetaData.VERSION); }

	protected static final DFA[] _decisionToDFA;
	protected static final PredictionContextCache _sharedContextCache =
		new PredictionContextCache();
	public static final int
		PUNTO=1, PTCOMA=2, COMA=3, COMILLAS=4, DIFERENTE=5, DOSPUNTOS=6, IGUAL=7, 
		AMPER=8, MENOR=9, MAYORIGUAL=10, MENORIGUAL=11, OPMATCH=12, D_IGUAL=13, 
		NOT_E=14, MAYOR=15, OR=16, AND=17, MUL=18, DIV=19, ADD=20, SUB=21, MOD=22, 
		POT=23, PARIZQ=24, PARDER=25, LLAVEIZQ=26, LLAVEDER=27, CORIZQ=28, CORDER=29, 
		INTDER=30, OR_COND=31, DDPUNTO=32, CONSOLE=33, LOG=34, PRINTLN=35, P_NUMBER=36, 
		P_STRING=37, P_IF=38, P_ELSE=39, P_WHILE=40, P_FOR=41, P_IN=42, P_AS=43, 
		POWF=44, POW=45, BREAK=46, CONTINUE=47, RETURN=48, LET=49, INT=50, FLOAT=51, 
		BOOL=52, CHAR=53, STR=54, USIZE=55, MUT=56, P_STRUCT=57, FUNCTION=58, 
		MODULE=59, PUB=60, TRUE=61, FALSE=62, MATCH=63, LOOP=64, PUSH=65, T_STRING=66, 
		F_ABS=67, F_SQRT=68, F_CLONE=69, INSERT=70, REMOVE=71, CONTAINS=72, LEN=73, 
		CAPACITY=74, DEF=75, VEC=76, VECT=77, CAP=78, NEW=79, COMENTARIO=80, NUMBER=81, 
		DECIMAL=82, STRING=83, ID=84, CARACTER=85, WHITESPACE=86;
	public static final int
		RULE_start = 0, RULE_instrucciones = 1, RULE_instruccion = 2, RULE_instruccion_wc = 3, 
		RULE_asignacion_var = 4, RULE_declaracion_var = 5, RULE_mutable = 6, RULE_types = 7, 
		RULE_tipo_d = 8, RULE_asignar_Array = 9, RULE_dimensiones = 10, RULE_tipo_vector = 11, 
		RULE_vectores = 12, RULE_loops = 13, RULE_ifs = 14, RULE_llaves_if = 15, 
		RULE_elses = 16, RULE_breaks = 17, RULE_continues = 18, RULE_impr = 19, 
		RULE_formato = 20, RULE_matches = 21, RULE_casos = 22, RULE_cases = 23, 
		RULE_conditions = 24, RULE_defaults = 25, RULE_set_match = 26, RULE_rfor = 27, 
		RULE_iter_for = 28, RULE_expression = 29, RULE_expr_arit = 30, RULE_primitivo = 31, 
		RULE_listValues = 32, RULE_arrayAcc = 33, RULE_listArray = 34;
	private static String[] makeRuleNames() {
		return new String[] {
			"start", "instrucciones", "instruccion", "instruccion_wc", "asignacion_var", 
			"declaracion_var", "mutable", "types", "tipo_d", "asignar_Array", "dimensiones", 
			"tipo_vector", "vectores", "loops", "ifs", "llaves_if", "elses", "breaks", 
			"continues", "impr", "formato", "matches", "casos", "cases", "conditions", 
			"defaults", "set_match", "rfor", "iter_for", "expression", "expr_arit", 
			"primitivo", "listValues", "arrayAcc", "listArray"
		};
	}
	public static final String[] ruleNames = makeRuleNames();

	private static String[] makeLiteralNames() {
		return new String[] {
			null, "'.'", "';'", "','", "'\"'", "'!'", "':'", "'='", "'&'", "'<'", 
			"'>='", "'<='", "'=>'", "'=='", "'!='", "'>'", "'||'", "'&&'", "'*'", 
			"'/'", "'+'", "'-'", "'%'", "'^'", "'('", "')'", "'{'", "'}'", "'['", 
			"']'", "'?'", "'|'", "'::'", "'console'", "'log'", "'println!'", "'number'", 
			"'String'", "'if'", "'else'", "'while'", "'for'", "'in'", "'as'", "'powf'", 
			"'pow'", "'break'", "'continue'", "'return'", "'let'", "'i64'", "'f64'", 
			"'bool'", "'char'", "'&str'", "'usize'", "'mut'", "'struct'", "'fn'", 
			"'mod'", "'pub'", "'true'", "'false'", "'match'", "'loop'", "'push'", 
			"'to_string'", "'abs'", "'sqrt'", "'clone'", "'insert'", "'remove'", 
			"'contains'", "'len'", "'capacity'", "'_'", "'vec!'", "'Vec'", "'with_capacity'", 
			"'new'"
		};
	}
	private static final String[] _LITERAL_NAMES = makeLiteralNames();
	private static String[] makeSymbolicNames() {
		return new String[] {
			null, "PUNTO", "PTCOMA", "COMA", "COMILLAS", "DIFERENTE", "DOSPUNTOS", 
			"IGUAL", "AMPER", "MENOR", "MAYORIGUAL", "MENORIGUAL", "OPMATCH", "D_IGUAL", 
			"NOT_E", "MAYOR", "OR", "AND", "MUL", "DIV", "ADD", "SUB", "MOD", "POT", 
			"PARIZQ", "PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ", "CORDER", "INTDER", 
			"OR_COND", "DDPUNTO", "CONSOLE", "LOG", "PRINTLN", "P_NUMBER", "P_STRING", 
			"P_IF", "P_ELSE", "P_WHILE", "P_FOR", "P_IN", "P_AS", "POWF", "POW", 
			"BREAK", "CONTINUE", "RETURN", "LET", "INT", "FLOAT", "BOOL", "CHAR", 
			"STR", "USIZE", "MUT", "P_STRUCT", "FUNCTION", "MODULE", "PUB", "TRUE", 
			"FALSE", "MATCH", "LOOP", "PUSH", "T_STRING", "F_ABS", "F_SQRT", "F_CLONE", 
			"INSERT", "REMOVE", "CONTAINS", "LEN", "CAPACITY", "DEF", "VEC", "VECT", 
			"CAP", "NEW", "COMENTARIO", "NUMBER", "DECIMAL", "STRING", "ID", "CARACTER", 
			"WHITESPACE"
		};
	}
	private static final String[] _SYMBOLIC_NAMES = makeSymbolicNames();
	public static final Vocabulary VOCABULARY = new VocabularyImpl(_LITERAL_NAMES, _SYMBOLIC_NAMES);

	/**
	 * @deprecated Use {@link #VOCABULARY} instead.
	 */
	@Deprecated
	public static final String[] tokenNames;
	static {
		tokenNames = new String[_SYMBOLIC_NAMES.length];
		for (int i = 0; i < tokenNames.length; i++) {
			tokenNames[i] = VOCABULARY.getLiteralName(i);
			if (tokenNames[i] == null) {
				tokenNames[i] = VOCABULARY.getSymbolicName(i);
			}

			if (tokenNames[i] == null) {
				tokenNames[i] = "<INVALID>";
			}
		}
	}

	@Override
	@Deprecated
	public String[] getTokenNames() {
		return tokenNames;
	}

	@Override

	public Vocabulary getVocabulary() {
		return VOCABULARY;
	}

	@Override
	public String getGrammarFileName() { return "Chems.g4"; }

	@Override
	public String[] getRuleNames() { return ruleNames; }

	@Override
	public String getSerializedATN() { return _serializedATN; }

	@Override
	public ATN getATN() { return _ATN; }

	public Chems(TokenStream input) {
		super(input);
		_interp = new ParserATNSimulator(this,_ATN,_decisionToDFA,_sharedContextCache);
	}

	public static class StartContext extends ParserRuleContext {
		public *arrayList.List lista;
		public InstruccionesContext instrucciones;
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public StartContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_start; }
	}

	public final StartContext start() throws RecognitionException {
		StartContext _localctx = new StartContext(_ctx, getState());
		enterRule(_localctx, 0, RULE_start);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(70);
			((StartContext)_localctx).instrucciones = instrucciones();
			_localctx.lista = ((StartContext)_localctx).instrucciones.l
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionesContext extends ParserRuleContext {
		public *arrayList.List l;
		public InstruccionContext instruccion;
		public List<InstruccionContext> e = new ArrayList<InstruccionContext>();
		public List<InstruccionContext> instruccion() {
			return getRuleContexts(InstruccionContext.class);
		}
		public InstruccionContext instruccion(int i) {
			return getRuleContext(InstruccionContext.class,i);
		}
		public InstruccionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instrucciones; }
	}

	public final InstruccionesContext instrucciones() throws RecognitionException {
		InstruccionesContext _localctx = new InstruccionesContext(_ctx, getState());
		enterRule(_localctx, 2, RULE_instrucciones);

		    _localctx.l =  arrayList.New()
		  
		int _la;
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(76);
			_errHandler.sync(this);
			_la = _input.LA(1);
			while (((((_la - 33)) & ~0x3f) == 0 && ((1L << (_la - 33)) & ((1L << (CONSOLE - 33)) | (1L << (PRINTLN - 33)) | (1L << (P_IF - 33)) | (1L << (P_WHILE - 33)) | (1L << (P_FOR - 33)) | (1L << (BREAK - 33)) | (1L << (CONTINUE - 33)) | (1L << (LET - 33)) | (1L << (MATCH - 33)) | (1L << (LOOP - 33)) | (1L << (ID - 33)))) != 0)) {
				{
				{
				setState(73);
				((InstruccionesContext)_localctx).instruccion = instruccion();
				((InstruccionesContext)_localctx).e.add(((InstruccionesContext)_localctx).instruccion);
				}
				}
				setState(78);
				_errHandler.sync(this);
				_la = _input.LA(1);
			}

			      listInt := localctx.(*InstruccionesContext).GetE()
			      		for _, e := range listInt {
			            _localctx.l.Add(e.GetInstr())
			          }
			    
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class InstruccionContext extends ParserRuleContext {
		public interfaces.Instruction instr;
		public ExpressionContext expression;
		public Declaracion_varContext declaracion_var;
		public Asignacion_varContext asignacion_var;
		public Token LLAVEIZQ;
		public InstruccionesContext instrucciones;
		public Token LLAVEDER;
		public BreaksContext breaks;
		public ContinuesContext continues;
		public IfsContext ifs;
		public LoopsContext loops;
		public ImprContext impr;
		public MatchesContext matches;
		public RforContext rfor;
		public TerminalNode CONSOLE() { return getToken(Chems.CONSOLE, 0); }
		public TerminalNode PUNTO() { return getToken(Chems.PUNTO, 0); }
		public TerminalNode LOG() { return getToken(Chems.LOG, 0); }
		public TerminalNode PARIZQ() { return getToken(Chems.PARIZQ, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Chems.PARDER, 0); }
		public TerminalNode PTCOMA() { return getToken(Chems.PTCOMA, 0); }
		public Declaracion_varContext declaracion_var() {
			return getRuleContext(Declaracion_varContext.class,0);
		}
		public Asignacion_varContext asignacion_var() {
			return getRuleContext(Asignacion_varContext.class,0);
		}
		public TerminalNode P_WHILE() { return getToken(Chems.P_WHILE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public BreaksContext breaks() {
			return getRuleContext(BreaksContext.class,0);
		}
		public ContinuesContext continues() {
			return getRuleContext(ContinuesContext.class,0);
		}
		public IfsContext ifs() {
			return getRuleContext(IfsContext.class,0);
		}
		public LoopsContext loops() {
			return getRuleContext(LoopsContext.class,0);
		}
		public ImprContext impr() {
			return getRuleContext(ImprContext.class,0);
		}
		public MatchesContext matches() {
			return getRuleContext(MatchesContext.class,0);
		}
		public RforContext rfor() {
			return getRuleContext(RforContext.class,0);
		}
		public InstruccionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion; }
	}

	public final InstruccionContext instruccion() throws RecognitionException {
		InstruccionContext _localctx = new InstruccionContext(_ctx, getState());
		enterRule(_localctx, 4, RULE_instruccion);
		try {
			setState(127);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONSOLE:
				enterOuterAlt(_localctx, 1);
				{
				setState(81);
				match(CONSOLE);
				setState(82);
				match(PUNTO);
				setState(83);
				match(LOG);
				setState(84);
				match(PARIZQ);
				setState(85);
				((InstruccionContext)_localctx).expression = expression();
				setState(86);
				match(PARDER);
				setState(87);
				match(PTCOMA);
				_localctx.instr = instruction.NewImprimir(((InstruccionContext)_localctx).expression.p)
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(90);
				((InstruccionContext)_localctx).declaracion_var = declaracion_var();
				setState(91);
				match(PTCOMA);
				_localctx.instr=((InstruccionContext)_localctx).declaracion_var.i
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(94);
				((InstruccionContext)_localctx).asignacion_var = asignacion_var();
				setState(95);
				match(PTCOMA);
				_localctx.instr=((InstruccionContext)_localctx).asignacion_var.i
				}
				break;
			case P_WHILE:
				enterOuterAlt(_localctx, 4);
				{
				setState(98);
				match(P_WHILE);
				setState(99);
				((InstruccionContext)_localctx).expression = expression();
				setState(100);
				((InstruccionContext)_localctx).LLAVEIZQ = match(LLAVEIZQ);
				setState(101);
				((InstruccionContext)_localctx).instrucciones = instrucciones();
				setState(102);
				((InstruccionContext)_localctx).LLAVEDER = match(LLAVEDER);
				_localctx.instr = instruction.NewWhile(((InstruccionContext)_localctx).expression.p, ((InstruccionContext)_localctx).instrucciones.l,((InstruccionContext)_localctx).LLAVEIZQ.GetLine(),((InstruccionContext)_localctx).LLAVEDER.GetColumn())
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 5);
				{
				setState(105);
				((InstruccionContext)_localctx).breaks = breaks();
				_localctx.instr=((InstruccionContext)_localctx).breaks.i
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 6);
				{
				setState(108);
				((InstruccionContext)_localctx).continues = continues();
				_localctx.instr=((InstruccionContext)_localctx).continues.i
				}
				break;
			case P_IF:
				enterOuterAlt(_localctx, 7);
				{
				setState(111);
				((InstruccionContext)_localctx).ifs = ifs();
				_localctx.instr=((InstruccionContext)_localctx).ifs.p
				}
				break;
			case LOOP:
				enterOuterAlt(_localctx, 8);
				{
				setState(114);
				((InstruccionContext)_localctx).loops = loops();
				_localctx.instr=((InstruccionContext)_localctx).loops.i
				}
				break;
			case PRINTLN:
				enterOuterAlt(_localctx, 9);
				{
				setState(117);
				((InstruccionContext)_localctx).impr = impr();
				setState(118);
				match(PTCOMA);
				_localctx.instr=((InstruccionContext)_localctx).impr.p
				}
				break;
			case MATCH:
				enterOuterAlt(_localctx, 10);
				{
				setState(121);
				((InstruccionContext)_localctx).matches = matches();
				_localctx.instr=((InstruccionContext)_localctx).matches.m
				}
				break;
			case P_FOR:
				enterOuterAlt(_localctx, 11);
				{
				setState(124);
				((InstruccionContext)_localctx).rfor = rfor();
				_localctx.instr=((InstruccionContext)_localctx).rfor.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Instruccion_wcContext extends ParserRuleContext {
		public interfaces.Instruction instr;
		public ExpressionContext expression;
		public Declaracion_varContext declaracion_var;
		public Asignacion_varContext asignacion_var;
		public Token LLAVEIZQ;
		public InstruccionesContext instrucciones;
		public Token LLAVEDER;
		public BreaksContext breaks;
		public ContinuesContext continues;
		public IfsContext ifs;
		public LoopsContext loops;
		public ImprContext impr;
		public RforContext rfor;
		public TerminalNode CONSOLE() { return getToken(Chems.CONSOLE, 0); }
		public TerminalNode PUNTO() { return getToken(Chems.PUNTO, 0); }
		public TerminalNode LOG() { return getToken(Chems.LOG, 0); }
		public TerminalNode PARIZQ() { return getToken(Chems.PARIZQ, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Chems.PARDER, 0); }
		public Declaracion_varContext declaracion_var() {
			return getRuleContext(Declaracion_varContext.class,0);
		}
		public Asignacion_varContext asignacion_var() {
			return getRuleContext(Asignacion_varContext.class,0);
		}
		public TerminalNode P_WHILE() { return getToken(Chems.P_WHILE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public BreaksContext breaks() {
			return getRuleContext(BreaksContext.class,0);
		}
		public ContinuesContext continues() {
			return getRuleContext(ContinuesContext.class,0);
		}
		public IfsContext ifs() {
			return getRuleContext(IfsContext.class,0);
		}
		public LoopsContext loops() {
			return getRuleContext(LoopsContext.class,0);
		}
		public ImprContext impr() {
			return getRuleContext(ImprContext.class,0);
		}
		public RforContext rfor() {
			return getRuleContext(RforContext.class,0);
		}
		public Instruccion_wcContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_instruccion_wc; }
	}

	public final Instruccion_wcContext instruccion_wc() throws RecognitionException {
		Instruccion_wcContext _localctx = new Instruccion_wcContext(_ctx, getState());
		enterRule(_localctx, 6, RULE_instruccion_wc);
		try {
			setState(168);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case CONSOLE:
				enterOuterAlt(_localctx, 1);
				{
				setState(129);
				match(CONSOLE);
				setState(130);
				match(PUNTO);
				setState(131);
				match(LOG);
				setState(132);
				match(PARIZQ);
				setState(133);
				((Instruccion_wcContext)_localctx).expression = expression();
				setState(134);
				match(PARDER);
				_localctx.instr = instruction.NewImprimir(((Instruccion_wcContext)_localctx).expression.p)
				}
				break;
			case LET:
				enterOuterAlt(_localctx, 2);
				{
				setState(137);
				((Instruccion_wcContext)_localctx).declaracion_var = declaracion_var();
				_localctx.instr=((Instruccion_wcContext)_localctx).declaracion_var.i
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 3);
				{
				setState(140);
				((Instruccion_wcContext)_localctx).asignacion_var = asignacion_var();
				_localctx.instr=((Instruccion_wcContext)_localctx).asignacion_var.i
				}
				break;
			case P_WHILE:
				enterOuterAlt(_localctx, 4);
				{
				setState(143);
				match(P_WHILE);
				setState(144);
				((Instruccion_wcContext)_localctx).expression = expression();
				setState(145);
				((Instruccion_wcContext)_localctx).LLAVEIZQ = match(LLAVEIZQ);
				setState(146);
				((Instruccion_wcContext)_localctx).instrucciones = instrucciones();
				setState(147);
				((Instruccion_wcContext)_localctx).LLAVEDER = match(LLAVEDER);
				_localctx.instr = instruction.NewWhile(((Instruccion_wcContext)_localctx).expression.p, ((Instruccion_wcContext)_localctx).instrucciones.l,((Instruccion_wcContext)_localctx).LLAVEIZQ.GetLine(),((Instruccion_wcContext)_localctx).LLAVEDER.GetColumn())
				}
				break;
			case BREAK:
				enterOuterAlt(_localctx, 5);
				{
				setState(150);
				((Instruccion_wcContext)_localctx).breaks = breaks();
				_localctx.instr=((Instruccion_wcContext)_localctx).breaks.i
				}
				break;
			case CONTINUE:
				enterOuterAlt(_localctx, 6);
				{
				setState(153);
				((Instruccion_wcContext)_localctx).continues = continues();
				_localctx.instr=((Instruccion_wcContext)_localctx).continues.i
				}
				break;
			case P_IF:
				enterOuterAlt(_localctx, 7);
				{
				setState(156);
				((Instruccion_wcContext)_localctx).ifs = ifs();
				_localctx.instr=((Instruccion_wcContext)_localctx).ifs.p
				}
				break;
			case LOOP:
				enterOuterAlt(_localctx, 8);
				{
				setState(159);
				((Instruccion_wcContext)_localctx).loops = loops();
				_localctx.instr=((Instruccion_wcContext)_localctx).loops.i
				}
				break;
			case PRINTLN:
				enterOuterAlt(_localctx, 9);
				{
				setState(162);
				((Instruccion_wcContext)_localctx).impr = impr();
				_localctx.instr=((Instruccion_wcContext)_localctx).impr.p
				}
				break;
			case P_FOR:
				enterOuterAlt(_localctx, 10);
				{
				setState(165);
				((Instruccion_wcContext)_localctx).rfor = rfor();
				_localctx.instr=((Instruccion_wcContext)_localctx).rfor.p
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Asignacion_varContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public Token id;
		public ExpressionContext expression;
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public Asignacion_varContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignacion_var; }
	}

	public final Asignacion_varContext asignacion_var() throws RecognitionException {
		Asignacion_varContext _localctx = new Asignacion_varContext(_ctx, getState());
		enterRule(_localctx, 8, RULE_asignacion_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(170);
			((Asignacion_varContext)_localctx).id = match(ID);
			setState(171);
			match(IGUAL);
			setState(172);
			((Asignacion_varContext)_localctx).expression = expression();

			    linea:=_localctx.id.GetLine()
			    col:=_localctx.id.GetColumn()
			    _localctx.i=instruction.NewAssignment((((Asignacion_varContext)_localctx).id!=null?((Asignacion_varContext)_localctx).id.getText():null),((Asignacion_varContext)_localctx).expression.p,linea,col)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Declaracion_varContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public Token LET;
		public MutableContext mutable;
		public Token id;
		public TypesContext types;
		public ExpressionContext expression;
		public TerminalNode LET() { return getToken(Chems.LET, 0); }
		public MutableContext mutable() {
			return getRuleContext(MutableContext.class,0);
		}
		public TypesContext types() {
			return getRuleContext(TypesContext.class,0);
		}
		public TerminalNode IGUAL() { return getToken(Chems.IGUAL, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public Declaracion_varContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_declaracion_var; }
	}

	public final Declaracion_varContext declaracion_var() throws RecognitionException {
		Declaracion_varContext _localctx = new Declaracion_varContext(_ctx, getState());
		enterRule(_localctx, 10, RULE_declaracion_var);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(175);
			((Declaracion_varContext)_localctx).LET = match(LET);
			setState(176);
			((Declaracion_varContext)_localctx).mutable = mutable();
			setState(177);
			((Declaracion_varContext)_localctx).id = match(ID);
			setState(178);
			((Declaracion_varContext)_localctx).types = types();
			setState(179);
			match(IGUAL);
			setState(180);
			((Declaracion_varContext)_localctx).expression = expression();
			_localctx.i=instruction.NewDeclaracion((((Declaracion_varContext)_localctx).id!=null?((Declaracion_varContext)_localctx).id.getText():null),((Declaracion_varContext)_localctx).types.l,((Declaracion_varContext)_localctx).expression.p,((Declaracion_varContext)_localctx).mutable.mut,((Declaracion_varContext)_localctx).LET.GetLine(),((Declaracion_varContext)_localctx).LET.GetColumn())
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MutableContext extends ParserRuleContext {
		public bool mut;
		public TerminalNode MUT() { return getToken(Chems.MUT, 0); }
		public MutableContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_mutable; }
	}

	public final MutableContext mutable() throws RecognitionException {
		MutableContext _localctx = new MutableContext(_ctx, getState());
		enterRule(_localctx, 12, RULE_mutable);
		try {
			setState(186);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case MUT:
				enterOuterAlt(_localctx, 1);
				{
				setState(183);
				match(MUT);
				_localctx.mut=true
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 2);
				{
				_localctx.mut=false
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class TypesContext extends ParserRuleContext {
		public interfaces.TipoSimbolo l;
		public Tipo_vectorContext tipo_vector;
		public Asignar_ArrayContext a;
		public Tipo_dContext tipo_d;
		public TerminalNode DOSPUNTOS() { return getToken(Chems.DOSPUNTOS, 0); }
		public Tipo_vectorContext tipo_vector() {
			return getRuleContext(Tipo_vectorContext.class,0);
		}
		public Asignar_ArrayContext asignar_Array() {
			return getRuleContext(Asignar_ArrayContext.class,0);
		}
		public Tipo_dContext tipo_d() {
			return getRuleContext(Tipo_dContext.class,0);
		}
		public TypesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_types; }
	}

	public final TypesContext types() throws RecognitionException {
		TypesContext _localctx = new TypesContext(_ctx, getState());
		enterRule(_localctx, 14, RULE_types);
		try {
			setState(199);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,4,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(188);
				match(DOSPUNTOS);
				setState(189);
				((TypesContext)_localctx).tipo_vector = tipo_vector();
				_localctx.l=((TypesContext)_localctx).tipo_vector.t
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(192);
				((TypesContext)_localctx).a = asignar_Array();

				    dim:=arrayList.New()
				    dim.Add(((TypesContext)_localctx).a.d)
				    _localctx.l=interfaces.TipoSimbolo{interfaces.ARRAY,dim}
				  
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(195);
				match(DOSPUNTOS);
				setState(196);
				((TypesContext)_localctx).tipo_d = tipo_d();
				_localctx.l=interfaces.TipoSimbolo{((TypesContext)_localctx).tipo_d.t,arrayList.New()}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tipo_dContext extends ParserRuleContext {
		public interfaces.TipoExpresion t;
		public TerminalNode INT() { return getToken(Chems.INT, 0); }
		public TerminalNode FLOAT() { return getToken(Chems.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(Chems.BOOL, 0); }
		public TerminalNode CHAR() { return getToken(Chems.CHAR, 0); }
		public TerminalNode STR() { return getToken(Chems.STR, 0); }
		public TerminalNode P_STRING() { return getToken(Chems.P_STRING, 0); }
		public TerminalNode USIZE() { return getToken(Chems.USIZE, 0); }
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public Tipo_dContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_d; }
	}

	public final Tipo_dContext tipo_d() throws RecognitionException {
		Tipo_dContext _localctx = new Tipo_dContext(_ctx, getState());
		enterRule(_localctx, 16, RULE_tipo_d);
		try {
			setState(217);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(201);
				match(INT);
				_localctx.t=interfaces.INTEGER
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(203);
				match(FLOAT);
				_localctx.t=interfaces.FLOAT
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(205);
				match(BOOL);
				_localctx.t=interfaces.BOOLEAN
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(207);
				match(CHAR);
				_localctx.t=interfaces.CHAR
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 5);
				{
				setState(209);
				match(STR);
				_localctx.t=interfaces.STRING
				}
				break;
			case P_STRING:
				enterOuterAlt(_localctx, 6);
				{
				setState(211);
				match(P_STRING);
				_localctx.t=interfaces.STR
				}
				break;
			case USIZE:
				enterOuterAlt(_localctx, 7);
				{
				setState(213);
				match(USIZE);
				_localctx.t=interfaces.USIZE
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 8);
				{
				setState(215);
				match(ID);
				_localctx.t=interfaces.STRUCT
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Asignar_ArrayContext extends ParserRuleContext {
		public interfaces.Dimensions d;
		public DimensionesContext dimensiones;
		public TerminalNode DOSPUNTOS() { return getToken(Chems.DOSPUNTOS, 0); }
		public DimensionesContext dimensiones() {
			return getRuleContext(DimensionesContext.class,0);
		}
		public Asignar_ArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_asignar_Array; }
	}

	public final Asignar_ArrayContext asignar_Array() throws RecognitionException {
		Asignar_ArrayContext _localctx = new Asignar_ArrayContext(_ctx, getState());
		enterRule(_localctx, 18, RULE_asignar_Array);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(219);
			match(DOSPUNTOS);
			setState(220);
			((Asignar_ArrayContext)_localctx).dimensiones = dimensiones();
			_localctx.d=((Asignar_ArrayContext)_localctx).dimensiones.d
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class DimensionesContext extends ParserRuleContext {
		public interfaces.Dimensions d;
		public Tipo_dContext tipo_d;
		public ExpressionContext expression;
		public DimensionesContext dimensiones;
		public TerminalNode CORIZQ() { return getToken(Chems.CORIZQ, 0); }
		public Tipo_dContext tipo_d() {
			return getRuleContext(Tipo_dContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Chems.PTCOMA, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(Chems.CORDER, 0); }
		public DimensionesContext dimensiones() {
			return getRuleContext(DimensionesContext.class,0);
		}
		public DimensionesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_dimensiones; }
	}

	public final DimensionesContext dimensiones() throws RecognitionException {
		DimensionesContext _localctx = new DimensionesContext(_ctx, getState());
		enterRule(_localctx, 20, RULE_dimensiones);
		try {
			setState(237);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,6,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(223);
				match(CORIZQ);
				setState(224);
				((DimensionesContext)_localctx).tipo_d = tipo_d();
				setState(225);
				match(PTCOMA);
				setState(226);
				((DimensionesContext)_localctx).expression = expression();
				setState(227);
				match(CORDER);

				    list:=arrayList.New()
				    list.Add(((DimensionesContext)_localctx).expression.p)
				    _localctx.d=interfaces.Dimensions{((DimensionesContext)_localctx).tipo_d.t,list}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(230);
				match(CORIZQ);
				setState(231);
				((DimensionesContext)_localctx).dimensiones = dimensiones();
				setState(232);
				match(PTCOMA);
				setState(233);
				((DimensionesContext)_localctx).expression = expression();
				setState(234);
				match(CORDER);
				((DimensionesContext)_localctx).dimensiones.d.Dimensions.Add(((DimensionesContext)_localctx).expression.p)
				                                                  ((DimensionesContext)_localctx).d = ((DimensionesContext)_localctx).dimensiones.d;
				                                                
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Tipo_vectorContext extends ParserRuleContext {
		public interfaces.TipoSimbolo t;
		public VectoresContext vectores;
		public TerminalNode VECT() { return getToken(Chems.VECT, 0); }
		public TerminalNode MENOR() { return getToken(Chems.MENOR, 0); }
		public VectoresContext vectores() {
			return getRuleContext(VectoresContext.class,0);
		}
		public TerminalNode MAYOR() { return getToken(Chems.MAYOR, 0); }
		public Tipo_vectorContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_tipo_vector; }
	}

	public final Tipo_vectorContext tipo_vector() throws RecognitionException {
		Tipo_vectorContext _localctx = new Tipo_vectorContext(_ctx, getState());
		enterRule(_localctx, 22, RULE_tipo_vector);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(239);
			match(VECT);
			setState(240);
			match(MENOR);
			setState(241);
			((Tipo_vectorContext)_localctx).vectores = vectores();
			setState(242);
			match(MAYOR);
			_localctx.t=interfaces.TipoSimbolo{interfaces.VECTOR,((Tipo_vectorContext)_localctx).vectores.l}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class VectoresContext extends ParserRuleContext {
		public *arrayList.List l;
		public Token id;
		public VectoresContext vectores;
		public TerminalNode INT() { return getToken(Chems.INT, 0); }
		public TerminalNode FLOAT() { return getToken(Chems.FLOAT, 0); }
		public TerminalNode BOOL() { return getToken(Chems.BOOL, 0); }
		public TerminalNode CHAR() { return getToken(Chems.CHAR, 0); }
		public TerminalNode STR() { return getToken(Chems.STR, 0); }
		public TerminalNode P_STRING() { return getToken(Chems.P_STRING, 0); }
		public TerminalNode USIZE() { return getToken(Chems.USIZE, 0); }
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public TerminalNode VECT() { return getToken(Chems.VECT, 0); }
		public TerminalNode MENOR() { return getToken(Chems.MENOR, 0); }
		public VectoresContext vectores() {
			return getRuleContext(VectoresContext.class,0);
		}
		public TerminalNode MAYOR() { return getToken(Chems.MAYOR, 0); }
		public VectoresContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_vectores; }
	}

	public final VectoresContext vectores() throws RecognitionException {
		VectoresContext _localctx = new VectoresContext(_ctx, getState());
		enterRule(_localctx, 24, RULE_vectores);
		try {
			setState(267);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case INT:
				enterOuterAlt(_localctx, 1);
				{
				setState(245);
				match(INT);
				_localctx.l=arrayList.New()
				            _localctx.l.Add(interfaces.INTEGER)
				}
				break;
			case FLOAT:
				enterOuterAlt(_localctx, 2);
				{
				setState(247);
				match(FLOAT);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.FLOAT)
				}
				break;
			case BOOL:
				enterOuterAlt(_localctx, 3);
				{
				setState(249);
				match(BOOL);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.BOOLEAN)
				}
				break;
			case CHAR:
				enterOuterAlt(_localctx, 4);
				{
				setState(251);
				match(CHAR);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.CHAR)
				}
				break;
			case STR:
				enterOuterAlt(_localctx, 5);
				{
				setState(253);
				match(STR);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.STRING)
				}
				break;
			case P_STRING:
				enterOuterAlt(_localctx, 6);
				{
				setState(255);
				match(P_STRING);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.STR)
				}
				break;
			case USIZE:
				enterOuterAlt(_localctx, 7);
				{
				setState(257);
				match(USIZE);
				_localctx.l=arrayList.New()
				              _localctx.l.Add(interfaces.USIZE)
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 8);
				{
				setState(259);
				((VectoresContext)_localctx).id = match(ID);
				_localctx.l=arrayList.New()
				              _localctx.l.Add((((VectoresContext)_localctx).id!=null?((VectoresContext)_localctx).id.getText():null))
				              _localctx.l.Add(interfaces.STRUCT)
				}
				break;
			case VECT:
				enterOuterAlt(_localctx, 9);
				{
				setState(261);
				match(VECT);
				setState(262);
				match(MENOR);
				setState(263);
				((VectoresContext)_localctx).vectores = vectores();
				setState(264);
				match(MAYOR);

				    ((VectoresContext)_localctx).vectores.l.Add(interfaces.VECTOR)
				    _localctx.l=((VectoresContext)_localctx).vectores.l
				  
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class LoopsContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public InstruccionesContext instrucciones;
		public TerminalNode LOOP() { return getToken(Chems.LOOP, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public LoopsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_loops; }
	}

	public final LoopsContext loops() throws RecognitionException {
		LoopsContext _localctx = new LoopsContext(_ctx, getState());
		enterRule(_localctx, 26, RULE_loops);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(269);
			match(LOOP);
			setState(270);
			match(LLAVEIZQ);
			setState(271);
			((LoopsContext)_localctx).instrucciones = instrucciones();
			setState(272);
			match(LLAVEDER);
			_localctx.i=instruction.NewLoop(((LoopsContext)_localctx).instrucciones.l)
			    fmt.Println("TOY EN EL ANALISIS")
			  
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class IfsContext extends ParserRuleContext {
		public interfaces.Instruction p;
		public Token P_IF;
		public ExpressionContext expression;
		public Llaves_ifContext llaves_if;
		public ElsesContext elses;
		public TerminalNode P_IF() { return getToken(Chems.P_IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public Llaves_ifContext llaves_if() {
			return getRuleContext(Llaves_ifContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public ElsesContext elses() {
			return getRuleContext(ElsesContext.class,0);
		}
		public IfsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_ifs; }
	}

	public final IfsContext ifs() throws RecognitionException {
		IfsContext _localctx = new IfsContext(_ctx, getState());
		enterRule(_localctx, 28, RULE_ifs);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(275);
			((IfsContext)_localctx).P_IF = match(P_IF);
			setState(276);
			((IfsContext)_localctx).expression = expression();
			setState(277);
			match(LLAVEIZQ);
			setState(278);
			((IfsContext)_localctx).llaves_if = llaves_if(0);
			setState(279);
			match(LLAVEDER);
			setState(280);
			((IfsContext)_localctx).elses = elses();
			_localctx.p=instruction.NewIf(((IfsContext)_localctx).expression.p,((IfsContext)_localctx).llaves_if.l,((IfsContext)_localctx).elses.e,((IfsContext)_localctx).P_IF.GetLine(),((IfsContext)_localctx).P_IF.GetColumn(),0)
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Llaves_ifContext extends ParserRuleContext {
		public *arrayList.List l;
		public Llaves_ifContext k;
		public ExpressionContext expression;
		public InstruccionContext instruccion;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public InstruccionContext instruccion() {
			return getRuleContext(InstruccionContext.class,0);
		}
		public Llaves_ifContext llaves_if() {
			return getRuleContext(Llaves_ifContext.class,0);
		}
		public Llaves_ifContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_llaves_if; }
	}

	public final Llaves_ifContext llaves_if() throws RecognitionException {
		return llaves_if(0);
	}

	private Llaves_ifContext llaves_if(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Llaves_ifContext _localctx = new Llaves_ifContext(_ctx, _parentState);
		Llaves_ifContext _prevctx = _localctx;
		int _startState = 30;
		enterRecursionRule(_localctx, 30, RULE_llaves_if, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(290);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,8,_ctx) ) {
			case 1:
				{
				setState(284);
				((Llaves_ifContext)_localctx).expression = expression();
				_localctx.l=arrayList.New()
				                            i:=interfaces.OpcionIf{0,((Llaves_ifContext)_localctx).expression.p}
				                            _localctx.l.Add(i)
				}
				break;
			case 2:
				{
				setState(287);
				((Llaves_ifContext)_localctx).instruccion = instruccion();
				_localctx.l=arrayList.New()
				                            i:=interfaces.OpcionIf{1,((Llaves_ifContext)_localctx).instruccion.instr}
				                            _localctx.l.Add(i)
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(302);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(300);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,9,_ctx) ) {
					case 1:
						{
						_localctx = new Llaves_ifContext(_parentctx, _parentState);
						_localctx.k = _prevctx;
						_localctx.k = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_llaves_if);
						setState(292);
						if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
						setState(293);
						((Llaves_ifContext)_localctx).expression = expression();

						                                      i:=interfaces.OpcionIf{0,((Llaves_ifContext)_localctx).expression.p}
						                                      ((Llaves_ifContext)_localctx).k.l.Add(i)
						                                      _localctx.l=((Llaves_ifContext)_localctx).k.l
						}
						break;
					case 2:
						{
						_localctx = new Llaves_ifContext(_parentctx, _parentState);
						_localctx.k = _prevctx;
						_localctx.k = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_llaves_if);
						setState(296);
						if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
						setState(297);
						((Llaves_ifContext)_localctx).instruccion = instruccion();

						                                      i:=interfaces.OpcionIf{1,((Llaves_ifContext)_localctx).instruccion.instr}
						                                      ((Llaves_ifContext)_localctx).k.l.Add(i)
						                                      _localctx.l=((Llaves_ifContext)_localctx).k.l
						}
						break;
					}
					} 
				}
				setState(304);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,10,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ElsesContext extends ParserRuleContext {
		public interfaces.Instruction e;
		public Llaves_ifContext llaves_if;
		public Token P_IF;
		public ExpressionContext expression;
		public ElsesContext elses;
		public TerminalNode P_ELSE() { return getToken(Chems.P_ELSE, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public Llaves_ifContext llaves_if() {
			return getRuleContext(Llaves_ifContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public TerminalNode P_IF() { return getToken(Chems.P_IF, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public ElsesContext elses() {
			return getRuleContext(ElsesContext.class,0);
		}
		public ElsesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_elses; }
	}

	public final ElsesContext elses() throws RecognitionException {
		ElsesContext _localctx = new ElsesContext(_ctx, getState());
		enterRule(_localctx, 32, RULE_elses);
		try {
			setState(321);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,11,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(305);
				match(P_ELSE);
				setState(306);
				match(LLAVEIZQ);
				setState(307);
				((ElsesContext)_localctx).llaves_if = llaves_if(0);
				setState(308);
				match(LLAVEDER);
				_localctx.e=instruction.NewIf(expresion.NewPrimitivo(1,interfaces.BOOLEAN,0,0),((ElsesContext)_localctx).llaves_if.l,instruction.NewElseNull("null"),0,0,3)
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(311);
				match(P_ELSE);
				setState(312);
				((ElsesContext)_localctx).P_IF = match(P_IF);
				setState(313);
				((ElsesContext)_localctx).expression = expression();
				setState(314);
				match(LLAVEIZQ);
				setState(315);
				((ElsesContext)_localctx).llaves_if = llaves_if(0);
				setState(316);
				match(LLAVEDER);
				setState(317);
				((ElsesContext)_localctx).elses = elses();
				_localctx.e=instruction.NewIf(((ElsesContext)_localctx).expression.p,((ElsesContext)_localctx).llaves_if.l,((ElsesContext)_localctx).elses.e,((ElsesContext)_localctx).P_IF.GetLine(),((ElsesContext)_localctx).P_IF.GetColumn(),2)
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				_localctx.e= instruction.NewElseNull("null")
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class BreaksContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public Token BREAK;
		public ExpressionContext expression;
		public TerminalNode BREAK() { return getToken(Chems.BREAK, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode PTCOMA() { return getToken(Chems.PTCOMA, 0); }
		public BreaksContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_breaks; }
	}

	public final BreaksContext breaks() throws RecognitionException {
		BreaksContext _localctx = new BreaksContext(_ctx, getState());
		enterRule(_localctx, 34, RULE_breaks);
		try {
			setState(331);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,12,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(323);
				((BreaksContext)_localctx).BREAK = match(BREAK);
				setState(324);
				((BreaksContext)_localctx).expression = expression();
				setState(325);
				match(PTCOMA);
				_localctx.i=instruction.NewBreak(((BreaksContext)_localctx).expression.p,true,((BreaksContext)_localctx).BREAK.GetLine(),((BreaksContext)_localctx).BREAK.GetColumn())
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(328);
				((BreaksContext)_localctx).BREAK = match(BREAK);
				setState(329);
				match(PTCOMA);
				_localctx.i=instruction.NewBreak(expresion.NewPrimitivo(1,interfaces.INTEGER,0,0),false,((BreaksContext)_localctx).BREAK.GetLine(),((BreaksContext)_localctx).BREAK.GetColumn())
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ContinuesContext extends ParserRuleContext {
		public interfaces.Instruction i;
		public Token CONTINUE;
		public TerminalNode CONTINUE() { return getToken(Chems.CONTINUE, 0); }
		public TerminalNode PTCOMA() { return getToken(Chems.PTCOMA, 0); }
		public ContinuesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_continues; }
	}

	public final ContinuesContext continues() throws RecognitionException {
		ContinuesContext _localctx = new ContinuesContext(_ctx, getState());
		enterRule(_localctx, 36, RULE_continues);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(333);
			((ContinuesContext)_localctx).CONTINUE = match(CONTINUE);
			setState(334);
			match(PTCOMA);
			_localctx.i=instruction.NewContinue("continue",((ContinuesContext)_localctx).CONTINUE.GetLine(),((ContinuesContext)_localctx).CONTINUE.GetColumn())
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ImprContext extends ParserRuleContext {
		public interfaces.Instruction p;
		public Token PARIZQ;
		public FormatoContext formato;
		public ListValuesContext listValues;
		public TerminalNode PRINTLN() { return getToken(Chems.PRINTLN, 0); }
		public TerminalNode PARIZQ() { return getToken(Chems.PARIZQ, 0); }
		public FormatoContext formato() {
			return getRuleContext(FormatoContext.class,0);
		}
		public ListValuesContext listValues() {
			return getRuleContext(ListValuesContext.class,0);
		}
		public TerminalNode PARDER() { return getToken(Chems.PARDER, 0); }
		public ImprContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_impr; }
	}

	public final ImprContext impr() throws RecognitionException {
		ImprContext _localctx = new ImprContext(_ctx, getState());
		enterRule(_localctx, 38, RULE_impr);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(337);
			match(PRINTLN);
			setState(338);
			((ImprContext)_localctx).PARIZQ = match(PARIZQ);
			setState(339);
			((ImprContext)_localctx).formato = formato();
			setState(340);
			((ImprContext)_localctx).listValues = listValues(0);
			setState(341);
			match(PARDER);
			_localctx.p=instruction.NewPrint(((ImprContext)_localctx).listValues.l,((ImprContext)_localctx).formato.s,((ImprContext)_localctx).PARIZQ.GetLine(),((ImprContext)_localctx).PARIZQ.GetColumn())
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class FormatoContext extends ParserRuleContext {
		public string s;
		public Token STRING;
		public TerminalNode STRING() { return getToken(Chems.STRING, 0); }
		public TerminalNode COMA() { return getToken(Chems.COMA, 0); }
		public FormatoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_formato; }
	}

	public final FormatoContext formato() throws RecognitionException {
		FormatoContext _localctx = new FormatoContext(_ctx, getState());
		enterRule(_localctx, 40, RULE_formato);
		try {
			setState(348);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,13,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(344);
				((FormatoContext)_localctx).STRING = match(STRING);
				setState(345);
				match(COMA);
				str:= (((FormatoContext)_localctx).STRING!=null?((FormatoContext)_localctx).STRING.getText():null)[1:len((((FormatoContext)_localctx).STRING!=null?((FormatoContext)_localctx).STRING.getText():null))-1]
				      _localctx.s=str
				   
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				_localctx.s=""
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class MatchesContext extends ParserRuleContext {
		public interfaces.Instruction m;
		public Token MATCH;
		public ExpressionContext expression;
		public CasosContext casos;
		public TerminalNode MATCH() { return getToken(Chems.MATCH, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public CasosContext casos() {
			return getRuleContext(CasosContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public MatchesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_matches; }
	}

	public final MatchesContext matches() throws RecognitionException {
		MatchesContext _localctx = new MatchesContext(_ctx, getState());
		enterRule(_localctx, 42, RULE_matches);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(350);
			((MatchesContext)_localctx).MATCH = match(MATCH);
			setState(351);
			((MatchesContext)_localctx).expression = expression();
			setState(352);
			match(LLAVEIZQ);
			setState(353);
			((MatchesContext)_localctx).casos = casos(0);
			setState(354);
			match(LLAVEDER);
			_localctx.m=instruction.NewMatch(((MatchesContext)_localctx).expression.p,((MatchesContext)_localctx).casos.l,((MatchesContext)_localctx).MATCH.GetLine(),((MatchesContext)_localctx).MATCH.GetColumn())
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class CasosContext extends ParserRuleContext {
		public *arrayList.List l;
		public CasosContext cs;
		public CasesContext cases;
		public DefaultsContext defaults;
		public CasesContext cases() {
			return getRuleContext(CasesContext.class,0);
		}
		public DefaultsContext defaults() {
			return getRuleContext(DefaultsContext.class,0);
		}
		public CasosContext casos() {
			return getRuleContext(CasosContext.class,0);
		}
		public CasosContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_casos; }
	}

	public final CasosContext casos() throws RecognitionException {
		return casos(0);
	}

	private CasosContext casos(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		CasosContext _localctx = new CasosContext(_ctx, _parentState);
		CasosContext _prevctx = _localctx;
		int _startState = 44;
		enterRecursionRule(_localctx, 44, RULE_casos, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(364);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case DIFERENTE:
			case SUB:
			case PARIZQ:
			case CORIZQ:
			case P_IF:
			case INT:
			case FLOAT:
			case TRUE:
			case FALSE:
			case MATCH:
			case LOOP:
			case NUMBER:
			case DECIMAL:
			case STRING:
			case ID:
			case CARACTER:
				{
				setState(358);
				((CasosContext)_localctx).cases = cases();
				_localctx.l=arrayList.New()
				                _localctx.l.Add(((CasosContext)_localctx).cases.c)
				}
				break;
			case DEF:
				{
				setState(361);
				((CasosContext)_localctx).defaults = defaults();
				_localctx.l=arrayList.New()
				                _localctx.l.Add(((CasosContext)_localctx).defaults.c)
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
			_ctx.stop = _input.LT(-1);
			setState(376);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(374);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,15,_ctx) ) {
					case 1:
						{
						_localctx = new CasosContext(_parentctx, _parentState);
						_localctx.cs = _prevctx;
						_localctx.cs = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_casos);
						setState(366);
						if (!(precpred(_ctx, 4))) throw new FailedPredicateException(this, "precpred(_ctx, 4)");
						setState(367);
						((CasosContext)_localctx).cases = cases();
						((CasosContext)_localctx).cs.l.Add(((CasosContext)_localctx).cases.c)
						                                  _localctx.l=((CasosContext)_localctx).cs.l
						}
						break;
					case 2:
						{
						_localctx = new CasosContext(_parentctx, _parentState);
						_localctx.cs = _prevctx;
						_localctx.cs = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_casos);
						setState(370);
						if (!(precpred(_ctx, 3))) throw new FailedPredicateException(this, "precpred(_ctx, 3)");
						setState(371);
						((CasosContext)_localctx).defaults = defaults();
						((CasosContext)_localctx).cs.l.Add(((CasosContext)_localctx).defaults.c)
						                                  _localctx.l=((CasosContext)_localctx).cs.l
						}
						break;
					}
					} 
				}
				setState(378);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,16,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class CasesContext extends ParserRuleContext {
		public interfaces.Cases c;
		public ConditionsContext conditions;
		public Set_matchContext set_match;
		public ConditionsContext conditions() {
			return getRuleContext(ConditionsContext.class,0);
		}
		public TerminalNode OPMATCH() { return getToken(Chems.OPMATCH, 0); }
		public Set_matchContext set_match() {
			return getRuleContext(Set_matchContext.class,0);
		}
		public CasesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_cases; }
	}

	public final CasesContext cases() throws RecognitionException {
		CasesContext _localctx = new CasesContext(_ctx, getState());
		enterRule(_localctx, 46, RULE_cases);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(379);
			((CasesContext)_localctx).conditions = conditions(0);
			setState(380);
			match(OPMATCH);
			setState(381);
			((CasesContext)_localctx).set_match = set_match();
			_localctx.c=interfaces.Cases{((CasesContext)_localctx).conditions.l,((CasesContext)_localctx).set_match.cs.Cuerpo,((CasesContext)_localctx).set_match.cs.Retorno,0,false}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ConditionsContext extends ParserRuleContext {
		public *arrayList.List l;
		public ConditionsContext cond;
		public ExpressionContext expression;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode OR_COND() { return getToken(Chems.OR_COND, 0); }
		public ConditionsContext conditions() {
			return getRuleContext(ConditionsContext.class,0);
		}
		public ConditionsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_conditions; }
	}

	public final ConditionsContext conditions() throws RecognitionException {
		return conditions(0);
	}

	private ConditionsContext conditions(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ConditionsContext _localctx = new ConditionsContext(_ctx, _parentState);
		ConditionsContext _prevctx = _localctx;
		int _startState = 48;
		enterRecursionRule(_localctx, 48, RULE_conditions, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(385);
			((ConditionsContext)_localctx).expression = expression();
			_localctx.l=arrayList.New() 
			                                      _localctx.l.Add(((ConditionsContext)_localctx).expression.p) 
			}
			_ctx.stop = _input.LT(-1);
			setState(395);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ConditionsContext(_parentctx, _parentState);
					_localctx.cond = _prevctx;
					_localctx.cond = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_conditions);
					setState(388);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(389);
					match(OR_COND);
					setState(390);
					((ConditionsContext)_localctx).expression = expression();
					((ConditionsContext)_localctx).cond.l.Add(((ConditionsContext)_localctx).expression.p)
					                                                _localctx.l=((ConditionsContext)_localctx).cond.l
					}
					} 
				}
				setState(397);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,17,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class DefaultsContext extends ParserRuleContext {
		public interfaces.Cases c;
		public Set_matchContext set_match;
		public TerminalNode DEF() { return getToken(Chems.DEF, 0); }
		public TerminalNode OPMATCH() { return getToken(Chems.OPMATCH, 0); }
		public Set_matchContext set_match() {
			return getRuleContext(Set_matchContext.class,0);
		}
		public DefaultsContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_defaults; }
	}

	public final DefaultsContext defaults() throws RecognitionException {
		DefaultsContext _localctx = new DefaultsContext(_ctx, getState());
		enterRule(_localctx, 50, RULE_defaults);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(398);
			match(DEF);
			setState(399);
			match(OPMATCH);
			setState(400);
			((DefaultsContext)_localctx).set_match = set_match();
			_localctx.c=interfaces.Cases{((DefaultsContext)_localctx).set_match.cs.Condicion,((DefaultsContext)_localctx).set_match.cs.Cuerpo,((DefaultsContext)_localctx).set_match.cs.Retorno,1,false}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Set_matchContext extends ParserRuleContext {
		public interfaces.Cases cs;
		public ExpressionContext expression;
		public InstruccionesContext instrucciones;
		public Instruccion_wcContext instruccion_wc;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COMA() { return getToken(Chems.COMA, 0); }
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public Instruccion_wcContext instruccion_wc() {
			return getRuleContext(Instruccion_wcContext.class,0);
		}
		public Set_matchContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_set_match; }
	}

	public final Set_matchContext set_match() throws RecognitionException {
		Set_matchContext _localctx = new Set_matchContext(_ctx, getState());
		enterRule(_localctx, 52, RULE_set_match);
		try {
			setState(416);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,18,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(403);
				((Set_matchContext)_localctx).expression = expression();
				setState(404);
				match(COMA);
				arr:=arrayList.New()
				    arr.Add(instruction.NewElseNull("null"))
				    _localctx.cs=interfaces.Cases{ arrayList.New(),arr,((Set_matchContext)_localctx).expression.p,0,false}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(407);
				match(LLAVEIZQ);
				setState(408);
				((Set_matchContext)_localctx).instrucciones = instrucciones();
				setState(409);
				match(LLAVEDER);
				_localctx.cs=interfaces.Cases{ arrayList.New(),((Set_matchContext)_localctx).instrucciones.l,expresion.NewPrimitivo (1,interfaces.INTEGER,0,0),0,false}
				}
				break;
			case 3:
				enterOuterAlt(_localctx, 3);
				{
				setState(412);
				((Set_matchContext)_localctx).instruccion_wc = instruccion_wc();
				setState(413);
				match(COMA);
				arr:=arrayList.New()
				    arr.Add(((Set_matchContext)_localctx).instruccion_wc.instr)
				    _localctx.cs=interfaces.Cases{ arrayList.New(),arr,expresion.NewPrimitivo (1,interfaces.INTEGER,0,0),0,false}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class RforContext extends ParserRuleContext {
		public interfaces.Instruction p;
		public Token P_FOR;
		public Token id;
		public Iter_forContext iter_for;
		public InstruccionesContext instrucciones;
		public TerminalNode P_FOR() { return getToken(Chems.P_FOR, 0); }
		public TerminalNode P_IN() { return getToken(Chems.P_IN, 0); }
		public Iter_forContext iter_for() {
			return getRuleContext(Iter_forContext.class,0);
		}
		public TerminalNode LLAVEIZQ() { return getToken(Chems.LLAVEIZQ, 0); }
		public InstruccionesContext instrucciones() {
			return getRuleContext(InstruccionesContext.class,0);
		}
		public TerminalNode LLAVEDER() { return getToken(Chems.LLAVEDER, 0); }
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public RforContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_rfor; }
	}

	public final RforContext rfor() throws RecognitionException {
		RforContext _localctx = new RforContext(_ctx, getState());
		enterRule(_localctx, 54, RULE_rfor);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(418);
			((RforContext)_localctx).P_FOR = match(P_FOR);
			setState(419);
			((RforContext)_localctx).id = match(ID);
			setState(420);
			match(P_IN);
			setState(421);
			((RforContext)_localctx).iter_for = iter_for();
			setState(422);
			match(LLAVEIZQ);
			setState(423);
			((RforContext)_localctx).instrucciones = instrucciones();
			setState(424);
			match(LLAVEDER);
			_localctx.p=instruction.NewFor(((RforContext)_localctx).instrucciones.l,(((RforContext)_localctx).id!=null?((RforContext)_localctx).id.getText():null),((RforContext)_localctx).iter_for.p,((RforContext)_localctx).P_FOR.GetLine(),((RforContext)_localctx).P_FOR.GetColumn())
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Iter_forContext extends ParserRuleContext {
		public interfaces.For_Range p;
		public ExpressionContext exp1;
		public ExpressionContext exp2;
		public List<TerminalNode> PUNTO() { return getTokens(Chems.PUNTO); }
		public TerminalNode PUNTO(int i) {
			return getToken(Chems.PUNTO, i);
		}
		public List<ExpressionContext> expression() {
			return getRuleContexts(ExpressionContext.class);
		}
		public ExpressionContext expression(int i) {
			return getRuleContext(ExpressionContext.class,i);
		}
		public Iter_forContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_iter_for; }
	}

	public final Iter_forContext iter_for() throws RecognitionException {
		Iter_forContext _localctx = new Iter_forContext(_ctx, getState());
		enterRule(_localctx, 56, RULE_iter_for);
		try {
			setState(436);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,19,_ctx) ) {
			case 1:
				enterOuterAlt(_localctx, 1);
				{
				setState(427);
				((Iter_forContext)_localctx).exp1 = expression();
				setState(428);
				match(PUNTO);
				setState(429);
				match(PUNTO);
				setState(430);
				((Iter_forContext)_localctx).exp2 = expression();
				_localctx.p=interfaces.For_Range{((Iter_forContext)_localctx).exp1.p,((Iter_forContext)_localctx).exp2.p,0}
				}
				break;
			case 2:
				enterOuterAlt(_localctx, 2);
				{
				setState(433);
				((Iter_forContext)_localctx).exp1 = expression();
				_localctx.p=interfaces.For_Range{((Iter_forContext)_localctx).exp1.p,expresion.NewPrimitivo (1,interfaces.INTEGER,0,0),1}
				}
				break;
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ExpressionContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Expr_aritContext expr_arit;
		public Expr_aritContext expr_arit() {
			return getRuleContext(Expr_aritContext.class,0);
		}
		public ExpressionContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expression; }
	}

	public final ExpressionContext expression() throws RecognitionException {
		ExpressionContext _localctx = new ExpressionContext(_ctx, getState());
		enterRule(_localctx, 58, RULE_expression);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(438);
			((ExpressionContext)_localctx).expr_arit = expr_arit(0);
			_localctx.p = ((ExpressionContext)_localctx).expr_arit.p
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class Expr_aritContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Expr_aritContext exp;
		public Expr_aritContext opIz;
		public Token PARIZQ;
		public Expr_aritContext opDe;
		public Token DIFERENTE;
		public Token SUB;
		public PrimitivoContext primitivo;
		public ExpressionContext expression;
		public LoopsContext loops;
		public IfsContext ifs;
		public MatchesContext matches;
		public ListValuesContext listValues;
		public ArrayAccContext arrayAcc;
		public Token op;
		public Tipo_dContext tipo_d;
		public Token PUNTO;
		public TerminalNode INT() { return getToken(Chems.INT, 0); }
		public TerminalNode DDPUNTO() { return getToken(Chems.DDPUNTO, 0); }
		public TerminalNode POW() { return getToken(Chems.POW, 0); }
		public TerminalNode PARIZQ() { return getToken(Chems.PARIZQ, 0); }
		public TerminalNode COMA() { return getToken(Chems.COMA, 0); }
		public TerminalNode PARDER() { return getToken(Chems.PARDER, 0); }
		public List<Expr_aritContext> expr_arit() {
			return getRuleContexts(Expr_aritContext.class);
		}
		public Expr_aritContext expr_arit(int i) {
			return getRuleContext(Expr_aritContext.class,i);
		}
		public TerminalNode FLOAT() { return getToken(Chems.FLOAT, 0); }
		public TerminalNode POWF() { return getToken(Chems.POWF, 0); }
		public TerminalNode DIFERENTE() { return getToken(Chems.DIFERENTE, 0); }
		public TerminalNode SUB() { return getToken(Chems.SUB, 0); }
		public PrimitivoContext primitivo() {
			return getRuleContext(PrimitivoContext.class,0);
		}
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public LoopsContext loops() {
			return getRuleContext(LoopsContext.class,0);
		}
		public IfsContext ifs() {
			return getRuleContext(IfsContext.class,0);
		}
		public MatchesContext matches() {
			return getRuleContext(MatchesContext.class,0);
		}
		public TerminalNode CORIZQ() { return getToken(Chems.CORIZQ, 0); }
		public ListValuesContext listValues() {
			return getRuleContext(ListValuesContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(Chems.CORDER, 0); }
		public ArrayAccContext arrayAcc() {
			return getRuleContext(ArrayAccContext.class,0);
		}
		public TerminalNode MUL() { return getToken(Chems.MUL, 0); }
		public TerminalNode DIV() { return getToken(Chems.DIV, 0); }
		public TerminalNode MOD() { return getToken(Chems.MOD, 0); }
		public TerminalNode ADD() { return getToken(Chems.ADD, 0); }
		public TerminalNode D_IGUAL() { return getToken(Chems.D_IGUAL, 0); }
		public TerminalNode NOT_E() { return getToken(Chems.NOT_E, 0); }
		public TerminalNode MENOR() { return getToken(Chems.MENOR, 0); }
		public TerminalNode MENORIGUAL() { return getToken(Chems.MENORIGUAL, 0); }
		public TerminalNode MAYORIGUAL() { return getToken(Chems.MAYORIGUAL, 0); }
		public TerminalNode MAYOR() { return getToken(Chems.MAYOR, 0); }
		public TerminalNode OR() { return getToken(Chems.OR, 0); }
		public TerminalNode AND() { return getToken(Chems.AND, 0); }
		public TerminalNode P_AS() { return getToken(Chems.P_AS, 0); }
		public Tipo_dContext tipo_d() {
			return getRuleContext(Tipo_dContext.class,0);
		}
		public TerminalNode PUNTO() { return getToken(Chems.PUNTO, 0); }
		public TerminalNode T_STRING() { return getToken(Chems.T_STRING, 0); }
		public Expr_aritContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_expr_arit; }
	}

	public final Expr_aritContext expr_arit() throws RecognitionException {
		return expr_arit(0);
	}

	private Expr_aritContext expr_arit(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		Expr_aritContext _localctx = new Expr_aritContext(_ctx, _parentState);
		Expr_aritContext _prevctx = _localctx;
		int _startState = 60;
		enterRecursionRule(_localctx, 60, RULE_expr_arit, _p);
		int _la;
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			setState(495);
			_errHandler.sync(this);
			switch ( getInterpreter().adaptivePredict(_input,20,_ctx) ) {
			case 1:
				{
				setState(442);
				match(INT);
				setState(443);
				match(DDPUNTO);
				setState(444);
				match(POW);
				setState(445);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(446);
				((Expr_aritContext)_localctx).opIz = expr_arit(0);
				setState(447);
				match(COMA);
				setState(448);
				((Expr_aritContext)_localctx).opDe = expr_arit(0);
				setState(449);
				match(PARDER);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"^",((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).PARIZQ.GetLine(),((Expr_aritContext)_localctx).PARIZQ.GetColumn())
				}
				break;
			case 2:
				{
				setState(452);
				match(FLOAT);
				setState(453);
				match(DDPUNTO);
				setState(454);
				match(POWF);
				setState(455);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(456);
				((Expr_aritContext)_localctx).opIz = expr_arit(0);
				setState(457);
				match(COMA);
				setState(458);
				((Expr_aritContext)_localctx).opDe = expr_arit(0);
				setState(459);
				match(PARDER);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"^",((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).PARIZQ.GetLine(),((Expr_aritContext)_localctx).PARIZQ.GetColumn())
				}
				break;
			case 3:
				{
				setState(462);
				((Expr_aritContext)_localctx).DIFERENTE = match(DIFERENTE);
				setState(463);
				((Expr_aritContext)_localctx).opIz = expr_arit(13);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"!",((Expr_aritContext)_localctx).opIz.p,true,((Expr_aritContext)_localctx).DIFERENTE.GetLine(),((Expr_aritContext)_localctx).DIFERENTE.GetColumn())
				}
				break;
			case 4:
				{
				setState(466);
				((Expr_aritContext)_localctx).SUB = match(SUB);
				setState(467);
				((Expr_aritContext)_localctx).opIz = expr_arit(12);
				_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,"°",((Expr_aritContext)_localctx).opIz.p,true,((Expr_aritContext)_localctx).SUB.GetLine(),((Expr_aritContext)_localctx).SUB.GetColumn())
				}
				break;
			case 5:
				{
				setState(470);
				((Expr_aritContext)_localctx).primitivo = primitivo();
				_localctx.p = ((Expr_aritContext)_localctx).primitivo.p
				}
				break;
			case 6:
				{
				setState(473);
				((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
				setState(474);
				((Expr_aritContext)_localctx).expression = expression();
				setState(475);
				match(PARDER);
				_localctx.p = ((Expr_aritContext)_localctx).expression.p
				}
				break;
			case 7:
				{
				setState(478);
				((Expr_aritContext)_localctx).loops = loops();
				_localctx.p=expresion.NewDevLoop(((Expr_aritContext)_localctx).loops.i)
				}
				break;
			case 8:
				{
				setState(481);
				((Expr_aritContext)_localctx).ifs = ifs();
				_localctx.p=expresion.NewDevLoop(((Expr_aritContext)_localctx).ifs.p)
				}
				break;
			case 9:
				{
				setState(484);
				((Expr_aritContext)_localctx).matches = matches();
				_localctx.p=expresion.NewDevLoop(((Expr_aritContext)_localctx).matches.m)
				}
				break;
			case 10:
				{
				setState(487);
				match(CORIZQ);
				setState(488);
				((Expr_aritContext)_localctx).listValues = listValues(0);
				setState(489);
				match(CORDER);
				 _localctx.p = expresion.NewArray(((Expr_aritContext)_localctx).listValues.l) 
				}
				break;
			case 11:
				{
				setState(492);
				((Expr_aritContext)_localctx).arrayAcc = arrayAcc();
				_localctx.p=((Expr_aritContext)_localctx).arrayAcc.p
				}
				break;
			}
			_ctx.stop = _input.LT(-1);
			setState(530);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					setState(528);
					_errHandler.sync(this);
					switch ( getInterpreter().adaptivePredict(_input,21,_ctx) ) {
					case 1:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(497);
						if (!(precpred(_ctx, 11))) throw new FailedPredicateException(this, "precpred(_ctx, 11)");
						setState(498);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MUL) | (1L << DIV) | (1L << MOD))) != 0)) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(499);
						((Expr_aritContext)_localctx).opDe = expr_arit(12);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 2:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(502);
						if (!(precpred(_ctx, 10))) throw new FailedPredicateException(this, "precpred(_ctx, 10)");
						setState(503);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==ADD || _la==SUB) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(504);
						((Expr_aritContext)_localctx).opDe = expr_arit(11);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 3:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(507);
						if (!(precpred(_ctx, 9))) throw new FailedPredicateException(this, "precpred(_ctx, 9)");
						setState(508);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !((((_la) & ~0x3f) == 0 && ((1L << _la) & ((1L << MENOR) | (1L << MAYORIGUAL) | (1L << MENORIGUAL) | (1L << D_IGUAL) | (1L << NOT_E) | (1L << MAYOR))) != 0)) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(509);
						((Expr_aritContext)_localctx).opDe = expr_arit(10);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 4:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.opIz = _prevctx;
						_localctx.opIz = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(512);
						if (!(precpred(_ctx, 8))) throw new FailedPredicateException(this, "precpred(_ctx, 8)");
						setState(513);
						((Expr_aritContext)_localctx).op = _input.LT(1);
						_la = _input.LA(1);
						if ( !(_la==OR || _la==AND) ) {
							((Expr_aritContext)_localctx).op = (Token)_errHandler.recoverInline(this);
						}
						else {
							if ( _input.LA(1)==Token.EOF ) matchedEOF = true;
							_errHandler.reportMatch(this);
							consume();
						}
						setState(514);
						((Expr_aritContext)_localctx).opDe = expr_arit(9);
						_localctx.p = expresion.NewOperacion(((Expr_aritContext)_localctx).opIz.p,(((Expr_aritContext)_localctx).op!=null?((Expr_aritContext)_localctx).op.getText():null),((Expr_aritContext)_localctx).opDe.p,false,((Expr_aritContext)_localctx).op.GetLine(),((Expr_aritContext)_localctx).op.GetColumn())
						}
						break;
					case 5:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.exp = _prevctx;
						_localctx.exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(517);
						if (!(precpred(_ctx, 15))) throw new FailedPredicateException(this, "precpred(_ctx, 15)");
						setState(518);
						match(P_AS);
						setState(519);
						((Expr_aritContext)_localctx).tipo_d = tipo_d();
						_localctx.p=expresion.NewCast(((Expr_aritContext)_localctx).exp.p,((Expr_aritContext)_localctx).tipo_d.t)
						}
						break;
					case 6:
						{
						_localctx = new Expr_aritContext(_parentctx, _parentState);
						_localctx.exp = _prevctx;
						_localctx.exp = _prevctx;
						pushNewRecursionContext(_localctx, _startState, RULE_expr_arit);
						setState(522);
						if (!(precpred(_ctx, 14))) throw new FailedPredicateException(this, "precpred(_ctx, 14)");
						setState(523);
						((Expr_aritContext)_localctx).PUNTO = match(PUNTO);
						setState(524);
						match(T_STRING);
						setState(525);
						((Expr_aritContext)_localctx).PARIZQ = match(PARIZQ);
						setState(526);
						match(PARDER);
						_localctx.p=expresion.NewToString(((Expr_aritContext)_localctx).exp.p,((Expr_aritContext)_localctx).PUNTO.GetLine(),((Expr_aritContext)_localctx).PUNTO.GetColumn())
						}
						break;
					}
					} 
				}
				setState(532);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,22,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class PrimitivoContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Token NUMBER;
		public Token STRING;
		public Token TRUE;
		public Token FALSE;
		public Token DECIMAL;
		public Token CARACTER;
		public Token id;
		public TerminalNode NUMBER() { return getToken(Chems.NUMBER, 0); }
		public TerminalNode STRING() { return getToken(Chems.STRING, 0); }
		public TerminalNode TRUE() { return getToken(Chems.TRUE, 0); }
		public TerminalNode FALSE() { return getToken(Chems.FALSE, 0); }
		public TerminalNode DECIMAL() { return getToken(Chems.DECIMAL, 0); }
		public TerminalNode CARACTER() { return getToken(Chems.CARACTER, 0); }
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public PrimitivoContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_primitivo; }
	}

	public final PrimitivoContext primitivo() throws RecognitionException {
		PrimitivoContext _localctx = new PrimitivoContext(_ctx, getState());
		enterRule(_localctx, 62, RULE_primitivo);
		try {
			setState(547);
			_errHandler.sync(this);
			switch (_input.LA(1)) {
			case NUMBER:
				enterOuterAlt(_localctx, 1);
				{
				setState(533);
				((PrimitivoContext)_localctx).NUMBER = match(NUMBER);

				            	num,err := strconv.Atoi((((PrimitivoContext)_localctx).NUMBER!=null?((PrimitivoContext)_localctx).NUMBER.getText():null))
				                if err!= nil{
				                    fmt.Println(err)
				                }
				            linea:=_localctx._NUMBER.GetLine()
				            col:=_localctx._NUMBER.GetColumn()
				            _localctx.p = expresion.NewPrimitivo (num,interfaces.INTEGER,col,linea)
				            
				       
				}
				break;
			case STRING:
				enterOuterAlt(_localctx, 2);
				{
				setState(535);
				((PrimitivoContext)_localctx).STRING = match(STRING);
				 
				      linea:=_localctx._STRING.GetLine()
				      col:=_localctx._STRING.GetColumn()
				      str:= (((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null)[1:len((((PrimitivoContext)_localctx).STRING!=null?((PrimitivoContext)_localctx).STRING.getText():null))-1]
				      _localctx.p = expresion.NewPrimitivo(str,interfaces.STRING,col,linea)
				}
				break;
			case TRUE:
				enterOuterAlt(_localctx, 3);
				{
				setState(537);
				((PrimitivoContext)_localctx).TRUE = match(TRUE);
				            
				        _localctx.p=expresion.NewPrimitivo(1,interfaces.BOOLEAN,((PrimitivoContext)_localctx).TRUE.GetColumn(),((PrimitivoContext)_localctx).TRUE.GetLine())      
				    
				}
				break;
			case FALSE:
				enterOuterAlt(_localctx, 4);
				{
				setState(539);
				((PrimitivoContext)_localctx).FALSE = match(FALSE);

				        _localctx.p=expresion.NewPrimitivo(0,interfaces.BOOLEAN,((PrimitivoContext)_localctx).FALSE.GetColumn(),((PrimitivoContext)_localctx).FALSE.GetLine())      
				    
				}
				break;
			case DECIMAL:
				enterOuterAlt(_localctx, 5);
				{
				setState(541);
				((PrimitivoContext)_localctx).DECIMAL = match(DECIMAL);

				            linea:=_localctx._DECIMAL.GetLine()
				            col:=_localctx._DECIMAL.GetColumn()
				              num,err:=strconv.ParseFloat((((PrimitivoContext)_localctx).DECIMAL!=null?((PrimitivoContext)_localctx).DECIMAL.getText():null),64)
				              if err!=nil{ 
				                fmt.Println(err)
				              }
				              _localctx.p=expresion.NewPrimitivo(num,interfaces.FLOAT,col,linea)
				}
				break;
			case CARACTER:
				enterOuterAlt(_localctx, 6);
				{
				setState(543);
				((PrimitivoContext)_localctx).CARACTER = match(CARACTER);

				      linea:=_localctx._CARACTER.GetLine()
				      col:=_localctx._CARACTER.GetColumn()
				      str:= (((PrimitivoContext)_localctx).CARACTER!=null?((PrimitivoContext)_localctx).CARACTER.getText():null)[1:len((((PrimitivoContext)_localctx).CARACTER!=null?((PrimitivoContext)_localctx).CARACTER.getText():null))-1]
				      _localctx.p = expresion.NewPrimitivo(str,interfaces.CHAR,col,linea)
				    
				}
				break;
			case ID:
				enterOuterAlt(_localctx, 7);
				{
				setState(545);
				((PrimitivoContext)_localctx).id = match(ID);

				      linea:=_localctx.id.GetLine()
				      col:=_localctx.id.GetColumn()
				      _localctx.p=expresion.NewCallVariable((((PrimitivoContext)_localctx).id!=null?((PrimitivoContext)_localctx).id.getText():null),linea,col)
				    
				}
				break;
			default:
				throw new NoViableAltException(this);
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListValuesContext extends ParserRuleContext {
		public *arrayList.List l;
		public ListValuesContext list;
		public ExpressionContext expression;
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode COMA() { return getToken(Chems.COMA, 0); }
		public ListValuesContext listValues() {
			return getRuleContext(ListValuesContext.class,0);
		}
		public ListValuesContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listValues; }
	}

	public final ListValuesContext listValues() throws RecognitionException {
		return listValues(0);
	}

	private ListValuesContext listValues(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListValuesContext _localctx = new ListValuesContext(_ctx, _parentState);
		ListValuesContext _prevctx = _localctx;
		int _startState = 64;
		enterRecursionRule(_localctx, 64, RULE_listValues, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(550);
			((ListValuesContext)_localctx).expression = expression();
			 
			                    _localctx.l = arrayList.New()
			                    _localctx.l.Add(((ListValuesContext)_localctx).expression.p)
			                
			}
			_ctx.stop = _input.LT(-1);
			setState(560);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListValuesContext(_parentctx, _parentState);
					_localctx.list = _prevctx;
					_localctx.list = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listValues);
					setState(553);
					if (!(precpred(_ctx, 2))) throw new FailedPredicateException(this, "precpred(_ctx, 2)");
					setState(554);
					match(COMA);
					setState(555);
					((ListValuesContext)_localctx).expression = expression();
					 
					                                                  ((ListValuesContext)_localctx).list.l.Add(((ListValuesContext)_localctx).expression.p)
					                                                  _localctx.l = ((ListValuesContext)_localctx).list.l
					                                              
					}
					} 
				}
				setState(562);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,24,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public static class ArrayAccContext extends ParserRuleContext {
		public interfaces.Expresion p;
		public Token id;
		public ListArrayContext list;
		public TerminalNode ID() { return getToken(Chems.ID, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public ArrayAccContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_arrayAcc; }
	}

	public final ArrayAccContext arrayAcc() throws RecognitionException {
		ArrayAccContext _localctx = new ArrayAccContext(_ctx, getState());
		enterRule(_localctx, 66, RULE_arrayAcc);
		try {
			enterOuterAlt(_localctx, 1);
			{
			setState(563);
			((ArrayAccContext)_localctx).id = match(ID);
			setState(564);
			((ArrayAccContext)_localctx).list = listArray(0);
			_localctx.p=expresion.NewArrayAccess((((ArrayAccContext)_localctx).id!=null?((ArrayAccContext)_localctx).id.getText():null),((ArrayAccContext)_localctx).list.l,((ArrayAccContext)_localctx).id.GetLine(),((ArrayAccContext)_localctx).id.GetColumn())
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			exitRule();
		}
		return _localctx;
	}

	public static class ListArrayContext extends ParserRuleContext {
		public *arrayList.List l;
		public ListArrayContext lista;
		public ExpressionContext expression;
		public TerminalNode CORIZQ() { return getToken(Chems.CORIZQ, 0); }
		public ExpressionContext expression() {
			return getRuleContext(ExpressionContext.class,0);
		}
		public TerminalNode CORDER() { return getToken(Chems.CORDER, 0); }
		public ListArrayContext listArray() {
			return getRuleContext(ListArrayContext.class,0);
		}
		public ListArrayContext(ParserRuleContext parent, int invokingState) {
			super(parent, invokingState);
		}
		@Override public int getRuleIndex() { return RULE_listArray; }
	}

	public final ListArrayContext listArray() throws RecognitionException {
		return listArray(0);
	}

	private ListArrayContext listArray(int _p) throws RecognitionException {
		ParserRuleContext _parentctx = _ctx;
		int _parentState = getState();
		ListArrayContext _localctx = new ListArrayContext(_ctx, _parentState);
		ListArrayContext _prevctx = _localctx;
		int _startState = 68;
		enterRecursionRule(_localctx, 68, RULE_listArray, _p);
		try {
			int _alt;
			enterOuterAlt(_localctx, 1);
			{
			{
			setState(568);
			match(CORIZQ);
			setState(569);
			((ListArrayContext)_localctx).expression = expression();
			setState(570);
			match(CORDER);

			    _localctx.l=arrayList.New()
			    _localctx.l.Add(((ListArrayContext)_localctx).expression.p)
			  
			}
			_ctx.stop = _input.LT(-1);
			setState(581);
			_errHandler.sync(this);
			_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
			while ( _alt!=2 && _alt!=org.antlr.v4.runtime.atn.ATN.INVALID_ALT_NUMBER ) {
				if ( _alt==1 ) {
					if ( _parseListeners!=null ) triggerExitRuleEvent();
					_prevctx = _localctx;
					{
					{
					_localctx = new ListArrayContext(_parentctx, _parentState);
					_localctx.lista = _prevctx;
					_localctx.lista = _prevctx;
					pushNewRecursionContext(_localctx, _startState, RULE_listArray);
					setState(573);
					if (!(precpred(_ctx, 1))) throw new FailedPredicateException(this, "precpred(_ctx, 1)");
					setState(574);
					match(CORIZQ);
					setState(575);
					((ListArrayContext)_localctx).expression = expression();
					setState(576);
					match(CORDER);

					              ((ListArrayContext)_localctx).lista.l.Add(((ListArrayContext)_localctx).expression.p)
					              _localctx.l=((ListArrayContext)_localctx).lista.l
					            
					}
					} 
				}
				setState(583);
				_errHandler.sync(this);
				_alt = getInterpreter().adaptivePredict(_input,25,_ctx);
			}
			}
		}
		catch (RecognitionException re) {
			_localctx.exception = re;
			_errHandler.reportError(this, re);
			_errHandler.recover(this, re);
		}
		finally {
			unrollRecursionContexts(_parentctx);
		}
		return _localctx;
	}

	public boolean sempred(RuleContext _localctx, int ruleIndex, int predIndex) {
		switch (ruleIndex) {
		case 15:
			return llaves_if_sempred((Llaves_ifContext)_localctx, predIndex);
		case 22:
			return casos_sempred((CasosContext)_localctx, predIndex);
		case 24:
			return conditions_sempred((ConditionsContext)_localctx, predIndex);
		case 30:
			return expr_arit_sempred((Expr_aritContext)_localctx, predIndex);
		case 32:
			return listValues_sempred((ListValuesContext)_localctx, predIndex);
		case 34:
			return listArray_sempred((ListArrayContext)_localctx, predIndex);
		}
		return true;
	}
	private boolean llaves_if_sempred(Llaves_ifContext _localctx, int predIndex) {
		switch (predIndex) {
		case 0:
			return precpred(_ctx, 2);
		case 1:
			return precpred(_ctx, 1);
		}
		return true;
	}
	private boolean casos_sempred(CasosContext _localctx, int predIndex) {
		switch (predIndex) {
		case 2:
			return precpred(_ctx, 4);
		case 3:
			return precpred(_ctx, 3);
		}
		return true;
	}
	private boolean conditions_sempred(ConditionsContext _localctx, int predIndex) {
		switch (predIndex) {
		case 4:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean expr_arit_sempred(Expr_aritContext _localctx, int predIndex) {
		switch (predIndex) {
		case 5:
			return precpred(_ctx, 11);
		case 6:
			return precpred(_ctx, 10);
		case 7:
			return precpred(_ctx, 9);
		case 8:
			return precpred(_ctx, 8);
		case 9:
			return precpred(_ctx, 15);
		case 10:
			return precpred(_ctx, 14);
		}
		return true;
	}
	private boolean listValues_sempred(ListValuesContext _localctx, int predIndex) {
		switch (predIndex) {
		case 11:
			return precpred(_ctx, 2);
		}
		return true;
	}
	private boolean listArray_sempred(ListArrayContext _localctx, int predIndex) {
		switch (predIndex) {
		case 12:
			return precpred(_ctx, 1);
		}
		return true;
	}

	public static final String _serializedATN =
		"\3\u608b\ua72a\u8133\ub9ed\u417c\u3be7\u7786\u5964\3X\u024b\4\2\t\2\4"+
		"\3\t\3\4\4\t\4\4\5\t\5\4\6\t\6\4\7\t\7\4\b\t\b\4\t\t\t\4\n\t\n\4\13\t"+
		"\13\4\f\t\f\4\r\t\r\4\16\t\16\4\17\t\17\4\20\t\20\4\21\t\21\4\22\t\22"+
		"\4\23\t\23\4\24\t\24\4\25\t\25\4\26\t\26\4\27\t\27\4\30\t\30\4\31\t\31"+
		"\4\32\t\32\4\33\t\33\4\34\t\34\4\35\t\35\4\36\t\36\4\37\t\37\4 \t \4!"+
		"\t!\4\"\t\"\4#\t#\4$\t$\3\2\3\2\3\2\3\3\7\3M\n\3\f\3\16\3P\13\3\3\3\3"+
		"\3\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4"+
		"\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3"+
		"\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\3\4\5\4\u0082\n\4\3\5\3\5\3"+
		"\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5"+
		"\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3\5\3"+
		"\5\3\5\5\5\u00ab\n\5\3\6\3\6\3\6\3\6\3\6\3\7\3\7\3\7\3\7\3\7\3\7\3\7\3"+
		"\7\3\b\3\b\3\b\5\b\u00bd\n\b\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3\t\3"+
		"\t\5\t\u00ca\n\t\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3\n\3"+
		"\n\3\n\3\n\5\n\u00dc\n\n\3\13\3\13\3\13\3\13\3\f\3\f\3\f\3\f\3\f\3\f\3"+
		"\f\3\f\3\f\3\f\3\f\3\f\3\f\3\f\5\f\u00f0\n\f\3\r\3\r\3\r\3\r\3\r\3\r\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\3"+
		"\16\3\16\3\16\3\16\3\16\3\16\3\16\3\16\5\16\u010e\n\16\3\17\3\17\3\17"+
		"\3\17\3\17\3\17\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\20\3\21\3\21\3\21"+
		"\3\21\3\21\3\21\3\21\5\21\u0125\n\21\3\21\3\21\3\21\3\21\3\21\3\21\3\21"+
		"\3\21\7\21\u012f\n\21\f\21\16\21\u0132\13\21\3\22\3\22\3\22\3\22\3\22"+
		"\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\3\22\5\22\u0144\n\22"+
		"\3\23\3\23\3\23\3\23\3\23\3\23\3\23\3\23\5\23\u014e\n\23\3\24\3\24\3\24"+
		"\3\24\3\25\3\25\3\25\3\25\3\25\3\25\3\25\3\26\3\26\3\26\3\26\5\26\u015f"+
		"\n\26\3\27\3\27\3\27\3\27\3\27\3\27\3\27\3\30\3\30\3\30\3\30\3\30\3\30"+
		"\3\30\5\30\u016f\n\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\3\30\7\30\u0179"+
		"\n\30\f\30\16\30\u017c\13\30\3\31\3\31\3\31\3\31\3\31\3\32\3\32\3\32\3"+
		"\32\3\32\3\32\3\32\3\32\3\32\7\32\u018c\n\32\f\32\16\32\u018f\13\32\3"+
		"\33\3\33\3\33\3\33\3\33\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3\34\3"+
		"\34\3\34\3\34\3\34\5\34\u01a3\n\34\3\35\3\35\3\35\3\35\3\35\3\35\3\35"+
		"\3\35\3\35\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\3\36\5\36\u01b7\n\36"+
		"\3\37\3\37\3\37\3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 "+
		"\3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 "+
		"\3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \5 \u01f2\n \3 \3 \3 \3 \3 \3 "+
		"\3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 \3 "+
		"\3 \3 \7 \u0213\n \f \16 \u0216\13 \3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3!\3"+
		"!\3!\3!\5!\u0226\n!\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\3\"\7\"\u0231\n\""+
		"\f\"\16\"\u0234\13\"\3#\3#\3#\3#\3$\3$\3$\3$\3$\3$\3$\3$\3$\3$\3$\3$\7"+
		"$\u0246\n$\f$\16$\u0249\13$\3$\2\b .\62>BF%\2\4\6\b\n\f\16\20\22\24\26"+
		"\30\32\34\36 \"$&(*,.\60\62\64\668:<>@BDF\2\6\4\2\24\25\30\30\3\2\26\27"+
		"\4\2\13\r\17\21\3\2\22\23\2\u0274\2H\3\2\2\2\4N\3\2\2\2\6\u0081\3\2\2"+
		"\2\b\u00aa\3\2\2\2\n\u00ac\3\2\2\2\f\u00b1\3\2\2\2\16\u00bc\3\2\2\2\20"+
		"\u00c9\3\2\2\2\22\u00db\3\2\2\2\24\u00dd\3\2\2\2\26\u00ef\3\2\2\2\30\u00f1"+
		"\3\2\2\2\32\u010d\3\2\2\2\34\u010f\3\2\2\2\36\u0115\3\2\2\2 \u0124\3\2"+
		"\2\2\"\u0143\3\2\2\2$\u014d\3\2\2\2&\u014f\3\2\2\2(\u0153\3\2\2\2*\u015e"+
		"\3\2\2\2,\u0160\3\2\2\2.\u016e\3\2\2\2\60\u017d\3\2\2\2\62\u0182\3\2\2"+
		"\2\64\u0190\3\2\2\2\66\u01a2\3\2\2\28\u01a4\3\2\2\2:\u01b6\3\2\2\2<\u01b8"+
		"\3\2\2\2>\u01f1\3\2\2\2@\u0225\3\2\2\2B\u0227\3\2\2\2D\u0235\3\2\2\2F"+
		"\u0239\3\2\2\2HI\5\4\3\2IJ\b\2\1\2J\3\3\2\2\2KM\5\6\4\2LK\3\2\2\2MP\3"+
		"\2\2\2NL\3\2\2\2NO\3\2\2\2OQ\3\2\2\2PN\3\2\2\2QR\b\3\1\2R\5\3\2\2\2ST"+
		"\7#\2\2TU\7\3\2\2UV\7$\2\2VW\7\32\2\2WX\5<\37\2XY\7\33\2\2YZ\7\4\2\2Z"+
		"[\b\4\1\2[\u0082\3\2\2\2\\]\5\f\7\2]^\7\4\2\2^_\b\4\1\2_\u0082\3\2\2\2"+
		"`a\5\n\6\2ab\7\4\2\2bc\b\4\1\2c\u0082\3\2\2\2de\7*\2\2ef\5<\37\2fg\7\34"+
		"\2\2gh\5\4\3\2hi\7\35\2\2ij\b\4\1\2j\u0082\3\2\2\2kl\5$\23\2lm\b\4\1\2"+
		"m\u0082\3\2\2\2no\5&\24\2op\b\4\1\2p\u0082\3\2\2\2qr\5\36\20\2rs\b\4\1"+
		"\2s\u0082\3\2\2\2tu\5\34\17\2uv\b\4\1\2v\u0082\3\2\2\2wx\5(\25\2xy\7\4"+
		"\2\2yz\b\4\1\2z\u0082\3\2\2\2{|\5,\27\2|}\b\4\1\2}\u0082\3\2\2\2~\177"+
		"\58\35\2\177\u0080\b\4\1\2\u0080\u0082\3\2\2\2\u0081S\3\2\2\2\u0081\\"+
		"\3\2\2\2\u0081`\3\2\2\2\u0081d\3\2\2\2\u0081k\3\2\2\2\u0081n\3\2\2\2\u0081"+
		"q\3\2\2\2\u0081t\3\2\2\2\u0081w\3\2\2\2\u0081{\3\2\2\2\u0081~\3\2\2\2"+
		"\u0082\7\3\2\2\2\u0083\u0084\7#\2\2\u0084\u0085\7\3\2\2\u0085\u0086\7"+
		"$\2\2\u0086\u0087\7\32\2\2\u0087\u0088\5<\37\2\u0088\u0089\7\33\2\2\u0089"+
		"\u008a\b\5\1\2\u008a\u00ab\3\2\2\2\u008b\u008c\5\f\7\2\u008c\u008d\b\5"+
		"\1\2\u008d\u00ab\3\2\2\2\u008e\u008f\5\n\6\2\u008f\u0090\b\5\1\2\u0090"+
		"\u00ab\3\2\2\2\u0091\u0092\7*\2\2\u0092\u0093\5<\37\2\u0093\u0094\7\34"+
		"\2\2\u0094\u0095\5\4\3\2\u0095\u0096\7\35\2\2\u0096\u0097\b\5\1\2\u0097"+
		"\u00ab\3\2\2\2\u0098\u0099\5$\23\2\u0099\u009a\b\5\1\2\u009a\u00ab\3\2"+
		"\2\2\u009b\u009c\5&\24\2\u009c\u009d\b\5\1\2\u009d\u00ab\3\2\2\2\u009e"+
		"\u009f\5\36\20\2\u009f\u00a0\b\5\1\2\u00a0\u00ab\3\2\2\2\u00a1\u00a2\5"+
		"\34\17\2\u00a2\u00a3\b\5\1\2\u00a3\u00ab\3\2\2\2\u00a4\u00a5\5(\25\2\u00a5"+
		"\u00a6\b\5\1\2\u00a6\u00ab\3\2\2\2\u00a7\u00a8\58\35\2\u00a8\u00a9\b\5"+
		"\1\2\u00a9\u00ab\3\2\2\2\u00aa\u0083\3\2\2\2\u00aa\u008b\3\2\2\2\u00aa"+
		"\u008e\3\2\2\2\u00aa\u0091\3\2\2\2\u00aa\u0098\3\2\2\2\u00aa\u009b\3\2"+
		"\2\2\u00aa\u009e\3\2\2\2\u00aa\u00a1\3\2\2\2\u00aa\u00a4\3\2\2\2\u00aa"+
		"\u00a7\3\2\2\2\u00ab\t\3\2\2\2\u00ac\u00ad\7V\2\2\u00ad\u00ae\7\t\2\2"+
		"\u00ae\u00af\5<\37\2\u00af\u00b0\b\6\1\2\u00b0\13\3\2\2\2\u00b1\u00b2"+
		"\7\63\2\2\u00b2\u00b3\5\16\b\2\u00b3\u00b4\7V\2\2\u00b4\u00b5\5\20\t\2"+
		"\u00b5\u00b6\7\t\2\2\u00b6\u00b7\5<\37\2\u00b7\u00b8\b\7\1\2\u00b8\r\3"+
		"\2\2\2\u00b9\u00ba\7:\2\2\u00ba\u00bd\b\b\1\2\u00bb\u00bd\b\b\1\2\u00bc"+
		"\u00b9\3\2\2\2\u00bc\u00bb\3\2\2\2\u00bd\17\3\2\2\2\u00be\u00bf\7\b\2"+
		"\2\u00bf\u00c0\5\30\r\2\u00c0\u00c1\b\t\1\2\u00c1\u00ca\3\2\2\2\u00c2"+
		"\u00c3\5\24\13\2\u00c3\u00c4\b\t\1\2\u00c4\u00ca\3\2\2\2\u00c5\u00c6\7"+
		"\b\2\2\u00c6\u00c7\5\22\n\2\u00c7\u00c8\b\t\1\2\u00c8\u00ca\3\2\2\2\u00c9"+
		"\u00be\3\2\2\2\u00c9\u00c2\3\2\2\2\u00c9\u00c5\3\2\2\2\u00ca\21\3\2\2"+
		"\2\u00cb\u00cc\7\64\2\2\u00cc\u00dc\b\n\1\2\u00cd\u00ce\7\65\2\2\u00ce"+
		"\u00dc\b\n\1\2\u00cf\u00d0\7\66\2\2\u00d0\u00dc\b\n\1\2\u00d1\u00d2\7"+
		"\67\2\2\u00d2\u00dc\b\n\1\2\u00d3\u00d4\78\2\2\u00d4\u00dc\b\n\1\2\u00d5"+
		"\u00d6\7\'\2\2\u00d6\u00dc\b\n\1\2\u00d7\u00d8\79\2\2\u00d8\u00dc\b\n"+
		"\1\2\u00d9\u00da\7V\2\2\u00da\u00dc\b\n\1\2\u00db\u00cb\3\2\2\2\u00db"+
		"\u00cd\3\2\2\2\u00db\u00cf\3\2\2\2\u00db\u00d1\3\2\2\2\u00db\u00d3\3\2"+
		"\2\2\u00db\u00d5\3\2\2\2\u00db\u00d7\3\2\2\2\u00db\u00d9\3\2\2\2\u00dc"+
		"\23\3\2\2\2\u00dd\u00de\7\b\2\2\u00de\u00df\5\26\f\2\u00df\u00e0\b\13"+
		"\1\2\u00e0\25\3\2\2\2\u00e1\u00e2\7\36\2\2\u00e2\u00e3\5\22\n\2\u00e3"+
		"\u00e4\7\4\2\2\u00e4\u00e5\5<\37\2\u00e5\u00e6\7\37\2\2\u00e6\u00e7\b"+
		"\f\1\2\u00e7\u00f0\3\2\2\2\u00e8\u00e9\7\36\2\2\u00e9\u00ea\5\26\f\2\u00ea"+
		"\u00eb\7\4\2\2\u00eb\u00ec\5<\37\2\u00ec\u00ed\7\37\2\2\u00ed\u00ee\b"+
		"\f\1\2\u00ee\u00f0\3\2\2\2\u00ef\u00e1\3\2\2\2\u00ef\u00e8\3\2\2\2\u00f0"+
		"\27\3\2\2\2\u00f1\u00f2\7O\2\2\u00f2\u00f3\7\13\2\2\u00f3\u00f4\5\32\16"+
		"\2\u00f4\u00f5\7\21\2\2\u00f5\u00f6\b\r\1\2\u00f6\31\3\2\2\2\u00f7\u00f8"+
		"\7\64\2\2\u00f8\u010e\b\16\1\2\u00f9\u00fa\7\65\2\2\u00fa\u010e\b\16\1"+
		"\2\u00fb\u00fc\7\66\2\2\u00fc\u010e\b\16\1\2\u00fd\u00fe\7\67\2\2\u00fe"+
		"\u010e\b\16\1\2\u00ff\u0100\78\2\2\u0100\u010e\b\16\1\2\u0101\u0102\7"+
		"\'\2\2\u0102\u010e\b\16\1\2\u0103\u0104\79\2\2\u0104\u010e\b\16\1\2\u0105"+
		"\u0106\7V\2\2\u0106\u010e\b\16\1\2\u0107\u0108\7O\2\2\u0108\u0109\7\13"+
		"\2\2\u0109\u010a\5\32\16\2\u010a\u010b\7\21\2\2\u010b\u010c\b\16\1\2\u010c"+
		"\u010e\3\2\2\2\u010d\u00f7\3\2\2\2\u010d\u00f9\3\2\2\2\u010d\u00fb\3\2"+
		"\2\2\u010d\u00fd\3\2\2\2\u010d\u00ff\3\2\2\2\u010d\u0101\3\2\2\2\u010d"+
		"\u0103\3\2\2\2\u010d\u0105\3\2\2\2\u010d\u0107\3\2\2\2\u010e\33\3\2\2"+
		"\2\u010f\u0110\7B\2\2\u0110\u0111\7\34\2\2\u0111\u0112\5\4\3\2\u0112\u0113"+
		"\7\35\2\2\u0113\u0114\b\17\1\2\u0114\35\3\2\2\2\u0115\u0116\7(\2\2\u0116"+
		"\u0117\5<\37\2\u0117\u0118\7\34\2\2\u0118\u0119\5 \21\2\u0119\u011a\7"+
		"\35\2\2\u011a\u011b\5\"\22\2\u011b\u011c\b\20\1\2\u011c\37\3\2\2\2\u011d"+
		"\u011e\b\21\1\2\u011e\u011f\5<\37\2\u011f\u0120\b\21\1\2\u0120\u0125\3"+
		"\2\2\2\u0121\u0122\5\6\4\2\u0122\u0123\b\21\1\2\u0123\u0125\3\2\2\2\u0124"+
		"\u011d\3\2\2\2\u0124\u0121\3\2\2\2\u0125\u0130\3\2\2\2\u0126\u0127\f\4"+
		"\2\2\u0127\u0128\5<\37\2\u0128\u0129\b\21\1\2\u0129\u012f\3\2\2\2\u012a"+
		"\u012b\f\3\2\2\u012b\u012c\5\6\4\2\u012c\u012d\b\21\1\2\u012d\u012f\3"+
		"\2\2\2\u012e\u0126\3\2\2\2\u012e\u012a\3\2\2\2\u012f\u0132\3\2\2\2\u0130"+
		"\u012e\3\2\2\2\u0130\u0131\3\2\2\2\u0131!\3\2\2\2\u0132\u0130\3\2\2\2"+
		"\u0133\u0134\7)\2\2\u0134\u0135\7\34\2\2\u0135\u0136\5 \21\2\u0136\u0137"+
		"\7\35\2\2\u0137\u0138\b\22\1\2\u0138\u0144\3\2\2\2\u0139\u013a\7)\2\2"+
		"\u013a\u013b\7(\2\2\u013b\u013c\5<\37\2\u013c\u013d\7\34\2\2\u013d\u013e"+
		"\5 \21\2\u013e\u013f\7\35\2\2\u013f\u0140\5\"\22\2\u0140\u0141\b\22\1"+
		"\2\u0141\u0144\3\2\2\2\u0142\u0144\b\22\1\2\u0143\u0133\3\2\2\2\u0143"+
		"\u0139\3\2\2\2\u0143\u0142\3\2\2\2\u0144#\3\2\2\2\u0145\u0146\7\60\2\2"+
		"\u0146\u0147\5<\37\2\u0147\u0148\7\4\2\2\u0148\u0149\b\23\1\2\u0149\u014e"+
		"\3\2\2\2\u014a\u014b\7\60\2\2\u014b\u014c\7\4\2\2\u014c\u014e\b\23\1\2"+
		"\u014d\u0145\3\2\2\2\u014d\u014a\3\2\2\2\u014e%\3\2\2\2\u014f\u0150\7"+
		"\61\2\2\u0150\u0151\7\4\2\2\u0151\u0152\b\24\1\2\u0152\'\3\2\2\2\u0153"+
		"\u0154\7%\2\2\u0154\u0155\7\32\2\2\u0155\u0156\5*\26\2\u0156\u0157\5B"+
		"\"\2\u0157\u0158\7\33\2\2\u0158\u0159\b\25\1\2\u0159)\3\2\2\2\u015a\u015b"+
		"\7U\2\2\u015b\u015c\7\5\2\2\u015c\u015f\b\26\1\2\u015d\u015f\b\26\1\2"+
		"\u015e\u015a\3\2\2\2\u015e\u015d\3\2\2\2\u015f+\3\2\2\2\u0160\u0161\7"+
		"A\2\2\u0161\u0162\5<\37\2\u0162\u0163\7\34\2\2\u0163\u0164\5.\30\2\u0164"+
		"\u0165\7\35\2\2\u0165\u0166\b\27\1\2\u0166-\3\2\2\2\u0167\u0168\b\30\1"+
		"\2\u0168\u0169\5\60\31\2\u0169\u016a\b\30\1\2\u016a\u016f\3\2\2\2\u016b"+
		"\u016c\5\64\33\2\u016c\u016d\b\30\1\2\u016d\u016f\3\2\2\2\u016e\u0167"+
		"\3\2\2\2\u016e\u016b\3\2\2\2\u016f\u017a\3\2\2\2\u0170\u0171\f\6\2\2\u0171"+
		"\u0172\5\60\31\2\u0172\u0173\b\30\1\2\u0173\u0179\3\2\2\2\u0174\u0175"+
		"\f\5\2\2\u0175\u0176\5\64\33\2\u0176\u0177\b\30\1\2\u0177\u0179\3\2\2"+
		"\2\u0178\u0170\3\2\2\2\u0178\u0174\3\2\2\2\u0179\u017c\3\2\2\2\u017a\u0178"+
		"\3\2\2\2\u017a\u017b\3\2\2\2\u017b/\3\2\2\2\u017c\u017a\3\2\2\2\u017d"+
		"\u017e\5\62\32\2\u017e\u017f\7\16\2\2\u017f\u0180\5\66\34\2\u0180\u0181"+
		"\b\31\1\2\u0181\61\3\2\2\2\u0182\u0183\b\32\1\2\u0183\u0184\5<\37\2\u0184"+
		"\u0185\b\32\1\2\u0185\u018d\3\2\2\2\u0186\u0187\f\4\2\2\u0187\u0188\7"+
		"!\2\2\u0188\u0189\5<\37\2\u0189\u018a\b\32\1\2\u018a\u018c\3\2\2\2\u018b"+
		"\u0186\3\2\2\2\u018c\u018f\3\2\2\2\u018d\u018b\3\2\2\2\u018d\u018e\3\2"+
		"\2\2\u018e\63\3\2\2\2\u018f\u018d\3\2\2\2\u0190\u0191\7M\2\2\u0191\u0192"+
		"\7\16\2\2\u0192\u0193\5\66\34\2\u0193\u0194\b\33\1\2\u0194\65\3\2\2\2"+
		"\u0195\u0196\5<\37\2\u0196\u0197\7\5\2\2\u0197\u0198\b\34\1\2\u0198\u01a3"+
		"\3\2\2\2\u0199\u019a\7\34\2\2\u019a\u019b\5\4\3\2\u019b\u019c\7\35\2\2"+
		"\u019c\u019d\b\34\1\2\u019d\u01a3\3\2\2\2\u019e\u019f\5\b\5\2\u019f\u01a0"+
		"\7\5\2\2\u01a0\u01a1\b\34\1\2\u01a1\u01a3\3\2\2\2\u01a2\u0195\3\2\2\2"+
		"\u01a2\u0199\3\2\2\2\u01a2\u019e\3\2\2\2\u01a3\67\3\2\2\2\u01a4\u01a5"+
		"\7+\2\2\u01a5\u01a6\7V\2\2\u01a6\u01a7\7,\2\2\u01a7\u01a8\5:\36\2\u01a8"+
		"\u01a9\7\34\2\2\u01a9\u01aa\5\4\3\2\u01aa\u01ab\7\35\2\2\u01ab\u01ac\b"+
		"\35\1\2\u01ac9\3\2\2\2\u01ad\u01ae\5<\37\2\u01ae\u01af\7\3\2\2\u01af\u01b0"+
		"\7\3\2\2\u01b0\u01b1\5<\37\2\u01b1\u01b2\b\36\1\2\u01b2\u01b7\3\2\2\2"+
		"\u01b3\u01b4\5<\37\2\u01b4\u01b5\b\36\1\2\u01b5\u01b7\3\2\2\2\u01b6\u01ad"+
		"\3\2\2\2\u01b6\u01b3\3\2\2\2\u01b7;\3\2\2\2\u01b8\u01b9\5> \2\u01b9\u01ba"+
		"\b\37\1\2\u01ba=\3\2\2\2\u01bb\u01bc\b \1\2\u01bc\u01bd\7\64\2\2\u01bd"+
		"\u01be\7\"\2\2\u01be\u01bf\7/\2\2\u01bf\u01c0\7\32\2\2\u01c0\u01c1\5>"+
		" \2\u01c1\u01c2\7\5\2\2\u01c2\u01c3\5> \2\u01c3\u01c4\7\33\2\2\u01c4\u01c5"+
		"\b \1\2\u01c5\u01f2\3\2\2\2\u01c6\u01c7\7\65\2\2\u01c7\u01c8\7\"\2\2\u01c8"+
		"\u01c9\7.\2\2\u01c9\u01ca\7\32\2\2\u01ca\u01cb\5> \2\u01cb\u01cc\7\5\2"+
		"\2\u01cc\u01cd\5> \2\u01cd\u01ce\7\33\2\2\u01ce\u01cf\b \1\2\u01cf\u01f2"+
		"\3\2\2\2\u01d0\u01d1\7\7\2\2\u01d1\u01d2\5> \17\u01d2\u01d3\b \1\2\u01d3"+
		"\u01f2\3\2\2\2\u01d4\u01d5\7\27\2\2\u01d5\u01d6\5> \16\u01d6\u01d7\b "+
		"\1\2\u01d7\u01f2\3\2\2\2\u01d8\u01d9\5@!\2\u01d9\u01da\b \1\2\u01da\u01f2"+
		"\3\2\2\2\u01db\u01dc\7\32\2\2\u01dc\u01dd\5<\37\2\u01dd\u01de\7\33\2\2"+
		"\u01de\u01df\b \1\2\u01df\u01f2\3\2\2\2\u01e0\u01e1\5\34\17\2\u01e1\u01e2"+
		"\b \1\2\u01e2\u01f2\3\2\2\2\u01e3\u01e4\5\36\20\2\u01e4\u01e5\b \1\2\u01e5"+
		"\u01f2\3\2\2\2\u01e6\u01e7\5,\27\2\u01e7\u01e8\b \1\2\u01e8\u01f2\3\2"+
		"\2\2\u01e9\u01ea\7\36\2\2\u01ea\u01eb\5B\"\2\u01eb\u01ec\7\37\2\2\u01ec"+
		"\u01ed\b \1\2\u01ed\u01f2\3\2\2\2\u01ee\u01ef\5D#\2\u01ef\u01f0\b \1\2"+
		"\u01f0\u01f2\3\2\2\2\u01f1\u01bb\3\2\2\2\u01f1\u01c6\3\2\2\2\u01f1\u01d0"+
		"\3\2\2\2\u01f1\u01d4\3\2\2\2\u01f1\u01d8\3\2\2\2\u01f1\u01db\3\2\2\2\u01f1"+
		"\u01e0\3\2\2\2\u01f1\u01e3\3\2\2\2\u01f1\u01e6\3\2\2\2\u01f1\u01e9\3\2"+
		"\2\2\u01f1\u01ee\3\2\2\2\u01f2\u0214\3\2\2\2\u01f3\u01f4\f\r\2\2\u01f4"+
		"\u01f5\t\2\2\2\u01f5\u01f6\5> \16\u01f6\u01f7\b \1\2\u01f7\u0213\3\2\2"+
		"\2\u01f8\u01f9\f\f\2\2\u01f9\u01fa\t\3\2\2\u01fa\u01fb\5> \r\u01fb\u01fc"+
		"\b \1\2\u01fc\u0213\3\2\2\2\u01fd\u01fe\f\13\2\2\u01fe\u01ff\t\4\2\2\u01ff"+
		"\u0200\5> \f\u0200\u0201\b \1\2\u0201\u0213\3\2\2\2\u0202\u0203\f\n\2"+
		"\2\u0203\u0204\t\5\2\2\u0204\u0205\5> \13\u0205\u0206\b \1\2\u0206\u0213"+
		"\3\2\2\2\u0207\u0208\f\21\2\2\u0208\u0209\7-\2\2\u0209\u020a\5\22\n\2"+
		"\u020a\u020b\b \1\2\u020b\u0213\3\2\2\2\u020c\u020d\f\20\2\2\u020d\u020e"+
		"\7\3\2\2\u020e\u020f\7D\2\2\u020f\u0210\7\32\2\2\u0210\u0211\7\33\2\2"+
		"\u0211\u0213\b \1\2\u0212\u01f3\3\2\2\2\u0212\u01f8\3\2\2\2\u0212\u01fd"+
		"\3\2\2\2\u0212\u0202\3\2\2\2\u0212\u0207\3\2\2\2\u0212\u020c\3\2\2\2\u0213"+
		"\u0216\3\2\2\2\u0214\u0212\3\2\2\2\u0214\u0215\3\2\2\2\u0215?\3\2\2\2"+
		"\u0216\u0214\3\2\2\2\u0217\u0218\7S\2\2\u0218\u0226\b!\1\2\u0219\u021a"+
		"\7U\2\2\u021a\u0226\b!\1\2\u021b\u021c\7?\2\2\u021c\u0226\b!\1\2\u021d"+
		"\u021e\7@\2\2\u021e\u0226\b!\1\2\u021f\u0220\7T\2\2\u0220\u0226\b!\1\2"+
		"\u0221\u0222\7W\2\2\u0222\u0226\b!\1\2\u0223\u0224\7V\2\2\u0224\u0226"+
		"\b!\1\2\u0225\u0217\3\2\2\2\u0225\u0219\3\2\2\2\u0225\u021b\3\2\2\2\u0225"+
		"\u021d\3\2\2\2\u0225\u021f\3\2\2\2\u0225\u0221\3\2\2\2\u0225\u0223\3\2"+
		"\2\2\u0226A\3\2\2\2\u0227\u0228\b\"\1\2\u0228\u0229\5<\37\2\u0229\u022a"+
		"\b\"\1\2\u022a\u0232\3\2\2\2\u022b\u022c\f\4\2\2\u022c\u022d\7\5\2\2\u022d"+
		"\u022e\5<\37\2\u022e\u022f\b\"\1\2\u022f\u0231\3\2\2\2\u0230\u022b\3\2"+
		"\2\2\u0231\u0234\3\2\2\2\u0232\u0230\3\2\2\2\u0232\u0233\3\2\2\2\u0233"+
		"C\3\2\2\2\u0234\u0232\3\2\2\2\u0235\u0236\7V\2\2\u0236\u0237\5F$\2\u0237"+
		"\u0238\b#\1\2\u0238E\3\2\2\2\u0239\u023a\b$\1\2\u023a\u023b\7\36\2\2\u023b"+
		"\u023c\5<\37\2\u023c\u023d\7\37\2\2\u023d\u023e\b$\1\2\u023e\u0247\3\2"+
		"\2\2\u023f\u0240\f\3\2\2\u0240\u0241\7\36\2\2\u0241\u0242\5<\37\2\u0242"+
		"\u0243\7\37\2\2\u0243\u0244\b$\1\2\u0244\u0246\3\2\2\2\u0245\u023f\3\2"+
		"\2\2\u0246\u0249\3\2\2\2\u0247\u0245\3\2\2\2\u0247\u0248\3\2\2\2\u0248"+
		"G\3\2\2\2\u0249\u0247\3\2\2\2\34N\u0081\u00aa\u00bc\u00c9\u00db\u00ef"+
		"\u010d\u0124\u012e\u0130\u0143\u014d\u015e\u016e\u0178\u017a\u018d\u01a2"+
		"\u01b6\u01f1\u0212\u0214\u0225\u0232\u0247";
	public static final ATN _ATN =
		new ATNDeserializer().deserialize(_serializedATN.toCharArray());
	static {
		_decisionToDFA = new DFA[_ATN.getNumberOfDecisions()];
		for (int i = 0; i < _ATN.getNumberOfDecisions(); i++) {
			_decisionToDFA[i] = new DFA(_ATN.getDecisionState(i), i);
		}
	}
}