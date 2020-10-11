import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './plugins/iview.js'
import moment from "moment";

Vue.prototype.$moment = moment;
Vue.config.ignoredElements = [/^ion-/]
Vue.config.productionTip = false

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
