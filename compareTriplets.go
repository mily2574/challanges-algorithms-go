func compareTriplets(a []int32, b []int32) []int32 {
  var a_score, b_score int32
  for i, j := 0, 0 ; i < len(a) && j < len(b);  i, j = i+1, j+1 {
    if a[i] > b[j] {
      a_score ++
    } else if a[i] < b[j] {
      b_score ++
    }
  }
  return []int32{a_score, b_score}
}
