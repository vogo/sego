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

package sego

import (
	"fmt"
	"testing"
)

// expect 检查实际值是否与期望值相等，不相等时记录错误
func expect(t *testing.T, expect string, actual interface{}) {
	actualString := fmt.Sprint(actual)
	if expect != actualString {
		t.Errorf("期待值=\"%s\", 实际=\"%s\"", expect, actualString)
	}
}

// printTokens 将 Token 列表转换为字符串输出
func printTokens(tokens []*Token, numTokens int) (output string) {
	for iToken := 0; iToken < numTokens; iToken++ {
		for _, word := range tokens[iToken].text {
			output += fmt.Sprint(string(word))
		}
		output += " "
	}
	return
}

// toWords 将字符串切片转换为 Text 切片
func toWords(strings ...string) []Text {
	words := []Text{}
	for _, s := range strings {
		words = append(words, []byte(s))
	}
	return words
}

// bytesToString 将 Text 切片转换为字符串，单词之间用 "/" 分隔
func bytesToString(bytes []Text) (output string) {
	for _, b := range bytes {
		output += (string(b) + "/")
	}
	return
}
