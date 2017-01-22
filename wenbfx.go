package main

import (
        "fmt"
       // "strconv"
        "unicode"
        "net/http"
        "github.com/wangbin/jiebago"
)

var seg jiebago.Segmenter

func init() {
        seg.LoadDictionary("dict.txt")
}

func print(ch <-chan string) {
        for word := range ch {
                fmt.Printf(" %s /", word)
        }
        fmt.Println()
}
func IsChineseChar(str string) bool {
        for _, r := range str {
        if unicode.Is(unicode.Scripts["Han"], r) {
                        return true
                }
        }
        return false
}
func Getwords(ch <-chan string)  string {
        words :=""
        for word := range ch {
               if IsChineseChar(word) == false ||len([]rune(word)) == 1 {
                       continue
               }
               words=words+" "+ word 
        }
        return words
}

func Example() {
        fmt.Print("【全模式】：")
        print(seg.CutAll("我来到北京清华大学"))

        fmt.Print("【精确模式】：")
        print(seg.Cut("我来到北京清华大学", false))

        fmt.Print("【新词识别】：")
        print(seg.Cut("他来到了网易杭研大厦", true))

        fmt.Print("【搜索引擎模式】：")
        print(seg.CutForSearch("小明硕士毕业于中国科学院计算所，后在日本京都大学深造", true))
}
func mainpage(w http.ResponseWriter, r *http.Request) {
       r.ParseForm()
       contents := r.FormValue("content")
       b :=contents
       b=b+"\n【模式1】：\n"
       b=b+ Getwords(seg.CutAll(contents))
       b=b+"\n【模式2】：\n"+ Getwords(seg.Cut(contents,false))
       b=b+"\n【模式3】：\n"
       b=b+ Getwords(seg.Cut(contents,true))
       b=b+"\n【模式4】：\n"
       b=b+ Getwords(seg.CutForSearch(contents,true))

       w.Write([]byte(b))
}
func main(){
       Example()

       http.HandleFunc("/mainpage", mainpage)
       http.Handle("/", http.StripPrefix("/", http.FileServer(http.Dir("./"))))
       err := http.ListenAndServe(":28080", nil)
       if err != nil {
              fmt.Println("ListenAndServer:" , err)

	}
}
