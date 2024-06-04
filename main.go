// Hello world
package main

import (
	"fmt"

	_ "github.com/xxliu95/learningGo/HelloWorld"
	"github.com/xxliu95/learningGo/Random"
)

func main() {
	// fmt.Println("你好，世界！")
	fmt.Println(Random.RandRunes(12, "鼠牛虎兔龙蛇马羊猴鸡狗猪"))
}
