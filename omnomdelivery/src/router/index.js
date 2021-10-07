import Vue from "vue";
import VueRouter from "vue-router";
import Home from "../views/Home.vue";
import Example from "@/components/Example";
import Cart from "@/components/Cart";

Vue.use(VueRouter);

const routes = [
    {
        path: "/",
        name: "Home",
        component: Home,
    },
    {
        path: "/cart",
        name: "Cart",
        component: Cart,
    },
    {
        path: "/about",
        name: "About",

        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () =>
            import(/* webpackChunkName: "about" */ "../views/About.vue"),
        children: [
            {
                path: "example",
                name: "Example",
                component: Example
            }
        ]
    },
];

const router = new VueRouter({
    mode: "history",
    base: process.env.BASE_URL,
    routes,
});

export default router;
