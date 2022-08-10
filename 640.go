func getAandB(e string) (int, int) {
    e = e + "+"
    flag := 1
    a, b := 0, 0
    fac := 0
    haveX := false
    haveNum := false
    for _, c := range e {
        if 'x' == c {
            haveX = true
        } else if '+' == c || '-' == c{
            if haveX {
                if 0 == fac && !haveNum {
                    a += flag
                } else {
                    a += (flag * fac)
                }
            } else {
                b += (flag * fac)
            }
            haveX = false
            fac = 0
            if '+' ==c {
                flag = 1
            } else {
                flag = -1
            }
            haveNum = false
        } else {
            haveNum = true
            tmpN, _ := strconv.Atoi(string(c))
            fac = fac * 10 + tmpN
        }
    }
    // fmt.Println(a, b)
    return a, b
}
func solveEquation(equation string) string {
    s := strings.Split(equation, "=")
    // fmt.Println(s)
    a1, b1 := getAandB(s[0])
    a2, b2 := getAandB(s[1])
    a := a1 - a2
    b := b2 - b1
    if 0 == a {
        if 0 == b {
            return "Infinite solutions"
        } else {
            return "No solution"
        }
    }
    return "x=" + strconv.Itoa(b / a)
}
