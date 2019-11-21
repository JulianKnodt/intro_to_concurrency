package main

import (
  "log"
  "net/http"
  "bufio"
  "strings"
  "io/ioutil"
)

const linkFile = "links.txt"

func main() {
  links, _ := ioutil.ReadFile(linkFile)
  reader := bufio.NewReader(strings.NewReader(string(links)))
  scanner := bufio.NewScanner(reader);

  total := 0
  for scanner.Scan() {
    total += 1
    text := scanner.Text()
    response, _ := http.Get(text)
    body, _ := ioutil.ReadAll(response.Body)
    resp := string(body)
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
