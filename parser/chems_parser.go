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
	3, 24715, 42794, 33075, 47597, 16764, 15335, 30598, 22884, 3, 88, 614,
	4, 2, 9, 2, 4, 3, 9, 3, 4, 4, 9, 4, 4, 5, 9, 5, 4, 6, 9, 6, 4, 7, 9, 7,
	4, 8, 9, 8, 4, 9, 9, 9, 4, 10, 9, 10, 4, 11, 9, 11, 4, 12, 9, 12, 4, 13,
	9, 13, 4, 14, 9, 14, 4, 15, 9, 15, 4, 16, 9, 16, 4, 17, 9, 17, 4, 18, 9,
	18, 4, 19, 9, 19, 4, 20, 9, 20, 4, 21, 9, 21, 4, 22, 9, 22, 4, 23, 9, 23,
	4, 24, 9, 24, 4, 25, 9, 25, 4, 26, 9, 26, 4, 27, 9, 27, 4, 28, 9, 28, 4,
	29, 9, 29, 4, 30, 9, 30, 4, 31, 9, 31, 4, 32, 9, 32, 4, 33, 9, 33, 4, 34,
	9, 34, 4, 35, 9, 35, 4, 36, 9, 36, 4, 37, 9, 37, 4, 38, 9, 38, 3, 2, 3,
	2, 3, 2, 3, 3, 7, 3, 81, 10, 3, 12, 3, 14, 3, 84, 11, 3, 3, 3, 3, 3, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3, 4, 3,
	4, 3, 4, 5, 4, 138, 10, 4, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3,
	5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 3, 5, 5, 5, 182,
	10, 5, 3, 6, 3, 6, 3, 6, 3, 6, 3, 6, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7, 3, 7,
	3, 7, 3, 7, 3, 8, 3, 8, 3, 8, 5, 8, 200, 10, 8, 3, 9, 3, 9, 3, 9, 3, 9,
	3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 5, 9, 213, 10, 9, 3, 10, 3, 10,
	3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3, 10, 3,
	10, 3, 10, 3, 10, 3, 10, 5, 10, 231, 10, 10, 3, 11, 3, 11, 3, 11, 3, 11,
	3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3, 12, 3,
	12, 3, 12, 3, 12, 3, 12, 5, 12, 251, 10, 12, 3, 13, 3, 13, 3, 13, 3, 13,
	3, 13, 3, 13, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3,
	14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14, 3, 14,
	3, 14, 3, 14, 3, 14, 5, 14, 281, 10, 14, 3, 15, 3, 15, 3, 15, 3, 15, 3,
	15, 3, 15, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 16, 3, 17,
	3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 5, 17, 304, 10, 17, 3, 17, 3,
	17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 3, 17, 7, 17, 314, 10, 17, 12, 17,
	14, 17, 317, 11, 17, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3,
	18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 3, 18, 5, 18, 335,
	10, 18, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 3, 19, 5, 19,
	345, 10, 19, 3, 20, 3, 20, 3, 20, 3, 20, 3, 21, 3, 21, 3, 21, 3, 21, 3,
	21, 3, 21, 3, 21, 3, 22, 3, 22, 3, 22, 3, 22, 5, 22, 362, 10, 22, 3, 23,
	3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 23, 3, 24, 3, 24, 3, 24, 3, 24, 3,
	24, 3, 24, 3, 24, 5, 24, 378, 10, 24, 3, 24, 3, 24, 3, 24, 3, 24, 3, 24,
	3, 24, 3, 24, 3, 24, 7, 24, 388, 10, 24, 12, 24, 14, 24, 391, 11, 24, 3,
	25, 3, 25, 3, 25, 3, 25, 3, 25, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26, 3, 26,
	3, 26, 3, 26, 3, 26, 7, 26, 407, 10, 26, 12, 26, 14, 26, 410, 11, 26, 3,
	27, 3, 27, 3, 27, 3, 27, 3, 27, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28,
	3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 3, 28, 5, 28, 430, 10, 28, 3,
	29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 29, 3, 30, 3, 30,
	3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 3, 30, 5, 30, 450, 10, 30, 3,
	31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 31, 3, 32, 3, 32, 3, 32, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 5, 33, 518, 10, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3,
	33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 3, 33,
	3, 33, 3, 33, 3, 33, 3, 33, 3, 33, 7, 33, 551, 10, 33, 12, 33, 14, 33,
	554, 11, 33, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 3,
	34, 3, 34, 3, 34, 3, 34, 3, 34, 3, 34, 5, 34, 570, 10, 34, 3, 35, 3, 35,
	3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 3, 35, 7, 35, 581, 10, 35, 12,
	35, 14, 35, 584, 11, 35, 3, 36, 3, 36, 3, 36, 3, 36, 3, 37, 3, 37, 3, 37,
	3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 3, 37, 7, 37, 602,
	10, 37, 12, 37, 14, 37, 605, 11, 37, 3, 38, 3, 38, 3, 38, 3, 38, 3, 38,
	3, 38, 3, 38, 3, 38, 2, 8, 32, 46, 50, 64, 68, 72, 39, 2, 4, 6, 8, 10,
	12, 14, 16, 18, 20, 22, 24, 26, 28, 30, 32, 34, 36, 38, 40, 42, 44, 46,
	48, 50, 52, 54, 56, 58, 60, 62, 64, 66, 68, 70, 72, 74, 2, 6, 4, 2, 20,
	21, 24, 24, 3, 2, 22, 23, 4, 2, 11, 13, 15, 17, 3, 2, 18, 19, 2, 656, 2,
	76, 3, 2, 2, 2, 4, 82, 3, 2, 2, 2, 6, 137, 3, 2, 2, 2, 8, 181, 3, 2, 2,
	2, 10, 183, 3, 2, 2, 2, 12, 188, 3, 2, 2, 2, 14, 199, 3, 2, 2, 2, 16, 212,
	3, 2, 2, 2, 18, 230, 3, 2, 2, 2, 20, 232, 3, 2, 2, 2, 22, 250, 3, 2, 2,
	2, 24, 252, 3, 2, 2, 2, 26, 280, 3, 2, 2, 2, 28, 282, 3, 2, 2, 2, 30, 288,
	3, 2, 2, 2, 32, 303, 3, 2, 2, 2, 34, 334, 3, 2, 2, 2, 36, 344, 3, 2, 2,
	2, 38, 346, 3, 2, 2, 2, 40, 350, 3, 2, 2, 2, 42, 361, 3, 2, 2, 2, 44, 363,
	3, 2, 2, 2, 46, 377, 3, 2, 2, 2, 48, 392, 3, 2, 2, 2, 50, 397, 3, 2, 2,
	2, 52, 411, 3, 2, 2, 2, 54, 429, 3, 2, 2, 2, 56, 431, 3, 2, 2, 2, 58, 449,
	3, 2, 2, 2, 60, 451, 3, 2, 2, 2, 62, 457, 3, 2, 2, 2, 64, 517, 3, 2, 2,
	2, 66, 569, 3, 2, 2, 2, 68, 571, 3, 2, 2, 2, 70, 585, 3, 2, 2, 2, 72, 589,
	3, 2, 2, 2, 74, 606, 3, 2, 2, 2, 76, 77, 5, 4, 3, 2, 77, 78, 8, 2, 1, 2,
	78, 3, 3, 2, 2, 2, 79, 81, 5, 6, 4, 2, 80, 79, 3, 2, 2, 2, 81, 84, 3, 2,
	2, 2, 82, 80, 3, 2, 2, 2, 82, 83, 3, 2, 2, 2, 83, 85, 3, 2, 2, 2, 84, 82,
	3, 2, 2, 2, 85, 86, 8, 3, 1, 2, 86, 5, 3, 2, 2, 2, 87, 88, 7, 35, 2, 2,
	88, 89, 7, 3, 2, 2, 89, 90, 7, 36, 2, 2, 90, 91, 7, 26, 2, 2, 91, 92, 5,
	62, 32, 2, 92, 93, 7, 27, 2, 2, 93, 94, 7, 4, 2, 2, 94, 95, 8, 4, 1, 2,
	95, 138, 3, 2, 2, 2, 96, 97, 5, 12, 7, 2, 97, 98, 7, 4, 2, 2, 98, 99, 8,
	4, 1, 2, 99, 138, 3, 2, 2, 2, 100, 101, 5, 10, 6, 2, 101, 102, 7, 4, 2,
	2, 102, 103, 8, 4, 1, 2, 103, 138, 3, 2, 2, 2, 104, 105, 7, 42, 2, 2, 105,
	106, 5, 62, 32, 2, 106, 107, 7, 28, 2, 2, 107, 108, 5, 4, 3, 2, 108, 109,
	7, 29, 2, 2, 109, 110, 8, 4, 1, 2, 110, 138, 3, 2, 2, 2, 111, 112, 5, 36,
	19, 2, 112, 113, 8, 4, 1, 2, 113, 138, 3, 2, 2, 2, 114, 115, 5, 38, 20,
	2, 115, 116, 8, 4, 1, 2, 116, 138, 3, 2, 2, 2, 117, 118, 5, 30, 16, 2,
	118, 119, 8, 4, 1, 2, 119, 138, 3, 2, 2, 2, 120, 121, 5, 28, 15, 2, 121,
	122, 8, 4, 1, 2, 122, 138, 3, 2, 2, 2, 123, 124, 5, 40, 21, 2, 124, 125,
	7, 4, 2, 2, 125, 126, 8, 4, 1, 2, 126, 138, 3, 2, 2, 2, 127, 128, 5, 44,
	23, 2, 128, 129, 8, 4, 1, 2, 129, 138, 3, 2, 2, 2, 130, 131, 5, 56, 29,
	2, 131, 132, 8, 4, 1, 2, 132, 138, 3, 2, 2, 2, 133, 134, 5, 60, 31, 2,
	134, 135, 7, 4, 2, 2, 135, 136, 8, 4, 1, 2, 136, 138, 3, 2, 2, 2, 137,
	87, 3, 2, 2, 2, 137, 96, 3, 2, 2, 2, 137, 100, 3, 2, 2, 2, 137, 104, 3,
	2, 2, 2, 137, 111, 3, 2, 2, 2, 137, 114, 3, 2, 2, 2, 137, 117, 3, 2, 2,
	2, 137, 120, 3, 2, 2, 2, 137, 123, 3, 2, 2, 2, 137, 127, 3, 2, 2, 2, 137,
	130, 3, 2, 2, 2, 137, 133, 3, 2, 2, 2, 138, 7, 3, 2, 2, 2, 139, 140, 7,
	35, 2, 2, 140, 141, 7, 3, 2, 2, 141, 142, 7, 36, 2, 2, 142, 143, 7, 26,
	2, 2, 143, 144, 5, 62, 32, 2, 144, 145, 7, 27, 2, 2, 145, 146, 8, 5, 1,
	2, 146, 182, 3, 2, 2, 2, 147, 148, 5, 12, 7, 2, 148, 149, 8, 5, 1, 2, 149,
	182, 3, 2, 2, 2, 150, 151, 5, 10, 6, 2, 151, 152, 8, 5, 1, 2, 152, 182,
	3, 2, 2, 2, 153, 154, 7, 42, 2, 2, 154, 155, 5, 62, 32, 2, 155, 156, 7,
	28, 2, 2, 156, 157, 5, 4, 3, 2, 157, 158, 7, 29, 2, 2, 158, 159, 8, 5,
	1, 2, 159, 182, 3, 2, 2, 2, 160, 161, 5, 36, 19, 2, 161, 162, 8, 5, 1,
	2, 162, 182, 3, 2, 2, 2, 163, 164, 5, 38, 20, 2, 164, 165, 8, 5, 1, 2,
	165, 182, 3, 2, 2, 2, 166, 167, 5, 30, 16, 2, 167, 168, 8, 5, 1, 2, 168,
	182, 3, 2, 2, 2, 169, 170, 5, 28, 15, 2, 170, 171, 8, 5, 1, 2, 171, 182,
	3, 2, 2, 2, 172, 173, 5, 40, 21, 2, 173, 174, 8, 5, 1, 2, 174, 182, 3,
	2, 2, 2, 175, 176, 5, 56, 29, 2, 176, 177, 8, 5, 1, 2, 177, 182, 3, 2,
	2, 2, 178, 179, 5, 60, 31, 2, 179, 180, 8, 5, 1, 2, 180, 182, 3, 2, 2,
	2, 181, 139, 3, 2, 2, 2, 181, 147, 3, 2, 2, 2, 181, 150, 3, 2, 2, 2, 181,
	153, 3, 2, 2, 2, 181, 160, 3, 2, 2, 2, 181, 163, 3, 2, 2, 2, 181, 166,
	3, 2, 2, 2, 181, 169, 3, 2, 2, 2, 181, 172, 3, 2, 2, 2, 181, 175, 3, 2,
	2, 2, 181, 178, 3, 2, 2, 2, 182, 9, 3, 2, 2, 2, 183, 184, 7, 86, 2, 2,
	184, 185, 7, 9, 2, 2, 185, 186, 5, 62, 32, 2, 186, 187, 8, 6, 1, 2, 187,
	11, 3, 2, 2, 2, 188, 189, 7, 51, 2, 2, 189, 190, 5, 14, 8, 2, 190, 191,
	7, 86, 2, 2, 191, 192, 5, 16, 9, 2, 192, 193, 7, 9, 2, 2, 193, 194, 5,
	62, 32, 2, 194, 195, 8, 7, 1, 2, 195, 13, 3, 2, 2, 2, 196, 197, 7, 58,
	2, 2, 197, 200, 8, 8, 1, 2, 198, 200, 8, 8, 1, 2, 199, 196, 3, 2, 2, 2,
	199, 198, 3, 2, 2, 2, 200, 15, 3, 2, 2, 2, 201, 202, 7, 8, 2, 2, 202, 203,
	5, 24, 13, 2, 203, 204, 8, 9, 1, 2, 204, 213, 3, 2, 2, 2, 205, 206, 5,
	20, 11, 2, 206, 207, 8, 9, 1, 2, 207, 213, 3, 2, 2, 2, 208, 209, 7, 8,
	2, 2, 209, 210, 5, 18, 10, 2, 210, 211, 8, 9, 1, 2, 211, 213, 3, 2, 2,
	2, 212, 201, 3, 2, 2, 2, 212, 205, 3, 2, 2, 2, 212, 208, 3, 2, 2, 2, 213,
	17, 3, 2, 2, 2, 214, 215, 7, 52, 2, 2, 215, 231, 8, 10, 1, 2, 216, 217,
	7, 53, 2, 2, 217, 231, 8, 10, 1, 2, 218, 219, 7, 54, 2, 2, 219, 231, 8,
	10, 1, 2, 220, 221, 7, 55, 2, 2, 221, 231, 8, 10, 1, 2, 222, 223, 7, 56,
	2, 2, 223, 231, 8, 10, 1, 2, 224, 225, 7, 39, 2, 2, 225, 231, 8, 10, 1,
	2, 226, 227, 7, 57, 2, 2, 227, 231, 8, 10, 1, 2, 228, 229, 7, 86, 2, 2,
	229, 231, 8, 10, 1, 2, 230, 214, 3, 2, 2, 2, 230, 216, 3, 2, 2, 2, 230,
	218, 3, 2, 2, 2, 230, 220, 3, 2, 2, 2, 230, 222, 3, 2, 2, 2, 230, 224,
	3, 2, 2, 2, 230, 226, 3, 2, 2, 2, 230, 228, 3, 2, 2, 2, 231, 19, 3, 2,
	2, 2, 232, 233, 7, 8, 2, 2, 233, 234, 5, 22, 12, 2, 234, 235, 8, 11, 1,
	2, 235, 21, 3, 2, 2, 2, 236, 237, 7, 30, 2, 2, 237, 238, 5, 18, 10, 2,
	238, 239, 7, 4, 2, 2, 239, 240, 5, 62, 32, 2, 240, 241, 7, 31, 2, 2, 241,
	242, 8, 12, 1, 2, 242, 251, 3, 2, 2, 2, 243, 244, 7, 30, 2, 2, 244, 245,
	5, 22, 12, 2, 245, 246, 7, 4, 2, 2, 246, 247, 5, 62, 32, 2, 247, 248, 7,
	31, 2, 2, 248, 249, 8, 12, 1, 2, 249, 251, 3, 2, 2, 2, 250, 236, 3, 2,
	2, 2, 250, 243, 3, 2, 2, 2, 251, 23, 3, 2, 2, 2, 252, 253, 7, 79, 2, 2,
	253, 254, 7, 11, 2, 2, 254, 255, 5, 26, 14, 2, 255, 256, 7, 17, 2, 2, 256,
	257, 8, 13, 1, 2, 257, 25, 3, 2, 2, 2, 258, 259, 7, 52, 2, 2, 259, 281,
	8, 14, 1, 2, 260, 261, 7, 53, 2, 2, 261, 281, 8, 14, 1, 2, 262, 263, 7,
	54, 2, 2, 263, 281, 8, 14, 1, 2, 264, 265, 7, 55, 2, 2, 265, 281, 8, 14,
	1, 2, 266, 267, 7, 56, 2, 2, 267, 281, 8, 14, 1, 2, 268, 269, 7, 39, 2,
	2, 269, 281, 8, 14, 1, 2, 270, 271, 7, 57, 2, 2, 271, 281, 8, 14, 1, 2,
	272, 273, 7, 86, 2, 2, 273, 281, 8, 14, 1, 2, 274, 275, 7, 79, 2, 2, 275,
	276, 7, 11, 2, 2, 276, 277, 5, 26, 14, 2, 277, 278, 7, 17, 2, 2, 278, 279,
	8, 14, 1, 2, 279, 281, 3, 2, 2, 2, 280, 258, 3, 2, 2, 2, 280, 260, 3, 2,
	2, 2, 280, 262, 3, 2, 2, 2, 280, 264, 3, 2, 2, 2, 280, 266, 3, 2, 2, 2,
	280, 268, 3, 2, 2, 2, 280, 270, 3, 2, 2, 2, 280, 272, 3, 2, 2, 2, 280,
	274, 3, 2, 2, 2, 281, 27, 3, 2, 2, 2, 282, 283, 7, 66, 2, 2, 283, 284,
	7, 28, 2, 2, 284, 285, 5, 4, 3, 2, 285, 286, 7, 29, 2, 2, 286, 287, 8,
	15, 1, 2, 287, 29, 3, 2, 2, 2, 288, 289, 7, 40, 2, 2, 289, 290, 5, 62,
	32, 2, 290, 291, 7, 28, 2, 2, 291, 292, 5, 32, 17, 2, 292, 293, 7, 29,
	2, 2, 293, 294, 5, 34, 18, 2, 294, 295, 8, 16, 1, 2, 295, 31, 3, 2, 2,
	2, 296, 297, 8, 17, 1, 2, 297, 298, 5, 62, 32, 2, 298, 299, 8, 17, 1, 2,
	299, 304, 3, 2, 2, 2, 300, 301, 5, 6, 4, 2, 301, 302, 8, 17, 1, 2, 302,
	304, 3, 2, 2, 2, 303, 296, 3, 2, 2, 2, 303, 300, 3, 2, 2, 2, 304, 315,
	3, 2, 2, 2, 305, 306, 12, 4, 2, 2, 306, 307, 5, 62, 32, 2, 307, 308, 8,
	17, 1, 2, 308, 314, 3, 2, 2, 2, 309, 310, 12, 3, 2, 2, 310, 311, 5, 6,
	4, 2, 311, 312, 8, 17, 1, 2, 312, 314, 3, 2, 2, 2, 313, 305, 3, 2, 2, 2,
	313, 309, 3, 2, 2, 2, 314, 317, 3, 2, 2, 2, 315, 313, 3, 2, 2, 2, 315,
	316, 3, 2, 2, 2, 316, 33, 3, 2, 2, 2, 317, 315, 3, 2, 2, 2, 318, 319, 7,
	41, 2, 2, 319, 320, 7, 28, 2, 2, 320, 321, 5, 32, 17, 2, 321, 322, 7, 29,
	2, 2, 322, 323, 8, 18, 1, 2, 323, 335, 3, 2, 2, 2, 324, 325, 7, 41, 2,
	2, 325, 326, 7, 40, 2, 2, 326, 327, 5, 62, 32, 2, 327, 328, 7, 28, 2, 2,
	328, 329, 5, 32, 17, 2, 329, 330, 7, 29, 2, 2, 330, 331, 5, 34, 18, 2,
	331, 332, 8, 18, 1, 2, 332, 335, 3, 2, 2, 2, 333, 335, 8, 18, 1, 2, 334,
	318, 3, 2, 2, 2, 334, 324, 3, 2, 2, 2, 334, 333, 3, 2, 2, 2, 335, 35, 3,
	2, 2, 2, 336, 337, 7, 48, 2, 2, 337, 338, 5, 62, 32, 2, 338, 339, 7, 4,
	2, 2, 339, 340, 8, 19, 1, 2, 340, 345, 3, 2, 2, 2, 341, 342, 7, 48, 2,
	2, 342, 343, 7, 4, 2, 2, 343, 345, 8, 19, 1, 2, 344, 336, 3, 2, 2, 2, 344,
	341, 3, 2, 2, 2, 345, 37, 3, 2, 2, 2, 346, 347, 7, 49, 2, 2, 347, 348,
	7, 4, 2, 2, 348, 349, 8, 20, 1, 2, 349, 39, 3, 2, 2, 2, 350, 351, 7, 37,
	2, 2, 351, 352, 7, 26, 2, 2, 352, 353, 5, 42, 22, 2, 353, 354, 5, 68, 35,
	2, 354, 355, 7, 27, 2, 2, 355, 356, 8, 21, 1, 2, 356, 41, 3, 2, 2, 2, 357,
	358, 7, 85, 2, 2, 358, 359, 7, 5, 2, 2, 359, 362, 8, 22, 1, 2, 360, 362,
	8, 22, 1, 2, 361, 357, 3, 2, 2, 2, 361, 360, 3, 2, 2, 2, 362, 43, 3, 2,
	2, 2, 363, 364, 7, 65, 2, 2, 364, 365, 5, 62, 32, 2, 365, 366, 7, 28, 2,
	2, 366, 367, 5, 46, 24, 2, 367, 368, 7, 29, 2, 2, 368, 369, 8, 23, 1, 2,
	369, 45, 3, 2, 2, 2, 370, 371, 8, 24, 1, 2, 371, 372, 5, 48, 25, 2, 372,
	373, 8, 24, 1, 2, 373, 378, 3, 2, 2, 2, 374, 375, 5, 52, 27, 2, 375, 376,
	8, 24, 1, 2, 376, 378, 3, 2, 2, 2, 377, 370, 3, 2, 2, 2, 377, 374, 3, 2,
	2, 2, 378, 389, 3, 2, 2, 2, 379, 380, 12, 6, 2, 2, 380, 381, 5, 48, 25,
	2, 381, 382, 8, 24, 1, 2, 382, 388, 3, 2, 2, 2, 383, 384, 12, 5, 2, 2,
	384, 385, 5, 52, 27, 2, 385, 386, 8, 24, 1, 2, 386, 388, 3, 2, 2, 2, 387,
	379, 3, 2, 2, 2, 387, 383, 3, 2, 2, 2, 388, 391, 3, 2, 2, 2, 389, 387,
	3, 2, 2, 2, 389, 390, 3, 2, 2, 2, 390, 47, 3, 2, 2, 2, 391, 389, 3, 2,
	2, 2, 392, 393, 5, 50, 26, 2, 393, 394, 7, 14, 2, 2, 394, 395, 5, 54, 28,
	2, 395, 396, 8, 25, 1, 2, 396, 49, 3, 2, 2, 2, 397, 398, 8, 26, 1, 2, 398,
	399, 5, 62, 32, 2, 399, 400, 8, 26, 1, 2, 400, 408, 3, 2, 2, 2, 401, 402,
	12, 4, 2, 2, 402, 403, 7, 33, 2, 2, 403, 404, 5, 62, 32, 2, 404, 405, 8,
	26, 1, 2, 405, 407, 3, 2, 2, 2, 406, 401, 3, 2, 2, 2, 407, 410, 3, 2, 2,
	2, 408, 406, 3, 2, 2, 2, 408, 409, 3, 2, 2, 2, 409, 51, 3, 2, 2, 2, 410,
	408, 3, 2, 2, 2, 411, 412, 7, 77, 2, 2, 412, 413, 7, 14, 2, 2, 413, 414,
	5, 54, 28, 2, 414, 415, 8, 27, 1, 2, 415, 53, 3, 2, 2, 2, 416, 417, 5,
	62, 32, 2, 417, 418, 7, 5, 2, 2, 418, 419, 8, 28, 1, 2, 419, 430, 3, 2,
	2, 2, 420, 421, 7, 28, 2, 2, 421, 422, 5, 4, 3, 2, 422, 423, 7, 29, 2,
	2, 423, 424, 8, 28, 1, 2, 424, 430, 3, 2, 2, 2, 425, 426, 5, 8, 5, 2, 426,
	427, 7, 5, 2, 2, 427, 428, 8, 28, 1, 2, 428, 430, 3, 2, 2, 2, 429, 416,
	3, 2, 2, 2, 429, 420, 3, 2, 2, 2, 429, 425, 3, 2, 2, 2, 430, 55, 3, 2,
	2, 2, 431, 432, 7, 43, 2, 2, 432, 433, 7, 86, 2, 2, 433, 434, 7, 44, 2,
	2, 434, 435, 5, 58, 30, 2, 435, 436, 7, 28, 2, 2, 436, 437, 5, 4, 3, 2,
	437, 438, 7, 29, 2, 2, 438, 439, 8, 29, 1, 2, 439, 57, 3, 2, 2, 2, 440,
	441, 5, 62, 32, 2, 441, 442, 7, 3, 2, 2, 442, 443, 7, 3, 2, 2, 443, 444,
	5, 62, 32, 2, 444, 445, 8, 30, 1, 2, 445, 450, 3, 2, 2, 2, 446, 447, 5,
	62, 32, 2, 447, 448, 8, 30, 1, 2, 448, 450, 3, 2, 2, 2, 449, 440, 3, 2,
	2, 2, 449, 446, 3, 2, 2, 2, 450, 59, 3, 2, 2, 2, 451, 452, 7, 86, 2, 2,
	452, 453, 5, 72, 37, 2, 453, 454, 7, 9, 2, 2, 454, 455, 5, 62, 32, 2, 455,
	456, 8, 31, 1, 2, 456, 61, 3, 2, 2, 2, 457, 458, 5, 64, 33, 2, 458, 459,
	8, 32, 1, 2, 459, 63, 3, 2, 2, 2, 460, 461, 8, 33, 1, 2, 461, 462, 7, 52,
	2, 2, 462, 463, 7, 34, 2, 2, 463, 464, 7, 47, 2, 2, 464, 465, 7, 26, 2,
	2, 465, 466, 5, 64, 33, 2, 466, 467, 7, 5, 2, 2, 467, 468, 5, 64, 33, 2,
	468, 469, 7, 27, 2, 2, 469, 470, 8, 33, 1, 2, 470, 518, 3, 2, 2, 2, 471,
	472, 7, 53, 2, 2, 472, 473, 7, 34, 2, 2, 473, 474, 7, 46, 2, 2, 474, 475,
	7, 26, 2, 2, 475, 476, 5, 64, 33, 2, 476, 477, 7, 5, 2, 2, 477, 478, 5,
	64, 33, 2, 478, 479, 7, 27, 2, 2, 479, 480, 8, 33, 1, 2, 480, 518, 3, 2,
	2, 2, 481, 482, 7, 7, 2, 2, 482, 483, 5, 64, 33, 16, 483, 484, 8, 33, 1,
	2, 484, 518, 3, 2, 2, 2, 485, 486, 7, 23, 2, 2, 486, 487, 5, 64, 33, 15,
	487, 488, 8, 33, 1, 2, 488, 518, 3, 2, 2, 2, 489, 490, 5, 66, 34, 2, 490,
	491, 8, 33, 1, 2, 491, 518, 3, 2, 2, 2, 492, 493, 7, 26, 2, 2, 493, 494,
	5, 62, 32, 2, 494, 495, 7, 27, 2, 2, 495, 496, 8, 33, 1, 2, 496, 518, 3,
	2, 2, 2, 497, 498, 5, 28, 15, 2, 498, 499, 8, 33, 1, 2, 499, 518, 3, 2,
	2, 2, 500, 501, 5, 30, 16, 2, 501, 502, 8, 33, 1, 2, 502, 518, 3, 2, 2,
	2, 503, 504, 5, 44, 23, 2, 504, 505, 8, 33, 1, 2, 505, 518, 3, 2, 2, 2,
	506, 507, 7, 30, 2, 2, 507, 508, 5, 68, 35, 2, 508, 509, 7, 31, 2, 2, 509,
	510, 8, 33, 1, 2, 510, 518, 3, 2, 2, 2, 511, 512, 5, 70, 36, 2, 512, 513,
	8, 33, 1, 2, 513, 518, 3, 2, 2, 2, 514, 515, 5, 74, 38, 2, 515, 516, 8,
	33, 1, 2, 516, 518, 3, 2, 2, 2, 517, 460, 3, 2, 2, 2, 517, 471, 3, 2, 2,
	2, 517, 481, 3, 2, 2, 2, 517, 485, 3, 2, 2, 2, 517, 489, 3, 2, 2, 2, 517,
	492, 3, 2, 2, 2, 517, 497, 3, 2, 2, 2, 517, 500, 3, 2, 2, 2, 517, 503,
	3, 2, 2, 2, 517, 506, 3, 2, 2, 2, 517, 511, 3, 2, 2, 2, 517, 514, 3, 2,
	2, 2, 518, 552, 3, 2, 2, 2, 519, 520, 12, 14, 2, 2, 520, 521, 9, 2, 2,
	2, 521, 522, 5, 64, 33, 15, 522, 523, 8, 33, 1, 2, 523, 551, 3, 2, 2, 2,
	524, 525, 12, 13, 2, 2, 525, 526, 9, 3, 2, 2, 526, 527, 5, 64, 33, 14,
	527, 528, 8, 33, 1, 2, 528, 551, 3, 2, 2, 2, 529, 530, 12, 12, 2, 2, 530,
	531, 9, 4, 2, 2, 531, 532, 5, 64, 33, 13, 532, 533, 8, 33, 1, 2, 533, 551,
	3, 2, 2, 2, 534, 535, 12, 11, 2, 2, 535, 536, 9, 5, 2, 2, 536, 537, 5,
	64, 33, 12, 537, 538, 8, 33, 1, 2, 538, 551, 3, 2, 2, 2, 539, 540, 12,
	18, 2, 2, 540, 541, 7, 45, 2, 2, 541, 542, 5, 18, 10, 2, 542, 543, 8, 33,
	1, 2, 543, 551, 3, 2, 2, 2, 544, 545, 12, 17, 2, 2, 545, 546, 7, 3, 2,
	2, 546, 547, 7, 68, 2, 2, 547, 548, 7, 26, 2, 2, 548, 549, 7, 27, 2, 2,
	549, 551, 8, 33, 1, 2, 550, 519, 3, 2, 2, 2, 550, 524, 3, 2, 2, 2, 550,
	529, 3, 2, 2, 2, 550, 534, 3, 2, 2, 2, 550, 539, 3, 2, 2, 2, 550, 544,
	3, 2, 2, 2, 551, 554, 3, 2, 2, 2, 552, 550, 3, 2, 2, 2, 552, 553, 3, 2,
	2, 2, 553, 65, 3, 2, 2, 2, 554, 552, 3, 2, 2, 2, 555, 556, 7, 83, 2, 2,
	556, 570, 8, 34, 1, 2, 557, 558, 7, 85, 2, 2, 558, 570, 8, 34, 1, 2, 559,
	560, 7, 63, 2, 2, 560, 570, 8, 34, 1, 2, 561, 562, 7, 64, 2, 2, 562, 570,
	8, 34, 1, 2, 563, 564, 7, 84, 2, 2, 564, 570, 8, 34, 1, 2, 565, 566, 7,
	87, 2, 2, 566, 570, 8, 34, 1, 2, 567, 568, 7, 86, 2, 2, 568, 570, 8, 34,
	1, 2, 569, 555, 3, 2, 2, 2, 569, 557, 3, 2, 2, 2, 569, 559, 3, 2, 2, 2,
	569, 561, 3, 2, 2, 2, 569, 563, 3, 2, 2, 2, 569, 565, 3, 2, 2, 2, 569,
	567, 3, 2, 2, 2, 570, 67, 3, 2, 2, 2, 571, 572, 8, 35, 1, 2, 572, 573,
	5, 62, 32, 2, 573, 574, 8, 35, 1, 2, 574, 582, 3, 2, 2, 2, 575, 576, 12,
	4, 2, 2, 576, 577, 7, 5, 2, 2, 577, 578, 5, 62, 32, 2, 578, 579, 8, 35,
	1, 2, 579, 581, 3, 2, 2, 2, 580, 575, 3, 2, 2, 2, 581, 584, 3, 2, 2, 2,
	582, 580, 3, 2, 2, 2, 582, 583, 3, 2, 2, 2, 583, 69, 3, 2, 2, 2, 584, 582,
	3, 2, 2, 2, 585, 586, 7, 86, 2, 2, 586, 587, 5, 72, 37, 2, 587, 588, 8,
	36, 1, 2, 588, 71, 3, 2, 2, 2, 589, 590, 8, 37, 1, 2, 590, 591, 7, 30,
	2, 2, 591, 592, 5, 62, 32, 2, 592, 593, 7, 31, 2, 2, 593, 594, 8, 37, 1,
	2, 594, 603, 3, 2, 2, 2, 595, 596, 12, 3, 2, 2, 596, 597, 7, 30, 2, 2,
	597, 598, 5, 62, 32, 2, 598, 599, 7, 31, 2, 2, 599, 600, 8, 37, 1, 2, 600,
	602, 3, 2, 2, 2, 601, 595, 3, 2, 2, 2, 602, 605, 3, 2, 2, 2, 603, 601,
	3, 2, 2, 2, 603, 604, 3, 2, 2, 2, 604, 73, 3, 2, 2, 2, 605, 603, 3, 2,
	2, 2, 606, 607, 7, 30, 2, 2, 607, 608, 5, 62, 32, 2, 608, 609, 7, 4, 2,
	2, 609, 610, 5, 62, 32, 2, 610, 611, 7, 31, 2, 2, 611, 612, 8, 38, 1, 2,
	612, 75, 3, 2, 2, 2, 28, 82, 137, 181, 199, 212, 230, 250, 280, 303, 313,
	315, 334, 344, 361, 377, 387, 389, 408, 429, 449, 517, 550, 552, 569, 582,
	603,
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
	"start", "instrucciones", "instruccion", "instruccion_wc", "asignacion_var",
	"declaracion_var", "mutable", "types", "tipo_d", "asignar_Array", "dimensiones",
	"tipo_vector", "vectores", "loops", "ifs", "llaves_if", "elses", "breaks",
	"continues", "impr", "formato", "matches", "casos", "cases", "conditions",
	"defaults", "set_match", "rfor", "iter_for", "mod_Array", "expression",
	"expr_arit", "primitivo", "listValues", "arrayAcc", "listArray", "creatArray",
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
	ChemsRULE_instruccion_wc  = 3
	ChemsRULE_asignacion_var  = 4
	ChemsRULE_declaracion_var = 5
	ChemsRULE_mutable         = 6
	ChemsRULE_types           = 7
	ChemsRULE_tipo_d          = 8
	ChemsRULE_asignar_Array   = 9
	ChemsRULE_dimensiones     = 10
	ChemsRULE_tipo_vector     = 11
	ChemsRULE_vectores        = 12
	ChemsRULE_loops           = 13
	ChemsRULE_ifs             = 14
	ChemsRULE_llaves_if       = 15
	ChemsRULE_elses           = 16
	ChemsRULE_breaks          = 17
	ChemsRULE_continues       = 18
	ChemsRULE_impr            = 19
	ChemsRULE_formato         = 20
	ChemsRULE_matches         = 21
	ChemsRULE_casos           = 22
	ChemsRULE_cases           = 23
	ChemsRULE_conditions      = 24
	ChemsRULE_defaults        = 25
	ChemsRULE_set_match       = 26
	ChemsRULE_rfor            = 27
	ChemsRULE_iter_for        = 28
	ChemsRULE_mod_Array       = 29
	ChemsRULE_expression      = 30
	ChemsRULE_expr_arit       = 31
	ChemsRULE_primitivo       = 32
	ChemsRULE_listValues      = 33
	ChemsRULE_arrayAcc        = 34
	ChemsRULE_listArray       = 35
	ChemsRULE_creatArray      = 36
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
		p.SetState(74)

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
	p.SetState(80)
	p.GetErrorHandler().Sync(p)
	_la = p.GetTokenStream().LA(1)

	for (((_la-33)&-(0x1f+1)) == 0 && ((1<<uint((_la-33)))&((1<<(ChemsCONSOLE-33))|(1<<(ChemsPRINTLN-33))|(1<<(ChemsP_IF-33))|(1<<(ChemsP_WHILE-33))|(1<<(ChemsP_FOR-33))|(1<<(ChemsBREAK-33))|(1<<(ChemsCONTINUE-33))|(1<<(ChemsLET-33))|(1<<(ChemsMATCH-33))|(1<<(ChemsLOOP-33)))) != 0) || _la == ChemsID {
		{
			p.SetState(77)

			var _x = p.Instruccion()

			localctx.(*InstruccionesContext)._instruccion = _x
		}
		localctx.(*InstruccionesContext).e = append(localctx.(*InstruccionesContext).e, localctx.(*InstruccionesContext)._instruccion)

		p.SetState(82)
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

	// Get_ifs returns the _ifs rule contexts.
	Get_ifs() IIfsContext

	// Get_loops returns the _loops rule contexts.
	Get_loops() ILoopsContext

	// Get_impr returns the _impr rule contexts.
	Get_impr() IImprContext

	// Get_matches returns the _matches rule contexts.
	Get_matches() IMatchesContext

	// Get_rfor returns the _rfor rule contexts.
	Get_rfor() IRforContext

	// Get_mod_Array returns the _mod_Array rule contexts.
	Get_mod_Array() IMod_ArrayContext

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

	// Set_ifs sets the _ifs rule contexts.
	Set_ifs(IIfsContext)

	// Set_loops sets the _loops rule contexts.
	Set_loops(ILoopsContext)

	// Set_impr sets the _impr rule contexts.
	Set_impr(IImprContext)

	// Set_matches sets the _matches rule contexts.
	Set_matches(IMatchesContext)

	// Set_rfor sets the _rfor rule contexts.
	Set_rfor(IRforContext)

	// Set_mod_Array sets the _mod_Array rule contexts.
	Set_mod_Array(IMod_ArrayContext)

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
	_ifs             IIfsContext
	_loops           ILoopsContext
	_impr            IImprContext
	_matches         IMatchesContext
	_rfor            IRforContext
	_mod_Array       IMod_ArrayContext
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

func (s *InstruccionContext) Get_ifs() IIfsContext { return s._ifs }

func (s *InstruccionContext) Get_loops() ILoopsContext { return s._loops }

func (s *InstruccionContext) Get_impr() IImprContext { return s._impr }

func (s *InstruccionContext) Get_matches() IMatchesContext { return s._matches }

func (s *InstruccionContext) Get_rfor() IRforContext { return s._rfor }

func (s *InstruccionContext) Get_mod_Array() IMod_ArrayContext { return s._mod_Array }

func (s *InstruccionContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *InstruccionContext) Set_declaracion_var(v IDeclaracion_varContext) { s._declaracion_var = v }

func (s *InstruccionContext) Set_asignacion_var(v IAsignacion_varContext) { s._asignacion_var = v }

func (s *InstruccionContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *InstruccionContext) Set_breaks(v IBreaksContext) { s._breaks = v }

func (s *InstruccionContext) Set_continues(v IContinuesContext) { s._continues = v }

func (s *InstruccionContext) Set_ifs(v IIfsContext) { s._ifs = v }

func (s *InstruccionContext) Set_loops(v ILoopsContext) { s._loops = v }

func (s *InstruccionContext) Set_impr(v IImprContext) { s._impr = v }

func (s *InstruccionContext) Set_matches(v IMatchesContext) { s._matches = v }

func (s *InstruccionContext) Set_rfor(v IRforContext) { s._rfor = v }

func (s *InstruccionContext) Set_mod_Array(v IMod_ArrayContext) { s._mod_Array = v }

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

func (s *InstruccionContext) Ifs() IIfsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfsContext)
}

