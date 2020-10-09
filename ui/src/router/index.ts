import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'
import Home from '../views/Home.vue'
import Dashboard from '../views/Dashboard.vue'
import FileManage from '../views/FileManage.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
    {
        path: '/',
        name: 'Home',
        component: Home,
        children: [
            {
                path: '/dashboard',
                component: Dashboard,
                name: 'Dashboard',
                meta: {
                    title: '仪表盘',
                    icon: 'ios-speedometer-outline'
                }
            },
            {
                path: '/file-manage',
                component: FileManage,
                name: 'FileManage',
                meta: {
                    title: '文件管理',
                    icon: 'ios-folder-open-outline'
                }
            },
        ]
    }
]

const router = new VueRouter({
    routes
})

export default router
