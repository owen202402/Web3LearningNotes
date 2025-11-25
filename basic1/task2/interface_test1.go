/*
✅面向对象
1.题目 ：定义一个 Shape 接口，包含 Area() 和 Perimeter() 两个方法。然后创建 Rectangle 和 Circle 结构体，实现 Shape 接口。在主函数中，创建这两个结构体的实例，并调用它们的 Area() 和 Perimeter() 方法。
考察点 ：接口的定义与实现、面向对象编程风格。
2.题目 ：使用组合的方式创建一个 Person 结构体，包含 Name 和 Age 字段，再创建一个 Employee 结构体，组合 Person 结构体并添加 EmployeeID 字段。为 Employee 结构体实现一个 PrintInfo() 方法，输出员工的信息。
考察点 ：组合的使用、方法接收者。
*/

package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	l int
	w int
}

type Circle struct {
	r int
}

func (rect *Rectangle) Area() float64 {
	fmt.Printf("Rectangle, l:%d w:%d \n", rect.l, rect.w)
	return float64(rect.l * rect.w)
}

func (rect *Rectangle) Perimeter() float64 {
	fmt.Printf("Rectangle, l:%d w:%d \n", rect.l, rect.w)
	return float64(rect.l+rect.w) * 2
}

func (cir *Circle) Area() float64 {
	fmt.Print("Circle, r ", cir.r, "  ")
	return float64(cir.r*cir.r) * math.Pi
}

func (cir *Circle) Perimeter() float64 {
	fmt.Print("Circle, r ", cir.r, "  ")
	return float64(cir.r) * 2 * math.Pi
}

func PrintArea(s Shape) {
	fmt.Println("Area: ", s.Area())
}

func PrintPerimeter(s Shape) {
	fmt.Println("Perimeter: ", s.Perimeter())
}

func main() {
	circle := &Circle{r: 3}
	rectangle := &Rectangle{l: 5, w: 2}
	PrintPerimeter(circle)
	PrintPerimeter(rectangle)
	PrintArea(circle)
	PrintArea(rectangle)

}
