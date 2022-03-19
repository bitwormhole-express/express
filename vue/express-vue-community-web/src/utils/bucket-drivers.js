
function aliyunOSS() {
    let p = {}
    return {
        properties: p
    }
}

function baiduBOS() {
    let p = {}
    return {
        properties: p
    }
}

function huaweiOBS() {
    let p = {}
    return {
        properties: p
    }
}

function qqCOS() {
    let p = {}
    return {
        properties: p
    }
}


export default {
    name: "bucket-drivers",

    listDrivers() {
        const list = []
        list.push(aliyunOSS())
        list.push(baiduBOS())
        list.push(huaweiOBS())
        list.push(qqCOS())
        return list
    },
}
