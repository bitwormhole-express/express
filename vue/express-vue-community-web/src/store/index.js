import { createStore } from 'vuex'
import AxiosModule from './AxiosModule.js'
import BucketModule from './BucketModule.js'
import SessionModule from './SessionModule.js'

export default createStore({
  state: {
  },
  mutations: {
  },
  actions: {
  },
  modules: {
    "axios": AxiosModule,
    "session": SessionModule,
    "bucket": BucketModule,
  }
})
