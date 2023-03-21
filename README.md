# golang_helper

## Motivation
> Improve code reusability and encapsulate commonly used functions

### reference
> source folder introduce
- `http` Common functions related to http request and response
- `json`JSON complex data type processing common functions

#### json
```go
key := "complex.0.os.1"
	jsonData := `{"name":"kings.chan","lang":["java","go"],"complex":[{"os":["linux","mac","windows"],"reassure":true}]}`
	exists, val := json.GetByExp(jsonData, key)
	if exists {
		fmt.Println(key,"key is exists and the value is :", val)
	} else {
		fmt.Println(key, "is not exists !")
	}
```
The result of running the above code is
`complex.0.os.1 key is exists and the value is : mac`