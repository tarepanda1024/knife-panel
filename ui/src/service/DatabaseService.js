import axios from "@/http/axios";

export const listDB = (data) => {
    return axios.request({
        url: '/xhr/v1/dbs/list',
        method: 'get'
    })
};

export const listDBTables = (data) => {
    return axios.request({
        url: '/xhr/v1/dbs/listTables',
        params: data,
        method: 'get'
    })
};

export const listTableColumns = (data) => {
    return axios.request({
        url: '/xhr/v1/dbs/listColumns',
        params: data,
        method: 'get'
    })
};