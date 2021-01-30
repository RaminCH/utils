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

//InSliceInt function ...
func InSliceInt(a []int, x int) bool {
	for _, n := range a {
	   if x == n {
		  return true
	   }
	}
	return false
  }