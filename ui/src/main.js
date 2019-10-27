import Vue from "vue";
import VueMaterial from "vue-material";

import App from "./App.vue";
import router from "./router";
import store from "./store";

Vue.use(VueMaterial);
Vue.config.productionTip = false;

import "./assets/fonts/mr-robot.ttf";
import "vue-material/dist/vue-material.min.css";
import "vue-material/dist/theme/default.css";

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
