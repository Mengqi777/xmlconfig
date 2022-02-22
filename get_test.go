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
	"testing"

	"github.com/stretchr/testify/assert"
)

func newCase(value string) map[string]*property {
	return map[string]*property{"name1": {
		XMLName: xml.Name{
			Space: "",
			Local: "property",
		},
		Name:        "name1",
		Value:       value,
		Tag:         "tag1,tag2",
		Description: "demo1",
	}}
}

func TestXmlConfig_GetBool(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key         string
		defaultBool bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		// TODO: Add test cases.
		{
			name:   "获取bool/true",
			fields: fields{configurations: newCase("true")},
			args:   args{"name1", true},
			want:   true,
		},
		{
			name:   "获取bool/false",
			fields: fields{configurations: newCase("false")},
			args:   args{"name1", false},
			want:   false,
		},
		{
			name:   "获取bool/value异常",
			fields: fields{configurations: newCase("")},
			args:   args{"name1", false},
			want:   false,
		},
		{
			name:   "获取bool/key不存在",
			fields: fields{configurations: newCase("true")},
			args:   args{"name11", false},
			want:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			assert.Equalf(t, tt.want, x.GetBool(tt.args.key, tt.args.defaultBool), "GetBool(%v, %v)", tt.args.key, tt.args.defaultBool)
		})
	}
}

func TestXmlConfig_GetConfigKeys(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		// TODO: Add test cases.
		{
			name:   "获取所有key",
			fields: fields{configurations: newConfigurations()},
			want:   []string{"name1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			assert.Equalf(t, tt.want, x.GetConfigKeys(), "GetConfigKeys()")
		})
	}
}

func TestXmlConfig_GetInt(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key        string
		defaultInt int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取int",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 0},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		}, {
			name:   "获取int/小数点",
			fields: fields{configurations: newCase("1.2")},
			args:   args{"name1", 0},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取int/与默认值不等",
			fields: fields{configurations: newCase("3")},
			args:   args{"name1", 1},
			want:   3,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取int默认值",
			fields: fields{configurations: newCase("3")},
			args:   args{"name3", 1},
			want:   1,
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
			got, err := x.GetInt(tt.args.key, tt.args.defaultInt)
			if !tt.wantErr(t, err, fmt.Sprintf("GetInt(%v, %v)", tt.args.key, tt.args.defaultInt)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetInt(%v, %v)", tt.args.key, tt.args.defaultInt)
		})
	}
}

func TestXmlConfig_GetInt16(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key          string
		defaultInt16 int16
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int16
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取int16",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 0},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取int16默认值",
			fields: fields{configurations: newCase("3")},
			args:   args{"name3", 1},
			want:   1,
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
			got, err := x.GetInt16(tt.args.key, tt.args.defaultInt16)
			if !tt.wantErr(t, err, fmt.Sprintf("GetInt16(%v, %v)", tt.args.key, tt.args.defaultInt16)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetInt16(%v, %v)", tt.args.key, tt.args.defaultInt16)
		})
	}
}

func TestXmlConfig_GetInt32(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key          string
		defaultInt32 int32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int32
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取int32",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 0},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取int32默认值",
			fields: fields{configurations: newCase("3")},
			args:   args{"name3", 1},
			want:   1,
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
			got, err := x.GetInt32(tt.args.key, tt.args.defaultInt32)
			if !tt.wantErr(t, err, fmt.Sprintf("GetInt32(%v, %v)", tt.args.key, tt.args.defaultInt32)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetInt32(%v, %v)", tt.args.key, tt.args.defaultInt32)
		})
	}
}

func TestXmlConfig_GetInt64(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key          string
		defaultInt64 int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int64
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取int64",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 0},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取int64默认值",
			fields: fields{configurations: newCase("3")},
			args:   args{"name3", 1},
			want:   1,
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
			got, err := x.GetInt64(tt.args.key, tt.args.defaultInt64)
			if !tt.wantErr(t, err, fmt.Sprintf("GetInt64(%v, %v)", tt.args.key, tt.args.defaultInt64)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetInt64(%v, %v)", tt.args.key, tt.args.defaultInt64)
		})
	}
}

func TestXmlConfig_GetInt8(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key         string
		defaultInt8 int8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int8
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取int8",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 0},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取int8默认值",
			fields: fields{configurations: newCase("3")},
			args:   args{"name3", 1},
			want:   1,
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
			got, err := x.GetInt8(tt.args.key, tt.args.defaultInt8)
			if !tt.wantErr(t, err, fmt.Sprintf("GetInt8(%v, %v)", tt.args.key, tt.args.defaultInt8)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetInt8(%v, %v)", tt.args.key, tt.args.defaultInt8)
		})
	}
}

func TestXmlConfig_GetPropsWithPrefix(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		prefix string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   map[string]string
	}{
		// TODO: Add test cases.
		{
			name:   "根据前缀获取kv",
			fields: fields{configurations: newCase("1")},
			args: args{
				prefix: "name",
			},
			want: map[string]string{"name1": "1"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			assert.Equalf(t, tt.want, x.GetPropsWithPrefix(tt.args.prefix), "GetPropsWithPrefix(%v)", tt.args.prefix)
		})
	}
}

func TestXmlConfig_GetString(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key           string
		defaultString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "获取string",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1", "string"},
			want:   "value1",
		},
		{
			name:   "获取默认string",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name2", "string"},
			want:   "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			assert.Equalf(t, tt.want, x.GetString(tt.args.key, tt.args.defaultString), "GetString(%v, %v)", tt.args.key, tt.args.defaultString)
		})
	}
}

