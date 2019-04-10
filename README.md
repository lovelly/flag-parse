# flag-parse  

```
f, _ := parse.ParseArgs("-k -sm 20 --H 'Host: --baidu.com ' -H 'Accept-Encoding: gzip, deflate, sdch' -L -v -o /dev/null ")
for _, v := range f.GetStringSlice("H") {
	fmt.Println(v)
}
```
