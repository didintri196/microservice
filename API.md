---
title: microservice v1.0.0
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
highlight_theme: darkula
headingLevel: 2

---

<!-- Generator: Widdershins v4.0.1 -->

<h1 id="microservice">microservice v1.0.0</h1>

> Scroll down for code samples, example requests and responses. Select a language for code samples from the tabs above or the mobile navigation menu.

Base URLs:

* <a href="http://localhost:3000">http://localhost:3000</a>

* <a href="http://localhost:4000">http://localhost:4000</a>

 License: 

# Authentication

- HTTP Authentication, scheme: bearer 

<h1 id="microservice-ms-auth">ms-auth</h1>

## Register

<a id="opIdRegister"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:3000/user/signup \
  -H 'Content-Type: text/plain' \
  -H 'Accept: application/json; charset=utf-8' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST http://localhost:3000/user/signup HTTP/1.1
Host: localhost:3000
Content-Type: text/plain
Accept: application/json; charset=utf-8

```

```javascript
const inputBody = '{
  "username": "didintri196921",
  "name": "Didin Tri Anggoro",
  "phone": "085895567978",
  "role": "admin"
}';
const headers = {
  'Content-Type':'text/plain',
  'Accept':'application/json; charset=utf-8',
  'Authorization':'Bearer {access-token}'
};

