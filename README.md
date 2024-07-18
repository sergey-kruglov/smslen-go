# smslen

**smslen** is a small Go package designed to help developers calculate the number of SMS messages required to send a given text.

Inspired by [sms-length](https://github.com/inkOfPixel/sms-length) Node.JS package.

## Installation

```bash
go get github.com/sergey-kruglov/smslen-go
```

## Usage

```go
package main

import "github.com/sergey-kruglov/smslen-go"

func main(text string) {
  text := "Hello, World!"
  sms := smslen.Count(text)
  //  {
  //    Encoding: "GSM_7BIT",
  //    Chars: 13,
  //    CharsInPart: 160,
  //    Parts: 1,
  //  }
}
```
