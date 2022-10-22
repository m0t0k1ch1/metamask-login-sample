"use strict";

import "babel-polyfill";

import Vue from "vue";
import ElementUI from "element-ui";
import locale from "element-ui/lib/locale/lang/en";
import Web3 from "web3";

import AppClient from "./client.js";
import AppError from "./error.js";

import "../css/reset.css";
import "element-ui/lib/theme-chalk/index.css";
import "../css/style.css";

Vue.use(ElementUI, { locale });

const appClient = new AppClient();

new Vue({
  el: "#app",
  data: {
    isLoginButtonDisabled: true,
    isObservationEnabled: false,
    user: null,
  },
  created: async function () {
    try {
      await initWeb3();
      this.isLoginButtonDisabled = false;
      this.observe();
    } catch (e) {
      this.handleError(e);
    }
  },
  methods: {
    observe: async function () {
      try {
        setTimeout(this.observe, 1000);

        if (!this.isObservationEnabled) {
          return;
        }

        let address = await getAddress();
        if (address === this.user.address) {
          return;
        } else {
          this.$alert("Account has changed");
          this.logout();
        }
      } catch (e) {
        this.handleError(e);
      }
    },
    login: async function () {
      try {
        if (window.ethereum) {
          await ethereum.enable();
        }

        let address = await getAddress();
        if (address === null) {
          throw new AppError("Please unlock MetaMask account");
        }

        let networkId = await getNetworkId();
        if (networkId !== 5) {
          throw new AppError("Please connect MetaMask to Goerli Test Network");
        }

        let challengeResult = await appClient.challenge(address);
        let signature = await signTypedData(address, challengeResult.challenge);
        let authorizeResult = await appClient.authorize(address, signature);

        appClient.setToken(authorizeResult.token);

        this.user = await appClient.getUser(address);
        this.isObservationEnabled = true;
      } catch (e) {
        this.handleError(e);
      }
    },
    updateUser: async function () {
      try {
        await appClient.updateUser(this.user.address, {
          name: this.user.name,
        });
        this.info("Success");
      } catch (e) {
        this.handleError(e);
      }
    },
    deleteUser: async function () {
      try {
        await this.$confirm("Are you sure to delete the account?");
        await appClient.deleteUser(this.user.address);
        this.logout();
        this.info("Success");
      } catch (e) {
        this.handleError(e);
      }
    },
    info: function (msg) {
      this.$message({
        showClose: true,
        message: msg,
        type: "info",
      });
    },
    warn: function (msg) {
      this.$message({
        showClose: true,
        message: msg,
        type: "warning",
      });
    },
    handleError: function (e) {
      if (typeof e === "string") {
        if (e !== "cancel") {
          throw e;
        }
      } else if (e instanceof AppError) {
        if (e.code > 0) {
          this.warn(e.message + " [" + e.code + "]");
        } else {
          this.warn(e.message);
        }
      } else if (e.message.match(/User denied message signature\./)) {
        this.warn("Please accept the signature request");
      } else {
        throw e;
      }
    },
    logout: function () {
      appClient.initToken();

      this.isObservationEnabled = false;
      this.user = null;
    },
  },
});

function initWeb3() {
  return new Promise((resolve, reject) => {
    if (window.ethereum) {
      window.web3 = new Web3(ethereum);
    } else if (window.web3) {
      window.web3 = new Web3(web3.currentProvider);
    } else {
      return reject(new AppError("Please install MetaMask"));
    }

    try {
      window.localStorage;
    } catch (e) {
      return reject(
        new AppError("Please allow 3rd party cookies for web3.js 1.0.0")
      );
    }

    window.web3.defaultChain = "goerli";

    window.web3.extend({
      property: "app",
      methods: [
        {
          name: "signTypedData",
          call: "eth_signTypedData",
          params: 2,
        },
      ],
    });

    resolve();
  });
}

function getAddress() {
  return new Promise((resolve, reject) => {
    web3.eth.getAccounts().then((accounts) => {
      if (accounts.length <= 0) {
        return resolve(null);
      }
      resolve(accounts[0]);
    });
  });
}

function getNetworkId() {
  return web3.eth.net.getId();
}

function signTypedData(address, value) {
  return web3.app.signTypedData(
    [
      {
        type: "string",
        name: "challenge",
        value: value,
      },
    ],
    address
  );
}
