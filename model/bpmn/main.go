package main

import (
	"encoding/xml"
	"fmt"
)

type Person struct {
	XMLName xml.Name `xml:"person"`
	Name    string   `xml:"name"`
	Age     int      `xml:"age"`
}

type Animal struct {
	XMLName xml.Name `xml:"animal"`
	Name    string   `xml:"name"`
	Type    string   `xml:"type"`
}

type Item struct {
	XMLName xml.Name `xml:"item"`
	Value   interface{}
}

type Data struct {
	XMLName xml.Name `xml:"data"`
	Items   []Item   `xml:",any"`
}

func main() {
	xmlData := `
        <data>
            <item>
                <person>
                    <name>Alice</name>
                    <age>30</age>
                </person>
            </item>
            <item>
                <animal>
                    <name>Fluffy</name>
                    <type>Cat</type>
                </animal>
            </item>
        </data>
    `
	var data Data
	err := xml.Unmarshal([]byte(xmlData), &data)
	if err != nil {
		panic(err)
	}
	for _, item := range data.Items {
		switch v := item.Value.(type) {
		case Person:
			fmt.Printf("Person: %s, %d\n", v.Name, v.Age)
		case Animal:
			fmt.Printf("Animal: %s, %s\n", v.Name, v.Type)
		}
	}
}
