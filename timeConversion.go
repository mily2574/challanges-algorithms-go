func timeConversion(s string) string {
        if strings.HasSuffix(s, "AM") && s[0:2] == "12" {
        s = strings.Replace(s, "12", "00", 1)
    } else if strings.HasSuffix(s, "PM") && s[0:2] != "12" {
        hour := fmt.Sprintf("%c%c", s[0]+1, s[1]+2)
        s = strings.Replace(s, s[0:2], hour, 1)
    }
    return s[0:8]
}
