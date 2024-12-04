package day3

func SetInstructions(memory *string, index *int, state bool) bool{
    if *index > 4 && (*memory)[*index-5:*index] == "don't" {
        return false
    }
    if *index > 2 && (*memory)[*index-2:*index] == "do" {
        return true
    }
    return state
}
