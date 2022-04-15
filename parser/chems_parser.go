// Code generated from Chems.g4 by ANTLR 4.9.3. DO NOT EDIT.

package parser // Chems

import (
	"fmt"
	"reflect"
	"strconv"

	"github.com/antlr/antlr4/runtime/Go/antlr"
)

import "OLC2/interfaces"
import "OLC2/expresion"
import "OLC2/instruction"
import arrayList "github.com/colegno/arraylist"

// Suppress unused import errors
var _ = fmt.Printf
var _ = reflect.Copy
var _ = strconv.Itoa

var parserATN = []uint16{
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 88, 287,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 3, 2, 3, 2, 3, 2, 3, 3, 7, 3, 41, 10, 3, 12, 3, 14, 3, 44, 11, 3, 3,
	3, 3, 3, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 5, 4, 78, 10, 4, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3,
	7, 3, 7, 3, 7, 5, 7, 96, 10, 7, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3, 8, 3,
	8, 3, 8, 3, 8, 3, 8, 3, 8, 5, 8, 109, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9, 3,
	9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5,
	9, 127, 10, 9, 3, 10, 3, 10, 3, 10, 3, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 3, 11, 5,
	11, 147, 10, 11, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3,
	13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 3, 13, 5, 13,
	177, 10, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 5,
	14, 187, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 233, 10, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 266, 10, 17, 12,
	17, 14, 17, 269, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18,
	3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 285, 10, 18, 3,
	18, 2, 3, 32, 19, 2, 4, 6, 8, 10, 12, 14, 16, 18, 20, 22, 24, 26, 28, 30,
	32, 34, 2, 6, 4, 2, 20, 21, 24, 24, 3, 2, 22, 23, 4, 2, 11, 13, 15, 17,
	3, 2, 18, 19, 2, 312, 2, 36, 3, 2, 2, 2, 4, 42, 3, 2, 2, 2, 6, 77, 3, 2,
	2, 2, 8, 79, 3, 2, 2, 2, 10, 84, 3, 2, 2, 2, 12, 95, 3, 2, 2, 2, 14, 108,
	3, 2, 2, 2, 16, 126, 3, 2, 2, 2, 18, 128, 3, 2, 2, 2, 20, 146, 3, 2, 2,
	2, 22, 148, 3, 2, 2, 2, 24, 176, 3, 2, 2, 2, 26, 186, 3, 2, 2, 2, 28, 188,
	3, 2, 2, 2, 30, 192, 3, 2, 2, 2, 32, 232, 3, 2, 2, 2, 34, 284, 3, 2, 2,
	2, 36, 37, 5, 4, 3, 2, 37, 38, 8, 2, 1, 2, 38, 3, 3, 2, 2, 2, 39, 41, 5,
	6, 4, 2, 40, 39, 3, 2, 2, 2, 41, 44, 3, 2, 2, 2, 42, 40, 3, 2, 2, 2, 42,
	43, 3, 2, 2, 2, 43, 45, 3, 2, 2, 2, 44, 42, 3, 2, 2, 2, 45, 46, 8, 3, 1,
	2, 46, 5, 3, 2, 2, 2, 47, 48, 7, 35, 2, 2, 48, 49, 7, 3, 2, 2, 49, 50,
	7, 36, 2, 2, 50, 51, 7, 26, 2, 2, 51, 52, 5, 30, 16, 2, 52, 53, 7, 27,
	2, 2, 53, 54, 7, 4, 2, 2, 54, 55, 8, 4, 1, 2, 55, 78, 3, 2, 2, 2, 56, 57,
	5, 10, 6, 2, 57, 58, 7, 4, 2, 2, 58, 59, 8, 4, 1, 2, 59, 78, 3, 2, 2, 2,
	60, 61, 5, 8, 5, 2, 61, 62, 7, 4, 2, 2, 62, 63, 8, 4, 1, 2, 63, 78, 3,
	2, 2, 2, 64, 65, 7, 42, 2, 2, 65, 66, 5, 30, 16, 2, 66, 67, 7, 28, 2, 2,
	67, 68, 5, 4, 3, 2, 68, 69, 7, 29, 2, 2, 69, 70, 8, 4, 1, 2, 70, 78, 3,
	2, 2, 2, 71, 72, 5, 26, 14, 2, 72, 73, 8, 4, 1, 2, 73, 78, 3, 2, 2, 2,
	74, 75, 5, 28, 15, 2, 75, 76, 8, 4, 1, 2, 76, 78, 3, 2, 2, 2, 77, 47, 3,
	2, 2, 2, 77, 56, 3, 2, 2, 2, 77, 60, 3, 2, 2, 2, 77, 64, 3, 2, 2, 2, 77,
	71, 3, 2, 2, 2, 77, 74, 3, 2, 2, 2, 78, 7, 3, 2, 2, 2, 79, 80, 7, 86, 2,
	2, 80, 81, 7, 9, 2, 2, 81, 82, 5, 30, 16, 2, 82, 83, 8, 5, 1, 2, 83, 9,
	3, 2, 2, 2, 84, 85, 7, 51, 2, 2, 85, 86, 5, 12, 7, 2, 86, 87, 7, 86, 2,
	2, 87, 88, 5, 14, 8, 2, 88, 89, 7, 9, 2, 2, 89, 90, 5, 30, 16, 2, 90, 91,
	8, 6, 1, 2, 91, 11, 3, 2, 2, 2, 92, 93, 7, 58, 2, 2, 93, 96, 8, 7, 1, 2,
	94, 96, 8, 7, 1, 2, 95, 92, 3, 2, 2, 2, 95, 94, 3, 2, 2, 2, 96, 13, 3,
	2, 2, 2, 97, 98, 7, 8, 2, 2, 98, 99, 5, 22, 12, 2, 99, 100, 8, 8, 1, 2,
	100, 109, 3, 2, 2, 2, 101, 102, 5, 18, 10, 2, 102, 103, 8, 8, 1, 2, 103,
	109, 3, 2, 2, 2, 104, 105, 7, 8, 2, 2, 105, 106, 5, 16, 9, 2, 106, 107,
	8, 8, 1, 2, 107, 109, 3, 2, 2, 2, 108, 97, 3, 2, 2, 2, 108, 101, 3, 2,
	2, 2, 108, 104, 3, 2, 2, 2, 109, 15, 3, 2, 2, 2, 110, 111, 7, 52, 2, 2,
	111, 127, 8, 9, 1, 2, 112, 113, 7, 53, 2, 2, 113, 127, 8, 9, 1, 2, 114,
	115, 7, 54, 2, 2, 115, 127, 8, 9, 1, 2, 116, 117, 7, 55, 2, 2, 117, 127,
	8, 9, 1, 2, 118, 119, 7, 56, 2, 2, 119, 127, 8, 9, 1, 2, 120, 121, 7, 39,
	2, 2, 121, 127, 8, 9, 1, 2, 122, 123, 7, 57, 2, 2, 123, 127, 8, 9, 1, 2,
	124, 125, 7, 86, 2, 2, 125, 127, 8, 9, 1, 2, 126, 110, 3, 2, 2, 2, 126,
	112, 3, 2, 2, 2, 126, 114, 3, 2, 2, 2, 126, 116, 3, 2, 2, 2, 126, 118,
	3, 2, 2, 2, 126, 120, 3, 2, 2, 2, 126, 122, 3, 2, 2, 2, 126, 124, 3, 2,
	2, 2, 127, 17, 3, 2, 2, 2, 128, 129, 7, 8, 2, 2, 129, 130, 5, 20, 11, 2,
	130, 131, 8, 10, 1, 2, 131, 19, 3, 2, 2, 2, 132, 133, 7, 30, 2, 2, 133,
	134, 5, 16, 9, 2, 134, 135, 7, 4, 2, 2, 135, 136, 5, 30, 16, 2, 136, 137,
	7, 31, 2, 2, 137, 138, 8, 11, 1, 2, 138, 147, 3, 2, 2, 2, 139, 140, 7,
	30, 2, 2, 140, 141, 5, 20, 11, 2, 141, 142, 7, 4, 2, 2, 142, 143, 5, 30,
	16, 2, 143, 144, 7, 31, 2, 2, 144, 145, 8, 11, 1, 2, 145, 147, 3, 2, 2,
	2, 146, 132, 3, 2, 2, 2, 146, 139, 3, 2, 2, 2, 147, 21, 3, 2, 2, 2, 148,
	149, 7, 79, 2, 2, 149, 150, 7, 11, 2, 2, 150, 151, 5, 24, 13, 2, 151, 152,
	7, 17, 2, 2, 152, 153, 8, 12, 1, 2, 153, 23, 3, 2, 2, 2, 154, 155, 7, 52,
	2, 2, 155, 177, 8, 13, 1, 2, 156, 157, 7, 53, 2, 2, 157, 177, 8, 13, 1,
	2, 158, 159, 7, 54, 2, 2, 159, 177, 8, 13, 1, 2, 160, 161, 7, 55, 2, 2,
	161, 177, 8, 13, 1, 2, 162, 163, 7, 56, 2, 2, 163, 177, 8, 13, 1, 2, 164,
	165, 7, 39, 2, 2, 165, 177, 8, 13, 1, 2, 166, 167, 7, 57, 2, 2, 167, 177,
	8, 13, 1, 2, 168, 169, 7, 86, 2, 2, 169, 177, 8, 13, 1, 2, 170, 171, 7,
	79, 2, 2, 171, 172, 7, 11, 2, 2, 172, 173, 5, 24, 13, 2, 173, 174, 7, 17,
	2, 2, 174, 175, 8, 13, 1, 2, 175, 177, 3, 2, 2, 2, 176, 154, 3, 2, 2, 2,
	176, 156, 3, 2, 2, 2, 176, 158, 3, 2, 2, 2, 176, 160, 3, 2, 2, 2, 176,
	162, 3, 2, 2, 2, 176, 164, 3, 2, 2, 2, 176, 166, 3, 2, 2, 2, 176, 168,
	3, 2, 2, 2, 176, 170, 3, 2, 2, 2, 177, 25, 3, 2, 2, 2, 178, 179, 7, 48,
	2, 2, 179, 180, 5, 30, 16, 2, 180, 181, 7, 4, 2, 2, 181, 182, 8, 14, 1,
	2, 182, 187, 3, 2, 2, 2, 183, 184, 7, 48, 2, 2, 184, 185, 7, 4, 2, 2, 185,
	187, 8, 14, 1, 2, 186, 178, 3, 2, 2, 2, 186, 183, 3, 2, 2, 2, 187, 27,
	3, 2, 2, 2, 188, 189, 7, 49, 2, 2, 189, 190, 7, 4, 2, 2, 190, 191, 8, 15,
	1, 2, 191, 29, 3, 2, 2, 2, 192, 193, 5, 32, 17, 2, 193, 194, 8, 16, 1,
	2, 194, 31, 3, 2, 2, 2, 195, 196, 8, 17, 1, 2, 196, 197, 7, 52, 2, 2, 197,
	198, 7, 34, 2, 2, 198, 199, 7, 47, 2, 2, 199, 200, 7, 26, 2, 2, 200, 201,
	5, 32, 17, 2, 201, 202, 7, 5, 2, 2, 202, 203, 5, 32, 17, 2, 203, 204, 7,
	27, 2, 2, 204, 205, 8, 17, 1, 2, 205, 233, 3, 2, 2, 2, 206, 207, 7, 53,
	2, 2, 207, 208, 7, 34, 2, 2, 208, 209, 7, 46, 2, 2, 209, 210, 7, 26, 2,
	2, 210, 211, 5, 32, 17, 2, 211, 212, 7, 5, 2, 2, 212, 213, 5, 32, 17, 2,
	213, 214, 7, 27, 2, 2, 214, 215, 8, 17, 1, 2, 215, 233, 3, 2, 2, 2, 216,
	217, 7, 7, 2, 2, 217, 218, 5, 32, 17, 10, 218, 219, 8, 17, 1, 2, 219, 233,
	3, 2, 2, 2, 220, 221, 7, 23, 2, 2, 221, 222, 5, 32, 17, 9, 222, 223, 8,
	17, 1, 2, 223, 233, 3, 2, 2, 2, 224, 225, 5, 34, 18, 2, 225, 226, 8, 17,
	1, 2, 226, 233, 3, 2, 2, 2, 227, 228, 7, 26, 2, 2, 228, 229, 5, 30, 16,
	2, 229, 230, 7, 27, 2, 2, 230, 231, 8, 17, 1, 2, 231, 233, 3, 2, 2, 2,
	232, 195, 3, 2, 2, 2, 232, 206, 3, 2, 2, 2, 232, 216, 3, 2, 2, 2, 232,
	220, 3, 2, 2, 2, 232, 224, 3, 2, 2, 2, 232, 227, 3, 2, 2, 2, 233, 267,
	3, 2, 2, 2, 234, 235, 12, 8, 2, 2, 235, 236, 9, 2, 2, 2, 236, 237, 5, 32,
	17, 9, 237, 238, 8, 17, 1, 2, 238, 266, 3, 2, 2, 2, 239, 240, 12, 7, 2,
	2, 240, 241, 9, 3, 2, 2, 241, 242, 5, 32, 17, 8, 242, 243, 8, 17, 1, 2,
	243, 266, 3, 2, 2, 2, 244, 245, 12, 6, 2, 2, 245, 246, 9, 4, 2, 2, 246,
	247, 5, 32, 17, 7, 247, 248, 8, 17, 1, 2, 248, 266, 3, 2, 2, 2, 249, 250,
	12, 5, 2, 2, 250, 251, 9, 5, 2, 2, 251, 252, 5, 32, 17, 6, 252, 253, 8,
	17, 1, 2, 253, 266, 3, 2, 2, 2, 254, 255, 12, 12, 2, 2, 255, 256, 7, 45,
	2, 2, 256, 257, 5, 16, 9, 2, 257, 258, 8, 17, 1, 2, 258, 266, 3, 2, 2,
	2, 259, 260, 12, 11, 2, 2, 260, 261, 7, 3, 2, 2, 261, 262, 7, 68, 2, 2,
	262, 263, 7, 26, 2, 2, 263, 264, 7, 27, 2, 2, 264, 266, 8, 17, 1, 2, 265,
	234, 3, 2, 2, 2, 265, 239, 3, 2, 2, 2, 265, 244, 3, 2, 2, 2, 265, 249,
	3, 2, 2, 2, 265, 254, 3, 2, 2, 2, 265, 259, 3, 2, 2, 2, 266, 269, 3, 2,
	2, 2, 267, 265, 3, 2, 2, 2, 267, 268, 3, 2, 2, 2, 268, 33, 3, 2, 2, 2,
	269, 267, 3, 2, 2, 2, 270, 271, 7, 83, 2, 2, 271, 285, 8, 18, 1, 2, 272,
	273, 7, 85, 2, 2, 273, 285, 8, 18, 1, 2, 274, 275, 7, 63, 2, 2, 275, 285,
	8, 18, 1, 2, 276, 277, 7, 64, 2, 2, 277, 285, 8, 18, 1, 2, 278, 279, 7,
	84, 2, 2, 279, 285, 8, 18, 1, 2, 280, 281, 7, 87, 2, 2, 281, 285, 8, 18,
	1, 2, 282, 283, 7, 86, 2, 2, 283, 285, 8, 18, 1, 2, 284, 270, 3, 2, 2,
	2, 284, 272, 3, 2, 2, 2, 284, 274, 3, 2, 2, 2, 284, 276, 3, 2, 2, 2, 284,
	278, 3, 2, 2, 2, 284, 280, 3, 2, 2, 2, 284, 282, 3, 2, 2, 2, 285, 35, 3,
	2, 2, 2, 14, 42, 77, 95, 108, 126, 146, 176, 186, 232, 265, 267, 284,
}
var literalNames = []string{
	"", "'.'", "';'", "','", "'\"'", "'!'", "':'", "'='", "'&'", "'<'", "'>='",
	"'<='", "'=>'", "'=='", "'!='", "'>'", "'||'", "'&&'", "'*'", "'/'", "'+'",
	"'-'", "'%'", "'^'", "'('", "')'", "'{'", "'}'", "'['", "']'", "'?'", "'|'",
	"'::'", "'console'", "'log'", "'println!'", "'number'", "'String'", "'if'",
	"'else'", "'while'", "'for'", "'in'", "'as'", "'powf'", "'pow'", "'break'",
	"'continue'", "'return'", "'let'", "'i64'", "'f64'", "'bool'", "'char'",
	"'&str'", "'usize'", "'mut'", "'struct'", "'fn'", "'mod'", "'pub'", "'true'",
	"'false'", "'match'", "'loop'", "'push'", "'to_string'", "'abs'", "'sqrt'",
	"'clone'", "'insert'", "'remove'", "'contains'", "'len'", "'capacity'",
	"'_'", "'vec!'", "'Vec'", "'with_capacity'", "'new'",
}
var symbolicNames = []string{
	"", "PUNTO", "PTCOMA", "COMA", "COMILLAS", "DIFERENTE", "DOSPUNTOS", "IGUAL",
	"AMPER", "MENOR", "MAYORIGUAL", "MENORIGUAL", "OPMATCH", "D_IGUAL", "NOT_E",
	"MAYOR", "OR", "AND", "MUL", "DIV", "ADD", "SUB", "MOD", "POT", "PARIZQ",
	"PARDER", "LLAVEIZQ", "LLAVEDER", "CORIZQ", "CORDER", "INTDER", "OR_COND",
	"DDPUNTO", "CONSOLE", "LOG", "PRINTLN", "P_NUMBER", "P_STRING", "P_IF",
	"P_ELSE", "P_WHILE", "P_FOR", "P_IN", "P_AS", "POWF", "POW", "BREAK", "CONTINUE",
	"RETURN", "LET", "INT", "FLOAT", "BOOL", "CHAR", "STR", "USIZE", "MUT",
	"P_STRUCT", "FUNCTION", "MODULE", "PUB", "TRUE", "FALSE", "MATCH", "LOOP",
	"PUSH", "T_STRING", "F_ABS", "F_SQRT", "F_CLONE", "INSERT", "REMOVE", "CONTAINS",
	"LEN", "CAPACITY", "DEF", "VEC", "VECT", "CAP", "NEW", "COMENTARIO", "NUMBER",
	"DECIMAL", "STRING", "ID", "CARACTER", "WHITESPACE",
}

