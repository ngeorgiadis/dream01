import Vue from "vue";
import VueMaterial from "vue-material";

import "vue-material/dist/vue-material.min.css";
import "vue-material/dist/theme/default.css";

import App from "./App.vue";
import router from "./router";
import store from "./store";
import "./assets/fonts/mr-robot.ttf";
import "./assets/1.png";

Vue.use(VueMaterial);

Vue.config.productionTip = false;

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount("#app");
