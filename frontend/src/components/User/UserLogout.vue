<script setup>
import { useLocalStorage } from '@vueuse/core';
import { onBeforeMount } from 'vue';
import { userLogOut } from '../../../lib/api/userApi';
import { useRouter } from 'vue-router';
import { alertError } from '../../../lib/alert';

const token = useLocalStorage("token", "");
const router = useRouter();

onBeforeMount(async () => {
  const response = await userLogOut(token.value);
  const responseBody = await response.json();
  console.log(responseBody);
  if (response.status === 200) {
    token.value = ""
    router.push({
      path: "/login"
    })
  } else {
    await alertError(responseBody.errors);
  }
})

</script>

<template>

</template>