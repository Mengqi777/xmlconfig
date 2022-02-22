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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXmlConfig_SetBool(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key   string
		value bool
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantB  string
	}{
		// TODO: Add test cases.
		{
			name:   "设置bool/true",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1", true},
			wantB:  "true",
		},
		{
			name:   "设置bool/false",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1", false},
			wantB:  "false",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				Configurations: tt.fields.configurations,
			}
			x.SetBool(tt.args.key, tt.args.value)
			value, _ := x.Get("name1")
			assert.Equalf(t, tt.wantB, value, "Write(%v)", x)
		})
	}
	t.Cleanup(func() {
		tests = nil
	})
}

func TestXmlConfig_SetIfUnset(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  string
	}{
		// TODO: Add test cases.
		{
			name:   "不存在则设置",
			fields: fields{configurations: newConfigurations()},
			args:   args{"set", "setIfUnset"},
			wantS:  "setIfUnset",
		},
		{
			name:   "存在则不设置",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1", "setIfUnset"},
			wantS:  "value1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				Configurations: tt.fields.configurations,
			}
			x.SetIfUnset(tt.args.key, tt.args.value)
			if tt.name == "不存在则设置" {
				value, _ := x.Get("set")
				assert.Equalf(t, tt.wantS, value, "Write(%v)", x)
			}
			if tt.name == "存在则不设置" {
				value, _ := x.Get("name1")
				assert.Equalf(t, tt.wantS, value, "Write(%v)", x)
			}
		})
	}
	t.Cleanup(func() {
		tests = nil
	})
}

func TestXmlConfig_SetInt(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key   string
		value int64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  string
	}{
		// TODO: Add test cases.
		{
			name:   "设置int",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1", 1},
			wantS:  "1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				Configurations: tt.fields.configurations,
			}
			x.SetInt(tt.args.key, tt.args.value)
			value, _ := x.Get("name1")
			assert.Equalf(t, tt.wantS, value, "Write(%v)", x)
		})
	}
	t.Cleanup(func() {
		tests = nil
	})
}

func TestXmlConfig_SetString(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  string
	}{
		// TODO: Add test cases.
		{
			name:   "设置string",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1", "val"},
			wantS:  "val",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				Configurations: tt.fields.configurations,
			}
			x.SetString(tt.args.key, tt.args.value)
			value, _ := x.Get("name1")
			assert.Equalf(t, tt.wantS, value, "Write(%v)", x)
		})
	}
	t.Cleanup(func() {
		tests = nil
	})
}

func TestXmlConfig_SetUint(t *testing.T) {
	type fields struct {
		configurations map[string]*property
	}
	type args struct {
		key   string
		value uint64
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		wantS  string
	}{
		// TODO: Add test cases.
		{
			name:   "设置uint",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1", 12333331231231},
			wantS:  "12333331231231",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				Configurations: tt.fields.configurations,
			}
			x.SetUint(tt.args.key, tt.args.value)
			value, _ := x.Get("name1")
			assert.Equalf(t, tt.wantS, value, "Write(%v)", x)
		})
	}
}

func TestXmlConfig_unset(t *testing.T) {
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
		wantS  string
	}{
		// TODO: Add test cases.
		{
			name:   "删除",
			fields: fields{configurations: newConfigurations()},
			args:   args{"name1"},
			wantS:  "value1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			x := &XmlConfig{
				Configurations: tt.fields.configurations,
			}
			value, err := x.Get("name1")
			assert.Nil(t, err)
			assert.NotNil(t, value)
			assert.Equalf(t, tt.wantS, value, "Write(%v)", x)
			x.unset(tt.args.key)
			assert.Equal(t, 0, len(x.GetConfigKeys()))
			value, err = x.Get("name1")
			assert.NotNil(t, err)
			assert.Equalf(t, "", value, "Write(%v)", x)
		})
	}
}
