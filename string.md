# STRING

To figure out 1 string exist in other string, it's depends on the problem. Therefore, it's generally contains 2 or more approaches

## Use 2 pointers
- Example:
`func isSequence(s, sub string) bool {
    i, j := 0, 0
    for i < len(s) && j < len(sub) {
        if s[i] == sub[j] {
            j++
        }
        i++
    }
    return j == len(s)
}`

## Use hash Map
Use hash map to store each character, 2 kind of types are byte and rune