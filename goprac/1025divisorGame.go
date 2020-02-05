/*
1025. Divisor Game
Easy

Alice and Bob take turns playing a game, with Alice starting first.

Initially, there is a number N on the chalkboard.  On each player's turn, that player makes a move consisting of:

Choosing any x with 0 < x < N and N % x == 0.
Replacing the number N on the chalkboard with N - x.
Also, if a player cannot make a move, they lose the game.

Return True if and only if Alice wins the game, assuming both players play optimally.

 

Example 1:

Input: 2
Output: true
Explanation: Alice chooses 1, and Bob has no more moves.
Example 2:

Input: 3
Output: false
Explanation: Alice chooses 1, Bob chooses 1, and Alice has no more moves.
 

Note:

1 <= N <= 1000

soltion:
O(1) solution is check if it is odd or even, since alice always wins if N is even

dp way is that you make a boolean array of size N,
we iterate through the possible divisors which are always less than equal to sqrt(N) which means we travel through
each possible choice for each player at every turn
to cut the seach space we use dp to remember if the player will loose at N number at their turn no matter one so at dp[i] we see that it is already true
we dont need to go further down that path again and just return 

*/
func divisorGame(N int) bool {
    dp := make([]bool,N+1)
    dp[0] = true
    return dfs(N, dp)
}

func dfs (N int, dp []bool) bool {
    if dp[N] {
        return true
    }
    
    for i := 1; i * i <= N; i++ {
        if N % i == 0 {
            //everytime we call dfs its the other player turn, so if dfs returns true, the caller of that dfs function loose
            if !dfs(N-i, dp) {
                fmt.Println(N)
                dp[N] = true
                break
            }
        }
    }
    return dp[N]
    
}
