import Vue from "vue";
import App from "./App.vue";
import VueSvgGauge from "vue-svg-gauge";

Vue.config.productionTip = false;

Vue.use(VueSvgGauge);

new Vue({
  render: (h) => h(App),
}).$mount("#app");
