
// // 判断登录是否成功
// function isLoginOK(data) {
//     if (data == null) {
//         return false;
//     }
//     let status = data.status
//     let auth = data.auth
//     if (auth == null || status == null) {
//         return false
//     }
//     return (auth.success && (status == 200))
// }

export default {
    name: "BucketModule",
    namespaced: true,

    state: {
        bucketlist: []
    },

    getters: {
        items(state) {
            return state.bucketlist
        }
    },

    mutations: {
        updateItems(state, p) {
            state.bucketlist = p;
        }
    },

    actions: {
        fetch(context) {
            let p = {
                method: "GET",
                url: "/api/v1/buckets/"
            }
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                context.commit('updateItems', res.data.session)
                return res;
            })
        },
    },
}
