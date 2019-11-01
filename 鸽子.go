package main

import
	"fmt"


type person struct{
	name  string
	age   int
	gender  string
	saying  string
}
type pigeon interface{
     gugugege()
}
type repeater interface{
     repeat(string)
}
type lemonmonster interface {
     suan(string)
}
type smellgood interface {
     xiang (string)
}
func (p *person) guguge(){
	fmt.Print(p.name,"鸽死你")
}
func (p *person) repeat(word string){
	fmt.Print(word)
}
func (p *person) suan(){
	fmt.Print(p.gender,"我酸吗")
}
func (p *person) xiang(){
	fmt.Print(p.saying,"今天从外面跳下去也不会吃你们家一口饭")
}
func main(){

	p:= &person {
		name:  "傻叉",
		age :   18,
		gender:  "你猜猜",
		saying:  "我王境泽",
	}
	p.guguge()
	p.repeat("哈哈哈哈" )
	p.suan()
	p.xiang()

}