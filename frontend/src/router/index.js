import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import TodoView from "@/views/TodoView.vue";
import DiagView from "@/views/DiagView.vue";

const routes = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  // {
  //   path: '/about',
  //   name: 'about',
  //   // route level code-splitting
  //   // this generates a separate chunk (about.[hash].js) for this route
  //   // which is lazy-loaded when the route is visited.
  //   component: function () {
  //     return import(/* webpackChunkName: "about" */ '../views/AboutView.vue')
  //   }
  // },
  {
    path: '/todo',
    name: 'todo',
    component: TodoView
  },
  {
    path: '/diag',
    name: 'diag',
    component: DiagView
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
