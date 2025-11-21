sego - https://github.com/vogo/sego
====

> Fork from:  https://github.com/huichen/sego

Go中文分词

<a href="dictionary.go">词典</a>用双数组trie（Double-Array Trie）实现，
<a href="segmenter.go">分词器</a>算法为基于词频的最短路径加动态规划。

支持普通和搜索引擎两种分词模式，支持用户词典、词性标注，可运行<a href="server/server.go">JSON RPC服务</a>。

分词速度<a href="tools/benchmark/benchmark.go">单线程</a>9MB/s，<a href="tools/goroutines/goroutines.go">goroutines并发</a>42MB/s（8核Macbook Pro）。

# 安装/更新

```
go get -u github.com/vogo/sego
```

# 使用


```go
package main

import (
	"fmt"
	"github.com/vogo/sego"
)

func main() {
	// 载入词典
	var segmenter sego.Segmenter
	segmenter.Load("/opt/data/user.txt,/opt/data/dictionary.txt", []*sego.ConfigToken{
		{Text: "猫猫", Frequency: 2, Pos: "n"},
		{Text: "狗狗", Frequency: 2, Pos: "n"},
	})

	// 分词
	text := []byte("中华人民共和国中央人民政府")
	segments := segmenter.Segment(text)
  
	// 处理分词结果
	// 支持普通模式和搜索模式两种分词，见代码中SegmentsToString函数的注释。
	fmt.Println(sego.SegmentsToString(segments, false)) 
}
```
