package main

import ( "./ascii"
	"./iso"
)

func main()  {
	//ascii.IterateOverASCIIStringLiteral(ascii.Ascii)
	ascii.GreetingASCII()
	sl := iso.CreateExtendedASCII()
	iso.IterateOverASCIIStringLiteral(sl)



}
