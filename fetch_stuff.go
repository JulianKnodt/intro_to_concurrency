package main

import (
  "log"
  "net/http"
  "bufio"
  "io/ioutil"
  "os"
  "sync"
)

const linkFile = "links.txt"

func main() {
  // responseMutex := sync.Mutex{}
  responses := make([]string, 0)

  var wg sync.WaitGroup
  links, _ := os.Open(linkFile)
  reader := bufio.NewReader(links)
  scanner := bufio.NewScanner(reader);
  total := 0
  for scanner.Scan() {
    total += 1
    wg.Add(1)
    go func() {
      resp, err := http.Get(scanner.Text())
      if err != nil {
        log.Println(err)
        return
      }
      /* Is this mutex necessary?
      responseMutex.Lock()
      defer responseMutex.Unlock()
      */
      body, _ := ioutil.ReadAll(resp.Body)
      responses = append(responses, string(body))
      wg.Done()
    }()
  }
  wg.Wait()
  for _, resp := range responses {
    log.Println(resp)
  }
  log.Printf("Done fetching %d links", total)
}
