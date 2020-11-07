import Vue from 'vue'
import VueRouter, {RouteConfig} from 'vue-router'
import Home from '../views/Home.vue'
import Dashboard from '../views/Dashboard.vue'
import FileManage from '../views/FileManage.vue'
import Docker from '../views/Docker.vue'
import WebServer from '../views/WebServer.vue'
import AppStore from '../views/AppStore.vue'
import Database from '../views/Database.vue'
import Setting from '../views/SystemSetting.vue'
import SystemSetting from "@/views/SystemSetting.vue";
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
                    icon: 'speedometer'
                }
            },
            {
                path: '/database',
                component: Database,
                name: 'Database',
                meta: {
                    title: '数据库',
                    icon: 'server'
                }
            },
            {
                path: '/file-manage',
                component: FileManage,
                name: 'FileManage',
                meta: {
                    title: '文件管理',
                    icon: 'folder'
                }
            },
            {
                path: '/web-server',
                component: WebServer,
                name: 'WebServer',
                meta: {
                    title: '服务管理',
                    icon: 'logo-firefox'
                }
            },
            {
                path: '/docker',
                component: Docker,
                name: 'Docker',
                meta: {
                    title: '容器管理',
                    icon: 'logo-docker'
                }
            },
            {
                path: '/app-store',
                component: AppStore,
                name: 'AppStore',
                meta: {
                    title: '应用市场',
                    icon: 'logo-apple-appstore'
                }
            },
            {
                path: '/system-setting',
                component: SystemSetting,
                name: 'SystemSetting',
                meta: {
                    title: '系统设置',
                    icon: 'settings'
                }
            },
        ]
    }
]

const router = new VueRouter({
    routes
})

export default router
