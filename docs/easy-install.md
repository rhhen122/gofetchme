# Easy Install

> [!NOTE]
> This will open a `nano session` please add the following to the bottom of the file:
> `alias gofetch='$HOME/gofetchme/./gofetch'`

### for a `bash` system
```
git clone https://github.com/rhhen122/gofetchme.git
cd gofetchme
go build gofetch.go
./gofetch
cd ~
nano .bashrc
```
### for a `zsh` system
```
git clone https://github.com/rhhen122/gofetchme.git
cd gofetchme
go build gofetch.go
./gofetch
cd ~
nano .zshrc
```
Now Learn how to Config it:
<a href="../docs/learn-to-config.md">Config Tutorial</a>