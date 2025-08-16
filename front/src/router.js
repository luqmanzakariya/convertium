import { createRouter, createWebHistory } from "vue-router";
import RoomSearch from "./views/RoomSearch.vue";
import Rooms from "./views/Rooms.vue";
import ContactInformation from "./views/ContactInformation.vue";
import Confirmation from "./views/Confirmation.vue";
import Login from "./views/Login.vue";
import Register from "./views/Register.vue";
import Dashboard from "./views/Dashboard.vue";

const routes = [
  { path: "/", name: "RoomSearch", component: RoomSearch },
  { path: "/rooms", name: "Rooms", component: Rooms },
  {
    path: "/contact",
    name: "ContactInformation",
    component: ContactInformation,
  },
  { path: "/confirmation", name: "Confirmation", component: Confirmation },
  { path: "/login", name: "Login", component: Login },
  { path: "/register", name: "Register", component: Register },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: Dashboard,
    meta: { requiresAuth: true },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
