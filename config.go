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
	"io"
	"io/ioutil"
	"os"
)

// configuration TODO
type configuration struct {
	XMLName    xml.Name   `xml:"configuration"`
	Properties []property `xml:"property"`
}

// property TODO
type property struct {
	XMLName     xml.Name `xml:"property"`
	Name        string   `xml:"name"`
	Value       string   `xml:"value"`
	Tag         string   `xml:"tag"`
	Description string   `xml:"description"`
}

// String TODO
func (p *property) String() string {
	f := "<property>\n" +
		"    <name>%s</name>\n" +
		"    <value>%s</value>\n" +
		"    <tag>%s</tag>\n" +
		"    <description>%s</description>\n" +
		"</property>\n"
	return fmt.Sprintf(f, p.Name, p.Value, p.Tag, p.Description)
}

// Equal TODO
func (p *property) Equal(o *property) bool {
	return p.Name == o.Name && p.Value == o.Value && p.Tag == o.Tag
}

// XmlConfig TODO
type XmlConfig struct {
	configurations map[string]*property
}

// NewXmlConfig TODO
func NewXmlConfig() *XmlConfig {
	return &XmlConfig{
		configurations: make(map[string]*property),
	}
}

// String TODO
func (x *XmlConfig) String() string {
	str := ""
	for _, property := range x.configurations {
		str += property.String()
	}
	return str
}

// ParseXmlData TODO
func (x *XmlConfig) ParseXmlData(data []byte) error {
	c := &configuration{
		XMLName:    xml.Name{},
		Properties: make([]property, 0),
	}
	if err := xml.Unmarshal(data, c); err != nil {
		return err
	}
	for _, p := range c.Properties {
		x.configurations[p.Name] = &property{
			XMLName: xml.Name{
				Space: p.XMLName.Space,
				Local: p.XMLName.Local,
			},
			Name:        p.Name,
			Value:       p.Value,
			Tag:         p.Tag,
			Description: p.Description,
		}
	}
	return nil
}

// BuildXmlData 构建xml配置
func (x *XmlConfig) BuildXmlData() ([]byte, error) {
	var properties []property
	for _, property := range x.configurations {
		properties = append(properties, *property)
	}
	c := &configuration{
		Properties: properties,
	}
	data, err := xml.MarshalIndent(c, "", "    ")
	if err != nil {
		return nil, err
	}
	head := []byte(xml.Header)
	data = append(head, data...)
	return data, nil
}

// ReadXmlFile 从xml文件中读取配置
func (x *XmlConfig) ReadXmlFile(xmlFilePath string) error {
	xmlFile, err := os.Open(xmlFilePath)
	if err != nil {
		return err
	}
	defer xmlFile.Close()
	data, err := ioutil.ReadAll(xmlFile)
	if err != nil {
		return err
	}
	return x.ParseXmlData(data)
}

// Read 从IO中读取配置
func (x *XmlConfig) Read(data []byte, r io.Reader) error {
	if _, err := r.Read(data); err != nil {
		return err
	}
	return x.ParseXmlData(data)
}

// WriteXmlFile 将配置信息写入xml文件
func (x *XmlConfig) WriteXmlFile(xmlFilePath string) error {
	data, err := x.BuildXmlData()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(xmlFilePath, data, 0644)
}

// Write  将配置信息写入IO
func (x *XmlConfig) Write(w io.Writer) error {
	data, err := x.BuildXmlData()
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
