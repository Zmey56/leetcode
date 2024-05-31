package lemonadeChange

//At a lemonade stand, each lemonade costs $5. Customers are standing in a queue to buy from you and order one at a time
//(in the order specified by bills). Each customer will only buy one lemonade and pay with either a $5, $10, or $20 bill.
//You must provide the correct change to each customer so that the net transaction is that the customer pays $5.
//
//Note that you do not have any change in hand at first.
//
//Given an integer array bills where bills[i] is the bill the ith customer pays, return true if you can provide every
//customer with the correct change, or false otherwise.

func lemonadeChange(bills []int) bool {
	countFive := 0
	countTen := 0

	for _, bill := range bills {
		if bill == 5 {
			countFive++
		} else if bill == 10 {
			if countFive == 0 {
				return false
			}
			countFive--
			countTen++
		} else if bill == 20 {
			if countTen > 0 {
				countTen--
				countFive--
			} else {
				countFive -= 3
			}
			if countFive < 0 {
				return false
			}
		}
	}

	return true

}
