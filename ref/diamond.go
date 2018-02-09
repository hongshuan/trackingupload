package main

import "fmt"

// 关于Go，有件事令我很惊讶 —— 可以实现多重继承。确实很糟糕。
// Go并不能很好的处理继承的歧义。代码展示了著名的“可怕的钻石问题”
// (Dreaded diamond problem）：
//
//    T4(foo)
//    /    \
//   /      \
//  T3     T2(foo)
//   \      /
//    \    /
//      T1

type T1 struct {
    T2
    T3
}

type T2 struct {
    T4
    foo int
}

type T3 struct {
    T4
}

type T4 struct {
    foo int
}

func main() {
    t2 := T2{ T4{ 9000 }, 2 }
    t3 := T3{ T4{ 3 } }
    fmt.Printf("t2.foo=%d\n", t2.foo)
    fmt.Printf("t3.foo=%d\n", t3.foo)
    t1 := T1{
        t2,
        t3,
    }
    fmt.Printf("t1.foo=%d\n", t1.foo)

    fmt.Printf("t2.T4.foo=%d\n", t2.T4.foo)
}
