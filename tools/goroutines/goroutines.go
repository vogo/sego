/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// 测试sego并行分词速度

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"time"

	"github.com/vogo/sego"
)

var (
	segmenter  = sego.Segmenter{}
	numThreads = runtime.NumCPU()
	task       = make(chan []byte, numThreads*40)
	done       = make(chan bool, numThreads)
	numRuns    = 50
)

func worker() {
	for line := range task {
		segmenter.Segment(line)
	}
	done <- true
}

func main() {
	// 将线程数设置为CPU数
	runtime.GOMAXPROCS(numThreads)

	// 载入词典
	segmenter.LoadDictionary("../data/dictionary.txt")

	// 打开将要分词的文件
	file, err := os.Open("../testdata/bailuyuan.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// 逐行读入
	scanner := bufio.NewScanner(file)
	size := 0
	lines := [][]byte{}
	for scanner.Scan() {
		var text string
		fmt.Sscanf(scanner.Text(), "%s", &text)
		content := []byte(text)
		size += len(content)
		lines = append(lines, content)
	}

	// 启动工作线程
	for i := 0; i < numThreads; i++ {
		go worker()
	}
	log.Print("开始分词")

	// 记录时间
	t0 := time.Now()

	// 并行分词
	for i := 0; i < numRuns; i++ {
		for _, l := range lines {
			task <- l
		}
	}
	close(task)

	// 确保分词完成
	for i := 0; i < numThreads; i++ {
		<-done
	}

	// 记录时间并计算分词速度
	t1 := time.Now()
	log.Printf("分词花费时间 %v", t1.Sub(t0))
	log.Printf("分词速度 %f MB/s", float64(size*numRuns)/t1.Sub(t0).Seconds()/(1024*1024))
}
