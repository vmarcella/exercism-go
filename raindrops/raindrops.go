package raindrops

import "strconv"

func Convert(num int) string {
    /*
        Convert the number into raindrop sounds if the raindrop is divisble by 
        either 3, 5, 7

        Arguments:
            :num: A number  
    */
    outputString := ""
    factors := [3]int{3, 5, 7}
    rainSounds := [3]string{"Pling", "Plang", "Plong"}

    // iterate through all the factors   
    for i := 0; i < len(factors); i++ {
        if num % factors[i] == 0 {
            outputString += rainSounds[i]
        }
    }
   
    // Check to see if the length of the output string
    if len(outputString) == 0 {
        return strconv.Itoa(num) 
    }

    return outputString
}
