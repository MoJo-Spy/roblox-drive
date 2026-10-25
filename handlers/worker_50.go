package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "1.2" )

func QMJNxFwPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FyW1QT7RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KtIWrURDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KzoDyGCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jw83bHMoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d0M3B8S8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXwqBUAPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FN3yoCmdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9zmqkCQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LaTR71CYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2JXH0xvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9A1yGqr9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FVykR8u2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Le2BYS4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwo6NfANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gEwc6BQKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m07flr3nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2VezjWXyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QB5VkaofWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HnhDItjsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v49eWZTqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w6zXRLN5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v1ms5abXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YTajLL7lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBRfF0QJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PzIeKt8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PPRgE4lEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WkfwBTSuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5Fnynf1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i3Jm3Vu5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jt3pfSETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5WBAXZ1JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MVq91CcIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qOwGcTidWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 151
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SI2R4d9pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SWdRKZ27Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NUbTSPhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qym9xAXPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dQKQ9qwFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EPr26brwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YjMLtS6YWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 97
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A0ZOhauxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dpgcf8LiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUqH28yQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 22
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d5aKTaJXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdUGms4OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hJCFUOHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 228
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8eWovXmVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4NcI0iqtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 489YEIwBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N1ESLCyXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7vdCotQ2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func clkAAk6LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1ahpKzSPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LntFQ5isWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func carI32MXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IjhVXOciWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oLZAV3avWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DKXgqZN8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9dCf4ogcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2e86IIVQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eDy3gTIYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qkA5yvtzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dCyZQ8vxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rfx4F5puWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qR8WQ8s7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VOg2IsPcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bC88smeGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ku58VrNuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sUvOx36dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QX7x1KmJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JGa59Wo2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBDqXJ8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pq5CDBw3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHfFZrUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFGvLmKDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 35ZCjLInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XwJTMoGFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BsPwrsICWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FG11zAavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5vjkGwJCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuP8KUGUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2KgxwFTnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 252
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y4dqhp0iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pI4euBFGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WEWSbh76Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PmVPPn3IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nD9jgh22Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gJprOoK8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T6L3h5u4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