func (s *InstruccionContext) Loops() ILoopsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopsContext)
}

func (s *InstruccionContext) Impr() IImprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImprContext)
}

func (s *InstruccionContext) Matches() IMatchesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchesContext)
}

func (s *InstruccionContext) Rfor() IRforContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRforContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRforContext)
}

func (s *InstruccionContext) Mod_Array() IMod_ArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMod_ArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMod_ArrayContext)
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

	p.SetState(135)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 1, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(85)
			p.Match(ChemsCONSOLE)
		}
		{
			p.SetState(86)
			p.Match(ChemsPUNTO)
		}
		{
			p.SetState(87)
			p.Match(ChemsLOG)
		}
		{
			p.SetState(88)
			p.Match(ChemsPARIZQ)
		}
		{
			p.SetState(89)

			var _x = p.Expression()

			localctx.(*InstruccionContext)._expression = _x
		}
		{
			p.SetState(90)
			p.Match(ChemsPARDER)
		}
		{
			p.SetState(91)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*InstruccionContext).instr = instruction.NewImprimir(localctx.(*InstruccionContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(94)

			var _x = p.Declaracion_var()

			localctx.(*InstruccionContext)._declaracion_var = _x
		}
		{
			p.SetState(95)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_declaracion_var().GetI()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(98)

			var _x = p.Asignacion_var()

			localctx.(*InstruccionContext)._asignacion_var = _x
		}
		{
			p.SetState(99)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_asignacion_var().GetI()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(102)
			p.Match(ChemsP_WHILE)
		}
		{
			p.SetState(103)

			var _x = p.Expression()

			localctx.(*InstruccionContext)._expression = _x
		}
		{
			p.SetState(104)

			var _m = p.Match(ChemsLLAVEIZQ)

			localctx.(*InstruccionContext)._LLAVEIZQ = _m
		}
		{
			p.SetState(105)

			var _x = p.Instrucciones()

			localctx.(*InstruccionContext)._instrucciones = _x
		}
		{
			p.SetState(106)

			var _m = p.Match(ChemsLLAVEDER)

			localctx.(*InstruccionContext)._LLAVEDER = _m
		}
		localctx.(*InstruccionContext).instr = instruction.NewWhile(localctx.(*InstruccionContext).Get_expression().GetP(), localctx.(*InstruccionContext).Get_instrucciones().GetL(), localctx.(*InstruccionContext).Get_LLAVEIZQ().GetLine(), localctx.(*InstruccionContext).Get_LLAVEDER().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(109)

			var _x = p.Breaks()

			localctx.(*InstruccionContext)._breaks = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_breaks().GetI()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(112)

			var _x = p.Continues()

			localctx.(*InstruccionContext)._continues = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_continues().GetI()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(115)

			var _x = p.Ifs()

			localctx.(*InstruccionContext)._ifs = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_ifs().GetP()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(118)

			var _x = p.Loops()

			localctx.(*InstruccionContext)._loops = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_loops().GetI()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(121)

			var _x = p.Impr()

			localctx.(*InstruccionContext)._impr = _x
		}
		{
			p.SetState(122)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_impr().GetP()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(125)

			var _x = p.Matches()

			localctx.(*InstruccionContext)._matches = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_matches().GetM()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(128)

			var _x = p.Rfor()

			localctx.(*InstruccionContext)._rfor = _x
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_rfor().GetP()

	case 12:
		p.EnterOuterAlt(localctx, 12)
		{
			p.SetState(131)

			var _x = p.Mod_Array()

			localctx.(*InstruccionContext)._mod_Array = _x
		}
		{
			p.SetState(132)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*InstruccionContext).instr = localctx.(*InstruccionContext).Get_mod_Array().GetP()

	}

	return localctx
}

