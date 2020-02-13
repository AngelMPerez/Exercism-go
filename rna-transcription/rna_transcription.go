package strand

func ToRNA(dna string) string {

	var x="" 
	for i:=0;i<len(dna);i++{
		s := string([]byte{dna[i]})
			switch{
			case s == "A": x= "U"
			case s == "C": x= "G"
			case s == "T": x= "A"
			case s == "G": x= "C"
		}
	}	
	return x
}
