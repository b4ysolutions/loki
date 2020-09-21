import Vue from 'vue'
import { BootstrapVue, IconsPlugin } from 'bootstrap-vue'
import VueClipboard from 'vue-clipboard2'
import Notifications from 'vue-notification'
import App from './App.vue'


Vue.config.productionTip = false
Vue.use(BootstrapVue)
Vue.use(IconsPlugin)
Vue.use(VueClipboard)
Vue.use(Notifications)


var VueScrollTo = require('vue-scrollto')
Vue.use(VueScrollTo, {
  container: "body",
  duration: 1500,
  easing: "ease",
  offset: -45,
  force: true,
  cancelable: true,
  onStart: false,
  onDone: false,
  onCancel: false,
  x: false,
  y: true
})


import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import 'vue-material-design-icons/styles.css';

new Vue({
  render: h => h(App),
}).$mount('#app')