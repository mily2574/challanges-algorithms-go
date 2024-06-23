//Solution 1:
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v2 > v1 || v1-v2 == 0 {
		return "NO"
	}
	if (x2-x1)%(v1-v2) == 0 {
		return "YES"
	}
	return "NO"
}

//Recursive solution:
func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	if v2 >= v1 || x1 > x2 {
		return "NO"
	}
	if x1 == x2 {
		return "YES"
	}
	return kangaroo(x1+v1, v1, x2+v2, v2)
}
