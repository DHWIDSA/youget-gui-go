
# You-get GUI Golang

A GUI written in Golang for you-get.

![Image](./images/2022-01-28_17-52.png)

## Usage

```shell
./youget-gui
```

Just like the picture above, input the video
URL and click `download`,everything is easy!

## How to build

### Install dependencies

you should install Golang first, for example:

#### Debian

```shell
apt install golang-go
```

#### Arch

```shell
pacman -S golang
```

Then you must install you-get library via `pip install`.

```shell
pip install you-get
```

### Clone

```shell
git clone https://github.com/DHWIDSA/youget-gui-go
```

### Build

#### For Linux and Darwin users
Linux and Darwin:

Remove as follows in `main.go` download function:

```Golang
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
```

```shell
go build
```

Others to Windows:

```shell
env CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows go build -ldflags="-H windowsgui"
```

#### for Windows users

```shell
go build -ldflags="-H windowsgui"
```

Thanks [laishulu](https://github.com/) who offers [Sarasa-Mono-SC-Nerd](https://github.com/laishulu/Sarasa-Mono-SC-Nerd) font.