func TestXmlConfig_GetTrimmedString(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key           string
		defaultString string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "获取TrimmedString",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1", "string"},
			want:   "value1",
		},
		{
			name:   "获取默认TrimmedString",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name2", " string "},
			want:   "string",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			x.SetString("name1", "  value1  ")
			assert.Equalf(t, tt.want, x.GetTrimmedString(tt.args.key, tt.args.defaultString), "GetTrimmedString(%v, %v)", tt.args.key, tt.args.defaultString)
		})
	}
}

func TestXmlConfig_GetUint(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key         string
		defaultUint uint
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取uint",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint/小数",
			fields: fields{configurations: newCase("1.3")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint/默认值",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name2", 1},
			want:   1,
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
			got, err := x.GetUint(tt.args.key, tt.args.defaultUint)
			if !tt.wantErr(t, err, fmt.Sprintf("GetUint(%v, %v)", tt.args.key, tt.args.defaultUint)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetUint(%v, %v)", tt.args.key, tt.args.defaultUint)
		})
	}
}

func TestXmlConfig_GetUint16(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key           string
		defaultUint16 uint16
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint16
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取uint16",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint16/小数",
			fields: fields{configurations: newCase("1.3")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint16/默认值",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name2", 1},
			want:   1,
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
			got, err := x.GetUint16(tt.args.key, tt.args.defaultUint16)
			if !tt.wantErr(t, err, fmt.Sprintf("GetUint16(%v, %v)", tt.args.key, tt.args.defaultUint16)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetUint16(%v, %v)", tt.args.key, tt.args.defaultUint16)
		})
	}
}

func TestXmlConfig_GetUint32(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key           string
		defaultUint32 uint32
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint32
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取uint32",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint32/小数",
			fields: fields{configurations: newCase("1.3")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint32/默认值",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name2", 1},
			want:   1,
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
			got, err := x.GetUint32(tt.args.key, tt.args.defaultUint32)
			if !tt.wantErr(t, err, fmt.Sprintf("GetUint32(%v, %v)", tt.args.key, tt.args.defaultUint32)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetUint32(%v, %v)", tt.args.key, tt.args.defaultUint32)
		})
	}
}

func TestXmlConfig_GetUint64(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key           string
		defaultUint64 uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint64
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取uint64",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint64/小数",
			fields: fields{configurations: newCase("1.3")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint64/默认值",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name2", 1},
			want:   1,
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
			got, err := x.GetUint64(tt.args.key, tt.args.defaultUint64)
			if !tt.wantErr(t, err, fmt.Sprintf("GetUint64(%v, %v)", tt.args.key, tt.args.defaultUint64)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetUint64(%v, %v)", tt.args.key, tt.args.defaultUint64)
		})
	}
}

func TestXmlConfig_GetUint8(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key          string
		defaultUint8 uint8
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    uint8
		wantErr assert.ErrorAssertionFunc
	}{
		// TODO: Add test cases.
		{
			name:   "获取uint8",
			fields: fields{configurations: newCase("1")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint8/小数",
			fields: fields{configurations: newCase("1.3")},
			args:   args{"name1", 1},
			want:   1,
			wantErr: func(t assert.TestingT, err error, i ...interface{}) bool {
				return err == nil
			},
		},
		{
			name:   "获取uint8/默认值",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name2", 1},
			want:   1,
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
			got, err := x.GetUint8(tt.args.key, tt.args.defaultUint8)
			if !tt.wantErr(t, err, fmt.Sprintf("GetUint8(%v, %v)", tt.args.key, tt.args.defaultUint8)) {
				return
			}
			assert.Equalf(t, tt.want, got, "GetUint8(%v, %v)", tt.args.key, tt.args.defaultUint8)
		})
	}
}

func TestXmlConfig_Get(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		// TODO: Add test cases.
		{
			name:   "get",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1"},
			want:   "value1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			value, err := x.Get(tt.args.key)
			assert.Nil(t, err, "Get(%v)", tt.args.key)
			assert.Equalf(t, tt.want, value, "Get(%v)", tt.args.key)
		})
	}
}

func TestXmlConfig_GetTrimmedStrings(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key string
		sep string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name:   "获取strings",
			fields: fields{configurations: newCase(" t1, t2 ,t3 ")},
			args:   args{"name1", ","},
			want:   []string{"t1", "t2", "t3"},
		},
		{
			name:   "获取strings",
			fields: fields{configurations: newCase("t1,t")},
			args:   args{"name2", ","},
			want:   []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			assert.Equalf(t, tt.want, x.GetTrimmedStrings(tt.args.key, tt.args.sep), "GetTrimmedStrings(%v, %v)", tt.args.key, tt.args.sep)
		})
	}
}

func TestXmlConfig_GetStrings(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key string
		sep string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []string
	}{
		// TODO: Add test cases.
		{
			name:   "获取strings",
			fields: fields{configurations: newCase("t1,t2,t3")},
			args:   args{"name1", ","},
			want:   []string{"t1", "t2", "t3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				configurations: tt.fields.configurations,
			}
			assert.Equalf(t, tt.want, x.GetStrings(tt.args.key, tt.args.sep), "GetStrings(%v, %v)", tt.args.key, tt.args.sep)
		})
	}
}
