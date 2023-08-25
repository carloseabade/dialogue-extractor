package main

import (
  "bufio"
  "errors"
  "fmt"
  "log"
  "os"
)

func main() {
  extractDialogue("sub.ass")
}

func extractDialogue(path string) {
  f, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  scanner := bufio.NewScanner(f)
  scanner, err = removePreDialogueInfo(scanner)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  fmt.Println(scanner.Text())
}

func removePreDialogueInfo(s *bufio.Scanner) (*bufio.Scanner, error) {
  for s.Scan() {
    if s.Text() == "[Events]" {
      return s, nil
    }
  }
  return nil, errors.New("There is no '[Events]' tag in the file.")
}
