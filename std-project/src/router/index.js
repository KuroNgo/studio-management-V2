import HomeViewVue from '@/views/User/HomeView.vue';
import * as vueRouter from 'vue-router';
const routes = [
  {
    path: '/',
    name: 'Trang chủ | Cỏ Studio',
    component: HomeViewVue,
  },

  {
    path: '/about',
    name: 'Giới thiệu | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/User/AboutView.vue'),
   
  },

  {
    path: '/services',
    name: 'Dịch vụ | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/User/ServiceView.vue')
  },

  {
    path: '/contact',
    name: 'Liên hệ | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/User/ContactView.vue')
  },

  {
    path: '/product',
    name: 'Thư viện | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/User/ProductView.vue')
  },

  {
    path: '/booking',
    name: 'Đặt lịch | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/User/BookingView.vue')
  },

  {
    path: '/payment',
    name: 'Thanh toán | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/User/PaymentView.vue')
  },

  
  // Admin
  {
    path: '/account',
    name: 'Quản lý tài khoản | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Admin/AccountManagerView.vue')
  },
  
  {
    path: '/admin',
    name: 'Trang Admin | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Admin/AdminView.vue')
  },

  {
    path: '/booking-manager',
    name: 'Quản lý đặt lịch | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Admin/BookingManagerView.vue')
  },

  {
    path: '/customer-manager',
    name: 'Quản lý khách hàng | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Admin/CustomerManagerView.vue')
  },

  {
    path: '/product-manager',
    name: 'Quản lý sản phẩm | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Admin/ProductManagerView.vue')
  },

  {
    path: '/schedule',
    name: 'Quản lý lịch làm việc | Cỏ Studio',
    component: () => import(/* webpackChunkName: "about" */ '../views/Admin/ScheduleView.vue')
  },

  // Client
  {
    path: '/login',
    name: 'Login | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/LoginView.vue')
  },
  {
    path: '/register',
    name: 'Register | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/RegisterView.vue')
  },
  
  {
    path: '/interaction',
    name: 'Interaction | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/InteractionView.vue')
  },

  {
    path: '/blog',
    name: 'Blog | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/BlogView.vue')
  },

  {
    name:"OPPS | Page Not Found",
    path:'/:pathMatch(.*)*',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component:()=> import ('@/views/Client/NotFoundView.vue')
  }

];

const router = vueRouter.createRouter({
  history: vueRouter.createWebHistory(),
  routes: routes,

  // Cuộn lên đầu trang khi chuyển Route
  scrollBehavior(to, from, savedPosition) {
    if (to.hash) {
      // Nếu route đích có anchor (ví dụ: #section1)
      return { el: to.hash, behavior: 'smooth' };
    } else if (savedPosition) {
      // Nếu đã lưu vị trí cuộn trước đó
      return savedPosition;
    } else {
      // Mặc định, kéo lên đầu trang
      return { top: 0 };
    }
  },
});

router.beforeEach((to, from, next) => {
  document.title=` ${ to.name } `
  next()
})
export default router