# assert.go

> golang assert library

![go report](https://goreportcard.com/badge/github.com/ooopSnake/assert.go)
![license](https://img.shields.io/badge/license-MIT-brightgreen.svg)
[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://github.com/ooopSnake/assert.go)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat)](https://github.com/ooopSnake/assert.go/pulls)
[![Ask Me Anything !](https://img.shields.io/badge/Ask%20me-anything-1abc9c.svg)](https://github.com/ooopSnake/assert.go/issues)


## Get Start

```bash
# install
go get -u "github.com/ooopSnake/assert.go"
```



```go
// import and use
package your_package

import "github.com/ooopSnake/assert.go"
```


## Example

### direct panic
```go
func foo(value int) error {
    assert.Must(value > 10,"value too small").Panic()
    return nil
}
```

### convert panic to an error
```go
func foo(value int) error {
	if err := assert.Must(value > 10, "value too small").Error(); err != nil {
		return err
	}
	return nil
}

```

## License

Released under the [MIT License](https://github.com/ooopSnake/assert.go/blob/master/LICENSE)





