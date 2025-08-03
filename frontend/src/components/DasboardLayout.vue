<script setup>

</script>
<template>
  <div class="bg-gradient-to-br from-gray-900 to-gray-800 min-h-screen flex flex-col">
    <!-- Header with right-aligned menu -->
    <header class="bg-gradient shadow-lg">
      <div class="container mx-auto px-4 py-4 flex justify-between items-center">
        <router-link to="/dashboard/contacts"
          class="flex items-center hover:opacity-90 transition-opacity duration-200">
          <i class="fas fa-address-book text-white text-2xl mr-3"></i>
          <div class="text-white font-bold text-xl">Contact Management</div>
        </router-link>
        <nav>
          <ul class="flex space-x-6">
            <li>
              <router-link to="/dashboard/users/profile"
                class="text-gray-100 hover:text-white flex items-center transition-colors duration-200">
                <i class="fas fa-user-circle mr-2"></i>
                <span>Profile</span>
              </router-link>
            </li>
            <li>
              <router-link to="/dashboard/users/logout"
                class="text-gray-100 hover:text-white flex items-center transition-colors duration-200">
                <i class="fas fa-sign-out-alt mr-2"></i>
                <span>Logout</span>
              </router-link>
            </li>
          </ul>
        </nav>
      </div>
    </header>

    <!-- Main content -->
    <main class="container mx-auto px-4 py-8 flex-grow">
      <router-view></router-view>
      <!-- Footer -->
      <div class="mt-10 mb-6 text-center text-gray-400 text-sm animate-fade-in">
        <p>© 2025 Contact Management. All rights reserved.</p>
      </div>
    </main>
  </div>
</template>