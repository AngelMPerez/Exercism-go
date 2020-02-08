package reverse
import "bytes"


func Reverse(x string)(string){
	var y bytes.Buffer
	// var index int
	if(len(x)==0){
		return x
	}else{
		for i := len(x)-1; i >= 0; i-- {
			s := string([]byte{x[i]})
			y.WriteString(s)
		}
		return y.String()
	}
}