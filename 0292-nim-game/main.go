package main

func canWinNim(n int) bool {
    if n % 4 > 0 {
        return true
    }

    return false
}
