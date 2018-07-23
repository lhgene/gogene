接口说明:
1. 新增一个文件: http://host:9900/f, Post请求, post的body是json格式, 例如:
   {"filetype": "fastq", "sourceFileIds":["5b52eb729e80004e781f225d"]}
2. 删除一个文件: http://host:9900/f/{id}, Delete请求, 例如:
   http://host:9900/f/5b52eb729e80004e781f225d
   直接在数据库里删除，没留备份
3. 查询文件列表: http://host:9900/f, Get请求
   查询条件在Request的form里，例如 127.0.0.1:9000/f?_id=5b54c4ec9e80002f10f817b8
   返回josn数组格式, 例如:
    [
        {
            "ID": "5b54c4ec9e80002f10f817b9",
            "filetype": "fastq",
            "filename": "",
            "filesize": 0,
            "path": "",
            "sourceFileIds": [
                "5b52eb729e80004e781f225d"
            ],
            "creationTime": "2018-07-23T01:54:52.309+08:00",
            "modificationTime": "2018-07-23T01:54:52.309+08:00",
            "taskId": "",
            "sequencingType": "",
            "probe": "",
            "sequencer": "",
            "platform": "",
            "referenceGenome": "",
            "duplicationRemoved": false,
            "bed": "",
            "commented": false
        }
    ]

    文件的字段都写在了gdata/model.go里，包括:<br />
    `json:"_id" bson:"_id"`<br />
    `json:"filetype,omitempty" bson:"filetype,omitempty"`<br/>
    `json:"filename,omitempty" bson:"filename,omitempty"`<br />
    `json:"filesize,omitempty" bson:"filesize,omitempty"`<br />
    `json:"path,omitempty" bson:"path,omitempty"`<br />
    `json:"ip,omitempty" bson:"ip,omitempty"`<br />
    `json:"sourceFileIds,omitempty" bson:"sourceFileIds,<br />
    `json:"creationTime" bson:"creationTime"`<br />
    `json:"modificationTime" bson:"modificationTime"`<br />
    `json:"taskId,omitempty" bson:"taskId,omitempty"`<br />
    `json:"sequencingType,omitempty" ncingType,omitempty"`<br />
    `json:"probe,omitempty" bson:"probe,omitempty"`<br />
    `json:"sequencer,omitempty" bson:"sequencer,<br />
    `json:"platform,omitempty" bson:"platform,omitempty"`<br />
    `json:"referenceGenome,omitempty" enceGenome,omitempty"`<br />
    `json:"duplicationRemoved,omitempty" cationRemoved,omitempty"`<br />
    `json:"bed,omitempty" bson:"bed,omitempty"`<br />
    `json:"commented,omitempty" bson:"commented,<br />