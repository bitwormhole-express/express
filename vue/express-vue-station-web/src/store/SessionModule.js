
// 判断登录是否成功
function isLoginOK(data) {
    if (data == null) {
        return false;
    }
    let status = data.status
    let auth = data.auth
    if (auth == null || status == null) {
        return false
    }
    return (auth.success && (status == 200))
}

export default {
    name: "SessionModule",
    namespaced: true,

    state: {
        session: {}
    },

    getters: {
        info(state) {
            return state.session
        }
    },

    mutations: {
        updateSessionInfo(state, p) {
            // p.nickname = 'demo'
            // p.authenticated = true;
            state.session = p;
        }
    },

    actions: {
        fetch(context) {
            let p = {
                method: "GET",
                url: "/api/v1/session/"
            }
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                context.commit('updateSessionInfo', res.data.session)
                return res;
            })
        },

        login(context, p) {
            p.method = 'POST'
            p.url = '/api/v1/auth/'
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                if (isLoginOK(res.data)) {
                    context.dispatch('fetch')
                } else {
                    throw new Error("登录失败!")
                }
                return res
            })
        },

        logout(context) {
            let p = {
                method: "POST",
                url: "/api/v1/logout",
            }
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                context.dispatch('fetch')
                return res;
            })
        }
    },
}
