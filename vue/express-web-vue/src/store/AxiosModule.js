
import axios from 'axios'


const storage = sessionStorage

function useToken(req) {
    const key = "x-keeper-token"
    let headers = req.headers;
    let token = storage["token"]
    headers[key] = token
}

function setToken(resp) {
    const key = "x-keeper-set-token"
    let headers = resp.headers;
    let token = headers[key]
    if (token != null) {
        storage["token"] = token;
    }
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
        execute(context, p) {
            return axios(p)
        }
    },
}
