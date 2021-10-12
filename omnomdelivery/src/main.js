import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import Card from "@/components/Card";
import Product from "./components/Product";
import Products_section from "@/components/Products_section";
import Supplier_section from "@/components/Supplier_section";
import Supplier from "@/components/Supplier";
import Registration from "@/components/Registration";
import Login from "@/components/Login";
import UserProfile from "@/components/UserProfile";
import Logout from "@/components/Logout";
import DialogWindow from "./components/UI/DialogWindow";
import Cart from "./components/Cart";
import Header from "./components/Header";

Vue.config.productionTip = false;
Vue.component("card", Card).default;
Vue.component("product", Product).default;
Vue.component("products", Products_section).default;
Vue.component("suppliers", Supplier_section).default;
Vue.component("supplier", Supplier).default;
Vue.component("registration", Registration).default;
Vue.component("login", Login).default;
Vue.component("user_profile", UserProfile).default;
Vue.component("logout", Logout).default;
Vue.component("dialog_window", DialogWindow).default;
Vue.component("cart", Cart).default
Vue.component('omnom_header', Header).default


new Vue({
    router,
    store,
    render: (h) => h(App),
}).$mount("#app");
