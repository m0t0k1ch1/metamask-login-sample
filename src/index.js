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

          eth.net_version().then((netVersion) => {
            if (netVersion !== '3') {
              app.errorMessage = 'Please connect MetaMask to Ropsten Test Network'
              throw app.errorMessage
            }

            // clean up error message
            app.errorMessage = null

            const msgParams = [
              {
                type: 'string',
                name: 'message',
                value: 'Hi, Alice!',
              },
              {
                type: 'uint32',
                name: 'value',
                value: 42,
              },
            ]

            web3.currentProvider.sendAsync({
              method: 'eth_signTypedData',
              params: [msgParams, accounts[0]],
              from: accounts[0],
            }, (err, data) => {
              if (err) throw err
              if (data.error) {
                if (data.error.code === -32603) {
                  console.log('User denied message signature')
                  return
                }
                else {
                  throw data.error
                }
              }

              let sig = data.result

              console.log(sig) // TODO
            })
          })
        })
    },
  },
})
