import axios from 'axios'
import { API_BASE_URL } from './config'

const apiClient=axios.create({
    baseURL:API_BASE_URL,
    timeout: 10000,
    headers: {
        "Content-Type": "application/json"
    }    
})



apiClient.interceptors.request.use(function (config) {
    console.log("config", config);
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;}, function (error) {
    return Promise.reject(error);
  });


apiClient.interceptors.response.use(function (response) {
    return response;
  }, function (error) {
    return Promise.reject(error);
  });


  export const Request={
    get:(url,params={})=>apiClient.get(url,{params}),
    post:(url,data={})=>apiClient.post(url,data),
    put:(url,data={})=>apiClient.put(url,data),
    delete:(url,data={})=>apiClient.delete(url,data)
  }