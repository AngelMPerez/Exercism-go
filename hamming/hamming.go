package hamming
import "errors"

func Distance(a, b string) (int, error) {
	var index int
	var x=0 
	
	if(len(a)!=len(b)){
		return x, errors.New("error")
	}else{
		for index=0;index<len(a);index++{
			if(a[index]!=b[index]){
				x=x+1
			}
		}
		return x,nil
	}
}
