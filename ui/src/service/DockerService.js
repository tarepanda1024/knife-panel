import axios from '../http/axios'

export const overview = (data) => {
    return axios.request({
        url: '/xhr/v1/dockers/overview',
        method: 'get'
    })
};

export const listContainers = (data) => {
    return axios.request({
        url: '/xhr/v1/dockers/listContainers',
        method: 'get'
    })
};

export const listImages = (data) => {
    return axios.request({
        url: '/xhr/v1/dockers/listImages',
        method: 'get'
    })
};

export const listNets = (data) => {
    return axios.request({
        url: '/xhr/v1/dockers/listNets',
        method: 'get'
    })
};

export const listVolumes = (data) => {
    return axios.request({
        url: '/xhr/v1/dockers/listVolumes',
        method: 'get'
    })
};