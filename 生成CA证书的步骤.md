#### 根证书

它包含的文件如下:
- 公钥
- 私钥

##### 生成Key

```shell
openssl genrsa -out ca.key 2048
```

##### 生成秘钥

```shell
openssl req -new -x509 -days 7200 -key ca.key -out ca.pem
```

##### 填写信息

```shell
Country Name (2 letter code) []: CN
State or Province Name (full name) []: guangdong
Locality Name (eg, city) []: shenzhen
Organization Name (eg, company) []: shenzhen
Organizational Unit Name (eg, section) []: shenzhen
Common Name (eg, fully qualified host name) []: algorithm
Email Address []: wangcc.peaut@gmail.com
```

#### Server

##### 生成key

```shell
openssl ecparam -genkey -name secp384r1 -out server.key
```

##### 生成CSR 

```shell
openssl req -new -key server.key -out server.csr
```

##### 填写信息

```shell
Country Name (2 letter code) []: CN
State or Province Name (full name) []: guangdong
Locality Name (eg, city) []: shenzhen
Organization Name (eg, company) []: shenzhen
Organizational Unit Name (eg, section) []: shenzhen
Common Name (eg, fully qualified host name) []: algorithm
Email Address []: wangcc.peaut@gmail.com

Please enter the following 'extra' attributes
to be sent with your certificate request
A challenge password []: anewtime
An optional company name []: shenzhen

```

##### 基于 CA 签发

```shell
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -days 3650 -in server.csr -out server.pem
```

#### Client

##### 生成Key

```shell
openssl ecparam -genkey -name secp384r1 -out client.key
```

##### 生成CSR

```shell
openssl req -new -key client.key -out client.csr
```

##### 基于 CA 签发

```shell
openssl x509 -req -sha256 -CA ca.pem -CAkey ca.key -CAcreateserial -day 3650 -in client.csr -out client.pem
```

#### 整理目录

```shell
$ tree conf 
conf
├── ca.key
├── ca.pem
├── ca.srl
├── client
│   ├── client.csr
│   ├── client.key
│   └── client.pem
└── server
    ├── server.csr
    ├── server.key
    └── server.pem

```
