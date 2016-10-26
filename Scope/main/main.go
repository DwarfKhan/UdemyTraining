package main
import "fmt"
func main(){
	x := 12
	fmt.Println(x)
	duck()
}
func duck(){
	fmt.Println("suckamyducka")
	fmt.Println("But ",y," tho")
	junk()
}
func junk() {
	var x int = 14
	fmt.Println(x)
}
var y string = "why"
