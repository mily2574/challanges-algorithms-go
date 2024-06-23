func miniMaxSum(arr []int32) {
   min, max := arr[0], arr[0]
   var sum int
   for _, num := range arr{
       if num < min {
           min = num
       }
       if num > max {
           max = num
       }
       sum += int(num)
   }
   minSum := sum-int(max)
   maxSum := sum-int(min)
   fmt.Printf("%d %d\n", minSum,maxSum)
}
