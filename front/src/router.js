import { createRouter, createWebHistory } from "vue-router";
import RoomSearch from "./views/RoomSearch.vue";
import Rooms from "./views/Rooms.vue";
import ContactInformation from "./views/ContactInformation.vue";
import Confirmation from "./views/Confirmation.vue";
import Login from "./views/Login.vue";
import Register from "./views/Register.vue";
import Dashboard from "./views/Dashboard.vue";
import { APP_ROUTE } from "./constant/url";

const routes = [
  { path: APP_ROUTE.HOME, name: "RoomSearch", component: RoomSearch },
  { path: APP_ROUTE.ROOMS, name: "Rooms", component: Rooms },
  {
    path: APP_ROUTE.CONTACT_INFORMATION,
    name: "ContactInformation",
    component: ContactInformation,
  },
  {
    path: APP_ROUTE.CONFIRMATION,
    name: "Confirmation",
    component: Confirmation,
    meta: { requiresAuth: true },
  },
  { path: APP_ROUTE.LOGIN, name: "Login", component: Login },
  { path: APP_ROUTE.REGISTER, name: "Register", component: Register },
  {
    path: APP_ROUTE.DASHBOARD,
    name: "Dashboard",
    component: Dashboard,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

router.beforeEach((to, from, next) => {
  const token = localStorage.getItem("token");
  if (to.meta.requiresAuth && !token) {
    next(APP_ROUTE.LOGIN);
  } else {
    next();
  }
});

export default router;
