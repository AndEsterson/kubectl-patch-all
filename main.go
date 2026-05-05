package main

import (
  "context"
  "os"
  "os/signal"
  "syscall"
  "github.com/andesterson/kubectl-patch_all/pkg/cmd"
)

var version = "HEAD"

func main() {
  ctx := context.Background()
  ctx, stop := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
  defer stop()
  os.Exit(cmd.Run())  
}
