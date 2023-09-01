# mac-wake-display

A simple Go lib to wake up your display on macOS.

## Install

Install the package with:

```bash
go get github.com/asmoker/mac-wake-display
```

## Example

```go
package main

import (
	mwd "github.com/asmoker/mac-wake-display"
)

func main() {
	mwd.WakeDisplay()
}
```

## Thanks

- [BLEUnlock](https://github.com/ts1/BLEUnlock)
  - [lowlevel.c](https://github.com/ts1/BLEUnlock/blob/master/BLEUnlock/lowlevel.c)
  - [lowlevel.h](https://github.com/ts1/BLEUnlock/blob/master/BLEUnlock/lowlevel.h)

## License

The MIT License (MIT)

Copyright (c) 2023 asmoker, https://github.com/asmoker/mac-wake-display

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
