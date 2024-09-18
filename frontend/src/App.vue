<template>
  <div id="app" class="container">
    <nav class="navbar navbar-expand-lg navbar-light bg-light">
      <a class="navbar-brand mx-auto" href="#">My Blog</a>
      <div class="collapse navbar-collapse">
        <ul class="navbar-nav ml-auto">
          <li class="nav-item">
            <router-link class="nav-link btn btn-primary" to="/register">Register</router-link>
          </li>
          <li class="nav-item">
            <router-link class="nav-link btn btn-primary" to="/login">Login</router-link>
          </li>
        </ul>
      </div>
    </nav>

    <router-view />
  </div>
</template>

<script>
export default {
  name: "App",
};
</script>

<style>
/* Centering the brand in the navbar */
.navbar-brand {
  margin: 0 auto;
}

/* Aligning the nav items to the right */
.navbar-nav {
  margin-left: auto;
}

/* Ensure nav-links appear as buttons with primary background color */
.nav-link {
  color: #fff; /* Ensure text color is white */
  background-color: #007bff; /* Bootstrap primary blue color */
  border-radius: 0.25rem;
  padding: 0.5rem 1rem;
  margin-left: 0.5rem; /* Space between buttons */
  text-align: center; /* Center text */
  display: inline-block; /* Make sure the link appears as a button */
}

.nav-link:hover {
  background-color: #0056b3; /* Darker blue on hover */
}
</style>
