package utils

//InSlice function ...
func InSlice(a []string, x string) bool {
	for _, n := range a {
	   if x == n {
		  return true
	   }
	}
	return false
}

//ContainsInt function ...
func ContainsInt(a []int, x int) bool {
	for _, n := range a {
	   if x == n {
		  return true
	   }
	}
	return false
  }