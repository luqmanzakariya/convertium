<template>
  <header class="masthead">
    <nav>
      <ul>
        <li><router-link :to="homeLink">Home</router-link></li>
      </ul>
    </nav>
    <nav>
      <ul>
        <template v-if="isAuthenticated">
          <li>
            <router-link :to="dashboardLink"
              ><NavbarButton title="Dashboard"
            /></router-link>
          </li>
          <li>
            <NavbarButton title="Logout" @click="logout" />
          </li>
        </template>
        <template v-else>
          <li>
            <router-link :to="loginLink"
              ><NavbarButton title="Login"
            /></router-link>
          </li>
          <li>
            <router-link :to="registerLink"
              ><NavbarButton title="Register"
            /></router-link>
          </li>
        </template>
      </ul>
    </nav>
  </header>
</template>

<script>
import NavbarButton from "../../atoms/button/NavbarButton.vue";
import { APP_ROUTE } from "../../../constant/url";

export default {
  name: "Mashead",
  components: { NavbarButton },
  data() {
    return {
      homeLink: APP_ROUTE.HOME,
      loginLink: APP_ROUTE.LOGIN,
      registerLink: APP_ROUTE.REGISTER,
      dashboardLink: APP_ROUTE.DASHBOARD,
      isAuthenticated: false,
    };
  },
  created() {
    this.checkAuthStatus();
    // Listen for auth changes from other components
    window.addEventListener("storage", this.checkAuthStatus);
  },
  beforeUnmount() {
    // Clean up event listener
    window.removeEventListener("storage", this.checkAuthStatus);
  },
  methods: {
    checkAuthStatus() {
      this.isAuthenticated = !!localStorage.getItem("token");
    },
    logout() {
      localStorage.removeItem("token");
      this.isAuthenticated = false;
      this.$router.push(APP_ROUTE.HOME);

      window.dispatchEvent(new Event("storage"));
    },
  },
};
</script>

<style scoped>
.masthead {
  background: #333;
  color: white;
  padding: 1rem;
  font-weight: 700;
  display: flex;
  justify-content: space-between;
  align-items: center;
}
ul {
  padding: 0;
  margin: 0;
}
nav ul {
  display: flex;
  justify-content: end;
  list-style: none;
  gap: 1rem;
}

nav a {
  color: white;
  text-decoration: none;
  transition: color 0.3s, border-bottom 0.3s;
  border-bottom: 2px solid transparent;
}

nav a:hover {
  color: #ffd700;
  border-bottom: 2px solid #ffd700;
}
</style>
