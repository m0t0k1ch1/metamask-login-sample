import 'babel-polyfill'
import Vue from 'vue/dist/vue.esm.js'
import Eth from 'ethjs'

let eth

new Vue({
  el: '#app',
  data: {
    errorMessage: null,
  },
  methods: {
    login: function() {
      let app = this

      if (typeof web3 === 'undefined') {
        app.errorMessage = 'Please install MetaMask'
        return
      }

      eth = new Eth(web3.currentProvider)

      eth.accounts()
        .then((accounts) => {
          if (accounts.length <= 0) {
            app.errorMessage = 'Please unlock MetaMask account'
            throw app.errorMessage
          }
          return eth.net_version()
        })
        .then((netVersion) => {
          if (netVersion !== '3') {
            app.errorMessage = 'Please connect MetaMask to Ropsten Test Network'
            throw app.errorMessage
          }
          app.errorMessage = null
          alert('ok') // TODO
        })
        .catch((e) => {
          throw e
        })
    },
  },
})
