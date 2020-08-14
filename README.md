# go-common
### logs
Simple and fast, leveled log library.  
Usage 
```
logs.Info("info")
log := logs.New("MyLog")
log.SetLevel("trace")
log.Debug("debug")

output message like:
    common-logs: 2020-08-08T22:27:25+08:00 [info] logs.go@119: info
    MyLog: 2020-08-08T22:25:18+08:00 [debug] main.go@12: debug
```
-------
### http
Lightweight and high performance based on net/http. http router library.  
Usage
```
type HelloController struct {
}

func (h *HelloController) Path() string{
	return "/"
}

func (h *HelloController) Execute(ctx router.Context) {
	user := ctx.Get("user")
	ctx.Put(user)
}

main.go
func main() {
	router.Registry(&HelloController{})

	router.Run(":8080")
}

```
### config
simple conf file parser  : file format.
* key = value in a line.  
* namespace use . split in key.   
  db.mysql.ip = 127.0.0.1  
* group by , like db.mysql.user = admin,abc  
template
```
 db.mysql.ip = 127.0.0.1
 db.mysql.port = 3306
```
code example
```
sc := config.New("test.conf")
ip, ok := sc.Get("db.mysql.ip")
if ok {
    fmt.Println(ip)
}
```
