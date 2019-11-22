package main

import (
  "sync"
)

func main() {
  items := make([]int, 0)
  // var mu sync.Mutex
  var wg sync.WaitGroup
  wg.Add(1)
  for i := 0; i < 10; i++ {
    wg.Add(1)
    go func() {
      for j := 0; j < 100; j ++ {
        // mu.Lock()
        items = append(items, i)
        // mu.Unlock()
      }
      wg.Done()
    }()
  }
  wg.Done()
  wg.Wait()
  println(len(items))
}
