//A tuple is defined as a finite sorted list of elements
package main

import "fmt"

//IProcess interface
type IProcess interface{
	process()
}
//Adapter struct
type Adapter struct{
	adaptee Adaptee
}
//Adapter class method process
func (adapter Adapter)process(){
fmt.Println("Adapter process")
adapter.adaptee.convert()
}

//Adaptee struct
type Adaptee struct{
	adapterType int
}
//Adaptee class method convert
func (adaptee Adaptee) convert(){
	fmt.Println("Adaptee convert method")
}

	
func main(){

var processor IProcess = Adapter{}
processor.process()

	
}
/* Result
Adapter process
Adaptee convert method

*/
