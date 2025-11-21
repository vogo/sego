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

package cedar

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"io"
	"os"
)

// Save saves the cedar to an io.Writer,
// where dataType is either "json" or "gob".
func (da *Cedar) Save(out io.Writer, dataType string) error {
	switch dataType {
	case "gob", "GOB":
		dataEecoder := gob.NewEncoder(out)
		return dataEecoder.Encode(da.cedar)
	case "json", "JSON":
		dataEecoder := json.NewEncoder(out)
		return dataEecoder.Encode(da.cedar)
	}
	return ErrInvalidDataType
}

// SaveToFile saves the cedar to a file,
// where dataType is either "json" or "gob".
func (da *Cedar) SaveToFile(fileName string, dataType string) error {
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0o666)
	if err != nil {
		return err
	}
	defer file.Close()
	out := bufio.NewWriter(file)
	defer out.Flush()
	da.Save(out, dataType)
	return nil
}

// Load loads the cedar from an io.Writer,
// where dataType is either "json" or "gob".
func (da *Cedar) Load(in io.Reader, dataType string) error {
	switch dataType {
	case "gob", "GOB":
		dataDecoder := gob.NewDecoder(in)
		return dataDecoder.Decode(da.cedar)
	case "json", "JSON":
		dataDecoder := json.NewDecoder(in)
		return dataDecoder.Decode(da.cedar)
	}
	return ErrInvalidDataType
}

// LoadFromFile loads the cedar from a file,
// where dataType is either "json" or "gob".
func (da *Cedar) LoadFromFile(fileName string, dataType string) error {
	file, err := os.OpenFile(fileName, os.O_RDONLY, 0o600)
	defer file.Close()
	if err != nil {
		return err
	}
	in := bufio.NewReader(file)
	return da.Load(in, dataType)
}
