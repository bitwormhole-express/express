
import BucketDrivers from './../utils/bucket-drivers.js'


function makeDriverTable() {
    let dst = {};
    let src = BucketDrivers.listDrivers()
    for (var i in src) {
        let item = src[i]
        dst[item.driver] = item
    }
    return dst
}

const theDriverTable = makeDriverTable();

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

        driverTable() {
            return theDriverTable
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

        update(context, bucket) {
            let id = bucket.id;
            let p = {
                method: "PUT",
                url: "/api/v1/buckets/" + id,
                data: {
                    buckets: [bucket]
                }
            }
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                context.dispatch("fetch")
                return res;
            })
        },

        test(context, bucket) {
            let p = {
                method: "POST",
                url: "/api/v1/buckets/",
                data: {
                    params: { "test": "1" },
                    buckets: [bucket]
                }
            }
            return context.dispatch('axios/execute', p, { root: true }).then((res) => {
                return res;
            })
        },

        loadDrivers(context) {
            let list = BucketDrivers.listDrivers()
            context.commit('updateDrivers', list)
        },
    },
}
