package main

import (
  "log"
  "net/http"
  "bufio"
  "io/ioutil"
  "strings"
  "sync"
)

const linkFile = "links.txt"

func main() {

  links, _ := ioutil.ReadFile(linkFile)
  responses := make(chan string, strings.Count(string(links), "\n"))

  reader := bufio.NewReader(strings.NewReader(string(links)))
  scanner := bufio.NewScanner(reader);

  var wg sync.WaitGroup
  var once sync.Once
  for scanner.Scan() {
    wg.Add(1)
    go func(text string) {
      resp, _ := http.Get(text)
      body, _ := ioutil.ReadAll(resp.Body)
      responses <- string(body)
      wg.Done()
      wg.Wait()
      once.Do(func() {
        close(responses)
      })
      // or
      // close(responses)
    }(scanner.Text())
  }
  total := 0
  for resp := range responses {
    total += 1
    log.Println(resp[0:min(len(resp), 50)])
  }
  log.Printf("Done fetching %d links", total)
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}
