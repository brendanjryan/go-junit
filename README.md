# go-junitxml

A collection of types used to represent the [JUnit XML Schema Spec](https://www.ibm.com/support/knowledgecenter/en/SSUFAU_1.0.0/com.ibm.rsar.analysis.codereview.cobol.doc/topics/cac_useresults_junit.html).

Example Usage:
```go
package main 

import (
  junitxml "github.com/brendanjryan/go-junit/xml"
  "encoding/xml"
  "fmt"
)

func main(){
	
  ss := &junitxml.Suites{
  	Id: "com.foo.bar.test",
    Name: "The label of the scan",
    Tests: 1,
    Failures: 1,
    Suites: []Suite{
    	{
          Id: "com.foo.bar.test",
          Name: "The label of the provider",
          Tests: 1,
          Failures: 1,
          Testcases: []Testcase{
          	{
                Id: "com.foo.bar.test.always_fails",
                Name: "A test that always fails",
                Time: 0.01,
                Failures: []*Failure{
                    {
                      Message: "This test always fails",
                    }
                },
            },
         }
      }
    }
  }
  
  out, err := xml.Marshal(ss)
  if err != nil {
  	log.Fatal("error marshaling data: ", err)
  }
  
  log.Println(out)
}
```

