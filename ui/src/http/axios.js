const baseUrl = "";
import axios from 'axios'
import {Message} from 'view-design';

class HttpRequest {
    constructor(baseUrl) {
        this.baseUrl = baseUrl
        this.queue = {}
    }

    getInsideConfig() {
        return {
            baseURL: this.baseUrl,
        }
    }

    destroy(url) {
        delete this.queue[url]
    }

    interceptors(instance, url) {
        // 请求拦截
        instance.interceptors.request.use(config => {
            if (localStorage.getItem('token')) {  // 判断是否存在token，如果存在的话，则每个http header都加上token
                config.headers.Authorization = localStorage.getItem('token');
            }
            this.queue[url] = true;
            return config
        }, error => {
            console.log("req-error")
            return Promise.reject(error)
        });
        // 响应拦截
        instance.interceptors.response.use(res => {
            this.destroy(url);
            const {data, status} = res;
            if (status !== 200) {
                Message.error("网络错误");
                return false;
            }
            if (data.code !== 200) {
                if(data.msg){
                    Message.error(data.msg);
                }
                return false;
            }
            if (data.data != null) {
                return data.data
            }
            return true
        }, error => {
            Message.error("内部错误")
            this.destroy(url);
            return Promise.reject(error)
        })
    }

    request(options) {
        const instance = axios.create()
        options = Object.assign(this.getInsideConfig(), options)
        this.interceptors(instance, options.url)
        return instance(options)
    }
}


export default new HttpRequest(baseUrl);
