<template>
  <div class="dashboard-layout">
    <aside class="sidebar">
      <div class="user-profile">
        <div class="avatar">Admin</div>
        <span>پنل مدیریت اموال</span>
      </div>

      <nav class="nav-menu">
        <router-link to="/dashboard" class="nav-item" active-class="active">
          <i class="icon">📊</i> داشبورد
        </router-link>

        <router-link to="/assets/create" class="nav-item" active-class="active">
          <i class="icon">➕</i> ثبت اموال جدید
        </router-link>

        <router-link to="/reports" class="nav-item" active-class="active">
          <i class="icon">📋</i> گزارشات
        </router-link>
      </nav>

      <div class="sidebar-footer">
        <button @click="handleLogout" class="nav-item logout-btn">
          <i class="icon">🚪</i> خروج از سیستم
        </button>
      </div>
    </aside>

    <main class="main-content">
      <header class="top-bar">
        <P style="font-size: 40px">داشبورد مدیریتی</P>
        <div class="search-box">
          <input type="text" v-model="searchQuery" placeholder="جستجو در اموال..." />
          <i class="search-icon">🔍</i>
        </div>
      </header>

      <div class="sections-container">
        <section class="asset-section">
          <div class="section-header">
            <h2>لیست اموال اداری</h2>
            <span class="badge">اداری</span>
          </div>
          <div class="table-container">
            <table>
              <thead>
              <tr>
                <th>کد اموال</th>
                <th>نام کالا</th>
                <th>وضعیت</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="item in filteredOfficeAssets" :key="item.id">
                <td>{{ item.code }}</td>
                <td>{{ item.name }}</td>
                <td><span class="status-ok">موجود</span></td>
              </tr>
              </tbody>
            </table>
          </div>
          <div class="pagination">
            <button>بعدی</button>
            <span>صفحه ۱ از ۵</span>
            <button>قبلی</button>
          </div>
        </section>

        <section class="asset-section">
          <div class="section-header">
            <h2>لیست اموال فنی</h2>
            <span class="badge tech">فنی</span>
          </div>
          <div class="table-container">
            <table>
              <thead>
              <tr>
                <th>کد اموال</th>
                <th>نام کالا</th>
                <th>بخش اختصاصی</th>
              </tr>
              </thead>
              <tbody>
              <tr v-for="item in filteredTechAssets" :key="item.id">
                <td>{{ item.code }}</td>
                <td>{{ item.name }}</td>
                <td>{{ item.department }}</td>
              </tr>
              </tbody>
            </table>
          </div>
          <div class="pagination">
            <button>بعدی</button>
            <span>صفحه ۱ از ۳</span>
            <button>قبلی</button>
          </div>
        </section>
      </div>
    </main>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue'
import { useRouter } from 'vue-router'

const router = useRouter()
const searchQuery = ref('')

// داده‌های نمونه (اتصال به بک‌اند Go در آینده)
const officeAssets = ref([
  { id: 1, code: 'OFF-101', name: 'صندلی مدیریتی' },
  { id: 2, code: 'OFF-102', name: 'میز کنفرانس' },
])

const techAssets = ref([
  { id: 1, code: 'TEC-501', name: 'سرور HP G10', department: 'IT' },
  { id: 2, code: 'TEC-502', name: 'سوییچ سیسکو', department: 'شبکه' },
])

// فیلتر کردن لیست‌ها بر اساس جستجو
const filteredOfficeAssets = computed(() => {
  return officeAssets.value.filter(item => item.name.includes(searchQuery.value))
})

const filteredTechAssets = computed(() => {
  return techAssets.value.filter(item => item.name.includes(searchQuery.value))
})

// منطق خروج
const handleLogout = () => {
  if (confirm('آیا قصد خروج از سامانه را دارید؟')) {
    localStorage.removeItem('token') // پاک کردن توکن
    router.push('/') // هدایت به صفحه ورود
  }
}
</script>

<style scoped>
.dashboard-layout {
  display: flex;
  height: 100vh;
  background-color: #f4f7f6;
  direction: rtl;
  font-family: Tahoma, sans-serif;
}

/* Sidebar */
.sidebar {
  width: 260px;
  background-color: #2c3e50;
  color: white;
  padding: 20px;
  display: flex;
  flex-direction: column;
}

.user-profile {
  text-align: center;
  padding-bottom: 20px;
  border-bottom: 1px solid #34495e;
  margin-bottom: 20px;
}

.avatar {
  width: 60px;
  height: 60px;
  background: #3498db;
  border-radius: 50%;
  margin: 0 auto 10px;
  display: flex;
  align-items: center;
  justify-content: center;
  font-weight: bold;
}

.nav-menu {
  flex-grow: 1;
}

.nav-item {
  display: flex;
  align-items: center;
  gap: 10px;
  text-decoration: none;
  background: none;
  border: none;
  color: #bdc3c7;
  padding: 15px;
  text-align: right;
  width: 100%;
  cursor: pointer;
  font-size: 0.95rem;
  border-radius: 8px;
  transition: 0.3s;
  margin-bottom: 5px;
  box-sizing: border-box;
}

.nav-item:hover, .nav-item.active {
  background: #34495e;
  color: white;
}

.icon {
  font-style: normal;
}

/* Logout Button */
.sidebar-footer {
  border-top: 1px solid #34495e;
  padding-top: 10px;
}

.logout-btn {
  color: #fab1a0;
}

.logout-btn:hover {
  background-color: #c0392b;
  color: white;
}

/* Main Content */
.main-content {
  flex: 1;
  padding: 30px;
  overflow-y: auto;
}

.top-bar {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 30px;
}

.search-box {
  position: relative;
}

.search-box input {
  padding: 10px 40px 10px 15px;
  border-radius: 20px;
  border: 1px solid #ddd;
  width: 300px;
  font-family: inherit;
}

/* Sections & Tables */
.sections-container {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 20px;
}

.asset-section {
  background: white;
  padding: 20px;
  border-radius: 12px;
  box-shadow: 0 4px 6px rgba(0,0,0,0.05);
}

.section-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 20px;
}

.badge {
  background: #e8f4fd;
  color: #3498db;
  padding: 4px 12px;
  border-radius: 20px;
  font-size: 0.8rem;
}

.badge.tech {
  background: #fef4e8;
  color: #f39c12;
}

table {
  width: 100%;
  border-collapse: collapse;
}

th {
  text-align: right;
  padding: 12px;
  background: #f8f9fa;
  color: #7f8c8d;
}

td {
  padding: 12px;
  border-bottom: 1px solid #eee;
}

.pagination {
  margin-top: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 15px;
}

.pagination button {
  padding: 5px 15px;
  border: 1px solid #ddd;
  background: white;
  cursor: pointer;
  border-radius: 4px;
}
</style>
