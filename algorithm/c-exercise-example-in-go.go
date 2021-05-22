/**
 * @authors Lewis Tian (taseikyo@gmail.com)
 * @date    2021-05-18 15:11:23
 * @link    github.com/taseikyo
 */

package main

import (
    "fmt"
    "math"
)

func t1() {
    sum := 1 + 2 + 3 + 4
    for i := 1; i < 5; i++ {
        for j := 1; j < 5; j++ {
            for k := 1; k < 5; k++ {
                if i != j && j != k && i != k {
                    fmt.Println(i, j, k, sum-i-j-k)
                }
            }
        }
    }
}

func t2() {
    var money float64
    fmt.Print("Input your money: ")
    fmt.Scanf("%f", &money)
    var bonus float64
    bonus1 := 100000 * 0.1
    bonus2 := bonus1 + 100000*0.075
    bonus4 := bonus2 + 200000*0.05
    bonus6 := bonus4 + 200000*0.03
    bonus10 := bonus6 + 400000*0.015
    if money <= 100000 {
        bonus = money * 0.1
    } else if money <= 200000 {
        bonus = bonus1 + (money-100000)*0.075
    } else if money <= 400000 {
        bonus = bonus2 + (money-200000)*0.05
    } else if money <= 600000 {
        bonus = bonus4 + (money-400000)*0.03
    } else if money <= 1000000 {
        bonus = bonus6 + (money-600000)*0.015
    } else {
        bonus = bonus10 + (money-1000000)*0.01
    }
    fmt.Println(bonus)
}

func t3() {
    var i, j, m, n, x int
    for i = 1; i < 168/2+1; i++ {
        if 168%i == 0 {
            j = 168 / i
            if i > j && (i+j)%2 == 0 && (i-j)%2 == 0 {
                m = (i + j) / 2
                n = (i - j) / 2
                x = n*n - 100
                fmt.Printf("%d + 100 = %d * %d\n", x, n, n)
                fmt.Printf("%d + 268 = %d * %d\n", x, m, m)
            }
        }
    }
}

func t4() {
    var day, month, year, sum, leap int
    fmt.Print("Input: year month day: ")
    fmt.Scanf("%d %d %d", &year, &month, &day)
    switch month {
    case 1:
        sum = 0
    case 2:
        sum = 31
    case 3:
        sum = 59
    case 4:
        sum = 90
    case 5:
        sum = 120
    case 6:
        sum = 151
    case 7:
        sum = 181
    case 8:
        sum = 212
    case 9:
        sum = 243
    case 10:
        sum = 273
    case 11:
        sum = 304
    case 12:
        sum = 334
    default:
        sum = 0
    }
    sum = sum + day
    if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
        leap = 1
    } else {
        leap = 0
    }
    // *如果是闰年且月份大于2,总天数应该加一天
    if leap == 1 && month > 2 {
        sum++
    }
    fmt.Printf("这是这一年的第 %d 天。\n", sum)
}

func t5() {
    var x, y, z int
    fmt.Print("请输入三个数字：")
    fmt.Scanf("%d %d %d", &x, &y, &z)
    if x > y {
        x, y = y, x
    }
    if x > z {
        x, z = z, x
    }
    if y > z {
        y, z = z, y
    }
    fmt.Printf("从小到大排序: %d %d %d\n", x, y, z)
}

func t6() {
    fmt.Printf("用 * 号输出字母 C!\n")
    fmt.Printf(" ****\n")
    fmt.Printf(" *\n")
    fmt.Printf(" * \n")
    fmt.Printf(" ****\n")
}

func t7() {
    var a, b = 180, 210
    fmt.Printf("%c%c%c%c%c\n", b, a, a, a, b)
    fmt.Printf("%c%c%c%c%c\n", a, b, a, b, a)
    fmt.Printf("%c%c%c%c%c\n", a, a, b, a, a)
    fmt.Printf("%c%c%c%c%c\n", a, b, a, b, a)
    fmt.Printf("%c%c%c%c%c\n", b, a, a, a, b)
}

func t8() {
    size := 9
    for i := 1; i <= size; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("%2d * %2d = %2d ", i, j, i*j)
        }
        fmt.Println()
    }

}

func t9() {
    size := 8
    for i := 1; i <= size; i++ {
        for j := 1; j <= size; j++ {
            if (i+j)%2 == 0 {
                fmt.Print("⬛")
            } else {
                fmt.Print("⬜")
            }
        }
        fmt.Println()
    }

}

