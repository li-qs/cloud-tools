package main

import "fmt"

func HelpInfo() {
	fmt.Println("help	获取帮助信息：help")
	fmt.Println("apply	申请免费证书：apply [www.example.com]")
	fmt.Println("list	获取证书列表：list 或 list [page_number]")
	fmt.Println("detail	获取证书信息：detail [ID]")
	fmt.Println("keys	获取证书密钥：keys [ID]")
	fmt.Println("download	下载证书：download [ID] [/target_dir/target_file.zip]")
	fmt.Println("url	获取证书下载链接：url [ID]")
	fmt.Println("revoke	吊销证书：revoke [ID]")
	fmt.Println("delete	删除证书：delete [ID]")
}
