import Vue from 'vue'
import "./static/bundle.scss";
import App from './App.vue'
import {securedAxiosInstance} from './utils/ifetch'
import vuetify from './plugins/vuetify';
import VueAxios from 'vue-axios';
import VueRouter from 'vue-router';
import Vuex from 'vuex';
import storeInstances from './store';

import router_setup from './router';

Vue.use(VueAxios, {
  secured: securedAxiosInstance,
})
const router = new VueRouter({
  ...router_setup
})
Vue.use(VueRouter)
Vue.use(Vuex)
const store = new Vuex.Store(storeInstances)
Vue.config.productionTip = false

new Vue({
  vuetify,
  router,
  store,
  render: h => h(App)
}).$mount('#app')
