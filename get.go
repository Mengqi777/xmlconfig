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
	"errors"
	"strconv"
	"strings"
)

// GetInt TODO
func (x *XmlConfig) GetInt(key string, defaultInt int) (int, error) {
	if value, ok := x.configurations[key]; ok {
		return strconv.Atoi(value.Value)
	}
	return defaultInt, nil
}

// GetInt8 TODO
func (x *XmlConfig) GetInt8(key string, defaultInt8 int8) (int8, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseInt(value.Value, 10, 8)
		return int8(i), err
	}
	return defaultInt8, nil
}

// GetInt16 TODO
func (x *XmlConfig) GetInt16(key string, defaultInt16 int16) (int16, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseInt(value.Value, 10, 16)
		return int16(i), err
	}
	return defaultInt16, nil
}

// GetInt32 TODO
func (x *XmlConfig) GetInt32(key string, defaultInt32 int32) (int32, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseInt(value.Value, 10, 32)
		return int32(i), err
	}
	return defaultInt32, nil

}

// GetInt64 TODO
func (x *XmlConfig) GetInt64(key string, defaultInt64 int64) (int64, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseInt(value.Value, 10, 64)
		return i, err
	}
	return defaultInt64, nil
}

// GetUint TODO
func (x *XmlConfig) GetUint(key string, defaultUint uint) (uint, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseUint(value.Value, 10, 32)
		return uint(i), err
	}
	return defaultUint, nil
}

// GetUint8 TODO
func (x *XmlConfig) GetUint8(key string, defaultUint8 uint8) (uint8, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseUint(value.Value, 10, 8)
		return uint8(i), err
	}
	return defaultUint8, nil
}

// GetUint16 TODO
func (x *XmlConfig) GetUint16(key string, defaultUint16 uint16) (uint16, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseUint(value.Value, 10, 16)
		return uint16(i), err
	}
	return defaultUint16, nil
}

// GetUint32 TODO
func (x *XmlConfig) GetUint32(key string, defaultUint32 uint32) (uint32, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseUint(value.Value, 10, 32)
		return uint32(i), err
	}
	return defaultUint32, nil
}

// GetUint64 TODO
func (x *XmlConfig) GetUint64(key string, defaultUint64 uint64) (uint64, error) {
	if value, ok := x.configurations[key]; ok {
		i, err := strconv.ParseUint(value.Value, 10, 64)
		return i, err
	}
	return defaultUint64, nil
}

// GetBool TODO
func (x *XmlConfig) GetBool(key string, defaultBool bool) bool {
	if value, ok := x.configurations[key]; ok {
		return strings.ToLower(strings.TrimSpace(value.Value)) == "true"
	} else {
		return defaultBool
	}
}

// Get TODO
func (x *XmlConfig) Get(key string) (string, error) {
	if value, ok := x.configurations[key]; ok {
		return value.Value, nil
	}
	return "", errors.New("not exist key: " + key)
}

// GetString TODO
func (x *XmlConfig) GetString(key string, defaultString string) string {
	if value, ok := x.configurations[key]; ok {
		return value.Value
	} else {
		return defaultString
	}
}

// GetTrimmedString TODO
func (x *XmlConfig) GetTrimmedString(key string, defaultString string) string {
	if value, ok := x.configurations[key]; ok {
		return strings.TrimSpace(value.Value)
	} else {
		return strings.TrimSpace(defaultString)
	}
}

// GetStrings TODO
func (x *XmlConfig) GetStrings(key, sep string) []string {
	if len(sep) == 0 {
		sep = ","
	}
	s := x.GetString(key, "")
	if len(s) == 0 {
		return []string{}
	}
	return strings.Split(s, sep)
}

// GetTrimmedStrings TODO
func (x *XmlConfig) GetTrimmedStrings(key, sep string) []string {
	if len(sep) == 0 {
		sep = ","
	}
	s := x.GetString(key, "")
	if len(s) == 0 {
		return []string{}
	}
	tmpArr := strings.Split(s, sep)
	var arr []string
	for i := range tmpArr {
		arr = append(arr, strings.TrimSpace(tmpArr[i]))
	}
	return arr
}

// GetPropsWithPrefix TODO
func (x *XmlConfig) GetPropsWithPrefix(prefix string) map[string]string {
	props := make(map[string]string)
	for key, value := range x.configurations {
		if strings.HasPrefix(key, prefix) {
			props[key] = value.Value
		}
	}
	return props
}

// GetConfigKeys TODO
func (x *XmlConfig) GetConfigKeys() []string {
	var keys []string
	for k, _ := range x.configurations {
		keys = append(keys, k)
	}
	return keys
}
