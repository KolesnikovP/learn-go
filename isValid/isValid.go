package main
// BEGIN (write your solution here)
func IsValid(id int, text string) bool {
  isText := text != ""
  isMoreZero := id != 0
  
  res := isMoreZero && isText
  return res
}
// END
