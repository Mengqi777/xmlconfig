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
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var stringCase = "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n" +
	"<configuration>\n" +
	"    <property>\n" +
	"        <name>name1</name>\n" +
	"        <value>value1</value>\n" +
	"        <tag>tag1,tag2</tag>\n" +
	"        <description>demo1</description>\n" +
	"    </property>\n" +
	"</configuration>"

var writeXml = "w.xml"

func newConfigurations() map[string]*property {
	return map[string]*property{"name1": {
		XMLName: xml.Name{
			Space: "",
			Local: "property",
		},
		Name:        "name1",
		Value:       "value1",
		Tag:         "tag1,tag2",
		Description: "demo1",
	}}
}

func TestXmlConfig_ReadXmlFile(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "错误路径",
			fields: fields{configurations: make(map[string]*property)},
			args:   args{path: "empty.xml"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err != nil
			},
		},
		{
			name:   "正常路径",
			fields: fields{configurations: newConfigurations()},
			args:   args{path: "test.xml"},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: make(map[string]*property),
			}
			err := x.ReadXmlFile(tt.args.path)
			if !tt.wantErr(t, err, fmt.Sprintf("Write(%v)", x)) {
				panic(err)
			}
			if tt.name == "正常路径" {
				assert.Equal(t, newConfigurations(), x.configurations)
			}
		})
	}
}

func TestXmlConfig_Read(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		r io.Reader
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "从IO读取",
			fields: fields{configurations: newConfigurations()},
			args:   args{r: strings.NewReader(stringCase)},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: make(map[string]*property),
			}
			data := make([]byte, len(stringCase))
			err := x.Read(data, tt.args.r)
			if !tt.wantErr(t, err, fmt.Sprintf("Read(%v)", tt.args.r)) {
				panic(err)
			}
			assert.Equal(t, x.configurations, newConfigurations())
		})
	}
}

func TestXmlConfig_Write(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	tests := []struct {
		name    string
		fields  fields
		wantW   string
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "写入IO",
			fields: fields{configurations: newConfigurations()},
			wantW:  stringCase,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			w := &bytes.Buffer{}
			err := x.Write(w)
			if !tt.wantErr(t, err, fmt.Sprintf("Write(%v)", w)) {
				return
			}
			assert.Equalf(t, tt.wantW, w.String(), "Write(%v)", w)
		})
	}
}

func TestXmlConfig_WriteXmlFile(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		xmlFilePath string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "写入xml文件",
			fields: fields{configurations: newConfigurations()},
			args:   args{xmlFilePath: writeXml},
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			err := x.WriteXmlFile(tt.args.xmlFilePath)
			if !tt.wantErr(t, err, fmt.Sprintf("WriteXmlFile(%v)", tt.args.xmlFilePath)) {
				panic(err)
			}
			data, _ := ioutil.ReadFile(writeXml)
			assert.Equal(t, stringCase, string(data))
			os.Remove(writeXml)
		})
	}
}
