#### 该cloudgo功能
实现了类似课程上给出的示例的功能，即访问hello/id后返回hello id的json
#### 框架
使用了martini框架
原因：该框架路由方式等参考了express，较为熟悉容易上手
#### 使用curl测试
使用命令
```
curl -v http://localhost:8888/hello/your
```
测试结果
```
caijh@caijh-PC:~/goWork/src/github.com/caijh23/goWeb/web/cloudgo$ curl -v http://localhost:8888/hello/your
*   Trying ::1...
* Connected to localhost (::1) port 8888 (#0)
> GET /hello/your HTTP/1.1
> Host: localhost:8888
> User-Agent: curl/7.47.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: application/json; charset=UTF-8
< Date: Sat, 11 Nov 2017 13:48:23 GMT
< Content-Length: 26
< 
{
  "Test": "Hello your"
* Connection #0 to host localhost left intact
}
```
#### 使用ab测试
测试命令以及测试结果
```
caijh@caijh-PC:~$ ab -n 1000 -c 100 http://localhost:8888/hello/your
This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
Licensed to The Apache Software Foundation, http://www.apache.org/

Benchmarking localhost (be patient)
Completed 100 requests
Completed 200 requests
Completed 300 requests
Completed 400 requests
Completed 500 requests
Completed 600 requests
Completed 700 requests
Completed 800 requests
Completed 900 requests
Completed 1000 requests
Finished 1000 requests


Server Software:        
Server Hostname:        localhost
Server Port:            8888

Document Path:          /hello/your
Document Length:        26 bytes

Concurrency Level:      100
Time taken for tests:   0.530 seconds
Complete requests:      1000
Failed requests:        0
Total transferred:      149000 bytes
HTML transferred:       26000 bytes
Requests per second:    1886.53 [#/sec] (mean)
Time per request:       53.007 [ms] (mean)
Time per request:       0.530 [ms] (mean, across all concurrent requests)
Transfer rate:          274.51 [Kbytes/sec] received

Connection Times (ms)
              min  mean[+/-sd] median   max
Connect:        0    2   2.4      1      15
Processing:     1   49  21.8     46     127
Waiting:        1   48  22.4     45     126
Total:          1   51  22.3     48     129

Percentage of the requests served within a certain time (ms)
  50%     48
  66%     59
  75%     66
  80%     70
  90%     78
  95%     90
  98%    105
  99%    108
 100%    129 (longest request)
```
关于ab的参数</br>
-n 测试的总请求数。默认时，仅执行一个请求</br>
-c 一次并发请求个数。默认是一次一个。</br>
-H 添加请求头，例如 ‘Accept-Encoding: gzip’，以gzip方式请求。</br>
-t 测试所进行的最大秒数。其内部隐含值是-n 50000。它可以使对服务器的测试限制在一个固定的总时间以内。默认时，没有时间限制。</br>
-p 包含了需要POST的数据的文件.</br>
-T POST数据所使用的Content-type头信息。</br>
-v 设置显示信息的详细程度 – 4或更大值会显示头信息， 3或更大值可以显示响应代码(404, 200等), 2或更大值可以显示警告和其他信息。 -V 显示版本号并退出。</br>
-w 以HTML表的格式输出结果。默认时，它是白色背景的两列宽度的一张表。</br>
-i 执行HEAD请求，而不是GET。</br>
-C -C cookie-name=value 对请求附加一个Cookie:行。 其典型形式是name=value的一个参数对。此参数可以重复。</br>
#### ab测试结果重要参数的解释
对于测试结果各个参数其实按照字面意义理解就可以，其中几个重要的参数分别是</br>
```
Requests per second:    1886.53 [#/sec]
```
表示当前测试的服务器每秒可以处理1886.53个静态html的请求事务，后面的mean表示平均。这个数值表示当前机器的整体性能，值越大越好。</br>
计算公式：Request per second=Complete requests/Time taken for tests
```
Time per request:       53.007 [ms] (mean)
```
单个并发的延迟时间，后面的mean表示平均。用户平均请求等待时间。</br>
计算公式：Time per request=Time taken for tests/（Complete requests/Concurrency Level）
```
Time per request:       0.530 [ms] (mean, across all concurrent requests)
```
隔离开当前并发，单独完成一个请求需要的平均时间。服务器平均请求等待时间。</br>
计算公式：Time taken for/testsComplete requests</br>
同时，它也等于用户平均请求等待时间/并发用户数，即
Time per request/Concurrency Level。
#### 关于拓展作业
我读了Negroni库的源代码，具体见我博客http://blog.csdn.net/caijhBlog/article/details/78507334