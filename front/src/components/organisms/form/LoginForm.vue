<template>
  <div class="login-form">
    <form @submit.prevent="handleSubmit" class="form">
      <div class="form-group">
        <input
          type="email"
          id="email"
          v-model="form.email"
          required
          placeholder="Enter your email"
        />
      </div>
      <div class="form-group">
        <input
          type="password"
          id="password"
          v-model="form.password"
          required
          placeholder="Enter your password"
          minlength="4"
        />
      </div>
      <button type="submit" class="submit-button" :disabled="loading">
        {{ loading ? "Logging in..." : "Login" }}
      </button>
      <div v-if="error" class="error-message">{{ error }}</div>
    </form>
  </div>
</template>

<script>
import axios from "axios";
import { APP_ROUTE } from "../../../constant/url";

export default {
  name: "LoginForm",
  data() {
    return {
      form: {
        email: "",
        password: "",
      },
      loading: false,
      error: "",
    };
  },
  methods: {
    async handleSubmit() {
      this.loading = true;
      this.error = "";

      try {
        const response = await axios.post(
          `${import.meta.env.VITE_API_URL}/login`,
          {
            email: this.form.email,
            password: this.form.password,
          },
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );

        const token = response.data.token;

        if (token) {
          localStorage.setItem("token", token);
          this.$router.push(APP_ROUTE.HOME);
          this.$emit("success");
        } else {
          throw new Error("No token received");
        }
      } catch (err) {
        if (err?.response?.status === 401) {
          this.error = "Invalid email or password.";
        } else if (err?.response?.status === 403) {
          this.error =
            "Your account is not activated. Please check your email.";
        } else {
          this.error =
            err.response?.data?.message ||
            err.message ||
            "Login failed. Please try again.";
        }
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
.login-form {
  max-width: 400px;
  margin: 0 auto;
  padding: 2rem;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
}

.form {
  width: 100%;
}

.form-group {
  margin-bottom: 1.5rem;
  display: flex;
}

label {
  display: block;
  margin-bottom: 0.5rem;
  font-weight: 500;
  color: #333;
}

input {
  width: 100%;
  padding: 0.75rem;
  border: 1px solid #ddd;
  border-radius: 4px;
  font-size: 1rem;
}

input:focus {
  outline: none;
  border-color: #000;
}

.submit-button {
  width: 100%;
  padding: 0.75rem;
  background-color: #000;
  color: white;
  border: none;
  border-radius: 4px;
  font-size: 1rem;
  font-weight: 500;
  cursor: pointer;
  transition: background-color 0.3s;
}

.submit-button:hover {
  background-color: #333;
}

.submit-button:disabled {
  background-color: #ccc;
  cursor: not-allowed;
}

.error-message {
  margin-top: 1rem;
  color: #e74c3c;
  text-align: center;
}
</style>
