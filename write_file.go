package main
import ( 
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)
func check(e error){
if e != nil{
	panic(e)	
   }
}

func main(){
	d1 := []byte("hello\nGo\nI\nam\nRashmi\n")
	err := ioutil.WriteFile("/tmp/dat1",d1,0644)
	check(err)
	
	f,err := os.Create("/tmp/dat2")
	check(err)

	defer f.Close()
	
	d2 := []byte{10,20,30,40,50}
	n2,err := f.Write(d2)
	check(err)
	fmt.Printf("wrote %d bytes\n",n2)
	
		
	n3,err := f.WriteString("writes\n")
	fmt.Printf("wrote %d bytes\n",n3)

	f.Sync()
	
	w := bufio.NewWriter(f)
	n4,err := w.WriteString("I\nlove\nyou\nGo\n")
	fmt.Printf("wrote %d bytes\n",n4)

	w.Flush()
}