func t10() {
    size := 11
    fmt.Printf("😁😁\n")
    for i := 1; i <= size; i++ {
        for j := 1; j <= i; j++ {
            fmt.Printf("⬛⬛")
        }
        fmt.Println()
    }

}

func t11() {
    f1, f2 := 1, 1
    size := 40
    for i := 0; i < size; i++ {
        fmt.Printf("%20d %20d", f1, f2)
        f1 = f1 + f2
        f2 = f1 + f2
        if i%2 == 0 {
            fmt.Println()
        }
    }
}

func t12() {
    lower, upper := 101, 201
    for i := lower; i < upper; i++ {
        flag := false
        for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
            if i%j == 0 {
                flag = true
                break
            }
        }
        if !flag {
            fmt.Printf("%d ", i)
        }

    }
}

func t13() {
    var x, y, z int
    for i := 100; i < 1000; i++ {
        x = i % 10
        y = i / 10 % 10
        z = i / 100 % 10
        if i == (x*x*x + y*y*y + z*z*z) {
            fmt.Printf("%d ", i)
        }
    }
}

func t14() {
    var n int
    fmt.Printf("请输入整数：")
    fmt.Scanf("%d", &n)
    fmt.Printf("%d=", n)
    for i := 2; i <= n; i++ {
        for n%i == 0 {
            fmt.Printf("%d", i)
            n /= i
            if n != 1 {
                fmt.Printf("*")
            }
        }
    }
}

func t15() {
    var score int
    var grade int8
    fmt.Printf("请输入分数： ")
    fmt.Scanf("%d", &score)
    if score >= 90 {
        grade = 'A'
    } else if score >= 60 {
        grade = 'B'
    } else {
        grade = 'C'
    }
    fmt.Printf("%c\n", grade)
}

func t16() {
    var a, b, r, n int
    fmt.Printf("请输入两个数字：")
    fmt.Scanf("%d %d", &a, &b)
    if a < b {
        a, b = b, a
    }
    r = a % b
    n = a * b
    for r != 0 {
        a = b
        b = r
        r = a % b
    }
    fmt.Printf("这两个数的最大公约数是%d，最小公倍数是%d\n", b, n/b)
}

func t17() {
    var c int8
    letters, spaces, digits, others := 0, 0, 0, 0
    fmt.Printf("请输入一些字母：")
    for {
        fmt.Scanf("%c", &c)
        if c == '\n' {
            break
        }
        if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
            letters++
        } else if c >= '0' && c <= '9' {
            digits++
        } else if c == ' ' {
            spaces++
        } else {
            others++
        }
    }
    fmt.Printf("字母=%d,数字=%d,空格=%d,其他=%d\n", letters, digits, spaces, others)
}

func t18() {
    a, n, sum := 0, 0, 0
    fmt.Printf("请输入 a 和 n：")
    fmt.Scanf("%d%d", &a, &n)
    for n > 0 {
        sum += a * n
        a *= 10
        n--
    }
    fmt.Printf("sum: %d\n", sum)
}

func t19() {
    upper := 1000
    for i := 1; i < upper; i++ {
        sum := 0
        for j := 1; j <= i/2; j++ {
            if i%j == 0 {
                sum += j
            }
        }
        if i == sum {
            fmt.Printf("%d ", i)
        }
    }
}

func t20() {
    count := 10
    h := 100.0
    sum := h
    h /= 2
    for i := 1; i < count; i++ {
        sum += 2*h
        h /= 2
    }
    fmt.Printf("第 10 次落地时，共经过 %f 米，第 10 次反弹高 %f 米\n", sum, h)
}

func main() {
    func_list := []func(){t1,
        t2,
        t3,
        t4,
        t5,
        t6,
        t7,
        t8,
        t9,
        t10,
        t11,
        t12,
        t13,
        t14,
        t15,
        t16,
        t17,
        t18,
        t19,
        t20,
    }
    var idx int
    var ch int8

    for {
        fmt.Printf("Input function idx (0 ~ %d, 0 means execute all functions): ", len(func_list))
        fmt.Scan(&idx)

        if idx > len(func_list) {
            idx = len(func_list)
        }
        if idx == 0 {
            for i := 0; i < len(func_list); i++ {
                fmt.Printf("==========%d==========\n\n", i+1)
                func_list[i]()
            }
        } else {
            func_list[idx-1]()
        }

        fmt.Printf("\nInput q/Q to quit: ")
        fmt.Scanf("%c", &ch)
        if ch == 'q' || ch == 'Q' {
            break
        }
    }
}
