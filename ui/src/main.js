import Vue from "vue";

import App from "./App.vue";
import router from "./router";
import store from "./store";

import "./assets/fonts/mr-robot.ttf";
import "./assets/1.png";
import vuetify from "./plugins/vuetify";

Vue.config.productionTip = false;

new Vue({
	router,
	store,
	vuetify,
	render: h => h(App),
}).$mount("#app");
