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

// Package cedar implements double-array trie.
//
// It is a golang port of cedar (http://www.tkl.iis.u-tokyo.ac.jp/~ynaga/cedar) which is written in C++ by Naoki Yoshinaga.
// Currently cedar implements the `reduced` verion of cedar.
// This package is not thread safe if there is one goroutine doing
// insertions or deletions.
//
// # Note
//
// key must be `[]byte` without zero items,
// while value must be integer in the range [0, 2<<63-2] or [0, 2<<31-2] depends on the platform.
package cedar
