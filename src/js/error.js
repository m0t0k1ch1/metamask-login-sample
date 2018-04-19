'use strict';

let AppError = function(code, message) {
  this.code = code;
  this.message = message;
};
AppError.prototype.name = "AppError";

module.exports = AppError;
