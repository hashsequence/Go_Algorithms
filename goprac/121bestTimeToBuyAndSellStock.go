/*
121. Best Time to Buy and Sell Stock
Easy

Say you have an array for which the ith element is the price of a given stock on day i.

If you were only permitted to complete at most one transaction (i.e., buy one and sell one share of the stock), design an algorithm to find the maximum profit.

Note that you cannot sell a stock before you buy one.

Example 1:

Input: [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
             Not 7-1 = 6, as selling price needs to be larger than buying price.
Example 2:

Input: [7,6,4,3,1]
Output: 0
Explanation: In this case, no transaction is done, i.e. max profit = 0.

solution:
must do in one pass
pls think it through first even though it seems easy.
first case is first two days
set selling price, buying price, and profit
iterate over prices 
if the price on day i is greater than selling price, then set the selling price as this new price
if the price on day i is smaller than the buying price, reset the buying price to day i's and set selling price to the next day as the minimum day to sell,
we can do this because we accounted for all the profit margin with the previous buying day
*/
func maxProfit(prices []int) int {
    if len(prices) <= 1{
        return 0
    }
    
    buy := prices[0]
    sell := prices[1]
    profit := sell-buy
    for i:= 1; i < len(prices); i++ {
        if prices[i] > sell {
            sell = prices[i]
        }
        if i < len(prices) - 1 && prices[i] < buy {
            buy = prices[i]
            sell = prices[i+1]
        }
        if sell - buy > profit {
            profit = sell - buy
        }
    }
    if profit < 0 {
        return 0
    }
    return profit
}
