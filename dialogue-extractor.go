package main

import (
  "bufio"
  "errors"
  "log"
  "os"
  "strings"
)

func extractDialogue(path string) string {
  f, err := os.Open(path)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  scanner := bufio.NewScanner(f)
  err = removePreDialogueInfo(scanner)
  if err != nil {
    log.Fatal(err)
    os.Exit(1)
  }
  dialogueLines := getDialogues(scanner)
  return strings.Join(dialogueLines, "\n")
}

func removePreDialogueInfo(s *bufio.Scanner) error {
  for s.Scan() {
    if s.Text() == "[Events]" {
      s.Scan()
      return nil
    }
  }
  return errors.New("There is no '[Events]' tag in the file.")
}

func getDialoguePosition(s *bufio.Scanner) int {
  pos := strings.Count(s.Text(), ",")
  return pos
}

func getDialogues(s *bufio.Scanner) []string {
  dPos := getDialoguePosition(s)

  var lines []string
  var text, line string
  for s.Scan() {
    text = s.Text()
    if strings.HasPrefix(text, "Dialogue") {
      line = strings.Join(strings.Split(s.Text(), ",")[dPos:] ,"")
      lines = append(lines, line)
    }
  }
  return lines
}
