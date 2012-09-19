package main

import (
  "fmt"
  "strings"
)

func main() {
  fmt.Println(
    strings.Contains("test", "es"),   // true
    strings.Count("test", "t"),       // 2
    strings.HasPrefix("test", "te"),  // true
    strings.HasSuffix("test", "st"),  // true
    strings.Index("test", "s"),       // 2
    strings.Join([]string{"a","b","c"},"-"),  // "a-b-c"
    strings.Repeat("x", 5),           // "xxxxx"
    strings.Replace("aaaaa","a","b",2),       // "bbaaa"
    strings.Split("a-b-c-d-e","-"),   // [a b c d e] == []string{"a","b","c","d","e"}
    strings.ToLower("TEST"),          // "test"
    strings.ToUpper("test"),          // "TEST"
  )
  arr := []byte("test")
  str := string([]byte{'t','e','s','t'})
  fmt.Println(arr, str)
  fmt.Println(string(arr))
}

