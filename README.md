# Used to parse Gitlab files

## Usage

Each flag is dedicated for parsing a particular file.


### Parsing api_json.log files
```
./gitlab-parser.go --apijsonlog api_json.log

```

Parse for a particular field
```

Example log entry

{
	"time": "2019-07-03T00:12:08.815Z",
	"severity": "INFO",
	"duration": 42.7,
	"db": 3.56,
	"view": 39.14,
	"status": 200,
	"method": "GET",
	"path": "/api/v3/projects/someproject/repository/commits/master/blob",
	"params": {
		"filepath": "com/liaison/system/task/xmlv2/xml-split_v1_1.ns"
	},
	"host": "myhost",
	"ip": "10.123.20.69, 10.123.20.4",
	"ua": "org.gitlab.api.GitlabAPI/1.8.0_162"
}


./gitlab-parser.go --apijsonlog api_json.log --field host

mygitlabserver.com
mygitlabserver.com
mygitlabserver.com
mygitlabserver.com

```
