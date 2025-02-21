func MyFunc(x int) (int, error) {
    if x == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return 10 / x, nil
}

func main() {
    result, err := MyFunc(0)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)

    result, err = MyFunc(2)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Println("Result:", result)
}