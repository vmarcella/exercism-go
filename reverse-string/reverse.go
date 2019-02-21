package reverse

func String(inputStr string) string {
    var outputStr []byte;

    for i := len(inputStr) - 1; i >= 0; i-- {
        outputStr = append(outputStr, inputStr[i])
    }

    return string(outputStr)
}
