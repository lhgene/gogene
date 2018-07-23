接口说明:
1. 新增一个文件: http://host:9000/f, Post请求, post的body是json格式, 例如:
   {"filetype": "fastq", "sourceFileIds":["5b52eb729e80004e781f225d"]}
2. 删除一个文件: http://host:9000/f/{id}, Delete请求, 例如:
   http://host:9000/f/5b52eb729e80004e781f225d
   直接在数据库里删除，没留备份
3. 查询文件列表: http://host:9000/f, Get请求
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