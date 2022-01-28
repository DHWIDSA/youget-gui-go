
# youget-gui-go

A gui written in Golang for you-get.

![image](./images/2022-01-28_13-15.png)

## Useage

```shell
./youget-gui
```

Just like the picturn above,input the video
url and click `download`,everything is easy!

## How to build

### Install dependencies

```shell
pip install you-get
```

### Build

```shell
git clone https://github.com/DHWIDSA/youget-gui-go
```

#### for linux users

linux:

```shell
go build
```

linux to windows:

```shell
env CC=x86_64-w64-mingw32-gcc CGO_ENABLED=1 GOOS=windows go build -ldflags="-H windowsgui"
```

#### for windows users

```shell
go build -ldflags="-H windowsgui"
```