// IInstruccion_wcContext is an interface to support dynamic dispatch.
type IInstruccion_wcContext interface {
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

	// Get_ifs returns the _ifs rule contexts.
	Get_ifs() IIfsContext

	// Get_loops returns the _loops rule contexts.
	Get_loops() ILoopsContext

	// Get_impr returns the _impr rule contexts.
	Get_impr() IImprContext

	// Get_rfor returns the _rfor rule contexts.
	Get_rfor() IRforContext

	// Get_mod_Array returns the _mod_Array rule contexts.
	Get_mod_Array() IMod_ArrayContext

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

	// Set_ifs sets the _ifs rule contexts.
	Set_ifs(IIfsContext)

	// Set_loops sets the _loops rule contexts.
	Set_loops(ILoopsContext)

	// Set_impr sets the _impr rule contexts.
	Set_impr(IImprContext)

	// Set_rfor sets the _rfor rule contexts.
	Set_rfor(IRforContext)

	// Set_mod_Array sets the _mod_Array rule contexts.
	Set_mod_Array(IMod_ArrayContext)

	// GetInstr returns the instr attribute.
	GetInstr() interfaces.Instruction

	// SetInstr sets the instr attribute.
	SetInstr(interfaces.Instruction)

	// IsInstruccion_wcContext differentiates from other interfaces.
	IsInstruccion_wcContext()
}

type Instruccion_wcContext struct {
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
	_ifs             IIfsContext
	_loops           ILoopsContext
	_impr            IImprContext
	_rfor            IRforContext
	_mod_Array       IMod_ArrayContext
}

func NewEmptyInstruccion_wcContext() *Instruccion_wcContext {
	var p = new(Instruccion_wcContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_instruccion_wc
	return p
}

func (*Instruccion_wcContext) IsInstruccion_wcContext() {}

func NewInstruccion_wcContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Instruccion_wcContext {
	var p = new(Instruccion_wcContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_instruccion_wc

	return p
}

func (s *Instruccion_wcContext) GetParser() antlr.Parser { return s.parser }

func (s *Instruccion_wcContext) Get_LLAVEIZQ() antlr.Token { return s._LLAVEIZQ }

func (s *Instruccion_wcContext) Get_LLAVEDER() antlr.Token { return s._LLAVEDER }

func (s *Instruccion_wcContext) Set_LLAVEIZQ(v antlr.Token) { s._LLAVEIZQ = v }

func (s *Instruccion_wcContext) Set_LLAVEDER(v antlr.Token) { s._LLAVEDER = v }

func (s *Instruccion_wcContext) Get_expression() IExpressionContext { return s._expression }

func (s *Instruccion_wcContext) Get_declaracion_var() IDeclaracion_varContext {
	return s._declaracion_var
}

func (s *Instruccion_wcContext) Get_asignacion_var() IAsignacion_varContext { return s._asignacion_var }

func (s *Instruccion_wcContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Instruccion_wcContext) Get_breaks() IBreaksContext { return s._breaks }

func (s *Instruccion_wcContext) Get_continues() IContinuesContext { return s._continues }

func (s *Instruccion_wcContext) Get_ifs() IIfsContext { return s._ifs }

func (s *Instruccion_wcContext) Get_loops() ILoopsContext { return s._loops }

func (s *Instruccion_wcContext) Get_impr() IImprContext { return s._impr }

func (s *Instruccion_wcContext) Get_rfor() IRforContext { return s._rfor }

func (s *Instruccion_wcContext) Get_mod_Array() IMod_ArrayContext { return s._mod_Array }

func (s *Instruccion_wcContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Instruccion_wcContext) Set_declaracion_var(v IDeclaracion_varContext) {
	s._declaracion_var = v
}

func (s *Instruccion_wcContext) Set_asignacion_var(v IAsignacion_varContext) { s._asignacion_var = v }

func (s *Instruccion_wcContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Instruccion_wcContext) Set_breaks(v IBreaksContext) { s._breaks = v }

func (s *Instruccion_wcContext) Set_continues(v IContinuesContext) { s._continues = v }

func (s *Instruccion_wcContext) Set_ifs(v IIfsContext) { s._ifs = v }

func (s *Instruccion_wcContext) Set_loops(v ILoopsContext) { s._loops = v }

func (s *Instruccion_wcContext) Set_impr(v IImprContext) { s._impr = v }

func (s *Instruccion_wcContext) Set_rfor(v IRforContext) { s._rfor = v }

func (s *Instruccion_wcContext) Set_mod_Array(v IMod_ArrayContext) { s._mod_Array = v }

func (s *Instruccion_wcContext) GetInstr() interfaces.Instruction { return s.instr }

func (s *Instruccion_wcContext) SetInstr(v interfaces.Instruction) { s.instr = v }

func (s *Instruccion_wcContext) CONSOLE() antlr.TerminalNode {
	return s.GetToken(ChemsCONSOLE, 0)
}

func (s *Instruccion_wcContext) PUNTO() antlr.TerminalNode {
	return s.GetToken(ChemsPUNTO, 0)
}

func (s *Instruccion_wcContext) LOG() antlr.TerminalNode {
	return s.GetToken(ChemsLOG, 0)
}

func (s *Instruccion_wcContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsPARIZQ, 0)
}

func (s *Instruccion_wcContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Instruccion_wcContext) PARDER() antlr.TerminalNode {
	return s.GetToken(ChemsPARDER, 0)
}

func (s *Instruccion_wcContext) Declaracion_var() IDeclaracion_varContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDeclaracion_varContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDeclaracion_varContext)
}

func (s *Instruccion_wcContext) Asignacion_var() IAsignacion_varContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IAsignacion_varContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IAsignacion_varContext)
}

func (s *Instruccion_wcContext) P_WHILE() antlr.TerminalNode {
	return s.GetToken(ChemsP_WHILE, 0)
}

func (s *Instruccion_wcContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEIZQ, 0)
}

func (s *Instruccion_wcContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Instruccion_wcContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEDER, 0)
}

func (s *Instruccion_wcContext) Breaks() IBreaksContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IBreaksContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IBreaksContext)
}

func (s *Instruccion_wcContext) Continues() IContinuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IContinuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IContinuesContext)
}

func (s *Instruccion_wcContext) Ifs() IIfsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfsContext)
}

func (s *Instruccion_wcContext) Loops() ILoopsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopsContext)
}

func (s *Instruccion_wcContext) Impr() IImprContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IImprContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IImprContext)
}

func (s *Instruccion_wcContext) Rfor() IRforContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IRforContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IRforContext)
}

func (s *Instruccion_wcContext) Mod_Array() IMod_ArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMod_ArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMod_ArrayContext)
}

func (s *Instruccion_wcContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Instruccion_wcContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Instruccion_wcContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterInstruccion_wc(s)
	}
}

func (s *Instruccion_wcContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitInstruccion_wc(s)
	}
}

