# Array of Unit Test Cases

Test case can be written in two ways such as,
* Single Test 
* Array of test cases

### Single Test 
Test case attribute declare as a single test unit before executing testable function or method.
```
func TestXxx(t *testing.T) {
    // test attribute declaration
    // here
    out,err := testableFunction(in)
    if err!=nil{
        t.Errro(err)
    }
    }
```
### Test Array
Map of test case structs use for implementing array of test cases. It is easy to organise different test scenarios in single test function.
```
func TestXxx(t *testing.T) {
tests := map[string]struct{
    in interface{}
    out interface{}
    wantError bool
  }{
  `test case 1` :{
    in input1,
    out output1,
    wantError false,
  },
  // more test case
}

// test attribute declaration
// here
for name, test := range tests {
  t.Run(name, func(t *testing.T) {
    out,err := testableFunction()
    if err!=nil{
      t.Errro(err)
    }
  })
}
}
```