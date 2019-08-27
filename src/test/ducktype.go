package main
import (
	"fmt"
	"reflect"
)

// 如果某个变量 t（person） 的类型 T（Person） 实现了某个接口 I（ISayHello） 所要求的所有方法，
// 那么这个变量 t（person） 就能被赋值给 I（ISayHello） 的接口变量 i。
// 调用 i 的方法，
// 最终就是调用 t（person） 的方法
type ISayHello interface {
	// 这里只是保证会被赋值  
	// 类型a的变量a可以赋值给类型b的变量b
    SayHello()
}
func greeting(i ISayHello) {
	fmt.Println("本身interface");
	// 此时这个i就已经等于person了  就是person的SayHello()方法
	i.SayHello()
	fmt.Println(reflect.TypeOf(i));//main.Person
	fmt.Println("本身interface");
	// 本身interface
	// Hello!
	// 本身interface
}

type Person struct {
	name string
}
func (person Person) SayHello() {
    fmt.Printf("Hello!")
}

func main () {
	// person是一个变量  
	// Person是一种数据类型  表明person是一个Person类型的变量
	// 结构体的赋值  就是这样  
	// 前面是结构体的名字
	// 后面{}是对属性进行赋值
	person := Person{name:"duck type"}
	fmt.Println(person.name,person);//person.name是一个字符串  person是一个结构体 ;duck type {duck type}
	
	var i ISayHello 	// 一个接口  变量为i
	// 因为person实现了ISayHello接口的所有方法  
	// 所以person能够赋值给i
	// 什么时候能够赋值呢？就是一个接口与一个变量的方法都一样的时候 
	// 就可以进行赋值
	// 这里 i一被赋值  i的类型变了  变成了Person
	i = person			// 这个赋值是在说：person变量赋值给了变量i
	fmt.Println(i);		// {duck type}
	fmt.Println(reflect.TypeOf(i));		//main.Person
	fmt.Println(reflect.TypeOf(person));//main.Person
	greeting(i)			// 最后输出： Hello! 表明i有了person的属性  是Person数据类型了
	person.SayHello() 	//Hello!
	i.SayHello();		//Hello!
	
}

// 第一，类型 T 不需要显式地声明它实现了接口 I（就是说不需要Person=ISayHello），
// 只要类型 T 实现了所有接口 I 规定的函数，它就自动地实现了接口 I。 
// 这样就像动态语言一样省了很多代码，少了许多限制。

// 第二，在把 duck 或者 person 传递给 greeting 前，
// 需要显式或者隐式地把它们转换为接口 I 的接口变量 i。
// 这样就可以和其它静态类型语言一样，在编译时检查参数的合法性。