package isbn

import "strconv"

func IsValidISBN(isbn string) bool {

    // Check if the isbn is too small
    if len(isbn) < 10 {
        return false
    }

    var numList = []int{};

    // Process the ISBN into a num list
    for pos, char := range isbn {
        if char != '-' {
            if value, err := strconv.Atoi(string(char)); err == nil {
                numList = append(numList, value)
            }else{
                if (char == 'X' || char == 'x') && pos == len(isbn) - 1 {
                    numList = append(numList, 10) 
                } else {
                    return false
                }
            }
        }
    }

    if len(numList) < 10 || len(numList) > 10{
        return false 
    } else {
        counter := 10
        sum := 0
        for i := 0; i < len(numList); i++ {
            sum += numList[i] * counter
            counter -= 1
        }
        return sum % 11 == 0
    }
}
