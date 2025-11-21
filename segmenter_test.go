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
	"testing"
)

var prodSeg = Segmenter{}

func TestSplit(t *testing.T) {
	expect(t, "中/国/有/十/三/亿/人/口/",
		bytesToString(splitTextToWords([]byte(
			"中国有十三亿人口"))))

	expect(t, "github/ /is/ /a/ /web/-/based/ /hosting/ /service/,/ /for/ /software/ /development/ /projects/./",
		bytesToString(splitTextToWords([]byte(
			"GitHub is a web-based hosting service, for software development projects."))))

	expect(t, "中/国/雅/虎/yahoo/!/ /china/致/力/于/，/领/先/的/公/益/民/生/门/户/网/站/。/",
		bytesToString(splitTextToWords([]byte(
			"中国雅虎Yahoo! China致力于，领先的公益民生门户网站。"))))

	expect(t, "こ/ん/に/ち/は/", bytesToString(splitTextToWords([]byte("こんにちは"))))

	expect(t, "안/녕/하/세/요/", bytesToString(splitTextToWords([]byte("안녕하세요"))))

	expect(t, "Я/ /тоже/ /рада/ /Вас/ /видеть/", bytesToString(splitTextToWords([]byte("Я тоже рада Вас видеть"))))

	expect(t, "¿/cómo/ /van/ /las/ /cosas/", bytesToString(splitTextToWords([]byte("¿Cómo van las cosas"))))

	expect(t, "wie/ /geht/ /es/ /ihnen/", bytesToString(splitTextToWords([]byte("Wie geht es Ihnen"))))

	expect(t, "je/ /suis/ /enchanté/ /de/ /cette/ /pièce/",
		bytesToString(splitTextToWords([]byte("Je suis enchanté de cette pièce"))))
}

func TestSegment(t *testing.T) {
	var seg Segmenter
	configTokens := []*ConfigToken{
		{Text: "猫猫", Frequency: 2, Pos: "n"},
		{Text: "狗狗", Frequency: 2, Pos: "n"},
	}
	seg.Load("testdata/test_dict1.txt,testdata/test_dict2.txt", configTokens)
	expect(t, "14", seg.dict.NumTokens())
	segments := seg.Segment([]byte("中国有十三亿人口"))
	expect(t, "中国/ 有/p3 十三亿/ 人口/p12 ", SegmentsToString(segments, false))
	expect(t, "4", len(segments))
	expect(t, "0", segments[0].start)
	expect(t, "6", segments[0].end)
	expect(t, "6", segments[1].start)
	expect(t, "9", segments[1].end)
	expect(t, "9", segments[2].start)
	expect(t, "18", segments[2].end)
	expect(t, "18", segments[3].start)
	expect(t, "24", segments[3].end)
	segments = seg.Segment([]byte("猫猫狗狗"))
	expect(t, "2", len(segments))
	expect(t, "0", segments[0].start)
	expect(t, "6", segments[0].end)
	expect(t, "6", segments[1].start)
	expect(t, "12", segments[1].end)
	expect(t, "n", segments[0].token.pos)
	expect(t, "n", segments[1].token.pos)
}

func TestLargeDictionary(t *testing.T) {
	prodSeg.LoadDictionary("data/dictionary.txt")
	expect(t, "中国/ns 人口/n ", SegmentsToString(prodSeg.Segment(
		[]byte("中国人口")), false))

	expect(t, "中国/ns 人口/n ", SegmentsToString(prodSeg.internalSegment(
		[]byte("中国人口"), false), false))

	expect(t, "中国/ns 人口/n ", SegmentsToString(prodSeg.internalSegment(
		[]byte("中国人口"), true), false))

	expect(t, "中华人民共和国/ns 中央人民政府/nt ", SegmentsToString(prodSeg.internalSegment(
		[]byte("中华人民共和国中央人民政府"), true), false))

	expect(t, "中华人民共和国中央人民政府/nt ", SegmentsToString(prodSeg.internalSegment(
		[]byte("中华人民共和国中央人民政府"), false), false))

	expect(t, "中华/nz 人民/n 共和/nz 国/n 共和国/ns 人民共和国/nt 中华人民共和国/ns 中央/n 人民/n 政府/n 人民政府/nt 中央人民政府/nt 中华人民共和国中央人民政府/nt ", SegmentsToString(prodSeg.Segment(
		[]byte("中华人民共和国中央人民政府")), true))
}