fetch('http://localhost:3000/user/signup',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'text/plain',
  'Accept' => 'application/json; charset=utf-8',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post 'http://localhost:3000/user/signup',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'text/plain',
  'Accept': 'application/json; charset=utf-8',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('http://localhost:3000/user/signup', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'text/plain',
    'Accept' => 'application/json; charset=utf-8',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','http://localhost:3000/user/signup', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:3000/user/signup");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"text/plain"},
        "Accept": []string{"application/json; charset=utf-8"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "http://localhost:3000/user/signup", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /user/signup`

*Register*

> Body parameter

```
username: didintri196921
name: Didin Tri Anggoro
phone: "085895567978"
role: admin

```

<h3 id="register-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|string|true|none|

> Example responses

> 201 Response

```json
{
  "message": "User created",
  "data": {
    "username": "didintri196921",
    "name": "Didin Tri Anggoro",
    "phone": "085895567978",
    "role": "admin",
    "password": "Jl78",
    "_id": "855bd7146fa44166a696ddd0c42c4251"
  }
}
```

<h3 id="register-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|None|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Register Success|[201](#schema201)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearer, bearer, bearer, bearer
</aside>

## Login

<a id="opIdLogin"></a>

> Code samples

```shell
# You can also use wget
curl -X POST http://localhost:3000/user/login \
  -H 'Content-Type: text/plain' \
  -H 'Accept: application/json; charset=utf-8' \
  -H 'Authorization: Bearer {access-token}'

```

```http
POST http://localhost:3000/user/login HTTP/1.1
Host: localhost:3000
Content-Type: text/plain
Accept: application/json; charset=utf-8

```

```javascript
const inputBody = '{
  "phone": "1",
  "password": "dF25"
}';
const headers = {
  'Content-Type':'text/plain',
  'Accept':'application/json; charset=utf-8',
  'Authorization':'Bearer {access-token}'
};

fetch('http://localhost:3000/user/login',
{
  method: 'POST',
  body: inputBody,
  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Content-Type' => 'text/plain',
  'Accept' => 'application/json; charset=utf-8',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.post 'http://localhost:3000/user/login',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Content-Type': 'text/plain',
  'Accept': 'application/json; charset=utf-8',
  'Authorization': 'Bearer {access-token}'
}

r = requests.post('http://localhost:3000/user/login', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Content-Type' => 'text/plain',
    'Accept' => 'application/json; charset=utf-8',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('POST','http://localhost:3000/user/login', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:3000/user/login");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("POST");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Content-Type": []string{"text/plain"},
        "Accept": []string{"application/json; charset=utf-8"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("POST", "http://localhost:3000/user/login", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`POST /user/login`

*Login*

> Body parameter

```
phone: "1"
password: dF25

```

<h3 id="login-parameters">Parameters</h3>

|Name|In|Type|Required|Description|
|---|---|---|---|---|
|body|body|string|true|none|

> Example responses

> 201 Response

```json
{
  "message": "Success Login",
  "data": {
    "username": "didintri196921",
    "name": "Didin Tri Anggoro",
    "phone": "085895567978",
    "role": "admin",
    "password": "Jl78",
    "_id": "855bd7146fa44166a696ddd0c42c4251"
  },
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImRpZGludHJpMTk2OTIxIiwicGhvbmUiOiIwODU4OTU1Njc5NzgiLCJyb2xlIjoiYWRtaW4iLCJ0aW1lc3RhbXAiOjE2NDQ2MDM5OTQ5MjEsImlhdCI6MTY0NDYwMzk5NCwiZXhwIjoxNjQ0NjA3NTk0fQ.GGGipd0Cqu5UTgkbrRk0TNwVnZ_KmhIM0C7QTkvveoA"
}
```

<h3 id="login-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|None|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|Login Success|[201](#schema201)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearer, bearer, bearer, bearer
</aside>

## Profile

<a id="opIdProfile"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:3000/user/profile \
  -H 'Accept: application/json; charset=utf-8' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET http://localhost:3000/user/profile HTTP/1.1
Host: localhost:3000
Accept: application/json; charset=utf-8

```

```javascript

const headers = {
  'Accept':'application/json; charset=utf-8',
  'Authorization':'Bearer {access-token}'
};

fetch('http://localhost:3000/user/profile',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json; charset=utf-8',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get 'http://localhost:3000/user/profile',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json; charset=utf-8',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('http://localhost:3000/user/profile', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json; charset=utf-8',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','http://localhost:3000/user/profile', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:3000/user/profile");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json; charset=utf-8"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://localhost:3000/user/profile", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /user/profile`

*Profile*

> Example responses

> 201 Response

```json
{
  "message": "User Authenticated",
  "data": {
    "username": "didintri196921",
    "phone": "085895567978",
    "role": "admin",
    "timestamp": 1644603994921,
    "iat": 1644603994,
    "exp": 1644607594
  }
}
```

<h3 id="profile-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|None|
|201|[Created](https://tools.ietf.org/html/rfc7231#section-6.3.2)|User Authenticated|[201](#schema201)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearer, bearer, bearer, bearer
</aside>

<h1 id="microservice-ms-fetch">ms-fetch</h1>

## Resource

<a id="opIdResource"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:3000/fetch/user/resource \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET http://localhost:3000/fetch/user/resource HTTP/1.1
Host: localhost:3000
Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('http://localhost:3000/fetch/user/resource',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get 'http://localhost:3000/fetch/user/resource',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('http://localhost:3000/fetch/user/resource', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','http://localhost:3000/fetch/user/resource', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:3000/fetch/user/resource");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://localhost:3000/fetch/user/resource", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fetch/user/resource`

*Resource*

> Example responses

> 200 Response

```json
{
  "message": "success",
  "data": [
    {
      "uuid": "2171f26a-4140-490b-a890-ea26472e2d0c",
      "komoditas": "Ikan Paus Edit 1",
      "area_provinsi": "SUMATERA BARAT",
      "area_kota": "PADANG PARIAMAN",
      "size": "30",
      "price": "500000",
      "tgl_parsed": "2022-02-11 16:02:52",
      "timestamp": "1644595372938"
    },
    {
      "uuid": "70ba57bb-5de2-4600-aa6f-ac57172e597d",
      "komoditas": "Ikan Arwanas",
      "area_provinsi": "JAWA TENGAH",
      "area_kota": "PEMALANG",
      "size": "20",
      "price": "10000",
      "tgl_parsed": "2022-01-24 12:24:29",
      "timestamp": "2022-01-24T12:24:29.951Z"
    },
    ...
  ]
}
```

<h3 id="resource-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|[200](#schema200)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearer, bearer, bearer, bearer
</aside>

## ResourcewithUSD

<a id="opIdResourcewithUSD"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:3000/fetch/user/resource/currency-usd \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET http://localhost:3000/fetch/user/resource/currency-usd HTTP/1.1
Host: localhost:3000
Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('http://localhost:3000/fetch/user/resource/currency-usd',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get 'http://localhost:3000/fetch/user/resource/currency-usd',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('http://localhost:3000/fetch/user/resource/currency-usd', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','http://localhost:3000/fetch/user/resource/currency-usd', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:3000/fetch/user/resource/currency-usd");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://localhost:3000/fetch/user/resource/currency-usd", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fetch/user/resource/currency-usd`

*Resource with USD*

> Example responses

> 200 Response

```json
{
  "message": "success",
  "data": [
    {
      "uuid": "2171f26a-4140-490b-a890-ea26472e2d0c",
      "komoditas": "Ikan Paus Edit 1",
      "area_provinsi": "SUMATERA BARAT",
      "area_kota": "PADANG PARIAMAN",
      "size": "30",
      "price": "500000",
      "price_usd": "$35.00",
      "tgl_parsed": "2022-02-11 16:02:52",
      "timestamp": "1644595372938"
    },
    {
      "uuid": "70ba57bb-5de2-4600-aa6f-ac57172e597d",
      "komoditas": "Ikan Arwanas",
      "area_provinsi": "JAWA TENGAH",
      "area_kota": "PEMALANG",
      "size": "20",
      "price": "10000",
      "price_usd": "$0.70",
      "tgl_parsed": "2022-01-24 12:24:29",
      "timestamp": "2022-01-24T12:24:29.951Z"
    },
    {
      "uuid": "76fd9d63-f093-414a-8c6e-a822d7445ea1",
      "komoditas": "Ikan Piranha",
      "area_provinsi": "LAMPUNG",
      "area_kota": "LAMPUNG TIMUR",
      "size": "76",
      "price": "150000",
      "price_usd": "$10.50",
      "tgl_parsed": "2022-01-24 12:25:36",
      "timestamp": "2022-01-24T12:25:36.401Z"
    }
    ...
  ]
}
```

<h3 id="resourcewithusd-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|[200](#schema200)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearer, bearer, bearer, bearer
</aside>

## ReportResourceWeekly

<a id="opIdReportResourceWeekly"></a>

> Code samples

```shell
# You can also use wget
curl -X GET http://localhost:3000/fetch/admin/resource/report-weekly \
  -H 'Accept: application/json' \
  -H 'Authorization: Bearer {access-token}'

```

```http
GET http://localhost:3000/fetch/admin/resource/report-weekly HTTP/1.1
Host: localhost:3000
Accept: application/json

```

```javascript

const headers = {
  'Accept':'application/json',
  'Authorization':'Bearer {access-token}'
};

fetch('http://localhost:3000/fetch/admin/resource/report-weekly',
{
  method: 'GET',

  headers: headers
})
.then(function(res) {
    return res.json();
}).then(function(body) {
    console.log(body);
});

```

```ruby
require 'rest-client'
require 'json'

headers = {
  'Accept' => 'application/json',
  'Authorization' => 'Bearer {access-token}'
}

result = RestClient.get 'http://localhost:3000/fetch/admin/resource/report-weekly',
  params: {
  }, headers: headers

p JSON.parse(result)

```

```python
import requests
headers = {
  'Accept': 'application/json',
  'Authorization': 'Bearer {access-token}'
}

r = requests.get('http://localhost:3000/fetch/admin/resource/report-weekly', headers = headers)

print(r.json())

```

```php
<?php

require 'vendor/autoload.php';

$headers = array(
    'Accept' => 'application/json',
    'Authorization' => 'Bearer {access-token}',
);

$client = new \GuzzleHttp\Client();

// Define array of request body.
$request_body = array();

try {
    $response = $client->request('GET','http://localhost:3000/fetch/admin/resource/report-weekly', array(
        'headers' => $headers,
        'json' => $request_body,
       )
    );
    print_r($response->getBody()->getContents());
 }
 catch (\GuzzleHttp\Exception\BadResponseException $e) {
    // handle exception or api errors.
    print_r($e->getMessage());
 }

 // ...

```

```java
URL obj = new URL("http://localhost:3000/fetch/admin/resource/report-weekly");
HttpURLConnection con = (HttpURLConnection) obj.openConnection();
con.setRequestMethod("GET");
int responseCode = con.getResponseCode();
BufferedReader in = new BufferedReader(
    new InputStreamReader(con.getInputStream()));
String inputLine;
StringBuffer response = new StringBuffer();
while ((inputLine = in.readLine()) != null) {
    response.append(inputLine);
}
in.close();
System.out.println(response.toString());

```

```go
package main

import (
       "bytes"
       "net/http"
)

func main() {

    headers := map[string][]string{
        "Accept": []string{"application/json"},
        "Authorization": []string{"Bearer {access-token}"},
    }

    data := bytes.NewBuffer([]byte{jsonReq})
    req, err := http.NewRequest("GET", "http://localhost:3000/fetch/admin/resource/report-weekly", data)
    req.Header = headers

    client := &http.Client{}
    resp, err := client.Do(req)
    // ...
}

```

`GET /fetch/admin/resource/report-weekly`

*Report Resource Weekly*

> Example responses

> 200 Response

```json
{
  "message": "success",
  "data": [
    {
      "uuid": "2171f26a-4140-490b-a890-ea26472e2d0c",
      "komoditas": "Ikan Paus Edit 1",
      "area_provinsi": "SUMATERA BARAT",
      "area_kota": "PADANG PARIAMAN",
      "size": "30",
      "price": "500000",
      "price_usd": "$35.00",
      "tgl_parsed": "2022-02-11 16:02:52",
      "timestamp": "1644595372938"
    },
    {
      "uuid": "70ba57bb-5de2-4600-aa6f-ac57172e597d",
      "komoditas": "Ikan Arwanas",
      "area_provinsi": "JAWA TENGAH",
      "area_kota": "PEMALANG",
      "size": "20",
      "price": "10000",
      "price_usd": "$0.70",
      "tgl_parsed": "2022-01-24 12:24:29",
      "timestamp": "2022-01-24T12:24:29.951Z"
    },
    {
      "uuid": "76fd9d63-f093-414a-8c6e-a822d7445ea1",
      "komoditas": "Ikan Piranha",
      "area_provinsi": "LAMPUNG",
      "area_kota": "LAMPUNG TIMUR",
      "size": "76",
      "price": "150000",
      "price_usd": "$10.50",
      "tgl_parsed": "2022-01-24 12:25:36",
      "timestamp": "2022-01-24T12:25:36.401Z"
    },
    ...
  ]
}
```

<h3 id="reportresourceweekly-responses">Responses</h3>

|Status|Meaning|Description|Schema|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|Success|[200](#schema200)|

<aside class="warning">
To perform this operation, you must be authenticated by means of one of the following methods:
bearer, bearer, bearer, bearer
</aside>

# Schemas

<h2 id="tocS_200">200</h2>
<!-- backwards compatibility -->
<a id="schema200"></a>
<a id="schema_200"></a>
<a id="tocS200"></a>
<a id="tocs200"></a>

```json
{
  "message": "success",
  "data": [
    {
      "uuid": "2171f26a-4140-490b-a890-ea26472e2d0c",
      "komoditas": "Ikan Paus Edit 1",
      "area_provinsi": "SUMATERA BARAT",
      "area_kota": "PADANG PARIAMAN",
      "size": "30",
      "price": "500000",
      "tgl_parsed": "2022-02-11 16:02:52",
      "timestamp": "1644595372938"
    }
  ]
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|message|string|false|none|none|
|data|[object]|false|none|none|
|» uuid|string|false|none|none|
|» komoditas|string|false|none|none|
|» area_provinsi|string|false|none|none|
|» area_kota|string|false|none|none|
|» size|string|false|none|none|
|» price|string|false|none|none|
|» tgl_parsed|string|false|none|none|
|» timestamp|string|false|none|none|

<h2 id="tocS_201">201</h2>
<!-- backwards compatibility -->
<a id="schema201"></a>
<a id="schema_201"></a>
<a id="tocS201"></a>
<a id="tocs201"></a>

```json
{
  "message": "User created",
  "data": {
    "username": "didintri196921",
    "name": "Didin Tri Anggoro",
    "phone": "085895567978",
    "role": "admin",
    "password": "Jl78",
    "_id": "855bd7146fa44166a696ddd0c42c4251"
  }
}

```

### Properties

|Name|Type|Required|Restrictions|Description|
|---|---|---|---|---|
|message|string|false|none|none|
|data|object|false|none|none|
|» username|string|false|none|none|
|» name|string|false|none|none|
|» phone|string|false|none|none|
|» role|string|false|none|none|
|» password|string|false|none|none|
|» _id|string|false|none|none|

