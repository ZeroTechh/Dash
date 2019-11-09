# Dash
Easily Manage Config Files In Go

## Example
config.yaml
```yaml
test:
  testInt: 32
  testFloat: 32.2
  testBool: true
  testString: "hey"
  testList:
    - "a"
    - "b"
    - "c'
```
main.go
```go
// provide it with the filename and different path where config file could be
config := dash.GetConfig("test.yaml", []string{"path1", "path2"})
test := config.Map("test")
testInt := test.Int("testInt")
testFloat := test.Float("testFloat")
testBool := test.Bool("testBool")
testList := test.List("testList")
```