func (p *Chems) Instruccion_wc() (localctx IInstruccion_wcContext) {
	this := p
	_ = this

	localctx = NewInstruccion_wcContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 6, ChemsRULE_instruccion_wc)

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

	p.SetState(179)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 2, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(137)
			p.Match(ChemsCONSOLE)
		}
		{
			p.SetState(138)
			p.Match(ChemsPUNTO)
		}
		{
			p.SetState(139)
			p.Match(ChemsLOG)
		}
		{
			p.SetState(140)
			p.Match(ChemsPARIZQ)
		}
		{
			p.SetState(141)

			var _x = p.Expression()

			localctx.(*Instruccion_wcContext)._expression = _x
		}
		{
			p.SetState(142)
			p.Match(ChemsPARDER)
		}
		localctx.(*Instruccion_wcContext).instr = instruction.NewImprimir(localctx.(*Instruccion_wcContext).Get_expression().GetP())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(145)

			var _x = p.Declaracion_var()

			localctx.(*Instruccion_wcContext)._declaracion_var = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_declaracion_var().GetI()

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(148)

			var _x = p.Asignacion_var()

			localctx.(*Instruccion_wcContext)._asignacion_var = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_asignacion_var().GetI()

	case 4:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(151)
			p.Match(ChemsP_WHILE)
		}
		{
			p.SetState(152)

			var _x = p.Expression()

			localctx.(*Instruccion_wcContext)._expression = _x
		}
		{
			p.SetState(153)

			var _m = p.Match(ChemsLLAVEIZQ)

			localctx.(*Instruccion_wcContext)._LLAVEIZQ = _m
		}
		{
			p.SetState(154)

			var _x = p.Instrucciones()

			localctx.(*Instruccion_wcContext)._instrucciones = _x
		}
		{
			p.SetState(155)

			var _m = p.Match(ChemsLLAVEDER)

			localctx.(*Instruccion_wcContext)._LLAVEDER = _m
		}
		localctx.(*Instruccion_wcContext).instr = instruction.NewWhile(localctx.(*Instruccion_wcContext).Get_expression().GetP(), localctx.(*Instruccion_wcContext).Get_instrucciones().GetL(), localctx.(*Instruccion_wcContext).Get_LLAVEIZQ().GetLine(), localctx.(*Instruccion_wcContext).Get_LLAVEDER().GetColumn())

	case 5:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(158)

			var _x = p.Breaks()

			localctx.(*Instruccion_wcContext)._breaks = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_breaks().GetI()

	case 6:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(161)

			var _x = p.Continues()

			localctx.(*Instruccion_wcContext)._continues = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_continues().GetI()

	case 7:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(164)

			var _x = p.Ifs()

			localctx.(*Instruccion_wcContext)._ifs = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_ifs().GetP()

	case 8:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(167)

			var _x = p.Loops()

			localctx.(*Instruccion_wcContext)._loops = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_loops().GetI()

	case 9:
		p.EnterOuterAlt(localctx, 9)
		{
			p.SetState(170)

			var _x = p.Impr()

			localctx.(*Instruccion_wcContext)._impr = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_impr().GetP()

	case 10:
		p.EnterOuterAlt(localctx, 10)
		{
			p.SetState(173)

			var _x = p.Rfor()

			localctx.(*Instruccion_wcContext)._rfor = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_rfor().GetP()

	case 11:
		p.EnterOuterAlt(localctx, 11)
		{
			p.SetState(176)

			var _x = p.Mod_Array()

			localctx.(*Instruccion_wcContext)._mod_Array = _x
		}
		localctx.(*Instruccion_wcContext).instr = localctx.(*Instruccion_wcContext).Get_mod_Array().GetP()

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
	p.EnterRule(localctx, 8, ChemsRULE_asignacion_var)

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
		p.SetState(181)

		var _m = p.Match(ChemsID)

		localctx.(*Asignacion_varContext).id = _m
	}
	{
		p.SetState(182)
		p.Match(ChemsIGUAL)
	}
	{
		p.SetState(183)

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
	p.EnterRule(localctx, 10, ChemsRULE_declaracion_var)

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

		var _m = p.Match(ChemsLET)

		localctx.(*Declaracion_varContext)._LET = _m
	}
	{
		p.SetState(187)

		var _x = p.Mutable()

		localctx.(*Declaracion_varContext)._mutable = _x
	}
	{
		p.SetState(188)

		var _m = p.Match(ChemsID)

		localctx.(*Declaracion_varContext).id = _m
	}
	{
		p.SetState(189)

		var _x = p.Types()

		localctx.(*Declaracion_varContext)._types = _x
	}
	{
		p.SetState(190)
		p.Match(ChemsIGUAL)
	}
	{
		p.SetState(191)

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
	p.EnterRule(localctx, 12, ChemsRULE_mutable)

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

	p.SetState(197)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsMUT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(194)
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
	p.EnterRule(localctx, 14, ChemsRULE_types)

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

	p.SetState(210)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 4, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(199)
			p.Match(ChemsDOSPUNTOS)
		}
		{
			p.SetState(200)

			var _x = p.Tipo_vector()

			localctx.(*TypesContext)._tipo_vector = _x
		}
		localctx.(*TypesContext).l = localctx.(*TypesContext).Get_tipo_vector().GetT()

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(203)

			var _x = p.Asignar_Array()

			localctx.(*TypesContext).a = _x
		}

		dim := arrayList.New()
		dim.Add(localctx.(*TypesContext).GetA().GetD())
		localctx.(*TypesContext).l = interfaces.TipoSimbolo{interfaces.ARRAY, dim}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(206)
			p.Match(ChemsDOSPUNTOS)
		}
		{
			p.SetState(207)

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
	p.EnterRule(localctx, 16, ChemsRULE_tipo_d)

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

	p.SetState(228)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(212)
			p.Match(ChemsINT)
		}
		localctx.(*Tipo_dContext).t = interfaces.INTEGER

	case ChemsFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(214)
			p.Match(ChemsFLOAT)
		}
		localctx.(*Tipo_dContext).t = interfaces.FLOAT

	case ChemsBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(216)
			p.Match(ChemsBOOL)
		}
		localctx.(*Tipo_dContext).t = interfaces.BOOLEAN

	case ChemsCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(218)
			p.Match(ChemsCHAR)
		}
		localctx.(*Tipo_dContext).t = interfaces.CHAR

	case ChemsSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(220)
			p.Match(ChemsSTR)
		}
		localctx.(*Tipo_dContext).t = interfaces.STRING

	case ChemsP_STRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(222)
			p.Match(ChemsP_STRING)
		}
		localctx.(*Tipo_dContext).t = interfaces.STR

	case ChemsUSIZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(224)
			p.Match(ChemsUSIZE)
		}
		localctx.(*Tipo_dContext).t = interfaces.USIZE

	case ChemsID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(226)
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
	p.EnterRule(localctx, 18, ChemsRULE_asignar_Array)

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
		p.SetState(230)
		p.Match(ChemsDOSPUNTOS)
	}
	{
		p.SetState(231)

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
	p.EnterRule(localctx, 20, ChemsRULE_dimensiones)

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

	p.SetState(248)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 6, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(234)
			p.Match(ChemsCORIZQ)
		}
		{
			p.SetState(235)

			var _x = p.Tipo_d()

			localctx.(*DimensionesContext)._tipo_d = _x
		}
		{
			p.SetState(236)
			p.Match(ChemsPTCOMA)
		}
		{
			p.SetState(237)

			var _x = p.Expression()

			localctx.(*DimensionesContext)._expression = _x
		}
		{
			p.SetState(238)
			p.Match(ChemsCORDER)
		}

		list := arrayList.New()
		list.Add(localctx.(*DimensionesContext).Get_expression().GetP())
		localctx.(*DimensionesContext).d = interfaces.Dimensions{localctx.(*DimensionesContext).Get_tipo_d().GetT(), list}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(241)
			p.Match(ChemsCORIZQ)
		}
		{
			p.SetState(242)

			var _x = p.Dimensiones()

			localctx.(*DimensionesContext)._dimensiones = _x
		}
		{
			p.SetState(243)
			p.Match(ChemsPTCOMA)
		}
		{
			p.SetState(244)

			var _x = p.Expression()

			localctx.(*DimensionesContext)._expression = _x
		}
		{
			p.SetState(245)
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
	p.EnterRule(localctx, 22, ChemsRULE_tipo_vector)

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
		p.SetState(250)
		p.Match(ChemsVECT)
	}
	{
		p.SetState(251)
		p.Match(ChemsMENOR)
	}
	{
		p.SetState(252)

		var _x = p.Vectores()

		localctx.(*Tipo_vectorContext)._vectores = _x
	}
	{
		p.SetState(253)
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
	p.EnterRule(localctx, 24, ChemsRULE_vectores)

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

	p.SetState(278)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsINT:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(256)
			p.Match(ChemsINT)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.INTEGER)

	case ChemsFLOAT:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(258)
			p.Match(ChemsFLOAT)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.FLOAT)

	case ChemsBOOL:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(260)
			p.Match(ChemsBOOL)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.BOOLEAN)

	case ChemsCHAR:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(262)
			p.Match(ChemsCHAR)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.CHAR)

	case ChemsSTR:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(264)
			p.Match(ChemsSTR)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.STRING)

	case ChemsP_STRING:
		p.EnterOuterAlt(localctx, 6)
		{
			p.SetState(266)
			p.Match(ChemsP_STRING)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.STR)

	case ChemsUSIZE:
		p.EnterOuterAlt(localctx, 7)
		{
			p.SetState(268)
			p.Match(ChemsUSIZE)
		}
		localctx.(*VectoresContext).l = arrayList.New()
		localctx.(*VectoresContext).l.Add(interfaces.USIZE)

	case ChemsID:
		p.EnterOuterAlt(localctx, 8)
		{
			p.SetState(270)

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
			p.SetState(272)
			p.Match(ChemsVECT)
		}
		{
			p.SetState(273)
			p.Match(ChemsMENOR)
		}
		{
			p.SetState(274)

			var _x = p.Vectores()

			localctx.(*VectoresContext)._vectores = _x
		}
		{
			p.SetState(275)
			p.Match(ChemsMAYOR)
		}

		localctx.(*VectoresContext).Get_vectores().GetL().Add(interfaces.VECTOR)
		localctx.(*VectoresContext).l = localctx.(*VectoresContext).Get_vectores().GetL()

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}

	return localctx
}

// ILoopsContext is an interface to support dynamic dispatch.
type ILoopsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetI returns the i attribute.
	GetI() interfaces.Instruction

	// SetI sets the i attribute.
	SetI(interfaces.Instruction)

	// IsLoopsContext differentiates from other interfaces.
	IsLoopsContext()
}

type LoopsContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	i              interfaces.Instruction
	_instrucciones IInstruccionesContext
}

func NewEmptyLoopsContext() *LoopsContext {
	var p = new(LoopsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_loops
	return p
}

func (*LoopsContext) IsLoopsContext() {}

func NewLoopsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *LoopsContext {
	var p = new(LoopsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_loops

	return p
}

func (s *LoopsContext) GetParser() antlr.Parser { return s.parser }

func (s *LoopsContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *LoopsContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *LoopsContext) GetI() interfaces.Instruction { return s.i }

func (s *LoopsContext) SetI(v interfaces.Instruction) { s.i = v }

func (s *LoopsContext) LOOP() antlr.TerminalNode {
	return s.GetToken(ChemsLOOP, 0)
}

func (s *LoopsContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEIZQ, 0)
}

func (s *LoopsContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *LoopsContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEDER, 0)
}

func (s *LoopsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *LoopsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *LoopsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterLoops(s)
	}
}

func (s *LoopsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitLoops(s)
	}
}

func (p *Chems) Loops() (localctx ILoopsContext) {
	this := p
	_ = this

	localctx = NewLoopsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 26, ChemsRULE_loops)

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
		p.SetState(280)
		p.Match(ChemsLOOP)
	}
	{
		p.SetState(281)
		p.Match(ChemsLLAVEIZQ)
	}
	{
		p.SetState(282)

		var _x = p.Instrucciones()

		localctx.(*LoopsContext)._instrucciones = _x
	}
	{
		p.SetState(283)
		p.Match(ChemsLLAVEDER)
	}
	localctx.(*LoopsContext).i = instruction.NewLoop(localctx.(*LoopsContext).Get_instrucciones().GetL())
	fmt.Println("TOY EN EL ANALISIS")

	return localctx
}

// IIfsContext is an interface to support dynamic dispatch.
type IIfsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_P_IF returns the _P_IF token.
	Get_P_IF() antlr.Token

	// Set_P_IF sets the _P_IF token.
	Set_P_IF(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_llaves_if returns the _llaves_if rule contexts.
	Get_llaves_if() ILlaves_ifContext

	// Get_elses returns the _elses rule contexts.
	Get_elses() IElsesContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_llaves_if sets the _llaves_if rule contexts.
	Set_llaves_if(ILlaves_ifContext)

	// Set_elses sets the _elses rule contexts.
	Set_elses(IElsesContext)

	// GetP returns the p attribute.
	GetP() interfaces.Instruction

	// SetP sets the p attribute.
	SetP(interfaces.Instruction)

	// IsIfsContext differentiates from other interfaces.
	IsIfsContext()
}

type IfsContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Instruction
	_P_IF       antlr.Token
	_expression IExpressionContext
	_llaves_if  ILlaves_ifContext
	_elses      IElsesContext
}

func NewEmptyIfsContext() *IfsContext {
	var p = new(IfsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_ifs
	return p
}

func (*IfsContext) IsIfsContext() {}

func NewIfsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *IfsContext {
	var p = new(IfsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_ifs

	return p
}

func (s *IfsContext) GetParser() antlr.Parser { return s.parser }

func (s *IfsContext) Get_P_IF() antlr.Token { return s._P_IF }

func (s *IfsContext) Set_P_IF(v antlr.Token) { s._P_IF = v }

func (s *IfsContext) Get_expression() IExpressionContext { return s._expression }

func (s *IfsContext) Get_llaves_if() ILlaves_ifContext { return s._llaves_if }

func (s *IfsContext) Get_elses() IElsesContext { return s._elses }

func (s *IfsContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *IfsContext) Set_llaves_if(v ILlaves_ifContext) { s._llaves_if = v }

func (s *IfsContext) Set_elses(v IElsesContext) { s._elses = v }

func (s *IfsContext) GetP() interfaces.Instruction { return s.p }

func (s *IfsContext) SetP(v interfaces.Instruction) { s.p = v }

func (s *IfsContext) P_IF() antlr.TerminalNode {
	return s.GetToken(ChemsP_IF, 0)
}

func (s *IfsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *IfsContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEIZQ, 0)
}

func (s *IfsContext) Llaves_if() ILlaves_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILlaves_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILlaves_ifContext)
}

func (s *IfsContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEDER, 0)
}

func (s *IfsContext) Elses() IElsesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElsesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElsesContext)
}

func (s *IfsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *IfsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *IfsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterIfs(s)
	}
}

func (s *IfsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitIfs(s)
	}
}

func (p *Chems) Ifs() (localctx IIfsContext) {
	this := p
	_ = this

	localctx = NewIfsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 28, ChemsRULE_ifs)

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
		p.SetState(286)

		var _m = p.Match(ChemsP_IF)

		localctx.(*IfsContext)._P_IF = _m
	}
	{
		p.SetState(287)

		var _x = p.Expression()

		localctx.(*IfsContext)._expression = _x
	}
	{
		p.SetState(288)
		p.Match(ChemsLLAVEIZQ)
	}
	{
		p.SetState(289)

		var _x = p.llaves_if(0)

		localctx.(*IfsContext)._llaves_if = _x
	}
	{
		p.SetState(290)
		p.Match(ChemsLLAVEDER)
	}
	{
		p.SetState(291)

		var _x = p.Elses()

		localctx.(*IfsContext)._elses = _x
	}
	localctx.(*IfsContext).p = instruction.NewIf(localctx.(*IfsContext).Get_expression().GetP(), localctx.(*IfsContext).Get_llaves_if().GetL(), localctx.(*IfsContext).Get_elses().GetE(), localctx.(*IfsContext).Get_P_IF().GetLine(), localctx.(*IfsContext).Get_P_IF().GetColumn(), 0)

	return localctx
}

// ILlaves_ifContext is an interface to support dynamic dispatch.
type ILlaves_ifContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetK returns the k rule contexts.
	GetK() ILlaves_ifContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_instruccion returns the _instruccion rule contexts.
	Get_instruccion() IInstruccionContext

	// SetK sets the k rule contexts.
	SetK(ILlaves_ifContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_instruccion sets the _instruccion rule contexts.
	Set_instruccion(IInstruccionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsLlaves_ifContext differentiates from other interfaces.
	IsLlaves_ifContext()
}

type Llaves_ifContext struct {
	*antlr.BaseParserRuleContext
	parser       antlr.Parser
	l            *arrayList.List
	k            ILlaves_ifContext
	_expression  IExpressionContext
	_instruccion IInstruccionContext
}

func NewEmptyLlaves_ifContext() *Llaves_ifContext {
	var p = new(Llaves_ifContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_llaves_if
	return p
}

func (*Llaves_ifContext) IsLlaves_ifContext() {}

func NewLlaves_ifContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Llaves_ifContext {
	var p = new(Llaves_ifContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_llaves_if

	return p
}

func (s *Llaves_ifContext) GetParser() antlr.Parser { return s.parser }

func (s *Llaves_ifContext) GetK() ILlaves_ifContext { return s.k }

func (s *Llaves_ifContext) Get_expression() IExpressionContext { return s._expression }

func (s *Llaves_ifContext) Get_instruccion() IInstruccionContext { return s._instruccion }

func (s *Llaves_ifContext) SetK(v ILlaves_ifContext) { s.k = v }

func (s *Llaves_ifContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Llaves_ifContext) Set_instruccion(v IInstruccionContext) { s._instruccion = v }

func (s *Llaves_ifContext) GetL() *arrayList.List { return s.l }

func (s *Llaves_ifContext) SetL(v *arrayList.List) { s.l = v }

func (s *Llaves_ifContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Llaves_ifContext) Instruccion() IInstruccionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionContext)
}

func (s *Llaves_ifContext) Llaves_if() ILlaves_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILlaves_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILlaves_ifContext)
}

func (s *Llaves_ifContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Llaves_ifContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Llaves_ifContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterLlaves_if(s)
	}
}

func (s *Llaves_ifContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitLlaves_if(s)
	}
}

func (p *Chems) Llaves_if() (localctx ILlaves_ifContext) {
	return p.llaves_if(0)
}

func (p *Chems) llaves_if(_p int) (localctx ILlaves_ifContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewLlaves_ifContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ILlaves_ifContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 30
	p.EnterRecursionRule(localctx, 30, ChemsRULE_llaves_if, _p)

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
	p.SetState(301)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 8, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(295)

			var _x = p.Expression()

			localctx.(*Llaves_ifContext)._expression = _x
		}
		localctx.(*Llaves_ifContext).l = arrayList.New()
		i := interfaces.OpcionIf{0, localctx.(*Llaves_ifContext).Get_expression().GetP()}
		localctx.(*Llaves_ifContext).l.Add(i)

	case 2:
		{
			p.SetState(298)

			var _x = p.Instruccion()

			localctx.(*Llaves_ifContext)._instruccion = _x
		}
		localctx.(*Llaves_ifContext).l = arrayList.New()
		i := interfaces.OpcionIf{1, localctx.(*Llaves_ifContext).Get_instruccion().GetInstr()}
		localctx.(*Llaves_ifContext).l.Add(i)

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(313)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(311)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 9, p.GetParserRuleContext()) {
			case 1:
				localctx = NewLlaves_ifContext(p, _parentctx, _parentState)
				localctx.(*Llaves_ifContext).k = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_llaves_if)
				p.SetState(303)

				if !(p.Precpred(p.GetParserRuleContext(), 2)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
				}
				{
					p.SetState(304)

					var _x = p.Expression()

					localctx.(*Llaves_ifContext)._expression = _x
				}

				i := interfaces.OpcionIf{0, localctx.(*Llaves_ifContext).Get_expression().GetP()}
				localctx.(*Llaves_ifContext).GetK().GetL().Add(i)
				localctx.(*Llaves_ifContext).l = localctx.(*Llaves_ifContext).GetK().GetL()

			case 2:
				localctx = NewLlaves_ifContext(p, _parentctx, _parentState)
				localctx.(*Llaves_ifContext).k = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_llaves_if)
				p.SetState(307)

				if !(p.Precpred(p.GetParserRuleContext(), 1)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
				}
				{
					p.SetState(308)

					var _x = p.Instruccion()

					localctx.(*Llaves_ifContext)._instruccion = _x
				}

				i := interfaces.OpcionIf{1, localctx.(*Llaves_ifContext).Get_instruccion().GetInstr()}
				localctx.(*Llaves_ifContext).GetK().GetL().Add(i)
				localctx.(*Llaves_ifContext).l = localctx.(*Llaves_ifContext).GetK().GetL()

			}

		}
		p.SetState(315)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 10, p.GetParserRuleContext())
	}

	return localctx
}

// IElsesContext is an interface to support dynamic dispatch.
type IElsesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_P_IF returns the _P_IF token.
	Get_P_IF() antlr.Token

	// Set_P_IF sets the _P_IF token.
	Set_P_IF(antlr.Token)

	// Get_llaves_if returns the _llaves_if rule contexts.
	Get_llaves_if() ILlaves_ifContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_elses returns the _elses rule contexts.
	Get_elses() IElsesContext

	// Set_llaves_if sets the _llaves_if rule contexts.
	Set_llaves_if(ILlaves_ifContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_elses sets the _elses rule contexts.
	Set_elses(IElsesContext)

	// GetE returns the e attribute.
	GetE() interfaces.Instruction

	// SetE sets the e attribute.
	SetE(interfaces.Instruction)

	// IsElsesContext differentiates from other interfaces.
	IsElsesContext()
}

type ElsesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	e           interfaces.Instruction
	_llaves_if  ILlaves_ifContext
	_P_IF       antlr.Token
	_expression IExpressionContext
	_elses      IElsesContext
}

