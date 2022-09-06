# run command

```bash
 ✘ nemkovich_iu@stws3344  ~/golang/gopl.io/ch1/fetchall   master  
 go run main.go http://google.com http://facebook.com http://youtube.com http://baidu.com http://yahoo.com http://amazon.com http://qq.com http://twitter.com http://taobao.com http://live.com http://google.co.in https://twitch.tv https://tut.by
```

## output

```bash
0.40s    14558  http://google.com
0.45s    14699  http://google.co.in
0.59s   116966  https://twitch.tv
0.68s    76014  http://facebook.com
0.71s    37607  http://live.com
Get "http://baidu.com": EOF
0.93s   125192  http://twitter.com
0.94s   174689  http://qq.com
1.17s   687459  http://youtube.com
1.26s     6591  http://amazon.com
1.95s   639716  http://yahoo.com
5.14s    10540  http://taobao.com
Get "https://tut.by": dial tcp 178.172.160.4:443: i/o timeout
30.01s elapsed
```
