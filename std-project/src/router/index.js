import HomeViewVue from '@/views/Client/HomeView.vue';
import * as vueRouter from 'vue-router';
const routes = [
  //********************************************************CLIENT********************************************** */
  {
    path: '/',
    name: 'Trang chủ | Cỏ Studio',
    component: HomeViewVue,
    meta:{
      showHeaderFooter: true
    }
  },

  {
    path: '/about',
    name: 'Giới thiệu | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/AboutView.vue'),
    meta:{
      showHeaderFooter: true
    }
  },

  {
    path: '/services',
    name: 'Dịch vụ | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/ServicesView.vue'),
    children: [
      {
        path: 'eventservices',
        name: 'Dịch vụ Sự kiện | Cỏ Studio',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Client/Services/EVServiceView.vue'),
        meta:{
          showHeaderFooter: true
        }
      },
      {
        path: 'comercialservice',
        name: 'Dịch vụ Thương mại | Cỏ Studio',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Client/Services/ComercialServiceView.vue'),
        meta:{
          showHeaderFooter: true
        }
      },
      {
        path: 'flycamservice',
        name: 'Dịch vụ Flycam | Cỏ Studio',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Client/Services/FlycamServiceView.vue'),
        meta:{
          showHeaderFooter: true
        }
      },
      {
        path: 'livestream',
        name: 'Dịch vụ Livestream | Cỏ Studio',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Client/Services/LiveServiceView.vue'),
        meta:{
          showHeaderFooter: true
        }
      },
      {
        path: 'weddingservices',
        name: 'Dịch vụ Đám cưới | Cỏ Studio',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/Client/Services/WedServiceView.vue'),
        meta:{
          showHeaderFooter: true
        }
      },
    ]
  },

  {
    path: '/contact',
    name: 'Liên hệ | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/ContactView.vue'),
    meta:{
      showHeaderFooter: true
    }
  },

  {
    path: '/collection',
    name: 'Bộ sưu tập | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/CollectionView.vue'),
    meta:{
      showHeaderFooter: true
    }
  },

  {
    name: "OPPS | Page Not Found",
    path: '/:pathMatch(.*)*',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import('@/views/Client/NotFoundView.vue')
  },

  //********************************************************CLIENT/ACCOUNT********************************************** */
  {
    path: '/login',
    name: 'Login | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/Account/LoginView.vue')
  },
  {
    path: '/register',
    name: 'Register | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/Account/RegisterView.vue')
  },
  //********************************************************CLIENT/ACCOUNT********************************************** */

  //********************************************************CLIENT/ACCOUNT/ACCOUNT-RECOVERY********************************************** */
  {
    path: '/forgetpassword',
    name: 'Account Recovery | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/Account/AccountRecovery/ForgetPassView.vue')
  },
  {
    path: '/recoverpassword',
    name: 'Account Recovery | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/Account/AccountRecovery/RecoverPassView.vue')
  },
  {
    path: '/resetpassword',
    name: 'Account Recovery | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/Account/AccountRecovery/ResetPassView.vue')
  },
  {
    path: '/sentcode',
    name: 'Account Recovery | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/Client/Account/AccountRecovery/SentCodeView.vue')
  },
  //********************************************************CLIENT/ACCOUNT/ACCOUNT-RECOVERY********************************************** */


  //********************************************************CLIENT********************************************** */


  //********************************************************USER********************************************** */

  {
    path: '/payment',
    name: 'Thanh toán | Cỏ Studio',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/User/PaymentView.vue')
  },

  //********************************************************USER********************************************** */

  //********************************************************ADMIN********************************************** */

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

  //********************************************************ADMIN********************************************** */



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
  document.title = ` ${to.name} `
  next()
})
export default router
