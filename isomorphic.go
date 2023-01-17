func isIsomorphic(s string, t string) bool {
    m := make(map[byte]byte)
    inverse_m := make(map[byte]byte)
    for i := range s {
        _, ok := m[s[i]]
        _, ok2 := inverse_m[t[i]]
        if ok {
            if m[s[i]] != t[i] {
                return false
            }
        } else if ok2 {
            if inverse_m[t[i]] != s[i] {
                return false
            }
        } else {
            m[s[i]] = t[i]
            inverse_m[t[i]] = s[i]
        }
    }
    return true
}