func NewEmptyElsesContext() *ElsesContext {
	var p = new(ElsesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_elses
	return p
}

func (*ElsesContext) IsElsesContext() {}

func NewElsesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ElsesContext {
	var p = new(ElsesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_elses

	return p
}

func (s *ElsesContext) GetParser() antlr.Parser { return s.parser }

func (s *ElsesContext) Get_P_IF() antlr.Token { return s._P_IF }

func (s *ElsesContext) Set_P_IF(v antlr.Token) { s._P_IF = v }

func (s *ElsesContext) Get_llaves_if() ILlaves_ifContext { return s._llaves_if }

func (s *ElsesContext) Get_expression() IExpressionContext { return s._expression }

func (s *ElsesContext) Get_elses() IElsesContext { return s._elses }

func (s *ElsesContext) Set_llaves_if(v ILlaves_ifContext) { s._llaves_if = v }

func (s *ElsesContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ElsesContext) Set_elses(v IElsesContext) { s._elses = v }

func (s *ElsesContext) GetE() interfaces.Instruction { return s.e }

func (s *ElsesContext) SetE(v interfaces.Instruction) { s.e = v }

func (s *ElsesContext) P_ELSE() antlr.TerminalNode {
	return s.GetToken(ChemsP_ELSE, 0)
}

func (s *ElsesContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEIZQ, 0)
}

func (s *ElsesContext) Llaves_if() ILlaves_ifContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILlaves_ifContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILlaves_ifContext)
}

func (s *ElsesContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEDER, 0)
}

func (s *ElsesContext) P_IF() antlr.TerminalNode {
	return s.GetToken(ChemsP_IF, 0)
}

func (s *ElsesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ElsesContext) Elses() IElsesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IElsesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IElsesContext)
}

func (s *ElsesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ElsesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ElsesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterElses(s)
	}
}

func (s *ElsesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitElses(s)
	}
}

func (p *Chems) Elses() (localctx IElsesContext) {
	this := p
	_ = this

	localctx = NewElsesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 32, ChemsRULE_elses)

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

	p.SetState(332)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 11, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(316)
			p.Match(ChemsP_ELSE)
		}
		{
			p.SetState(317)
			p.Match(ChemsLLAVEIZQ)
		}
		{
			p.SetState(318)

			var _x = p.llaves_if(0)

			localctx.(*ElsesContext)._llaves_if = _x
		}
		{
			p.SetState(319)
			p.Match(ChemsLLAVEDER)
		}
		localctx.(*ElsesContext).e = instruction.NewIf(expresion.NewPrimitivo(1, interfaces.BOOLEAN, 0, 0), localctx.(*ElsesContext).Get_llaves_if().GetL(), instruction.NewElseNull("null"), 0, 0, 3)

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(322)
			p.Match(ChemsP_ELSE)
		}
		{
			p.SetState(323)

			var _m = p.Match(ChemsP_IF)

			localctx.(*ElsesContext)._P_IF = _m
		}
		{
			p.SetState(324)

			var _x = p.Expression()

			localctx.(*ElsesContext)._expression = _x
		}
		{
			p.SetState(325)
			p.Match(ChemsLLAVEIZQ)
		}
		{
			p.SetState(326)

			var _x = p.llaves_if(0)

			localctx.(*ElsesContext)._llaves_if = _x
		}
		{
			p.SetState(327)
			p.Match(ChemsLLAVEDER)
		}
		{
			p.SetState(328)

			var _x = p.Elses()

			localctx.(*ElsesContext)._elses = _x
		}
		localctx.(*ElsesContext).e = instruction.NewIf(localctx.(*ElsesContext).Get_expression().GetP(), localctx.(*ElsesContext).Get_llaves_if().GetL(), localctx.(*ElsesContext).Get_elses().GetE(), localctx.(*ElsesContext).Get_P_IF().GetLine(), localctx.(*ElsesContext).Get_P_IF().GetColumn(), 2)

	case 3:
		p.EnterOuterAlt(localctx, 3)
		localctx.(*ElsesContext).e = instruction.NewElseNull("null")

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
	p.EnterRule(localctx, 34, ChemsRULE_breaks)

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

	p.SetState(342)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 12, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(334)

			var _m = p.Match(ChemsBREAK)

			localctx.(*BreaksContext)._BREAK = _m
		}
		{
			p.SetState(335)

			var _x = p.Expression()

			localctx.(*BreaksContext)._expression = _x
		}
		{
			p.SetState(336)
			p.Match(ChemsPTCOMA)
		}
		localctx.(*BreaksContext).i = instruction.NewBreak(localctx.(*BreaksContext).Get_expression().GetP(), true, localctx.(*BreaksContext).Get_BREAK().GetLine(), localctx.(*BreaksContext).Get_BREAK().GetColumn())

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(339)

			var _m = p.Match(ChemsBREAK)

			localctx.(*BreaksContext)._BREAK = _m
		}
		{
			p.SetState(340)
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
	p.EnterRule(localctx, 36, ChemsRULE_continues)

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
		p.SetState(344)

		var _m = p.Match(ChemsCONTINUE)

		localctx.(*ContinuesContext)._CONTINUE = _m
	}
	{
		p.SetState(345)
		p.Match(ChemsPTCOMA)
	}
	localctx.(*ContinuesContext).i = instruction.NewContinue("continue", localctx.(*ContinuesContext).Get_CONTINUE().GetLine(), localctx.(*ContinuesContext).Get_CONTINUE().GetColumn())

	return localctx
}

// IImprContext is an interface to support dynamic dispatch.
type IImprContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_PARIZQ returns the _PARIZQ token.
	Get_PARIZQ() antlr.Token

	// Set_PARIZQ sets the _PARIZQ token.
	Set_PARIZQ(antlr.Token)

	// Get_formato returns the _formato rule contexts.
	Get_formato() IFormatoContext

	// Get_listValues returns the _listValues rule contexts.
	Get_listValues() IListValuesContext

	// Set_formato sets the _formato rule contexts.
	Set_formato(IFormatoContext)

	// Set_listValues sets the _listValues rule contexts.
	Set_listValues(IListValuesContext)

	// GetP returns the p attribute.
	GetP() interfaces.Instruction

	// SetP sets the p attribute.
	SetP(interfaces.Instruction)

	// IsImprContext differentiates from other interfaces.
	IsImprContext()
}

type ImprContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Instruction
	_PARIZQ     antlr.Token
	_formato    IFormatoContext
	_listValues IListValuesContext
}

func NewEmptyImprContext() *ImprContext {
	var p = new(ImprContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_impr
	return p
}

func (*ImprContext) IsImprContext() {}

func NewImprContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ImprContext {
	var p = new(ImprContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_impr

	return p
}

func (s *ImprContext) GetParser() antlr.Parser { return s.parser }

func (s *ImprContext) Get_PARIZQ() antlr.Token { return s._PARIZQ }

func (s *ImprContext) Set_PARIZQ(v antlr.Token) { s._PARIZQ = v }

func (s *ImprContext) Get_formato() IFormatoContext { return s._formato }

func (s *ImprContext) Get_listValues() IListValuesContext { return s._listValues }

func (s *ImprContext) Set_formato(v IFormatoContext) { s._formato = v }

func (s *ImprContext) Set_listValues(v IListValuesContext) { s._listValues = v }

func (s *ImprContext) GetP() interfaces.Instruction { return s.p }

func (s *ImprContext) SetP(v interfaces.Instruction) { s.p = v }

func (s *ImprContext) PRINTLN() antlr.TerminalNode {
	return s.GetToken(ChemsPRINTLN, 0)
}

func (s *ImprContext) PARIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsPARIZQ, 0)
}

func (s *ImprContext) Formato() IFormatoContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IFormatoContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IFormatoContext)
}

func (s *ImprContext) ListValues() IListValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListValuesContext)
}

func (s *ImprContext) PARDER() antlr.TerminalNode {
	return s.GetToken(ChemsPARDER, 0)
}

func (s *ImprContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ImprContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ImprContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterImpr(s)
	}
}

func (s *ImprContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitImpr(s)
	}
}

func (p *Chems) Impr() (localctx IImprContext) {
	this := p
	_ = this

	localctx = NewImprContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 38, ChemsRULE_impr)

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
		p.SetState(348)
		p.Match(ChemsPRINTLN)
	}
	{
		p.SetState(349)

		var _m = p.Match(ChemsPARIZQ)

		localctx.(*ImprContext)._PARIZQ = _m
	}
	{
		p.SetState(350)

		var _x = p.Formato()

		localctx.(*ImprContext)._formato = _x
	}
	{
		p.SetState(351)

		var _x = p.listValues(0)

		localctx.(*ImprContext)._listValues = _x
	}
	{
		p.SetState(352)
		p.Match(ChemsPARDER)
	}
	localctx.(*ImprContext).p = instruction.NewPrint(localctx.(*ImprContext).Get_listValues().GetL(), localctx.(*ImprContext).Get_formato().GetS(), localctx.(*ImprContext).Get_PARIZQ().GetLine(), localctx.(*ImprContext).Get_PARIZQ().GetColumn())

	return localctx
}

// IFormatoContext is an interface to support dynamic dispatch.
type IFormatoContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_STRING returns the _STRING token.
	Get_STRING() antlr.Token

	// Set_STRING sets the _STRING token.
	Set_STRING(antlr.Token)

	// GetS returns the s attribute.
	GetS() string

	// SetS sets the s attribute.
	SetS(string)

	// IsFormatoContext differentiates from other interfaces.
	IsFormatoContext()
}

type FormatoContext struct {
	*antlr.BaseParserRuleContext
	parser  antlr.Parser
	s       string
	_STRING antlr.Token
}

func NewEmptyFormatoContext() *FormatoContext {
	var p = new(FormatoContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_formato
	return p
}

func (*FormatoContext) IsFormatoContext() {}

func NewFormatoContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *FormatoContext {
	var p = new(FormatoContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_formato

	return p
}

func (s *FormatoContext) GetParser() antlr.Parser { return s.parser }

func (s *FormatoContext) Get_STRING() antlr.Token { return s._STRING }

func (s *FormatoContext) Set_STRING(v antlr.Token) { s._STRING = v }

func (s *FormatoContext) GetS() string { return s.s }

func (s *FormatoContext) SetS(v string) { s.s = v }

func (s *FormatoContext) STRING() antlr.TerminalNode {
	return s.GetToken(ChemsSTRING, 0)
}

func (s *FormatoContext) COMA() antlr.TerminalNode {
	return s.GetToken(ChemsCOMA, 0)
}

func (s *FormatoContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *FormatoContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *FormatoContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterFormato(s)
	}
}

func (s *FormatoContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitFormato(s)
	}
}

func (p *Chems) Formato() (localctx IFormatoContext) {
	this := p
	_ = this

	localctx = NewFormatoContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 40, ChemsRULE_formato)

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

	p.SetState(359)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 13, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(355)

			var _m = p.Match(ChemsSTRING)

			localctx.(*FormatoContext)._STRING = _m
		}
		{
			p.SetState(356)
			p.Match(ChemsCOMA)
		}
		str := (func() string {
			if localctx.(*FormatoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*FormatoContext).Get_STRING().GetText()
			}
		}())[1 : len((func() string {
			if localctx.(*FormatoContext).Get_STRING() == nil {
				return ""
			} else {
				return localctx.(*FormatoContext).Get_STRING().GetText()
			}
		}()))-1]
		localctx.(*FormatoContext).s = str

	case 2:
		p.EnterOuterAlt(localctx, 2)
		localctx.(*FormatoContext).s = ""

	}

	return localctx
}

// IMatchesContext is an interface to support dynamic dispatch.
type IMatchesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_MATCH returns the _MATCH token.
	Get_MATCH() antlr.Token

	// Set_MATCH sets the _MATCH token.
	Set_MATCH(antlr.Token)

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_casos returns the _casos rule contexts.
	Get_casos() ICasosContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_casos sets the _casos rule contexts.
	Set_casos(ICasosContext)

	// GetM returns the m attribute.
	GetM() interfaces.Instruction

	// SetM sets the m attribute.
	SetM(interfaces.Instruction)

	// IsMatchesContext differentiates from other interfaces.
	IsMatchesContext()
}

type MatchesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	m           interfaces.Instruction
	_MATCH      antlr.Token
	_expression IExpressionContext
	_casos      ICasosContext
}

func NewEmptyMatchesContext() *MatchesContext {
	var p = new(MatchesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_matches
	return p
}

func (*MatchesContext) IsMatchesContext() {}

func NewMatchesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *MatchesContext {
	var p = new(MatchesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_matches

	return p
}

func (s *MatchesContext) GetParser() antlr.Parser { return s.parser }

func (s *MatchesContext) Get_MATCH() antlr.Token { return s._MATCH }

func (s *MatchesContext) Set_MATCH(v antlr.Token) { s._MATCH = v }

func (s *MatchesContext) Get_expression() IExpressionContext { return s._expression }

func (s *MatchesContext) Get_casos() ICasosContext { return s._casos }

func (s *MatchesContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *MatchesContext) Set_casos(v ICasosContext) { s._casos = v }

func (s *MatchesContext) GetM() interfaces.Instruction { return s.m }

func (s *MatchesContext) SetM(v interfaces.Instruction) { s.m = v }

func (s *MatchesContext) MATCH() antlr.TerminalNode {
	return s.GetToken(ChemsMATCH, 0)
}

func (s *MatchesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *MatchesContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEIZQ, 0)
}

func (s *MatchesContext) Casos() ICasosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasosContext)
}

func (s *MatchesContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEDER, 0)
}

func (s *MatchesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *MatchesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *MatchesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterMatches(s)
	}
}

func (s *MatchesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitMatches(s)
	}
}

func (p *Chems) Matches() (localctx IMatchesContext) {
	this := p
	_ = this

	localctx = NewMatchesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 42, ChemsRULE_matches)

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
		p.SetState(361)

		var _m = p.Match(ChemsMATCH)

		localctx.(*MatchesContext)._MATCH = _m
	}
	{
		p.SetState(362)

		var _x = p.Expression()

		localctx.(*MatchesContext)._expression = _x
	}
	{
		p.SetState(363)
		p.Match(ChemsLLAVEIZQ)
	}
	{
		p.SetState(364)

		var _x = p.casos(0)

		localctx.(*MatchesContext)._casos = _x
	}
	{
		p.SetState(365)
		p.Match(ChemsLLAVEDER)
	}
	localctx.(*MatchesContext).m = instruction.NewMatch(localctx.(*MatchesContext).Get_expression().GetP(), localctx.(*MatchesContext).Get_casos().GetL(), localctx.(*MatchesContext).Get_MATCH().GetLine(), localctx.(*MatchesContext).Get_MATCH().GetColumn())

	return localctx
}

// ICasosContext is an interface to support dynamic dispatch.
type ICasosContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCs returns the cs rule contexts.
	GetCs() ICasosContext

	// Get_cases returns the _cases rule contexts.
	Get_cases() ICasesContext

	// Get_defaults returns the _defaults rule contexts.
	Get_defaults() IDefaultsContext

	// SetCs sets the cs rule contexts.
	SetCs(ICasosContext)

	// Set_cases sets the _cases rule contexts.
	Set_cases(ICasesContext)

	// Set_defaults sets the _defaults rule contexts.
	Set_defaults(IDefaultsContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsCasosContext differentiates from other interfaces.
	IsCasosContext()
}

type CasosContext struct {
	*antlr.BaseParserRuleContext
	parser    antlr.Parser
	l         *arrayList.List
	cs        ICasosContext
	_cases    ICasesContext
	_defaults IDefaultsContext
}

func NewEmptyCasosContext() *CasosContext {
	var p = new(CasosContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_casos
	return p
}

func (*CasosContext) IsCasosContext() {}

func NewCasosContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasosContext {
	var p = new(CasosContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_casos

	return p
}

func (s *CasosContext) GetParser() antlr.Parser { return s.parser }

func (s *CasosContext) GetCs() ICasosContext { return s.cs }

func (s *CasosContext) Get_cases() ICasesContext { return s._cases }

func (s *CasosContext) Get_defaults() IDefaultsContext { return s._defaults }

func (s *CasosContext) SetCs(v ICasosContext) { s.cs = v }

func (s *CasosContext) Set_cases(v ICasesContext) { s._cases = v }

func (s *CasosContext) Set_defaults(v IDefaultsContext) { s._defaults = v }

func (s *CasosContext) GetL() *arrayList.List { return s.l }

func (s *CasosContext) SetL(v *arrayList.List) { s.l = v }

func (s *CasosContext) Cases() ICasesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasesContext)
}

func (s *CasosContext) Defaults() IDefaultsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IDefaultsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IDefaultsContext)
}

func (s *CasosContext) Casos() ICasosContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICasosContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICasosContext)
}

func (s *CasosContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasosContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasosContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterCasos(s)
	}
}

func (s *CasosContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitCasos(s)
	}
}

func (p *Chems) Casos() (localctx ICasosContext) {
	return p.casos(0)
}

func (p *Chems) casos(_p int) (localctx ICasosContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewCasosContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx ICasosContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 44
	p.EnterRecursionRule(localctx, 44, ChemsRULE_casos, _p)

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
	p.SetState(375)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsDIFERENTE, ChemsSUB, ChemsPARIZQ, ChemsCORIZQ, ChemsP_IF, ChemsINT, ChemsFLOAT, ChemsTRUE, ChemsFALSE, ChemsMATCH, ChemsLOOP, ChemsNUMBER, ChemsDECIMAL, ChemsSTRING, ChemsID, ChemsCARACTER:
		{
			p.SetState(369)

			var _x = p.Cases()

			localctx.(*CasosContext)._cases = _x
		}
		localctx.(*CasosContext).l = arrayList.New()
		localctx.(*CasosContext).l.Add(localctx.(*CasosContext).Get_cases().GetC())

	case ChemsDEF:
		{
			p.SetState(372)

			var _x = p.Defaults()

			localctx.(*CasosContext)._defaults = _x
		}
		localctx.(*CasosContext).l = arrayList.New()
		localctx.(*CasosContext).l.Add(localctx.(*CasosContext).Get_defaults().GetC())

	default:
		panic(antlr.NewNoViableAltException(p, nil, nil, nil, nil, nil))
	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(387)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(385)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 15, p.GetParserRuleContext()) {
			case 1:
				localctx = NewCasosContext(p, _parentctx, _parentState)
				localctx.(*CasosContext).cs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_casos)
				p.SetState(377)

				if !(p.Precpred(p.GetParserRuleContext(), 4)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 4)", ""))
				}
				{
					p.SetState(378)

					var _x = p.Cases()

					localctx.(*CasosContext)._cases = _x
				}
				localctx.(*CasosContext).GetCs().GetL().Add(localctx.(*CasosContext).Get_cases().GetC())
				localctx.(*CasosContext).l = localctx.(*CasosContext).GetCs().GetL()

			case 2:
				localctx = NewCasosContext(p, _parentctx, _parentState)
				localctx.(*CasosContext).cs = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_casos)
				p.SetState(381)

				if !(p.Precpred(p.GetParserRuleContext(), 3)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 3)", ""))
				}
				{
					p.SetState(382)

					var _x = p.Defaults()

					localctx.(*CasosContext)._defaults = _x
				}
				localctx.(*CasosContext).GetCs().GetL().Add(localctx.(*CasosContext).Get_defaults().GetC())
				localctx.(*CasosContext).l = localctx.(*CasosContext).GetCs().GetL()

			}

		}
		p.SetState(389)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 16, p.GetParserRuleContext())
	}

	return localctx
}

