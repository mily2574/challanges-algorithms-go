func countApplesAndOranges(s int32, t int32, a int32, b int32, apples []int32, oranges []int32) {
    var n_apples, n_oranges int32
    //The number of apples that fall on Sam's house
    for i :=  range apples {
        apples[i] += a 
        if apples[i] >= s && apples[i] <= t {
                n_apples++
        }
    }
    fmt.Println(n_apples)
    //The number of oranges that fall on Sam's home
    for i :=  range oranges {
        oranges[i] += b 
        if oranges[i] >= s && oranges[i] <= t {
            n_oranges++
        }
    }
    fmt.Println(n_oranges)
}
