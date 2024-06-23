func birthdayCakeCandles(candles []int32) int32 {
    var max, count int32
    for _, num := range candles {
        if num > max {
            max = (num)
        }
    }
    for _, num := range candles {
        if num >= max {
            count++
        }
    }
   return count
}
