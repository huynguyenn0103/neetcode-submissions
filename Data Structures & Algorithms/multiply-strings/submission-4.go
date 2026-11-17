
func multiply(num1 string, num2 string) string {
    m, n := len(num1), len(num2)
    result := make([]int, m+n)

    for i := m - 1; i >= 0; i-- {
        for j := n - 1; j >= 0; j-- {
            a := int(num1[i] - '0')
            b := int(num2[j] - '0')

            product := a*b + result[i+j+1]

            result[i+j+1] = product % 10
            result[i+j] += product / 10
        }
    }

    // Convert the result array into a string.
    var sb strings.Builder

    for _, digit := range result {
        if sb.Len() == 0 && digit == 0 {
            continue
        }
        sb.WriteByte(byte(digit) + '0')
    }

    if sb.Len() == 0 {
        return "0"
    }

    return sb.String()
}