var ruleNames = []string{
	"start", "instrucciones", "instruccion", "asignacion_var", "declaracion_var",
	"mutable", "types", "tipo_d", "asignar_Array", "dimensiones", "tipo_vector",
	"vectores", "breaks", "continues", "expression", "expr_arit", "primitivo",
}

type Chems struct {
	*antlr.BaseParser
}

// NewChems produces a new parser instance for the optional input antlr.TokenStream.
//
// The *Chems instance produced may be reused by calling the SetInputStream method.
// The initial parser configuration is expensive to construct, and the object is not thread-safe;
// however, if used within a Golang sync.Pool, the construction cost amortizes well and the
// objects can be used in a thread-safe manner.
func NewChems(input antlr.TokenStream) *Chems {
	this := new(Chems)
	deserializer := antlr.NewATNDeserializer(nil)
	deserializedATN := deserializer.DeserializeFromUInt16(parserATN)
	decisionToDFA := make([]*antlr.DFA, len(deserializedATN.DecisionToState))
	for index, ds := range deserializedATN.DecisionToState {
		decisionToDFA[index] = antlr.NewDFA(ds, index)
	}
	this.BaseParser = antlr.NewBaseParser(input)

	this.Interpreter = antlr.NewParserATNSimulator(this, deserializedATN, decisionToDFA, antlr.NewPredictionContextCache())
	this.RuleNames = ruleNames
	this.LiteralNames = literalNames
	this.SymbolicNames = symbolicNames
	this.GrammarFileName = "Chems.g4"

	return this
}

// Chems tokens.
const (
	ChemsEOF        = antlr.TokenEOF
	ChemsPUNTO      = 1
	ChemsPTCOMA     = 2
	ChemsCOMA       = 3
	ChemsCOMILLAS   = 4
	ChemsDIFERENTE  = 5
	ChemsDOSPUNTOS  = 6
	ChemsIGUAL      = 7
	ChemsAMPER      = 8
	ChemsMENOR      = 9
	ChemsMAYORIGUAL = 10
	ChemsMENORIGUAL = 11
	ChemsOPMATCH    = 12
	ChemsD_IGUAL    = 13
	ChemsNOT_E      = 14
	ChemsMAYOR      = 15
	ChemsOR         = 16
	ChemsAND        = 17
	ChemsMUL        = 18
	ChemsDIV        = 19
	ChemsADD        = 20
	ChemsSUB        = 21
	ChemsMOD        = 22
	ChemsPOT        = 23
	ChemsPARIZQ     = 24
	ChemsPARDER     = 25
	ChemsLLAVEIZQ   = 26
	ChemsLLAVEDER   = 27
	ChemsCORIZQ     = 28
	ChemsCORDER     = 29
	ChemsINTDER     = 30
	ChemsOR_COND    = 31
	ChemsDDPUNTO    = 32
	ChemsCONSOLE    = 33
	ChemsLOG        = 34
	ChemsPRINTLN    = 35
	ChemsP_NUMBER   = 36
	ChemsP_STRING   = 37
	ChemsP_IF       = 38
	ChemsP_ELSE     = 39
	ChemsP_WHILE    = 40
	ChemsP_FOR      = 41
	ChemsP_IN       = 42
	ChemsP_AS       = 43
	ChemsPOWF       = 44
	ChemsPOW        = 45
	ChemsBREAK      = 46
	ChemsCONTINUE   = 47
	ChemsRETURN     = 48
	ChemsLET        = 49
	ChemsINT        = 50
	ChemsFLOAT      = 51
	ChemsBOOL       = 52
	ChemsCHAR       = 53
	ChemsSTR        = 54
	ChemsUSIZE      = 55
	ChemsMUT        = 56
	ChemsP_STRUCT   = 57
	ChemsFUNCTION   = 58
	ChemsMODULE     = 59
	ChemsPUB        = 60
	ChemsTRUE       = 61
	ChemsFALSE      = 62
	ChemsMATCH      = 63
	ChemsLOOP       = 64
	ChemsPUSH       = 65
	ChemsT_STRING   = 66
	ChemsF_ABS      = 67
	ChemsF_SQRT     = 68
	ChemsF_CLONE    = 69
	ChemsINSERT     = 70
	ChemsREMOVE     = 71
	ChemsCONTAINS   = 72
	ChemsLEN        = 73
	ChemsCAPACITY   = 74
	ChemsDEF        = 75
	ChemsVEC        = 76
	ChemsVECT       = 77
	ChemsCAP        = 78
	ChemsNEW        = 79
	ChemsCOMENTARIO = 80
	ChemsNUMBER     = 81
	ChemsDECIMAL    = 82
	ChemsSTRING     = 83
	ChemsID         = 84
	ChemsCARACTER   = 85
	ChemsWHITESPACE = 86
)

// Chems rules.
const (
	ChemsRULE_start           = 0
	ChemsRULE_instrucciones   = 1
	ChemsRULE_instruccion     = 2
	ChemsRULE_asignacion_var  = 3
	ChemsRULE_declaracion_var = 4
	ChemsRULE_mutable         = 5
	ChemsRULE_types           = 6
	ChemsRULE_tipo_d          = 7
	ChemsRULE_asignar_Array   = 8
	ChemsRULE_dimensiones     = 9
	ChemsRULE_tipo_vector     = 10
	ChemsRULE_vectores        = 11
	ChemsRULE_breaks          = 12
	ChemsRULE_continues       = 13
	ChemsRULE_expression      = 14
	ChemsRULE_expr_arit       = 15
	ChemsRULE_primitivo       = 16
)

