
function aliyunOSS() {

    const c = true; // c is 'const'
    const v = true; // v is 'visible'

    let p = [
        { name: "bucket", label: "桶", value: "b1", description: "", c },
        { name: "bucket.b1.name", label: "名称", value: "", description: "", v },
        { name: "bucket.b1.driver", label: "驱动", value: "aliyun", description: "", c },
        { name: "bucket.b1.dn-bucket", label: "桶域名", value: "", description: "" },
        { name: "bucket.b1.dn-endpoint", label: "端点域名", value: "", description: "" },
        { name: "bucket.b1.credential", label: "凭证", value: "c1", description: "", c, v },
        { name: "credential.c1.access-key-id", label: "密钥ID", value: "", description: "", v },
        { name: "credential.c1.access-key-secret", label: "密钥密文", value: "", description: "", v },
    ]


    return {
        driver: "aliyun",
        label: "阿里云 OSS",
        properties: p
    }
}

function baiduBOS() {

    const c = true; // c is 'const'
    const v = true; // v is 'visible'

    let p = [
        { name: "bucket", label: "桶", value: "b1", description: "", c },
        { name: "bucket.b1.name", label: "名称", value: "", description: "", v },
        { name: "bucket.b1.driver", label: "驱动", value: "baidu", description: "", c },
        { name: "bucket.b1.dn-bucket", label: "桶域名", value: "", description: "" },
        { name: "bucket.b1.dn-endpoint", label: "端点域名", value: "", description: "" },
        { name: "bucket.b1.credential", label: "凭证", value: "c1", description: "", c, v },
        { name: "credential.c1.access-key-id", label: "密钥ID", value: "", description: "", v },
        { name: "credential.c1.access-key-secret", label: "密钥密文", value: "", description: "", v },
    ]

    return {
        driver: "baidu",
        label: "百度云 BOS",
        properties: p
    }
}

function huaweiOBS() {

    // "bucket": "main",
    // "bucket.main.name": "xxx",
    // "bucket.main.driver": "huawei",
    // "bucket.main.dn-bucket":    "xxx",
    // "bucket.main.dn-endpoint": "xxx",
    // "bucket.main.credential": "c1",
    // "credential.c1.access-key-secret":     "xxx",
    // "credential.c1.access-key-id": "xxx",

    const c = true; // c is 'const'
    const v = true; // v is 'visible'

    let p = [
        { name: "bucket", label: "桶", value: "b1", description: "", c },
        { name: "bucket.b1.name", label: "名称", value: "", description: "", v },
        { name: "bucket.b1.driver", label: "驱动", value: "huawei", description: "", c },
        { name: "bucket.b1.dn-bucket", label: "桶域名", value: "", description: "" },
        { name: "bucket.b1.dn-endpoint", label: "端点域名", value: "", description: "" },
        { name: "bucket.b1.credential", label: "凭证", value: "c1", description: "", c, v },
        { name: "credential.c1.access-key-id", label: "密钥ID", value: "", description: "", v },
        { name: "credential.c1.access-key-secret", label: "密钥密文", value: "", description: "", v },
    ]

    return {
        driver: "huawei",
        label: "华为云 OBS",
        properties: p
    }
}

function qqCOS() {

    const c = true; // c is 'const'
    const v = true; // v is 'visible'

    let p = [
        { name: "bucket", label: "桶", value: "b1", description: "", c },
        { name: "bucket.b1.name", label: "名称", value: "", description: "", v },
        { name: "bucket.b1.driver", label: "驱动", value: "qq", description: "", c },
        { name: "bucket.b1.dn-bucket", label: "桶域名", value: "", description: "" },
        { name: "bucket.b1.dn-endpoint", label: "端点域名", value: "", description: "" },
        { name: "bucket.b1.credential", label: "凭证", value: "c1", description: "", c, v },

        { name: "bucket.b1.bucket-url", label: "桶 URL", value: "", description: "", v },
        { name: "bucket.b1.service-url", label: "服务 URL", value: "", description: "", v },

        { name: "credential.c1.access-key-id", label: "密钥ID", value: "", description: "", v },
        { name: "credential.c1.access-key-secret", label: "密钥密文", value: "", description: "", v },
    ]

    return {
        driver: "qq",
        label: "腾讯云 COS",
        properties: p
    }
}



function makeCopy(src) {
    if (src == null) {
        return src;
    }
    let str = JSON.stringify(src)
    return JSON.parse(str)
}

///////////////////////////////////////////

function getDriver(name) {
    let src = listDrivers()
    let dst = null;
    for (var i in src) {
        let x = src[i]
        if (x.driver == name) {
            dst = x;
            break;
        }
    }
    return makeCopy(dst)
}

function listDrivers() {
    const list = []
    list.push(aliyunOSS())
    list.push(baiduBOS())
    list.push(huaweiOBS())
    list.push(qqCOS())
    return makeCopy(list)
}


function propertiesListToTable(list) {
    let table = {}
    for (var i in list) {
        let item = list[i]
        table[item.name] = item.value
    }
    return table;
}

function propertiesTableToList(table, defaultDriverID) {
    let list = []
    if (table == null) {
        table = {}
    }
    let bucketID = table["bucket"]
    let driverID = table["bucket." + bucketID + ".driver"]
    let driver = getDriver(driverID)
    if (driver == null) {
        driver = getDriver(defaultDriverID)
    }
    if (driver == null) {
        driver = getDriver("aliyun")
    }
    if (driver.properties == null) {
        return list
    }
    list = driver.properties
    for (var i in list) {
        let item = list[i]
        let k = item.name;
        let v = table[k]
        if (v == null) {
            continue
        }
        item.value = v
    }
    return list
}

export default {
    name: "bucket-drivers",

    getDriver,
    listDrivers,
    propertiesListToTable,
    propertiesTableToList,
}
