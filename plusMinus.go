func plusMinus(arr []int32) {
    l := len(arr)
    var n_pos, n_neg, n_zero float64  
    for i:=0; i < l; i++ {
        if arr[i] > 0 {
            n_pos++
        } else if arr[i] < 0 {
            n_neg++
        } else {
            n_zero++
        }
    }
    ratio_pos := (n_pos/float64(l))
    ration_neg := (n_neg/float64(l))
    ration_zero := (n_zero/float64(l))
    fmt.Printf("%.6f\n",ratio_pos)
    fmt.Printf("%.6f\n",ration_neg)
    fmt.Printf("%.6f\n",ration_zero)    
}
