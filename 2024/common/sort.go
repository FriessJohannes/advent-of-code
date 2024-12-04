package common

func MergeSort(slice []int) []int {
    length := len(slice)

    if len(slice) < 2 {
        return slice
    }

    s1 := MergeSort(slice[:length/2])
    s2 := MergeSort(slice[length/2:])

    var (
        merged = []int{}
        ps1 = 0
        ps2 = 0
    )

    for  ps1 < len(s1) && ps2 < len(s2) {
        if s1[ps1] < s2[ps2] {
            merged = append(merged, s1[ps1])
            ps1++;
        } else {
            merged = append(merged, s2[ps2])
            ps2++
        }
    }

    merged = append(merged, s1[ps1:]...)
    merged = append(merged, s2[ps2:]...)

    return merged
}
