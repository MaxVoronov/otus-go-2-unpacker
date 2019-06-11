# Otus Homework #2
## Распаковка строки

Создать Go функцию, осуществляющую примитивную распаковку строки, содержащую повторяющиеся символы / руны, например:

* "a4bc2d5e" => "aaaabccddddde"
* "abcd" => "abcd"
* "45" => "" (некорректная строка)
* \`qwe\4\5\` => \`qwe45\' (*)
* \`qwe\45\` => \`qwe44444\` (*)
* \`qwe\\5\` => \`qwe\\\\\\\\\\\` (*)

### Как использовать

```
package main

import (
    "github.com/maxvoronov/otus-go-2-unpacker"
    "log"
)

func main() {
    result, err := unpacker.Unpack("a4bc2d5e")
    if err != nil {
        log.Panic("Error: %s", err)
    }

    log.Printf("Result: %s", result)
    // 2019/06/11 22:32:31 Result: aaaabccddddde
}
```
