<template>
  <div class="register-form">
    <form @submit.prevent="handleSubmit">
      <div class="form-group">
        <input
          type="text"
          id="name"
          v-model="form.name"
          required
          placeholder="Enter your full name"
        />
      </div>

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
          minlength="8"
        />
      </div>

      <div class="form-group" style="margin-bottom: 0px">
        <input
          type="password"
          id="password_confirmation"
          v-model="form.password_confirmation"
          required
          placeholder="Confirm your password"
          minlength="8"
        />
      </div>
      <div
        v-if="form.password && form.password_confirmation && !passwordsMatch"
        class="validation-error"
      >
        Passwords do not match
      </div>

      <button
        type="submit"
        class="submit-button"
        :disabled="loading || !passwordsMatch"
      >
        {{ loading ? "Creating account..." : "Register" }}
      </button>

      <div v-if="error" class="error-message">{{ error }}</div>

      <div class="login-link">
        Already have an account?
        <router-link to="/login">Login here</router-link>
      </div>
    </form>
  </div>
</template>

<script>
import apiService from "../../../service/api";
import { API_ROUTE } from "../../../constant/url";

export default {
  name: "RegisterForm",
  data() {
    return {
      form: {
        name: "",
        email: "",
        password: "",
        password_confirmation: "",
      },
      loading: false,
      error: "",
    };
  },
  computed: {
    passwordsMatch() {
      return this.form.password === this.form.password_confirmation;
    },
  },
  methods: {
    async handleSubmit() {
      if (!this.passwordsMatch) return;

      this.loading = true;
      this.error = "";

      try {
        const response = await apiService.post(
          `${API_ROUTE.REGISTER}`,
          this.form,
          {
            headers: {
              "Content-Type": "application/json",
            },
          }
        );

        this.$emit("success");
      } catch (err) {
        this.error =
          err.response?.data?.message ||
          err.response?.data?.error ||
          "Registration failed. Please try again.";
      } finally {
        this.loading = false;
      }
    },
  },
};
</script>

<style scoped>
.register-form {
  max-width: 400px;
  margin: 0 auto;
  padding: 2rem;
  background: #fff;
  border-radius: 8px;
  box-shadow: 0 2px 10px rgba(0, 0, 0, 0.1);
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

.validation-error {
  color: #e74c3c;
  font-size: 0.8rem;
  margin-top: 0.25rem;
  text-align: left;
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
  margin-top: 1.5rem;
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

.login-link {
  margin-top: 1.5rem;
  font-size: 0.9rem;
  color: #666;
  text-align: center;
}

.login-link a {
  color: #000;
  text-decoration: none;
  font-weight: 500;
}

.login-link a:hover {
  text-decoration: underline;
}
</style>
