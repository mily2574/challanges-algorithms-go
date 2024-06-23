func aVeryBigSum(ar []int64) int64 {
    // Write your code here
    var long int64
    for i := 0; i <len(ar); i++{
        long += ar[i]
    }
    return long 
}
