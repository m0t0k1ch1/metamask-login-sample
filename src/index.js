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

let app = new Vue({
  el: '#app',
  data: {
    isInitialized: false,
    errorMessage: null,
  },
  methods: {
    login: function() {
      let $this = this

      if (!$this.isInitialized) {
        throw new AppError('Initialization has not completed yet')
      }

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
        .then((result) => {
          // Does MetaMask connect to Ropeten?
          if (result !== 3) {
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
          ], accounts[0]) // TODO
        })
        .then((result) => {
          console.log(result) // TODO
        })
        .catch((e) => {
          if (e instanceof AppError) {
            app.errorMessage = e.message
          } else if (e.message.match(/User denied message signature\./)) {
            app.errorMessage = 'Request is cancelled'
          } else {
            throw e
          }
        })
    },
  },
})

window.addEventListener('load', () => {
  // Is MetaMask installed?
  if (typeof web3 === 'undefined') {
    app.errorMessage = 'Please install MetaMask'
  }

  window.web3 = new Web3(web3.currentProvider)
  window.web3.extend({
    property: 'app',
    methods: [{
      name: 'signTypedData',
      call: 'eth_signTypedData',
      params: 2,
    }],
  })

  app.isInitialized = true
})
