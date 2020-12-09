import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import vueHeadful from "vue-headful";
import { InitializeAction } from "@/store/actions";

Vue.component("vue-headful", vueHeadful);

Vue.config.productionTip = false;

store.dispatch(new InitializeAction());

new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount("#app");