// ICasesContext is an interface to support dynamic dispatch.
type ICasesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_conditions returns the _conditions rule contexts.
	Get_conditions() IConditionsContext

	// Get_set_match returns the _set_match rule contexts.
	Get_set_match() ISet_matchContext

	// Set_conditions sets the _conditions rule contexts.
	Set_conditions(IConditionsContext)

	// Set_set_match sets the _set_match rule contexts.
	Set_set_match(ISet_matchContext)

	// GetC returns the c attribute.
	GetC() interfaces.Cases

	// SetC sets the c attribute.
	SetC(interfaces.Cases)

	// IsCasesContext differentiates from other interfaces.
	IsCasesContext()
}

type CasesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	c           interfaces.Cases
	_conditions IConditionsContext
	_set_match  ISet_matchContext
}

func NewEmptyCasesContext() *CasesContext {
	var p = new(CasesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_cases
	return p
}

func (*CasesContext) IsCasesContext() {}

func NewCasesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CasesContext {
	var p = new(CasesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_cases

	return p
}

func (s *CasesContext) GetParser() antlr.Parser { return s.parser }

func (s *CasesContext) Get_conditions() IConditionsContext { return s._conditions }

func (s *CasesContext) Get_set_match() ISet_matchContext { return s._set_match }

func (s *CasesContext) Set_conditions(v IConditionsContext) { s._conditions = v }

func (s *CasesContext) Set_set_match(v ISet_matchContext) { s._set_match = v }

func (s *CasesContext) GetC() interfaces.Cases { return s.c }

func (s *CasesContext) SetC(v interfaces.Cases) { s.c = v }

func (s *CasesContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *CasesContext) OPMATCH() antlr.TerminalNode {
	return s.GetToken(ChemsOPMATCH, 0)
}

func (s *CasesContext) Set_match() ISet_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_matchContext)
}

func (s *CasesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CasesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CasesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterCases(s)
	}
}

func (s *CasesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitCases(s)
	}
}

func (p *Chems) Cases() (localctx ICasesContext) {
	this := p
	_ = this

	localctx = NewCasesContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 46, ChemsRULE_cases)

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
		p.SetState(390)

		var _x = p.conditions(0)

		localctx.(*CasesContext)._conditions = _x
	}
	{
		p.SetState(391)
		p.Match(ChemsOPMATCH)
	}
	{
		p.SetState(392)

		var _x = p.Set_match()

		localctx.(*CasesContext)._set_match = _x
	}
	localctx.(*CasesContext).c = interfaces.Cases{localctx.(*CasesContext).Get_conditions().GetL(), localctx.(*CasesContext).Get_set_match().GetCs().Cuerpo, localctx.(*CasesContext).Get_set_match().GetCs().Retorno, 0, false}

	return localctx
}

// IConditionsContext is an interface to support dynamic dispatch.
type IConditionsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetCond returns the cond rule contexts.
	GetCond() IConditionsContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetCond sets the cond rule contexts.
	SetCond(IConditionsContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsConditionsContext differentiates from other interfaces.
	IsConditionsContext()
}

type ConditionsContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	cond        IConditionsContext
	_expression IExpressionContext
}

func NewEmptyConditionsContext() *ConditionsContext {
	var p = new(ConditionsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_conditions
	return p
}

func (*ConditionsContext) IsConditionsContext() {}

func NewConditionsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ConditionsContext {
	var p = new(ConditionsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_conditions

	return p
}

func (s *ConditionsContext) GetParser() antlr.Parser { return s.parser }

func (s *ConditionsContext) GetCond() IConditionsContext { return s.cond }

func (s *ConditionsContext) Get_expression() IExpressionContext { return s._expression }

func (s *ConditionsContext) SetCond(v IConditionsContext) { s.cond = v }

func (s *ConditionsContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ConditionsContext) GetL() *arrayList.List { return s.l }

func (s *ConditionsContext) SetL(v *arrayList.List) { s.l = v }

func (s *ConditionsContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ConditionsContext) OR_COND() antlr.TerminalNode {
	return s.GetToken(ChemsOR_COND, 0)
}

func (s *ConditionsContext) Conditions() IConditionsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IConditionsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IConditionsContext)
}

func (s *ConditionsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ConditionsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ConditionsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterConditions(s)
	}
}

func (s *ConditionsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitConditions(s)
	}
}

func (p *Chems) Conditions() (localctx IConditionsContext) {
	return p.conditions(0)
}

func (p *Chems) conditions(_p int) (localctx IConditionsContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewConditionsContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IConditionsContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 48
	p.EnterRecursionRule(localctx, 48, ChemsRULE_conditions, _p)

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
	{
		p.SetState(396)

		var _x = p.Expression()

		localctx.(*ConditionsContext)._expression = _x
	}
	localctx.(*ConditionsContext).l = arrayList.New()
	localctx.(*ConditionsContext).l.Add(localctx.(*ConditionsContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(406)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewConditionsContext(p, _parentctx, _parentState)
			localctx.(*ConditionsContext).cond = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_conditions)
			p.SetState(399)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(400)
				p.Match(ChemsOR_COND)
			}
			{
				p.SetState(401)

				var _x = p.Expression()

				localctx.(*ConditionsContext)._expression = _x
			}
			localctx.(*ConditionsContext).GetCond().GetL().Add(localctx.(*ConditionsContext).Get_expression().GetP())
			localctx.(*ConditionsContext).l = localctx.(*ConditionsContext).GetCond().GetL()

		}
		p.SetState(408)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 17, p.GetParserRuleContext())
	}

	return localctx
}

// IDefaultsContext is an interface to support dynamic dispatch.
type IDefaultsContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_set_match returns the _set_match rule contexts.
	Get_set_match() ISet_matchContext

	// Set_set_match sets the _set_match rule contexts.
	Set_set_match(ISet_matchContext)

	// GetC returns the c attribute.
	GetC() interfaces.Cases

	// SetC sets the c attribute.
	SetC(interfaces.Cases)

	// IsDefaultsContext differentiates from other interfaces.
	IsDefaultsContext()
}

type DefaultsContext struct {
	*antlr.BaseParserRuleContext
	parser     antlr.Parser
	c          interfaces.Cases
	_set_match ISet_matchContext
}

func NewEmptyDefaultsContext() *DefaultsContext {
	var p = new(DefaultsContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_defaults
	return p
}

func (*DefaultsContext) IsDefaultsContext() {}

func NewDefaultsContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *DefaultsContext {
	var p = new(DefaultsContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_defaults

	return p
}

func (s *DefaultsContext) GetParser() antlr.Parser { return s.parser }

func (s *DefaultsContext) Get_set_match() ISet_matchContext { return s._set_match }

func (s *DefaultsContext) Set_set_match(v ISet_matchContext) { s._set_match = v }

func (s *DefaultsContext) GetC() interfaces.Cases { return s.c }

func (s *DefaultsContext) SetC(v interfaces.Cases) { s.c = v }

func (s *DefaultsContext) DEF() antlr.TerminalNode {
	return s.GetToken(ChemsDEF, 0)
}

func (s *DefaultsContext) OPMATCH() antlr.TerminalNode {
	return s.GetToken(ChemsOPMATCH, 0)
}

func (s *DefaultsContext) Set_match() ISet_matchContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ISet_matchContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ISet_matchContext)
}

func (s *DefaultsContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *DefaultsContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *DefaultsContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterDefaults(s)
	}
}

func (s *DefaultsContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitDefaults(s)
	}
}

func (p *Chems) Defaults() (localctx IDefaultsContext) {
	this := p
	_ = this

	localctx = NewDefaultsContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 50, ChemsRULE_defaults)

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
		p.SetState(409)
		p.Match(ChemsDEF)
	}
	{
		p.SetState(410)
		p.Match(ChemsOPMATCH)
	}
	{
		p.SetState(411)

		var _x = p.Set_match()

		localctx.(*DefaultsContext)._set_match = _x
	}
	localctx.(*DefaultsContext).c = interfaces.Cases{localctx.(*DefaultsContext).Get_set_match().GetCs().Condicion, localctx.(*DefaultsContext).Get_set_match().GetCs().Cuerpo, localctx.(*DefaultsContext).Get_set_match().GetCs().Retorno, 1, false}

	return localctx
}

// ISet_matchContext is an interface to support dynamic dispatch.
type ISet_matchContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Get_instruccion_wc returns the _instruccion_wc rule contexts.
	Get_instruccion_wc() IInstruccion_wcContext

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// Set_instruccion_wc sets the _instruccion_wc rule contexts.
	Set_instruccion_wc(IInstruccion_wcContext)

	// GetCs returns the cs attribute.
	GetCs() interfaces.Cases

	// SetCs sets the cs attribute.
	SetCs(interfaces.Cases)

	// IsSet_matchContext differentiates from other interfaces.
	IsSet_matchContext()
}

type Set_matchContext struct {
	*antlr.BaseParserRuleContext
	parser          antlr.Parser
	cs              interfaces.Cases
	_expression     IExpressionContext
	_instrucciones  IInstruccionesContext
	_instruccion_wc IInstruccion_wcContext
}

func NewEmptySet_matchContext() *Set_matchContext {
	var p = new(Set_matchContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_set_match
	return p
}

func (*Set_matchContext) IsSet_matchContext() {}

func NewSet_matchContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Set_matchContext {
	var p = new(Set_matchContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_set_match

	return p
}

func (s *Set_matchContext) GetParser() antlr.Parser { return s.parser }

func (s *Set_matchContext) Get_expression() IExpressionContext { return s._expression }

func (s *Set_matchContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *Set_matchContext) Get_instruccion_wc() IInstruccion_wcContext { return s._instruccion_wc }

func (s *Set_matchContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Set_matchContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *Set_matchContext) Set_instruccion_wc(v IInstruccion_wcContext) { s._instruccion_wc = v }

func (s *Set_matchContext) GetCs() interfaces.Cases { return s.cs }

func (s *Set_matchContext) SetCs(v interfaces.Cases) { s.cs = v }

func (s *Set_matchContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Set_matchContext) COMA() antlr.TerminalNode {
	return s.GetToken(ChemsCOMA, 0)
}

func (s *Set_matchContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEIZQ, 0)
}

func (s *Set_matchContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *Set_matchContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEDER, 0)
}

func (s *Set_matchContext) Instruccion_wc() IInstruccion_wcContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccion_wcContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccion_wcContext)
}

func (s *Set_matchContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Set_matchContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Set_matchContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterSet_match(s)
	}
}

func (s *Set_matchContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitSet_match(s)
	}
}

func (p *Chems) Set_match() (localctx ISet_matchContext) {
	this := p
	_ = this

	localctx = NewSet_matchContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 52, ChemsRULE_set_match)

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

	p.SetState(427)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 18, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(414)

			var _x = p.Expression()

			localctx.(*Set_matchContext)._expression = _x
		}
		{
			p.SetState(415)
			p.Match(ChemsCOMA)
		}
		arr := arrayList.New()
		arr.Add(instruction.NewElseNull("null"))
		localctx.(*Set_matchContext).cs = interfaces.Cases{arrayList.New(), arr, localctx.(*Set_matchContext).Get_expression().GetP(), 0, false}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(418)
			p.Match(ChemsLLAVEIZQ)
		}
		{
			p.SetState(419)

			var _x = p.Instrucciones()

			localctx.(*Set_matchContext)._instrucciones = _x
		}
		{
			p.SetState(420)
			p.Match(ChemsLLAVEDER)
		}
		localctx.(*Set_matchContext).cs = interfaces.Cases{arrayList.New(), localctx.(*Set_matchContext).Get_instrucciones().GetL(), expresion.NewPrimitivo(1, interfaces.INTEGER, 0, 0), 0, false}

	case 3:
		p.EnterOuterAlt(localctx, 3)
		{
			p.SetState(423)

			var _x = p.Instruccion_wc()

			localctx.(*Set_matchContext)._instruccion_wc = _x
		}
		{
			p.SetState(424)
			p.Match(ChemsCOMA)
		}
		arr := arrayList.New()
		arr.Add(localctx.(*Set_matchContext).Get_instruccion_wc().GetInstr())
		localctx.(*Set_matchContext).cs = interfaces.Cases{arrayList.New(), arr, expresion.NewPrimitivo(1, interfaces.INTEGER, 0, 0), 0, false}

	}

	return localctx
}

// IRforContext is an interface to support dynamic dispatch.
type IRforContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// Get_P_FOR returns the _P_FOR token.
	Get_P_FOR() antlr.Token

	// GetId returns the id token.
	GetId() antlr.Token

	// Set_P_FOR sets the _P_FOR token.
	Set_P_FOR(antlr.Token)

	// SetId sets the id token.
	SetId(antlr.Token)

	// Get_iter_for returns the _iter_for rule contexts.
	Get_iter_for() IIter_forContext

	// Get_instrucciones returns the _instrucciones rule contexts.
	Get_instrucciones() IInstruccionesContext

	// Set_iter_for sets the _iter_for rule contexts.
	Set_iter_for(IIter_forContext)

	// Set_instrucciones sets the _instrucciones rule contexts.
	Set_instrucciones(IInstruccionesContext)

	// GetP returns the p attribute.
	GetP() interfaces.Instruction

	// SetP sets the p attribute.
	SetP(interfaces.Instruction)

	// IsRforContext differentiates from other interfaces.
	IsRforContext()
}

type RforContext struct {
	*antlr.BaseParserRuleContext
	parser         antlr.Parser
	p              interfaces.Instruction
	_P_FOR         antlr.Token
	id             antlr.Token
	_iter_for      IIter_forContext
	_instrucciones IInstruccionesContext
}

func NewEmptyRforContext() *RforContext {
	var p = new(RforContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_rfor
	return p
}

func (*RforContext) IsRforContext() {}

func NewRforContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *RforContext {
	var p = new(RforContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_rfor

	return p
}

func (s *RforContext) GetParser() antlr.Parser { return s.parser }

func (s *RforContext) Get_P_FOR() antlr.Token { return s._P_FOR }

func (s *RforContext) GetId() antlr.Token { return s.id }

func (s *RforContext) Set_P_FOR(v antlr.Token) { s._P_FOR = v }

func (s *RforContext) SetId(v antlr.Token) { s.id = v }

func (s *RforContext) Get_iter_for() IIter_forContext { return s._iter_for }

func (s *RforContext) Get_instrucciones() IInstruccionesContext { return s._instrucciones }

func (s *RforContext) Set_iter_for(v IIter_forContext) { s._iter_for = v }

func (s *RforContext) Set_instrucciones(v IInstruccionesContext) { s._instrucciones = v }

func (s *RforContext) GetP() interfaces.Instruction { return s.p }

func (s *RforContext) SetP(v interfaces.Instruction) { s.p = v }

func (s *RforContext) P_FOR() antlr.TerminalNode {
	return s.GetToken(ChemsP_FOR, 0)
}

func (s *RforContext) P_IN() antlr.TerminalNode {
	return s.GetToken(ChemsP_IN, 0)
}

func (s *RforContext) Iter_for() IIter_forContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIter_forContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIter_forContext)
}

func (s *RforContext) LLAVEIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEIZQ, 0)
}

func (s *RforContext) Instrucciones() IInstruccionesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IInstruccionesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IInstruccionesContext)
}

func (s *RforContext) LLAVEDER() antlr.TerminalNode {
	return s.GetToken(ChemsLLAVEDER, 0)
}

func (s *RforContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *RforContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *RforContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *RforContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterRfor(s)
	}
}

func (s *RforContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitRfor(s)
	}
}

func (p *Chems) Rfor() (localctx IRforContext) {
	this := p
	_ = this

	localctx = NewRforContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 54, ChemsRULE_rfor)

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
		p.SetState(429)

		var _m = p.Match(ChemsP_FOR)

		localctx.(*RforContext)._P_FOR = _m
	}
	{
		p.SetState(430)

		var _m = p.Match(ChemsID)

		localctx.(*RforContext).id = _m
	}
	{
		p.SetState(431)
		p.Match(ChemsP_IN)
	}
	{
		p.SetState(432)

		var _x = p.Iter_for()

		localctx.(*RforContext)._iter_for = _x
	}
	{
		p.SetState(433)
		p.Match(ChemsLLAVEIZQ)
	}
	{
		p.SetState(434)

		var _x = p.Instrucciones()

		localctx.(*RforContext)._instrucciones = _x
	}
	{
		p.SetState(435)
		p.Match(ChemsLLAVEDER)
	}
	localctx.(*RforContext).p = instruction.NewFor(localctx.(*RforContext).Get_instrucciones().GetL(), (func() string {
		if localctx.(*RforContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*RforContext).GetId().GetText()
		}
	}()), localctx.(*RforContext).Get_iter_for().GetP(), localctx.(*RforContext).Get_P_FOR().GetLine(), localctx.(*RforContext).Get_P_FOR().GetColumn())

	return localctx
}

// IIter_forContext is an interface to support dynamic dispatch.
type IIter_forContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp1 returns the exp1 rule contexts.
	GetExp1() IExpressionContext

	// GetExp2 returns the exp2 rule contexts.
	GetExp2() IExpressionContext

	// SetExp1 sets the exp1 rule contexts.
	SetExp1(IExpressionContext)

	// SetExp2 sets the exp2 rule contexts.
	SetExp2(IExpressionContext)

	// GetP returns the p attribute.
	GetP() interfaces.For_Range

	// SetP sets the p attribute.
	SetP(interfaces.For_Range)

	// IsIter_forContext differentiates from other interfaces.
	IsIter_forContext()
}

type Iter_forContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.For_Range
	exp1   IExpressionContext
	exp2   IExpressionContext
}

