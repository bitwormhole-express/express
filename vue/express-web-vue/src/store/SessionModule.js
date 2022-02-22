
// 判断登录是否成功
function isLoginOK(data) {
    // TODO ... 
    if (data == null) {
        return false;
    }
    return false
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
            context.dispatch('axios/execute', p, { root: true }).then((res) => {
                context.commit('updateSessionInfo', res.data.session)
            })
        },

        login(context, p) {
            p.method = 'POST'
            p.url = '/api/v1/auth/'
            context.dispatch('axios/execute', p, { root: true }).then((res) => {
                if (isLoginOK(res.data)) {
                    context.dispatch('fetch')
                }
            })
        },

        logout(context) {
            let p = {
                method: "DELETE",
                url: "/api/v1/session/",
            }
            context.dispatch('axios/execute', p, { root: true }).then(() => {
                context.dispatch('fetch')
            })
        }
    },
}