// IStartContext is an interface to support dynamic dispatch.
type IStartContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetLista returns the lista attribute.
	GetLista() *arrayList.List

	// SetLista sets the lista attribute.
	SetLista(*arrayList.List)

	// IsStartContext differentiates from other interfaces.
	IsStartContext()
}

type StartContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	lista          *arrayList.List
	_instrucciones IInstruccionesContext
}

func NewEmptyStartContext() *StartContext {
	var p = new(StartContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_start
	return p
}

func (*StartContext) IsStartContext() {}

func NewStartContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *StartContext {
	var p = new(StartContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_start

	return p
}

func (s *StartContext) GetParser() antlr.Parser { return s.parser }

func (s *StartContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *StartContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *StartContext) GetLista() *arrayList.List { return s.lista }

func (s *StartContext) SetLista(v *arrayList.List) { s.lista = v }

func (s *StartContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *StartContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *StartContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *StartContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterStart(s)
	}
}

func (s *StartContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitStart(s)
	}
}

func (p *Chems) Start() (localctx IStartContext) {
	this := p
	_ = this

	localctx = NewStartContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 0, ChemsRULE_start)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(34)

		var _x = p.Instrucciones()

		localctx.(*StartContext)._instrucciones = _x
	}
	localctx.(*StartContext).lista = localctx.(*StartContext).Get_instrucciones().GetL()

	return localctx
}

// IInstruccionesContext is an interface to support dynamic dispatch.
type IInstruccionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetE returns the e rule context list.
	GetE() []IInstruccionContext

	// SetE sets the e rule context list.
	SetE([]IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsInstruccionesContext differentiates from other interfaces.
	IsInstruccionesContext()
}

type InstruccionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	_instruccion IInstruccionContext
	e            []IInstruccionContext
}

func NewEmptyInstruccionesContext() *InstruccionesContext {
	var p = new(InstruccionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instrucciones
	return p
}

func (*InstruccionesContext) IsInstruccionesContext() {}

func NewInstruccionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionesContext {
	var p = new(InstruccionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instrucciones

	return p
}

func (s *InstruccionesContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionesContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *InstruccionesContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *InstruccionesContext) GetE() []IInstruccionContext { return s.e }

func (s *InstruccionesContext) SetE(v []IInstruccionContext) { s.e = v }

func (s *InstruccionesContext) GetL() *arrayList.List { return s.l }

func (s *InstruccionesContext) SetL(v *arrayList.List) { s.l = v }

func (s *InstruccionesContext) AllInstruccion() []IInstruccionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IInstruccionContext)(nil)).Elem())
	var tst = make([]IInstruccionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IInstruccionContext)
		}
	}

	return tst
}

func (s *InstruccionesContext) Instruccion(i int) IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *InstruccionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstrucciones(s)
	}
}

func (s *InstruccionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstrucciones(s)
	}
}

func (p *Chems) Instrucciones() (localctx IInstruccionesContext) {
	this := p
	_ = this

	localctx = NewInstruccionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 2, ChemsRULE_instrucciones)

	localctx.(*InstruccionesContext).l = arrayList.New()

	var _la int

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	p.SetState(40)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ChemsCONSOLE-33))|(1<<(ChemsP_WHILE-33))|(1<<(ChemsBREAK-33))|(1<<(ChemsCONTINUE-33))|(1<<(ChemsLET-33)))) != 0) || _la == ChemsID {
		{
			p.SetState(37)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(42)
		p.GetErrorHandler().Sync(p)
		_la = p.GetTokenStream().LA(1)
	}

	listInt := localctx.(*InstruccionesContext).GetE()
	for _, e := range listInt {
		localctx.(*InstruccionesContext).l.Add(e.GetInstr())
	}

	return localctx
}

// IInstruccionContext is an interface to support dynamic dispatch.
type IInstruccionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LLAVEIZQ returns the _LLAVEIZQ token.
	Get_LLAVEIZQ() antlr.Token

	// Get_LLAVEDER returns the _LLAVEDER token.
	Get_LLAVEDER() antlr.Token

	// Set_LLAVEIZQ sets the _LLAVEIZQ token.
	Set_LLAVEIZQ(antlr.Token)

	// Set_LLAVEDER sets the _LLAVEDER token.
	Set_LLAVEDER(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_declaracion_var returns the _declaracion_var rule contexts.
	Get_declaracion_var() IDeclaracion_varContext

	// Get_asignacion_var returns the _asignacion_var rule contexts.
	Get_asignacion_var() IAsignacion_varContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_breaks returns the _breaks rule contexts.
	Get_breaks() IBreaksContext

	// Get_continues returns the _continues rule contexts.
	Get_continues() IContinuesContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_declaracion_var sets the _declaracion_var rule contexts.
	Set_declaracion_var(IDeclaracion_varContext)

	// Set_asignacion_var sets the _asignacion_var rule contexts.
	Set_asignacion_var(IAsignacion_varContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_breaks sets the _breaks rule contexts.
	Set_breaks(IBreaksContext)

	// Set_continues sets the _continues rule contexts.
	Set_continues(IContinuesContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstruccionContext differentiates from other interfaces.
	IsInstruccionContext()
}

type InstruccionContext struct {
	*antlr.BaseParserRuleContext
	parser           antlr.Parser
	instr            interfaces.Instruction
	_expression      IExpressionContext
	_declaracion_var IDeclaracion_varContext
	_asignacion_var  IAsignacion_varContext
	_LLAVEIZQ        antlr.Token
	_instrucciones   IInstruccionesContext
	_LLAVEDER        antlr.Token
	_breaks          IBreaksContext
	_continues       IContinuesContext
}

func NewEmptyInstruccionContext() *InstruccionContext {
	var p = new(InstruccionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instruccion
	return p
}

func (*InstruccionContext) IsInstruccionContext() {}

func NewInstruccionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *InstruccionContext {
	var p = new(InstruccionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instruccion

	return p
}

func (s *InstruccionContext) GetParser() antlr.Parser { return s.parser }

func (s *InstruccionContext) Get_LLAVEIZQ() antlr.Token { return s._LLAVEIZQ }

func (s *InstruccionContext) Get_LLAVEDER() antlr.Token { return s._LLAVEDER }

func (s *InstruccionContext) Set_LLAVEIZQ(v antlr.Token) { s._LLAVEIZQ = v }

func (s *InstruccionContext) Set_LLAVEDER(v antlr.Token) { s._LLAVEDER = v }

func (s *InstruccionContext) Get_expression() IExpressionContext { return s._expression }

func (s *InstruccionContext) Get_declaracion_var() IDeclaracion_varContext { return s._declaracion_var }

func (s *InstruccionContext) Get_asignacion_var() IAsignacion_varContext { return s._asignacion_var }

func (s *InstruccionContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *InstruccionContext) Get_breaks() IBreaksContext { return s._breaks }

func (s *InstruccionContext) Get_continues() IContinuesContext { return s._continues }

func (s *InstruccionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *InstruccionContext) Set_declaracion_var(v IDeclaracion_varContext) { s._declaracion_var = v }

func (s *InstruccionContext) Set_asignacion_var(v IAsignacion_varContext) { s._asignacion_var = v }

func (s *InstruccionContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *InstruccionContext) Set_breaks(v IBreaksContext) { s._breaks = v }

func (s *InstruccionContext) Set_continues(v IContinuesContext) { s._continues = v }

func (s *InstruccionContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *InstruccionContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *InstruccionContext) CONSOLE() antlr.TerminalNode {
	return s.GetToken(ChemsCONSOLE, 0)
}

func (s *InstruccionContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(ChemsPUNTO, 0)
}

func (s *InstruccionContext) LOG() antlr.TerminalNode {
	return s.GetToken(ChemsLOG, 0)
}

func (s *InstruccionContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsPARIZQ, 0)
}

func (s *InstruccionContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *InstruccionContext) PARDER() antlr.TerminalNode {
	return s.GetToken(ChemsPARDER, 0)
}

func (s *InstruccionContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsPTCOMA, 0)
}

func (s *InstruccionContext) Declaracion_var() IDeclaracion_varContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracion_varContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_varContext)
}

func (s *InstruccionContext) Asignacion_var() IAsignacion_varContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacion_varContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacion_varContext)
}

func (s *InstruccionContext) P_WHILE() antlr.TerminalNode {
	return s.GetToken(ChemsP_WHILE, 0)
}

func (s *InstruccionContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEIZQ, 0)
}

func (s *InstruccionContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *InstruccionContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEDER, 0)
}

func (s *InstruccionContext) Breaks() IBreaksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreaksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreaksContext)
}

func (s *InstruccionContext) Continues() IContinuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinuesContext)
}

func (s *InstruccionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *InstruccionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *InstruccionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstruccion(s)
	}
}

func (s *InstruccionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstruccion(s)
	}
}

func (p *Chems) Instruccion() (localctx IInstruccionContext) {
	this := p
	_ = this

	localctx = NewInstruccionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 4, ChemsRULE_instruccion)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(75)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsCONSOLE:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(45)
			p.Match(ChemsCONSOLE)
		}
		{
			p.SetState(46)
			p.Match(ChemsPUNTO)
		}
		{
			p.SetState(47)
			p.Match(ChemsLOG)
		}
		{
			p.SetState(48)
			p.Match(ChemsPARIZQ)
		}
		{
			p.SetState(49)

			var _x = p.Expression()

			localctx.(*InstruccionContext)._expression = _x
		}
		{
			p.SetState(50)
			p.Match(ChemsPARDER)
		}
		{
			p.SetState(51)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*InstruccionContext).instr = instruction.NewImprimir(localctx.(*InstruccionContext).Get_expression().GetP())

	case ChemsLET:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(54)

			var _x = p.Declaracion_var()

			localctx.(*InstruccionContext)._declaracion_var = _x
		}
		{
			p.SetState(55)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_declaracion_var().GetI()

	case ChemsID:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(58)

			var _x = p.Asignacion_var()

			localctx.(*InstruccionContext)._asignacion_var = _x
		}
		{
			p.SetState(59)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_asignacion_var().GetI()

	case ChemsP_WHILE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(62)
			p.Match(ChemsP_WHILE)
		}
		{
			p.SetState(63)

			var _x = p.Expression()

			localctx.(*InstruccionContext)._expression = _x
		}
		{
			p.SetState(64)

			var _m = p.Match(ChemsLLAVEIZQ)

			localctx.(*InstruccionContext)._LLAVEIZQ = _m
		}
		{
			p.SetState(65)

			var _x = p.Instrucciones()

			localctx.(*InstruccionContext)._instrucciones = _x
		}
		{
			p.SetState(66)

			var _m = p.Match(ChemsLLAVEDER)

			localctx.(*InstruccionContext)._LLAVEDER = _m
		}
		localctx.(*InstruccionContext).instr = instruction.NewWhile(localctx.(*InstruccionContext).Get_expression().GetP(), localctx.(*InstruccionContext).Get_instrucciones().GetL(), localctx.(*InstruccionContext).Get_LLAVEIZQ().GetLine(), localctx.(*InstruccionContext).Get_LLAVEDER().GetColumn())

	case ChemsBREAK:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(69)

			var _x = p.Breaks()

			localctx.(*InstruccionContext)._breaks = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_breaks().GetI()

	case ChemsCONTINUE:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(72)

			var _x = p.Continues()

			localctx.(*InstruccionContext)._continues = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_continues().GetI()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAsignacion_varContext is an interface to support dynamic dispatch.
type IAsignacion_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetI returns the i attribute.
	GetI() interfaces.Instruction

	// SetI sets the i attribute.
	SetI(interfaces.Instruction)

	// IsAsignacion_varContext differentiates from other interfaces.
	IsAsignacion_varContext()
}

type Asignacion_varContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	i           interfaces.Instruction
	id          antlr.Token
	_expression IExpressionContext
}

