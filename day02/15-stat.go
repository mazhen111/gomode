package main
import "fmt"

//映射练习
func main() {
	// 投票结果
	names := []string{
		"小明",
		"小红",
		"小花",
		"小黑",
		"大黄",
		"小明",
		"大黄",
		"小黑",
		"大黄",
		"大黄",
		"大黄",
		"大黄",
		"大黄",
		"大黄",


	}
	/*
		小明 1+1
		小红 1
		小花 1
		小黑 1+1
		大黄 1+1+1+1
	*/
	var stat =make(map[string]int)
	for _,i :=range names {
		v ,ok := stat[i]
		if !ok {
			stat[i]=1

		}else {

			stat[i]=v+1
		}
	}
	fmt.Println(stat)
}
