# go-junitxml

A collection of types used to represent the [JUnit XML Schema Spec](https://www.ibm.com/support/knowledgecenter/en/SSUFAU_1.0.0/com.ibm.rsar.analysis.codereview.cobol.doc/topics/cac_useresults_junit.html).

Example Usage:

```go
package main

import (
	"fmt"

	junitxml "github.com/brendanjryan/go-junit/xml"
)

func main() {
	ss := &junitxml.Suites{
		ID:       "com.foo.bar.test",
		Name:     "The label of the scan",
		Tests:    1,
		Failures: 1,
		Suites: []junitxml.Suite{
			{
				ID:       "com.foo.bar.test",
				Name:     "The label of the provider",
				Tests:    1,
				Failures: 1,
				TestCases: []junitxml.TestCase{
					{
						ID:   "com.foo.bar.test.always_fails",
						Name: "A test that always fails",
						Time: 0.01,
						Failures: []junitxml.Failure{
							{
								Message: "This test always fails",
							},
						},
					},
				},
			},
		},
	}

	fmt.Println("ss: %+v", ss)
}
```