func NewEmptyAsignacion_varContext() *Asignacion_varContext {
	var p = new(Asignacion_varContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_asignacion_var
	return p
}

func (*Asignacion_varContext) IsAsignacion_varContext() {}

func NewAsignacion_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignacion_varContext {
	var p = new(Asignacion_varContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_asignacion_var

	return p
}

func (s *Asignacion_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignacion_varContext) GetId() antlr.Token { return s.id }

func (s *Asignacion_varContext) SetId(v antlr.Token) { s.id = v }

func (s *Asignacion_varContext) Get_expression() IExpressionContext { return s._expression }

func (s *Asignacion_varContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Asignacion_varContext) GetI() interfaces.Instruction { return s.i }

func (s *Asignacion_varContext) SetI(v interfaces.Instruction) { s.i = v }

func (s *Asignacion_varContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *Asignacion_varContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Asignacion_varContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Asignacion_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignacion_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Asignacion_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterAsignacion_var(s)
	}
}

func (s *Asignacion_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitAsignacion_var(s)
	}
}

func (p *Chems) Asignacion_var() (localctx IAsignacion_varContext) {
	this := p
	_ = this

	localctx = NewAsignacion_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_asignacion_var)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(77)

		var _m = p.Match(ChemsID)

		localctx.(*Asignacion_varContext).id = _m
	}
	{
		p.SetState(78)
		p.Match(ChemsIGUAL)
	}
	{
		p.SetState(79)

		var _x = p.Expression()

		localctx.(*Asignacion_varContext)._expression = _x
	}

	linea := localctx.(*Asignacion_varContext).id.GetLine()
	col := localctx.(*Asignacion_varContext).id.GetColumn()
	localctx.(*Asignacion_varContext).i = instruction.NewAssignment((func() string {
		if localctx.(*Asignacion_varContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Asignacion_varContext).GetId().GetText()
		}
	}()), localctx.(*Asignacion_varContext).Get_expression().GetP(), linea, col)

	return localctx
}

// IDeclaracion_varContext is an interface to support dynamic dispatch.
type IDeclaracion_varContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_LET returns the _LET token.
	Get_LET() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// Set_LET sets the _LET token.
	Set_LET(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// Get_mutable returns the _mutable rule contexts.
	Get_mutable() IMutableContext

	// Get_types returns the _types rule contexts.
	Get_types() ITypesContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_mutable sets the _mutable rule contexts.
	Set_mutable(IMutableContext)

	// Set_types sets the _types rule contexts.
	Set_types(ITypesContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetI returns the i attribute.
	GetI() interfaces.Instruction

	// SetI sets the i attribute.
	SetI(interfaces.Instruction)

	// IsDeclaracion_varContext differentiates from other interfaces.
	IsDeclaracion_varContext()
}

type Declaracion_varContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	i           interfaces.Instruction
	_LET        antlr.Token
	_mutable    IMutableContext
	id          antlr.Token
	_types      ITypesContext
	_expression IExpressionContext
}

func NewEmptyDeclaracion_varContext() *Declaracion_varContext {
	var p = new(Declaracion_varContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_declaracion_var
	return p
}

func (*Declaracion_varContext) IsDeclaracion_varContext() {}

func NewDeclaracion_varContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Declaracion_varContext {
	var p = new(Declaracion_varContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_declaracion_var

	return p
}

func (s *Declaracion_varContext) GetParser() antlr.Parser { return s.parser }

func (s *Declaracion_varContext) Get_LET() antlr.Token { return s._LET }

func (s *Declaracion_varContext) GetId() antlr.Token { return s.id }

func (s *Declaracion_varContext) Set_LET(v antlr.Token) { s._LET = v }

func (s *Declaracion_varContext) SetId(v antlr.Token) { s.id = v }

func (s *Declaracion_varContext) Get_mutable() IMutableContext { return s._mutable }

func (s *Declaracion_varContext) Get_types() ITypesContext { return s._types }

func (s *Declaracion_varContext) Get_expression() IExpressionContext { return s._expression }

func (s *Declaracion_varContext) Set_mutable(v IMutableContext) { s._mutable = v }

func (s *Declaracion_varContext) Set_types(v ITypesContext) { s._types = v }

func (s *Declaracion_varContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Declaracion_varContext) GetI() interfaces.Instruction { return s.i }

func (s *Declaracion_varContext) SetI(v interfaces.Instruction) { s.i = v }

func (s *Declaracion_varContext) LET() antlr.TerminalNode {
	return s.GetToken(ChemsLET, 0)
}

func (s *Declaracion_varContext) Mutable() IMutableContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMutableContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMutableContext)
}

func (s *Declaracion_varContext) Types() ITypesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITypesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITypesContext)
}

func (s *Declaracion_varContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *Declaracion_varContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Declaracion_varContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Declaracion_varContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Declaracion_varContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Declaracion_varContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterDeclaracion_var(s)
	}
}

func (s *Declaracion_varContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitDeclaracion_var(s)
	}
}

func (p *Chems) Declaracion_var() (localctx IDeclaracion_varContext) {
	this := p
	_ = this

	localctx = NewDeclaracion_varContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 8, ChemsRULE_declaracion_var)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(82)

		var _m = p.Match(ChemsLET)

		localctx.(*Declaracion_varContext)._LET = _m
	}
	{
		p.SetState(83)

		var _x = p.Mutable()

		localctx.(*Declaracion_varContext)._mutable = _x
	}
	{
		p.SetState(84)

		var _m = p.Match(ChemsID)

		localctx.(*Declaracion_varContext).id = _m
	}
	{
		p.SetState(85)

		var _x = p.Types()

		localctx.(*Declaracion_varContext)._types = _x
	}
	{
		p.SetState(86)
		p.Match(ChemsIGUAL)
	}
	{
		p.SetState(87)

		var _x = p.Expression()

		localctx.(*Declaracion_varContext)._expression = _x
	}
	localctx.(*Declaracion_varContext).i = instruction.NewDeclaracion((func() string {
		if localctx.(*Declaracion_varContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Declaracion_varContext).GetId().GetText()
		}
	}()), localctx.(*Declaracion_varContext).Get_types().GetL(), localctx.(*Declaracion_varContext).Get_expression().GetP(), localctx.(*Declaracion_varContext).Get_mutable().GetMut(), localctx.(*Declaracion_varContext).Get_LET().GetLine(), localctx.(*Declaracion_varContext).Get_LET().GetColumn())

	return localctx
}

// IMutableContext is an interface to support dynamic dispatch.
type IMutableContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetMut returns the mut attribute.
	GetMut() bool

	// SetMut sets the mut attribute.
	SetMut(bool)

	// IsMutableContext differentiates from other interfaces.
	IsMutableContext()
}

type MutableContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	mut    bool
}

func NewEmptyMutableContext() *MutableContext {
	var p = new(MutableContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_mutable
	return p
}

func (*MutableContext) IsMutableContext() {}

func NewMutableContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MutableContext {
	var p = new(MutableContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_mutable

	return p
}

func (s *MutableContext) GetParser() antlr.Parser { return s.parser }

func (s *MutableContext) GetMut() bool { return s.mut }

func (s *MutableContext) SetMut(v bool) { s.mut = v }

func (s *MutableContext) MUT() antlr.TerminalNode {
	return s.GetToken(ChemsMUT, 0)
}

func (s *MutableContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MutableContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MutableContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterMutable(s)
	}
}

func (s *MutableContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitMutable(s)
	}
}

func (p *Chems) Mutable() (localctx IMutableContext) {
	this := p
	_ = this

	localctx = NewMutableContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 10, ChemsRULE_mutable)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(93)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsMUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(90)
			p.Match(ChemsMUT)
		}
		localctx.(*MutableContext).mut = true

	case ChemsID:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*MutableContext).mut = false

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ITypesContext is an interface to support dynamic dispatch.
type ITypesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_tipo_vector returns the _tipo_vector rule contexts.
	Get_tipo_vector() ITipo_vectorContext

	// GetA returns the a rule contexts.
	GetA() IAsignar_ArrayContext

	// Get_tipo_d returns the _tipo_d rule contexts.
	Get_tipo_d() ITipo_dContext

	// Set_tipo_vector sets the _tipo_vector rule contexts.
	Set_tipo_vector(ITipo_vectorContext)

	// SetA sets the a rule contexts.
	SetA(IAsignar_ArrayContext)

	// Set_tipo_d sets the _tipo_d rule contexts.
	Set_tipo_d(ITipo_dContext)

	// GetL returns the l attribute.
	GetL() interfaces.TipoSimbolo

	// SetL sets the l attribute.
	SetL(interfaces.TipoSimbolo)

	// IsTypesContext differentiates from other interfaces.
	IsTypesContext()
}

type TypesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            interfaces.TipoSimbolo
	_tipo_vector ITipo_vectorContext
	a            IAsignar_ArrayContext
	_tipo_d      ITipo_dContext
}

func NewEmptyTypesContext() *TypesContext {
	var p = new(TypesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_types
	return p
}

func (*TypesContext) IsTypesContext() {}

func NewTypesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *TypesContext {
	var p = new(TypesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_types

	return p
}

func (s *TypesContext) GetParser() antlr.Parser { return s.parser }

func (s *TypesContext) Get_tipo_vector() ITipo_vectorContext { return s._tipo_vector }

func (s *TypesContext) GetA() IAsignar_ArrayContext { return s.a }

func (s *TypesContext) Get_tipo_d() ITipo_dContext { return s._tipo_d }

func (s *TypesContext) Set_tipo_vector(v ITipo_vectorContext) { s._tipo_vector = v }

func (s *TypesContext) SetA(v IAsignar_ArrayContext) { s.a = v }

func (s *TypesContext) Set_tipo_d(v ITipo_dContext) { s._tipo_d = v }

func (s *TypesContext) GetL() interfaces.TipoSimbolo { return s.l }

func (s *TypesContext) SetL(v interfaces.TipoSimbolo) { s.l = v }

func (s *TypesContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(ChemsDOSPUNTOS, 0)
}

func (s *TypesContext) Tipo_vector() ITipo_vectorContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_vectorContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_vectorContext)
}

func (s *TypesContext) Asignar_Array() IAsignar_ArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignar_ArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignar_ArrayContext)
}

func (s *TypesContext) Tipo_d() ITipo_dContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_dContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_dContext)
}

func (s *TypesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *TypesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *TypesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterTypes(s)
	}
}

func (s *TypesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitTypes(s)
	}
}

func (p *Chems) Types() (localctx ITypesContext) {
	this := p
	_ = this

	localctx = NewTypesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 12, ChemsRULE_types)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(106)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 3, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(95)
			p.Match(ChemsDOSPUNTOS)
		}
		{
			p.SetState(96)

			var _x = p.Tipo_vector()

			localctx.(*TypesContext)._tipo_vector = _x
		}
		localctx.(*TypesContext).l = localctx.(*TypesContext).Get_tipo_vector().GetT()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(99)

			var _x = p.Asignar_Array()

			localctx.(*TypesContext).a = _x
		}

		dim := arrayList.New()
		dim.Add(localctx.(*TypesContext).GetA().GetD())
		localctx.(*TypesContext).l = interfaces.TipoSimbolo{interfaces.ARRAY, dim}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(102)
			p.Match(ChemsDOSPUNTOS)
		}
		{
			p.SetState(103)

			var _x = p.Tipo_d()

			localctx.(*TypesContext)._tipo_d = _x
		}
		localctx.(*TypesContext).l = interfaces.TipoSimbolo{localctx.(*TypesContext).Get_tipo_d().GetT(), arrayList.New()}

	}

	return localctx
}

// ITipo_dContext is an interface to support dynamic dispatch.
type ITipo_dContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetT returns the t attribute.
	GetT() interfaces.TipoExpresion

	// SetT sets the t attribute.
	SetT(interfaces.TipoExpresion)

	// IsTipo_dContext differentiates from other interfaces.
	IsTipo_dContext()
}

type Tipo_dContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	t      interfaces.TipoExpresion
}

