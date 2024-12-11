package day9

func SumFormula(from int, to int) int {
    return (to * (to + 1) - from * (from - 1)) / 2
}


func Compress(diskMap []int) int {
    var (
        checksum  int
        index     int = 0
        rightPtr  int = len(diskMap) - 1
    )

    for i := 0; i < len(diskMap); i++ {
        if i % 2 != 0 {
            free := diskMap[i]
            for free > 0 && rightPtr > i {
                if diskMap[rightPtr] == 0 {
                    rightPtr -= 2
                    continue
                }
                if free >= diskMap[rightPtr] {
                    checksum += SumFormula(index, index+diskMap[rightPtr]-1) * (rightPtr / 2)
                    free -= diskMap[rightPtr]
                    index += diskMap[rightPtr]
                    diskMap[rightPtr] = 0
                    rightPtr -= 2
                } else {
                    checksum += SumFormula(index, index+free-1) * (rightPtr / 2)
                    diskMap[rightPtr] -= free
                    index += free
                    free = 0
                }
            }
        } else {
            checksum += SumFormula(index, index+diskMap[i]-1) * (i / 2)
            index += diskMap[i]
        }
    }
    return checksum
}
