package day9

import (
    "slices"
)

type File struct {
    id int
    size int
    moved bool
}

type Partition struct {
    id int
    size int
    moveable bool
}

func ToPartitions(diskMap []int) []Partition {
    partitions := []Partition{}
    for id, size := range diskMap {
        if id % 2 == 0 {
            partitions = append(partitions, Partition{id: id/2, size: size, moveable: true})
        } else {
            partitions = append(partitions, Partition{id: -1, size: size, moveable: false})
        }
    }
    return partitions
}

func ReorderFiles(partitions []Partition) {
    for i := len(partitions)-1; i >= 0; i-- {
        file := partitions[i]
        if !file.moveable {
            continue
        }

        for index, space := range partitions {
            if space.id != -1 {
                continue
            }
            if index >= i {
                break
            }
            if space.size >= file.size {
                partitions[index].id = file.id
                partitions[index].moveable = false
                partitions[index].size = file.size

                partitions[i].id = -1
                partitions[i].moveable = false

                remainingSpace := space.size - file.size
                if remainingSpace > 0 {
                    freePartition := Partition{id: -1, size: remainingSpace, moveable: false}
                    partitions = slices.Insert(partitions, index+1, freePartition)
                    i++
                } 

                partitions = MergeFreePartitions(partitions)

                break
            }
        }
    }
}

func MergeFreePartitions(partitions []Partition) []Partition {
    length := len(partitions)
    for i := 1; i < length; {
        if partitions[i].id == -1  && partitions[i-1].id == -1 {
            partitions[i-1].size += partitions[i].size
            partitions = append(partitions[:i], partitions[i+1:]...)
            length--
            continue
        }
        i++
    }
    return partitions
}

func CalculateChecksum(partitions []Partition) int {
    var (
        checksum int = 0
        index int = 0
    )
    for _, file := range partitions {
        if file.id == -1 {
            index += file.size
            continue
        }
        checksum += SumFormula(index, index + file.size-1) * file.id
        index += file.size
    }

    return checksum
}
