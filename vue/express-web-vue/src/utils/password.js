import base64js from "base64-js";

function encodePassword(password) {
    let src = password + "";
    let dst = [];
    let len = src.length;
    for (let i = 0; i < len; i++) {
        let code = src.charCodeAt(i);
        dst.push(code);
    }
    return base64js.fromByteArray(dst);
}

export default {
    name: "PasswordUtil",
    encodePassword
}