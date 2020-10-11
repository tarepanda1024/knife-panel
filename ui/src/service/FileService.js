import axios from '../http/axios'

export const listFiles = (data) => {
    return axios.request({
        url: '/xhr/v1/files/list',
        params: data,
        method: 'get'
    })
};