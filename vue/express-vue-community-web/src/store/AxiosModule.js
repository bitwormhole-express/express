
import axios from 'axios'


const storage = sessionStorage

function useToken(req) {
    const key = "x-keeper-token"
    let headers = req.headers;
    let token = storage["token"]
    headers[key] = token
}

function setToken(resp) {
    const key = "x-set-keeper-token"
    let headers = resp.headers;
    let token = headers[key]
    if (token != null) {
        storage["token"] = token;
    }
}

function getResponseError(resp) {
    let status = resp.status;
    if (status != 200) {
        return new Error("HTTP " + status)
    }
    let data = resp.data;
    if (data == null) {
        return new Error("no response data")
    }
    let status2 = data.status;
    if (status2 != 200) {
        return new Error("status:" + status)
    }
    return null;
}

axios.interceptors.request.use((config) => {
    useToken(config)
    return config
}, (error) => {
    return Promise.reject(error);
})

axios.interceptors.response.use((response) => {
    setToken(response)
    return response
}, (error) => {
    return Promise.reject(error);
})

export default {
    name: "AxiosModule",
    namespaced: true,

    state: {
    },

    mutations: {
    },

    actions: {
        execute(context, params) {
            let result = axios(params).then((resp) => {
                let err = getResponseError(resp)
                if (err != null) {
                    throw err
                }
                return resp
            })
            return result
        }
    },
}
