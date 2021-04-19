/**
 * @author:  ccforevercn<1253705861@qq.com>
 * @link     http://ccforever.cn
 * @license  https://github.com/ccforevercn
 * @date:    2021/3/17
 */
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"strings"
)

var (
	// 指定扫描的文件格式
	ContentFileExt = map[string]bool{}
	// 扫描的关键字
	keywords string
	// 扫描的根目录
	address string
	// 写入扫描结果的文件名称
	writeFile string
	// 写入扫描结果的文件路径
	writeFilePath string
	// 当前文件目录
	currentPath string
	// 系统路径分隔符
	pathSeparator string
)

func init()  {
	// 设置当前文件目录
	currentPath, _ = os.Getwd()
	// 设置写入扫描结果的文件名称
	writeFile = "lookup.txt"
	// 设置系统路径分隔符
	pathSeparator = string(os.PathSeparator)
	// 设置写入扫描结果的文件路径
	writeFilePath = currentPath + pathSeparator + writeFile
}

func main()  {
	// 扫描指定文件
	receive()
	fmt.Println("扫描完成，输入Enter即可退出程序")
	fmt.Scanln()
}

func receive()  {
	// 请输入扫描的文件格式
	var fileExt string
	for  {
		fmt.Printf("请输入扫描的文件格式,类似.go,输入exit即可退出输入文件格式：")
		fmt.Scanln(&fileExt)
		fileExtLen := len(fileExt) > 0
		if fileExtLen {
			fileExtStatus := fileExt == "exit"
			if fileExtStatus {
				break
			}
			ContentFileExt[fileExt] = true
		}
	}
	// 获取扫描文件格式的总数
	ContentFileExtLen := len(ContentFileExt) > 0
	// 如果为空直接退出程序
	if !ContentFileExtLen {
		return
	}
	for  {
		fmt.Printf("请输入查找的关键字：")
		fmt.Scanln(&keywords)
		keywordsLen := len(keywords) > 0
		if keywordsLen {
			break
		}
	}
	// 获取扫描的根目录
	fmt.Printf("请输入扫描的根目录，不输入为查看当前工作目录：")
	fmt.Scanln(&address)
	// 验证扫描的目录是否为空，如果为空则重置扫描的目录为当前路径
	addressLen := len(address) > 0
	if !addressLen {
		address = currentPath
	}
	// 打开扫描的根目录
	fmt.Println("拼命扫描中")
	dir(address)
}

// 扫描目录
func dir(address string)  {
	// 打开扫描的目录
	fileInfoList, _ := ioutil.ReadDir(address)
	for i := range fileInfoList {
		// 打开的文件或者文件夹追加系统路径分隔符
		filePath :=  address + pathSeparator
		// 设置文件夹或者文件名称绝对路径
		fileName := filePath + fileInfoList[i].Name()
		if fileInfoList[i].IsDir() {
			// 文件夹重新调用打开目录函数
			dir(fileName)
		} else {
			// 文件内容扫描
			file(fileName)
		}
	}
}

// 验证文件扩展名及文件内容扫描并添加到扫描结果文件
func file(fileName string)  {
	// 文件扩展名
	fileExt := path.Ext(fileName)
	// 验证扩展名是否为指定扫描的文件格式
	fileExtBool, _ := ContentFileExt[fileExt]
	if fileExtBool {
		// 打开文件
		content, _ := ioutil.ReadFile(fileName)
		// 扫描文件内容中是否存在关键字
		check := strings.Contains(string(content), keywords)
		if check {
			// 检测写入扫描结果的文件是否存在
			_, err := os.Stat(writeFilePath)
			if err != nil {
				// 不存在则创建文件
				file, createErr := os.Create(writeFilePath)
				if createErr != nil {
					// 创建文件失败，打印失败原因并终止程序
					panic(createErr)
				}
				// 写入文件
				file.Write([]byte(fileName + "\r\n"))
				// 关闭文件
				file.Close()
			}else {
				// 打开写入扫描结果的文件
				file, _ := os.OpenFile(writeFilePath, os.O_APPEND, 0666)
				// 写入文件
				file.Write([]byte(fileName + "\r\n"))
				// 关闭文件
				file.Close()
			}
			fmt.Println("搜索结果写入的文件名称：" + writeFilePath)
		}
	}
}

