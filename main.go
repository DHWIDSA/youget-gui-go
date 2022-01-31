package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"syscall"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var num int = 0

//初始化下载数

func main() {
	app := app.New()
	//创建一个新的app
	app.Settings().SetTheme(&myTheme{})
	//设置客制化主题解决中文问题
	window := app.NewWindow("you-get-gui")
	window.Resize(fyne.NewSize(400, 0))
	//设置窗口大小
	url := widget.NewEntry()
	//url.SetPlaceHolder("input video url")
	url1 := widget.NewLabel("| |video url:")
	//视频地址输入栏
	path := widget.NewEntry()
	//path.SetPlaceHolder("input save path")
	path1 := widget.NewLabel("| |save path:")
	form := container.New(layout.NewFormLayout(), url1, url, path1, path)
	//保存位置输入栏
	status := widget.NewLabel("| |No downloading")
	//初始化下载状态栏
	download := widget.NewButton("| |download", func() {
		_, err := os.Lstat(path.Text)
		if os.IsNotExist(err) {
			status.SetText("| |path is not exist!")
			//判断目录是否有效
		} else {
			go func() {
				num++
				num1 := num
				status.SetText(strconv.Itoa(num1) + "| |download start!")
				err := download(url.Text, path.Text)
				if err != nil {
					status.SetText(strconv.Itoa(num1) + "| |download failed!")
				}
				//错误处理
				status.SetText(strconv.Itoa(num1) + "| |download complete!")
			}()
		}
	})
	//定义下载这个作用
	content := container.NewVBox(
		form,
		download,
		status,
	)
	//创建vbox容器
	window.SetContent(content)
	window.ShowAndRun()
	//显示窗口然后运行
}

func download(url string, path string) error {
	fmt.Println("download start")
	cmd := exec.Command("you-get", url, "-o", path)
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
	err := cmd.Run()
	if err != nil {
		fmt.Println("you-get download error,you need to install you-get first!")
	} else {
		fmt.Println("download complete!")
	}
	return err
}

//使用you-get下载视频
