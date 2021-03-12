# Webpack_Sass

Instead of having to open a third terminal
which is kind of annoying in my opinion, I
want to be able to run the webpack compiler
and sass compiler in the same terminal with just
one commmand and preferably communicate through std.
so you get error messages etc..

```go
package main

import (
	"fmt"

	"github.com/DustiasTheGuy/webpack_sass/webpack_sass"
)

func main() {
	cfg := webpack_sass.Initalize()

	fmt.Println(cfg)
}
```
