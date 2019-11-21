package main

func prod() chan int {
  out := make(chan int, 1)
  go func() {
    for i := 0;; i++ {
      out <- i
    }
  }()
  return out
}

func main() {
  out1 := prod()
  out2 := prod()
  remaining := 10
  for remaining > 0 {
    select {
    case i := <-out1:
      println(i)
    case i := <-out2:
      println(i)
    }
    remaining -= 1
  }
}
