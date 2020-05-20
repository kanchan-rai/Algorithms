package main

 import (
         "fmt"
 )

 func PercentageChange(math, hindi, english int) (delta float64) {
         diff := float64(math + hindi + english)
         delta = (diff / float64(300)) * 100
         return
 }

 func main() {

		 fmt.Printf("Kanchan Kumar Results = Math : 85, Hindi : 55 English : 43 Percentage : %0.2f %% \n", PercentageChange(85, 55, 43))

		 fmt.Printf("Rahul Kumar Results = Math : 65, Hindi : 82 English : 48 Percentage : %0.2f %% \n", PercentageChange(65, 82, 48))

		 fmt.Printf("Karan Rai Results = Math : 89, Hindi : 92 English : 86 Percentage : %0.2f %% \n", PercentageChange(89, 92, 86))

		 fmt.Printf("Akhil Kumar Results = Math : 94, Hindi : 56 English : 74 Percentage : %0.2f %% \n", PercentageChange(94, 56, 74))

		 fmt.Printf("Nilesh Kumar Results = Math : 34, Hindi : 56 English : 26 Percentage : %0.2f %% \n", PercentageChange(34, 56, 26))

 }