func NewEmptyTipo_dContext() *Tipo_dContext {
	var p = new(Tipo_dContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_tipo_d
	return p
}

func (*Tipo_dContext) IsTipo_dContext() {}

func NewTipo_dContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_dContext {
	var p = new(Tipo_dContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_tipo_d

	return p
}

func (s *Tipo_dContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_dContext) GetT() interfaces.TipoExpresion { return s.t }

func (s *Tipo_dContext) SetT(v interfaces.TipoExpresion) { s.t = v }

func (s *Tipo_dContext) INT() antlr.TerminalNode {
	return s.GetToken(ChemsINT, 0)
}

func (s *Tipo_dContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsFLOAT, 0)
}

func (s *Tipo_dContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ChemsBOOL, 0)
}

func (s *Tipo_dContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ChemsCHAR, 0)
}

func (s *Tipo_dContext) STR() antlr.TerminalNode {
	return s.GetToken(ChemsSTR, 0)
}

func (s *Tipo_dContext) P_STRING() antlr.TerminalNode {
	return s.GetToken(ChemsP_STRING, 0)
}

func (s *Tipo_dContext) USIZE() antlr.TerminalNode {
	return s.GetToken(ChemsUSIZE, 0)
}

func (s *Tipo_dContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Tipo_dContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_dContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_dContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterTipo_d(s)
	}
}

func (s *Tipo_dContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitTipo_d(s)
	}
}

func (p *Chems) Tipo_d() (localctx ITipo_dContext) {
	this := p
	_ = this

	localctx = NewTipo_dContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 14, ChemsRULE_tipo_d)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(124)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(108)
			p.Match(ChemsINT)
		}
		localctx.(*Tipo_dContext).t = interfaces.INTEGER

	case ChemsFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(110)
			p.Match(ChemsFLOAT)
		}
		localctx.(*Tipo_dContext).t = interfaces.FLOAT

	case ChemsBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(112)
			p.Match(ChemsBOOL)
		}
		localctx.(*Tipo_dContext).t = interfaces.BOOLEAN

	case ChemsCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(114)
			p.Match(ChemsCHAR)
		}
		localctx.(*Tipo_dContext).t = interfaces.CHAR

	case ChemsSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(116)
			p.Match(ChemsSTR)
		}
		localctx.(*Tipo_dContext).t = interfaces.STRING

	case ChemsP_STRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(118)
			p.Match(ChemsP_STRING)
		}
		localctx.(*Tipo_dContext).t = interfaces.STR

	case ChemsUSIZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(120)
			p.Match(ChemsUSIZE)
		}
		localctx.(*Tipo_dContext).t = interfaces.USIZE

	case ChemsID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(122)
			p.Match(ChemsID)
		}
		localctx.(*Tipo_dContext).t = interfaces.STRUCT

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IAsignar_ArrayContext is an interface to support dynamic dispatch.
type IAsignar_ArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_dimensiones returns the _dimensiones rule contexts.
	Get_dimensiones() IDimensionesContext

	// Set_dimensiones sets the _dimensiones rule contexts.
	Set_dimensiones(IDimensionesContext)

	// GetD returns the d attribute.
	GetD() interfaces.Dimensions

	// SetD sets the d attribute.
	SetD(interfaces.Dimensions)

	// IsAsignar_ArrayContext differentiates from other interfaces.
	IsAsignar_ArrayContext()
}

type Asignar_ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	d            interfaces.Dimensions
	_dimensiones IDimensionesContext
}

func NewEmptyAsignar_ArrayContext() *Asignar_ArrayContext {
	var p = new(Asignar_ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_asignar_Array
	return p
}

func (*Asignar_ArrayContext) IsAsignar_ArrayContext() {}

func NewAsignar_ArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Asignar_ArrayContext {
	var p = new(Asignar_ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_asignar_Array

	return p
}

func (s *Asignar_ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Asignar_ArrayContext) Get_dimensiones() IDimensionesContext { return s._dimensiones }

func (s *Asignar_ArrayContext) Set_dimensiones(v IDimensionesContext) { s._dimensiones = v }

func (s *Asignar_ArrayContext) GetD() interfaces.Dimensions { return s.d }

func (s *Asignar_ArrayContext) SetD(v interfaces.Dimensions) { s.d = v }

func (s *Asignar_ArrayContext) DOSPUNTOS() antlr.TerminalNode {
	return s.GetToken(ChemsDOSPUNTOS, 0)
}

func (s *Asignar_ArrayContext) Dimensiones() IDimensionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimensionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimensionesContext)
}

func (s *Asignar_ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Asignar_ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Asignar_ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterAsignar_Array(s)
	}
}

func (s *Asignar_ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitAsignar_Array(s)
	}
}

func (p *Chems) Asignar_Array() (localctx IAsignar_ArrayContext) {
	this := p
	_ = this

	localctx = NewAsignar_ArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 16, ChemsRULE_asignar_Array)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(126)
		p.Match(ChemsDOSPUNTOS)
	}
	{
		p.SetState(127)

		var _x = p.Dimensiones()

		localctx.(*Asignar_ArrayContext)._dimensiones = _x
	}
	localctx.(*Asignar_ArrayContext).d = localctx.(*Asignar_ArrayContext).Get_dimensiones().GetD()

	return localctx
}

// IDimensionesContext is an interface to support dynamic dispatch.
type IDimensionesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_tipo_d returns the _tipo_d rule contexts.
	Get_tipo_d() ITipo_dContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_dimensiones returns the _dimensiones rule contexts.
	Get_dimensiones() IDimensionesContext

	// Set_tipo_d sets the _tipo_d rule contexts.
	Set_tipo_d(ITipo_dContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_dimensiones sets the _dimensiones rule contexts.
	Set_dimensiones(IDimensionesContext)

	// GetD returns the d attribute.
	GetD() interfaces.Dimensions

	// SetD sets the d attribute.
	SetD(interfaces.Dimensions)

	// IsDimensionesContext differentiates from other interfaces.
	IsDimensionesContext()
}

type DimensionesContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	d            interfaces.Dimensions
	_tipo_d      ITipo_dContext
	_expression  IExpressionContext
	_dimensiones IDimensionesContext
}

func NewEmptyDimensionesContext() *DimensionesContext {
	var p = new(DimensionesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_dimensiones
	return p
}

func (*DimensionesContext) IsDimensionesContext() {}

func NewDimensionesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DimensionesContext {
	var p = new(DimensionesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_dimensiones

	return p
}

func (s *DimensionesContext) GetParser() antlr.Parser { return s.parser }

func (s *DimensionesContext) Get_tipo_d() ITipo_dContext { return s._tipo_d }

func (s *DimensionesContext) Get_expression() IExpressionContext { return s._expression }

func (s *DimensionesContext) Get_dimensiones() IDimensionesContext { return s._dimensiones }

func (s *DimensionesContext) Set_tipo_d(v ITipo_dContext) { s._tipo_d = v }

func (s *DimensionesContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *DimensionesContext) Set_dimensiones(v IDimensionesContext) { s._dimensiones = v }

func (s *DimensionesContext) GetD() interfaces.Dimensions { return s.d }

func (s *DimensionesContext) SetD(v interfaces.Dimensions) { s.d = v }

func (s *DimensionesContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsCORIZQ, 0)
}

func (s *DimensionesContext) Tipo_d() ITipo_dContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_dContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_dContext)
}

func (s *DimensionesContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsPTCOMA, 0)
}

func (s *DimensionesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *DimensionesContext) CORDER() antlr.TerminalNode {
	return s.GetToken(ChemsCORDER, 0)
}

func (s *DimensionesContext) Dimensiones() IDimensionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDimensionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDimensionesContext)
}

func (s *DimensionesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DimensionesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DimensionesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterDimensiones(s)
	}
}

func (s *DimensionesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitDimensiones(s)
	}
}

func (p *Chems) Dimensiones() (localctx IDimensionesContext) {
	this := p
	_ = this

	localctx = NewDimensionesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 18, ChemsRULE_dimensiones)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(144)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 5, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(130)
			p.Match(ChemsCORIZQ)
		}
		{
			p.SetState(131)

			var _x = p.Tipo_d()

			localctx.(*DimensionesContext)._tipo_d = _x
		}
		{
			p.SetState(132)
			p.Match(ChemsPTCOMA)
		}
		{
			p.SetState(133)

			var _x = p.Expression()

			localctx.(*DimensionesContext)._expression = _x
		}
		{
			p.SetState(134)
			p.Match(ChemsCORDER)
		}

		list := arrayList.New()
		list.Add(localctx.(*DimensionesContext).Get_expression().GetP())
		localctx.(*DimensionesContext).d = interfaces.Dimensions{localctx.(*DimensionesContext).Get_tipo_d().GetT(), list}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(137)
			p.Match(ChemsCORIZQ)
		}
		{
			p.SetState(138)

			var _x = p.Dimensiones()

			localctx.(*DimensionesContext)._dimensiones = _x
		}
		{
			p.SetState(139)
			p.Match(ChemsPTCOMA)
		}
		{
			p.SetState(140)

			var _x = p.Expression()

			localctx.(*DimensionesContext)._expression = _x
		}
		{
			p.SetState(141)
			p.Match(ChemsCORDER)
		}
		localctx.(*DimensionesContext).Get_dimensiones().GetD().Dimensions.Add(localctx.(*DimensionesContext).Get_expression().GetP())
		localctx.(*DimensionesContext).SetD(localctx.(*DimensionesContext).Get_dimensiones().GetD())

	}

	return localctx
}

// ITipo_vectorContext is an interface to support dynamic dispatch.
type ITipo_vectorContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_vectores returns the _vectores rule contexts.
	Get_vectores() IVectoresContext

	// Set_vectores sets the _vectores rule contexts.
	Set_vectores(IVectoresContext)

	// GetT returns the t attribute.
	GetT() interfaces.TipoSimbolo

	// SetT sets the t attribute.
	SetT(interfaces.TipoSimbolo)

	// IsTipo_vectorContext differentiates from other interfaces.
	IsTipo_vectorContext()
}

type Tipo_vectorContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	t         interfaces.TipoSimbolo
	_vectores IVectoresContext
}

func NewEmptyTipo_vectorContext() *Tipo_vectorContext {
	var p = new(Tipo_vectorContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_tipo_vector
	return p
}

func (*Tipo_vectorContext) IsTipo_vectorContext() {}

func NewTipo_vectorContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Tipo_vectorContext {
	var p = new(Tipo_vectorContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_tipo_vector

	return p
}

func (s *Tipo_vectorContext) GetParser() antlr.Parser { return s.parser }

func (s *Tipo_vectorContext) Get_vectores() IVectoresContext { return s._vectores }

func (s *Tipo_vectorContext) Set_vectores(v IVectoresContext) { s._vectores = v }

func (s *Tipo_vectorContext) GetT() interfaces.TipoSimbolo { return s.t }

func (s *Tipo_vectorContext) SetT(v interfaces.TipoSimbolo) { s.t = v }

func (s *Tipo_vectorContext) VECT() antlr.TerminalNode {
	return s.GetToken(ChemsVECT, 0)
}

func (s *Tipo_vectorContext) MENOR() antlr.TerminalNode {
	return s.GetToken(ChemsMENOR, 0)
}

func (s *Tipo_vectorContext) Vectores() IVectoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVectoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVectoresContext)
}

func (s *Tipo_vectorContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsMAYOR, 0)
}

func (s *Tipo_vectorContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Tipo_vectorContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Tipo_vectorContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterTipo_vector(s)
	}
}

func (s *Tipo_vectorContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitTipo_vector(s)
	}
}

func (p *Chems) Tipo_vector() (localctx ITipo_vectorContext) {
	this := p
	_ = this

	localctx = NewTipo_vectorContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 20, ChemsRULE_tipo_vector)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(146)
		p.Match(ChemsVECT)
	}
	{
		p.SetState(147)
		p.Match(ChemsMENOR)
	}
	{
		p.SetState(148)

		var _x = p.Vectores()

		localctx.(*Tipo_vectorContext)._vectores = _x
	}
	{
		p.SetState(149)
		p.Match(ChemsMAYOR)
	}
	localctx.(*Tipo_vectorContext).t = interfaces.TipoSimbolo{interfaces.VECTOR, localctx.(*Tipo_vectorContext).Get_vectores().GetL()}

	return localctx
}

