package strand
import "bytes"

func ToRNA(dna string) string {

	var x bytes.Buffer
	for i:=0;i<len(dna);i++{
		s := string([]byte{dna[i]})
			switch{
			case s == "A": x.WriteString("U")
			case s == "C": x.WriteString("G")
			case s == "T": x.WriteString("A")
			case s == "G": x.WriteString("C")
		}
	}	
	return x.String()
}
