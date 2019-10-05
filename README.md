# go_crud
simple app like postman

### Installation

Run the following command to install the package:

```
git clone github.com/syronz/
go install
```

create directory and inside that put these files and directories:

cofnig.json
```json
{
	"data_dir":"data",
	"base_path":"basic.json"
}
```

basic.json
```json
{
	"env": [
		{
			"local": {
				"_TOKEN1_":"Bearer 234235078971",
				"_TOKEN2_":"eyJhbGciOiJIUzI1...",
				"_URL_":"http://127.0.0.1:8080",
				"_NAME_":"Sample"
			}
		},
		{
			"global": {
				"_URL_":"http://192.168.1.1:8080",
				"_TOKEN2_":"Bearer xxxxxxxxxxxxxxxxxxxx"
			}
		}
	],

	"header":{
		"Content-Type": "application/json",
		"User-Agent": "go_crud/0.1",
		"Accept": "*/*",
		"Cache-Control": "no-cache",
		"Connection": "keep-alive"
	}
}
```

inside the direcotry should be like:

	├ basic.json
	├ config.json
	└ requests
		└ funds
			├ get.json
			├ getOne.json
			└ post.json

### Run
```shell
go_prod -config cofnig.json
```

## request fiels content

```json
{
	"method":"post",
	"url":"_URL_/m/funds",
	"authorization":"Bearer _TOKEN2_",
	"payload": {
		"type": "in",
		"invoice_id": 1,
		"peer_id": null,
		"amount": 3000
	}
}
```