// IVectoresContext is an interface to support dynamic dispatch.
type IVectoresContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// Get_vectores returns the _vectores rule contexts.
	Get_vectores() IVectoresContext

	// Set_vectores sets the _vectores rule contexts.
	Set_vectores(IVectoresContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsVectoresContext differentiates from other interfaces.
	IsVectoresContext()
}

type VectoresContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	l         *arrayList.List
	id        antlr.Token
	_vectores IVectoresContext
}

func NewEmptyVectoresContext() *VectoresContext {
	var p = new(VectoresContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_vectores
	return p
}

func (*VectoresContext) IsVectoresContext() {}

func NewVectoresContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *VectoresContext {
	var p = new(VectoresContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_vectores

	return p
}

func (s *VectoresContext) GetParser() antlr.Parser { return s.parser }

func (s *VectoresContext) GetId() antlr.Token { return s.id }

func (s *VectoresContext) SetId(v antlr.Token) { s.id = v }

func (s *VectoresContext) Get_vectores() IVectoresContext { return s._vectores }

func (s *VectoresContext) Set_vectores(v IVectoresContext) { s._vectores = v }

func (s *VectoresContext) GetL() *arrayList.List { return s.l }

func (s *VectoresContext) SetL(v *arrayList.List) { s.l = v }

func (s *VectoresContext) INT() antlr.TerminalNode {
	return s.GetToken(ChemsINT, 0)
}

func (s *VectoresContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsFLOAT, 0)
}

func (s *VectoresContext) BOOL() antlr.TerminalNode {
	return s.GetToken(ChemsBOOL, 0)
}

func (s *VectoresContext) CHAR() antlr.TerminalNode {
	return s.GetToken(ChemsCHAR, 0)
}

func (s *VectoresContext) STR() antlr.TerminalNode {
	return s.GetToken(ChemsSTR, 0)
}

func (s *VectoresContext) P_STRING() antlr.TerminalNode {
	return s.GetToken(ChemsP_STRING, 0)
}

func (s *VectoresContext) USIZE() antlr.TerminalNode {
	return s.GetToken(ChemsUSIZE, 0)
}

func (s *VectoresContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *VectoresContext) VECT() antlr.TerminalNode {
	return s.GetToken(ChemsVECT, 0)
}

func (s *VectoresContext) MENOR() antlr.TerminalNode {
	return s.GetToken(ChemsMENOR, 0)
}

func (s *VectoresContext) Vectores() IVectoresContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IVectoresContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IVectoresContext)
}

func (s *VectoresContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsMAYOR, 0)
}

func (s *VectoresContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *VectoresContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *VectoresContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterVectores(s)
	}
}

func (s *VectoresContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitVectores(s)
	}
}

func (p *Chems) Vectores() (localctx IVectoresContext) {
	this := p
	_ = this

	localctx = NewVectoresContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 22, ChemsRULE_vectores)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(174)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(152)
			p.Match(ChemsINT)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.INTEGER)

	case ChemsFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(154)
			p.Match(ChemsFLOAT)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.FLOAT)

	case ChemsBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(156)
			p.Match(ChemsBOOL)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.BOOLEAN)

	case ChemsCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(158)
			p.Match(ChemsCHAR)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.CHAR)

	case ChemsSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(160)
			p.Match(ChemsSTR)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.STRING)

	case ChemsP_STRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(162)
			p.Match(ChemsP_STRING)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.STR)

	case ChemsUSIZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(164)
			p.Match(ChemsUSIZE)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.USIZE)

	case ChemsID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(166)

			var _m = p.Match(ChemsID)

			localctx.(*VectoresContext).id = _m
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add((func() string {
			if localctx.(*VectoresContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*VectoresContext).GetId().GetText()
			}
		}()))
		localctx.(*VectoresContext).l.Add(interfaces.STRUCT)

	case ChemsVECT:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(168)
			p.Match(ChemsVECT)
		}
		{
			p.SetState(169)
			p.Match(ChemsMENOR)
		}
		{
			p.SetState(170)

			var _x = p.Vectores()

			localctx.(*VectoresContext)._vectores = _x
		}
		{
			p.SetState(171)
			p.Match(ChemsMAYOR)
		}

		localctx.(*VectoresContext).Get_vectores().GetL().Add(interfaces.VECTOR)
		localctx.(*VectoresContext).l = localctx.(*VectoresContext).Get_vectores().GetL()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// IBreaksContext is an interface to support dynamic dispatch.
type IBreaksContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_BREAK returns the _BREAK token.
	Get_BREAK() antlr.Token

	// Set_BREAK sets the _BREAK token.
	Set_BREAK(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetI returns the i attribute.
	GetI() interfaces.Instruction

	// SetI sets the i attribute.
	SetI(interfaces.Instruction)

	// IsBreaksContext differentiates from other interfaces.
	IsBreaksContext()
}

type BreaksContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	i           interfaces.Instruction
	_BREAK      antlr.Token
	_expression IExpressionContext
}

func NewEmptyBreaksContext() *BreaksContext {
	var p = new(BreaksContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_breaks
	return p
}

func (*BreaksContext) IsBreaksContext() {}

func NewBreaksContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *BreaksContext {
	var p = new(BreaksContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_breaks

	return p
}

func (s *BreaksContext) GetParser() antlr.Parser { return s.parser }

func (s *BreaksContext) Get_BREAK() antlr.Token { return s._BREAK }

func (s *BreaksContext) Set_BREAK(v antlr.Token) { s._BREAK = v }

func (s *BreaksContext) Get_expression() IExpressionContext { return s._expression }

func (s *BreaksContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *BreaksContext) GetI() interfaces.Instruction { return s.i }

func (s *BreaksContext) SetI(v interfaces.Instruction) { s.i = v }

func (s *BreaksContext) BREAK() antlr.TerminalNode {
	return s.GetToken(ChemsBREAK, 0)
}

func (s *BreaksContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *BreaksContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsPTCOMA, 0)
}

func (s *BreaksContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *BreaksContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *BreaksContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterBreaks(s)
	}
}

func (s *BreaksContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitBreaks(s)
	}
}

func (p *Chems) Breaks() (localctx IBreaksContext) {
	this := p
	_ = this

	localctx = NewBreaksContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 24, ChemsRULE_breaks)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(184)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 7, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(176)

			var _m = p.Match(ChemsBREAK)

			localctx.(*BreaksContext)._BREAK = _m
		}
		{
			p.SetState(177)

			var _x = p.Expression()

			localctx.(*BreaksContext)._expression = _x
		}
		{
			p.SetState(178)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*BreaksContext).i = instruction.NewBreak(localctx.(*BreaksContext).Get_expression().GetP(), true, localctx.(*BreaksContext).Get_BREAK().GetLine(), localctx.(*BreaksContext).Get_BREAK().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(181)

			var _m = p.Match(ChemsBREAK)

			localctx.(*BreaksContext)._BREAK = _m
		}
		{
			p.SetState(182)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*BreaksContext).i = instruction.NewBreak(expresion.NewPrimitivo(1, interfaces.INTEGER, 0, 0), false, localctx.(*BreaksContext).Get_BREAK().GetLine(), localctx.(*BreaksContext).Get_BREAK().GetColumn())

	}

	return localctx
}

// IContinuesContext is an interface to support dynamic dispatch.
type IContinuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_CONTINUE returns the _CONTINUE token.
	Get_CONTINUE() antlr.Token

	// Set_CONTINUE sets the _CONTINUE token.
	Set_CONTINUE(antlr.Token)

	// GetI returns the i attribute.
	GetI() interfaces.Instruction

	// SetI sets the i attribute.
	SetI(interfaces.Instruction)

	// IsContinuesContext differentiates from other interfaces.
	IsContinuesContext()
}

type ContinuesContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	i         interfaces.Instruction
	_CONTINUE antlr.Token
}

func NewEmptyContinuesContext() *ContinuesContext {
	var p = new(ContinuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_continues
	return p
}

func (*ContinuesContext) IsContinuesContext() {}

func NewContinuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ContinuesContext {
	var p = new(ContinuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_continues

	return p
}

func (s *ContinuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ContinuesContext) Get_CONTINUE() antlr.Token { return s._CONTINUE }

func (s *ContinuesContext) Set_CONTINUE(v antlr.Token) { s._CONTINUE = v }

func (s *ContinuesContext) GetI() interfaces.Instruction { return s.i }

func (s *ContinuesContext) SetI(v interfaces.Instruction) { s.i = v }

func (s *ContinuesContext) CONTINUE() antlr.TerminalNode {
	return s.GetToken(ChemsCONTINUE, 0)
}

func (s *ContinuesContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsPTCOMA, 0)
}

func (s *ContinuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ContinuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ContinuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterContinues(s)
	}
}

func (s *ContinuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitContinues(s)
	}
}

func (p *Chems) Continues() (localctx IContinuesContext) {
	this := p
	_ = this

	localctx = NewContinuesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ChemsRULE_continues)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(186)

		var _m = p.Match(ChemsCONTINUE)

		localctx.(*ContinuesContext)._CONTINUE = _m
	}
	{
		p.SetState(187)
		p.Match(ChemsPTCOMA)
	}
	localctx.(*ContinuesContext).i = instruction.NewContinue("continue", localctx.(*ContinuesContext).Get_CONTINUE().GetLine(), localctx.(*ContinuesContext).Get_CONTINUE().GetColumn())

	return localctx
}

// IExpressionContext is an interface to support dynamic dispatch.
type IExpressionContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expr_arit returns the _expr_arit rule contexts.
	Get_expr_arit() IExpr_aritContext

	// Set_expr_arit sets the _expr_arit rule contexts.
	Set_expr_arit(IExpr_aritContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsExpressionContext differentiates from other interfaces.
	IsExpressionContext()
}

type ExpressionContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	p          interfaces.Expresion
	_expr_arit IExpr_aritContext
}

func NewEmptyExpressionContext() *ExpressionContext {
	var p = new(ExpressionContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expression
	return p
}

func (*ExpressionContext) IsExpressionContext() {}

func NewExpressionContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ExpressionContext {
	var p = new(ExpressionContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expression

	return p
}

func (s *ExpressionContext) GetParser() antlr.Parser { return s.parser }

func (s *ExpressionContext) Get_expr_arit() IExpr_aritContext { return s._expr_arit }

func (s *ExpressionContext) Set_expr_arit(v IExpr_aritContext) { s._expr_arit = v }

func (s *ExpressionContext) GetP() interfaces.Expresion { return s.p }

func (s *ExpressionContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *ExpressionContext) Expr_arit() IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *ExpressionContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ExpressionContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ExpressionContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpression(s)
	}
}

func (s *ExpressionContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpression(s)
	}
}

func (p *Chems) Expression() (localctx IExpressionContext) {
	this := p
	_ = this

	localctx = NewExpressionContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ChemsRULE_expression)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.EnterOuterAlt(localctx, 1)
	{
		p.SetState(190)

		var _x = p.expr_arit(0)

		localctx.(*ExpressionContext)._expr_arit = _x
	}
	localctx.(*ExpressionContext).p = localctx.(*ExpressionContext).Get_expr_arit().GetP()

	return localctx
}

