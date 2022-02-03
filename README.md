# You-get GUI Golang

A GUI written in Golang for you-get.

![Image](./images/2022-01-28_17-52.png)

## Usage

```sh
./youget-gui
```

Just like the picture above, input the video
URL and click `download`,everything is easy!

## How to build

### Install dependencies

You should install Golang first, for example:

#### Debian

```sh
apt install golang-go
```

#### Arch

```sh
pacman -S golang
```

Then you must install you-get library via `pip install`.

```sh
pip install you-get
```

### Clone

```sh
git clone https://github.com/DHWIDSA/youget-gui-go
```

### Build

#### For Linux and MacOS users

Linux and MacOS:

Remove as follows in `main.go` download function:

```Golang
	if runtime.GOOS == "windows" {
		cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	}
```

```sh
go build
```

Others to Windows:

```sh
env CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows go build -ldflags="-H windowsgui"
```

#### For Windows users

```sh
go build -ldflags="-H windowsgui"
```

Thanks [laishulu](https://github.com/laishulu) who offers [Sarasa-Mono-SC-Nerd](https://github.com/laishulu/Sarasa-Mono-SC-Nerd) font.
