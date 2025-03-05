## Basics
  1. every go file starts with a package clause
     a package is a collection of code that al does similar things
  2. There are import statements, which help import the other packages 
     package line says all the rest of the code in this file belongs to the main package
     the import statement indicates that we will be importing a package and utilizing some functionalities from that package.
     "A place for everything and everything in its place."
  3. The `go fmt` command can be used to format the go file such that the code is clearly formatted.
  4. For supporting multiple imports use the parentheses and add the packages space/newline seperated.
  5. Strings are series of bytes that usually represent the text characters.
  ```go
  \\ string literal
  "hello there mate!"
  \\ characters like newlines, tabs and other characters that would be hard to include in a program code can be represented with escape characters.
  \\ a backslash followed by a character is the way to represent the special char
  "hi there \t adp.\nThe day is wonderful."
  ```
  6. The go runes are used to represent the single characters.
  String literals are written surrounded by double quotation marks \"
  Rune Literals are written surrounded by single quotation marks \'
  7. Go programs can use almost any character from almost any language, because go uses the unicode standard for storing runes, runes are kept as numeric codes.
  Runes are stored as numeric codes
  when printing the runes, we would notice the numeric values instead of the characters
  fmt.Println('A') --> 65
  8. Booleans