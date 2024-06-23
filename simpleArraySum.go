func simpleArraySum(ar []int32) int32 {
    // Write your code here
    l := len(ar)
    s := ar[0] 
    for i := 1; i< l; i++{
        s += ar[i]
    }
    return s
}
