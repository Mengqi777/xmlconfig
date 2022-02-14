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
	f := "<property>" +
		"    <name>%s</name>" +
		"    <value>%s</value>" +
		"    <tag>%s</tag>" +
		"    <description>%s</description>" +
		"</property>"
	return fmt.Sprintf(f, p.Name, p.Value, p.Tag, p.Description)
}

// Equal TODO
func (p *property) Equal() bool {
	return false
}

// XmlConfig TODO
type XmlConfig struct {
	configurations map[string]*property
}

func (x *XmlConfig) parseXmlData(data []byte) error {
	var c configuration
	if err := xml.Unmarshal(data, &c); err != nil {
		return err
	}
	for _, property := range c.Properties {
		x.configurations[property.Name] = &property
	}
	return nil
}

// buildXmlData 构建xml配置
func (x *XmlConfig) buildXmlData() ([]byte, error) {
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

// ReadXmlFile XmlConfig 解析xml配置文件为XmlConfig，xml配置文件格式如下
/**
<?xml version="1.0" encoding="UTF-8"?>
<?xml-stylesheet type="text/xsl" href="configuration.xsl"?>
<configuration>
    <property>
        <name>name1</name>
        <value>value1</value>
        <tag>tag1,tag2</tag>
        <description>demo2</description>
    </property>
</configuration>
*/
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
	return x.parseXmlData(data)
}

// Read 从IO中读取配置
func (x *XmlConfig) Read(data []byte, r io.Reader) error {
	if _, err := r.Read(data); err != nil {
		return err
	}
	return x.parseXmlData(data)
}

// WriteXmlFile 将配置信息写入xml文件
func (x *XmlConfig) WriteXmlFile(xmlFilePath string) error {
	data, err := x.buildXmlData()
	if err != nil {
		return err
	}
	return ioutil.WriteFile(xmlFilePath, data, 0644)
}

// Write  将配置信息写入IO
func (x *XmlConfig) Write(w io.Writer) error {
	data, err := x.buildXmlData()
	if err != nil {
		return err
	}
	_, err = w.Write(data)
	return err
}
