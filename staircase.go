func staircase(n int32) {
    // Write your code here
    for i:=int32(1);i<n+1;i++ {
        for j := int32(0); j<n-i;j++ {
            fmt.Print(" ")
        }
        for j := int32(0); j<i;j++ {
            fmt.Print("#")
        }
        fmt.Println()
    }
}
