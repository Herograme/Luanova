// Copyright 2024 Luau contributors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package luau

const (
	Illegal = iota
	EOF

	Comment // //
	CommentBlock // /**/
	
	Literal // Literal

	// Keywords
	Function // Function
	Local    // Local
	If       // If
	ElseIf   // ElseIf
	Else     // Else
	While    // While
	For      // For
	End      // End
	Then     // Then
	Repeat   // Repeat
	Continue // Continue
	Break    // Break
	In       // In
	Return   // Return
	Do       // Do

	// Operators
	Not          // Not
	NotEqual     // ~=
	Or           // Or
	Greater      // <
	Less         // >
	GreaterEqual // <=
	LessEqual    // >=
	And          // And
	Assign       // =
	Equal        // ==
	PlusAssign   // +=
	SubAssign    // -=
	MultiAssign  // *=
	DivAssign    // /=
	ModAssign    // %=
	Declaration  // =
	Plus         // +
	Sub          // -
	Multi        // *
	Div          // /
	Mod          // %
	Po           // ^
	Concat       // ..

	//Delimiters
	LParen   // (
	RParen   // )
	LBrace   // {
	RBrace   // }
	LBrack   // [
	RBrack   // ]
	Comma    // ,
	Semi     // ;
	Colom    // :
	Dot      // .
	Dots     // ...
	Arrow    // ->
	Question // ?

	// Primitive Types
	String // string
	Int    // int
	Float  // float
)
