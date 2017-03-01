package iso

import "fmt"

func CreateExtendedASCII() string{
	var extascii []byte;
	fmt.Println(extascii)

	// iterate over extended ascii
	for i := 0x80; i <= 0xFF; i++{
		extascii = append(extascii, byte(i)) // append byte to our array
	}

	// print out
	fmt.Printf("%q", extascii)

	return string(extascii) // konverterer byten til string

}


func IterateOverASCIIStringLiteral(sl string) {
	for i := 0; i < len(sl); i++ {
		c := []byte {sl[i]}
		fmt.Printf("%X %s %b\n", c, c, sl[i])
	}
}

// Kode for Oppgave 2b
func GreetingExtendedASCII(hello []byte) string{
	var str string;
	var strbytes []byte;
	for _, c := range hello {
		strbytes = append(strbytes, c)

	}

	str = string(strbytes)

	return str
}
