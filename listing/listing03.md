Что выведет программа? Объяснить вывод программы. Объяснить внутреннее устройство интерфейсов и их отличие от пустых интерфейсов.

```go
package main

import (
	"fmt"
	"os"
)

func Foo() error {
	var err *os.PathError = nil
	return err
}

func main() {
	err := Foo()
	fmt.Println(err)
	fmt.Println(err == nil)
}
```

Ответ:
```
os.PathError nill
false
Так как из Foo возвращается типизированный nil, то он не будет равен nil без типа. Чтобы получить true, нужно привести nil к (*os.PathError)(nil)

```
