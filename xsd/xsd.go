package xsd

import (
	"encoding/xml"
)

type Schema struct {
	XMLName            xml.Name      `xml:"http://www.w3.org/2001/XMLSchema schema"`
	TNS                string        `xml:"xmlns tns",attr`
	XS                 string        `xml:"xmlns xs,attr"`
	TargetNamespace    string        `xml:"targetNamespace,attr"`
	ElementFormDefault string        `xml:"elementFormDefault,attr"`
	Version            string        `xml:"version,attr"`
	Include            []Include     `xml:"http://www.w3.org/2001/XMLSchema include"`
	Imports            []Import      `xml:"http://www.w3.org/2001/XMLSchema import"`
	Elements           []Element     `xml:"http://www.w3.org/2001/XMLSchema element"`
	ComplexTypes       []ComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	SimpleTypes        []SimpleType  `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
}

type Include struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2001/XMLSchema include"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Namespace      string   `xml:"namespace,attr"`
}

type Element struct {
	XMLName      xml.Name     `xml:"http://www.w3.org/2001/XMLSchema element"`
	Type         string       `xml:"type,attr"`
	Nillable     string       `xml:"nillable,attr"`
	MinOccurs    string       `xml:"minOccurs,attr"`
	MaxOccurs    string       `xml:"maxOccurs,attr"`
	Form         string       `xml:"form,attr"`
	Name         string       `xml:"name,attr"`
	ComplexTypes *ComplexType `xml:"http://www.w3.org/2001/XMLSchema complexType"`
}

type SimpleType struct {
	XMLName     xml.Name `xml:"http://www.w3.org/2001/XMLSchema simpleType"`
	Restriction Restriction
}

type Restriction struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema restriction"`
	Base    string   `xml:"base,attr"`
	Pattern Pattern
}

type Pattern struct {
	XMLName xml.Name `xml:"http://www.w3.org/2001/XMLSchema pattern"`
	Value   string   `xml:"value,attr"`
}

type ComplexType struct {
	XMLName  xml.Name        `xml:"http://www.w3.org/2001/XMLSchema complexType"`
	Name     string          `xml:"name,attr"`
	Abstract bool            `xml:"abstract,attr"`
	Sequence []Element       `xml:"sequence>element"`
	Content  *ComplexContent `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
}

type ComplexContent struct {
	XMLName   xml.Name  `xml:"http://www.w3.org/2001/XMLSchema complexContent"`
	Extension Extension `xml:"http://www.w3.org/2001/XMLSchema extension"`
}

type Extension struct {
	XMLName  xml.Name  `xml:"http://www.w3.org/2001/XMLSchema extension"`
	Base     string    `xml:"base,attr"`
	Sequence []Element `xml:"sequence>element"`
}

type Import struct {
	XMLName        xml.Name `xml:"http://www.w3.org/2001/XMLSchema import"`
	SchemaLocation string   `xml:"schemaLocation,attr"`
	Namespace      string   `xml:"namespace,attr"`
}
