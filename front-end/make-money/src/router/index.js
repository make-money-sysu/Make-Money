import Vue from 'vue'
import Router from 'vue-router'
import MainPage from '@/components/MainPage'
import order from '@/components/order'
import personalInfo from '@/components/personalInfo'
import QuestionnaireList from '@/components/QuestionnaireList'
import Login from '@/components/Login'
import Register from '@/components/Register'
import HomePage from '@/components/HomePage'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'MainPage',
      component: MainPage
    },
    {
      path: '/order',
      name: 'order',
      component: order
    },
    {
      path: '/personalInfo',
      name: 'personalInfo',
      component: personalInfo
    },
    {
      path: '/QuestionnaireList',
      name: 'QuestionnaireList',
      component: QuestionnaireList
    },
    {
      path: '/Login',
      name: 'Login',
      component: Login
    },
    {
      path: '/Register',
      name: 'Register',
      component: Register
    },
    {
      path: '/HomePage',
      name: 'HomePage',
      component: HomePage
    }
  ]
})
