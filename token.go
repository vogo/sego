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

// 字串类型，可以用来表达
//  1. 一个字元，比如"中"又如"国", 英文的一个字元是一个词
//  2. 一个分词，比如"中国"又如"人口"
//  3. 一段文字，比如"中国有十三亿人口"
type Text []byte

// Token 表示一个分词
type Token struct {
	// 分词的字串，这实际上是个字元数组
	text []Text

	// 分词在语料库中的词频
	frequency int

	// log2(总词频/该分词词频)，这相当于log2(1/p(分词))，用作动态规划中
	// 该分词的路径长度。求解prod(p(分词))的最大值相当于求解
	// sum(distance(分词))的最小值，这就是“最短路径”的来历。
	distance float32

	// 词性标注
	// 沿用中文词性集（如 ICTCLAS/人民日报标注集），
	// 如： n (名词)、 v (动词)、 a (形容词)、 d (副词)、 m (数词)、 q (量词)、 r (代词)、 p (介词)、
	// c (连词)、 u (助词)、 t (时间词)、 s (处所词)、 f (方位词)、 i (成语)、 l (习用语)、 j (简称)、
	// h (前缀)、 k (后缀)、 g (语素)、 x (字符串/未知)、 w (标点)、 z (状态词)、 b (区别词)
	// 词典行格式通常为： 词语 频次 词性 ；不写词性时只给前两列，库会把词性置为空
	// 分词路径选择不使用词性，主要由词频和词典匹配驱动；词性仅作为 Token.Pos() 的附加信息输出
	// 词性不会改变切分边界或得分；是否提供词性标签对最终切分结果无直接影响
	// 可以在分词后基于词性做结果过滤/分组（例如保留名词、去掉助词），但这是后处理逻辑，不改变分词器内部决策
	pos string

	// 该分词文本的进一步分词划分，见Segments函数注释。
	segments []*Segment
}

// Text 返回分词文本
func (token *Token) Text() string {
	return textSliceToString(token.text)
}

// Frequency 返回分词在语料库中的词频
func (token *Token) Frequency() int {
	return token.frequency
}

// Pos 返回分词的词性标注
func (token *Token) Pos() string {
	return token.pos
}

// 该分词文本的进一步分词划分，比如"中华人民共和国中央人民政府"这个分词
// 有两个子分词"中华人民共和国"和"中央人民政府"。子分词也可以进一步有子分词
// 形成一个树结构，遍历这个树就可以得到该分词的所有细致分词划分，这主要
// 用于搜索引擎对一段文本进行全文搜索。
func (token *Token) Segments() []*Segment {
	return token.segments
}

// TextEquals 判断分词文本是否与给定字符串相等
func (token *Token) TextEquals(string string) bool {
	tokenLen := 0
	for _, t := range token.text {
		tokenLen += len(t)
	}
	if tokenLen != len(string) {
		return false
	}
	bytStr := []byte(string)
	index := 0
	for i := 0; i < len(token.text); i++ {
		textArray := []byte(token.text[i])
		for j := 0; j < len(textArray); j++ {
			if textArray[j] != bytStr[index] {
				index = index + 1
				return false
			}
			index = index + 1
		}
	}
	return true
}

type ConfigToken struct {
	Text      string `json:"text"`
	Frequency int    `json:"frequency"`
	Pos       string `json:"pos"`
}

func (c *ConfigToken) ToToken() *Token {
	return &Token{
		text:      splitTextToWords([]byte(c.Text)),
		frequency: c.Frequency,
		pos:       c.Pos,
	}
}
