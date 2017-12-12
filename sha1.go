package main
import "crypto/sha1"
import "fmt"

func main(){
	s := "I am Rashmi Nagpal"
	h := sha1.New()
		
	h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Println(s)
	fmt.Printf("%x\n",bs)
}
