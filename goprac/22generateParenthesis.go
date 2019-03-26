func generateParenthesis(n int) []string {
    res := []string{}
    if n == 0 {
        return res
    } 
    generateParensHelper(&res, "", n, n)
    return res
}

func generateParensHelper(res *[]string, s string, openParens int, closeParens int) {
    //fmt.Println(s)
    if openParens == 0 && closeParens == 0 {
        *res = append(*res, s)
    }
    if openParens > 0 {
        generateParensHelper(res, s+string('('),openParens-1,closeParens)
    }
    if closeParens > openParens {
        generateParensHelper(res, s+string(')'), openParens, closeParens-1)
    }
}
