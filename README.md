SHA1Check
=========
Check a certificate change for SHA1 signatures and the impending shapocolypse.

###Building 
```
https://github.com/steakejjs/SHA1Check.git
cd SHA1Check
go build sha1check.go
```

###Usage
```
./sha1check -n example.com
```
Specify a port other than 443 with -p
```
./sha1check -n example.com -p 8080
```
