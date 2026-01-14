<template>
  <div class="login-container">
    <div class="login-card">
      <h2>خوش آمدید</h2>
      <p>لطفاً نام کاربری و رمز عبور خود را وارد کنید</p>

      <form @submit.prevent="handleLogin">
        <div class="input-group">
          <label>نام کاربری</label>
          <input
            v-model="username"
            type="text"
            placeholder="Username"
            required
          />
        </div>

        <div class="input-group">
          <label>رمز عبور</label>
          <input
            v-model="password"
            type="password"
            placeholder="Password"
            required
          />
        </div>

        <button type="submit" :disabled="isLoading">
          {{ isLoading ? 'در حال ورود...' : 'ورود به پنل ادمین' }}
        </button>

        <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
      </form>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'


const username = ref('')
const password = ref('')
const isLoading = ref(false)
const errorMessage = ref('')
const router = useRouter()

const handleLogin = async () => {
  isLoading.value = true
  errorMessage.value = ''

  try {
    const response = await axios.post('http://localhost:8000/api/v1/login', {
      username: username.value,
      password: password.value
    })
    localStorage.setItem('token', response.data.token)
    router.push('/dashboard')
  } catch (error) {
    errorMessage.value = 'نام کاربری یا رمز عبور اشتباه است'
  } finally {
    isLoading.value = false
  }
}
</script>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  font-family: 'Tahoma', sans-serif;
  direction: rtl;
}

.login-card {
  background: rgba(255, 255, 255, 0.9);
  padding: 2rem;
  border-radius: 15px;
  box-shadow: 0 10px 25px rgba(0,0,0,0.2);
  width: 100%;
  max-width: 400px;
  text-align: center;
}

h2 { color: #333; margin-bottom: 0.5rem; }
p { color: #666; font-size: 0.9rem; margin-bottom: 1.5rem; }

.input-group {
  text-align: right;
  margin-bottom: 1rem;
}

label { display: block; margin-bottom: 0.5rem; color: #444; }

input {
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-sizing: border-box; /* برای جلوگیری از بیرون زدن اینپوت */
}

button {
  width: 100%;
  padding: 12px;
  background-color: #764ba2;
  color: white;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1rem;
  transition: 0.3s;
}

button:hover { background-color: #5a368c; }
button:disabled { background-color: #ccc; }

.error { color: #ff4d4d; margin-top: 1rem; font-size: 0.85rem; }
</style>
