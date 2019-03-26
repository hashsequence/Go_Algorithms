func removeInvalidParentheses(s string) []string {
    res := []string{}
    dfs_remove(s, &res, 0, 0, &Parens{byte('('), byte(')')})
    return res
}

func dfs_remove(s string, res *[]string, last_l int, last_r int, p *Parens) {
    stack := 0
    r := last_r
    //iterates from left to right and checks for extra closing brases
    for ;r < len(s); r++ {
        if s[r] == p.Left {
            stack++
        }
        if s[r] == p.Right {
            stack--
        }
        //if there is extra closing braces then proceed to iterate from last_l to r
        if stack < 0 {
            for l := last_l; l <= r; l++ {
                //check if the current character is the closing braces and if the current position is equal to the last removal position or if the previous 
                //character is not a closing brace then recursive call dfs_remove with the character taken out and the current l position and r
                if s[l] == p.Right && (l == last_l || s[l-1] != p.Right) {
                    dfs_remove(s[0:l]+s[l+1:len(s)], res, l, r, p)
                }
            }
            return
        }
    }
    //in the previous for loop we looked for invalid parens while iterating left to right, and there isnt any
    //we then reverse the string so we look for invalid parens in the other direction right to left
    reversed := Reverse(s)
    //if statement here make sure this is only called after we go left to right
    if p.Left == byte('(') {
        dfs_remove(reversed, res, 0, 0, &Parens{byte(')'), byte('(')})
        //we can only reach this statement if s in this recursive call is valid from left to right and right to left
        //so we know the string isvalid then add to our result
    } else {
        *res = append(*res, reversed)
    }
}

type Parens struct {
    Left byte
    Right byte
}

func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}
