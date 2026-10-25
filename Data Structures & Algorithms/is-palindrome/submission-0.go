func isPalindrome(s string) bool {
    i, j := 0, len(s)-1
    for i < j{
        leftChar := s[i]
        if leftChar >= 'A' && leftChar <= 'Z'{
            leftChar += 32
        }

        if (leftChar < 'a' || leftChar > 'z') &&
            (leftChar < '0' || leftChar > '9'){
                i++
                continue
        }

        rightChar := s[j]
        if rightChar >= 'A' && rightChar <= 'Z'{
            rightChar += 32
        }

        if (rightChar < 'a' || rightChar > 'z') &&
            (rightChar < '0' || rightChar > '9'){
                j--
                continue
        }

        if leftChar != rightChar{
            return false
        }

        i++
        j--
    }
    return true
}