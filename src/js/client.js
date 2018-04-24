'use strict';

import axios from 'axios';

import AppError from './error.js';

let AppClient = function() {
  this.axios = axios.create();
  this.axios.interceptors.response.use((response) => {
    let data = response.data;
    if (data.state === 'error') {
      let result = data.result;
      return Promise.reject(new AppError(result.message, result.code));
    }
    return data.result;
  })
};
AppClient.prototype.name = 'AppClient';

AppClient.prototype.initToken = function() {
  delete this.axios.defaults.headers.common['Authorization'];
};

AppClient.prototype.setToken = function(token) {
  this.axios.defaults.headers.common['Authorization'] = 'Bearer ' + token;
};

AppClient.prototype.challenge = function(address) {
  let p = new URLSearchParams();
  p.append('address', address);

  return this.axios.post('/auth/challenge', p);
};

AppClient.prototype.authorize = function(address, signature) {
  let p = new URLSearchParams();
  p.append('address', address);
  p.append('signature', signature);

  return this.axios.post('/auth/authorize', p);
};

AppClient.prototype.getUser = function(address) {
  return this.axios.get('/api/users/' + address);
};

AppClient.prototype.updateUser = function(address, params) {
  let p = new URLSearchParams();
  p.append('name', params.name);

  return this.axios.put('/api/users/' + address, p);
};

AppClient.prototype.deleteUser = function(address) {
  return this.axios.delete('/api/users/' + address);
};

module.exports = AppClient;
