package main

import (
	"fmt"
	"strings"
)

/* Strings are value types and immutable: once created you cannot
modify the contents of the string. Put in another way:
strings are immutable arrays of bytes.
*/

func main() {
	fmt.Println(
		strings.Contains("test", "es"),             // true
		strings.Count("test", "t"),                 // 2
		strings.HasPrefix("test", "te"),            // true
		strings.HasSuffix("test", "st"),            // true
		strings.Index("test", "s"),                 // 2
		strings.Join([]string{"a", "b", "c"}, "-"), // "a-b-c"
		strings.Repeat("x", 5),                     // "xxxxx"
		strings.Replace("aaaaa", "a", "b", 2),      // "bbaaa"
		strings.Split("a-b-c-d-e", "-"),            // [a b c d e] == []string{"a","b","c","d","e"}
		strings.ToLower("TEST"),                    // "test"
		strings.ToUpper("test"),                    // "TEST"
	)
	arr := []byte("test")
	str := string([]byte{'t', 'e', 's', 't'})
	fmt.Println(arr, str)    // [116 101 115 116] test
	fmt.Println(string(arr)) // test
	// raw string
	rstr := `this is Raw
string it can span multiple
lines. \n is not interpreted`
	fmt.Println(rstr)
	// 只有ASCII碼才能使用[index]，中文不行。
	str = "中文字體"
	fmt.Println(str[0], str[2]) //228 173
}