// IExpr_aritContext is an interface to support dynamic dispatch.
type IExpr_aritContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PARIZQ returns the _PARIZQ token.
	Get_PARIZQ() antlr.Token

	// Get_DIFERENTE returns the _DIFERENTE token.
	Get_DIFERENTE() antlr.Token

	// Get_SUB returns the _SUB token.
	Get_SUB() antlr.Token

	// GetOp returns the op token.
	GetOp() antlr.Token

	// Get_PUNTO returns the _PUNTO token.
	Get_PUNTO() antlr.Token

	// Set_PARIZQ sets the _PARIZQ token.
	Set_PARIZQ(antlr.Token)

	// Set_DIFERENTE sets the _DIFERENTE token.
	Set_DIFERENTE(antlr.Token)

	// Set_SUB sets the _SUB token.
	Set_SUB(antlr.Token)

	// SetOp sets the op token.
	SetOp(antlr.Token)

	// Set_PUNTO sets the _PUNTO token.
	Set_PUNTO(antlr.Token)

	// GetExp returns the exp rule contexts.
	GetExp() IExpr_aritContext

	// GetOpIz returns the opIz rule contexts.
	GetOpIz() IExpr_aritContext

	// GetOpDe returns the opDe rule contexts.
	GetOpDe() IExpr_aritContext

	// Get_primitivo returns the _primitivo rule contexts.
	Get_primitivo() IPrimitivoContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_tipo_d returns the _tipo_d rule contexts.
	Get_tipo_d() ITipo_dContext

	// SetExp sets the exp rule contexts.
	SetExp(IExpr_aritContext)

	// SetOpIz sets the opIz rule contexts.
	SetOpIz(IExpr_aritContext)

	// SetOpDe sets the opDe rule contexts.
	SetOpDe(IExpr_aritContext)

	// Set_primitivo sets the _primitivo rule contexts.
	Set_primitivo(IPrimitivoContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_tipo_d sets the _tipo_d rule contexts.
	Set_tipo_d(ITipo_dContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsExpr_aritContext differentiates from other interfaces.
	IsExpr_aritContext()
}

type Expr_aritContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Expresion
	exp         IExpr_aritContext
	opIz        IExpr_aritContext
	_PARIZQ     antlr.Token
	opDe        IExpr_aritContext
	_DIFERENTE  antlr.Token
	_SUB        antlr.Token
	_primitivo  IPrimitivoContext
	_expression IExpressionContext
	op          antlr.Token
	_tipo_d     ITipo_dContext
	_PUNTO      antlr.Token
}

func NewEmptyExpr_aritContext() *Expr_aritContext {
	var p = new(Expr_aritContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_expr_arit
	return p
}

func (*Expr_aritContext) IsExpr_aritContext() {}

func NewExpr_aritContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Expr_aritContext {
	var p = new(Expr_aritContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_expr_arit

	return p
}

func (s *Expr_aritContext) GetParser() antlr.Parser { return s.parser }

func (s *Expr_aritContext) Get_PARIZQ() antlr.Token { return s._PARIZQ }

func (s *Expr_aritContext) Get_DIFERENTE() antlr.Token { return s._DIFERENTE }

func (s *Expr_aritContext) Get_SUB() antlr.Token { return s._SUB }

func (s *Expr_aritContext) GetOp() antlr.Token { return s.op }

func (s *Expr_aritContext) Get_PUNTO() antlr.Token { return s._PUNTO }

func (s *Expr_aritContext) Set_PARIZQ(v antlr.Token) { s._PARIZQ = v }

func (s *Expr_aritContext) Set_DIFERENTE(v antlr.Token) { s._DIFERENTE = v }

func (s *Expr_aritContext) Set_SUB(v antlr.Token) { s._SUB = v }

func (s *Expr_aritContext) SetOp(v antlr.Token) { s.op = v }

func (s *Expr_aritContext) Set_PUNTO(v antlr.Token) { s._PUNTO = v }

func (s *Expr_aritContext) GetExp() IExpr_aritContext { return s.exp }

func (s *Expr_aritContext) GetOpIz() IExpr_aritContext { return s.opIz }

func (s *Expr_aritContext) GetOpDe() IExpr_aritContext { return s.opDe }

func (s *Expr_aritContext) Get_primitivo() IPrimitivoContext { return s._primitivo }

func (s *Expr_aritContext) Get_expression() IExpressionContext { return s._expression }

func (s *Expr_aritContext) Get_tipo_d() ITipo_dContext { return s._tipo_d }

func (s *Expr_aritContext) SetExp(v IExpr_aritContext) { s.exp = v }

func (s *Expr_aritContext) SetOpIz(v IExpr_aritContext) { s.opIz = v }

func (s *Expr_aritContext) SetOpDe(v IExpr_aritContext) { s.opDe = v }

func (s *Expr_aritContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_tipo_d(v ITipo_dContext) { s._tipo_d = v }

func (s *Expr_aritContext) GetP() interfaces.Expresion { return s.p }

func (s *Expr_aritContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *Expr_aritContext) INT() antlr.TerminalNode {
	return s.GetToken(ChemsINT, 0)
}

func (s *Expr_aritContext) DDPUNTO() antlr.TerminalNode {
	return s.GetToken(ChemsDDPUNTO, 0)
}

func (s *Expr_aritContext) POW() antlr.TerminalNode {
	return s.GetToken(ChemsPOW, 0)
}

func (s *Expr_aritContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsPARIZQ, 0)
}

func (s *Expr_aritContext) COMA() antlr.TerminalNode {
	return s.GetToken(ChemsCOMA, 0)
}

func (s *Expr_aritContext) PARDER() antlr.TerminalNode {
	return s.GetToken(ChemsPARDER, 0)
}

func (s *Expr_aritContext) AllExpr_arit() []IExpr_aritContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem())
	var tst = make([]IExpr_aritContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpr_aritContext)
		}
	}

	return tst
}

func (s *Expr_aritContext) Expr_arit(i int) IExpr_aritContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpr_aritContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpr_aritContext)
}

func (s *Expr_aritContext) FLOAT() antlr.TerminalNode {
	return s.GetToken(ChemsFLOAT, 0)
}

func (s *Expr_aritContext) POWF() antlr.TerminalNode {
	return s.GetToken(ChemsPOWF, 0)
}

func (s *Expr_aritContext) DIFERENTE() antlr.TerminalNode {
	return s.GetToken(ChemsDIFERENTE, 0)
}

func (s *Expr_aritContext) SUB() antlr.TerminalNode {
	return s.GetToken(ChemsSUB, 0)
}

func (s *Expr_aritContext) Primitivo() IPrimitivoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IPrimitivoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IPrimitivoContext)
}

func (s *Expr_aritContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Expr_aritContext) MUL() antlr.TerminalNode {
	return s.GetToken(ChemsMUL, 0)
}

func (s *Expr_aritContext) DIV() antlr.TerminalNode {
	return s.GetToken(ChemsDIV, 0)
}

func (s *Expr_aritContext) MOD() antlr.TerminalNode {
	return s.GetToken(ChemsMOD, 0)
}

func (s *Expr_aritContext) ADD() antlr.TerminalNode {
	return s.GetToken(ChemsADD, 0)
}

func (s *Expr_aritContext) D_IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsD_IGUAL, 0)
}

func (s *Expr_aritContext) NOT_E() antlr.TerminalNode {
	return s.GetToken(ChemsNOT_E, 0)
}

func (s *Expr_aritContext) MENOR() antlr.TerminalNode {
	return s.GetToken(ChemsMENOR, 0)
}

func (s *Expr_aritContext) MENORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsMENORIGUAL, 0)
}

func (s *Expr_aritContext) MAYORIGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsMAYORIGUAL, 0)
}

func (s *Expr_aritContext) MAYOR() antlr.TerminalNode {
	return s.GetToken(ChemsMAYOR, 0)
}

func (s *Expr_aritContext) OR() antlr.TerminalNode {
	return s.GetToken(ChemsOR, 0)
}

func (s *Expr_aritContext) AND() antlr.TerminalNode {
	return s.GetToken(ChemsAND, 0)
}

func (s *Expr_aritContext) P_AS() antlr.TerminalNode {
	return s.GetToken(ChemsP_AS, 0)
}

func (s *Expr_aritContext) Tipo_d() ITipo_dContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ITipo_dContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ITipo_dContext)
}

func (s *Expr_aritContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(ChemsPUNTO, 0)
}

func (s *Expr_aritContext) T_STRING() antlr.TerminalNode {
	return s.GetToken(ChemsT_STRING, 0)
}

func (s *Expr_aritContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Expr_aritContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Expr_aritContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterExpr_arit(s)
	}
}

func (s *Expr_aritContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitExpr_arit(s)
	}
}

func (p *Chems) Expr_arit() (localctx IExpr_aritContext) {
	return p.expr_arit(0)
}

