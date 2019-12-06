# intro
simple file upload and download tool written in go
# install
```bash
go get -u github.com/IanSmith123/fileupload
```
# usage
start this program
````
fileupload 
````
set port 
```bash
fileupload -port 2019 
```
help doc
```bash
fileupload -h
```

# upload
```bash
curl -F file=@filename.png 127.0.0.1:2019
```

# download
```bash
wget 127.0.0.1:2019/file
```

# notice
Everyone can upload file to your server and download file in current directory once you start this program.

So, please close it in time.