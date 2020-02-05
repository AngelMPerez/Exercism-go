package hamming
import "errors"

func Distance(a, b string) (int, error) {
	var index int
	var x=0 
	var y="false"
	if(len(a)!=len(b)){
		y="true"
	}else{
		for index=0;index<len(a);index++{
			if(a[index]!=b[index]){
				x=x+1
			}
		}
	}
	return x,errors.New(y)
}