func NewEmptyIter_forContext() *Iter_forContext {
	var p = new(Iter_forContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_iter_for
	return p
}

func (*Iter_forContext) IsIter_forContext() {}

func NewIter_forContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Iter_forContext {
	var p = new(Iter_forContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_iter_for

	return p
}

func (s *Iter_forContext) GetParser() antlr.Parser { return s.parser }

func (s *Iter_forContext) GetExp1() IExpressionContext { return s.exp1 }

func (s *Iter_forContext) GetExp2() IExpressionContext { return s.exp2 }

func (s *Iter_forContext) SetExp1(v IExpressionContext) { s.exp1 = v }

func (s *Iter_forContext) SetExp2(v IExpressionContext) { s.exp2 = v }

func (s *Iter_forContext) GetP() interfaces.For_Range { return s.p }

func (s *Iter_forContext) SetP(v interfaces.For_Range) { s.p = v }

func (s *Iter_forContext) AllPUNTO() []antlr.TerminalNode {
	return s.GetTokens(ChemsPUNTO)
}

func (s *Iter_forContext) PUNTO(i int) antlr.TerminalNode {
	return s.GetToken(ChemsPUNTO, i)
}

func (s *Iter_forContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *Iter_forContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Iter_forContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Iter_forContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Iter_forContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterIter_for(s)
	}
}

func (s *Iter_forContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitIter_for(s)
	}
}

func (p *Chems) Iter_for() (localctx IIter_forContext) {
	this := p
	_ = this

	localctx = NewIter_forContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 56, ChemsRULE_iter_for)

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

	p.SetState(447)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 19, p.GetParserRuleContext()) {
	case 1:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(438)

			var _x = p.Expression()

			localctx.(*Iter_forContext).exp1 = _x
		}
		{
			p.SetState(439)
			p.Match(ChemsPUNTO)
		}
		{
			p.SetState(440)
			p.Match(ChemsPUNTO)
		}
		{
			p.SetState(441)

			var _x = p.Expression()

			localctx.(*Iter_forContext).exp2 = _x
		}
		localctx.(*Iter_forContext).p = interfaces.For_Range{localctx.(*Iter_forContext).GetExp1().GetP(), localctx.(*Iter_forContext).GetExp2().GetP(), 0}

	case 2:
		p.EnterOuterAlt(localctx, 2)
		{
			p.SetState(444)

			var _x = p.Expression()

			localctx.(*Iter_forContext).exp1 = _x
		}
		localctx.(*Iter_forContext).p = interfaces.For_Range{localctx.(*Iter_forContext).GetExp1().GetP(), expresion.NewPrimitivo(1, interfaces.INTEGER, 0, 0), 1}

	}

	return localctx
}

// IMod_ArrayContext is an interface to support dynamic dispatch.
type IMod_ArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// Get_IGUAL returns the _IGUAL token.
	Get_IGUAL() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// Set_IGUAL sets the _IGUAL token.
	Set_IGUAL(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetP returns the p attribute.
	GetP() interfaces.Instruction

	// SetP sets the p attribute.
	SetP(interfaces.Instruction)

	// IsMod_ArrayContext differentiates from other interfaces.
	IsMod_ArrayContext()
}

type Mod_ArrayContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	p           interfaces.Instruction
	id          antlr.Token
	list        IListArrayContext
	_IGUAL      antlr.Token
	_expression IExpressionContext
}

func NewEmptyMod_ArrayContext() *Mod_ArrayContext {
	var p = new(Mod_ArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_mod_Array
	return p
}

func (*Mod_ArrayContext) IsMod_ArrayContext() {}

func NewMod_ArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *Mod_ArrayContext {
	var p = new(Mod_ArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_mod_Array

	return p
}

func (s *Mod_ArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *Mod_ArrayContext) GetId() antlr.Token { return s.id }

func (s *Mod_ArrayContext) Get_IGUAL() antlr.Token { return s._IGUAL }

func (s *Mod_ArrayContext) SetId(v antlr.Token) { s.id = v }

func (s *Mod_ArrayContext) Set_IGUAL(v antlr.Token) { s._IGUAL = v }

func (s *Mod_ArrayContext) GetList() IListArrayContext { return s.list }

func (s *Mod_ArrayContext) Get_expression() IExpressionContext { return s._expression }

func (s *Mod_ArrayContext) SetList(v IListArrayContext) { s.list = v }

func (s *Mod_ArrayContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Mod_ArrayContext) GetP() interfaces.Instruction { return s.p }

func (s *Mod_ArrayContext) SetP(v interfaces.Instruction) { s.p = v }

func (s *Mod_ArrayContext) IGUAL() antlr.TerminalNode {
	return s.GetToken(ChemsIGUAL, 0)
}

func (s *Mod_ArrayContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *Mod_ArrayContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *Mod_ArrayContext) ListArray() IListArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *Mod_ArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *Mod_ArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *Mod_ArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterMod_Array(s)
	}
}

func (s *Mod_ArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitMod_Array(s)
	}
}

func (p *Chems) Mod_Array() (localctx IMod_ArrayContext) {
	this := p
	_ = this

	localctx = NewMod_ArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 58, ChemsRULE_mod_Array)

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
		p.SetState(449)

		var _m = p.Match(ChemsID)

		localctx.(*Mod_ArrayContext).id = _m
	}
	{
		p.SetState(450)

		var _x = p.listArray(0)

		localctx.(*Mod_ArrayContext).list = _x
	}
	{
		p.SetState(451)

		var _m = p.Match(ChemsIGUAL)

		localctx.(*Mod_ArrayContext)._IGUAL = _m
	}
	{
		p.SetState(452)

		var _x = p.Expression()

		localctx.(*Mod_ArrayContext)._expression = _x
	}
	localctx.(*Mod_ArrayContext).p = instruction.NewModArray((func() string {
		if localctx.(*Mod_ArrayContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Mod_ArrayContext).GetId().GetText()
		}
	}()), expresion.NewArrayAccess((func() string {
		if localctx.(*Mod_ArrayContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*Mod_ArrayContext).GetId().GetText()
		}
	}()), localctx.(*Mod_ArrayContext).GetList().GetL(), localctx.(*Mod_ArrayContext).GetId().GetLine(), localctx.(*Mod_ArrayContext).GetId().GetColumn()), localctx.(*Mod_ArrayContext).Get_expression().GetP(), localctx.(*Mod_ArrayContext).Get_IGUAL().GetLine(), localctx.(*Mod_ArrayContext).Get_IGUAL().GetColumn())

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
	p.EnterRule(localctx, 60, ChemsRULE_expression)

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
		p.SetState(455)

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

	// Get_loops returns the _loops rule contexts.
	Get_loops() ILoopsContext

	// Get_ifs returns the _ifs rule contexts.
	Get_ifs() IIfsContext

	// Get_matches returns the _matches rule contexts.
	Get_matches() IMatchesContext

	// Get_listValues returns the _listValues rule contexts.
	Get_listValues() IListValuesContext

	// Get_arrayAcc returns the _arrayAcc rule contexts.
	Get_arrayAcc() IArrayAccContext

	// Get_creatArray returns the _creatArray rule contexts.
	Get_creatArray() ICreatArrayContext

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

	// Set_loops sets the _loops rule contexts.
	Set_loops(ILoopsContext)

	// Set_ifs sets the _ifs rule contexts.
	Set_ifs(IIfsContext)

	// Set_matches sets the _matches rule contexts.
	Set_matches(IMatchesContext)

	// Set_listValues sets the _listValues rule contexts.
	Set_listValues(IListValuesContext)

	// Set_arrayAcc sets the _arrayAcc rule contexts.
	Set_arrayAcc(IArrayAccContext)

	// Set_creatArray sets the _creatArray rule contexts.
	Set_creatArray(ICreatArrayContext)

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
	_loops      ILoopsContext
	_ifs        IIfsContext
	_matches    IMatchesContext
	_listValues IListValuesContext
	_arrayAcc   IArrayAccContext
	_creatArray ICreatArrayContext
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

func (s *Expr_aritContext) Get_loops() ILoopsContext { return s._loops }

func (s *Expr_aritContext) Get_ifs() IIfsContext { return s._ifs }

func (s *Expr_aritContext) Get_matches() IMatchesContext { return s._matches }

func (s *Expr_aritContext) Get_listValues() IListValuesContext { return s._listValues }

func (s *Expr_aritContext) Get_arrayAcc() IArrayAccContext { return s._arrayAcc }

func (s *Expr_aritContext) Get_creatArray() ICreatArrayContext { return s._creatArray }

func (s *Expr_aritContext) Get_tipo_d() ITipo_dContext { return s._tipo_d }

func (s *Expr_aritContext) SetExp(v IExpr_aritContext) { s.exp = v }

func (s *Expr_aritContext) SetOpIz(v IExpr_aritContext) { s.opIz = v }

func (s *Expr_aritContext) SetOpDe(v IExpr_aritContext) { s.opDe = v }

func (s *Expr_aritContext) Set_primitivo(v IPrimitivoContext) { s._primitivo = v }

func (s *Expr_aritContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *Expr_aritContext) Set_loops(v ILoopsContext) { s._loops = v }

func (s *Expr_aritContext) Set_ifs(v IIfsContext) { s._ifs = v }

func (s *Expr_aritContext) Set_matches(v IMatchesContext) { s._matches = v }

func (s *Expr_aritContext) Set_listValues(v IListValuesContext) { s._listValues = v }

func (s *Expr_aritContext) Set_arrayAcc(v IArrayAccContext) { s._arrayAcc = v }

func (s *Expr_aritContext) Set_creatArray(v ICreatArrayContext) { s._creatArray = v }

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

func (s *Expr_aritContext) Loops() ILoopsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ILoopsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ILoopsContext)
}

func (s *Expr_aritContext) Ifs() IIfsContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IIfsContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IIfsContext)
}

func (s *Expr_aritContext) Matches() IMatchesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IMatchesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IMatchesContext)
}

func (s *Expr_aritContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsCORIZQ, 0)
}

func (s *Expr_aritContext) ListValues() IListValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListValuesContext)
}

func (s *Expr_aritContext) CORDER() antlr.TerminalNode {
	return s.GetToken(ChemsCORDER, 0)
}

func (s *Expr_aritContext) ArrayAcc() IArrayAccContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IArrayAccContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IArrayAccContext)
}

func (s *Expr_aritContext) CreatArray() ICreatArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*ICreatArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(ICreatArrayContext)
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
	_startState := 62
	p.EnterRecursionRule(localctx, 62, ChemsRULE_expr_arit, _p)
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
	p.SetState(515)
	p.GetErrorHandler().Sync(p)
	switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 20, p.GetParserRuleContext()) {
	case 1:
		{
			p.SetState(459)
			p.Match(ChemsINT)
		}
		{
			p.SetState(460)
			p.Match(ChemsDDPUNTO)
		}
		{
			p.SetState(461)
			p.Match(ChemsPOW)
		}
		{
			p.SetState(462)

			var _m = p.Match(ChemsPARIZQ)

			localctx.(*Expr_aritContext)._PARIZQ = _m
		}
		{
			p.SetState(463)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opIz = _x
		}
		{
			p.SetState(464)
			p.Match(ChemsCOMA)
		}
		{
			p.SetState(465)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opDe = _x
		}
		{
			p.SetState(466)
			p.Match(ChemsPARDER)
		}
		localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), "^", localctx.(*Expr_aritContext).GetOpDe().GetP(), false, localctx.(*Expr_aritContext).Get_PARIZQ().GetLine(), localctx.(*Expr_aritContext).Get_PARIZQ().GetColumn())

	case 2:
		{
			p.SetState(469)
			p.Match(ChemsFLOAT)
		}
		{
			p.SetState(470)
			p.Match(ChemsDDPUNTO)
		}
		{
			p.SetState(471)
			p.Match(ChemsPOWF)
		}
		{
			p.SetState(472)

			var _m = p.Match(ChemsPARIZQ)

			localctx.(*Expr_aritContext)._PARIZQ = _m
		}
		{
			p.SetState(473)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opIz = _x
		}
		{
			p.SetState(474)
			p.Match(ChemsCOMA)
		}
		{
			p.SetState(475)

			var _x = p.expr_arit(0)

			localctx.(*Expr_aritContext).opDe = _x
		}
		{
			p.SetState(476)
			p.Match(ChemsPARDER)
		}
		localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), "^", localctx.(*Expr_aritContext).GetOpDe().GetP(), false, localctx.(*Expr_aritContext).Get_PARIZQ().GetLine(), localctx.(*Expr_aritContext).Get_PARIZQ().GetColumn())

	case 3:
		{
			p.SetState(479)

			var _m = p.Match(ChemsDIFERENTE)

			localctx.(*Expr_aritContext)._DIFERENTE = _m
		}
		{
			p.SetState(480)

			var _x = p.expr_arit(14)

			localctx.(*Expr_aritContext).opIz = _x
		}
		localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), "!", localctx.(*Expr_aritContext).GetOpIz().GetP(), true, localctx.(*Expr_aritContext).Get_DIFERENTE().GetLine(), localctx.(*Expr_aritContext).Get_DIFERENTE().GetColumn())

	case 4:
		{
			p.SetState(483)

			var _m = p.Match(ChemsSUB)

			localctx.(*Expr_aritContext)._SUB = _m
		}
		{
			p.SetState(484)

			var _x = p.expr_arit(13)

			localctx.(*Expr_aritContext).opIz = _x
		}
		localctx.(*Expr_aritContext).p = expresion.NewOperacion(localctx.(*Expr_aritContext).GetOpIz().GetP(), "°", localctx.(*Expr_aritContext).GetOpIz().GetP(), true, localctx.(*Expr_aritContext).Get_SUB().GetLine(), localctx.(*Expr_aritContext).Get_SUB().GetColumn())

	case 5:
		{
			p.SetState(487)

			var _x = p.Primitivo()

			localctx.(*Expr_aritContext)._primitivo = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_primitivo().GetP()

	case 6:
		{
			p.SetState(490)

			var _m = p.Match(ChemsPARIZQ)

			localctx.(*Expr_aritContext)._PARIZQ = _m
		}
		{
			p.SetState(491)

			var _x = p.Expression()

			localctx.(*Expr_aritContext)._expression = _x
		}
		{
			p.SetState(492)
			p.Match(ChemsPARDER)
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_expression().GetP()

	case 7:
		{
			p.SetState(495)

			var _x = p.Loops()

			localctx.(*Expr_aritContext)._loops = _x
		}
		localctx.(*Expr_aritContext).p = expresion.NewDevLoop(localctx.(*Expr_aritContext).Get_loops().GetI())

	case 8:
		{
			p.SetState(498)

			var _x = p.Ifs()

			localctx.(*Expr_aritContext)._ifs = _x
		}
		localctx.(*Expr_aritContext).p = expresion.NewDevLoop(localctx.(*Expr_aritContext).Get_ifs().GetP())

	case 9:
		{
			p.SetState(501)

			var _x = p.Matches()

			localctx.(*Expr_aritContext)._matches = _x
		}
		localctx.(*Expr_aritContext).p = expresion.NewDevLoop(localctx.(*Expr_aritContext).Get_matches().GetM())

	case 10:
		{
			p.SetState(504)
			p.Match(ChemsCORIZQ)
		}
		{
			p.SetState(505)

			var _x = p.listValues(0)

			localctx.(*Expr_aritContext)._listValues = _x
		}
		{
			p.SetState(506)
			p.Match(ChemsCORDER)
		}
		localctx.(*Expr_aritContext).p = expresion.NewArray(localctx.(*Expr_aritContext).Get_listValues().GetL())

	case 11:
		{
			p.SetState(509)

			var _x = p.ArrayAcc()

			localctx.(*Expr_aritContext)._arrayAcc = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_arrayAcc().GetP()

	case 12:
		{
			p.SetState(512)

			var _x = p.CreatArray()

			localctx.(*Expr_aritContext)._creatArray = _x
		}
		localctx.(*Expr_aritContext).p = localctx.(*Expr_aritContext).Get_creatArray().GetP()

	}
	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(550)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			p.SetState(548)
			p.GetErrorHandler().Sync(p)
			switch p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 21, p.GetParserRuleContext()) {
			case 1:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).opIz = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(517)

				if !(p.Precpred(p.GetParserRuleContext(), 12)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 12)", ""))
				}
				{
					p.SetState(518)

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
					p.SetState(519)

					var _x = p.expr_arit(13)

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
				p.SetState(522)

				if !(p.Precpred(p.GetParserRuleContext(), 11)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 11)", ""))
				}
				{
					p.SetState(523)

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
					p.SetState(524)

					var _x = p.expr_arit(12)

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
				p.SetState(527)

				if !(p.Precpred(p.GetParserRuleContext(), 10)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 10)", ""))
				}
				{
					p.SetState(528)

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
					p.SetState(529)

					var _x = p.expr_arit(11)

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
				p.SetState(532)

				if !(p.Precpred(p.GetParserRuleContext(), 9)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 9)", ""))
				}
				{
					p.SetState(533)

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
					p.SetState(534)

					var _x = p.expr_arit(10)

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
				p.SetState(537)

				if !(p.Precpred(p.GetParserRuleContext(), 16)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 16)", ""))
				}
				{
					p.SetState(538)
					p.Match(ChemsP_AS)
				}
				{
					p.SetState(539)

					var _x = p.Tipo_d()

					localctx.(*Expr_aritContext)._tipo_d = _x
				}
				localctx.(*Expr_aritContext).p = expresion.NewCast(localctx.(*Expr_aritContext).GetExp().GetP(), localctx.(*Expr_aritContext).Get_tipo_d().GetT())

			case 6:
				localctx = NewExpr_aritContext(p, _parentctx, _parentState)
				localctx.(*Expr_aritContext).exp = _prevctx
				p.PushNewRecursionContext(localctx, _startState, ChemsRULE_expr_arit)
				p.SetState(542)

				if !(p.Precpred(p.GetParserRuleContext(), 15)) {
					panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 15)", ""))
				}
				{
					p.SetState(543)

					var _m = p.Match(ChemsPUNTO)

					localctx.(*Expr_aritContext)._PUNTO = _m
				}
				{
					p.SetState(544)
					p.Match(ChemsT_STRING)
				}
				{
					p.SetState(545)

					var _m = p.Match(ChemsPARIZQ)

					localctx.(*Expr_aritContext)._PARIZQ = _m
				}
				{
					p.SetState(546)
					p.Match(ChemsPARDER)
				}
				localctx.(*Expr_aritContext).p = expresion.NewToString(localctx.(*Expr_aritContext).GetExp().GetP(), localctx.(*Expr_aritContext).Get_PUNTO().GetLine(), localctx.(*Expr_aritContext).Get_PUNTO().GetColumn())

			}

		}
		p.SetState(552)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 22, p.GetParserRuleContext())
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
	p.EnterRule(localctx, 64, ChemsRULE_primitivo)

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

	p.SetState(567)
	p.GetErrorHandler().Sync(p)

	switch p.GetTokenStream().LA(1) {
	case ChemsNUMBER:
		p.EnterOuterAlt(localctx, 1)
		{
			p.SetState(553)

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
			p.SetState(555)

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
			p.SetState(557)

			var _m = p.Match(ChemsTRUE)

			localctx.(*PrimitivoContext)._TRUE = _m
		}

		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(1, interfaces.BOOLEAN, localctx.(*PrimitivoContext).Get_TRUE().GetColumn(), localctx.(*PrimitivoContext).Get_TRUE().GetLine())

	case ChemsFALSE:
		p.EnterOuterAlt(localctx, 4)
		{
			p.SetState(559)

			var _m = p.Match(ChemsFALSE)

			localctx.(*PrimitivoContext)._FALSE = _m
		}

		localctx.(*PrimitivoContext).p = expresion.NewPrimitivo(0, interfaces.BOOLEAN, localctx.(*PrimitivoContext).Get_FALSE().GetColumn(), localctx.(*PrimitivoContext).Get_FALSE().GetLine())

	case ChemsDECIMAL:
		p.EnterOuterAlt(localctx, 5)
		{
			p.SetState(561)

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
			p.SetState(563)

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
			p.SetState(565)

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

// IListValuesContext is an interface to support dynamic dispatch.
type IListValuesContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetList returns the list rule contexts.
	GetList() IListValuesContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetList sets the list rule contexts.
	SetList(IListValuesContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListValuesContext differentiates from other interfaces.
	IsListValuesContext()
}

type ListValuesContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	list        IListValuesContext
	_expression IExpressionContext
}

func NewEmptyListValuesContext() *ListValuesContext {
	var p = new(ListValuesContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_listValues
	return p
}

func (*ListValuesContext) IsListValuesContext() {}

func NewListValuesContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListValuesContext {
	var p = new(ListValuesContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_listValues

	return p
}

func (s *ListValuesContext) GetParser() antlr.Parser { return s.parser }

func (s *ListValuesContext) GetList() IListValuesContext { return s.list }

func (s *ListValuesContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListValuesContext) SetList(v IListValuesContext) { s.list = v }

func (s *ListValuesContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListValuesContext) GetL() *arrayList.List { return s.l }

func (s *ListValuesContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListValuesContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListValuesContext) COMA() antlr.TerminalNode {
	return s.GetToken(ChemsCOMA, 0)
}

func (s *ListValuesContext) ListValues() IListValuesContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListValuesContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListValuesContext)
}

func (s *ListValuesContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListValuesContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListValuesContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterListValues(s)
	}
}

