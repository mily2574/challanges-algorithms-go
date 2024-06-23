func diagonalDifference(arr [][]int32) int32 {
  var firstDiagonal, secondDiagonal int32
    for r := 0; r < len(arr); r++ {
        firstDiagonal += arr[r][r]
    }
    for c := 0; c <len(arr); c++ {
        secondDiagonal += arr[c][len(arr)-1-c]
    }
    diff := math.Abs(float64(firstDiagonal - secondDiagonal))
  return int32(diff)
}
