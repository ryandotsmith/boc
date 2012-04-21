```
$ export DATABASE_URL=postgres://user:pass@localhost/dbname
$ export PORT=8080
$ ./boc -q="SELECT true"
```

meanwhile...

```
$ curl http://localhost:8080/check
HTTP/1.1 200 OK
Date: Sat, 21 Apr 2012 06:26:13 GMT
Transfer-Encoding: chunked
Content-Type: text/plain; charset=utf-8

OK
```

## OR

```
$ export DATABASE_URL=postgres://user:pass@localhost/dbname
$ export PORT=8080
$ ./boc -q="SELECT false"
```

meanwhile...

```
$ curl -i http://localhost:8080/check
HTTP/1.1 417 Expectation Failed
Date: Sat, 21 Apr 2012 06:27:03 GMT
Transfer-Encoding: chunked
Content-Type: text/plain; charset=utf-8

FAIL
```