func (p *Chems) expr_arit(_p int) (localctx IExpr_aritContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewExpr_aritContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IExpr_aritContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, ChemsRULE_expr_arit, _p)
	var _la int

	defer func() {
		p.UnrollRecursionContexts(_parentctx)
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	var _alt int

	p.EnterOuterAlt(localctx, 1)
	p.SetState(230)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsINT:
		{
			p.SetState(194)
			p.Match(ChemsINT)
		}
		{
			p.SetState(195)
			p.Match(ChemsDDPUNTO)
		}
		{
			p.SetState(196)
			p.Match(ChemsPOW)
		}
		{
			p.SetState(197)

			var _m = p.Match(ChemsPARIZQ)

			localctx.(*Expr_aritContext)._PARIZQ = _m
		}
		{
			p.SetState(198)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opIz = _x
		}
		{
			p.SetState(199)
			p.Match(ChemsCOMA)
		}
		{
			p.SetState(200)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opDe = _x
		}
		{
			p.SetState(201)
			p.Match(ChemsPARDER)
		}
		localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), "^", localctx.(*Expr_aritContext).GetOpDe().GetP(), false, localctx.(*Expr_aritContext).Get_PARIZQ().GetLine(), localctx.(*Expr_aritContext).Get_PARIZQ().GetColumn())

	case ChemsFLOAT:
		{
			p.SetState(204)
			p.Match(ChemsFLOAT)
		}
		{
			p.SetState(205)
			p.Match(ChemsDDPUNTO)
		}
		{
			p.SetState(206)
			p.Match(ChemsPOWF)
		}
		{
			p.SetState(207)

			var _m = p.Match(ChemsPARIZQ)

			localctx.(*Expr_aritContext)._PARIZQ = _m
		}
		{
			p.SetState(208)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opIz = _x
		}
		{
			p.SetState(209)
			p.Match(ChemsCOMA)
		}
		{
			p.SetState(210)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opDe = _x
		}
		{
			p.SetState(211)
			p.Match(ChemsPARDER)
		}
		localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), "^", localctx.(*Expr_aritContext).GetOpDe().GetP(), false, localctx.(*Expr_aritContext).Get_PARIZQ().GetLine(), localctx.(*Expr_aritContext).Get_PARIZQ().GetColumn())

	case ChemsDIFERENTE:
		{
			p.SetState(214)

			var _m = p.Match(ChemsDIFERENTE)

			localctx.(*Expr_aritContext)._DIFERENTE = _m
		}
		{
			p.SetState(215)

			var _x = p.expr_arit(8)

			localctx.(*Expr_aritContext).opIz = _x
		}
		localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), "!", localctx.(*Expr_aritContext).GetOpIz().GetP(), true, localctx.(*Expr_aritContext).Get_DIFERENTE().GetLine(), localctx.(*Expr_aritContext).Get_DIFERENTE().GetColumn())

	case ChemsSUB:
		{
			p.SetState(218)

			var _m = p.Match(ChemsSUB)

			localctx.(*Expr_aritContext)._SUB = _m
		}
		{
			p.SetState(219)

			var _x = p.expr_arit(7)

			localctx.(*Expr_aritContext).opIz = _x
		}
		localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), "°", localctx.(*Expr_aritContext).GetOpIz().GetP(), true, localctx.(*Expr_aritContext).Get_SUB().GetLine(), localctx.(*Expr_aritContext).Get_SUB().GetColumn())

	case ChemsTRUE, ChemsFALSE, ChemsNUMBER, ChemsDECIMAL, ChemsSTRING, ChemsID, ChemsCARACTER:
		{
			p.SetState(222)

			var _x = p.Primitivo()

			localctx.(*Expr_aritContext)._primitivo = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitivo().GetP()

	case ChemsPARIZQ:
		{
			p.SetState(225)

			var _m = p.Match(ChemsPARIZQ)

			localctx.(*Expr_aritContext)._PARIZQ = _m
		}
		{
			p.SetState(226)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(227)
			p.Match(ChemsPARDER)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(265)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(263)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(232)

				if !(p.Precpred(p.GetParserRuleContext(), 6)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 6)", ""))
				}
				{
					p.SetState(233)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsMUL)|(1<<ChemsDIV)|(1<<ChemsMOD))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(234)

					var _x = p.expr_arit(7)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetP(), false, localctx.(*Expr_aritContext).GetOp().GetLine(), localctx.(*Expr_aritContext).GetOp().GetColumn())

			case 2:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(237)

				if !(p.Precpred(p.GetParserRuleContext(), 5)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 5)", ""))
				}
				{
					p.SetState(238)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsADD || _la == ChemsSUB) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(239)

					var _x = p.expr_arit(6)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetP(), false, localctx.(*Expr_aritContext).GetOp().GetLine(), localctx.(*Expr_aritContext).GetOp().GetColumn())

			case 3:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(242)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(243)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(((_la)&-(0x1f+1)) == 0 && ((1<<uint(_la))&((1<<ChemsMENOR)|(1<<ChemsMAYORIGUAL)|(1<<ChemsMENORIGUAL)|(1<<ChemsD_IGUAL)|(1<<ChemsNOT_E)|(1<<ChemsMAYOR))) != 0) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(244)

					var _x = p.expr_arit(5)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetP(), false, localctx.(*Expr_aritContext).GetOp().GetLine(), localctx.(*Expr_aritContext).GetOp().GetColumn())

			case 4:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(247)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(248)

					var _lt = p.GetTokenStream().LT(1)

					localctx.(*Expr_aritContext).op = _lt

					_la = p.GetTokenStream().LA(1)

					if !(_la == ChemsOR || _la == ChemsAND) {
						var _ri = p.GetErrorHandler().RecoverInline(p)

						localctx.(*Expr_aritContext).op = _ri
					} else {
						p.GetErrorHandler().ReportMatch(p)
						p.Consume()
					}
				}
				{
					p.SetState(249)

					var _x = p.expr_arit(4)

					localctx.(*Expr_aritContext).opDe = _x
				}
				localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), (func() string {
					if localctx.(*Expr_aritContext).GetOp() == nil {
						return ""
					} else {
						return localctx.(*Expr_aritContext).GetOp().GetText()
					}
				}()), localctx.(*Expr_aritContext).GetOpDe().GetP(), false, localctx.(*Expr_aritContext).GetOp().GetLine(), localctx.(*Expr_aritContext).GetOp().GetColumn())

			case 5:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).exp = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(252)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(253)
					p.Match(ChemsP_AS)
				}
				{
					p.SetState(254)

					var _x = p.Tipo_d()

					localctx.(*Expr_aritContext)._tipo_d = _x
				}
				localctx.(*Expr_aritContext).p = expresion.NewCast(localctx.(*Expr_aritContext).GetExp().GetP(), localctx.(*Expr_aritContext).Get_tipo_d().GetT())

			case 6:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).exp = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(257)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(258)

					var _m = p.Match(ChemsPUNTO)

					localctx.(*Expr_aritContext)._PUNTO = _m
				}
				{
					p.SetState(259)
					p.Match(ChemsT_STRING)
				}
				{
					p.SetState(260)

					var _m = p.Match(ChemsPARIZQ)

					localctx.(*Expr_aritContext)._PARIZQ = _m
				}
				{
					p.SetState(261)
					p.Match(ChemsPARDER)
				}
				localctx.(*Expr_aritContext).p = expresion.NewToString(localctx.(*Expr_aritContext).GetExp().GetP(), localctx.(*Expr_aritContext).Get_PUNTO().GetLine(), localctx.(*Expr_aritContext).Get_PUNTO().GetColumn())

			}

		}
		p.SetState(267)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IPrimitivoContext is an interface to support dynamic dispatch.
type IPrimitivoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_NUMBER returns the _NUMBER token.
	Get_NUMBER() antlr.Token

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Get_TRUE returns the _TRUE token.
	Get_TRUE() antlr.Token

	// Get_FALSE returns the _FALSE token.
	Get_FALSE() antlr.Token

	// Get_DECIMAL returns the _DECIMAL token.
	Get_DECIMAL() antlr.Token

	// Get_CARACTER returns the _CARACTER token.
	Get_CARACTER() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// Set_NUMBER sets the _NUMBER token.
	Set_NUMBER(antlr.Token)

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// Set_TRUE sets the _TRUE token.
	Set_TRUE(antlr.Token)

	// Set_FALSE sets the _FALSE token.
	Set_FALSE(antlr.Token)

	// Set_DECIMAL sets the _DECIMAL token.
	Set_DECIMAL(antlr.Token)

	// Set_CARACTER sets the _CARACTER token.
	Set_CARACTER(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsPrimitivoContext differentiates from other interfaces.
	IsPrimitivoContext()
}

type PrimitivoContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	p         interfaces.Expresion
	_NUMBER   antlr.Token
	_STRING   antlr.Token
	_TRUE     antlr.Token
	_FALSE    antlr.Token
	_DECIMAL  antlr.Token
	_CARACTER antlr.Token
	id        antlr.Token
}

func NewEmptyPrimitivoContext() *PrimitivoContext {
	var p = new(PrimitivoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_primitivo
	return p
}

func (*PrimitivoContext) IsPrimitivoContext() {}

func NewPrimitivoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *PrimitivoContext {
	var p = new(PrimitivoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_primitivo

	return p
}

func (s *PrimitivoContext) GetParser() antlr.Parser { return s.parser }

func (s *PrimitivoContext) Get_NUMBER() antlr.Token { return s._NUMBER }

func (s *PrimitivoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *PrimitivoContext) Get_TRUE() antlr.Token { return s._TRUE }

func (s *PrimitivoContext) Get_FALSE() antlr.Token { return s._FALSE }

func (s *PrimitivoContext) Get_DECIMAL() antlr.Token { return s._DECIMAL }

func (s *PrimitivoContext) Get_CARACTER() antlr.Token { return s._CARACTER }

func (s *PrimitivoContext) GetId() antlr.Token { return s.id }

func (s *PrimitivoContext) Set_NUMBER(v antlr.Token) { s._NUMBER = v }

func (s *PrimitivoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *PrimitivoContext) Set_TRUE(v antlr.Token) { s._TRUE = v }

func (s *PrimitivoContext) Set_FALSE(v antlr.Token) { s._FALSE = v }

func (s *PrimitivoContext) Set_DECIMAL(v antlr.Token) { s._DECIMAL = v }

func (s *PrimitivoContext) Set_CARACTER(v antlr.Token) { s._CARACTER = v }

func (s *PrimitivoContext) SetId(v antlr.Token) { s.id = v }

func (s *PrimitivoContext) GetP() interfaces.Expresion { return s.p }

func (s *PrimitivoContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *PrimitivoContext) NUMBER() antlr.TerminalNode {
	return s.GetToken(ChemsNUMBER, 0)
}

func (s *PrimitivoContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *PrimitivoContext) TRUE() antlr.TerminalNode {
	return s.GetToken(ChemsTRUE, 0)
}

func (s *PrimitivoContext) FALSE() antlr.TerminalNode {
	return s.GetToken(ChemsFALSE, 0)
}

func (s *PrimitivoContext) DECIMAL() antlr.TerminalNode {
	return s.GetToken(ChemsDECIMAL, 0)
}

func (s *PrimitivoContext) CARACTER() antlr.TerminalNode {
	return s.GetToken(ChemsCARACTER, 0)
}

func (s *PrimitivoContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *PrimitivoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *PrimitivoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *PrimitivoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterPrimitivo(s)
	}
}

func (s *PrimitivoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitPrimitivo(s)
	}
}

func (p *Chems) Primitivo() (localctx IPrimitivoContext) {
	this := p
	_ = this

	localctx = NewPrimitivoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ChemsRULE_primitivo)

	defer func() {
		p.ExitRule()
	}()

	defer func() {
		if err := recover(); err != nil {
			if v, ok := err.(antlr.RecognitionException); ok {
				localctx.SetException(v)
				p.GetErrorHandler().ReportError(p, v)
				p.GetErrorHandler().Recover(p, v)
			} else {
				panic(err)
			}
		}
	}()

	p.SetState(282)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(268)

			var _m = p.Match(ChemsNUMBER)

			localctx.(*PrimitivoContext)._NUMBER = _m
		}

		num, err := strconv.Atoi((func() string {
			if localctx.(*PrimitivoContext).Get_NUMBER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_NUMBER().GetText()
			}
		}()))
		if err != nil {
			fmt.Println(err)
		}
		linea := localctx.(*PrimitivoContext)._NUMBER.GetLine()
		col := localctx.(*PrimitivoContext)._NUMBER.GetColumn()
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(num, interfaces.INTEGER, col, linea)

	case ChemsSTRING:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(270)

			var _m = p.Match(ChemsSTRING)

			localctx.(*PrimitivoContext)._STRING = _m
		}

		linea := localctx.(*PrimitivoContext)._STRING.GetLine()
		col := localctx.(*PrimitivoContext)._STRING.GetColumn()
		str := (func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(str, interfaces.STRING, col, linea)

	case ChemsTRUE:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(272)

			var _m = p.Match(ChemsTRUE)

			localctx.(*PrimitivoContext)._TRUE = _m
		}

		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(1, interfaces.BOOLEAN, localctx.(*PrimitivoContext).Get_TRUE().GetColumn(), localctx.(*PrimitivoContext).Get_TRUE().GetLine())

	case ChemsFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(274)

			var _m = p.Match(ChemsFALSE)

			localctx.(*PrimitivoContext)._FALSE = _m
		}

		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(0, interfaces.BOOLEAN, localctx.(*PrimitivoContext).Get_FALSE().GetColumn(), localctx.(*PrimitivoContext).Get_FALSE().GetLine())

	case ChemsDECIMAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(276)

			var _m = p.Match(ChemsDECIMAL)

			localctx.(*PrimitivoContext)._DECIMAL = _m
		}

		linea := localctx.(*PrimitivoContext)._DECIMAL.GetLine()
		col := localctx.(*PrimitivoContext)._DECIMAL.GetColumn()
		num, err := strconv.ParseFloat((func() string {
			if localctx.(*PrimitivoContext).Get_DECIMAL() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_DECIMAL().GetText()
			}
		}()), 64)
		if err != nil {
			fmt.Println(err)
		}
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(num, interfaces.FLOAT, col, linea)

	case ChemsCARACTER:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(278)

			var _m = p.Match(ChemsCARACTER)

			localctx.(*PrimitivoContext)._CARACTER = _m
		}

		linea := localctx.(*PrimitivoContext)._CARACTER.GetLine()
		col := localctx.(*PrimitivoContext)._CARACTER.GetColumn()
		str := (func() string {
			if localctx.(*PrimitivoContext).Get_CARACTER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CARACTER().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*PrimitivoContext).Get_CARACTER() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).Get_CARACTER().GetText()
			}
		}()))-1]
		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(str, interfaces.CHAR, col, linea)

	case ChemsID:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(280)

			var _m = p.Match(ChemsID)

			localctx.(*PrimitivoContext).id = _m
		}

		linea := localctx.(*PrimitivoContext).id.GetLine()
		col := localctx.(*PrimitivoContext).id.GetColumn()
		localctx.(*PrimitivoContext).p = expresion.NewCallVariable((func() string {
			if localctx.(*PrimitivoContext).GetId() == nil {
				return ""
			} else {
				return localctx.(*PrimitivoContext).GetId().GetText()
			}
		}()), linea, col)

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 6)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 5)

	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	case 4:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 5:
		return p.Precpred(p.GetParserRuleContext(), 9)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
