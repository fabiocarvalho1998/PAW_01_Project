import Vue from 'vue'
import VueRouter from 'vue-router'
import VueResource from 'vue-resource'
import VueGeolocation from 'vue-browser-geolocation'
import * as VueGoogleMaps from 'vue2-google-maps'
import BootstrapVue from 'bootstrap-vue'
import 'bootstrap/dist/css/bootstrap.css'
import 'bootstrap-vue/dist/bootstrap-vue.css'
import App from './App.vue'
import Login from './pages/Login.vue'
import Map from './pages/Map.vue'
import ListZones from './pages/ListZones.vue'


Vue.config.productionTip = false
Vue.use(VueRouter)
Vue.use(VueResource)
Vue.use(BootstrapVue)
Vue.use(VueGeolocation)
Vue.use(VueGoogleMaps,  {
  load: {
    key: 'AIzaSyDvzBG1YC-EqqOau_4BMMAgK7p-t9nrYjE',
    autobindAllEvents: true,
  },
  installComponents: true,
})

const router = new VueRouter({
  routes: [
    {
      path: '/',
      name: 'login',
      component: Login
    },
    {
      path: '/map',
      name: 'map',
      component: Map
    },
    {
      path: '/zones',
      name: 'zones',
      component: ListZones
    }
  ]
})

new Vue({
  render: h => h(App),
  router
}).$mount('#app')
