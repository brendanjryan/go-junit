package xml

import (
	"encoding/xml"
)

// Suites is a collection of JUnit test suites.
type Suites struct {
	XMLName  xml.Name `xml:"testsuites"`
	Id       string   `xml:"id,attr"`
	Name     string   `xml:"name,attr"`
	Tests    int      `xml:"tests,attr"`
	Failures int      `xml:"failures,attr"`
	Time     string   `xml:"time,attr"`

	Suites []Suite
}

// Suite is a single JUnit test suite which may contain many
// testcases.
type Suite struct {
	XMLName    xml.Name   `xml:"testsuite"`
	Id         string     `xml:"id,attr"`
	Tests      int        `xml:"tests,attr"`
	Failures   int        `xml:"failures,attr"`
	Time       string     `xml:"time,attr"`
	Name       string     `xml:"name,attr"`
	Properties []Property `xml:"properties>property,omitempty"`

	TestCases []TestCase
}

// TestCase is a single test case with its result.
type TestCase struct {
	XMLName xml.Name `xml:"testcase"`
	Name    string   `xml:"name,attr"`
	Time    string   `xml:"time,attr"`

	Failure *Failure `xml:"failure,omitempty"`
}

// Property represents a key/value pair used to define properties.
type Property struct {
	Name  string `xml:"name,attr"`
	Value string `xml:"value,attr"`
}

// Failure contains data related to a failed test.
type Failure struct {
	XMLName  xml.Name `xml:"failure"`
	Message  string   `xml:"message,attr"`
	Type     string   `xml:"type,attr"`
	Contents string   `xml:",chardata"`
}