func (s *ListValuesContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitListValues(s)
	}
}

func (p *Chems) ListValues() (localctx IListValuesContext) {
	return p.listValues(0)
}

func (p *Chems) listValues(_p int) (localctx IListValuesContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListValuesContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListValuesContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 66
	p.EnterRecursionRule(localctx, 66, ChemsRULE_listValues, _p)

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
	{
		p.SetState(570)

		var _x = p.Expression()

		localctx.(*ListValuesContext)._expression = _x
	}

	localctx.(*ListValuesContext).l = arrayList.New()
	localctx.(*ListValuesContext).l.Add(localctx.(*ListValuesContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(580)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListValuesContext(p, _parentctx, _parentState)
			localctx.(*ListValuesContext).list = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_listValues)
			p.SetState(573)

			if !(p.Precpred(p.GetParserRuleContext(), 2)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 2)", ""))
			}
			{
				p.SetState(574)
				p.Match(ChemsCOMA)
			}
			{
				p.SetState(575)

				var _x = p.Expression()

				localctx.(*ListValuesContext)._expression = _x
			}

			localctx.(*ListValuesContext).GetList().GetL().Add(localctx.(*ListValuesContext).Get_expression().GetP())
			localctx.(*ListValuesContext).l = localctx.(*ListValuesContext).GetList().GetL()

		}
		p.SetState(582)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 24, p.GetParserRuleContext())
	}

	return localctx
}

// IArrayAccContext is an interface to support dynamic dispatch.
type IArrayAccContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetId returns the id token.
	GetId() antlr.Token

	// SetId sets the id token.
	SetId(antlr.Token)

	// GetList returns the list rule contexts.
	GetList() IListArrayContext

	// SetList sets the list rule contexts.
	SetList(IListArrayContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsArrayAccContext differentiates from other interfaces.
	IsArrayAccContext()
}

type ArrayAccContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.Expresion
	id     antlr.Token
	list   IListArrayContext
}

func NewEmptyArrayAccContext() *ArrayAccContext {
	var p = new(ArrayAccContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_arrayAcc
	return p
}

func (*ArrayAccContext) IsArrayAccContext() {}

func NewArrayAccContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ArrayAccContext {
	var p = new(ArrayAccContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_arrayAcc

	return p
}

func (s *ArrayAccContext) GetParser() antlr.Parser { return s.parser }

func (s *ArrayAccContext) GetId() antlr.Token { return s.id }

func (s *ArrayAccContext) SetId(v antlr.Token) { s.id = v }

func (s *ArrayAccContext) GetList() IListArrayContext { return s.list }

func (s *ArrayAccContext) SetList(v IListArrayContext) { s.list = v }

func (s *ArrayAccContext) GetP() interfaces.Expresion { return s.p }

func (s *ArrayAccContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *ArrayAccContext) ID() antlr.TerminalNode {
	return s.GetToken(ChemsID, 0)
}

func (s *ArrayAccContext) ListArray() IListArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ArrayAccContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ArrayAccContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ArrayAccContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterArrayAcc(s)
	}
}

func (s *ArrayAccContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitArrayAcc(s)
	}
}

func (p *Chems) ArrayAcc() (localctx IArrayAccContext) {
	this := p
	_ = this

	localctx = NewArrayAccContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 68, ChemsRULE_arrayAcc)

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
		p.SetState(583)

		var _m = p.Match(ChemsID)

		localctx.(*ArrayAccContext).id = _m
	}
	{
		p.SetState(584)

		var _x = p.listArray(0)

		localctx.(*ArrayAccContext).list = _x
	}
	localctx.(*ArrayAccContext).p = expresion.NewArrayAccess((func() string {
		if localctx.(*ArrayAccContext).GetId() == nil {
			return ""
		} else {
			return localctx.(*ArrayAccContext).GetId().GetText()
		}
	}()), localctx.(*ArrayAccContext).GetList().GetL(), localctx.(*ArrayAccContext).GetId().GetLine(), localctx.(*ArrayAccContext).GetId().GetColumn())

	return localctx
}

// IListArrayContext is an interface to support dynamic dispatch.
type IListArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetLista returns the lista rule contexts.
	GetLista() IListArrayContext

	// Get_expression returns the _expression rule contexts.
	Get_expression() IExpressionContext

	// SetLista sets the lista rule contexts.
	SetLista(IListArrayContext)

	// Set_expression sets the _expression rule contexts.
	Set_expression(IExpressionContext)

	// GetL returns the l attribute.
	GetL() *arrayList.List

	// SetL sets the l attribute.
	SetL(*arrayList.List)

	// IsListArrayContext differentiates from other interfaces.
	IsListArrayContext()
}

type ListArrayContext struct {
	*antlr.BaseParserRuleContext
	parser      antlr.Parser
	l           *arrayList.List
	lista       IListArrayContext
	_expression IExpressionContext
}

func NewEmptyListArrayContext() *ListArrayContext {
	var p = new(ListArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_listArray
	return p
}

func (*ListArrayContext) IsListArrayContext() {}

func NewListArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *ListArrayContext {
	var p = new(ListArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_listArray

	return p
}

func (s *ListArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *ListArrayContext) GetLista() IListArrayContext { return s.lista }

func (s *ListArrayContext) Get_expression() IExpressionContext { return s._expression }

func (s *ListArrayContext) SetLista(v IListArrayContext) { s.lista = v }

func (s *ListArrayContext) Set_expression(v IExpressionContext) { s._expression = v }

func (s *ListArrayContext) GetL() *arrayList.List { return s.l }

func (s *ListArrayContext) SetL(v *arrayList.List) { s.l = v }

func (s *ListArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsCORIZQ, 0)
}

func (s *ListArrayContext) Expression() IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *ListArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(ChemsCORDER, 0)
}

func (s *ListArrayContext) ListArray() IListArrayContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IListArrayContext)(nil)).Elem(), 0)

	if t == nil {
		return nil
	}

	return t.(IListArrayContext)
}

func (s *ListArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *ListArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *ListArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterListArray(s)
	}
}

func (s *ListArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitListArray(s)
	}
}

func (p *Chems) ListArray() (localctx IListArrayContext) {
	return p.listArray(0)
}

func (p *Chems) listArray(_p int) (localctx IListArrayContext) {
	this := p
	_ = this

	var _parentctx antlr.ParserRuleContext = p.GetParserRuleContext()
	_parentState := p.GetState()
	localctx = NewListArrayContext(p, p.GetParserRuleContext(), _parentState)
	var _prevctx IListArrayContext = localctx
	var _ antlr.ParserRuleContext = _prevctx // TODO: To prevent unused variable warning.
	_startState := 70
	p.EnterRecursionRule(localctx, 70, ChemsRULE_listArray, _p)

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
	{
		p.SetState(588)
		p.Match(ChemsCORIZQ)
	}
	{
		p.SetState(589)

		var _x = p.Expression()

		localctx.(*ListArrayContext)._expression = _x
	}
	{
		p.SetState(590)
		p.Match(ChemsCORDER)
	}

	localctx.(*ListArrayContext).l = arrayList.New()
	localctx.(*ListArrayContext).l.Add(localctx.(*ListArrayContext).Get_expression().GetP())

	p.GetParserRuleContext().SetStop(p.GetTokenStream().LT(-1))
	p.SetState(601)
	p.GetErrorHandler().Sync(p)
	_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())

	for _alt != 2 && _alt != antlr.ATNInvalidAltNumber {
		if _alt == 1 {
			if p.GetParseListeners() != nil {
				p.TriggerExitRuleEvent()
			}
			_prevctx = localctx
			localctx = NewListArrayContext(p, _parentctx, _parentState)
			localctx.(*ListArrayContext).lista = _prevctx
			p.PushNewRecursionContext(localctx, _startState, ChemsRULE_listArray)
			p.SetState(593)

			if !(p.Precpred(p.GetParserRuleContext(), 1)) {
				panic(antlr.NewFailedPredicateException(p, "p.Precpred(p.GetParserRuleContext(), 1)", ""))
			}
			{
				p.SetState(594)
				p.Match(ChemsCORIZQ)
			}
			{
				p.SetState(595)

				var _x = p.Expression()

				localctx.(*ListArrayContext)._expression = _x
			}
			{
				p.SetState(596)
				p.Match(ChemsCORDER)
			}

			localctx.(*ListArrayContext).GetLista().GetL().Add(localctx.(*ListArrayContext).Get_expression().GetP())
			localctx.(*ListArrayContext).l = localctx.(*ListArrayContext).GetLista().GetL()

		}
		p.SetState(603)
		p.GetErrorHandler().Sync(p)
		_alt = p.GetInterpreter().AdaptivePredict(p.GetTokenStream(), 25, p.GetParserRuleContext())
	}

	return localctx
}

// ICreatArrayContext is an interface to support dynamic dispatch.
type ICreatArrayContext interface {
	antlr.ParserRuleContext

	// GetParser returns the parser.
	GetParser() antlr.Parser

	// GetExp1 returns the exp1 rule contexts.
	GetExp1() IExpressionContext

	// GetExp2 returns the exp2 rule contexts.
	GetExp2() IExpressionContext

	// SetExp1 sets the exp1 rule contexts.
	SetExp1(IExpressionContext)

	// SetExp2 sets the exp2 rule contexts.
	SetExp2(IExpressionContext)

	// GetP returns the p attribute.
	GetP() interfaces.Expresion

	// SetP sets the p attribute.
	SetP(interfaces.Expresion)

	// IsCreatArrayContext differentiates from other interfaces.
	IsCreatArrayContext()
}

type CreatArrayContext struct {
	*antlr.BaseParserRuleContext
	parser antlr.Parser
	p      interfaces.Expresion
	exp1   IExpressionContext
	exp2   IExpressionContext
}

func NewEmptyCreatArrayContext() *CreatArrayContext {
	var p = new(CreatArrayContext)
	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(nil, -1)
	p.RuleIndex = ChemsRULE_creatArray
	return p
}

func (*CreatArrayContext) IsCreatArrayContext() {}

func NewCreatArrayContext(parser antlr.Parser, parent antlr.ParserRuleContext, invokingState int) *CreatArrayContext {
	var p = new(CreatArrayContext)

	p.BaseParserRuleContext = antlr.NewBaseParserRuleContext(parent, invokingState)

	p.parser = parser
	p.RuleIndex = ChemsRULE_creatArray

	return p
}

func (s *CreatArrayContext) GetParser() antlr.Parser { return s.parser }

func (s *CreatArrayContext) GetExp1() IExpressionContext { return s.exp1 }

func (s *CreatArrayContext) GetExp2() IExpressionContext { return s.exp2 }

func (s *CreatArrayContext) SetExp1(v IExpressionContext) { s.exp1 = v }

func (s *CreatArrayContext) SetExp2(v IExpressionContext) { s.exp2 = v }

func (s *CreatArrayContext) GetP() interfaces.Expresion { return s.p }

func (s *CreatArrayContext) SetP(v interfaces.Expresion) { s.p = v }

func (s *CreatArrayContext) CORIZQ() antlr.TerminalNode {
	return s.GetToken(ChemsCORIZQ, 0)
}

func (s *CreatArrayContext) PTCOMA() antlr.TerminalNode {
	return s.GetToken(ChemsPTCOMA, 0)
}

func (s *CreatArrayContext) CORDER() antlr.TerminalNode {
	return s.GetToken(ChemsCORDER, 0)
}

func (s *CreatArrayContext) AllExpression() []IExpressionContext {
	var ts = s.GetTypedRuleContexts(reflect.TypeOf((*IExpressionContext)(nil)).Elem())
	var tst = make([]IExpressionContext, len(ts))

	for i, t := range ts {
		if t != nil {
			tst[i] = t.(IExpressionContext)
		}
	}

	return tst
}

func (s *CreatArrayContext) Expression(i int) IExpressionContext {
	var t = s.GetTypedRuleContext(reflect.TypeOf((*IExpressionContext)(nil)).Elem(), i)

	if t == nil {
		return nil
	}

	return t.(IExpressionContext)
}

func (s *CreatArrayContext) GetRuleContext() antlr.RuleContext {
	return s
}

func (s *CreatArrayContext) ToStringTree(ruleNames []string, recog antlr.Recognizer) string {
	return antlr.TreesStringTree(s, ruleNames, recog)
}

func (s *CreatArrayContext) EnterRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.EnterCreatArray(s)
	}
}

func (s *CreatArrayContext) ExitRule(listener antlr.ParseTreeListener) {
	if listenerT, ok := listener.(ChemsListener); ok {
		listenerT.ExitCreatArray(s)
	}
}

func (p *Chems) CreatArray() (localctx ICreatArrayContext) {
	this := p
	_ = this

	localctx = NewCreatArrayContext(p, p.GetParserRuleContext(), p.GetState())
	p.EnterRule(localctx, 72, ChemsRULE_creatArray)

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
		p.SetState(604)
		p.Match(ChemsCORIZQ)
	}
	{
		p.SetState(605)

		var _x = p.Expression()

		localctx.(*CreatArrayContext).exp1 = _x
	}
	{
		p.SetState(606)
		p.Match(ChemsPTCOMA)
	}
	{
		p.SetState(607)

		var _x = p.Expression()

		localctx.(*CreatArrayContext).exp2 = _x
	}
	{
		p.SetState(608)
		p.Match(ChemsCORDER)
	}
	localctx.(*CreatArrayContext).p = expresion.NewArray2(localctx.(*CreatArrayContext).GetExp1().GetP(), localctx.(*CreatArrayContext).GetExp2().GetP())

	return localctx
}

func (p *Chems) Sempred(localctx antlr.RuleContext, ruleIndex, predIndex int) bool {
	switch ruleIndex {
	case 15:
		var t *Llaves_ifContext = nil
		if localctx != nil {
			t = localctx.(*Llaves_ifContext)
		}
		return p.Llaves_if_Sempred(t, predIndex)

	case 22:
		var t *CasosContext = nil
		if localctx != nil {
			t = localctx.(*CasosContext)
		}
		return p.Casos_Sempred(t, predIndex)

	case 24:
		var t *ConditionsContext = nil
		if localctx != nil {
			t = localctx.(*ConditionsContext)
		}
		return p.Conditions_Sempred(t, predIndex)

	case 31:
		var t *Expr_aritContext = nil
		if localctx != nil {
			t = localctx.(*Expr_aritContext)
		}
		return p.Expr_arit_Sempred(t, predIndex)

	case 33:
		var t *ListValuesContext = nil
		if localctx != nil {
			t = localctx.(*ListValuesContext)
		}
		return p.ListValues_Sempred(t, predIndex)

	case 35:
		var t *ListArrayContext = nil
		if localctx != nil {
			t = localctx.(*ListArrayContext)
		}
		return p.ListArray_Sempred(t, predIndex)

	default:
		panic("No predicate with index: " + fmt.Sprint(ruleIndex))
	}
}

func (p *Chems) Llaves_if_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 0:
		return p.Precpred(p.GetParserRuleContext(), 2)

	case 1:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Casos_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 2:
		return p.Precpred(p.GetParserRuleContext(), 4)

	case 3:
		return p.Precpred(p.GetParserRuleContext(), 3)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Conditions_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 4:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) Expr_arit_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 5:
		return p.Precpred(p.GetParserRuleContext(), 12)

	case 6:
		return p.Precpred(p.GetParserRuleContext(), 11)

	case 7:
		return p.Precpred(p.GetParserRuleContext(), 10)

	case 8:
		return p.Precpred(p.GetParserRuleContext(), 9)

	case 9:
		return p.Precpred(p.GetParserRuleContext(), 16)

	case 10:
		return p.Precpred(p.GetParserRuleContext(), 15)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) ListValues_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 11:
		return p.Precpred(p.GetParserRuleContext(), 2)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}

func (p *Chems) ListArray_Sempred(localctx antlr.RuleContext, predIndex int) bool {
	this := p
	_ = this

	switch predIndex {
	case 12:
		return p.Precpred(p.GetParserRuleContext(), 1)

	default:
		panic("No predicate with index: " + fmt.Sprint(predIndex))
	}
}
