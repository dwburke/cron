# golang "github.com/robfig/cron" embedded


## main.go

```
package main

import (
	"github.com/dwburke/atexit"

	"github.com/dwburke/monner/cmd"
)

func main() {
	defer atexit.AtExit()
	cmd.Execute()
}
```

## cmd/root.go

```
package cmd

import (
   ...
	"github.com/dwburke/cron"
   ...
)

func init() {
	cobra.OnInitialize(initConfig)
	cobra.OnInitialize(cron.Run)
}

var rootCmd = &cobra.Command{
	...
}
```

