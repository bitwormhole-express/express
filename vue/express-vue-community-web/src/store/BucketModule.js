
import BucketDrivers from './../utils/bucket-drivers.js'

export default {
    name: "BucketModule",
    namespaced: true,

    state: {
        bucketlist: [],
        driverlist: [],
    },

    getters: {
        drivers(state) {
            return state.driverlist
        },
        items(state) {
            return state.bucketlist
        },
    },

    mutations: {
        updateDrivers(state, p) {
            state.driverlist = p;
        },
        updateItems(state, p) {
            state.bucketlist = p;
        }
    },

    actions: {

        delete(context, item) {
            let id = item.id;
            let p = {
                method: "DELETE",
                url: "/api/v1/buckets/" + id,
            }
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                context.dispatch("fetch")
                return res;
            })
        },

        fetch(context) {
            let p = {
                method: "GET",
                url: "/api/v1/buckets/"
            }
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                context.commit('updateItems', res.data.buckets)
                return res;
            })
        },

        insert(context, bucket) {
            let p = {
                method: "POST",
                url: "/api/v1/buckets/",
                data: {
                    buckets: [bucket]
                }
            }
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                context.dispatch("fetch")
                return res;
            })
        },

        loadDrivers(context) {
            let list = BucketDrivers.listDrivers()
            context.commit('updateDrivers', list)
        },
    },
}
