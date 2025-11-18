package main

import (
    "fmt"
)

func main() {
    m := make(map[string]int, 10)

    m["1"] = int(1)
    m["2"] = int(2)
    m["3"] = int(3)
    m["4"] = int(4)
    m["5"] = int(5)
    m["6"] = int(6)

    // 获取元素
    value1 := m["1"]
    fmt.Println("m[\"1\"] =", value1)

    value1, exist := m["1"]
    fmt.Println("m[\"1\"] =", value1, ", exist =", exist)

    valueUnexist, exist := m["10"]
    fmt.Println("m[\"10\"] =", valueUnexist, ", exist =", exist)

    // 修改值
    fmt.Println("before modify, m[\"2\"] =", m["2"])
    m["2"] = 20
    fmt.Println("after modify, m[\"2\"] =", m["2"])

    // 获取map的长度
    fmt.Println("before add, len(m) =", len(m))
    m["10"] = 10
    fmt.Println("after add, len(m) =", len(m))

    // 遍历map集合main
    for key, value := range m {
        fmt.Println("iterate map, m[", key, "] =", value)
    }

    // 使用内置函数删除指定的key
    _, exist_10 := m["10"]
    fmt.Println("before delete, exist 10: ", exist_10)
    delete(m, "10")
    _, exist_10 = m["10"]
    fmt.Println("after delete, exist 10: ", exist_10)

    // 在遍历时，删除map中的key
    for key := range m {
        fmt.Println("iterate map, will delete key:", key)
        delete(m, key)
    }
    fmt.Println("m = ", m)
}
