<template>
    <div class="container mt-5">
      <h1>Welcome to the Home Page</h1>
      <p>Here you can see your posts and explore the blog.</p>
  
      <div class="mt-4">
        <button class="btn btn-primary" @click="createPost">Create New Post</button>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'HomePage',
    methods: {
      createPost() {
        this.$router.push('/create-post');  // Navigate to create post page
      }
    }
  };
  </script>
  
  <style scoped>
  h1 {
    color: #007bff;
  }
  </style>
  