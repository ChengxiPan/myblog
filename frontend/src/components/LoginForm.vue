<template>
    <div class="container mt-5">
      <h2>Login</h2>
      <form @submit.prevent="loginUser">
        <div class="mb-3">
          <label for="email" class="form-label">Email address</label>
          <input
            type="email"
            class="form-control"
            id="email"
            v-model="email"
            required
          />
        </div>
        <div class="mb-3">
          <label for="password" class="form-label">Password</label>
          <input
            type="password"
            class="form-control"
            id="password"
            v-model="password"
            required
          />
        </div>
        <button type="submit" class="btn btn-primary">Login</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        email: "",
        password: "",
      };
    },
    methods: {
      async loginUser() {
        try {
          const response = await axios.post("http://localhost:8080/api/auth/login", {
            email: this.email,
            password: this.password,
          });
          alert("Login successful");
          console.log("Login successful:", response.data);
          // Redirect to the home page. home page: /home
            this.$router.push("/home");
        } catch (error) {
            if (error.response && error.response.data && error.response.data.error) {
                alert("Error: " + error.response.data.error);
            } else {
                alert("An error occurred during login.");
            }
            console.error("Error logging in:", error.response.data);
        }

      },
    },
  };
  </script>
  
  <style>
  /* Add any additional styling here if needed */
  </style>
  