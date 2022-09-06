/*
	通过os.Stat方法，我们可以获取文件的信息，比如文件大小、名字等
		func main() {
			fi,err:=os.Stat(文件路径)
			if err ==nil {
				fmt.Println("name:",fi.Name())
				fmt.Println("size:",fi.Size())
				fmt.Println("is dir:",fi.IsDir())
				fmt.Println("mode::",fi.Mode())
				fmt.Println("modTime:",fi.ModTime())
			}
		}
*/

package main

func main() {

}
