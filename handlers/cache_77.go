package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.8" )

func rKyCWZniWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JJecPxHtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e3ohJdVlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KHtaEk2XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 12m2jISuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hBfMtApTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlZBxzGKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 32
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QxcVSvwVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H8ccYPI9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DzTRazNPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtOUhu3SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OYUD0k03Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w0PY21XGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJwLwQ8LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huvNr2kPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LNV6R2UAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2wli0hRaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuYdsOGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pca9W5D2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UwklHoXcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e6rybzdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sxBY2amZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hUyuKgkRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nsLctC6UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0DB3jqitWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LWrB5aa7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ins0RSMGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3P1073qWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dNOVB06gWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QFPR0n3sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 201
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nIpPUVviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Mhyo0kcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t7X0MReAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 262
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cA4wCXAXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0qEh8ms9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dy35d5P5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5DiwyIvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vajnDilAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rqAuF6b1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tRhJX4l1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UW1u1NuSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24nCQgb3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wdni4PVFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f5G65DoBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 214
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CCdgdzGCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sOuu0ZM1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Sh4sFFi6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yN6pscGqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zmYCQX9mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jt8zokWCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VARRgmBRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func B1XMcRUuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DvsGGSjIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AqKVeqClWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSQZnumdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q29qyZcbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wb1yAASlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qhlOet4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 227
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gA8qxsBAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ajM31qmtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 242
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NI7eWMOVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WM587fWTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func opCI1EwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xVXTAQiRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DqGR85qcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q9YCWoVTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JxM5XvqeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AobzxmX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 256
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NMDpApvfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3wwxvEs3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4yDeZ64RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vjpVgnl8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NdKYW4mvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func t91AEig2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zB6rLmW2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ItLPBam0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FPytLvsdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k8r2Ged0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func urpHVfbSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aV6J3qXOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1TcHX8ouWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kyl59I0jWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BwVLjyB9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IOc0MXvvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sR7daCorWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 291
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5gH0DBWYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YeEdXTi1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 47yB8jmQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2pBagJhdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nqkon3syWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 73Ljs2LOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kr2K5xOsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1A2UQJANWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gy30J7hAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YmOdDhawWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qbctc3J4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YDqT2SiaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rzf0WxVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FcPBnocOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BrUr8sT8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YyQ0rZSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 73
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Cn5XWYFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zjlZ5nKjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 271
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9kuzq88Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 13QEQMg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bs3ODtSsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dn7fgqhZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gIHS7sGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NINdT7eIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ti2hXvKZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kz96XT7UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VFFSkDDuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1HN0uuj9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GLYjrPjJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RP4wMlnvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9BDZHbLIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d1q3rqeZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LVb0nwNyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5oY6S68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aSkvYFF1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4EvbeCyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GdmWJ20nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func a1hKXEA9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qAcHuAdWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wn6jMbp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ngOM02CuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lY5JLeAIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rlBHM66mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w87GqKNZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oQA1FgmaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3H05cY7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgi7dMQuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iI327qK5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XZmW2HN6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVx9KgBnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDuhehleWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DRCkOYc4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ub5tZNtWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 126
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WQSZdTMYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V304IDszWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FXCtwEtHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CnnYP3DfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 217
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tHlNLyWUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrcT7cU5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bDmh8LFfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bjGmowl1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBjaLbq7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hqIPo7GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jxAFGJ3TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RdJEhoISWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 54
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p3ErUTNUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zu5FynkTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0tE6jnDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func riJQeXwNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sTpBd2FnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C87xCMeUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SKPewKHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TrE2zmp5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 140
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QMmpA2QjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Zh1QMpCCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 15
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wteUszcDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PAFzIxFPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l9R6Pd9NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 66y5SuAbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MWSr95dsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FFw7BQwGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toAliozIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXk0V91EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMzRiiHUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8n8SFQ3wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RMgng4lIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GJZOWsAnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EoATJGtmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 57GRm89lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8oyHDVwSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XbFuMnadWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lEcSqbnRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0Og7uOLjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iWPmfeQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6nkUUPLfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0XYq6mgLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qOTG0VhJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KZOt2YWAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XGW83MSQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 141
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aXzcr2tGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 26
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DcDicJsSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AAUs8AM9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MUR3bRH0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1iabrEzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9AW1PNfqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yO5lqJoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LIzocd7HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZbACXfOxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 45
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pWGo08ejWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mChJc5gHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h30xbZo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ywrtvrO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uKRnJllsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HAvtfolpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func toVzmXR7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CYUjIDryWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Hvx0ft6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 92
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SGit9ILYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XJLJzOJSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ch5EzZ0RWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func orM8DpVvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dJK3zL2hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yd0x595vWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 24
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D1xfipLbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wuzxidCWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bw2e7QlWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Vd2PbuZYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 14
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SHfRLMpxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4pPn6DEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 58
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5PbheOfEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x1i3VYIXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jWbStfXXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YXtOdkpMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func u9Mr8DimWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qP0mWpJBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dHEeaZNMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJ2OmmzIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tQp8yKnTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f0FSu4tDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Mywk7q4pWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func noOrZaZjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3x2DnPanWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 152
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8uZXZ7oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EV0MWq5JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func G5gdT5dYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tZBfe3l3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QzTzPav5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ThH5bdUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ynb2SGdkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 51AH8ufkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wF819dB7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NlndNqDYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7IetOQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 128
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func P1j823ZcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func duG3MWM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oDnsRZpUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s24DQ9kwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UEH8in7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xBKD3iFoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dudIzd8kWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jFFKExa4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cttW4NntWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iysMd403Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLFmJo4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func InSyNLbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Il6AphRlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hVtceMvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JFyW6koGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NtlzeMZlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eqty8guOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nenlTDPYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5Z0qAPtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i01vSz0IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3dW5uYSmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XVAHEGjqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLw3vWpsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cPECxBE7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jAZ7TfbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKfVHTOLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 61
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func soruEk9DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mvpfby4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rxtb83odWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 94
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zekClCDrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eWaSzifwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kbCuaLnQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func avhV0tmeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ntITjI6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KwinozedWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AKeIGzGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m6CQjTHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6ygUwc4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5itS271UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xHG2pL1cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nbhr9fBdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MjpR2FjeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h3goOH6BWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nI2ggjNYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZQD6LA5fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z4HcIEFRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z2x8l2KQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8t88jz0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FSNhK6fRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jGJ0mkYJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XDmQvoJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 139
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q5tlGpSSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dwg9XwTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VULHaqNXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6R1vWxBYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 88
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TX6w6n8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2rK5NAtbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FW0SoYbwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y7cPes8zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nzjBbyVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6SHdohXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dTrfEjoJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 114
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlP9QHBeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5SQYoHHZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ioCYUYEFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LYXiKBa1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9S1i4aPmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q0ULG50XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KOjffsdHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQI7X0AQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Kg55hLrmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VYG925aHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r8lBkvIDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 230
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mc88D44aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GACm95yTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Pla5kW4KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HtH0bWZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func w7Ro7zSeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XSUltfxxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIWhNxDqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z3RXUiPGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Z8gpNd37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ruCUNC8bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UQ0QXhKoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56JSIVSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HWBmHpGbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DlRtyUNaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A6Q2GkTvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JNcJmdOjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ijbabPXgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func D14PgKGYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 174
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UFwDZg0OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y0QPisUWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 274
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

