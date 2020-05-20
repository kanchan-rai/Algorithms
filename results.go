package main

 import (
         "fmt"
 )

 func PercentageChange(subject1, subject2 int) (delta float64) {
         diff := float64(subject2 + subject1)
         delta = (diff / float64(200)) * 100
         return
 }

 func main() {

         fmt.Printf("subject1 : 85, subject2 : 55 Percentage : %0.2f %% \n", PercentageChange(85, 55))

         fmt.Printf("subject1 : 65, subject2 : 50 Percentage : %0.2f %% \n", PercentageChange(65, 50))

         fmt.Printf("subject1 : 95, subject2 : 30 Percentage : %0.2f %% \n", PercentageChange(95, 30))
		 

 }