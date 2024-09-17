<template>
    <div class="container mt-5">
      <h2>Register</h2>
      <form @submit.prevent="registerUser">
        <div class="mb-3">
          <label for="username" class="form-label">Username</label>
          <input
            type="text"
            class="form-control"
            id="username"
            v-model="username"
            required
          />
        </div>
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
        <button type="submit" class="btn btn-primary">Register</button>
      </form>
    </div>
  </template>
  
  <script>
  import axios from "axios";
  
  export default {
    data() {
      return {
        username: "",
        email: "",
        password: "",
      };
    },
    methods: {
      async registerUser() {
        try {
          const response = await axios.post("http://localhost:8080/api/auth/register", {
            username: this.username,
            email: this.email,
            password: this.password,
          });
          console.log("Registration successful:", response.data);
          //alert
          alert("Registration successful");
        } catch (error) {
          alert(error.response.data.message);
          console.error("Error registering user:", error.response.data);
        }
      },
    },
  };
  </script>
  
  <style>

  </style>
  