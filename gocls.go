packet gocls

import (
    "strings"
    //    "fmt"
)

func GetLCS(str1 string, str2 string) int {
    rune1 := []rune(str1)
    maxlen := 0
    //  maxstr := ""
    rlen1 := len(rune1)
    for k,_ := range rune1 {
        tmpstr := ""
        var i,j,tmplen int
        i = k
        if k+maxlen >= rlen1 {
            break
        }else{
            tmpstr = string(rune1[k])
            for j=k+1; j < k+maxlen ; j++ {
                tmpstr += string(rune1[j])
                i++
            }
        }
        if maxlen > 0{
            tmplen = maxlen -1
        }else{
            tmplen = 0
        }

        for{
            re := strings.Contains(str2, tmpstr)
            if re {
                tmplen ++
                if tmplen > maxlen {
                    maxlen = tmplen
                    //maxstr = tmpstr
                }
                i++
                if(i >= rlen1) {
                    break
                }
                tmpstr += string(rune1[i])
            }else{
                break
            }
        }
    }
    //  fmt.Printf("\n%s",maxstr)
    return maxlen
}
