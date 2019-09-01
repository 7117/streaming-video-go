package main
 
import(
	"fmt"
	"reflect"
)
 
type b struct{
	c int
	d string
}
 
type bb struct{
	ff b
	dd string
}
 
type bbb struct{
	// 这句话是在表明
	// c是一个引用传值
	// b对应地址
	c *b
	dd string
}
 
func main(){
	// 这个是结构体嵌套进行引用传值  
	// 因为此时在bbb中b是一个地址&{}
	// c后面是结构体进行赋值  是对
	// 此处的b只是针对结构体b的 
	ggggg:=bbb{c:&b{c:1,d:"gggg"},dd:"gggg"}
	fmt.Println(reflect.TypeOf(ggggg.c));//*main.b
	fmt.Println(reflect.TypeOf(*(ggggg.c)));//main.b
	fmt.Println(ggggg.c);				//&{1 gggg}
	fmt.Println(ggggg);					// {0xc04204c420 gggg}
 
	// 这个是结构体嵌套进行引用传值  
	// 这样是错误的哈，在bbb里面b是引用传值  直接用值传递会报错
	// cannot use b literal (type b) as type *b in field value
	// 不能在字段值中使用b文字(类型b)作为类型*b
	// 就是说我们用的是b类型  在结构体声明的时候是*b  所以错误
	// gggggg:=bbb{c:b{c:1,d:"gggg"},dd:"gggg"}
	// fmt.Println(gggggg);				//错误

	// invalid indirect of b literal (type b) b字母无法指向b类)
	// ggggggg:=bbb{c:*b{c:1,d:"gggg"},dd:"gggg"}
	// fmt.Println(ggggggg);			//错误

	// 结构体的定义要么是值传递(方式一)      要么是引用传递使用* 对类型定义（方式二）
	// 结构体的赋值要么是什么值传递(方式一)  要么是引用传递使用&类型名字{id:1,name:"aa"}（方式二）
	// 方式一对应方式一；方式二对应方式二
}
 
// 0xc042052058
// 1
// &{1 gggg}
// {1 gggg}
// {{1 gggg} gggg}
// *main.b
// &{1 gggg}
// {0xc04204c420 gggg}