'use strict';

import Vue from 'vue/dist/vue.esm.js';
import ElementUI from 'element-ui';
import locale from 'element-ui/lib/locale/lang/en';
import Web3 from 'web3';
import axios from 'axios';

import AppError from './error.js';

import '../css/reset.css';
import 'element-ui/lib/theme-chalk/index.css';
import '../css/style.css';

Vue.use(ElementUI, {locale});

let client = axios.create();
client.interceptors.response.use((response) => {
  let data = response.data;
  if (data.state === 'error') {
    let result = data.result;
    return Promise.reject(new AppError(result.code, result.message));
  }
  return data.result;
})

new Vue({
  el: '#app',
  data: {
    isLoginButtonDisabled: true,
    user: null,
  },
  created: function() {
    // Is MetaMask installed?
    if (typeof web3 === 'undefined') {
      this.warn('Please install MetaMask');
      return;
    }

    // LocalStorage is available?
    try {
      window.localStorage;
    }
    catch(e) {
      this.warn('Please allow 3rd party cookies for web3.js 1.0.0');
      return;
    }

    window.web3 = new Web3(web3.currentProvider);
    window.web3.extend({
      property: 'app',
      methods: [{
        name: 'signTypedData',
        call: 'eth_signTypedData',
        params: 2,
      }],
    });

    this.isLoginButtonDisabled = false;
  },
  methods: {
    login: function() {
      let $this = this;

      let accounts = [];

      web3.eth.getAccounts()
        .then((result) => {
          // Are there available accounts?
          if (result.length <= 0) {
            $this.throw('Please unlock MetaMask account');
          }

          accounts = result;

          return web3.eth.net.getId();
        })
        .then((result) => {
          // Does MetaMask connect to Ropeten?
          if (result !== 3) {
            $this.throw('Please connect MetaMask to Ropsten Test Network');
          }

          let params = new URLSearchParams();
          params.append('address', accounts[0]);

          return client.post('/auth/challenge', params);
        })
        .then((result) => {
          let typedData = [{
            type: 'string',
            name: 'challenge',
            value: result.challenge,
          }];

          return web3.app.signTypedData(typedData, accounts[0]);
        })
        .then((result) => {
          let params = new URLSearchParams();
          params.append('address', accounts[0]);
          params.append('signature', result);

          return client.post('/auth/authorize', params);
        })
        .then((result) => {
          client.defaults.headers.common['Authorization'] = 'Bearer ' + result.token;

          return client.get('/api/users/' + accounts[0]);
        })
        .then((result) => {
          $this.user = result;
        })
        .catch((e) => {
          $this.handleError(e);
        });
    },
    updateUser: function() {
      let $this = this;

      let params = new URLSearchParams();
      params.append('name', $this.user.name);

      client.put('/api/users/' + $this.user.address, params)
        .then((result) => {
          $this.info('success');
        })
        .catch((e) => {
          $this.handleError(e);
        });
    },
    deleteUser: function() {
      let $this = this;

      $this.$confirm('Are you sure to delete the account?')
        .then(_ => {
          return client.delete('/api/users/' + $this.user.address)
        })
        .then((result) => {
          $this.logout();
        })
        .catch((e) => {
          $this.handleError(e);
        });
    },
    throw: function(msg) {
      throw new AppError(0, msg);
    },
    handleError: function(e) {
      if (typeof e === 'string') {
        if (e !== 'cancel') {
          throw e;
        }
      }
      else if (e instanceof AppError) {
        if (e.code > 0) {
          this.warn(e.message + ' [' + e.code + ']');
        } else {
          this.warn(e.message);
        }
      }
      else if (e.message.match(/User denied message signature\./)) {
        this.warn('Please accept the signature request');
      }
      else {
        throw e;
      }
    },
    info: function(msg) {
      this.$message({
        showClose: true,
        message: msg,
        type: 'info',
      });
    },
    warn: function(msg) {
      this.$message({
        showClose: true,
        message: msg,
        type: 'warning',
      });
    },
    logout: function() {
      delete client.defaults.headers.common['Authorization'];
      this.user = null;
    },
  },
});
