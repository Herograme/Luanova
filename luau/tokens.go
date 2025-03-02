// Copyright 2024 Luau contributors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package syntax

type Token uint

type token = Token

const (
	_ token = iota

	// Special tokens
	_EOF     // EOF
	_Name    // name
	_Number  // number
	_String  // string
	_Comment // comment

	// Operators
	_Add    // +
	_Sub    // -
	_Mul    // *
	_Div    // /
	_Mod    // %
	_Pow    // ^
	_Concat // ..
	_Len    // #
	_Eq     // ==
	_Ne     // ~=
	_Le     // <=
	_Ge     // >=
	_Lt     // <
	_Gt     // >
	_And    // and
	_Or     // or
	_Not    // not

	// Assignment operators
	_Assign       // =
	_AddAssign    // +=
	_SubAssign    // -=
	_MulAssign    // *=
	_DivAssign    // /=
	_ModAssign    // %=
	_PowAssign    // ^=
	_ConcatAssign // ..=

	// Delimiters
	_Lparen   // (
	_Rparen   // )
	_Lbrace   // {
	_Rbrace   // }
	_Lbrack   // [
	_Rbrack   // ]
	_Comma    // ,
	_Semi     // ;
	_Colon    // :
	_Dot      // .
	_Dots     // ...
	_Arrow    // ->
	_Question // ?

	// Keywords
	_And
	_Break
	_Continue
	_Do
	_Else
	_Elseif
	_End
	_False
	_For
	_Function
	_If
	_In
	_Local
	_Nil
	_Not
	_Or
	_Repeat
	_Return
	_Then
	_True
	_Until
	_While

	// Luau-specific keywords
	_Type
	_Export
	_Enum
	_Interface
	_As
	_Is
	_Where

	tokenCount
)
