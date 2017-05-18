GoLang:


1、安装Intellij idea 的golang插件

2、建立目录：D:\SystemFile\GoWorkspace

在系统里面配置GOPATH=D:\SystemFile\GoWorkspace，然后在GOPATH目录下面建立 src, bin, pkg三个目录，然后把工程建立在src目录下面，例如golang-start

D:\SystemFile\GoWorkspace\bin
D:\SystemFile\GoWorkspace\pkg
D:\SystemFile\GoWorkspace\src\golang-start


工程结构：

golang-start ---com

                                  -----start

                                        -----main.Go

                                 ------im

                                        -----im.go

main.go

package main

import "fmt"

import "golang-start/im"

func main() {
       fmt.Println(111)
       fmt.Println("df")
       im.Han()

}


im.go


package im

import "fmt"

func Han() {
       fmt.Println("i am han")
       han1()
}

func han1() {
       fmt.Println("i am han11")
}


其他都不用管，直接能在start目录下面的main.go里面引用im目录下面的im.go里面的方法。

























#设置Golang的GOPATH
http://lib.csdn.net/article/go/34082

#golang工程约定，分包和编译
http://blog.csdn.net/king_sky/article/details/12065105





