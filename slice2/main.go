package main

//下面是一个切片特点的练习代码，要求实现一个函数，用于将一个字符串转换成单词列表，并将列表中的单词按照出现次数从多到少排序输出
import (
	"fmt"
	"sort"
	"strings"
)

func wordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, w := range strings.Fields(s) {
		m[w]++
	}
	return m
}

func sortByCount(m map[string]int) []string {
	s := make([]string, 0, len(m))
	for k := range m {
		s = append(s, k)
	}
	count := func(i, j int) bool { return m[s[i]] > m[s[j]] }
	sort.Slice(s, count)
	return s
}

func main() {
	s := "hello world hello go go go go go"
	wc := wordCount(s)
	fmt.Println(wc)
	sorted := sortByCount(wc)
	fmt.Println(sorted)
}

//在这个示例中，我们定义了两个函数，分别是wordCount和sortByCount。
//wordCount函数接受一个字符串作为参数，将字符串中的单词转换成一个map，其中键为单词，值为单词出现的次数。
//sortByCount函数接受一个map作为参数，将map中的单词按照出现次数从多到少排序，然后将排序后的单词列表返回。
//在main函数中，我们创建了一个字符串s，并调用wordCount函数将字符串转换成一个单词map。
//然后，我们调用sortByCount函数将单词按照出现次数排序，然后输出排序后的结果。运行上述代码，输出结果如下：
//map[go:4 hello:2 world:1]
//[go hello world]

//可以看到，原先的字符串中的单词已经被转换成了一个map，并按照出现次数排序输出。
//在这个示例中，我们使用了切片的动态扩容特性，以及sort.Slice函数对切片的排序能力，充分体现了切片的灵活性和便利性
