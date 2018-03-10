import 'babel-polyfill'
import Vue from 'vue/dist/vue.esm.js'
import Web3 from 'web3'

function AppError(message) {
  this.message = message
}
Object.setPrototypeOf(AppError, Error)
AppError.prototype = Object.create(Error.prototype)
AppError.prototype.name = "AppError"
AppError.prototype.message = ""
AppError.prototype.constructor = AppError

new Vue({
  el: '#app',
  data: {
    errorMessage: null,
  },
  methods: {
    login: function() {
      let app = this

      // Is MetaMask installed?
      if (typeof web3 === 'undefined') {
        app.errorMessage = 'Please install MetaMask'
        return
      }

      web3 = new Web3(web3.currentProvider)
      web3.extend({
        property: 'app',
        methods: [{
          name: 'signTypedData',
          call: 'eth_signTypedData',
          params: 2,
        }],
      })

      let accounts = []

      web3.eth.getAccounts()
        .then((result) => {
          // Are there available accounts?
          if (result.length <= 0) {
            throw new AppError('Please unlock MetaMask account')
          }

          accounts = result

          return web3.eth.net.getId()
        })
        .then((network) => {
          // Does MetaMask connect to Ropeten?
          if (network !== 3) {
            throw new AppError('Please connect MetaMask to Ropsten Test Network')
          }

          // Clean up error message
          app.errorMessage = null

          return web3.app.signTypedData([
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
          ], accounts[0])
        })
        .then((signature) => {
          console.log(signature) // TODO
        })
        .catch((e) => {
          if (e instanceof AppError) {
            app.errorMessage = e.message
          } else if (e.message.match(/User denied message signature\./)) {
            app.errorMessage = 'User denied message signature'
          } else {
            throw e
          }
        })
    },
  },
})
