// MIT License
//
// Copyright (c) 2022 孟琦
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.
package xmlconfig

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

// SetString TODO
func (x *XmlConfig) SetString(key string, value string) {
	if _, ok := x.configurations[key]; ok {
		x.configurations[key].Value = value
	} else {
		x.configurations[key] = &property{
			XMLName:     xml.Name{Local: "property"},
			Name:        key,
			Value:       value,
			Tag:         "",
			Description: "",
		}
	}
}

// SetBool TODO
func (x *XmlConfig) SetBool(key string, value bool) {
	if value {
		x.SetString(key, "true")
	} else {
		x.SetString(key, "false")
	}
}

// SetInt TODO
func (x *XmlConfig) SetInt(key string, value int64) {
	x.SetString(key, strconv.FormatInt(value, 10))
}

// SetUint TODO
func (x *XmlConfig) SetUint(key string, value uint64) {
	x.SetString(key, strconv.FormatUint(value, 10))
}

// SetIfUnset TODO
func (x *XmlConfig) SetIfUnset(key string, value string) {
	if _, ok := x.configurations[key]; !ok {
		fmt.Println(x.configurations[key])
		x.SetString(key, value)
	}
}

func (x *XmlConfig) unset(key string) {
	if _, ok := x.configurations[key]; ok {
		delete(x.configurations, key)
	}
}
